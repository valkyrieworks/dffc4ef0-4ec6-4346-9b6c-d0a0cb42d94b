package app

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/protocolio"
	cryptographyproto "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	strongmindkinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const (
	platformEdition                 = 1
	ballotAdditionToken    string = "REDACTED"
	ballotAdditionMaximumLength int64  = 1024 * 1024 * 128 //
	ballotAdditionMaximumItem int64  = 128
	headingPreservedToken   string = "REDACTED"
	endingSuccessionUUID       string = "REDACTED"
	endingBallotAddnAltitude string = "REDACTED"
	endingPrimaryAltitude string = "REDACTED"
)

//
//
//
type Platform struct {
	iface.FoundationPlatform
	tracer          log.Tracer
	status           *Status
	images       *ImageDepot
	cfg             *Settings
	recoverImage *iface.Image
	recoverSegments   [][]byte
	applicationTxpool      *ApplicationTxpool
}

//
//
type Settings struct {
	//
	Dir string `toml:"dir"`

	//
	//
	ImageDuration uint64 `toml:"image_duration"`

	//
	//
	//
	PreserveLedgers uint64 `toml:"preserve_ledgers"`

	//
	//
	TokenKind string `toml:"token_kind"`

	//
	//
	//
	EndureDuration uint64 `toml:"endure_duration"`

	//
	ApplicationFlankTxpool bool `toml:"application_flank_txpool"`

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
	AssessorRevisions map[string]map[string]uint8 `toml:"assessor_revise"`

	//
	//
	ArrangeNominationDeferral time.Duration `toml:"arrange_nomination_deferral"`
	HandleNominationDeferral time.Duration `toml:"handle_nomination_deferral"`
	InspectTransferDeferral         time.Duration `toml:"inspect_transfer_deferral"`
	CulminateLedgerDeferral   time.Duration `toml:"culminate_ledger_deferral"`
	BallotAdditionDeferral   time.Duration `toml:"ballot_addition_deferral"`

	//
	//
	//
	BallotAdditionsActivateAltitude int64 `toml:"ballot_additions_activate_altitude"`

	//
	//
	//
	//
	BallotAdditionsReviseAltitude int64 `toml:"ballot_additions_revise_altitude"`

	//
	BallotAdditionExtent uint `toml:"ballot_addition_extent"`
}

func FallbackSettings(dir string) *Settings {
	return &Settings{
		EndureDuration:  1,
		ImageDuration: 100,
		Dir:              dir,
	}
}

//
func FreshPlatform(cfg *Settings) (*Platform, error) {
	status, err := FreshStatus(cfg.Dir, cfg.EndureDuration)
	if err != nil {
		return nil, err
	}
	images, err := FreshImageDepot(filepath.Join(cfg.Dir, "REDACTED"))
	if err != nil {
		return nil, err
	}

	tracer := log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

	return &Platform{
		tracer:     tracer,
		status:      status,
		images:  images,
		cfg:        cfg,
		applicationTxpool: FreshApplicationTxpool(tracer),
	}, nil
}

//
func (app *Platform) Details(context.Context, *iface.SolicitDetails) (*iface.ReplyDetails, error) {
	altitude, digest := app.status.Details()
	return &iface.ReplyDetails{
		Edition:          edition.IfaceEdition,
		PlatformEdition:       platformEdition,
		FinalLedgerAltitude:  int64(altitude),
		FinalLedgerPlatformDigest: digest,
	}, nil
}

func (app *Platform) reviseBallotAdditionActivateAltitude(prevailingAltitude int64) *commitchema.AgreementSettings {
	var parameters *commitchema.AgreementSettings
	if app.cfg.BallotAdditionsReviseAltitude == prevailingAltitude {
		app.tracer.Details("REDACTED",
			"REDACTED", prevailingAltitude,
			"REDACTED", app.cfg.BallotAdditionsActivateAltitude)
		parameters = &commitchema.AgreementSettings{
			Iface: &commitchema.IfaceParameters{
				BallotAdditionsActivateAltitude: app.cfg.BallotAdditionsActivateAltitude,
			},
		}
		app.tracer.Details("REDACTED", "REDACTED", app.cfg.BallotAdditionsActivateAltitude)
		app.status.Set(headingPreservedToken+endingBallotAddnAltitude, strconv.FormatInt(app.cfg.BallotAdditionsActivateAltitude, 10))
	}
	return parameters
}

