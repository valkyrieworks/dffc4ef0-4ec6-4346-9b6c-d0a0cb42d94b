package txpool

import (
	"context"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/pkg/errors"
)

//
//
//
//
//
//
type ApplicationTxpool struct {
	ctx     context.Context
	settings  *settings.TxpoolSettings
	telemetry *Telemetry
	app     ApplicationTxpoolCustomer
	observed    TransferStash
	tracer  log.Tracer
}

//
type ApplicationTxpoolCustomer interface {
	//
	AppendTransfer(ctx context.Context, req *iface.SolicitAppendTransfer) (*iface.ReplyAppendTransfer, error)

	//
	HarvestTrans(ctx context.Context, req *iface.SolicitHarvestTrans) (*iface.ReplyHarvestTrans, error)

	//
	Purge(context.Context) error
}

//
type ApplicationTxpoolChoice func(*ApplicationTxpool)

//
const (
	observedStashExtent = 100_000
	harvestMaximumOctets  = 0
	harvestMaximumFuel    = 0
	harvestDuration  = 500 * time.Millisecond
)

var _ Txpool = &ApplicationTxpool{}

var (
	FaultNegationExecuted = errors.New("REDACTED")
	FaultBlankTransfer        = errors.New("REDACTED")
	FaultObservedTransfer         = errors.New("REDACTED")
)

func UsingMorningTelemetry(telemetry *Telemetry) ApplicationTxpoolChoice {
	return func(m *ApplicationTxpool) { m.telemetry = telemetry }
}

func UsingMorningTracer(tracer log.Tracer) ApplicationTxpoolChoice {
	return func(m *ApplicationTxpool) { m.tracer = tracer }
}

func FreshApplicationTxpool(
	settings *settings.TxpoolSettings,
	app ApplicationTxpoolCustomer,
	choices ...ApplicationTxpoolChoice,
) *ApplicationTxpool {
	//
	//
	observed := FreshLeastusedTransferStash(observedStashExtent)

	m := &ApplicationTxpool{
		ctx:     context.Background(),
		settings:  settings,
		app:     app,
		observed:    observed,
		telemetry: NooperationTelemetry(),
		tracer:  log.FreshNooperationTracer(),
	}

	for _, opt := range choices {
		opt(m)
	}

	return m
}

//
//
func (m *ApplicationTxpool) AppendTransfer(tx kinds.Tx) error {
	if err := m.shieldTransfer(tx); err != nil {
		return err
	}

	cipher, err := m.appendTransfer(tx)

	//
	switch {
	case err != nil:
		m.telemetry.UnsuccessfulTrans.Add(1)
		return encloseFaultCipher("REDACTED", cipher, err)
	case cipherReissue(cipher):
		//
		m.observed.Discard(tx)
		fallthrough
	case cipher != iface.CipherKindOKAY:
		m.telemetry.DeclinedTrans.Add(1)
		return encloseFaultCipher("REDACTED", cipher, err)
	default:
		m.telemetry.TransferExtentOctets.Observe(float64(len(tx)))
		return nil
	}
}

//
func (m *ApplicationTxpool) shieldTransfer(tx kinds.Tx) error {
	transferExtent := len(tx)

	if transferExtent == 0 {
		return FaultBlankTransfer
	}

	if m.settings.MaximumTransferOctets > 0 && transferExtent > m.settings.MaximumTransferOctets {
		return &FaultTransferExcessivelyAmple{
			Max:    m.settings.MaximumTransferOctets,
			Existing: transferExtent,
		}
	}

	propelled := m.observed.Propel(tx)
	if !propelled {
		m.telemetry.EarlierAcceptedTrans.Add(1)
		return FaultObservedTransfer
	}

	return nil
}

func (m *ApplicationTxpool) appendTransfer(tx kinds.Tx) (uint32, error) {
	//
	reply, err := m.app.AppendTransfer(m.ctx, &iface.SolicitAppendTransfer{Tx: tx})
	if err != nil {
		if reply != nil {
			return reply.Cipher, err
		}
		return 0, err
	}

	return reply.Cipher, nil
}

