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

	"github.com/valkyrieworks/iface/instance/objectdepot"
	iface "github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/vault"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/utils/protoio"
	cryptography "github.com/valkyrieworks/schema/consensuscore/vault"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	cometkinds "github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

const (
	applicationRelease                 = 1
	ballotAdditionKey    string = "REDACTED"
	ballotAdditionMaximumSize int64  = 1024 * 1024 * 128 //
	ballotAdditionMaximumValue int64  = 128
	prefixPreservedKey   string = "REDACTED"
	postfixSeriesUID       string = "REDACTED"
	postfixBallotExtensionLevel string = "REDACTED"
	postfixPrimaryLevel string = "REDACTED"
)

//
//
//
type Software struct {
	iface.RootSoftware
	tracer          log.Tracer
	status           *Status
	mirrors       *MirrorDepot
	cfg             *Settings
	recoverMirror *iface.Mirror
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
	MirrorCadence uint64 `toml:"mirror_cadence"`

	//
	//
	//
	PreserveLedgers uint64 `toml:"preserve_ledgers"`

	//
	//
	KeyKind string `toml:"key_kind"`

	//
	//
	//
	EndureCadence uint64 `toml:"endure_cadence"`

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
	RatifierRefreshes map[string]map[string]uint8 `toml:"ratifier_modify"`

	//
	//
	ArrangeNominationDeferral time.Duration `toml:"arrange_nomination_deferral"`
	HandleNominationDeferral time.Duration `toml:"handle_nomination_deferral"`
	InspectTransferDeferral         time.Duration `toml:"inspect_transfer_deferral"`
	CompleteLedgerDeferral   time.Duration `toml:"complete_ledger_deferral"`
	BallotAdditionDeferral   time.Duration `toml:"ballot_addition_deferral"`

	//
	//
	//
	BallotPluginsActivateLevel int64 `toml:"ballot_plugins_activate_level"`

	//
	//
	//
	//
	BallotPluginsModifyLevel int64 `toml:"ballot_plugins_modify_level"`

	//
	BallotAdditionVolume uint `toml:"ballot_addition_volume"`
}

func StandardSettings(dir string) *Settings {
	return &Settings{
		EndureCadence:  1,
		MirrorCadence: 100,
		Dir:              dir,
	}
}

//
func NewSoftware(cfg *Settings) (*Software, error) {
	status, err := NewStatus(cfg.Dir, cfg.EndureCadence)
	if err != nil {
		return nil, err
	}
	mirrors, err := NewMirrorDepot(filepath.Join(cfg.Dir, "REDACTED"))
	if err != nil {
		return nil, err
	}

	tracer := log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

	return &Software{
		tracer:     tracer,
		status:      status,
		mirrors:  mirrors,
		cfg:        cfg,
		applicationTxpool: NewApplicationTxpool(tracer),
	}, nil
}

//
func (app *Software) Details(context.Context, *iface.QueryDetails) (*iface.ReplyDetails, error) {
	level, digest := app.status.Details()
	return &iface.ReplyDetails{
		Release:          release.IfaceRelease,
		ApplicationRelease:       applicationRelease,
		FinalLedgerLevel:  int64(level),
		FinalLedgerApplicationDigest: digest,
	}, nil
}

func (app *Software) modifyBallotAdditionActivateLevel(ongoingLevel int64) *engineproto.AgreementOptions {
	var options *engineproto.AgreementOptions
	if app.cfg.BallotPluginsModifyLevel == ongoingLevel {
		app.tracer.Details("REDACTED",
			"REDACTED", ongoingLevel,
			"REDACTED", app.cfg.BallotPluginsActivateLevel)
		options = &engineproto.AgreementOptions{
			Iface: &engineproto.IfaceOptions{
				BallotPluginsActivateLevel: app.cfg.BallotPluginsActivateLevel,
			},
		}
		app.tracer.Details("REDACTED", "REDACTED", app.cfg.BallotPluginsActivateLevel)
		app.status.Set(prefixPreservedKey+postfixBallotExtensionLevel, strconv.FormatInt(app.cfg.BallotPluginsActivateLevel, 10))
	}
	return options
}