//
func (app *Platform) InitializeSuccession(_ context.Context, req *iface.SolicitInitializeSuccession) (*iface.ReplyInitializeSuccession, error) {
	var err error
	app.status.primaryAltitude = uint64(req.PrimaryAltitude)
	if len(req.ApplicationStatusOctets) > 0 {
		err = app.status.Ingest(0, req.ApplicationStatusOctets)
		if err != nil {
			return nil, err
		}
	}
	app.tracer.Details("REDACTED", "REDACTED", req.SuccessionUuid)
	app.status.Set(headingPreservedToken+endingSuccessionUUID, req.SuccessionUuid)
	app.tracer.Details("REDACTED", "REDACTED", req.AgreementSettings.Iface.BallotAdditionsActivateAltitude)
	app.status.Set(headingPreservedToken+endingBallotAddnAltitude, strconv.FormatInt(req.AgreementSettings.Iface.BallotAdditionsActivateAltitude, 10))
	app.tracer.Details("REDACTED", "REDACTED", req.PrimaryAltitude)
	app.status.Set(headingPreservedToken+endingPrimaryAltitude, strconv.FormatInt(req.PrimaryAltitude, 10))
	//
	if req.Assessors != nil {
		for _, val := range req.Assessors {

			if err := app.depotAssessor(&val); err != nil {
				return nil, err
			}
		}
	}

	parameters := app.reviseBallotAdditionActivateAltitude(0)

	reply := &iface.ReplyInitializeSuccession{
		AgreementSettings: parameters,
		PlatformDigest:         app.status.ObtainDigest(),
	}
	if reply.Assessors, err = app.assessorRevisions(0); err != nil {
		return nil, err
	}
	return reply, nil
}

//
func (app *Platform) InspectTransfer(_ context.Context, req *iface.SolicitInspectTransfer) (*iface.ReplyInspectTransfer, error) {
	key, _, err := analyzeTransfer(req.Tx)
	if err != nil || key == headingPreservedToken {
		return &iface.ReplyInspectTransfer{
			Cipher: statedepot.CipherKindSerializationFailure,
			Log:  err.Error(),
		}, nil
	}

	if app.cfg.InspectTransferDeferral != 0 {
		time.Sleep(app.cfg.InspectTransferDeferral)
	}

	return &iface.ReplyInspectTransfer{Cipher: statedepot.CipherKindOKAY, FuelDesired: 1}, nil
}

func (app *Platform) AppendTransfer(ctx context.Context, req *iface.SolicitAppendTransfer) (*iface.ReplyAppendTransfer, error) {
	if !app.cfg.ApplicationFlankTxpool {
		return nil, errors.New("REDACTED")
	}

	//
	key, _, err := analyzeTransfer(req.Tx)
	if err != nil || key == headingPreservedToken {
		return &iface.ReplyAppendTransfer{
			Cipher: statedepot.CipherKindSerializationFailure,
		}, nil
	}

	app.applicationTxpool.AppendTransfer(req.Tx)

	return &iface.ReplyAppendTransfer{Cipher: statedepot.CipherKindOKAY}, nil
}

func (app *Platform) HarvestTrans(ctx context.Context, req *iface.SolicitHarvestTrans) (*iface.ReplyHarvestTrans, error) {
	if !app.cfg.ApplicationFlankTxpool {
		return nil, errors.New("REDACTED")
	}

	txs := app.applicationTxpool.HarvestTrans(false)

	return &iface.ReplyHarvestTrans{Txs: txs.TowardSegmentBelongingOctets()}, nil
}

