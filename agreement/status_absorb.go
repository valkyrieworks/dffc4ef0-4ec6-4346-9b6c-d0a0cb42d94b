package agreement

import (
	"fmt"
	"time"

	controlkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agreement/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
	"github.com/pkg/errors"
)

//
type AbsorbNominee struct {
	ledger      *kinds.Ledger
	ledgerFragments *kinds.FragmentAssign
	endorse     *kinds.Endorse
	addnEndorse  *kinds.ExpandedEndorse

	attested bool

	//
	bufferedLedgerUUID kinds.LedgerUUID
}

type absorbAttestedLedgerSolicit struct {
	AbsorbNominee
	relayedLocated   time.Time
	reply chan absorbAttestedLedgerReply
}

type absorbAttestedLedgerReply struct {
	err error
}

//
func FreshAbsorbNominee(
	ledger *kinds.Ledger,
	ledgerFragments *kinds.FragmentAssign,
	endorse *kinds.Endorse,
	addnEndorse *kinds.ExpandedEndorse,
) (AbsorbNominee, error) {
	ic := AbsorbNominee{
		ledger:      ledger,
		ledgerFragments: ledgerFragments,
		endorse:     endorse,
		addnEndorse:  addnEndorse,
	}

	if err := ic.CertifyFundamental(); err != nil {
		return ic, err
	}

	return ic, nil
}

//
func (ic *AbsorbNominee) Altitude() int64 {
	return ic.ledger.Altitude
}

//
func (ic *AbsorbNominee) LedgerUUID() kinds.LedgerUUID {
	if !ic.bufferedLedgerUUID.EqualsNull() {
		return ic.bufferedLedgerUUID
	}

	ic.bufferedLedgerUUID = kinds.LedgerUUID{
		Digest:          ic.ledger.Digest(),
		FragmentAssignHeading: ic.ledgerFragments.Heading(),
	}

	return ic.bufferedLedgerUUID
}

//
func (ic *AbsorbNominee) CertifyFundamental() error {
	switch {
	case ic.ledger == nil:
		return errors.Wrap(FaultCertification, "REDACTED")
	case ic.ledgerFragments == nil:
		return errors.Wrap(FaultCertification, "REDACTED")
	case ic.endorse == nil:
		return errors.Wrap(FaultCertification, "REDACTED")
	}

	//
	var (
		ledgerUUID     = ic.LedgerUUID()
		ledgerAltitude = ic.ledger.Altitude
	)

	if ic.additionsActivated() {
		switch {
		case ic.addnEndorse.Altitude != ledgerAltitude:
			return errors.Wrapf(FaultCertification, "REDACTED", ic.addnEndorse.Altitude, ledgerAltitude)
		case !ic.addnEndorse.LedgerUUID.Matches(ledgerUUID):
			return errors.Wrap(FaultCertification, "REDACTED")
		default:
			return ic.addnEndorse.CertifyFundamental()
		}
	}

	switch {
	case ic.endorse.Altitude != ledgerAltitude:
		return errors.Wrapf(FaultCertification, "REDACTED", ic.endorse.Altitude, ledgerAltitude)
	case !ic.endorse.LedgerUUID.Matches(ledgerUUID):
		return errors.Wrap(FaultCertification, "REDACTED")
	default:
		return ic.endorse.CertifyFundamental()
	}
}

//
func (ic *AbsorbNominee) Validate(status status.Status) error {
	var (
		altitude            = ic.Altitude()
		ledgerUUID           = ic.LedgerUUID()
		successionUUID           = status.SuccessionUUID
		additionsExisting = status.AgreementSettings.Iface.BallotAdditionsActivated(altitude)
	)

	//
	if additionsExisting != ic.additionsActivated() {
		return fmt.Errorf(
			"REDACTED",
			ic.Altitude(), additionsExisting, ic.additionsActivated(),
		)
	}

	if err := status.CertifyLedger(ic.ledger); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	err := status.Assessors.ValidateEndorseAgile(successionUUID, ledgerUUID, altitude, ic.endorse)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if ic.additionsActivated() {
		if err = ic.addnEndorse.AssureAdditions(true); err != nil {
			return fmt.Errorf("REDACTED", err)
		}

		err = status.Assessors.ValidateEndorseAgile(successionUUID, ledgerUUID, altitude, ic.addnEndorse.TowardEndorse())
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
	}

	ic.attested = true

	return nil
}