//
func (app *Software) InitSeries(_ context.Context, req *iface.QueryInitSeries) (*iface.ReplyInitSeries, error) {
	var err error
	app.status.primaryLevel = uint64(req.PrimaryLevel)
	if len(req.ApplicationStatusOctets) > 0 {
		err = app.status.Include(0, req.ApplicationStatusOctets)
		if err != nil {
			return nil, err
		}
	}
	app.tracer.Details("REDACTED", "REDACTED", req.SeriesUid)
	app.status.Set(prefixPreservedKey+postfixSeriesUID, req.SeriesUid)
	app.tracer.Details("REDACTED", "REDACTED", req.AgreementOptions.Iface.BallotPluginsActivateLevel)
	app.status.Set(prefixPreservedKey+postfixBallotExtensionLevel, strconv.FormatInt(req.AgreementOptions.Iface.BallotPluginsActivateLevel, 10))
	app.tracer.Details("REDACTED", "REDACTED", req.PrimaryLevel)
	app.status.Set(prefixPreservedKey+postfixPrimaryLevel, strconv.FormatInt(req.PrimaryLevel, 10))
	//
	if req.Ratifiers != nil {
		for _, val := range req.Ratifiers {

			if err := app.depotRatifier(&val); err != nil {
				return nil, err
			}
		}
	}

	options := app.modifyBallotAdditionActivateLevel(0)

	reply := &iface.ReplyInitSeries{
		AgreementOptions: options,
		ApplicationDigest:         app.status.FetchDigest(),
	}
	if reply.Ratifiers, err = app.ratifierRefreshes(0); err != nil {
		return nil, err
	}
	return reply, nil
}

//
func (app *Software) InspectTransfer(_ context.Context, req *iface.QueryInspectTransfer) (*iface.ReplyInspectTransfer, error) {
	key, _, err := analyzeTransfer(req.Tx)
	if err != nil || key == prefixPreservedKey {
		return &iface.ReplyInspectTransfer{
			Code: objectdepot.CodeKindCodecFault,
			Log:  err.Error(),
		}, nil
	}

	if app.cfg.InspectTransferDeferral != 0 {
		time.Sleep(app.cfg.InspectTransferDeferral)
	}

	return &iface.ReplyInspectTransfer{Code: objectdepot.CodeKindSuccess, FuelDesired: 1}, nil
}

func (app *Software) EmbedTransfer(ctx context.Context, req *iface.QueryEmbedTransfer) (*iface.ReplyEmbedTransfer, error) {
	if !app.cfg.ApplicationFlankTxpool {
		return nil, errors.New("REDACTED")
	}

	//
	key, _, err := analyzeTransfer(req.Tx)
	if err != nil || key == prefixPreservedKey {
		return &iface.ReplyEmbedTransfer{
			Code: objectdepot.CodeKindCodecFault,
		}, nil
	}

	app.applicationTxpool.EmbedTransfer(req.Tx)

	return &iface.ReplyEmbedTransfer{Code: objectdepot.CodeKindSuccess}, nil
}

func (app *Software) HarvestTrans(ctx context.Context, req *iface.QueryHarvestTrans) (*iface.ReplyHarvestTrans, error) {
	if !app.cfg.ApplicationFlankTxpool {
		return nil, errors.New("REDACTED")
	}

	txs := app.applicationTxpool.HarvestTrans(false)

	return &iface.ReplyHarvestTrans{Txs: txs.ToSegmentOfOctets()}, nil
}

