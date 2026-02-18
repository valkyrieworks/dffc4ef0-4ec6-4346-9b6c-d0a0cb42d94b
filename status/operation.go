package status

import (
	"bytes"
	"context"
	"fmt"
	"time"

	iface "github.com/valkyrieworks/iface/kinds"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/utils/abort"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/txpool"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//

//
type LedgerRunner struct {
	//
	depot Depot

	//
	ledgerDepot LedgerDepot

	//
	gatewayApplication gateway.ApplicationLinkAgreement

	//
	eventBus kinds.LedgerEventBroadcaster

	//
	//
	txpool txpool.Txpool
	eventpool  ProofDepository

	//
	finalCertifiedLedger *kinds.Ledger

	tracer log.Tracer

	stats *Stats
}

type LedgerRunnerSetting func(runner *LedgerRunner)

func LedgerRunnerWithStats(stats *Stats) LedgerRunnerSetting {
	return func(ledgerExecute *LedgerRunner) {
		ledgerExecute.stats = stats
	}
}

//
//
func NewLedgerRunner(
	statusDepot Depot,
	tracer log.Tracer,
	gatewayApplication gateway.ApplicationLinkAgreement,
	txpool txpool.Txpool,
	eventpool ProofDepository,
	ledgerDepot LedgerDepot,
	options ...LedgerRunnerSetting,
) *LedgerRunner {
	res := &LedgerRunner{
		depot:      statusDepot,
		gatewayApplication:   gatewayApplication,
		eventBus:   kinds.NoopEventBus{},
		txpool:    txpool,
		eventpool:     eventpool,
		tracer:     tracer,
		stats:    NoopStats(),
		ledgerDepot: ledgerDepot,
	}

	for _, setting := range options {
		setting(res)
	}

	return res
}

func (ledgerExecute *LedgerRunner) Depot() Depot {
	return ledgerExecute.depot
}

//
//
func (ledgerExecute *LedgerRunner) AssignEventBus(eventBus kinds.LedgerEventBroadcaster) {
	ledgerExecute.eventBus = eventBus
}

//
//
//
//
//
//
func (ledgerExecute *LedgerRunner) InstantiateNominationLedger(
	ctx context.Context,
	level int64,
	status Status,
	finalExtensionEndorse *kinds.ExpandedEndorse,
	recommenderAddress []byte,
) (*kinds.Ledger, error) {
	maximumOctets := status.AgreementOptions.Ledger.MaximumOctets
	emptyMaximumOctets := maximumOctets == -1
	if emptyMaximumOctets {
		maximumOctets = int64(kinds.MaximumLedgerVolumeOctets)
	}

	maximumFuel := status.AgreementOptions.Ledger.MaximumFuel

	proof, evtVolume := ledgerExecute.eventpool.AwaitingProof(status.AgreementOptions.Proof.MaximumOctets)

	//
	maximumDataOctets := kinds.MaximumDataOctets(maximumOctets, evtVolume, status.Ratifiers.Volume())
	maximumHarvestOctets := maximumDataOctets
	if emptyMaximumOctets {
		maximumHarvestOctets = -1
	}

	txs := ledgerExecute.txpool.HarvestMaximumOctetsMaximumFuel(maximumHarvestOctets, maximumFuel)
	endorse := finalExtensionEndorse.ToEndorse()
	ledger, err := status.CreateLedger(level, txs, endorse, proof, recommenderAddress)
	if err != nil {
		return nil, err
	}
	rpp, err := ledgerExecute.gatewayApplication.ArrangeNomination(
		ctx,
		&iface.QueryArrangeNomination{
			MaximumTransferOctets:         maximumDataOctets,
			Txs:                ledger.Txs.ToSegmentOfOctets(),
			NativeFinalEndorse:    constructExpandedEndorseDetailsFromDepot(finalExtensionEndorse, ledgerExecute.depot, status.PrimaryLevel, status.AgreementOptions.Iface),
			Malpractice:        ledger.Proof.Proof.ToIface(),
			Level:             ledger.Level,
			Time:               ledger.Time,
			FollowingRatifiersDigest: ledger.FollowingRatifiersDigest,
			RecommenderLocation:    ledger.RecommenderLocation,
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

	txl := kinds.ToTrans(rpp.Txs)
	if err := txl.Certify(maximumDataOctets); err != nil {
		return nil, err
	}

	return status.CreateLedger(level, txl, endorse, proof, recommenderAddress)
}

func (ledgerExecute *LedgerRunner) HandleNomination(
	ledger *kinds.Ledger,
	status Status,
) (bool, error) {
	reply, err := ledgerExecute.gatewayApplication.HandleNomination(context.TODO(), &iface.QueryHandleNomination{
		Digest:               ledger.Heading.Digest(),
		Level:             ledger.Level,
		Time:               ledger.Time,
		Txs:                ledger.Txs.ToSegmentOfOctets(),
		NominatedFinalEndorse: constructFinalEndorseDetailsFromDepot(ledger, ledgerExecute.depot, status.PrimaryLevel),
		Malpractice:        ledger.Proof.Proof.ToIface(),
		RecommenderLocation:    ledger.RecommenderLocation,
		FollowingRatifiersDigest: ledger.FollowingRatifiersDigest,
	})
	if err != nil {
		return false, err
	}
	if reply.IsStateUnclear() {
		panic(fmt.Sprintf("REDACTED", reply.Status.String()))
	}

	return reply.IsApproved(), nil
}

//
//
//
//
func (ledgerExecute *LedgerRunner) CertifyLedger(status Status, ledger *kinds.Ledger) error {
	if !ledgerExecute.finalCertifiedLedger.DigestsTo(ledger.Digest()) {
		if err := certifyLedger(status, ledger); err != nil {
			return err
		}
		ledgerExecute.finalCertifiedLedger = ledger
	}
	return ledgerExecute.eventpool.InspectProof(ledger.Proof.Proof)
}

//
func (ledgerExecute *LedgerRunner) ExecuteValidatedLedger(
	status Status, ledgerUID kinds.LedgerUID, ledger *kinds.Ledger,
) (Status, error) {
	return ledgerExecute.executeLedger(status, ledgerUID, ledger)
}

//
//
//
//
//
//
func (ledgerExecute *LedgerRunner) ExecuteLedger(
	status Status, ledgerUID kinds.LedgerUID, ledger *kinds.Ledger,
) (Status, error) {
	if !ledgerExecute.finalCertifiedLedger.DigestsTo(ledger.Digest()) {
		if err := certifyLedger(status, ledger); err != nil {
			return status, ErrCorruptLedger(err)
		}
		ledgerExecute.finalCertifiedLedger = ledger
	}
	return ledgerExecute.executeLedger(status, ledgerUID, ledger)
}

func (ledgerExecute *LedgerRunner) executeLedger(status Status, ledgerUID kinds.LedgerUID, ledger *kinds.Ledger) (Status, error) {
	beginMoment := time.Now().UnixNano()
	ifaceReply, err := ledgerExecute.gatewayApplication.CompleteLedger(context.TODO(), &iface.QueryCompleteLedger{
		Digest:               ledger.Digest(),
		FollowingRatifiersDigest: ledger.FollowingRatifiersDigest,
		RecommenderLocation:    ledger.RecommenderLocation,
		Level:             ledger.Level,
		Time:               ledger.Time,
		ResolvedFinalEndorse:  constructFinalEndorseDetailsFromDepot(ledger, ledgerExecute.depot, status.PrimaryLevel),
		Malpractice:        ledger.Proof.Proof.ToIface(),
		Txs:                ledger.Txs.ToSegmentOfOctets(),
	})
	terminateTime := time.Now().UnixNano()
	ledgerExecute.stats.LedgerExecutionTime.Observe(float64(terminateTime-beginMoment) / 1000000)
	if err != nil {
		ledgerExecute.tracer.Fault("REDACTED", "REDACTED", err)
		return status, err
	}

	ledgerExecute.tracer.Details(
		"REDACTED",
		"REDACTED", ledger.Level,
		"REDACTED", len(ifaceReply.TransOutcomes),
		"REDACTED", len(ifaceReply.RatifierRefreshes),
		"REDACTED", fmt.Sprintf("REDACTED", ifaceReply.ApplicationDigest),
	)

	//
	if len(ledger.Txs) != len(ifaceReply.TransOutcomes) {
		return status, fmt.Errorf("REDACTED", len(ledger.Txs), len(ifaceReply.TransOutcomes))
	}

	ledgerExecute.tracer.Details("REDACTED", "REDACTED", ledger.Level, "REDACTED", fmt.Sprintf("REDACTED", ifaceReply.ApplicationDigest))

	abort.Abort() //

	//
	if err := ledgerExecute.depot.PersistCompleteLedgerReply(ledger.Level, ifaceReply); err != nil {
		return status, err
	}

	abort.Abort() //

	//
	err = certifyRatifierRefreshes(ifaceReply.RatifierRefreshes, status.AgreementOptions.Ratifier)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	ratifierRefreshes, err := kinds.Schema2tm.RatifierRefreshes(ifaceReply.RatifierRefreshes)
	if err != nil {
		return status, err
	}
	if len(ratifierRefreshes) > 0 {
		ledgerExecute.tracer.Details("REDACTED", "REDACTED", kinds.RatifierCatalogString(ratifierRefreshes))
		ledgerExecute.stats.RatifierCollectionRefreshes.Add(1)
	}
	if ifaceReply.AgreementArgumentRefreshes != nil {
		ledgerExecute.stats.AgreementArgumentRefreshes.Add(1)
	}

	//
	status, err = modifyStatus(status, ledgerUID, &ledger.Heading, ifaceReply, ratifierRefreshes)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	//
	preserveLevel, err := ledgerExecute.Endorse(status, ledger, ifaceReply)
	if err != nil {
		return status, fmt.Errorf("REDACTED", err)
	}

	//
	ledgerExecute.eventpool.Modify(status, ledger.Proof.Proof)

	abort.Abort() //

	//
	status.ApplicationDigest = ifaceReply.ApplicationDigest
	if err := ledgerExecute.depot.Persist(status); err != nil {
		return status, err
	}

	abort.Abort() //

	//
	if preserveLevel > 0 {
		trimmed, err := ledgerExecute.trimLedgers(preserveLevel, status)
		if err != nil {
			ledgerExecute.tracer.Fault("REDACTED", "REDACTED", preserveLevel, "REDACTED", err)
		} else {
			ledgerExecute.tracer.Diagnose("REDACTED", "REDACTED", trimmed, "REDACTED", preserveLevel)
		}
	}

	//
	//
	triggerEvents(ledgerExecute.tracer, ledgerExecute.eventBus, ledger, ledgerUID, ifaceReply, ratifierRefreshes)

	return status, nil
}

func (ledgerExecute *LedgerRunner) ExpandBallot(
	ctx context.Context,
	ballot *kinds.Ballot,
	ledger *kinds.Ledger,
	status Status,
) ([]byte, error) {
	if !ledger.DigestsTo(ballot.LedgerUID.Digest) {
		panic(fmt.Sprintf("REDACTED", ledger.Digest(), ballot.LedgerUID.Digest))
	}
	if ballot.Level != ledger.Level {
		panic(fmt.Sprintf("REDACTED", ledger.Level, ballot.Level))
	}

	req := iface.QueryExpandBallot{
		Digest:               ballot.LedgerUID.Digest,
		Level:             ballot.Level,
		Time:               ledger.Time,
		Txs:                ledger.Txs.ToSegmentOfOctets(),
		NominatedFinalEndorse: constructFinalEndorseDetailsFromDepot(ledger, ledgerExecute.depot, status.PrimaryLevel),
		Malpractice:        ledger.Proof.Proof.ToIface(),
		FollowingRatifiersDigest: ledger.FollowingRatifiersDigest,
		RecommenderLocation:    ledger.RecommenderLocation,
	}

	reply, err := ledgerExecute.gatewayApplication.ExpandBallot(ctx, &req)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return reply.BallotAddition, nil
}

func (ledgerExecute *LedgerRunner) ValidateBallotAddition(ctx context.Context, ballot *kinds.Ballot) error {
	req := iface.QueryValidateBallotAddition{
		Digest:             ballot.LedgerUID.Digest,
		RatifierLocation: ballot.RatifierLocation,
		Level:           ballot.Level,
		BallotAddition:    ballot.Addition,
	}

	reply, err := ledgerExecute.gatewayApplication.ValidateBallotAddition(ctx, &req)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if reply.IsStateUnclear() {
		panic(fmt.Sprintf("REDACTED", reply.Status.String()))
	}

	if !reply.IsApproved() {
		return kinds.ErrCorruptBallotAddition
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
func (ledgerExecute *LedgerRunner) Endorse(
	status Status,
	ledger *kinds.Ledger,
	ifaceReply *iface.ReplyCompleteLedger,
) (int64, error) {
	ledgerExecute.txpool.Secure()
	unsealTxpool := func() { ledgerExecute.txpool.Release() }

	//
	//
	err := ledgerExecute.txpool.PurgeApplicationLink()
	if err != nil {
		unsealTxpool()
		ledgerExecute.tracer.Fault("REDACTED", "REDACTED", err)
		return 0, err
	}

	//
	res, err := ledgerExecute.gatewayApplication.Endorse(context.TODO())
	if err != nil {
		unsealTxpool()
		ledgerExecute.tracer.Fault("REDACTED", "REDACTED", err)
		return 0, err
	}

	//
	ledgerExecute.tracer.Details(
		"REDACTED",
		"REDACTED", ledger.Level,
		"REDACTED", fmt.Sprintf("REDACTED", ledger.ApplicationDigest),
	)

	//
	go ledgerExecute.asyncModifyTxpool(unsealTxpool, ledger, status.Clone(), ifaceReply)

	return res.PreserveLevel, nil
}

//
func (ledgerExecute *LedgerRunner) asyncModifyTxpool(
	unsealTxpool func(),
	ledger *kinds.Ledger,
	status Status,
	ifaceReply *iface.ReplyCompleteLedger,
) {
	defer unsealTxpool()

	err := ledgerExecute.txpool.Modify(
		ledger.Level,
		ledger.Txs,
		ifaceReply.TransOutcomes,
		TransferPreInspect(status),
		TransferSubmitInspect(status),
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

func constructFinalEndorseDetailsFromDepot(ledger *kinds.Ledger, depot Depot, primaryLevel int64) iface.EndorseDetails {
	if ledger.Level == primaryLevel { //
		//
		//
		return iface.EndorseDetails{}
	}

	finalValueCollection, err := depot.ImportRatifiers(ledger.Level - 1)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ledger.Level-1, err))
	}

	return ConstructFinalEndorseDetails(ledger, finalValueCollection, primaryLevel)
}

//
//
//
func ConstructFinalEndorseDetails(ledger *kinds.Ledger, finalValueCollection *kinds.RatifierAssign, primaryLevel int64) iface.EndorseDetails {
	if ledger.Level == primaryLevel {
		//
		//
		return iface.EndorseDetails{}
	}

	var (
		endorseVolume = ledger.FinalEndorse.Volume()
		valueCollectionSize  = len(finalValueCollection.Ratifiers)
	)

	//
	//
	if endorseVolume != valueCollectionSize {
		panic(fmt.Sprintf(
			"REDACTED",
			endorseVolume, valueCollectionSize, ledger.Level, ledger.FinalEndorse.Endorsements, finalValueCollection.Ratifiers,
		))
	}

	ballots := make([]iface.BallotDetails, ledger.FinalEndorse.Volume())
	for i, val := range finalValueCollection.Ratifiers {
		endorseSignature := ledger.FinalEndorse.Endorsements[i]
		ballots[i] = iface.BallotDetails{
			Ratifier:   kinds.Tm2schema.Ratifier(val),
			LedgerUidMark: engineproto.LedgerUIDMark(endorseSignature.LedgerUIDMark),
		}
	}

	return iface.EndorseDetails{
		Cycle: ledger.FinalEndorse.Cycle,
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
func constructExpandedEndorseDetailsFromDepot(ec *kinds.ExpandedEndorse, depot Depot, primaryLevel int64, ap kinds.IfaceOptions) iface.ExpandedEndorseDetails {
	if ec.Level < primaryLevel {
		//
		return iface.ExpandedEndorseDetails{}
	}

	valueCollection, err := depot.ImportRatifiers(ec.Level)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ec.Level, primaryLevel, err))
	}

	return ConstructExpandedEndorseDetails(ec, valueCollection, primaryLevel, ap)
}

//
//
//
func ConstructExpandedEndorseDetails(ec *kinds.ExpandedEndorse, valueCollection *kinds.RatifierAssign, primaryLevel int64, ap kinds.IfaceOptions) iface.ExpandedEndorseDetails {
	if ec.Level < primaryLevel {
		//
		return iface.ExpandedEndorseDetails{}
	}

	var (
		ecVolume    = ec.Volume()
		valueCollectionSize = len(valueCollection.Ratifiers)
	)

	//
	//
	if ecVolume != valueCollectionSize {
		panic(fmt.Errorf(
			"REDACTED",
			ecVolume, valueCollectionSize, ec.Level, ec.ExpandedEndorsements, valueCollection.Ratifiers,
		))
	}

	ballots := make([]iface.ExpandedBallotDetails, ecVolume)
	for i, val := range valueCollection.Ratifiers {
		ecs := ec.ExpandedEndorsements[i]

		//
		//
		if ecs.LedgerUIDMark != kinds.LedgerUIDMarkMissing && !bytes.Equal(ecs.RatifierLocation, val.Location) {
			panic(fmt.Errorf("REDACTED",
				i, ecs.RatifierLocation, ec.Level, val.Location,
			))
		}

		//
		//
		//
		//
		//
		if err := ecs.AssureAddition(ap.BallotPluginsActivated(ec.Level)); err != nil {
			panic(fmt.Errorf("REDACTED", ec.Level, err))
		}

		ballots[i] = iface.ExpandedBallotDetails{
			Ratifier:          kinds.Tm2schema.Ratifier(val),
			LedgerUidMark:        engineproto.LedgerUIDMark(ecs.LedgerUIDMark),
			BallotAddition:      ecs.Addition,
			AdditionAutograph: ecs.AdditionAutograph,
		}
	}

	return iface.ExpandedEndorseDetails{
		Cycle: ec.Cycle,
		Ballots: ballots,
	}
}

func certifyRatifierRefreshes(ifaceRefreshes []iface.RatifierModify,
	options kinds.RatifierOptions,
) error {
	for _, valueModify := range ifaceRefreshes {
		if valueModify.FetchEnergy() < 0 {
			return fmt.Errorf("REDACTED", valueModify)
		} else if valueModify.FetchEnergy() == 0 {
			//
			//
			continue
		}

		//
		pk, err := cryptocode.PublicKeyFromSchema(valueModify.PublicKey)
		if err != nil {
			return err
		}

		if !kinds.IsSoundPublickeyKind(options, pk.Kind()) {
			return fmt.Errorf("REDACTED",
				valueModify, pk.Kind())
		}
	}
	return nil
}

//
func modifyStatus(
	status Status,
	ledgerUID kinds.LedgerUID,
	heading *kinds.Heading,
	ifaceReply *iface.ReplyCompleteLedger,
	ratifierRefreshes []*kinds.Ratifier,
) (Status, error) {
	//
	//
	nValueCollection := status.FollowingRatifiers.Clone()

	//
	finalLevelValuesModified := status.FinalLevelRatifiersModified
	if len(ratifierRefreshes) > 0 {
		err := nValueCollection.ModifyWithAlterCollection(ratifierRefreshes)
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}
		//
		finalLevelValuesModified = heading.Level + 1 + 1
	}

	//
	nValueCollection.AugmentRecommenderUrgency(1)

	//
	followingOptions := status.AgreementOptions
	finalLevelOptionsModified := status.FinalLevelAgreementOptionsModified
	if ifaceReply.AgreementArgumentRefreshes != nil {
		//
		followingOptions = status.AgreementOptions.Modify(ifaceReply.AgreementArgumentRefreshes)
		err := followingOptions.CertifySimple()
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}

		err = status.AgreementOptions.CertifyModify(ifaceReply.AgreementArgumentRefreshes, heading.Level)
		if err != nil {
			return status, fmt.Errorf("REDACTED", err)
		}

		status.Release.Agreement.App = followingOptions.Release.App

		//
		finalLevelOptionsModified = heading.Level + 1
	}

	followingRelease := status.Release

	//
	//
	return Status{
		Release:                          followingRelease,
		LedgerUID:                          status.LedgerUID,
		PrimaryLevel:                    status.PrimaryLevel,
		FinalLedgerLevel:                  heading.Level,
		FinalLedgerUID:                      ledgerUID,
		FinalLedgerTime:                    heading.Time,
		FollowingRatifiers:                   nValueCollection,
		Ratifiers:                       status.FollowingRatifiers.Clone(),
		FinalRatifiers:                   status.Ratifiers.Clone(),
		FinalLevelRatifiersModified:      finalLevelValuesModified,
		AgreementOptions:                  followingOptions,
		FinalLevelAgreementOptionsModified: finalLevelOptionsModified,
		FinalOutcomesDigest:                  TransferOutcomesDigest(ifaceReply.TransOutcomes),
		ApplicationDigest:                          nil,
	}, nil
}

//
//
//
func triggerEvents(
	tracer log.Tracer,
	eventBus kinds.LedgerEventBroadcaster,
	ledger *kinds.Ledger,
	ledgerUID kinds.LedgerUID,
	ifaceReply *iface.ReplyCompleteLedger,
	ratifierRefreshes []*kinds.Ratifier,
) {
	if err := eventBus.BroadcastEventNewLedger(kinds.EventDataNewLedger{
		Ledger:               ledger,
		LedgerUID:             ledgerUID,
		OutcomeCompleteLedger: *ifaceReply,
	}); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	if err := eventBus.BroadcastEventNewLedgerHeading(kinds.EventDataNewLedgerHeading{
		Heading: ledger.Heading,
	}); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	if err := eventBus.BroadcastEventNewLedgerEvents(kinds.EventDataNewLedgerEvents{
		Level: ledger.Level,
		Events: ifaceReply.Events,
		CountTrans: int64(len(ledger.Txs)),
	}); err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
	}

	if len(ledger.Proof.Proof) != 0 {
		for _, ev := range ledger.Proof.Proof {
			if err := eventBus.BroadcastEventNewProof(kinds.EventDataNewProof{
				Proof: ev,
				Level:   ledger.Level,
			}); err != nil {
				tracer.Fault("REDACTED", "REDACTED", err)
			}
		}
	}

	for i, tx := range ledger.Txs {
		if err := eventBus.BroadcastEventTransfer(kinds.EventDataTransfer{TransOutcome: iface.TransOutcome{
			Level: ledger.Level,
			Ordinal:  uint32(i),
			Tx:     tx,
			Outcome: *(ifaceReply.TransOutcomes[i]),
		}}); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	}

	if len(ratifierRefreshes) > 0 {
		if err := eventBus.BroadcastEventRatifierCollectionRefreshes(
			kinds.EventDataRatifierCollectionRefreshes{RatifierRefreshes: ratifierRefreshes}); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

//
//

//
//
func InvokeEndorseLedger(
	applicationLinkAgreement gateway.ApplicationLinkAgreement,
	ledger *kinds.Ledger,
	tracer log.Tracer,
	depot Depot,
	primaryLevel int64,
) ([]byte, error) {
	endorseDetails := constructFinalEndorseDetailsFromDepot(ledger, depot, primaryLevel)

	reply, err := applicationLinkAgreement.CompleteLedger(context.TODO(), &iface.QueryCompleteLedger{
		Digest:               ledger.Digest(),
		FollowingRatifiersDigest: ledger.FollowingRatifiersDigest,
		RecommenderLocation:    ledger.RecommenderLocation,
		Level:             ledger.Level,
		Time:               ledger.Time,
		ResolvedFinalEndorse:  endorseDetails,
		Malpractice:        ledger.Proof.Proof.ToIface(),
		Txs:                ledger.Txs.ToSegmentOfOctets(),
	})
	if err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return nil, err
	}

	//
	if len(ledger.Txs) != len(reply.TransOutcomes) {
		return nil, fmt.Errorf("REDACTED", len(ledger.Txs), len(reply.TransOutcomes))
	}

	tracer.Details("REDACTED", "REDACTED", ledger.Level, "REDACTED", fmt.Sprintf("REDACTED", reply.ApplicationDigest))

	//
	_, err = applicationLinkAgreement.Endorse(context.TODO())
	if err != nil {
		tracer.Fault("REDACTED", "REDACTED", err)
		return nil, err
	}

	//
	return reply.ApplicationDigest, nil
}

func (ledgerExecute *LedgerRunner) trimLedgers(preserveLevel int64, status Status) (uint64, error) {
	root := ledgerExecute.ledgerDepot.Root()
	if preserveLevel <= root {
		return 0, nil
	}

	quantityTrimmed, trimmedHeadingLevel, err := ledgerExecute.ledgerDepot.TrimLedgers(preserveLevel, status)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	err = ledgerExecute.Depot().TrimConditions(root, preserveLevel, trimmedHeadingLevel)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}
	return quantityTrimmed, nil
}
