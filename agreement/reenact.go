package agreement

import (
	"bytes"
	"context"
	"fmt"
	"hash/crc32"
	"io"
	"reflect"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var checksum32c = crc32.MakeTable(crc32.Castagnoli)

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
func (cs *Status) retrieveReenactSignal(msg *ScheduledJournalSignal, freshPhaseUnder kinds.Listening) error {
	//
	if _, ok := msg.Msg.(TerminateAltitudeSignal); ok {
		return nil
	}

	//
	switch m := msg.Msg.(type) {
	case kinds.IncidentDataIterationStatus:
		cs.Tracer.Details("REDACTED", "REDACTED", m.Altitude, "REDACTED", m.Iteration, "REDACTED", m.Phase)
		//
		metronome := time.After(time.Second * 2)
		if freshPhaseUnder != nil {
			select {
			case phaseSignal := <-freshPhaseUnder.Out():
				m2 := phaseSignal.Data().(kinds.IncidentDataIterationStatus)
				if m.Altitude != m2.Altitude || m.Iteration != m2.Iteration || m.Phase != m2.Phase {
					return fmt.Errorf("REDACTED", m2, m)
				}
			case <-freshPhaseUnder.Aborted():
				return fmt.Errorf("REDACTED")
			case <-metronome:
				return fmt.Errorf("REDACTED")
			}
		}
	case signalDetails:
		nodeUUID := m.NodeUUID
		if nodeUUID == "REDACTED" {
			nodeUUID = "REDACTED"
		}
		switch msg := m.Msg.(type) {
		case *NominationSignal:
			p := msg.Nomination
			cs.Tracer.Details("REDACTED", "REDACTED", p.Altitude, "REDACTED", p.Iteration, "REDACTED",
				p.LedgerUUID.FragmentAssignHeading, "REDACTED", p.PolicyIteration, "REDACTED", nodeUUID)
		case *LedgerFragmentSignal:
			cs.Tracer.Details("REDACTED", "REDACTED", msg.Altitude, "REDACTED", msg.Iteration, "REDACTED", nodeUUID)
		case *BallotSignal:
			v := msg.Ballot
			cs.Tracer.Details("REDACTED", "REDACTED", v.Altitude, "REDACTED", v.Iteration, "REDACTED", v.Kind,
				"REDACTED", v.LedgerUUID, "REDACTED", nodeUUID, "REDACTED", len(v.Addition), "REDACTED", len(v.AdditionNotation))
		}

		cs.processSignal(m)
	case deadlineDetails:
		cs.Tracer.Details("REDACTED", "REDACTED", m.Altitude, "REDACTED", m.Iteration, "REDACTED", m.Phase, "REDACTED", m.Interval)
		cs.processDeadline(m, cs.IterationStatus)
	default:
		return fmt.Errorf("REDACTED", reflect.TypeOf(msg.Msg))
	}
	return nil
}

//
//
func (cs *Status) overtakeReenact(controlAltitude int64) error {
	//
	cs.reenactStyle = true
	defer func() { cs.reenactStyle = false }()

	//
	//
	//
	//
	//
	//
	gr, detected, err := cs.wal.LookupForeachTerminateAltitude(controlAltitude, &JournalLookupChoices{BypassDataImpairmentFaults: true})
	if err != nil {
		return err
	}
	if gr != nil {
		if err := gr.Close(); err != nil {
			return err
		}
	}
	if detected {
		return fmt.Errorf("REDACTED", controlAltitude)
	}

	//
	//
	//
	if controlAltitude < cs.status.PrimaryAltitude {
		return fmt.Errorf("REDACTED", controlAltitude, cs.status.PrimaryAltitude)
	}
	terminateAltitude := controlAltitude - 1
	if controlAltitude == cs.status.PrimaryAltitude {
		terminateAltitude = 0
	}
	gr, detected, err = cs.wal.LookupForeachTerminateAltitude(terminateAltitude, &JournalLookupChoices{BypassDataImpairmentFaults: true})
	if err == io.EOF {
		cs.Tracer.Failure("REDACTED", "REDACTED", terminateAltitude)
	} else if err != nil {
		return err
	}
	if !detected {
		return fmt.Errorf("REDACTED", controlAltitude, terminateAltitude)
	}
	defer gr.Close()

	cs.Tracer.Details("REDACTED", "REDACTED", controlAltitude)

	var msg *ScheduledJournalSignal
	dec := JournalDeserializer{gr}

Cycle:
	for {
		msg, err = dec.Deserialize()
		switch {
		case err == io.EOF:
			break Cycle
		case EqualsDataImpairmentFailure(err):
			cs.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", controlAltitude)
			return err
		case err != nil:
			return err
		}

		//
		//
		//
		if err := cs.retrieveReenactSignal(msg, nil); err != nil {
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

type Negotiator struct {
	statusDepot   sm.Depot
	primaryStatus sm.Status
	depot        sm.LedgerDepot
	incidentChannel     kinds.LedgerIncidentBroadcaster
	producePaper       *kinds.OriginPaper
	tracer       log.Tracer

	nthLedgers int //
}

func FreshNegotiator(statusDepot sm.Depot, status sm.Status,
	depot sm.LedgerDepot, producePaper *kinds.OriginPaper,
) *Negotiator {
	return &Negotiator{
		statusDepot:   statusDepot,
		primaryStatus: status,
		depot:        depot,
		incidentChannel:     kinds.NooperationIncidentPipeline{},
		producePaper:       producePaper,
		tracer:       log.FreshNooperationTracer(),
		nthLedgers:      0,
	}
}

func (h *Negotiator) AssignTracer(l log.Tracer) {
	h.tracer = l
}

//
//
func (h *Negotiator) AssignIncidentChannel(incidentChannel kinds.LedgerIncidentBroadcaster) {
	h.incidentChannel = incidentChannel
}

//
func (h *Negotiator) NTHLedgers() int {
	return h.nthLedgers
}

//
func (h *Negotiator) Negotiation(delegatePlatform delegate.PlatformLinks) error {
	return h.NegotiationUsingEnv(context.TODO(), delegatePlatform)
}

//
func (h *Negotiator) NegotiationUsingEnv(ctx context.Context, delegatePlatform delegate.PlatformLinks) error {
	//
	res, err := delegatePlatform.Inquire().Details(ctx, delegate.SolicitDetails)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	ledgerAltitude := res.FinalLedgerAltitude
	if ledgerAltitude < 0 {
		return fmt.Errorf("REDACTED", ledgerAltitude)
	}
	platformDigest := res.FinalLedgerPlatformDigest

	h.tracer.Details("REDACTED",
		"REDACTED", ledgerAltitude,
		"REDACTED", log.FreshIdleFormat("REDACTED", platformDigest),
		"REDACTED", res.Edition,
		"REDACTED", res.PlatformEdition,
	)

	//
	if h.primaryStatus.FinalLedgerAltitude == 0 {
		h.primaryStatus.Edition.Agreement.App = res.PlatformEdition
	}

	//
	platformDigest, err = h.ReenactLedgersUsingEnv(ctx, h.primaryStatus, platformDigest, ledgerAltitude, delegatePlatform)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	h.tracer.Details("REDACTED",
		"REDACTED", ledgerAltitude, "REDACTED", log.FreshIdleFormat("REDACTED", platformDigest))

	//

	return nil
}

//
//
//
func (h *Negotiator) ReenactLedgers(
	status sm.Status,
	platformDigest []byte,
	applicationLedgerAltitude int64,
	delegatePlatform delegate.PlatformLinks,
) ([]byte, error) {
	return h.ReenactLedgersUsingEnv(context.TODO(), status, platformDigest, applicationLedgerAltitude, delegatePlatform)
}

//
func (h *Negotiator) ReenactLedgersUsingEnv(
	ctx context.Context,
	status sm.Status,
	platformDigest []byte,
	applicationLedgerAltitude int64,
	delegatePlatform delegate.PlatformLinks,
) ([]byte, error) {
	depotLedgerFoundation := h.depot.Foundation()
	depotLedgerAltitude := h.depot.Altitude()
	statusLedgerAltitude := status.FinalLedgerAltitude
	h.tracer.Details(
		"REDACTED",
		"REDACTED",
		applicationLedgerAltitude,
		"REDACTED",
		depotLedgerAltitude,
		"REDACTED",
		statusLedgerAltitude)

	//
	if applicationLedgerAltitude == 0 {
		assessors := make([]*kinds.Assessor, len(h.producePaper.Assessors))
		for i, val := range h.producePaper.Assessors {
			assessors[i] = kinds.FreshAssessor(val.PublicToken, val.Potency)
		}
		assessorAssign := kinds.FreshAssessorAssign(assessors)
		followingValues := kinds.Temp2buffer.AssessorRevisions(assessorAssign)
		bufferargs := h.producePaper.AgreementSettings.TowardSchema()
		req := &iface.SolicitInitializeSuccession{
			Moment:            h.producePaper.OriginMoment,
			SuccessionUuid:         h.producePaper.SuccessionUUID,
			PrimaryAltitude:   h.producePaper.PrimaryAltitude,
			AgreementSettings: &bufferargs,
			Assessors:      followingValues,
			ApplicationStatusOctets:   h.producePaper.ApplicationStatus,
		}
		res, err := delegatePlatform.Agreement().InitializeSuccession(context.TODO(), req)
		if err != nil {
			return nil, err
		}

		platformDigest = res.PlatformDigest

		if statusLedgerAltitude == 0 { //
			//
			//
			//
			if len(res.PlatformDigest) > 0 {
				status.PlatformDigest = res.PlatformDigest
			}
			//
			if len(res.Assessors) > 0 {
				values, err := kinds.Buffer2temp.AssessorRevisions(res.Assessors)
				if err != nil {
					return nil, err
				}
				status.Assessors = kinds.FreshAssessorAssign(values)
				status.FollowingAssessors = kinds.FreshAssessorAssign(values).DuplicateAdvanceNominatorUrgency(1)
			} else if len(h.producePaper.Assessors) == 0 {
				//
				return nil, fmt.Errorf("REDACTED")
			}

			if res.AgreementSettings != nil {
				status.AgreementSettings = status.AgreementSettings.Revise(res.AgreementSettings)
				status.Edition.Agreement.App = status.AgreementSettings.Edition.App
			}
			//
			status.FinalOutcomesDigest = hashmap.DigestOriginatingOctetSegments(nil)
			if err := h.statusDepot.Persist(status); err != nil {
				return nil, err
			}
		}
	}

	//
	switch {
	case depotLedgerAltitude == 0:
		attestApplicationDigestMatchesSingleOriginatingStatus(platformDigest, status)
		return platformDigest, nil

	case applicationLedgerAltitude == 0 && status.PrimaryAltitude < depotLedgerFoundation:
		//
		return platformDigest, sm.FaultApplicationLedgerAltitudeExcessivelyInferior{ApplicationAltitude: applicationLedgerAltitude, DepotFoundation: depotLedgerFoundation}

	case applicationLedgerAltitude > 0 && applicationLedgerAltitude < depotLedgerFoundation-1:
		//
		return platformDigest, sm.FaultApplicationLedgerAltitudeExcessivelyInferior{ApplicationAltitude: applicationLedgerAltitude, DepotFoundation: depotLedgerFoundation}

	case depotLedgerAltitude < applicationLedgerAltitude:
		//
		return platformDigest, sm.FaultApplicationLedgerAltitudeExcessivelySuperior{BaseAltitude: depotLedgerAltitude, ApplicationAltitude: applicationLedgerAltitude}

	case depotLedgerAltitude < statusLedgerAltitude:
		//
		panic(fmt.Sprintf("REDACTED", statusLedgerAltitude, depotLedgerAltitude))

	case depotLedgerAltitude > statusLedgerAltitude+1:
		//
		panic(fmt.Sprintf("REDACTED", depotLedgerAltitude, statusLedgerAltitude+1))
	}

	var err error
	//
	//
	switch depotLedgerAltitude {
	case statusLedgerAltitude:
		//
		//
		if applicationLedgerAltitude < depotLedgerAltitude {
			//
			return h.reenactLedgers(ctx, status, delegatePlatform, applicationLedgerAltitude, depotLedgerAltitude, false)
		} else if applicationLedgerAltitude == depotLedgerAltitude {
			//
			attestApplicationDigestMatchesSingleOriginatingStatus(platformDigest, status)
			return platformDigest, nil
		}

	case statusLedgerAltitude + 1:
		//
		//
		switch {
		case applicationLedgerAltitude < statusLedgerAltitude:
			//
			//
			return h.reenactLedgers(ctx, status, delegatePlatform, applicationLedgerAltitude, depotLedgerAltitude, true)

		case applicationLedgerAltitude == statusLedgerAltitude:
			//
			//
			//
			//
			h.tracer.Details("REDACTED")
			status, err = h.reenactLedger(status, depotLedgerAltitude, delegatePlatform.Agreement())
			return status.PlatformDigest, err

		case applicationLedgerAltitude == depotLedgerAltitude:
			//
			culminateLedgerReply, err := h.statusDepot.FetchFinalCulminateLedgerReply(depotLedgerAltitude)
			if err != nil {
				return nil, err
			}
			//
			//
			//
			//
			if len(culminateLedgerReply.PlatformDigest) == 0 {
				culminateLedgerReply.PlatformDigest = platformDigest
			}
			simulateApplication := freshSimulateDelegateApplication(culminateLedgerReply)
			h.tracer.Details("REDACTED")
			status, err = h.reenactLedger(status, depotLedgerAltitude, simulateApplication)
			return status.PlatformDigest, err
		}

	}

	panic(fmt.Sprintf("REDACTED",
		applicationLedgerAltitude, depotLedgerAltitude, statusLedgerAltitude))
}

func (h *Negotiator) reenactLedgers(
	ctx context.Context,
	status sm.Status,
	delegatePlatform delegate.PlatformLinks,
	applicationLedgerAltitude,
	depotLedgerAltitude int64,
	alterStatus bool,
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

	var platformDigest []byte
	var err error
	ultimateLedger := depotLedgerAltitude
	if alterStatus {
		ultimateLedger--
	}
	initialLedger := applicationLedgerAltitude + 1
	if initialLedger == 1 {
		initialLedger = status.PrimaryAltitude
	}
	for i := initialLedger; i <= ultimateLedger; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		h.tracer.Details("REDACTED", "REDACTED", i)
		ledger := h.depot.FetchLedger(i)
		//
		if len(platformDigest) > 0 {
			attestApplicationDigestMatchesSingleOriginatingLedger(platformDigest, ledger)
		}

		platformDigest, err = sm.InvokeEndorseLedger(delegatePlatform.Agreement(), ledger, h.tracer, h.statusDepot, h.producePaper.PrimaryAltitude)
		if err != nil {
			return nil, err
		}

		h.nthLedgers++
	}

	if alterStatus {
		//
		status, err = h.reenactLedger(status, depotLedgerAltitude, delegatePlatform.Agreement())
		if err != nil {
			return nil, err
		}
		platformDigest = status.PlatformDigest
	}

	attestApplicationDigestMatchesSingleOriginatingStatus(platformDigest, status)
	return platformDigest, nil
}

//
func (h *Negotiator) reenactLedger(status sm.Status, altitude int64, delegatePlatform delegate.ApplicationLinkAgreement) (sm.Status, error) {
	ledger := h.depot.FetchLedger(altitude)
	summary := h.depot.FetchLedgerSummary(altitude)

	//
	//
	ledgerExecute := sm.FreshLedgerHandler(h.statusDepot, h.tracer, delegatePlatform, blankTxpool{}, sm.VoidProofHub{}, h.depot)
	ledgerExecute.AssignIncidentChannel(h.incidentChannel)

	var err error
	status, err = ledgerExecute.ExecuteLedger(status, summary.LedgerUUID, ledger)
	if err != nil {
		return sm.Status{}, err
	}

	h.nthLedgers++

	return status, nil
}

func attestApplicationDigestMatchesSingleOriginatingLedger(platformDigest []byte, ledger *kinds.Ledger) {
	if !bytes.Equal(platformDigest, ledger.PlatformDigest) {
		panic(fmt.Sprintf(`REDACTED.

REDACTEDv
REDACTED`,
			platformDigest, ledger.PlatformDigest, ledger))
	}
}

func attestApplicationDigestMatchesSingleOriginatingStatus(platformDigest []byte, status sm.Status) {
	if !bytes.Equal(platformDigest, status.PlatformDigest) {
		panic(fmt.Sprintf(`REDACTEDt
REDACTED.

REDACTEDv

REDACTED`,
			platformDigest, status.PlatformDigest, status))
	}
}