//
func (app *Software) CompleteLedger(_ context.Context, req *iface.QueryCompleteLedger) (*iface.ReplyCompleteLedger, error) {
	txs := make([]*iface.InvokeTransferOutcome, len(req.Txs))

	for i, tx := range req.Txs {
		key, item, err := analyzeTransfer(tx)
		if err != nil {
			panic(err) //
		}
		if key == prefixPreservedKey {
			panic(fmt.Errorf("REDACTED", prefixPreservedKey))
		}
		app.status.Set(key, item)

		txs[i] = &iface.InvokeTransferOutcome{Code: objectdepot.CodeKindSuccess}
	}

	for _, ev := range req.Malpractice {
		app.tracer.Details("REDACTED",
			"REDACTED", ev.FetchRatifier().Location,
			"REDACTED", ev.FetchKind(),
			"REDACTED", ev.FetchLevel(),
			"REDACTED", ev.FetchTime(),
			"REDACTED", ev.FetchSumPollingEnergy(),
		)
	}

	valueRefreshes, err := app.ratifierRefreshes(uint64(req.Level))
	if err != nil {
		panic(err)
	}

	options := app.modifyBallotAdditionActivateLevel(req.Level)

	if app.cfg.CompleteLedgerDeferral != 0 {
		time.Sleep(app.cfg.CompleteLedgerDeferral)
	}

	return &iface.ReplyCompleteLedger{
		TransOutcomes:             txs,
		RatifierRefreshes:      valueRefreshes,
		ApplicationDigest:               app.status.Complete(),
		AgreementArgumentRefreshes: options,
		Events: []iface.Event{
			{
				Kind: "REDACTED",
				Properties: []iface.EventProperty{
					{
						Key:   "REDACTED",
						Item: strconv.Itoa(valueRefreshes.Len()),
					},
					{
						Key:   "REDACTED",
						Item: strconv.Itoa(int(req.Level)),
					},
				},
			},
		},
	}, nil
}