//
func (app *Platform) CulminateLedger(_ context.Context, req *iface.SolicitCulminateLedger) (*iface.ReplyCulminateLedger, error) {
	txs := make([]*iface.InvokeTransferOutcome, len(req.Txs))

	for i, tx := range req.Txs {
		key, datum, err := analyzeTransfer(tx)
		if err != nil {
			panic(err) //
		}
		if key == headingPreservedToken {
			panic(fmt.Errorf("REDACTED", headingPreservedToken))
		}
		app.status.Set(key, datum)

		txs[i] = &iface.InvokeTransferOutcome{Cipher: statedepot.CipherKindOKAY}
	}

	for _, ev := range req.Malpractice {
		app.tracer.Details("REDACTED",
			"REDACTED", ev.ObtainAssessor().Location,
			"REDACTED", ev.ObtainKind(),
			"REDACTED", ev.ObtainAltitude(),
			"REDACTED", ev.ObtainMoment(),
			"REDACTED", ev.ObtainSumBallotingPotency(),
		)
	}

	itemRevisions, err := app.assessorRevisions(uint64(req.Altitude))
	if err != nil {
		panic(err)
	}

	parameters := app.reviseBallotAdditionActivateAltitude(req.Altitude)

	if app.cfg.CulminateLedgerDeferral != 0 {
		time.Sleep(app.cfg.CulminateLedgerDeferral)
	}

	return &iface.ReplyCulminateLedger{
		TransferOutcomes:             txs,
		AssessorRevisions:      itemRevisions,
		PlatformDigest:               app.status.Culminate(),
		AgreementArgumentRevisions: parameters,
		Incidents: []iface.Incident{
			{
				Kind: "REDACTED",
				Properties: []iface.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: strconv.Itoa(itemRevisions.Len()),
					},
					{
						Key:   "REDACTED",
						Datum: strconv.Itoa(int(req.Altitude)),
					},
				},
			},
		},
	}, nil
}

