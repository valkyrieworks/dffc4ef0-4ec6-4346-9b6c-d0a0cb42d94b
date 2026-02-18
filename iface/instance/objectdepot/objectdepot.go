package objectdepot

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/-db"

	"github.com/valkyrieworks/iface/kinds"
	cryptography "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/utils/log"
	cryptography "github.com/valkyrieworks/schema/consensuscore/vault"
	"github.com/valkyrieworks/release"
)

var (
	statusKey        = []byte("REDACTED")
	objectCouplePrefixKey = []byte("REDACTED")
)

const (
	RatifierPrefix        = "REDACTED"
	ApplicationRelease      uint64 = 1
)

var _ kinds.Software = (*Software)(nil)

//
//
//
type Software struct {
	kinds.RootSoftware

	status        Status
	PreserveLedgers int64 //
	arrangedTrans    [][]byte
	tracer       log.Tracer

	//
	valueRefreshes         []kinds.RatifierModify
	valueAddressToPublicKeyIndex map[string]cryptography.PublicKey

	//
	//
	generateLedgerEvents bool
}

//
func NewSoftware(db dbm.DB) *Software {
	return &Software{
		tracer:             log.NewNoopTracer(),
		status:              importStatus(db),
		valueAddressToPublicKeyIndex: make(map[string]cryptography.PublicKey),
	}
}

//
func NewDurableSoftware(storeFolder string) *Software {
	label := "REDACTED"
	db, err := dbm.NewGoLayerStore(label, storeFolder)
	if err != nil {
		panic(fmt.Errorf("REDACTED", storeFolder, err))
	}
	return NewSoftware(db)
}

//
//
func NewInRamSoftware() *Software {
	return NewSoftware(dbm.NewMemoryStore())
}

func (app *Software) CollectionGenerateLedgerEvents() {
	app.generateLedgerEvents = true
}

//
//
//
//
func (app *Software) Details(context.Context, *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	//
	if len(app.valueAddressToPublicKeyIndex) == 0 && app.status.Level > 0 {
		ratifiers := app.fetchRatifiers()
		for _, v := range ratifiers {
			publickey, err := cryptography.PublicKeyFromSchema(v.PublicKey)
			if err != nil {
				panic(fmt.Errorf("REDACTED", err))
			}
			app.valueAddressToPublicKeyIndex[string(publickey.Location())] = v.PublicKey
		}
	}

	return &kinds.ReplyDetails{
		Data:             fmt.Sprintf("REDACTED", app.status.Volume),
		Release:          release.IfaceRelease,
		ApplicationRelease:       ApplicationRelease,
		FinalLedgerLevel:  app.status.Level,
		FinalLedgerApplicationDigest: app.status.Digest(),
	}, nil
}

//
//
//
func (app *Software) InitSeries(_ context.Context, req *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	for _, v := range req.Ratifiers {
		app.modifyRatifier(v)
	}
	applicationDigest := make([]byte, 8)
	binary.PutVarint(applicationDigest, app.status.Volume)
	return &kinds.ReplyInitSeries{
		ApplicationDigest: applicationDigest,
	}, nil
}

//
//
//
//
//
//
//
func (app *Software) InspectTransfer(_ context.Context, req *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	//
	if isRatifierTransfer(req.Tx) {
		if _, _, _, err := analyzeRatifierTransfer(req.Tx); err != nil {
			//
			return &kinds.ReplyInspectTransfer{Code: CodeKindCorruptTransferLayout}, nil
		}
	} else if !isSoundTransfer(req.Tx) {
		return &kinds.ReplyInspectTransfer{Code: CodeKindCorruptTransferLayout}, nil
	}

	return &kinds.ReplyInspectTransfer{Code: CodeKindSuccess, FuelDesired: 1}, nil
}

//
//
//
func isSoundTransfer(tx []byte) bool {
	if bytes.Count(tx, []byte("REDACTED")) == 1 && bytes.Count(tx, []byte("REDACTED")) == 0 {
		if !bytes.HasPrefix(tx, []byte("REDACTED")) && !bytes.HasSuffix(tx, []byte("REDACTED")) {
			return true
		}
	} else if bytes.Count(tx, []byte("REDACTED")) == 1 && bytes.Count(tx, []byte("REDACTED")) == 0 {
		if !bytes.HasPrefix(tx, []byte("REDACTED")) && !bytes.HasSuffix(tx, []byte("REDACTED")) {
			return true
		}
	}
	return false
}

//
//
//
//
func (app *Software) ArrangeNomination(ctx context.Context, req *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	return &kinds.ReplyArrangeNomination{Txs: app.layoutTrans(ctx, req.Txs)}, nil
}

