package status

import (
	"bytes"
	"context"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/abort"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/txpool"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//

//
type LedgerHandler struct {
	//
	depot Depot

	//
	ledgerDepot LedgerDepot

	//
	delegatePlatform delegate.ApplicationLinkAgreement

	//
	incidentPipeline kinds.LedgerIncidentBroadcaster

	//
	//
	txpool txpool.Txpool
	incidentpool  ProofHub

	//
	finalVerifiedLedger *kinds.Ledger

	tracer log.Tracer

	telemetry *Telemetry
}

type LedgerHandlerSelection func(handler *LedgerHandler)

func LedgerHandlerUsingTelemetry(telemetry *Telemetry) LedgerHandlerSelection {
	return func(ledgerExecute *LedgerHandler) {
		ledgerExecute.telemetry = telemetry
	}
}

//
//
func FreshLedgerHandler(
	statusDepot Depot,
	tracer log.Tracer,
	delegatePlatform delegate.ApplicationLinkAgreement,
	txpool txpool.Txpool,
	incidentpool ProofHub,
	ledgerDepot LedgerDepot,
	choices ...LedgerHandlerSelection,
) *LedgerHandler {
	res := &LedgerHandler{
		depot:      statusDepot,
		delegatePlatform:   delegatePlatform,
		incidentPipeline:   kinds.NooperationIncidentPipeline{},
		txpool:    txpool,
		incidentpool:     incidentpool,
		tracer:     tracer,
		telemetry:    NooperationTelemetry(),
		ledgerDepot: ledgerDepot,
	}

	for _, selection := range choices {
		selection(res)
	}

	return res
}

func (ledgerExecute *LedgerHandler) Depot() Depot {
	return ledgerExecute.depot
}

//
//
func (ledgerExecute *LedgerHandler) AssignIncidentChannel(incidentPipeline kinds.LedgerIncidentBroadcaster) {
	ledgerExecute.incidentPipeline = incidentPipeline
}

//
//
//
//
//
//
func (ledgerExecute *LedgerHandler) GenerateNominationLedger(
	ctx context.Context,
	altitude int64,
	status Status,
	finalAddnEndorse *kinds.ExpandedEndorse,
	nominatorLocation []byte,
) (*kinds.Ledger, error) {
	maximumOctets := status.AgreementSettings.Ledger.MaximumOctets
	blankMaximumOctets := maximumOctets == -1
	if blankMaximumOctets {
		maximumOctets = int64(kinds.MaximumLedgerExtentOctets)
	}

	maximumFuel := status.AgreementSettings.Ledger.MaximumFuel

	proof, occurenceExtent := ledgerExecute.incidentpool.AwaitingProof(status.AgreementSettings.Proof.MaximumOctets)

	//
	maximumDataOctets := kinds.MaximumDataOctets(maximumOctets, occurenceExtent, status.Assessors.Extent())
	maximumHarvestOctets := maximumDataOctets
	if blankMaximumOctets {
		maximumHarvestOctets = -1
	}

	txs := ledgerExecute.txpool.HarvestMaximumOctetsMaximumFuel(maximumHarvestOctets, maximumFuel)
	endorse := finalAddnEndorse.TowardEndorse()
	ledger, err := status.CreateLedger(altitude, txs, endorse, proof, nominatorLocation)
	if err != nil {
		return nil, err
	}
	rpp, err := ledgerExecute.delegatePlatform.ArrangeNomination(
		ctx,
		&iface.SolicitArrangeNomination{
			MaximumTransferOctets:         maximumDataOctets,
			Txs:                ledger.Txs.TowardSegmentBelongingOctets(),
			RegionalFinalEndorse:    constructExpandedEndorseDetailsOriginatingDepot(finalAddnEndorse, ledgerExecute.depot, status.PrimaryAltitude, status.AgreementSettings.Iface),
			Malpractice:        ledger.Proof.Proof.TowardIface(),
			Altitude:             ledger.Altitude,
			Moment:               ledger.Moment,
			FollowingAssessorsDigest: ledger.FollowingAssessorsDigest,
			NominatorLocation:    ledger.NominatorLocation,
		},
	)
	if err != nil {
		//
		//
		//
		//
		//
		//
		//
		//
		return nil, err
	}

	txl := kinds.TowardTrans(rpp.Txs)
	if err := txl.Certify(maximumDataOctets); err != nil {
		return nil, err
	}

	return status.CreateLedger(altitude, txl, endorse, proof, nominatorLocation)
}

func (ledgerExecute *LedgerHandler) HandleNomination(
	ledger *kinds.Ledger,
	status Status,
) (bool, error) {
	reply, err := ledgerExecute.delegatePlatform.HandleNomination(context.TODO(), &iface.SolicitHandleNomination{
		Digest:               ledger.Heading.Digest(),
		Altitude:             ledger.Altitude,
		Moment:               ledger.Moment,
		Txs:                ledger.Txs.TowardSegmentBelongingOctets(),
		ItemizedFinalEndorse: constructFinalEndorseDetailsOriginatingDepot(ledger, ledgerExecute.depot, status.PrimaryAltitude),
		Malpractice:        ledger.Proof.Proof.TowardIface(),
		NominatorLocation:    ledger.NominatorLocation,
		FollowingAssessorsDigest: ledger.FollowingAssessorsDigest,
	})
	if err != nil {
		return false, err
	}
	if reply.EqualsConditionUnfamiliar() {
		panic(fmt.Sprintf("REDACTED", reply.Condition.Text()))
	}

	return reply.EqualsApproved(), nil
}

//
//
//
//
func (ledgerExecute *LedgerHandler) CertifyLedger(status Status, ledger *kinds.Ledger) error {
	if !ledgerExecute.finalVerifiedLedger.DigestsToward(ledger.Digest()) {
		if err := certifyLedger(status, ledger); err != nil {
			return err
		}
		ledgerExecute.finalVerifiedLedger = ledger
	}
	return ledgerExecute.incidentpool.InspectProof(ledger.Proof.Proof)
}

//
func (ledgerExecute *LedgerHandler) ExecuteAttestedLedger(
	status Status, ledgerUUID kinds.LedgerUUID, ledger *kinds.Ledger,
) (Status, error) {
	return ledgerExecute.executeLedger(status, ledgerUUID, ledger)
}

//
//
//
//
//
//
func (ledgerExecute *LedgerHandler) ExecuteLedger(
	status Status, ledgerUUID kinds.LedgerUUID, ledger *kinds.Ledger,
) (Status, error) {
	if !ledgerExecute.finalVerifiedLedger.DigestsToward(ledger.Digest()) {
		if err := certifyLedger(status, ledger); err != nil {
			return status, FaultUnfitLedger(err)
		}
		ledgerExecute.finalVerifiedLedger = ledger
	}
	return ledgerExecute.executeLedger(status, ledgerUUID, ledger)
}

func (ledgerExecute *LedgerHandler) executeLedger(status Status, ledgerUUID kinds.LedgerUUID, ledger *kinds.Ledger) (Status, error) {
	initiateMoment := time.Now().UnixNano()
	ifaceReply, err := ledgerExecute.delegatePlatform.CulminateLedger(context.TODO(), &iface.SolicitCulminateLedger{
		Digest:               ledger.Digest(),
		FollowingAssessorsDigest: ledger.FollowingAssessorsDigest,
		NominatorLocation:    ledger.NominatorLocation,
		Altitude:             ledger.Altitude,
		Moment:               ledger.Moment,
		ResolvedFinalEndorse:  constructFinalEndorseDetailsOriginatingDepot(ledger, ledgerExecute.depot, status.PrimaryAltitude),
		Malpractice:        ledger.Proof.Proof.TowardIface(),
		Txs:                ledger.Txs.TowardSegmentBelongingOctets(),
	})
	terminateMoment := time.Now().UnixNano()
	ledgerExecute.telemetry.LedgerHandlingMoment.Observe(float64(terminateMoment-initiateMoment) / 1000000)
	if err != nil {
		ledgerExecute.tracer.Failure("REDACTED", "REDACTED", err)
		return status, err
	}

	ledgerExecute.tracer.Details(
		"REDACTED",
		"REDACTED", ledger.Altitude,
		"REDACTED", len(ifaceReply.TransferOutcomes),
		"REDACTED", len(ifaceReply.AssessorRevisions),
		"REDACTED", fmt.Sprintf("REDACTED", ifaceReply.PlatformDigest),
	)

	//
	if len(ledger.Txs) != len(ifaceReply.TransferOutcomes) {
		return status, fmt.Errorf("REDACTED", len(ledger.Txs), len(ifaceReply.TransferOutcomes))
	}

	ledgerExecute.tracer.Details("REDACTED", "REDACTED", ledger.Altitude, "REDACTED", fmt.Sprintf("REDACTED", ifaceReply.PlatformDigest))

	abort.Mishap() //

	//
	if err := ledgerExecute.depot.PersistCulminateLedgerReply(ledger.Altitude, ifaceReply); err != nil {
		return status, err
	}

	abort.Mishap() //

	//
	err = certifyAssessorRevisions(ifaceReply.AssessorRevisions, status.AgreementSettings.Assessor)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	assessorRevisions, err := kinds.Buffer2temp.AssessorRevisions(ifaceReply.AssessorRevisions)
	if err != nil {
		return status, err
	}
	if len(assessorRevisions) > 0 {
		ledgerExecute.tracer.Details("REDACTED", "REDACTED", kinds.AssessorCatalogText(assessorRevisions))
		ledgerExecute.telemetry.AssessorAssignRevisions.Add(1)
	}
	if ifaceReply.AgreementArgumentRevisions != nil {
		ledgerExecute.telemetry.AgreementArgumentRevisions.Add(1)
	}

	//
	status, err = reviseStatus(status, ledgerUUID, &ledger.Heading, ifaceReply, assessorRevisions)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	//
	preserveAltitude, err := ledgerExecute.Endorse(status, ledger, ifaceReply)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	//
	ledgerExecute.incidentpool.Revise(status, ledger.Proof.Proof)

	abort.Mishap() //

	//
	status.PlatformDigest = ifaceReply.PlatformDigest
	if err := ledgerExecute.depot.Persist(status); err != nil {
		return status, err
	}

	abort.Mishap() //

	//
	if preserveAltitude > 0 {
		trimmed, err := ledgerExecute.trimLedgers(preserveAltitude, status)
		if err != nil {
			ledgerExecute.tracer.Failure("REDACTED", "REDACTED", preserveAltitude, "REDACTED", err)
		} else {
			ledgerExecute.tracer.Diagnose("REDACTED", "REDACTED", trimmed, "REDACTED", preserveAltitude)
		}
	}

	//
	//
	triggerIncidents(ledgerExecute.tracer, ledgerExecute.incidentPipeline, ledger, ledgerUUID, ifaceReply, assessorRevisions)

	return status, nil
}

func (ledgerExecute *LedgerHandler) BroadenBallot(
	ctx context.Context,
	ballot *kinds.Ballot,
	ledger *kinds.Ledger,
	status Status,
) ([]byte, error) {
	if !ledger.DigestsToward(ballot.LedgerUUID.Digest) {
		panic(fmt.Sprintf("REDACTED", ledger.Digest(), ballot.LedgerUUID.Digest))
	}
	if ballot.Altitude != ledger.Altitude {
		panic(fmt.Sprintf("REDACTED", ledger.Altitude, ballot.Altitude))
	}

	req := iface.SolicitBroadenBallot{
		Digest:               ballot.LedgerUUID.Digest,
		Altitude:             ballot.Altitude,
		Moment:               ledger.Moment,
		Txs:                ledger.Txs.TowardSegmentBelongingOctets(),
		ItemizedFinalEndorse: constructFinalEndorseDetailsOriginatingDepot(ledger, ledgerExecute.depot, status.PrimaryAltitude),
		Malpractice:        ledger.Proof.Proof.TowardIface(),
		FollowingAssessorsDigest: ledger.FollowingAssessorsDigest,
		NominatorLocation:    ledger.NominatorLocation,
	}

	reply, err := ledgerExecute.delegatePlatform.BroadenBallot(ctx, &req)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return reply.BallotAddition, nil
}

func (ledgerExecute *LedgerHandler) ValidateBallotAddition(ctx context.Context, ballot *kinds.Ballot) error {
	req := iface.SolicitValidateBallotAddition{
		Digest:             ballot.LedgerUUID.Digest,
		AssessorLocation: ballot.AssessorLocation,
		Altitude:           ballot.Altitude,
		BallotAddition:    ballot.Addition,
	}

	reply, err := ledgerExecute.delegatePlatform.ValidateBallotAddition(ctx, &req)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if reply.EqualsConditionUnfamiliar() {
		panic(fmt.Sprintf("REDACTED", reply.Condition.Text()))
	}

	if !reply.EqualsApproved() {
		return kinds.FaultUnfitBallotAddition
	}
	return nil
}

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
func (ledgerExecute *LedgerHandler) Endorse(
	status Status,
	ledger *kinds.Ledger,
	ifaceReply *iface.ReplyCulminateLedger,
) (int64, error) {
	ledgerExecute.txpool.Secure()
	releaseTxpool := func() { ledgerExecute.txpool.Release() }

	//
	//
	err := ledgerExecute.txpool.PurgeApplicationLink()
	if err != nil {
		releaseTxpool()
		ledgerExecute.tracer.Failure("REDACTED", "REDACTED", err)
		return 0, err
	}

	//
	res, err := ledgerExecute.delegatePlatform.Endorse(context.TODO())
	if err != nil {
		releaseTxpool()
		ledgerExecute.tracer.Failure("REDACTED", "REDACTED", err)
		return 0, err
	}

	//
	ledgerExecute.tracer.Details(
		"REDACTED",
		"REDACTED", ledger.Altitude,
		"REDACTED", fmt.Sprintf("REDACTED", ledger.PlatformDigest),
	)

	//
	go ledgerExecute.asyncronousReviseTxpool(releaseTxpool, ledger, status.Duplicate(), ifaceReply)

	return res.PreserveAltitude, nil
}

//
func (ledgerExecute *LedgerHandler) asyncronousReviseTxpool(
	releaseTxpool func(),
	ledger *kinds.Ledger,
	status Status,
	ifaceReply *iface.ReplyCulminateLedger,
) {
	defer releaseTxpool()

	err := ledgerExecute.txpool.Revise(
		ledger.Altitude,
		ledger.Txs,
		ifaceReply.TransferOutcomes,
		TransferPriorInspect(status),
		TransferRelayInspect(status),
	)
	if err != nil {
		//
		//
		//
		//
		//
		//
		panic(fmt.Sprintf("REDACTED", err))
	}
}

//
//

func constructFinalEndorseDetailsOriginatingDepot(ledger *kinds.Ledger, depot Depot, primaryAltitude int64) iface.EndorseDetails {
	if ledger.Altitude == primaryAltitude { //
		//
		//
		return iface.EndorseDetails{}
	}

	finalItemAssign, err := depot.FetchAssessors(ledger.Altitude - 1)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ledger.Altitude-1, err))
	}

	return ConstructFinalEndorseDetails(ledger, finalItemAssign, primaryAltitude)
}