//
func (app *Platform) Endorse(_ context.Context, _ *iface.SolicitEndorse) (*iface.ReplyEndorse, error) {
	altitude, err := app.status.Endorse()
	if err != nil {
		panic(err)
	}
	if app.cfg.ImageDuration > 0 && altitude%app.cfg.ImageDuration == 0 {
		image, err := app.images.Generate(app.status)
		if err != nil {
			panic(err)
		}
		app.tracer.Details("REDACTED", "REDACTED", image.Altitude)
		err = app.images.Trim(maximumImageTally)
		if err != nil {
			app.tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
	preserveAltitude := int64(0)
	if app.cfg.PreserveLedgers > 0 {
		preserveAltitude = int64(altitude - app.cfg.PreserveLedgers + 1)
	}
	return &iface.ReplyEndorse{
		PreserveAltitude: preserveAltitude,
	}, nil
}

//
func (app *Platform) Inquire(_ context.Context, req *iface.SolicitInquire) (*iface.ReplyInquire, error) {
	datum, altitude := app.status.Inquire(string(req.Data))
	return &iface.ReplyInquire{
		Altitude: int64(altitude),
		Key:    req.Data,
		Datum:  []byte(datum),
	}, nil
}

//
func (app *Platform) CollectionImages(context.Context, *iface.SolicitCollectionImages) (*iface.ReplyCatalogImages, error) {
	images, err := app.images.Catalog()
	if err != nil {
		panic(err)
	}
	return &iface.ReplyCatalogImages{Images: images}, nil
}

//
func (app *Platform) FetchImageSegment(_ context.Context, req *iface.SolicitFetchImageSegment) (*iface.ReplyFetchImageSegment, error) {
	segment, err := app.images.FetchSegment(req.Altitude, req.Layout, req.Segment)
	if err != nil {
		panic(err)
	}
	return &iface.ReplyFetchImageSegment{Segment: segment}, nil
}

//
func (app *Platform) ExtendImage(_ context.Context, req *iface.SolicitExtendImage) (*iface.ReplyExtendImage, error) {
	if app.recoverImage != nil {
		panic("REDACTED")
	}
	app.recoverImage = req.Image
	app.recoverSegments = [][]byte{}
	return &iface.ReplyExtendImage{Outcome: iface.Replyextendimage_EMBRACE}, nil
}

//
func (app *Platform) ExecuteImageSegment(_ context.Context, req *iface.SolicitExecuteImageSegment) (*iface.ReplyExecuteImageSegment, error) {
	if app.recoverImage == nil {
		panic("REDACTED")
	}
	app.recoverSegments = append(app.recoverSegments, req.Segment)
	if len(app.recoverSegments) == int(app.recoverImage.Segments) {
		bz := []byte{}
		for _, segment := range app.recoverSegments {
			bz = append(bz, segment...)
		}
		err := app.status.Ingest(app.recoverImage.Altitude, bz)
		if err != nil {
			panic(err)
		}
		app.recoverImage = nil
		app.recoverSegments = nil
	}
	return &iface.ReplyExecuteImageSegment{Outcome: iface.Replyapplyimagefragment_EMBRACE}, nil
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
func (app *Platform) ArrangeNomination(
	_ context.Context, req *iface.SolicitArrangeNomination,
) (*iface.ReplyArrangeNomination, error) {
	//
	if app.cfg.ApplicationFlankTxpool {
		req.Txs = app.applicationTxpool.HarvestTrans(true).TowardSegmentBelongingOctets()
	}

	_, existAdditionsActivated := app.inspectAltitudeAlsoAdditions(true, req.Altitude, "REDACTED")

	txs := make([][]byte, 0, len(req.Txs)+1)
	var sumOctets int64
	addnTransferHeading := fmt.Sprintf("REDACTED", ballotAdditionToken)
	sum, err := app.validateAlsoTotal(existAdditionsActivated, req.Altitude, &req.RegionalFinalEndorse, "REDACTED")
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if existAdditionsActivated {
		addnEndorseOctets, err := req.RegionalFinalEndorse.Serialize()
		if err != nil {
			panic("REDACTED")
		}
		addnEndorseHexadecimal := hex.EncodeToString(addnEndorseOctets)
		addnTransfer := []byte(fmt.Sprintf("REDACTED", addnTransferHeading, sum, addnEndorseHexadecimal))
		addnTransferLength := strongmindkinds.CalculateSchemaExtentForeachTrans([]strongmindkinds.Tx{addnTransfer})
		app.tracer.Details("REDACTED", "REDACTED", addnTransferLength)
		if addnTransferLength > req.MaximumTransferOctets {
			panic(fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				addnTransferLength, req.MaximumTransferOctets))
		}
		txs = append(txs, addnTransfer)
		//
		sumOctets = addnTransferLength
	}
	for _, tx := range req.Txs {
		if existAdditionsActivated && strings.HasPrefix(string(tx), addnTransferHeading) {
			//
			//
			continue
		}
		if strings.HasPrefix(string(tx), headingPreservedToken) {
			app.tracer.Failure("REDACTED", "REDACTED", tx)
			continue
		}
		transferLength := strongmindkinds.CalculateSchemaExtentForeachTrans([]strongmindkinds.Tx{tx})
		if sumOctets+transferLength > req.MaximumTransferOctets {
			break
		}
		sumOctets += transferLength
		//
		txs = append(txs, tx)
	}

	if app.cfg.ArrangeNominationDeferral != 0 {
		time.Sleep(app.cfg.ArrangeNominationDeferral)
	}

	return &iface.ReplyArrangeNomination{Txs: txs}, nil
}

//
//
//
//
func (app *Platform) HandleNomination(_ context.Context, req *iface.SolicitHandleNomination) (*iface.ReplyHandleNomination, error) {
	_, existAdditionsActivated := app.inspectAltitudeAlsoAdditions(true, req.Altitude, "REDACTED")

	for _, tx := range req.Txs {
		k, v, err := analyzeTransfer(tx)
		if err != nil {
			app.tracer.Failure("REDACTED", "REDACTED", tx, "REDACTED", err)
			return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_DECLINE}, nil
		}
		switch {
		case existAdditionsActivated && k == ballotAdditionToken:
			//
			if err := app.validateAdditionTransfer(req.Altitude, v); err != nil {
				app.tracer.Failure("REDACTED", k, v, "REDACTED", err)
				return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_DECLINE}, nil
			}
		case strings.HasPrefix(k, headingPreservedToken):
			app.tracer.Failure("REDACTED", k)
			return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_DECLINE}, nil
		}
	}

	if app.cfg.HandleNominationDeferral != 0 {
		time.Sleep(app.cfg.HandleNominationDeferral)
	}

	return &iface.ReplyHandleNomination{Condition: iface.Responseexecuteitem_EMBRACE}, nil
}