//
func (app *Software) Endorse(_ context.Context, _ *iface.QueryEndorse) (*iface.ReplyEndorse, error) {
	level, err := app.status.Endorse()
	if err != nil {
		panic(err)
	}
	if app.cfg.MirrorCadence > 0 && level%app.cfg.MirrorCadence == 0 {
		mirror, err := app.mirrors.Instantiate(app.status)
		if err != nil {
			panic(err)
		}
		app.tracer.Details("REDACTED", "REDACTED", mirror.Level)
		err = app.mirrors.Trim(maximumMirrorNumber)
		if err != nil {
			app.tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
	preserveLevel := int64(0)
	if app.cfg.PreserveLedgers > 0 {
		preserveLevel = int64(level - app.cfg.PreserveLedgers + 1)
	}
	return &iface.ReplyEndorse{
		PreserveLevel: preserveLevel,
	}, nil
}

//
func (app *Software) Inquire(_ context.Context, req *iface.QueryInquire) (*iface.ReplyInquire, error) {
	item, level := app.status.Inquire(string(req.Data))
	return &iface.ReplyInquire{
		Level: int64(level),
		Key:    req.Data,
		Item:  []byte(item),
	}, nil
}

//
func (app *Software) CatalogMirrors(context.Context, *iface.QueryCatalogMirrors) (*iface.ReplyCatalogMirrors, error) {
	mirrors, err := app.mirrors.Catalog()
	if err != nil {
		panic(err)
	}
	return &iface.ReplyCatalogMirrors{Mirrors: mirrors}, nil
}

//
func (app *Software) ImportMirrorSegment(_ context.Context, req *iface.QueryImportMirrorSegment) (*iface.ReplyImportMirrorSegment, error) {
	segment, err := app.mirrors.ImportSegment(req.Level, req.Layout, req.Segment)
	if err != nil {
		panic(err)
	}
	return &iface.ReplyImportMirrorSegment{Segment: segment}, nil
}

//
func (app *Software) ProposalMirror(_ context.Context, req *iface.QueryProposalMirror) (*iface.ReplyProposalMirror, error) {
	if app.recoverMirror != nil {
		panic("REDACTED")
	}
	app.recoverMirror = req.Mirror
	app.recoverSegments = [][]byte{}
	return &iface.ReplyProposalMirror{Outcome: iface.Replymirrorsnapshot_ALLOW}, nil
}

//
func (app *Software) ExecuteMirrorSegment(_ context.Context, req *iface.QueryExecuteMirrorSegment) (*iface.ReplyExecuteMirrorSegment, error) {
	if app.recoverMirror == nil {
		panic("REDACTED")
	}
	app.recoverSegments = append(app.recoverSegments, req.Segment)
	if len(app.recoverSegments) == int(app.recoverMirror.Segments) {
		bz := []byte{}
		for _, segment := range app.recoverSegments {
			bz = append(bz, segment...)
		}
		err := app.status.Include(app.recoverMirror.Level, bz)
		if err != nil {
			panic(err)
		}
		app.recoverMirror = nil
		app.recoverSegments = nil
	}
	return &iface.ReplyExecuteMirrorSegment{Outcome: iface.Replyexecutemirrorsegment_ALLOW}, nil
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
func (app *Software) ArrangeNomination(
	_ context.Context, req *iface.QueryArrangeNomination,
) (*iface.ReplyArrangeNomination, error) {
	//
	if app.cfg.ApplicationFlankTxpool {
		req.Txs = app.applicationTxpool.HarvestTrans(true).ToSegmentOfOctets()
	}

	_, arePluginsActivated := app.inspectLevelAndPlugins(true, req.Level, "REDACTED")

	txs := make([][]byte, 0, len(req.Txs)+1)
	var sumOctets int64
	extensionTransferPrefix := fmt.Sprintf("REDACTED", ballotAdditionKey)
	sum, err := app.validateAndTotal(arePluginsActivated, req.Level, &req.NativeFinalEndorse, "REDACTED")
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	if arePluginsActivated {
		extensionEndorseOctets, err := req.NativeFinalEndorse.Serialize()
		if err != nil {
			panic("REDACTED")
		}
		extensionEndorseHex := hex.EncodeToString(extensionEndorseOctets)
		extensionTransfer := []byte(fmt.Sprintf("REDACTED", extensionTransferPrefix, sum, extensionEndorseHex))
		extensionTransferSize := cometkinds.CalculateSchemaVolumeForTrans([]cometkinds.Tx{extensionTransfer})
		app.tracer.Details("REDACTED", "REDACTED", extensionTransferSize)
		if extensionTransferSize > req.MaximumTransferOctets {
			panic(fmt.Errorf("REDACTED"+
				"REDACTED"+
				"REDACTED",
				extensionTransferSize, req.MaximumTransferOctets))
		}
		txs = append(txs, extensionTransfer)
		//
		sumOctets = extensionTransferSize
	}
	for _, tx := range req.Txs {
		if arePluginsActivated && strings.HasPrefix(string(tx), extensionTransferPrefix) {
			//
			//
			continue
		}
		if strings.HasPrefix(string(tx), prefixPreservedKey) {
			app.tracer.Fault("REDACTED", "REDACTED", tx)
			continue
		}
		transferSize := cometkinds.CalculateSchemaVolumeForTrans([]cometkinds.Tx{tx})
		if sumOctets+transferSize > req.MaximumTransferOctets {
			break
		}
		sumOctets += transferSize
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
func (app *Software) HandleNomination(_ context.Context, req *iface.QueryHandleNomination) (*iface.ReplyHandleNomination, error) {
	_, arePluginsActivated := app.inspectLevelAndPlugins(true, req.Level, "REDACTED")

	for _, tx := range req.Txs {
		k, v, err := analyzeTransfer(tx)
		if err != nil {
			app.tracer.Fault("REDACTED", "REDACTED", tx, "REDACTED", err)
			return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_DECLINE}, nil
		}
		switch {
		case arePluginsActivated && k == ballotAdditionKey:
			//
			if err := app.validateAdditionTransfer(req.Level, v); err != nil {
				app.tracer.Fault("REDACTED", k, v, "REDACTED", err)
				return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_DECLINE}, nil
			}
		case strings.HasPrefix(k, prefixPreservedKey):
			app.tracer.Fault("REDACTED", k)
			return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_DECLINE}, nil
		}
	}

	if app.cfg.HandleNominationDeferral != 0 {
		time.Sleep(app.cfg.HandleNominationDeferral)
	}

	return &iface.ReplyHandleNomination{Status: iface.Responseprocessnomination_ALLOW}, nil
}

//
//
//
//
//
//
//
func (app *Software) ExpandBallot(_ context.Context, req *iface.QueryExpandBallot) (*iface.ReplyExpandBallot, error) {
	applicationLevel, arePluginsActivated := app.inspectLevelAndPlugins(false, req.Level, "REDACTED")
	if !arePluginsActivated {
		panic(fmt.Errorf("REDACTED", applicationLevel))
	}

	if app.cfg.BallotAdditionDeferral != 0 {
		time.Sleep(app.cfg.BallotAdditionDeferral)
	}

	var ext []byte
	var extensionSize int
	if app.cfg.BallotAdditionVolume != 0 {
		ext = make([]byte, app.cfg.BallotAdditionVolume)
		if _, err := rand.Read(ext); err != nil {
			panic(fmt.Errorf("REDACTED", len(ext)))
		}
		extensionSize = len(ext)
	} else {
		ext = make([]byte, 8)
		if num, err := rand.Int(rand.Reader, big.NewInt(ballotAdditionMaximumValue)); err != nil {
			panic(fmt.Errorf("REDACTED", len(ext)))
		} else {
			extensionSize = binary.PutVarint(ext, num.Int64())
		}
	}

	app.tracer.Details("REDACTED", "REDACTED", applicationLevel, "REDACTED", fmt.Sprintf("REDACTED", ext[:4]), "REDACTED", extensionSize)
	return &iface.ReplyExpandBallot{
		BallotAddition: ext[:extensionSize],
	}, nil
}

//
//
//
func (app *Software) ValidateBallotAddition(_ context.Context, req *iface.QueryValidateBallotAddition) (*iface.ReplyValidateBallotAddition, error) {
	applicationLevel, arePluginsActivated := app.inspectLevelAndPlugins(false, req.Level, "REDACTED")
	if !arePluginsActivated {
		panic(fmt.Errorf("REDACTED", applicationLevel))
	}
	//
	if len(req.BallotAddition) == 0 {
		app.tracer.Fault("REDACTED")
		return &iface.ReplyValidateBallotAddition{
			Status: iface.Responseverifyballotextension_DECLINE,
		}, nil
	}

	num, err := analyzeBallotAddition(app.cfg, req.BallotAddition)
	if err != nil {
		app.tracer.Fault("REDACTED", "REDACTED", fmt.Sprintf("REDACTED", req.BallotAddition[:4]), "REDACTED", err)
		return &iface.ReplyValidateBallotAddition{
			Status: iface.Responseverifyballotextension_DECLINE,
		}, nil
	}

	if app.cfg.BallotAdditionDeferral != 0 {
		time.Sleep(app.cfg.BallotAdditionDeferral)
	}

	app.tracer.Details("REDACTED", "REDACTED", req.Level, "REDACTED", fmt.Sprintf("REDACTED", req.BallotAddition[:4]), "REDACTED", num)
	return &iface.ReplyValidateBallotAddition{
		Status: iface.Responseverifyballotextension_ALLOW,
	}, nil
}

func (app *Software) Revert() error {
	return app.status.Revert()
}

func (app *Software) fetchApplicationLevel() int64 {
	primaryLevelStr, level := app.status.Inquire(prefixPreservedKey + postfixPrimaryLevel)
	if len(primaryLevelStr) == 0 {
		panic("REDACTED")
	}
	primaryLevel, err := strconv.ParseInt(primaryLevelStr, 10, 64)
	if err != nil {
		panic(fmt.Errorf("REDACTED", primaryLevelStr))
	}

	applicationLevel := int64(level)
	if applicationLevel == 0 {
		applicationLevel = primaryLevel - 1
	}
	return applicationLevel + 1
}

func (app *Software) inspectLevelAndPlugins(isArrangeHandleNomination bool, level int64, invocation string) (int64, bool) {
	applicationLevel := app.fetchApplicationLevel()
	if level != applicationLevel {
		panic(fmt.Errorf(
			"REDACTED",
			invocation, applicationLevel, level,
		))
	}

	ballotExtensionLevelStr := app.status.Get(prefixPreservedKey + postfixBallotExtensionLevel)
	if len(ballotExtensionLevelStr) == 0 {
		panic("REDACTED")
	}
	ballotExtensionLevel, err := strconv.ParseInt(ballotExtensionLevelStr, 10, 64)
	if err != nil {
		panic(fmt.Errorf("REDACTED", ballotExtensionLevelStr))
	}
	ongoingLevel := applicationLevel
	if isArrangeHandleNomination {
		ongoingLevel-- //
	}

	return applicationLevel, ballotExtensionLevel != 0 && ongoingLevel >= ballotExtensionLevel
}

func (app *Software) depotRatifier(valueModify *iface.RatifierModify) error {
	//
	publicKey, err := cryptocode.PublicKeyFromSchema(valueModify.PublicKey)
	if err != nil {
		return err
	}
	address := publicKey.Location().String()
	if valueModify.Energy > 0 {
		publicKeyOctets, err := valueModify.PublicKey.Serialize()
		if err != nil {
			return err
		}
		app.tracer.Details("REDACTED", "REDACTED", address)
		app.status.Set(prefixPreservedKey+address, hex.EncodeToString(publicKeyOctets))
	}
	return nil
}

//
func (app *Software) ratifierRefreshes(level uint64) (iface.RatifierRefreshes, error) {
	refreshes := app.cfg.RatifierRefreshes[fmt.Sprintf("REDACTED", level)]
	if len(refreshes) == 0 {
		return nil, nil
	}

	valueRefreshes := iface.RatifierRefreshes{}
	for keyString, energy := range refreshes {

		keyOctets, err := base64.StdEncoding.DecodeString(keyString)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", keyString, err)
		}
		valueModify := iface.ModifyRatifier(keyOctets, int64(energy), app.cfg.KeyKind)
		valueRefreshes = append(valueRefreshes, valueModify)
		if err := app.depotRatifier(&valueModify); err != nil {
			return nil, err
		}
	}
	return valueRefreshes, nil
}

//
func analyzeTransfer(tx []byte) (string, string, error) {
	segments := bytes.Split(tx, []byte("REDACTED"))
	if len(segments) != 2 {
		return "REDACTED", "REDACTED", fmt.Errorf("REDACTED", string(tx))
	}
	if len(segments[0]) == 0 {
		return "REDACTED", "REDACTED", errors.New("REDACTED")
	}
	return string(segments[0]), string(segments[1]), nil
}

func (app *Software) validateAndTotal(
	arePluginsActivated bool,
	ongoingLevel int64,
	extensionEndorse *iface.ExpandedEndorseDetails,
	invocation string,
) (int64, error) {
	var sum int64
	var extensionNumber int
	for _, ballot := range extensionEndorse.Ballots {
		if ballot.LedgerUidMark == engineproto.LedgerUIDMarkUnclear || ballot.LedgerUidMark > engineproto.LedgerUIDMarkNull {
			return 0, fmt.Errorf("REDACTED", ongoingLevel, ballot.LedgerUidMark)
		}
		if ballot.LedgerUidMark == engineproto.LedgerUIDMarkMissing || ballot.LedgerUidMark == engineproto.LedgerUIDMarkNull {
			if len(ballot.BallotAddition) != 0 {
				return 0, fmt.Errorf("REDACTED",
					ongoingLevel, ballot.LedgerUidMark)
			}
			if len(ballot.AdditionAutograph) != 0 {
				return 0, fmt.Errorf("REDACTED",
					ongoingLevel, ballot.LedgerUidMark)
			}
			//
			continue
		}
		if !arePluginsActivated {
			if len(ballot.BallotAddition) != 0 {
				return 0, fmt.Errorf("REDACTED",
					ongoingLevel)
			}
			if len(ballot.AdditionAutograph) != 0 {
				return 0, fmt.Errorf("REDACTED",
					ongoingLevel)
			}
			continue
		}
		if len(ballot.BallotAddition) == 0 {
			return 0, fmt.Errorf("REDACTED"+
				"REDACTED", ballot.Ratifier, ongoingLevel)
		}
		//
		if len(ballot.AdditionAutograph) == 0 {
			return 0, fmt.Errorf("REDACTED", ongoingLevel)
		}

		//
		ledgerUID := app.status.Get(prefixPreservedKey + postfixSeriesUID)
		if len(ledgerUID) == 0 {
			panic("REDACTED")
		}
		cve := engineproto.StandardBallotAddition{
			Addition: ballot.BallotAddition,
			Level:    ongoingLevel - 1, //
			Cycle:     int64(extensionEndorse.Cycle),
			SeriesUid:   ledgerUID,
		}
		extensionAttestOctets, err := protoio.SerializeSeparated(&cve)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", err)
		}

		//
		valueAddress := vault.Location(ballot.Ratifier.Location).String()
		publicKeyHex := app.status.Get(prefixPreservedKey + valueAddress)
		if len(publicKeyHex) == 0 {
			return 0, fmt.Errorf("REDACTED", valueAddress)
		}
		publicKeyOctets, err := hex.DecodeString(publicKeyHex)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", valueAddress, err)
		}
		var publicKeySchema cryptography.PublicKey
		err = publicKeySchema.Unserialize(publicKeyOctets)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", valueAddress, err)
		}
		publicKey, err := cryptocode.PublicKeyFromSchema(publicKeySchema)
		if err != nil {
			return 0, fmt.Errorf("REDACTED", valueAddress, err)
		}
		if !publicKey.ValidateAutograph(extensionAttestOctets, ballot.AdditionAutograph) {
			return 0, errors.New("REDACTED")
		}

		extensionItem, err := analyzeBallotAddition(app.cfg, ballot.BallotAddition)
		//
		if err != nil {
			return 0, fmt.Errorf("REDACTED", err)
		}
		app.tracer.Details(
			"REDACTED",
			"REDACTED", ongoingLevel,
			"REDACTED", valueAddress,
			"REDACTED", extensionItem,
			"REDACTED", invocation,
		)
		sum += extensionItem
		extensionNumber++
	}

	if arePluginsActivated && (extensionNumber == 0) {
		return 0, errors.New("REDACTED")
	}
	return sum, nil
}