//
//
//
func ConstructFinalEndorseDetails(ledger *kinds.Ledger, finalItemAssign *kinds.AssessorAssign, primaryAltitude int64) iface.EndorseDetails {
	if ledger.Altitude == primaryAltitude {
		//
		//
		return iface.EndorseDetails{}
	}

	var (
		endorseExtent = ledger.FinalEndorse.Extent()
		itemAssignLength  = len(finalItemAssign.Assessors)
	)

	//
	//
	if endorseExtent != itemAssignLength {
		panic(fmt.Sprintf(
			"REDACTED",
			endorseExtent, itemAssignLength, ledger.Altitude, ledger.FinalEndorse.Notations, finalItemAssign.Assessors,
		))
	}

	ballots := make([]iface.BallotDetails, ledger.FinalEndorse.Extent())
	for i, val := range finalItemAssign.Assessors {
		endorseSignature := ledger.FinalEndorse.Notations[i]
		ballots[i] = iface.BallotDetails{
			Assessor:   kinds.Temp2buffer.Assessor(val),
			LedgerUuidMarker: commitchema.LedgerUUIDMarker(endorseSignature.LedgerUUIDMarker),
		}
	}

	return iface.EndorseDetails{
		Iteration: ledger.FinalEndorse.Iteration,
		Ballots: ballots,
	}
}

