package agreement

import (
	"bytes"
	"context"
	"fmt"
	"hash/crc32"
	"io"
	"reflect"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
)

var crc32c = crc32.MakeTable(crc32.Castagnoli)

//
//
//
//
//
//
//
//

//
//
//
//

//
//
//
func (cs *Status) scanResimulateSignal(msg *ScheduledJournalSignal, newPhaseSubtract kinds.Enrollment) error {
	//
	if _, ok := msg.Msg.(TerminateLevelSignal); ok {
		return nil
	}

	//
	switch m := msg.Msg.(type) {
	case kinds.EventDataDurationStatus:
		cs.Tracer.Details("REDACTED", "REDACTED", m.Level, "REDACTED", m.Cycle, "REDACTED", m.Phase)
		//
		timer := time.After(time.Second * 2)
		if newPhaseSubtract != nil {
			select {
			case phaseMessage := <-newPhaseSubtract.Out():
				m2 := phaseMessage.Data().(kinds.EventDataDurationStatus)
				if m.Level != m2.Level || m.Cycle != m2.Cycle || m.Phase != m2.Phase {
					return fmt.Errorf("REDACTED", m2, m)
				}
			case <-newPhaseSubtract.Revoked():
				return fmt.Errorf("REDACTED")
			case <-timer:
				return fmt.Errorf("REDACTED")
			}
		}
	case messageDetails:
		nodeUID := m.NodeUID
		if nodeUID == "REDACTED" {
			nodeUID = "REDACTED"
		}
		switch msg := m.Msg.(type) {
		case *NominationSignal:
			p := msg.Nomination
			cs.Tracer.Details("REDACTED", "REDACTED", p.Level, "REDACTED", p.Cycle, "REDACTED",
				p.LedgerUID.SegmentAssignHeading, "REDACTED", p.POLDuration, "REDACTED", nodeUID)
		case *LedgerSegmentSignal:
			cs.Tracer.Details("REDACTED", "REDACTED", msg.Level, "REDACTED", msg.Cycle, "REDACTED", nodeUID)
		case *BallotSignal:
			v := msg.Ballot
			cs.Tracer.Details("REDACTED", "REDACTED", v.Level, "REDACTED", v.Cycle, "REDACTED", v.Kind,
				"REDACTED", v.LedgerUID, "REDACTED", nodeUID, "REDACTED", len(v.Addition), "REDACTED", len(v.AdditionAutograph))
		}

		cs.processMessage(m)
	case deadlineDetails:
		cs.Tracer.Details("REDACTED", "REDACTED", m.Level, "REDACTED", m.Cycle, "REDACTED", m.Phase, "REDACTED", m.Period)
		cs.processDeadline(m, cs.DurationStatus)
	default:
		return fmt.Errorf("REDACTED", reflect.TypeOf(msg.Msg))
	}
	return nil
}

//
//
func (cs *Status) overtakeResimulate(csLevel int64) error {
	//
	cs.resimulateStyle = true
	defer func() { cs.resimulateStyle = false }()

	//
	//
	//
	//
	//
	//
	gr, located, err := cs.wal.ScanForTerminateLevel(csLevel, &JournalScanSettings{BypassDataImpairmentFaults: true})
	if err != nil {
		return err
	}
	if gr != nil {
		if err := gr.Close(); err != nil {
			return err
		}
	}
	if located {
		return fmt.Errorf("REDACTED", csLevel)
	}

	//
	//
	//
	if csLevel < cs.status.PrimaryLevel {
		return fmt.Errorf("REDACTED", csLevel, cs.status.PrimaryLevel)
	}
	terminateLevel := csLevel - 1
	if csLevel == cs.status.PrimaryLevel {
		terminateLevel = 0
	}
	gr, located, err = cs.wal.ScanForTerminateLevel(terminateLevel, &JournalScanSettings{BypassDataImpairmentFaults: true})
	if err == io.EOF {
		cs.Tracer.Fault("REDACTED", "REDACTED", terminateLevel)
	} else if err != nil {
		return err
	}
	if !located {
		return fmt.Errorf("REDACTED", csLevel, terminateLevel)
	}
	defer gr.Close()

	cs.Tracer.Details("REDACTED", "REDACTED", csLevel)

	var msg *ScheduledJournalSignal
	dec := JournalParser{gr}

Cycle:
	for {
		msg, err = dec.Parse()
		switch {
		case err == io.EOF:
			break Cycle
		case IsDataImpairmentFault(err):
			cs.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", csLevel)
			return err
		case err != nil:
			return err
		}

		//
		//
		//
		if err := cs.scanResimulateSignal(msg, nil); err != nil {
			return err
		}
	}
	cs.Tracer.Details("REDACTED")
	return nil
}

