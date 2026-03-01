package status_test

import (
	"bytes"
	"context"
	"fmt"
	"time"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/intrinsic/verify"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

type parametersAlterationVerifyInstance struct {
	altitude int64
	parameters kinds.AgreementSettings
}

func freshVerifyApplication() delegate.PlatformLinks {
	app := &verifyApplication{}
	cc := delegate.FreshRegionalCustomerOriginator(app)
	return delegate.FreshPlatformLinks(cc, delegate.NooperationTelemetry())
}

func createAlsoEndorseValidLedger(
	status sm.Status,
	altitude int64,
	finalEndorse *kinds.Endorse,
	nominatorLocation []byte,
	ledgerExecute *sm.LedgerHandler,
	privateItems map[string]kinds.PrivateAssessor,
	proof []kinds.Proof,
) (sm.Status, kinds.LedgerUUID, *kinds.ExpandedEndorse, error) {
	//
	status, ledgerUUID, err := createAlsoExecuteValidLedger(status, altitude, finalEndorse, nominatorLocation, ledgerExecute, proof)
	if err != nil {
		return status, kinds.LedgerUUID{}, nil, err
	}

	//
	endorse, _, err := createSoundEndorse(altitude, ledgerUUID, status.Assessors, privateItems)
	if err != nil {
		return status, kinds.LedgerUUID{}, nil, err
	}
	return status, ledgerUUID, endorse, nil
}

func createAlsoExecuteValidLedger(status sm.Status, altitude int64, finalEndorse *kinds.Endorse, nominatorLocation []byte,
	ledgerExecute *sm.LedgerHandler, proof []kinds.Proof,
) (sm.Status, kinds.LedgerUUID, error) {
	ledger, err := status.CreateLedger(altitude, verify.CreateNTHTrans(altitude, 10), finalEndorse, proof, nominatorLocation)
	if err != nil {
		return status, kinds.LedgerUUID{}, nil
	}
	fragmentAssign, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	if err != nil {
		return status, kinds.LedgerUUID{}, err
	}

	if err := ledgerExecute.CertifyLedger(status, ledger); err != nil {
		return status, kinds.LedgerUUID{}, err
	}
	ledgerUUID := kinds.LedgerUUID{
		Digest:          ledger.Digest(),
		FragmentAssignHeading: fragmentAssign.Heading(),
	}
	status, err = ledgerExecute.ExecuteLedger(status, ledgerUUID, ledger)
	if err != nil {
		return status, kinds.LedgerUUID{}, err
	}
	return status, ledgerUUID, nil
}

func createLedger(status sm.Status, altitude int64, c *kinds.Endorse) (*kinds.Ledger, error) {
	return status.CreateLedger(
		altitude,
		verify.CreateNTHTrans(status.FinalLedgerAltitude, 10),
		c,
		nil,
		status.Assessors.ObtainNominator().Location,
	)
}

func createSoundEndorse(
	altitude int64,
	ledgerUUID kinds.LedgerUUID,
	values *kinds.AssessorAssign,
	privateItems map[string]kinds.PrivateAssessor,
) (*kinds.ExpandedEndorse, []*kinds.Ballot, error) {
	signatures := make([]kinds.ExpandedEndorseSignature, values.Extent())
	ballots := make([]*kinds.Ballot, values.Extent())
	for i := 0; i < values.Extent(); i++ {
		_, val := values.ObtainViaOrdinal(int32(i))
		ballot, err := kinds.CreateBallot(
			privateItems[val.Location.Text()],
			successionUUID,
			int32(i),
			altitude,
			0,
			commitchema.PreendorseKind,
			ledgerUUID,
			time.Now(),
		)
		if err != nil {
			return nil, nil, err
		}
		signatures[i] = ballot.ExpandedEndorseSignature()
		ballots[i] = ballot
	}
	return &kinds.ExpandedEndorse{
		Altitude:             altitude,
		LedgerUUID:            ledgerUUID,
		ExpandedNotations: signatures,
	}, ballots, nil
}

func createStatus(nthValues, altitude int) (sm.Status, dbm.DB, map[string]kinds.PrivateAssessor) {
	values := make([]kinds.OriginAssessor, nthValues)
	privateItems := make(map[string]kinds.PrivateAssessor, nthValues)
	for i := 0; i < nthValues; i++ {
		credential := []byte(fmt.Sprintf("REDACTED", i))
		pk := edwards25519.ProducePrivateTokenOriginatingCredential(credential)
		itemLocation := pk.PublicToken().Location()
		values[i] = kinds.OriginAssessor{
			Location: itemLocation,
			PublicToken:  pk.PublicToken(),
			Potency:   1000,
			Alias:    fmt.Sprintf("REDACTED", i),
		}
		privateItems[itemLocation.Text()] = kinds.FreshSimulatePRVUsingParameters(pk, false, false)
	}
	s, _ := sm.CreateInaugurationStatus(&kinds.OriginPaper{
		SuccessionUUID:    successionUUID,
		Assessors: values,
		PlatformDigest:    nil,
	})

	statusDatastore := dbm.FreshMemoryDatastore()
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	if err := statusDepot.Persist(s); err != nil {
		panic(err)
	}

	for i := 1; i < altitude; i++ {
		s.FinalLedgerAltitude++
		s.FinalAssessors = s.Assessors.Duplicate()
		if err := statusDepot.Persist(s); err != nil {
			panic(err)
		}
	}

	return s, statusDatastore, privateItems
}

func produceItemCollection(extent int) *kinds.AssessorAssign {
	values := make([]*kinds.Assessor, extent)
	for i := 0; i < extent; i++ {
		values[i] = kinds.FreshAssessor(edwards25519.ProducePrivateToken().PublicToken(), 10)
	}
	return kinds.FreshAssessorAssign(values)
}

func createHeadlineFragmentsRepliesItemPublicTokenAlteration(
	status sm.Status,
	publickey security.PublicToken,
) (kinds.Heading, kinds.LedgerUUID, *iface.ReplyCulminateLedger) {
	ledger, err := createLedger(status, status.FinalLedgerAltitude+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUUID{}, nil
	}
	ifaceReplies := &iface.ReplyCulminateLedger{}
	//
	_, val := status.FollowingAssessors.ObtainViaOrdinal(0)
	if !bytes.Equal(publickey.Octets(), val.PublicToken.Octets()) {
		ifaceReplies.AssessorRevisions = []iface.AssessorRevise{
			kinds.Temp2buffer.FreshAssessorRevise(val.PublicToken, 0),
			kinds.Temp2buffer.FreshAssessorRevise(publickey, 10),
		}
	}

	return ledger.Heading, kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: kinds.FragmentAssignHeading{}}, ifaceReplies
}