//
//
//
//
//
//
//
//
//
func constructExpandedEndorseDetailsOriginatingDepot(ec *kinds.ExpandedEndorse, depot Depot, primaryAltitude int64, ap kinds.IfaceParameters) iface.ExpandedEndorseDetails {
	if ec.Altitude < primaryAltitude {
		//
		return iface.ExpandedEndorseDetails{}
	}

	itemAssign, err := depot.FetchAssessors(ec.Altitude)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ec.Altitude, primaryAltitude, err))
	}

	return ConstructExpandedEndorseDetails(ec, itemAssign, primaryAltitude, ap)
}

//
//
//
func ConstructExpandedEndorseDetails(ec *kinds.ExpandedEndorse, itemAssign *kinds.AssessorAssign, primaryAltitude int64, ap kinds.IfaceParameters) iface.ExpandedEndorseDetails {
	if ec.Altitude < primaryAltitude {
		//
		return iface.ExpandedEndorseDetails{}
	}

	var (
		eccodeExtent    = ec.Extent()
		itemAssignLength = len(itemAssign.Assessors)
	)

	//
	//
	if eccodeExtent != itemAssignLength {
		panic(fmt.Errorf(
			"REDACTED",
			eccodeExtent, itemAssignLength, ec.Altitude, ec.ExpandedNotations, itemAssign.Assessors,
		))
	}

	ballots := make([]iface.ExpandedBallotDetails, eccodeExtent)
	for i, val := range itemAssign.Assessors {
		ecs := ec.ExpandedNotations[i]

		//
		//
		if ecs.LedgerUUIDMarker != kinds.LedgerUUIDMarkerMissing && !bytes.Equal(ecs.AssessorLocation, val.Location) {
			panic(fmt.Errorf("REDACTED",
				i, ecs.AssessorLocation, ec.Altitude, val.Location,
			))
		}

		//
		//
		//
		//
		//
		if err := ecs.AssureAddition(ap.BallotAdditionsActivated(ec.Altitude)); err != nil {
			panic(fmt.Errorf("REDACTED", ec.Altitude, err))
		}

		ballots[i] = iface.ExpandedBallotDetails{
			Assessor:          kinds.Temp2buffer.Assessor(val),
			LedgerUuidMarker:        commitchema.LedgerUUIDMarker(ecs.LedgerUUIDMarker),
			BallotAddition:      ecs.Addition,
			AdditionNotation: ecs.AdditionNotation,
		}
	}

	return iface.ExpandedEndorseDetails{
		Iteration: ec.Iteration,
		Ballots: ballots,
	}
}