//

//
//
/**
{
{
)
)
{
)
}
)
{
)
}
{
l
{
l
{
l
}
}
*/

//
//
//
//
//

type Greeter struct {
	statusDepot   sm.Depot
	primaryStatus sm.Status
	depot        sm.LedgerDepot
	eventBus     kinds.LedgerEventBroadcaster
	generatePaper       *kinds.OriginPaper
	tracer       log.Tracer

	nLedgers int //
}

func NewGreeter(statusDepot sm.Depot, status sm.Status,
	depot sm.LedgerDepot, generatePaper *kinds.OriginPaper,
) *Greeter {
	return &Greeter{
		statusDepot:   statusDepot,
		primaryStatus: status,
		depot:        depot,
		eventBus:     kinds.NoopEventBus{},
		generatePaper:       generatePaper,
		tracer:       log.NewNoopTracer(),
		nLedgers:      0,
	}
}

func (h *Greeter) AssignTracer(l log.Tracer) {
	h.tracer = l
}

//
//
func (h *Greeter) AssignEventBus(eventBus kinds.LedgerEventBroadcaster) {
	h.eventBus = eventBus
}

//
func (h *Greeter) NLedgers() int {
	return h.nLedgers
}

//
func (h *Greeter) Greeting(gatewayApplication gateway.ApplicationLinks) error {
	return h.GreetingWithContext(context.TODO(), gatewayApplication)
}

//
func (h *Greeter) GreetingWithContext(ctx context.Context, gatewayApplication gateway.ApplicationLinks) error {
	//
	res, err := gatewayApplication.Inquire().Details(ctx, gateway.QueryDetails)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	ledgerLevel := res.FinalLedgerLevel
	if ledgerLevel < 0 {
		return fmt.Errorf("REDACTED", ledgerLevel)
	}
	applicationDigest := res.FinalLedgerApplicationDigest

	h.tracer.Details("REDACTED",
		"REDACTED", ledgerLevel,
		"REDACTED", log.NewIdleFormat("REDACTED", applicationDigest),
		"REDACTED", res.Release,
		"REDACTED", res.ApplicationRelease,
	)

	//
	if h.primaryStatus.FinalLedgerLevel == 0 {
		h.primaryStatus.Release.Agreement.App = res.ApplicationRelease
	}

	//
	applicationDigest, err = h.ResimulateLedgersWithContext(ctx, h.primaryStatus, applicationDigest, ledgerLevel, gatewayApplication)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	h.tracer.Details("REDACTED",
		"REDACTED", ledgerLevel, "REDACTED", log.NewIdleFormat("REDACTED", applicationDigest))

	//

	return nil
}

//
//
//
func (h *Greeter) ResimulateLedgers(
	status sm.Status,
	applicationDigest []byte,
	applicationLedgerLevel int64,
	gatewayApplication gateway.ApplicationLinks,
) ([]byte, error) {
	return h.ResimulateLedgersWithContext(context.TODO(), status, applicationDigest, applicationLedgerLevel, gatewayApplication)
}