//
func (app *Software) validateAdditionTransfer(level int64, shipment string) error {
	segments := strings.Split(shipment, "REDACTED")
	if len(segments) != 2 {
		return fmt.Errorf("REDACTED")
	}
	expirationTotalStr := segments[0]
	if len(expirationTotalStr) == 0 {
		return fmt.Errorf("REDACTED")
	}

	expirationTotal, err := strconv.Atoi(expirationTotalStr)
	if err != nil {
		return fmt.Errorf("REDACTED", expirationTotalStr)
	}

	extensionEndorseHex := segments[1]
	if len(extensionEndorseHex) == 0 {
		return fmt.Errorf("REDACTED")
	}

	extensionEndorseOctets, err := hex.DecodeString(extensionEndorseHex)
	if err != nil {
		return fmt.Errorf("REDACTED")
	}

	var extensionEndorse iface.ExpandedEndorseDetails
	if extensionEndorse.Unserialize(extensionEndorseOctets) != nil {
		return fmt.Errorf("REDACTED")
	}

	sum, err := app.validateAndTotal(true, level, &extensionEndorse, "REDACTED")
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
	if cfg.BallotAdditionVolume == 0 {
		num, errValue := binary.Varint(ext)
		if errValue == 0 {
			return 0, errors.New("REDACTED")
		}
		if errValue < 0 {
			return 0, errors.New("REDACTED")
		}
		if num >= ballotAdditionMaximumValue {
			return 0, fmt.Errorf("REDACTED", ballotAdditionMaximumValue, num)
		}
		return num, nil
	}
	return int64(len(ext)), nil
}