//
//
func (app *Software) layoutTrans(ctx context.Context, ledgerData [][]byte) [][]byte {
	txs := make([][]byte, 0, len(ledgerData))
	for _, tx := range ledgerData {
		if reply, err := app.InspectTransfer(ctx, &kinds.QueryInspectTransfer{Tx: tx}); err == nil && reply.Code == CodeKindSuccess {
			txs = append(txs, bytes.Replace(tx, []byte("REDACTED"), []byte("REDACTED"), 1))
		}
	}
	return txs
}

//
//
func (app *Software) HandleNomination(ctx context.Context, req *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	for _, tx := range req.Txs {
		//
		if reply, err := app.InspectTransfer(ctx, &kinds.QueryInspectTransfer{Tx: tx}); err != nil || reply.Code != CodeKindSuccess {
			return &kinds.ReplyHandleNomination{Status: kinds.Responseprocessnomination_DECLINE}, nil
		}
	}
	return &kinds.ReplyHandleNomination{Status: kinds.Responseprocessnomination_ALLOW}, nil
}

//
//
//
//
func (app *Software) CompleteLedger(_ context.Context, req *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	//
	app.valueRefreshes = make([]kinds.RatifierModify, 0)
	app.arrangedTrans = make([][]byte, 0)

	//
	for _, ev := range req.Malpractice {
		if ev.Kind == kinds.Misconductkind_REPLICATED_BALLOT {
			address := string(ev.Ratifier.Location)
			if publicKey, ok := app.valueAddressToPublicKeyIndex[address]; ok {
				app.valueRefreshes = append(app.valueRefreshes, kinds.RatifierModify{
					PublicKey: publicKey,
					Energy:  ev.Ratifier.Energy - 1,
				})
				app.tracer.Details("REDACTED",
					"REDACTED", address)
			} else {
				panic(fmt.Errorf("REDACTED", address))
			}
		}
	}

	replyTrans := make([]*kinds.InvokeTransferOutcome, len(req.Txs))
	for i, tx := range req.Txs {
		if isRatifierTransfer(tx) {
			keyKind, publicKey, energy, err := analyzeRatifierTransfer(tx)
			if err != nil {
				panic(err)
			}
			app.valueRefreshes = append(app.valueRefreshes, kinds.ModifyRatifier(publicKey, energy, keyKind))
		} else {
			app.arrangedTrans = append(app.arrangedTrans, tx)
		}

		var key, item string
		segments := bytes.Split(tx, []byte("REDACTED"))
		if len(segments) == 2 {
			key, item = string(segments[0]), string(segments[1])
		} else {
			key, item = string(tx), string(tx)
		}
		replyTrans[i] = &kinds.InvokeTransferOutcome{
			Code: CodeKindSuccess,
			//
			Events: []kinds.Event{
				{
					Kind: "REDACTED",
					Properties: []kinds.EventProperty{
						{Key: "REDACTED", Item: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Item: key, Ordinal: true},
						{Key: "REDACTED", Item: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Item: "REDACTED", Ordinal: false},
					},
				},
				{
					Kind: "REDACTED",
					Properties: []kinds.EventProperty{
						{Key: "REDACTED", Item: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Item: item, Ordinal: true},
						{Key: "REDACTED", Item: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Item: "REDACTED", Ordinal: false},
					},
				},
			},
		}
		app.status.Volume++
	}

	app.status.Level = req.Level

	reply := &kinds.ReplyCompleteLedger{TransOutcomes: replyTrans, RatifierRefreshes: app.valueRefreshes, ApplicationDigest: app.status.Digest()}
	if !app.generateLedgerEvents {
		return reply, nil
	}
	if app.status.Level%2 == 0 {
		reply.Events = []kinds.Event{
			{
				Kind: "REDACTED",
				Properties: []kinds.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []kinds.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		}
	} else {
		reply.Events = []kinds.Event{
			{
				Kind: "REDACTED",
				Properties: []kinds.EventProperty{
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Item: "REDACTED",
						Ordinal: true,
					},
				},
			},
		}
	}
	return reply, nil
}

//
//
//
func (app *Software) Endorse(context.Context, *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	//
	for _, valueModify := range app.valueRefreshes {
		app.modifyRatifier(valueModify)
	}

	//
	for _, tx := range app.arrangedTrans {
		segments := bytes.Split(tx, []byte("REDACTED"))
		if len(segments) != 2 {
			panic(fmt.Sprintf("REDACTED", len(segments), segments))
		}
		key, item := string(segments[0]), string(segments[1])
		err := app.status.db.Set(prefixKey([]byte(key)), []byte(item))
		if err != nil {
			panic(err)
		}
	}

	//
	persistStatus(app.status)

	reply := &kinds.ReplyEndorse{}
	if app.PreserveLedgers > 0 && app.status.Level >= app.PreserveLedgers {
		reply.PreserveLevel = app.status.Level - app.PreserveLedgers + 1
	}
	return reply, nil
}