//
//
//
//
//
//
//
func (app *Platform) BroadenBallot(_ context.Context, req *iface.SolicitBroadenBallot) (*iface.ReplyBroadenBallot, error) {
	applicationAltitude, existAdditionsActivated := app.inspectAltitudeAlsoAdditions(false, req.Altitude, "REDACTED")
	if !existAdditionsActivated {
		panic(fmt.Errorf("REDACTED", applicationAltitude))
	}

	if app.cfg.BallotAdditionDeferral != 0 {
		time.Sleep(app.cfg.BallotAdditionDeferral)
	}

	var ext []byte
	var addnLength int
	if app.cfg.BallotAdditionExtent != 0 {
		ext = make([]byte, app.cfg.BallotAdditionExtent)
		if _, err := rand.Read(ext); err != nil {
			panic(fmt.Errorf("REDACTED", len(ext)))
		}
		addnLength = len(ext)
	} else {
		ext = make([]byte, 8)
		if num, err := rand.Int(rand.Reader, big.NewInt(ballotAdditionMaximumItem)); err != nil {
			panic(fmt.Errorf("REDACTED", len(ext)))
		} else {
			addnLength = binary.PutVarint(ext, num.Int64())
		}
	}

	app.tracer.Details("REDACTED", "REDACTED", applicationAltitude, "REDACTED", fmt.Sprintf("REDACTED", ext[:4]), "REDACTED", addnLength)
	return &iface.ReplyBroadenBallot{
		BallotAddition: ext[:addnLength],
	}, nil
}

//
//
//
func (app *Platform) ValidateBallotAddition(_ context.Context, req *iface.SolicitValidateBallotAddition) (*iface.ReplyValidateBallotAddition, error) {
	applicationAltitude, existAdditionsActivated := app.inspectAltitudeAlsoAdditions(false, req.Altitude, "REDACTED")
	if !existAdditionsActivated {
		panic(fmt.Errorf("REDACTED", applicationAltitude))
	}
	//
	if len(req.BallotAddition) == 0 {
		app.tracer.Failure("REDACTED")
		return &iface.ReplyValidateBallotAddition{
			Condition: iface.Responsecertifyballotaddition_DECLINE,
		}, nil
	}

	num, err := analyzeBallotAddition(app.cfg, req.BallotAddition)
	if err != nil {
		app.tracer.Failure("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", req.BallotAddition[:4]), "REDACTED", err)
		return &iface.ReplyValidateBallotAddition{
			Condition: iface.Responsecertifyballotaddition_DECLINE,
		}, nil
	}

	if app.cfg.BallotAdditionDeferral != 0 {
		time.Sleep(app.cfg.BallotAdditionDeferral)
	}

	app.tracer.Details("REDACTED", "REDACTED", req.Altitude, "REDACTED", fmt.Sprintf("REDACTED", req.BallotAddition[:4]), "REDACTED", num)
	return &iface.ReplyValidateBallotAddition{
		Condition: iface.Responsecertifyballotaddition_EMBRACE,
	}, nil
}

func (app *Platform) Revert() error {
	return app.status.Revert()
}

func (app *Platform) obtainApplicationAltitude() int64 {
	primaryAltitudeTxt, altitude := app.status.Inquire(headingPreservedToken + endingPrimaryAltitude)
	if len(primaryAltitudeTxt) == 0 {
		panic("REDACTED")
	}
	primaryAltitude, err := strconv.ParseInt(primaryAltitudeTxt, 10, 64)
	if err != nil {
		panic(fmt.Errorf("REDACTED", primaryAltitudeTxt))
	}

	applicationAltitude := int64(altitude)
	if applicationAltitude == 0 {
		applicationAltitude = primaryAltitude - 1
	}
	return applicationAltitude + 1
}