//
func (h *Greeter) ResimulateLedgersWithContext(
	ctx context.Context,
	status sm.Status,
	applicationDigest []byte,
	applicationLedgerLevel int64,
	gatewayApplication gateway.ApplicationLinks,
) ([]byte, error) {
	depotLedgerRoot := h.depot.Root()
	depotLedgerLevel := h.depot.Level()
	statusLedgerLevel := status.FinalLedgerLevel
	h.tracer.Details(
		"REDACTED",
		"REDACTED",
		applicationLedgerLevel,
		"REDACTED",
		depotLedgerLevel,
		"REDACTED",
		statusLedgerLevel)

	//
	if applicationLedgerLevel == 0 {
		ratifiers := make([]*kinds.Ratifier, len(h.generatePaper.Ratifiers))
		for i, val := range h.generatePaper.Ratifiers {
			ratifiers[i] = kinds.NewRatifier(val.PublicKey, val.Energy)
		}
		ratifierCollection := kinds.NewRatifierCollection(ratifiers)
		followingValues := kinds.Tm2schema.RatifierRefreshes(ratifierCollection)
		schemasettings := h.generatePaper.AgreementOptions.ToSchema()
		req := &iface.QueryInitSeries{
			Time:            h.generatePaper.OriginMoment,
			SeriesUid:         h.generatePaper.LedgerUID,
			PrimaryLevel:   h.generatePaper.PrimaryLevel,
			AgreementOptions: &schemasettings,
			Ratifiers:      followingValues,
			ApplicationStatusOctets:   h.generatePaper.ApplicationStatus,
		}
		res, err := gatewayApplication.Agreement().InitSeries(context.TODO(), req)
		if err != nil {
			return nil, err
		}

		applicationDigest = res.ApplicationDigest

		if statusLedgerLevel == 0 { //
			//
			//
			//
			if len(res.ApplicationDigest) > 0 {
				status.ApplicationDigest = res.ApplicationDigest
			}
			//
			if len(res.Ratifiers) > 0 {
				values, err := kinds.Schema2tm.RatifierRefreshes(res.Ratifiers)
				if err != nil {
					return nil, err
				}
				status.Ratifiers = kinds.NewRatifierCollection(values)
				status.FollowingRatifiers = kinds.NewRatifierCollection(values).CloneAugmentRecommenderUrgency(1)
			} else if len(h.generatePaper.Ratifiers) == 0 {
				//
				return nil, fmt.Errorf("REDACTED")
			}

			if res.AgreementOptions != nil {
				status.AgreementOptions = status.AgreementOptions.Modify(res.AgreementOptions)
				status.Release.Agreement.App = status.AgreementOptions.Release.App
			}
			//
			status.FinalOutcomesDigest = merkle.DigestFromOctetSegments(nil)
			if err := h.statusDepot.Persist(status); err != nil {
				return nil, err
			}
		}
	}

	//
	switch {
	case depotLedgerLevel == 0:
		affirmApplicationDigestMatchesOneFromStatus(applicationDigest, status)
		return applicationDigest, nil

	case applicationLedgerLevel == 0 && status.PrimaryLevel < depotLedgerRoot:
		//
		return applicationDigest, sm.ErrApplicationLedgerLevelTooInferior{ApplicationLevel: applicationLedgerLevel, DepotRoot: depotLedgerRoot}

	case applicationLedgerLevel > 0 && applicationLedgerLevel < depotLedgerRoot-1:
		//
		return applicationDigest, sm.ErrApplicationLedgerLevelTooInferior{ApplicationLevel: applicationLedgerLevel, DepotRoot: depotLedgerRoot}

	case depotLedgerLevel < applicationLedgerLevel:
		//
		return applicationDigest, sm.ErrApplicationLedgerLevelTooSuperior{CoreLevel: depotLedgerLevel, ApplicationLevel: applicationLedgerLevel}

	case depotLedgerLevel < statusLedgerLevel:
		//
		panic(fmt.Sprintf("REDACTED", statusLedgerLevel, depotLedgerLevel))

	case depotLedgerLevel > statusLedgerLevel+1:
		//
		panic(fmt.Sprintf("REDACTED", depotLedgerLevel, statusLedgerLevel+1))
	}

	var err error
	//
	//
	switch depotLedgerLevel {
	case statusLedgerLevel:
		//
		//
		if applicationLedgerLevel < depotLedgerLevel {
			//
			return h.resimulateLedgers(ctx, status, gatewayApplication, applicationLedgerLevel, depotLedgerLevel, false)
		} else if applicationLedgerLevel == depotLedgerLevel {
			//
			affirmApplicationDigestMatchesOneFromStatus(applicationDigest, status)
			return applicationDigest, nil
		}

	case statusLedgerLevel + 1:
		//
		//
		switch {
		case applicationLedgerLevel < statusLedgerLevel:
			//
			//
			return h.resimulateLedgers(ctx, status, gatewayApplication, applicationLedgerLevel, depotLedgerLevel, true)

		case applicationLedgerLevel == statusLedgerLevel:
			//
			//
			//
			//
			h.tracer.Details("REDACTED")
			status, err = h.resimulateLedger(status, depotLedgerLevel, gatewayApplication.Agreement())
			return status.ApplicationDigest, err

		case applicationLedgerLevel == depotLedgerLevel:
			//
			completeLedgerReply, err := h.statusDepot.ImportFinalCompleteLedgerReply(depotLedgerLevel)
			if err != nil {
				return nil, err
			}
			//
			//
			//
			//
			if len(completeLedgerReply.ApplicationDigest) == 0 {
				completeLedgerReply.ApplicationDigest = applicationDigest
			}
			emulateApplication := newEmulateGatewayApplication(completeLedgerReply)
			h.tracer.Details("REDACTED")
			status, err = h.resimulateLedger(status, depotLedgerLevel, emulateApplication)
			return status.ApplicationDigest, err
		}

	}

	panic(fmt.Sprintf("REDACTED",
		applicationLedgerLevel, depotLedgerLevel, statusLedgerLevel))
}

