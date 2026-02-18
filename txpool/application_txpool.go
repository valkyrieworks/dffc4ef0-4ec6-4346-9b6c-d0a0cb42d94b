package txpool

import (
	"context"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/kinds"
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
	stats *Stats
	app     ApplicationTxpoolCustomer
	viewed    TransferRepository
	tracer  log.Tracer
}

//
type ApplicationTxpoolCustomer interface {
	//
	EmbedTransfer(ctx context.Context, req *iface.QueryEmbedTransfer) (*iface.ReplyEmbedTransfer, error)

	//
	HarvestTrans(ctx context.Context, req *iface.QueryHarvestTrans) (*iface.ReplyHarvestTrans, error)

	//
	Purge(context.Context) error
}

//
type ApplicationTxpoolOption func(*ApplicationTxpool)

//
const (
	viewedRepositoryVolume = 100_000
	harvestMaximumOctets  = 0
	harvestMaximumFuel    = 0
	harvestCadence  = 500 * time.Millisecond
)

var _ Txpool = &ApplicationTxpool{}

var (
	ErrNegateExecuted = errors.New("REDACTED")
	ErrEmptyTransfer        = errors.New("REDACTED")
	ErrViewedTransfer         = errors.New("REDACTED")
)

func WithMorningStats(stats *Stats) ApplicationTxpoolOption {
	return func(m *ApplicationTxpool) { m.stats = stats }
}

func WithMorningTracer(tracer log.Tracer) ApplicationTxpoolOption {
	return func(m *ApplicationTxpool) { m.tracer = tracer }
}

