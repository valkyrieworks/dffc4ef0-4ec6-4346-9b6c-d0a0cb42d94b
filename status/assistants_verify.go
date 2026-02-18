package status_test

import (
	"bytes"
	"context"
	"fmt"
	"time"

	dbm "github.com/valkyrieworks/-db"

	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/intrinsic/verify"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

type optionsAlterVerifyScenario struct {
	level int64
	options kinds.AgreementOptions
}

func newVerifyApplication() gateway.ApplicationLinks {
	app := &verifyApplication{}
	cc := gateway.NewNativeCustomerOriginator(app)
	return gateway.NewApplicationLinks(cc, gateway.NoopStats())
}

func createAndEndorseValidLedger(
	status sm.Status,
	level int64,
	finalEndorse *kinds.Endorse,
	recommenderAddress []byte,
	ledgerExecute *sm.LedgerRunner,
	privateValues map[string]kinds.PrivateRatifier,
	proof []kinds.Proof,
) (sm.Status, kinds.LedgerUID, *kinds.ExpandedEndorse, error) {
	//
	status, ledgerUID, err := createAndExecuteValidLedger(status, level, finalEndorse, recommenderAddress, ledgerExecute, proof)
	if err != nil {
		return status, kinds.LedgerUID{}, nil, err
	}

	//
	endorse, _, err := createSoundEndorse(level, ledgerUID, status.Ratifiers, privateValues)
	if err != nil {
		return status, kinds.LedgerUID{}, nil, err
	}
	return status, ledgerUID, endorse, nil
}

func createAndExecuteValidLedger(status sm.Status, level int64, finalEndorse *kinds.Endorse, recommenderAddress []byte,
	ledgerExecute *sm.LedgerRunner, proof []kinds.Proof,
) (sm.Status, kinds.LedgerUID, error) {
	ledger, err := status.CreateLedger(level, verify.CreateNTrans(level, 10), finalEndorse, proof, recommenderAddress)
	if err != nil {
		return status, kinds.LedgerUID{}, nil
	}
	sectionCollection, err := ledger.CreateSegmentAssign(kinds.LedgerSegmentVolumeOctets)
	if err != nil {
		return status, kinds.LedgerUID{}, err
	}

	if err := ledgerExecute.CertifyLedger(status, ledger); err != nil {
		return status, kinds.LedgerUID{}, err
	}
	ledgerUID := kinds.LedgerUID{
		Digest:          ledger.Digest(),
		SegmentAssignHeading: sectionCollection.Heading(),
	}
	status, err = ledgerExecute.ExecuteLedger(status, ledgerUID, ledger)
	if err != nil {
		return status, kinds.LedgerUID{}, err
	}
	return status, ledgerUID, nil
}

func createLedger(status sm.Status, level int64, c *kinds.Endorse) (*kinds.Ledger, error) {
	return status.CreateLedger(
		level,
		verify.CreateNTrans(status.FinalLedgerLevel, 10),
		c,
		nil,
		status.Ratifiers.FetchRecommender().Location,
	)
}

func createSoundEndorse(
	level int64,
	ledgerUID kinds.LedgerUID,
	values *kinds.RatifierAssign,
	privateValues map[string]kinds.PrivateRatifier,
) (*kinds.ExpandedEndorse, []*kinds.Ballot, error) {
	autographs := make([]kinds.ExpandedEndorseSignature, values.Volume())
	ballots := make([]*kinds.Ballot, values.Volume())
	for i := 0; i < values.Volume(); i++ {
		_, val := values.FetchByOrdinal(int32(i))
		ballot, err := kinds.CreateBallot(
			privateValues[val.Location.String()],
			ledgerUID,
			int32(i),
			level,
			0,
			engineproto.PreendorseKind,
			ledgerUID,
			time.Now(),
		)
		if err != nil {
			return nil, nil, err
		}
		autographs[i] = ballot.ExpandedEndorseSignature()
		ballots[i] = ballot
	}
	return &kinds.ExpandedEndorse{
		Level:             level,
		LedgerUID:            ledgerUID,
		ExpandedEndorsements: autographs,
	}, ballots, nil
}

func createStatus(nValues, level int) (sm.Status, dbm.DB, map[string]kinds.PrivateRatifier) {
	values := make([]kinds.OriginRatifier, nValues)
	privateValues := make(map[string]kinds.PrivateRatifier, nValues)
	for i := 0; i < nValues; i++ {
		key := []byte(fmt.Sprintf("REDACTED", i))
		pk := ed25519.GeneratePrivateKeyFromPrivatekey(key)
		valueAddress := pk.PublicKey().Location()
		values[i] = kinds.OriginRatifier{
			Location: valueAddress,
			PublicKey:  pk.PublicKey(),
			Energy:   1000,
			Label:    fmt.Sprintf("REDACTED", i),
		}
		privateValues[valueAddress.String()] = kinds.NewEmulatePVWithOptions(pk, false, false)
	}
	s, _ := sm.CreateOriginStatus(&kinds.OriginPaper{
		LedgerUID:    ledgerUID,
		Ratifiers: values,
		ApplicationDigest:    nil,
	})

	statusStore := dbm.NewMemoryStore()
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	if err := statusDepot.Persist(s); err != nil {
		panic(err)
	}

	for i := 1; i < level; i++ {
		s.FinalLedgerLevel++
		s.FinalRatifiers = s.Ratifiers.Clone()
		if err := statusDepot.Persist(s); err != nil {
			panic(err)
		}
	}

	return s, statusStore, privateValues
}

func generateValueCollection(volume int) *kinds.RatifierAssign {
	values := make([]*kinds.Ratifier, volume)
	for i := 0; i < volume; i++ {
		values[i] = kinds.NewRatifier(ed25519.GeneratePrivateKey().PublicKey(), 10)
	}
	return kinds.NewRatifierCollection(values)
}

func createHeadingSectionsRepliesValuePublicKeyAlter(
	status sm.Status,
	publickey vault.PublicKey,
) (kinds.Heading, kinds.LedgerUID, *iface.ReplyCompleteLedger) {
	ledger, err := createLedger(status, status.FinalLedgerLevel+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUID{}, nil
	}
	ifaceReplies := &iface.ReplyCompleteLedger{}
	//
	_, val := status.FollowingRatifiers.FetchByOrdinal(0)
	if !bytes.Equal(publickey.Octets(), val.PublicKey.Octets()) {
		ifaceReplies.RatifierRefreshes = []iface.RatifierModify{
			kinds.Tm2schema.NewRatifierModify(val.PublicKey, 0),
			kinds.Tm2schema.NewRatifierModify(publickey, 10),
		}
	}

	return ledger.Heading, kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: kinds.SegmentAssignHeading{}}, ifaceReplies
}