func createHeadlineFragmentsRepliesItemPotencyAlteration(
	status sm.Status,
	potency int64,
) (kinds.Heading, kinds.LedgerUUID, *iface.ReplyCulminateLedger) {
	ledger, err := createLedger(status, status.FinalLedgerAltitude+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUUID{}, nil
	}
	ifaceReplies := &iface.ReplyCulminateLedger{}

	//
	_, val := status.FollowingAssessors.ObtainViaOrdinal(0)
	if val.BallotingPotency != potency {
		ifaceReplies.AssessorRevisions = []iface.AssessorRevise{
			kinds.Temp2buffer.FreshAssessorRevise(val.PublicToken, potency),
		}
	}

	return ledger.Heading, kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: kinds.FragmentAssignHeading{}}, ifaceReplies
}

func createHeadlineFragmentsRepliesParameters(
	status sm.Status,
	parameters commitchema.AgreementSettings,
) (kinds.Heading, kinds.LedgerUUID, *iface.ReplyCulminateLedger) {
	ledger, err := createLedger(status, status.FinalLedgerAltitude+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUUID{}, nil
	}
	ifaceReplies := &iface.ReplyCulminateLedger{
		AgreementArgumentRevisions: &parameters,
	}
	return ledger.Heading, kinds.LedgerUUID{Digest: ledger.Digest(), FragmentAssignHeading: kinds.FragmentAssignHeading{}}, ifaceReplies
}

func unpredictableInaugurationPaper() *kinds.OriginPaper {
	publickey := edwards25519.ProducePrivateToken().PublicToken()
	return &kinds.OriginPaper{
		OriginMoment: committime.Now(),
		SuccessionUUID:     "REDACTED",
		Assessors: []kinds.OriginAssessor{
			{
				Location: publickey.Location(),
				PublicToken:  publickey,
				Potency:   10,
				Alias:    "REDACTED",
			},
		},
		AgreementSettings: kinds.FallbackAgreementSettings(),
	}
}

//

type verifyApplication struct {
	iface.FoundationPlatform

	EndorseBallots      []iface.BallotDetails
	Malpractice      []iface.Malpractice
	FinalMoment         time.Time
	AssessorRevisions []iface.AssessorRevise
	PlatformDigest          []byte
}

var _ iface.Platform = (*verifyApplication)(nil)

func (app *verifyApplication) CulminateLedger(_ context.Context, req *iface.SolicitCulminateLedger) (*iface.ReplyCulminateLedger, error) {
	app.EndorseBallots = req.ResolvedFinalEndorse.Ballots
	app.Malpractice = req.Malpractice
	app.FinalMoment = req.Moment
	transferOutcomes := make([]*iface.InvokeTransferOutcome, len(req.Txs))
	for idx := range req.Txs {
		transferOutcomes[idx] = &iface.InvokeTransferOutcome{
			Cipher: iface.CipherKindOKAY,
		}
	}

	return &iface.ReplyCulminateLedger{
		AssessorRevisions: app.AssessorRevisions,
		AgreementArgumentRevisions: &commitchema.AgreementSettings{
			Edition: &commitchema.EditionParameters{
				App: 1,
			},
		},
		TransferOutcomes: transferOutcomes,
		PlatformDigest:   app.PlatformDigest,
	}, nil
}

func (app *verifyApplication) Endorse(_ context.Context, _ *iface.SolicitEndorse) (*iface.ReplyEndorse, error) {
	return &iface.ReplyEndorse{PreserveAltitude: 1}, nil
}

func (app *verifyApplication) ArrangeNomination(
	_ context.Context,
	req *iface.SolicitArrangeNomination,
) (*iface.ReplyArrangeNomination, error) {
	txs := make([][]byte, 0, len(req.Txs))
	var sumOctets int64
	for _, tx := range req.Txs {
		if len(tx) == 0 {
			continue
		}
		sumOctets += int64(len(tx))
		if sumOctets > req.MaximumTransferOctets {
			break
		}
		txs = append(txs, tx)
	}
	return &iface.ReplyArrangeNomination{Txs: txs}, nil
}

func (app *verifyApplication) HandleNomination(
	_ context.Context,
	req *iface.SolicitHandleNomination,
) (*iface.ReplyHandleNomination, error) {
	for _, tx := range req.Txs {
		if len(tx) == 0 {
			return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_DECLINE}, nil
		}
	}
	return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil
}