func NewApplicationTxpool(
	settings *settings.TxpoolSettings,
	app ApplicationTxpoolCustomer,
	opts ...ApplicationTxpoolOption,
) *ApplicationTxpool {
	//
	//
	viewed := NewLRUTransferRepository(viewedRepositoryVolume)

	m := &ApplicationTxpool{
		ctx:     context.Background(),
		settings:  settings,
		app:     app,
		viewed:    viewed,
		stats: NoopStats(),
		tracer:  log.NewNoopTracer(),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

//
//
func (m *ApplicationTxpool) EmbedTransfer(tx kinds.Tx) error {
	if err := m.shieldTransfer(tx); err != nil {
		return err
	}

	code, err := m.embedTransfer(tx)

	//
	switch {
	case err != nil:
		m.stats.ErroredTrans.Add(1)
		return encloseErrCode("REDACTED", code, err)
	case codeReprocess(code):
		//
		m.viewed.Delete(tx)
		fallthrough
	case code != iface.CodeKindSuccess:
		m.stats.DeclinedTrans.Add(1)
		return encloseErrCode("REDACTED", code, err)
	default:
		m.stats.TransferVolumeOctets.Observe(float64(len(tx)))
		return nil
	}
}

//
func (m *ApplicationTxpool) shieldTransfer(tx kinds.Tx) error {
	transferVolume := len(tx)

	if transferVolume == 0 {
		return ErrEmptyTransfer
	}

	if m.settings.MaximumTransferOctets > 0 && transferVolume > m.settings.MaximumTransferOctets {
		return &ErrTransferTooBulky{
			Max:    m.settings.MaximumTransferOctets,
			Factual: transferVolume,
		}
	}

	impelled := m.viewed.Propel(tx)
	if !impelled {
		m.stats.YetAcceptedTrans.Add(1)
		return ErrViewedTransfer
	}

	return nil
}

func (m *ApplicationTxpool) embedTransfer(tx kinds.Tx) (uint32, error) {
	//
	reply, err := m.app.EmbedTransfer(m.ctx, &iface.QueryEmbedTransfer{Tx: tx})
	if err != nil {
		if reply != nil {
			return reply.Code, err
		}
		return 0, err
	}

	return reply.Code, nil
}

//
//
//
func (m *ApplicationTxpool) TransferFlow(ctx context.Context) <-chan kinds.Txs {
	ch := make(chan kinds.Txs, 1)

	go func() {
		defer func() {
			close(ch)

			if p := recover(); p != nil {
				m.tracer.Fault("REDACTED", "REDACTED", p)
			}
		}()

		m.harvestTrans(ctx, ch)
	}()

	return ch
}

func (m *ApplicationTxpool) harvestTrans(ctx context.Context, conduit chan<- kinds.Txs) {
	req := &iface.QueryHarvestTrans{
		MaximumOctets: harvestMaximumOctets,
		MaximumFuel:   harvestMaximumFuel,
	}

	for {
		select {
		case <-ctx.Done():
			m.tracer.Diagnose("REDACTED")
			return
		case <-time.After(harvestCadence):
			//
			res, err := m.app.HarvestTrans(ctx, req)
			switch {
			case isErrCtx(err):
				m.tracer.Diagnose("REDACTED")
				return
			case err != nil:
				m.tracer.Fault("REDACTED", "REDACTED", err)
				continue
			case len(res.Txs) == 0:
				//
				continue
			}

			txs := kinds.ToTrans(res.Txs)
			m.stats.HarvestedTrans.Add(float64(len(txs)))

			select {
			case <-ctx.Done():
				m.tracer.Diagnose("REDACTED")
				return
			case conduit <- txs:
				//
			}

			//
			for _, tx := range txs {
				if m.viewed.Has(tx) {
					continue
				}

				m.viewed.Propel(tx)
			}
		}
	}
}

//
func (m *ApplicationTxpool) PurgeApplicationLink() error {
	err := m.app.Purge(m.ctx)
	if err != nil {
		return ErrPurgeApplicationLink{Err: err}
	}

	return nil
}

//
//
//
func (m *ApplicationTxpool) InspectTransfer(tx kinds.Tx, callback func(res *iface.ReplyInspectTransfer), _ TransferDetails) error {
	if err := m.shieldTransfer(tx); err != nil {
		return err
	}

	go func() {
		defer func() {
			if p := recover(); p != nil {
				m.tracer.Fault("REDACTED", "REDACTED", p, "REDACTED", transferDigest(tx))
			}
		}()

		code, err := m.embedTransfer(tx)
		if err != nil {
			//
			m.tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", transferDigest(tx))
			return
		}

		//
		//
		//
		if callback != nil {
			callback(&iface.ReplyInspectTransfer{
				Code:      code,
				Data:      []byte{},
				Log:       "REDACTED",
				Details:      "REDACTED",
				FuelDesired: 0,
				FuelApplied:   0,
				Events:    []iface.Event{},
				Codex: "REDACTED",
			})
		}
	}()

	return nil
}

//
func (m *ApplicationTxpool) Modify(_ int64, _ kinds.Txs, _ []*iface.InvokeTransferOutcome, _ PreInspectFunction, _ SubmitInspectFunction) error {
	return nil
}

//
func (m *ApplicationTxpool) TransAccessible() <-chan struct{} { return nil }
func (m *ApplicationTxpool) ActivateTransAccessible()           {}

func (m *ApplicationTxpool) Volume() int        { return 0 }
func (m *ApplicationTxpool) VolumeOctets() int64 { return 0 }

func (m *ApplicationTxpool) HarvestMaximumOctetsMaximumFuel(_, _ int64) kinds.Txs { return nil }
func (m *ApplicationTxpool) HarvestMaximumTrans(_ int) kinds.Txs              { return nil }
func (m *ApplicationTxpool) DeleteTransferByKey(_ kinds.TransferKey) error       { return nil }
func (m *ApplicationTxpool) Purge()                                  {}

func (m *ApplicationTxpool) Secure()   {}
func (m *ApplicationTxpool) Release() {}

func isErrCtx(err error) bool {
	if err == nil {
		return false
	}

	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}

func codeReprocess(code uint32) bool {
	return code >= iface.CodeKindReprocess
}

func encloseErrCode(msg string, code uint32, err error) error {
	if err == nil {
		return fmt.Errorf("REDACTED", msg, code)
	}

	return errors.Wrapf(err, "REDACTED", msg, code)
}