func createHeadingSectionsRepliesValueEnergyAlter(
	status sm.Status,
	energy int64,
) (kinds.Heading, kinds.LedgerUID, *iface.ReplyCompleteLedger) {
	ledger, err := createLedger(status, status.FinalLedgerLevel+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUID{}, nil
	}
	ifaceReplies := &iface.ReplyCompleteLedger{}

	//
	_, val := status.FollowingRatifiers.FetchByOrdinal(0)
	if val.PollingEnergy != energy {
		ifaceReplies.RatifierRefreshes = []iface.RatifierModify{
			kinds.Tm2schema.NewRatifierModify(val.PublicKey, energy),
		}
	}

	return ledger.Heading, kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: kinds.SegmentAssignHeading{}}, ifaceReplies
}

func createHeadingSectionsRepliesOptions(
	status sm.Status,
	options engineproto.AgreementOptions,
) (kinds.Heading, kinds.LedgerUID, *iface.ReplyCompleteLedger) {
	ledger, err := createLedger(status, status.FinalLedgerLevel+1, new(kinds.Endorse))
	if err != nil {
		return kinds.Heading{}, kinds.LedgerUID{}, nil
	}
	ifaceReplies := &iface.ReplyCompleteLedger{
		AgreementArgumentRefreshes: &options,
	}
	return ledger.Heading, kinds.LedgerUID{Digest: ledger.Digest(), SegmentAssignHeading: kinds.SegmentAssignHeading{}}, ifaceReplies
}

func arbitraryOriginPaper() *kinds.OriginPaper {
	publickey := ed25519.GeneratePrivateKey().PublicKey()
	return &kinds.OriginPaper{
		OriginMoment: engineclock.Now(),
		LedgerUID:     "REDACTED",
		Ratifiers: []kinds.OriginRatifier{
			{
				Location: publickey.Location(),
				PublicKey:  publickey,
				Energy:   10,
				Label:    "REDACTED",
			},
		},
		AgreementOptions: kinds.StandardAgreementOptions(),
	}
}

//

type verifyApplication struct {
	iface.RootSoftware

	EndorseBallots      []iface.BallotDetails
	Malpractice      []iface.Malpractice
	FinalTime         time.Time
	RatifierRefreshes []iface.RatifierModify
	ApplicationDigest          []byte
}

var _ iface.Software = (*verifyApplication)(nil)

func (app *verifyApplication) CompleteLedger(_ context.Context, req *iface.QueryCompleteLedger) (*iface.ReplyCompleteLedger, error) {
	app.EndorseBallots = req.ResolvedFinalEndorse.Ballots
	app.Malpractice = req.Malpractice
	app.FinalTime = req.Time
	transferOutcomes := make([]*iface.InvokeTransferOutcome, len(req.Txs))
	for idx := range req.Txs {
		transferOutcomes[idx] = &iface.InvokeTransferOutcome{
			Code: iface.CodeKindSuccess,
		}
	}

	return &iface.ReplyCompleteLedger{
		RatifierRefreshes: app.RatifierRefreshes,
		AgreementArgumentRefreshes: &engineproto.AgreementOptions{
			Release: &engineproto.ReleaseOptions{
				App: 1,
			},
		},
		TransOutcomes: transferOutcomes,
		ApplicationDigest:   app.ApplicationDigest,
	}, nil
}

func (app *verifyApplication) Endorse(_ context.Context, _ *iface.QueryEndorse) (*iface.ReplyEndorse, error) {
	return &iface.ReplyEndorse{PreserveLevel: 1}, nil
}

func (app *verifyApplication) ArrangeNomination(
	_ context.Context,
	req *iface.QueryArrangeNomination,
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
	req *iface.QueryHandleNomination,
) (*iface.ReplyHandleNomination, error) {
	for _, tx := range req.Txs {
		if len(tx) == 0 {
			return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_DECLINE}, nil
		}
	}
	return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil
}