//
//
//
func (m *ApplicationTxpool) TransferInflux(ctx context.Context) <-chan kinds.Txs {
	ch := make(chan kinds.Txs, 1)

	go func() {
		defer func() {
			close(ch)

			if p := recover(); p != nil {
				m.tracer.Failure("REDACTED", "REDACTED", p)
			}
		}()

		m.harvestTrans(ctx, ch)
	}()

	return ch
}

func (m *ApplicationTxpool) harvestTrans(ctx context.Context, conduit chan<- kinds.Txs) {
	req := &iface.SolicitHarvestTrans{
		MaximumOctets: harvestMaximumOctets,
		MaximumFuel:   harvestMaximumFuel,
	}

	for {
		select {
		case <-ctx.Done():
			m.tracer.Diagnose("REDACTED")
			return
		case <-time.After(harvestDuration):
			//
			res, err := m.app.HarvestTrans(ctx, req)
			switch {
			case equalsFaultContext(err):
				m.tracer.Diagnose("REDACTED")
				return
			case err != nil:
				m.tracer.Failure("REDACTED", "REDACTED", err)
				continue
			case len(res.Txs) == 0:
				//
				continue
			}

			txs := kinds.TowardTrans(res.Txs)
			m.telemetry.HarvestedTrans.Add(float64(len(txs)))

			select {
			case <-ctx.Done():
				m.tracer.Diagnose("REDACTED")
				return
			case conduit <- txs:
				//
			}

			//
			for _, tx := range txs {
				if m.observed.Has(tx) {
					continue
				}

				m.observed.Propel(tx)
			}
		}
	}
}

//
func (m *ApplicationTxpool) PurgeApplicationLink() error {
	err := m.app.Purge(m.ctx)
	if err != nil {
		return FaultPurgeApplicationLink{Err: err}
	}

	return nil
}

//
//
//
func (m *ApplicationTxpool) InspectTransfer(tx kinds.Tx, clbk func(res *iface.ReplyInspectTransfer), _ TransferDetails) error {
	if err := m.shieldTransfer(tx); err != nil {
		return err
	}

	go func() {
		defer func() {
			if p := recover(); p != nil {
				m.tracer.Failure("REDACTED", "REDACTED", p, "REDACTED", transferDigest(tx))
			}
		}()

		cipher, err := m.appendTransfer(tx)
		if err != nil {
			//
			m.tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", transferDigest(tx))
			return
		}

		//
		//
		//
		if clbk != nil {
			clbk(&iface.ReplyInspectTransfer{
				Cipher:      cipher,
				Data:      []byte{},
				Log:       "REDACTED",
				Details:      "REDACTED",
				FuelDesired: 0,
				FuelUtilized:   0,
				Incidents:    []iface.Incident{},
				Codeset: "REDACTED",
			})
		}
	}()

	return nil
}

//
func (m *ApplicationTxpool) Revise(_ int64, _ kinds.Txs, _ []*iface.InvokeTransferOutcome, _ PriorInspectMethod, _ RelayInspectMethod) error {
	return nil
}

//
func (m *ApplicationTxpool) TransAccessible() <-chan struct{} { return nil }
func (m *ApplicationTxpool) ActivateTransAccessible()           {}

func (m *ApplicationTxpool) Extent() int        { return 0 }
func (m *ApplicationTxpool) ExtentOctets() int64 { return 0 }

func (m *ApplicationTxpool) HarvestMaximumOctetsMaximumFuel(_, _ int64) kinds.Txs { return nil }
func (m *ApplicationTxpool) HarvestMaximumTrans(_ int) kinds.Txs              { return nil }
func (m *ApplicationTxpool) DiscardTransferViaToken(_ kinds.TransferToken) error       { return nil }
func (m *ApplicationTxpool) Purge()                                  {}

func (m *ApplicationTxpool) Secure()   {}
func (m *ApplicationTxpool) Release() {}

func equalsFaultContext(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

func cipherReissue(cipher uint32) bool {
	return cipher >= iface.CipherKindReissue
}

func encloseFaultCipher(msg string, cipher uint32, err error) error {
	if err == nil {
		return fmt.Errorf("REDACTED", msg, cipher)
	}

	return errors.Wrapf(err, "REDACTED", msg, cipher)
}