//
func (app *Software) Inquire(_ context.Context, requestInquire *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	outcomeInquire := &kinds.ReplyInquire{}

	if requestInquire.Route == "REDACTED" {
		key := []byte(RatifierPrefix + string(requestInquire.Data))
		item, err := app.status.db.Get(key)
		if err != nil {
			panic(err)
		}

		return &kinds.ReplyInquire{
			Key:   requestInquire.Data,
			Item: item,
		}, nil
	}

	if requestInquire.Demonstrate {
		item, err := app.status.db.Get(prefixKey(requestInquire.Data))
		if err != nil {
			panic(err)
		}

		if item == nil {
			outcomeInquire.Log = "REDACTED"
		} else {
			outcomeInquire.Log = "REDACTED"
		}
		outcomeInquire.Ordinal = -1 //
		outcomeInquire.Key = requestInquire.Data
		outcomeInquire.Item = item
		outcomeInquire.Level = app.status.Level

		return outcomeInquire, nil
	}

	outcomeInquire.Key = requestInquire.Data
	item, err := app.status.db.Get(prefixKey(requestInquire.Data))
	if err != nil {
		panic(err)
	}
	if item == nil {
		outcomeInquire.Log = "REDACTED"
	} else {
		outcomeInquire.Log = "REDACTED"
	}
	outcomeInquire.Item = item
	outcomeInquire.Level = app.status.Level

	return outcomeInquire, nil
}

func (app *Software) End() error {
	return app.status.db.End()
}

func isRatifierTransfer(tx []byte) bool {
	return strings.HasPrefix(string(tx), RatifierPrefix)
}

func analyzeRatifierTransfer(tx []byte) (string, []byte, int64, error) {
	tx = tx[len(RatifierPrefix):]

	//
	kindPublicKeyAndEnergy := strings.Split(string(tx), "REDACTED")
	if len(kindPublicKeyAndEnergy) != 3 {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", kindPublicKeyAndEnergy)
	}
	keyKind, publickeyS, energyS := kindPublicKeyAndEnergy[0], kindPublicKeyAndEnergy[1], kindPublicKeyAndEnergy[2]

	//
	publickey, err := base64.StdEncoding.DecodeString(publickeyS)
	if err != nil {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", publickeyS)
	}

	//
	energy, err := strconv.ParseInt(energyS, 10, 64)
	if err != nil {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", energyS)
	}

	if energy < 0 {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", energy)
	}

	return keyKind, publickey, energy, nil
}

//
func (app *Software) modifyRatifier(v kinds.RatifierModify) {
	publickey, err := cryptography.PublicKeyFromSchema(v.PublicKey)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	key := []byte(RatifierPrefix + string(publickey.Octets()))

	if v.Energy == 0 {
		//
		hasKey, err := app.status.db.Has(key)
		if err != nil {
			panic(err)
		}
		if !hasKey {
			publicStr := base64.StdEncoding.EncodeToString(publickey.Octets())
			app.tracer.Details("REDACTED", "REDACTED", publicStr)
		}
		if err = app.status.db.Erase(key); err != nil {
			panic(err)
		}
		delete(app.valueAddressToPublicKeyIndex, string(publickey.Location()))
	} else {
		//
		item := bytes.NewBuffer(make([]byte, 0))
		if err := kinds.RecordSignal(&v, item); err != nil {
			panic(err)
		}
		if err = app.status.db.Set(key, item.Bytes()); err != nil {
			panic(err)
		}
		app.valueAddressToPublicKeyIndex[string(publickey.Location())] = v.PublicKey
	}
}

func (app *Software) fetchRatifiers() (ratifiers []kinds.RatifierModify) {
	itr, err := app.status.db.Repeater(nil, nil)
	if err != nil {
		panic(err)
	}
	for ; itr.Sound(); itr.Following() {
		if isRatifierTransfer(itr.Key()) {
			ratifier := new(kinds.RatifierModify)
			err := kinds.ScanSignal(bytes.NewBuffer(itr.Item()), ratifier)
			if err != nil {
				panic(err)
			}
			ratifiers = append(ratifiers, *ratifier)
		}
	}
	if err = itr.Fault(); err != nil {
		panic(err)
	}
	return
}

//

type Status struct {
	db dbm.DB
	//
	//
	Volume   int64 `json:"volume"`
	Level int64 `json:"level"`
}

func importStatus(db dbm.DB) Status {
	var status Status
	status.db = db
	statusOctets, err := db.Get(statusKey)
	if err != nil {
		panic(err)
	}
	if len(statusOctets) == 0 {
		return status
	}
	err = json.Unmarshal(statusOctets, &status)
	if err != nil {
		panic(err)
	}
	return status
}

func persistStatus(status Status) {
	statusOctets, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	err = status.db.Set(statusKey, statusOctets)
	if err != nil {
		panic(err)
	}
}

//
//
//
//
//
func (s Status) Digest() []byte {
	applicationDigest := make([]byte, 8)
	binary.PutVarint(applicationDigest, s.Volume)
	return applicationDigest
}

func prefixKey(key []byte) []byte {
	return append(objectCouplePrefixKey, key...)
}