func (ic *AbsorbNominee) additionsActivated() bool {
	return ic.addnEndorse != nil
}

//
func (ic *AbsorbNominee) endorseBalloting(successionUUID string, values *kinds.AssessorAssign) (iteration int32, ballotAssign *kinds.BallotAssign) {
	if ic.additionsActivated() {
		return ic.addnEndorse.Iteration, ic.addnEndorse.TowardExpandedBallotAssign(successionUUID, values)
	}

	return ic.endorse.Iteration, ic.endorse.TowardBallotAssign(successionUUID, values)
}

//
//
//
//
func (cs *Status) AbsorbAttestedLedger(ic AbsorbNominee) (err error) {
	tracer := cs.Tracer.Using("REDACTED", ic.Altitude())
	tracer.Details("REDACTED")

	defer func() {
		if err != nil {
			tracer.Details("REDACTED", "REDACTED", err)
		} else {
			tracer.Details("REDACTED")
		}
	}()

	//
	ch := make(chan absorbAttestedLedgerReply, 1)
	defer close(ch)

	req := &absorbAttestedLedgerSolicit{
		AbsorbNominee: ic,
		relayedLocated:          time.Now(),
		reply:        ch,
	}

	cs.transmitIntrinsicSignal(signalDetails{Msg: req})

	select {
	case <-cs.Exit():
		return fmt.Errorf("REDACTED")
	case res := <-req.reply:
		return res.err
	}
}

//
func (cs *Status) processAbsorbAttestedLedgerSolicit(req *absorbAttestedLedgerSolicit) {
	err := cs.absorbLedger(req.AbsorbNominee)

	req.reply <- absorbAttestedLedgerReply{err: err}
}

//
//
//
func (cs *Status) absorbLedger(ic AbsorbNominee) error {
	if !ic.attested {
		return errors.Wrap(FaultCertification, "REDACTED")
	}

	var (
		ledger           = ic.ledger
		ledgerFragments      = ic.ledgerFragments
		altitude          = ic.ledger.Altitude
		finalLedgerAltitude = cs.status.FinalLedgerAltitude
	)

	//
	//
	if altitude <= finalLedgerAltitude {
		return FaultEarlierComprised
	}

	//
	//
	if altitude != finalLedgerAltitude+1 {
		return errors.Wrapf(FaultAltitudeInterval, "REDACTED", altitude, finalLedgerAltitude+1)
	}

	//
	//
	var (
		statusDuplicate = cs.status.Duplicate()
		tracer    = cs.Tracer.Using("REDACTED", altitude)
	)

	//
	endorseIteration, endorseBallotAssign := ic.endorseBalloting(statusDuplicate.SuccessionUUID, statusDuplicate.Assessors)

	cs.reviseIterationPhase(endorseIteration, controlkinds.IterationPhaseEndorse)
	cs.EndorseIteration = endorseIteration
	cs.FinalEndorse = endorseBallotAssign
	cs.EndorseMoment = committime.Now()
	cs.freshPhase()

	//

	//
	//
	if ic.additionsActivated() {
		cs.ledgerDepot.PersistLedgerUsingExpandedEndorse(ledger, ledgerFragments, ic.addnEndorse)
	} else {
		cs.ledgerDepot.PersistLedger(ledger, ledgerFragments, ic.endorse)
	}

	//
	if err := cs.wal.PersistChronize(TerminateAltitudeSignal{altitude}); err != nil {
		panic(errors.Wrapf(err, "REDACTED", altitude))
	}

	//
	statusDuplicate, err := cs.ledgerExecute.ExecuteAttestedLedger(statusDuplicate, ic.LedgerUUID(), ledger)
	if err != nil {
		//
		panic(errors.Wrapf(err, "REDACTED", ledger.Altitude, ledger.Digest()))
	}

	//
	cs.logTelemetry(altitude, ledger)

	//
	//
	//
	//
	//
	cs.Ballots = nil
	cs.reviseTowardStatus(statusDuplicate)

	//
	if err := cs.revisePrivateAssessorPublicToken(); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	cs.timelineCycle0(&cs.IterationStatus)

	return nil
}