func (h *Greeter) resimulateLedgers(
	ctx context.Context,
	status sm.Status,
	gatewayApplication gateway.ApplicationLinks,
	applicationLedgerLevel,
	depotLedgerLevel int64,
	transformStatus bool,
) ([]byte, error) {
	//
	//
	//
	//
	//
	//
	//
	//
	//

	var applicationDigest []byte
	var err error
	ultimateLedger := depotLedgerLevel
	if transformStatus {
		ultimateLedger--
	}
	initialLedger := applicationLedgerLevel + 1
	if initialLedger == 1 {
		initialLedger = status.PrimaryLevel
	}
	for i := initialLedger; i <= ultimateLedger; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		h.tracer.Details("REDACTED", "REDACTED", i)
		ledger := h.depot.ImportLedger(i)
		//
		if len(applicationDigest) > 0 {
			affirmApplicationDigestMatchesOneFromLedger(applicationDigest, ledger)
		}

		applicationDigest, err = sm.InvokeEndorseLedger(gatewayApplication.Agreement(), ledger, h.tracer, h.statusDepot, h.generatePaper.PrimaryLevel)
		if err != nil {
			return nil, err
		}

		h.nLedgers++
	}

	if transformStatus {
		//
		status, err = h.resimulateLedger(status, depotLedgerLevel, gatewayApplication.Agreement())
		if err != nil {
			return nil, err
		}
		applicationDigest = status.ApplicationDigest
	}

	affirmApplicationDigestMatchesOneFromStatus(applicationDigest, status)
	return applicationDigest, nil
}

//
func (h *Greeter) resimulateLedger(status sm.Status, level int64, gatewayApplication gateway.ApplicationLinkAgreement) (sm.Status, error) {
	ledger := h.depot.ImportLedger(level)
	meta := h.depot.ImportLedgerMeta(level)

	//
	//
	ledgerExecute := sm.NewLedgerRunner(h.statusDepot, h.tracer, gatewayApplication, emptyTxpool{}, sm.EmptyProofDepository{}, h.depot)
	ledgerExecute.AssignEventBus(h.eventBus)

	var err error
	status, err = ledgerExecute.ExecuteLedger(status, meta.LedgerUID, ledger)
	if err != nil {
		return sm.Status{}, err
	}

	h.nLedgers++

	return status, nil
}

func affirmApplicationDigestMatchesOneFromLedger(applicationDigest []byte, ledger *kinds.Ledger) {
	if !bytes.Equal(applicationDigest, ledger.ApplicationDigest) {
		panic(fmt.Sprintf(`REDACTED.

REDACTEDv
REDACTED`,
			applicationDigest, ledger.ApplicationDigest, ledger))
	}
}

func affirmApplicationDigestMatchesOneFromStatus(applicationDigest []byte, status sm.Status) {
	if !bytes.Equal(applicationDigest, status.ApplicationDigest) {
		panic(fmt.Sprintf(`REDACTEDt
REDACTED.

REDACTEDv

REDACTED`,
			applicationDigest, status.ApplicationDigest, status))
	}
}
