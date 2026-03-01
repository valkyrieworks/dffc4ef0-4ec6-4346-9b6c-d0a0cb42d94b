package statedepot

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	cryptography "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	cryptographyproto "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

var (
	statusToken        = []byte("REDACTED")
	tokvalDuoHeadingToken = []byte("REDACTED")
)

const (
	AssessorHeading        = "REDACTED"
	PlatformEdition      uint64 = 1
)

var _ kinds.Platform = (*Platform)(nil)

//
//
//
type Platform struct {
	kinds.FoundationPlatform

	status        Status
	PreserveLedgers int64 //
	arrangedTrans    [][]byte
	tracer       log.Tracer

	//
	itemRevisions         []kinds.AssessorRevise
	itemLocationTowardPublicTokenIndex map[string]cryptographyproto.CommonToken

	//
	//
	produceLedgerIncidents bool
}

//
func FreshPlatform(db dbm.DB) *Platform {
	return &Platform{
		tracer:             log.FreshNooperationTracer(),
		status:              fetchStatus(db),
		itemLocationTowardPublicTokenIndex: make(map[string]cryptographyproto.CommonToken),
	}
}

//
func FreshEnduringPlatform(datastorePath string) *Platform {
	alias := "REDACTED"
	db, err := dbm.FreshProceedStratumDatastore(alias, datastorePath)
	if err != nil {
		panic(fmt.Errorf("REDACTED", datastorePath, err))
	}
	return FreshPlatform(db)
}

//
//
func FreshInsideRamPlatform() *Platform {
	return FreshPlatform(dbm.FreshMemoryDatastore())
}

func (app *Platform) AssignProduceLedgerIncidents() {
	app.produceLedgerIncidents = true
}

//
//
//
//
func (app *Platform) Details(context.Context, *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	//
	if len(app.itemLocationTowardPublicTokenIndex) == 0 && app.status.Altitude > 0 {
		assessors := app.obtainAssessors()
		for _, v := range assessors {
			publickey, err := cryptography.PublicTokenOriginatingSchema(v.PublicToken)
			if err != nil {
				panic(fmt.Errorf("REDACTED", err))
			}
			app.itemLocationTowardPublicTokenIndex[string(publickey.Location())] = v.PublicToken
		}
	}

	return &kinds.ReplyDetails{
		Data:             fmt.Sprintf("REDACTED", app.status.Extent),
		Edition:          edition.IfaceEdition,
		PlatformEdition:       PlatformEdition,
		FinalLedgerAltitude:  app.status.Altitude,
		FinalLedgerPlatformDigest: app.status.Digest(),
	}, nil
}

//
//
//
func (app *Platform) InitializeSuccession(_ context.Context, req *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	for _, v := range req.Assessors {
		app.reviseAssessor(v)
	}
	platformDigest := make([]byte, 8)
	binary.PutVarint(platformDigest, app.status.Extent)
	return &kinds.ReplyInitializeSuccession{
		PlatformDigest: platformDigest,
	}, nil
}

//
//
//
//
//
//
//
func (app *Platform) InspectTransfer(_ context.Context, req *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	//
	if equalsAssessorTransfer(req.Tx) {
		if _, _, _, err := analyzeAssessorTransfer(req.Tx); err != nil {
			//
			return &kinds.ReplyInspectTransfer{Cipher: CipherKindUnfitTransferLayout}, nil
		}
	} else if !equalsSoundTransfer(req.Tx) {
		return &kinds.ReplyInspectTransfer{Cipher: CipherKindUnfitTransferLayout}, nil
	}

	return &kinds.ReplyInspectTransfer{Cipher: CipherKindOKAY, FuelDesired: 1}, nil
}