func certifyAssessorRevisions(ifaceRevisions []iface.AssessorRevise,
	parameters kinds.AssessorParameters,
) error {
	for _, itemRevise := range ifaceRevisions {
		if itemRevise.ObtainPotency() < 0 {
			return fmt.Errorf("REDACTED", itemRevise)
		} else if itemRevise.ObtainPotency() == 0 {
			//
			//
			continue
		}

		//
		pk, err := cryptocode.PublicTokenOriginatingSchema(itemRevise.PublicToken)
		if err != nil {
			return err
		}

		if !kinds.EqualsSoundPublickeyKind(parameters, pk.Kind()) {
			return fmt.Errorf("REDACTED",
				itemRevise, pk.Kind())
		}
	}
	return nil
}

//
func reviseStatus(
	status Status,
	ledgerUUID kinds.LedgerUUID,
	heading *kinds.Heading,
	ifaceReply *iface.ReplyCulminateLedger,
	assessorRevisions []*kinds.Assessor,
) (Status, error) {
	//
	//
	nthItemAssign := status.FollowingAssessors.Duplicate()

	//
	finalAltitudeValuesAltered := status.FinalAltitudeAssessorsAltered
	if len(assessorRevisions) > 0 {
		err := nthItemAssign.ReviseUsingModifyAssign(assessorRevisions)
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}
		//
		finalAltitudeValuesAltered = heading.Altitude + 1 + 1
	}

	//
	nthItemAssign.AdvanceNominatorUrgency(1)

	//
	followingParameters := status.AgreementSettings
	finalAltitudeParametersAltered := status.FinalAltitudeAgreementParametersAltered
	if ifaceReply.AgreementArgumentRevisions != nil {
		//
		followingParameters = status.AgreementSettings.Revise(ifaceReply.AgreementArgumentRevisions)
		err := followingParameters.CertifyFundamental()
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}

		err = status.AgreementSettings.CertifyRevise(ifaceReply.AgreementArgumentRevisions, heading.Altitude)
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}

		status.Edition.Agreement.App = followingParameters.Edition.App

		//
		finalAltitudeParametersAltered = heading.Altitude + 1
	}

	followingEdition := status.Edition

	//
	//
	return Status{
		Edition:                          followingEdition,
		SuccessionUUID:                          status.SuccessionUUID,
		PrimaryAltitude:                    status.PrimaryAltitude,
		FinalLedgerAltitude:                  heading.Altitude,
		FinalLedgerUUID:                      ledgerUUID,
		FinalLedgerMoment:                    heading.Moment,
		FollowingAssessors:                   nthItemAssign,
		Assessors:                       status.FollowingAssessors.Duplicate(),
		FinalAssessors:                   status.Assessors.Duplicate(),
		FinalAltitudeAssessorsAltered:      finalAltitudeValuesAltered,
		AgreementSettings:                  followingParameters,
		FinalAltitudeAgreementParametersAltered: finalAltitudeParametersAltered,
		FinalOutcomesDigest:                  TransferOutcomesDigest(ifaceReply.TransferOutcomes),
		PlatformDigest:                          nil,
	}, nil
}