func (app *Platform) inspectAltitudeAlsoAdditions(equalsArrangeHandleNomination bool, altitude int64, invocation string) (int64, bool) {
	applicationAltitude := app.obtainApplicationAltitude()
	if altitude != applicationAltitude {
		panic(fmt.Errorf(
			"REDACTED",
			invocation, applicationAltitude, altitude,
		))
	}

	ballotAddnAltitudeTxt := app.status.Get(headingPreservedToken + endingBallotAddnAltitude)
	if len(ballotAddnAltitudeTxt) == 0 {
		panic("REDACTED")
	}
	ballotAddnAltitude, err := strconv.ParseInt(ballotAddnAltitudeTxt, 10, 64)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ballotAddnAltitudeTxt))
	}
	prevailingAltitude := applicationAltitude
	if equalsArrangeHandleNomination {
		prevailingAltitude-- //
	}

	return applicationAltitude, ballotAddnAltitude != 0 && prevailingAltitude >= ballotAddnAltitude
}

func (app *Platform) depotAssessor(itemRevise *iface.AssessorRevise) error {
	//
	publicToken, err := cryptocode.PublicTokenOriginatingSchema(itemRevise.PublicToken)
	if err != nil {
		return err
	}
	location := publicToken.Location().Text()
	if itemRevise.Potency > 0 {
		publicTokenOctets, err := itemRevise.PublicToken.Serialize()
		if err != nil {
			return err
		}
		app.tracer.Details("REDACTED", "REDACTED", location)
		app.status.Set(headingPreservedToken+location, hex.EncodeToString(publicTokenOctets))
	}
	return nil
}

//
func (app *Platform) assessorRevisions(altitude uint64) (iface.AssessorRevisions, error) {
	revisions := app.cfg.AssessorRevisions[fmt.Sprintf("REDACTED", altitude)]
	if len(revisions) == 0 {
		return nil, nil
	}

	itemRevisions := iface.AssessorRevisions{}
	for tokenText, potency := range revisions {

		tokenOctets, err := base64.StdEncoding.DecodeString(tokenText)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", tokenText, err)
		}
		itemRevise := iface.ReviseAssessor(tokenOctets, int64(potency), app.cfg.TokenKind)
		itemRevisions = append(itemRevisions, itemRevise)
		if err := app.depotAssessor(&itemRevise); err != nil {
			return nil, err
		}
	}
	return itemRevisions, nil
}

//
func analyzeTransfer(tx []byte) (string, string, error) {
	fragments := bytes.Split(tx, []byte("REDACTED"))
	if len(fragments) != 2 {
		return "REDACTED", "REDACTED", fmt.Errorf("REDACTED", string(tx))
	}
	if len(fragments[0]) == 0 {
		return "REDACTED", "REDACTED", errors.New("REDACTED")
	}
	return string(fragments[0]), string(fragments[1]), nil
}