//
//
//
func equalsSoundTransfer(tx []byte) bool {
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
func (app *Platform) ArrangeNomination(ctx context.Context, req *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	return &kinds.ReplyArrangeNomination{Txs: app.layoutTrans(ctx, req.Txs)}, nil
}

//
//
func (app *Platform) layoutTrans(ctx context.Context, ledgerData [][]byte) [][]byte {
	txs := make([][]byte, 0, len(ledgerData))
	for _, tx := range ledgerData {
		if reply, err := app.InspectTransfer(ctx, &kinds.SolicitInspectTransfer{Tx: tx}); err == nil && reply.Cipher == CipherKindOKAY {
			txs = append(txs, bytes.Replace(tx, []byte("REDACTED"), []byte("REDACTED"), 1))
		}
	}
	return txs
}

//
//
func (app *Platform) HandleNomination(ctx context.Context, req *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	for _, tx := range req.Txs {
		//
		if reply, err := app.InspectTransfer(ctx, &kinds.SolicitInspectTransfer{Tx: tx}); err != nil || reply.Cipher != CipherKindOKAY {
			return &kinds.ReplyHandleNomination{Condition: kinds.Responseexecuteitem_DECLINE}, nil
		}
	}
	return &kinds.ReplyHandleNomination{Condition: kinds.Responseexecuteitem_EMBRACE}, nil
}

//
//
//
//
func (app *Platform) CulminateLedger(_ context.Context, req *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	//
	app.itemRevisions = make([]kinds.AssessorRevise, 0)
	app.arrangedTrans = make([][]byte, 0)

	//
	for _, ev := range req.Malpractice {
		if ev.Kind == kinds.Malfunctionkind_REPLICATED_BALLOT {
			location := string(ev.Assessor.Location)
			if publicToken, ok := app.itemLocationTowardPublicTokenIndex[location]; ok {
				app.itemRevisions = append(app.itemRevisions, kinds.AssessorRevise{
					PublicToken: publicToken,
					Potency:  ev.Assessor.Potency - 1,
				})
				app.tracer.Details("REDACTED",
					"REDACTED", location)
			} else {
				panic(fmt.Errorf("REDACTED", location))
			}
		}
	}

	answerTrans := make([]*kinds.InvokeTransferOutcome, len(req.Txs))
	for i, tx := range req.Txs {
		if equalsAssessorTransfer(tx) {
			tokenKind, publicToken, potency, err := analyzeAssessorTransfer(tx)
			if err != nil {
				panic(err)
			}
			app.itemRevisions = append(app.itemRevisions, kinds.ReviseAssessor(publicToken, potency, tokenKind))
		} else {
			app.arrangedTrans = append(app.arrangedTrans, tx)
		}

		var key, datum string
		fragments := bytes.Split(tx, []byte("REDACTED"))
		if len(fragments) == 2 {
			key, datum = string(fragments[0]), string(fragments[1])
		} else {
			key, datum = string(tx), string(tx)
		}
		answerTrans[i] = &kinds.InvokeTransferOutcome{
			Cipher: CipherKindOKAY,
			//
			Incidents: []kinds.Incident{
				{
					Kind: "REDACTED",
					Properties: []kinds.IncidentProperty{
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Datum: key, Ordinal: true},
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: false},
					},
				},
				{
					Kind: "REDACTED",
					Properties: []kinds.IncidentProperty{
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Datum: datum, Ordinal: true},
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: true},
						{Key: "REDACTED", Datum: "REDACTED", Ordinal: false},
					},
				},
			},
		}
		app.status.Extent++
	}

	app.status.Altitude = req.Altitude

	reply := &kinds.ReplyCulminateLedger{TransferOutcomes: answerTrans, AssessorRevisions: app.itemRevisions, PlatformDigest: app.status.Digest()}
	if !app.produceLedgerIncidents {
		return reply, nil
	}
	if app.status.Altitude%2 == 0 {
		reply.Incidents = []kinds.Incident{
			{
				Kind: "REDACTED",
				Properties: []kinds.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
			{
				Kind: "REDACTED",
				Properties: []kinds.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
				},
			},
		}
	} else {
		reply.Incidents = []kinds.Incident{
			{
				Kind: "REDACTED",
				Properties: []kinds.IncidentProperty{
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
						Ordinal: true,
					},
					{
						Key:   "REDACTED",
						Datum: "REDACTED",
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
func (app *Platform) Endorse(context.Context, *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
	//
	for _, itemRevise := range app.itemRevisions {
		app.reviseAssessor(itemRevise)
	}

	//
	for _, tx := range app.arrangedTrans {
		fragments := bytes.Split(tx, []byte("REDACTED"))
		if len(fragments) != 2 {
			panic(fmt.Sprintf("REDACTED", len(fragments), fragments))
		}
		key, datum := string(fragments[0]), string(fragments[1])
		err := app.status.db.Set(headingToken([]byte(key)), []byte(datum))
		if err != nil {
			panic(err)
		}
	}

	//
	persistStatus(app.status)

	reply := &kinds.ReplyEndorse{}
	if app.PreserveLedgers > 0 && app.status.Altitude >= app.PreserveLedgers {
		reply.PreserveAltitude = app.status.Altitude - app.PreserveLedgers + 1
	}
	return reply, nil
}

//
func (app *Platform) Inquire(_ context.Context, requestInquire *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	outcomeInquire := &kinds.ReplyInquire{}

	if requestInquire.Route == "REDACTED" {
		key := []byte(AssessorHeading + string(requestInquire.Data))
		datum, err := app.status.db.Get(key)
		if err != nil {
			panic(err)
		}

		return &kinds.ReplyInquire{
			Key:   requestInquire.Data,
			Datum: datum,
		}, nil
	}

	if requestInquire.Validate {
		datum, err := app.status.db.Get(headingToken(requestInquire.Data))
		if err != nil {
			panic(err)
		}

		if datum == nil {
			outcomeInquire.Log = "REDACTED"
		} else {
			outcomeInquire.Log = "REDACTED"
		}
		outcomeInquire.Ordinal = -1 //
		outcomeInquire.Key = requestInquire.Data
		outcomeInquire.Datum = datum
		outcomeInquire.Altitude = app.status.Altitude

		return outcomeInquire, nil
	}

	outcomeInquire.Key = requestInquire.Data
	datum, err := app.status.db.Get(headingToken(requestInquire.Data))
	if err != nil {
		panic(err)
	}
	if datum == nil {
		outcomeInquire.Log = "REDACTED"
	} else {
		outcomeInquire.Log = "REDACTED"
	}
	outcomeInquire.Datum = datum
	outcomeInquire.Altitude = app.status.Altitude

	return outcomeInquire, nil
}

func (app *Platform) Shutdown() error {
	return app.status.db.Shutdown()
}

func equalsAssessorTransfer(tx []byte) bool {
	return strings.HasPrefix(string(tx), AssessorHeading)
}

func analyzeAssessorTransfer(tx []byte) (string, []byte, int64, error) {
	tx = tx[len(AssessorHeading):]

	//
	kindPublicTokenAlsoPotency := strings.Split(string(tx), "REDACTED")
	if len(kindPublicTokenAlsoPotency) != 3 {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", kindPublicTokenAlsoPotency)
	}
	tokenKind, publickeySTR, potencySTR := kindPublicTokenAlsoPotency[0], kindPublicTokenAlsoPotency[1], kindPublicTokenAlsoPotency[2]

	//
	publickey, err := base64.StdEncoding.DecodeString(publickeySTR)
	if err != nil {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", publickeySTR)
	}

	//
	potency, err := strconv.ParseInt(potencySTR, 10, 64)
	if err != nil {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", potencySTR)
	}

	if potency < 0 {
		return "REDACTED", nil, 0, fmt.Errorf("REDACTED", potency)
	}

	return tokenKind, publickey, potency, nil
}

//
func (app *Platform) reviseAssessor(v kinds.AssessorRevise) {
	publickey, err := cryptography.PublicTokenOriginatingSchema(v.PublicToken)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	key := []byte(AssessorHeading + string(publickey.Octets()))

	if v.Potency == 0 {
		//
		ownsToken, err := app.status.db.Has(key)
		if err != nil {
			panic(err)
		}
		if !ownsToken {
			publicTxt := base64.StdEncoding.EncodeToString(publickey.Octets())
			app.tracer.Details("REDACTED", "REDACTED", publicTxt)
		}
		if err = app.status.db.Erase(key); err != nil {
			panic(err)
		}
		delete(app.itemLocationTowardPublicTokenIndex, string(publickey.Location()))
	} else {
		//
		datum := bytes.NewBuffer(make([]byte, 0))
		if err := kinds.PersistArtifact(&v, datum); err != nil {
			panic(err)
		}
		if err = app.status.db.Set(key, datum.Bytes()); err != nil {
			panic(err)
		}
		app.itemLocationTowardPublicTokenIndex[string(publickey.Location())] = v.PublicToken
	}
}

func (app *Platform) obtainAssessors() (assessors []kinds.AssessorRevise) {
	itr, err := app.status.db.Traverser(nil, nil)
	if err != nil {
		panic(err)
	}
	for ; itr.Sound(); itr.Following() {
		if equalsAssessorTransfer(itr.Key()) {
			assessor := new(kinds.AssessorRevise)
			err := kinds.FetchArtifact(bytes.NewBuffer(itr.Datum()), assessor)
			if err != nil {
				panic(err)
			}
			assessors = append(assessors, *assessor)
		}
	}
	if err = itr.Failure(); err != nil {
		panic(err)
	}
	return
}

//

type Status struct {
	db dbm.DB
	//
	//
	Extent   int64 `json:"extent"`
	Altitude int64 `json:"altitude"`
}

func fetchStatus(db dbm.DB) Status {
	var status Status
	status.db = db
	statusOctets, err := db.Get(statusToken)
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
	err = status.db.Set(statusToken, statusOctets)
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
	platformDigest := make([]byte, 8)
	binary.PutVarint(platformDigest, s.Extent)
	return platformDigest
}

func headingToken(key []byte) []byte {
	return append(tokvalDuoHeadingToken, key...)
}