//
//
//
func triggerIncidents(
	tracer log.Tracer,
	incidentPipeline kinds.LedgerIncidentBroadcaster,
	ledger *kinds.Ledger,
	ledgerUUID kinds.LedgerUUID,
	ifaceReply *iface.ReplyCulminateLedger,
	assessorRevisions []*kinds.Assessor,
) {
	if err := incidentPipeline.BroadcastIncidentFreshLedger(kinds.IncidentDataFreshLedger{
		Ledger:               ledger,
		LedgerUUID:             ledgerUUID,
		OutcomeCulminateLedger: *ifaceReply,
	}); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	if err := incidentPipeline.BroadcastIncidentFreshLedgerHeading(kinds.IncidentDataFreshLedgerHeading{
		Heading: ledger.Heading,
	}); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	if err := incidentPipeline.BroadcastIncidentFreshLedgerIncidents(kinds.IncidentDataFreshLedgerIncidents{
		Altitude: ledger.Altitude,
		Incidents: ifaceReply.Incidents,
		CountTrans: int64(len(ledger.Txs)),
	}); err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
	}

	if len(ledger.Proof.Proof) != 0 {
		for _, ev := range ledger.Proof.Proof {
			if err := incidentPipeline.BroadcastIncidentFreshProof(kinds.IncidentDataFreshProof{
				Proof: ev,
				Altitude:   ledger.Altitude,
			}); err != nil {
				tracer.Failure("REDACTED", "REDACTED", err)
			}
		}
	}

	for i, tx := range ledger.Txs {
		if err := incidentPipeline.BroadcastIncidentTransfer(kinds.IncidentDataTransfer{TransferOutcome: iface.TransferOutcome{
			Altitude: ledger.Altitude,
			Ordinal:  uint32(i),
			Tx:     tx,
			Outcome: *(ifaceReply.TransferOutcomes[i]),
		}}); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}
	}

	if len(assessorRevisions) > 0 {
		if err := incidentPipeline.BroadcastIncidentAssessorAssignRevisions(
			kinds.IncidentDataAssessorAssignRevisions{AssessorRevisions: assessorRevisions}); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

//
//

//
//
func InvokeEndorseLedger(
	applicationLinkAgreement delegate.ApplicationLinkAgreement,
	ledger *kinds.Ledger,
	tracer log.Tracer,
	depot Depot,
	primaryAltitude int64,
) ([]byte, error) {
	endorseDetails := constructFinalEndorseDetailsOriginatingDepot(ledger, depot, primaryAltitude)

	reply, err := applicationLinkAgreement.CulminateLedger(context.TODO(), &iface.SolicitCulminateLedger{
		Digest:               ledger.Digest(),
		FollowingAssessorsDigest: ledger.FollowingAssessorsDigest,
		NominatorLocation:    ledger.NominatorLocation,
		Altitude:             ledger.Altitude,
		Moment:               ledger.Moment,
		ResolvedFinalEndorse:  endorseDetails,
		Malpractice:        ledger.Proof.Proof.TowardIface(),
		Txs:                ledger.Txs.TowardSegmentBelongingOctets(),
	})
	if err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return nil, err
	}

	//
	if len(ledger.Txs) != len(reply.TransferOutcomes) {
		return nil, fmt.Errorf("REDACTED", len(ledger.Txs), len(reply.TransferOutcomes))
	}

	tracer.Details("REDACTED", "REDACTED", ledger.Altitude, "REDACTED", fmt.Sprintf("REDACTED", reply.PlatformDigest))

	//
	_, err = applicationLinkAgreement.Endorse(context.TODO())
	if err != nil {
		tracer.Failure("REDACTED", "REDACTED", err)
		return nil, err
	}

	//
	return reply.PlatformDigest, nil
}

func (ledgerExecute *LedgerHandler) trimLedgers(preserveAltitude int64, status Status) (uint64, error) {
	foundation := ledgerExecute.ledgerDepot.Foundation()
	if preserveAltitude <= foundation {
		return 0, nil
	}

	quantityTrimmed, trimmedHeadlineAltitude, err := ledgerExecute.ledgerDepot.TrimLedgers(preserveAltitude, status)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	err = ledgerExecute.Depot().TrimStatuses(foundation, preserveAltitude, trimmedHeadlineAltitude)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}
	return quantityTrimmed, nil
}