func (app *Platform) validateAlsoTotal(
	existAdditionsActivated bool,
	prevailingAltitude int64,
	addnEndorse *iface.ExpandedEndorseDetails,
	invocation string,
) (int64, error) {
	var sum int64
	var addnTally int
	for _, ballot := range addnEndorse.Ballots {
		if ballot.LedgerUuidMarker == commitchema.LedgerUUIDMarkerUnfamiliar || ballot.LedgerUuidMarker > commitchema.LedgerUUIDMarkerVoid {
			return 0, fmt.Errorf("REDACTED", prevailingAltitude, ballot.LedgerUuidMarker)
		}
		if ballot.LedgerUuidMarker == commitchema.LedgerUUIDMarkerMissing || ballot.LedgerUuidMarker == commitchema.LedgerUUIDMarkerVoid {
			if len(ballot.BallotAddition) != 0 {
				return 0, fmt.Errorf("REDACTED",
					prevailingAltitude, ballot.LedgerUuidMarker)
			}
			if len(ballot.AdditionNotation) != 0 {
				return 0, fmt.Errorf("REDACTED",
					prevailingAltitude, ballot.LedgerUuidMarker)
			}
			//
			continue
		}
		if !existAdditionsActivated {
			if len(ballot.BallotAddition) != 0 {
				return 0, fmt.Errorf("REDACTED",
					prevailingAltitude)
			}
			if len(ballot.AdditionNotation) != 0 {
				return 0, fmt.Errorf("REDACTED",
					prevailingAltitude)
			}
			continue
		}
		if len(ballot.BallotAddition) == 0 {
			return 0, fmt.Errorf("REDACTED"+
				"REDACTED", ballot.Assessor, prevailingAltitude)
		}
		//
		if len(ballot.AdditionNotation) == 0 {
			return 0, fmt.Errorf("REDACTED", prevailingAltitude)
		}

		//
		successionUUID := app.status.Get(headingPreservedToken + endingSuccessionUUID)
		if len(successionUUID) == 0 {
			panic("REDACTED")
		}
		cve := commitchema.StandardBallotAddition{
			Addition: ballot.BallotAddition,
			Altitude:    prevailingAltitude - 1, //
			Iteration:     int64(addnEndorse.Iteration),
			SuccessionUuid:   successionUUID,
		}
		addnAttestOctets, err := protocolio.SerializeSeparated(&cve)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", err)
		}

		//
		itemLocation := security.Location(ballot.Assessor.Location).Text()
		publicTokenHexadecimal := app.status.Get(headingPreservedToken + itemLocation)
		if len(publicTokenHexadecimal) == 0 {
			return 0, fmt.Errorf("REDACTED", itemLocation)
		}
		publicTokenOctets, err := hex.DecodeString(publicTokenHexadecimal)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", itemLocation, err)
		}
		var publicTokenSchema cryptographyproto.CommonToken
		err = publicTokenSchema.Decode(publicTokenOctets)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", itemLocation, err)
		}
		publicToken, err := cryptocode.PublicTokenOriginatingSchema(publicTokenSchema)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", itemLocation, err)
		}
		if !publicToken.ValidateNotation(addnAttestOctets, ballot.AdditionNotation) {
			return 0, errors.New("REDACTED")
		}

		addnDatum, err := analyzeBallotAddition(app.cfg, ballot.BallotAddition)
		//
		if err != nil {
			return 0, fmt.Errorf("REDACTED", err)
		}
		app.tracer.Details(
			"REDACTED",
			"REDACTED", prevailingAltitude,
			"REDACTED", itemLocation,
			"REDACTED", addnDatum,
			"REDACTED", invocation,
		)
		sum += addnDatum
		addnTally++
	}

	if existAdditionsActivated && (addnTally == 0) {
		return 0, errors.New("REDACTED")
	}
	return sum, nil
}

//
func (app *Platform) validateAdditionTransfer(altitude int64, content string) error {
	fragments := strings.Split(content, "REDACTED")
	if len(fragments) != 2 {
		return fmt.Errorf("REDACTED")
	}
	expirationTotalTxt := fragments[0]
	if len(expirationTotalTxt) == 0 {
		return fmt.Errorf("REDACTED")
	}

	expirationTotal, err := strconv.Atoi(expirationTotalTxt)
	if err != nil {
		return fmt.Errorf("REDACTED", expirationTotalTxt)
	}

	addnEndorseHexadecimal := fragments[1]
	if len(addnEndorseHexadecimal) == 0 {
		return fmt.Errorf("REDACTED")
	}

	addnEndorseOctets, err := hex.DecodeString(addnEndorseHexadecimal)
	if err != nil {
		return fmt.Errorf("REDACTED")
	}

	var addnEndorse iface.ExpandedEndorseDetails
	if addnEndorse.Decode(addnEndorseOctets) != nil {
		return fmt.Errorf("REDACTED")
	}

	sum, err := app.validateAlsoTotal(true, altitude, &addnEndorse, "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if int64(expirationTotal) != sum {
		return fmt.Errorf("REDACTED", expirationTotal, sum)
	}
	return nil
}

//
//
//
func analyzeBallotAddition(cfg *Settings, ext []byte) (int64, error) {
	if cfg.BallotAdditionExtent == 0 {
		num, faultItem := binary.Varint(ext)
		if faultItem == 0 {
			return 0, errors.New("REDACTED")
		}
		if faultItem < 0 {
			return 0, errors.New("REDACTED")
		}
		if num >= ballotAdditionMaximumItem {
			return 0, fmt.Errorf("REDACTED", ballotAdditionMaximumItem, num)
		}
		return num, nil
	}
	return int64(len(ext)), nil
}
