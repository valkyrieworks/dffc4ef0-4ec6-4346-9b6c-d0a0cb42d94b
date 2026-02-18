package status_test

import (
	"fmt"
	"testing"
	"time"

	dbm "github.com/valkyrieworks/-db"
	cmtcrypto "github.com/valkyrieworks/schema/consensuscore/vault"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	sm "github.com/valkyrieworks/status"

	iface "github.com/valkyrieworks/iface/kinds"
	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	"github.com/stretchr/testify/require"
)

//

func computeIfaceRepliesKey(level int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", level))
}

var finalIfaceReplyKey = []byte("REDACTED")

var (
	_ sm.Depot    = (*MultipleDepot)(nil)
	_ PastDepot = (*MultipleDepot)(nil)
)

//
//
type MultipleDepot struct {
	sm.Depot
	db dbm.DB
	sm.DepotSettings
}

//
//
func NewMultipleDepot(db dbm.DB, options sm.DepotSettings, depot sm.Depot) *MultipleDepot {
	return &MultipleDepot{
		Depot:        depot,
		db:           db,
		DepotSettings: options,
	}
}

//
type PastDepot interface {
	PersistIfaceReplies(level int64, ifaceReplies *cometstatus.PastIfaceReplies) error
}

//
//
//
//
func (multiple MultipleDepot) PersistIfaceReplies(level int64, ifaceReplies *cometstatus.PastIfaceReplies) error {
	var dtrans []*iface.InvokeTransferOutcome
	//
	for _, tx := range ifaceReplies.DispatchTrans {
		if tx != nil {
			dtrans = append(dtrans, tx)
		}
	}
	ifaceReplies.DispatchTrans = dtrans

	//
	//
	if !multiple.DropIfaceReplies {
		bz, err := ifaceReplies.Serialize()
		if err != nil {
			return err
		}
		if err := multiple.db.Set(computeIfaceRepliesKey(level), bz); err != nil {
			return err
		}
	}

	//
	//
	reply := &cometstatus.IfaceRepliesDetails{
		PastIfaceReplies: ifaceReplies,
		Level:              level,
	}
	bz, err := reply.Serialize()
	if err != nil {
		return err
	}

	return multiple.db.CollectionAlign(finalIfaceReplyKey, bz)
}

//
//
//
//
func VerifyPastPersistAndImportCompleteLedger(t *testing.T) {
	dismantleBelow, statusStore, _, depot := configureVerifyScenarioWithDepot(t)
	defer dismantleBelow(t)
	options := sm.DepotSettings{
		DropIfaceReplies: false,
	}

	level := int64(1)
	multipleDepot := NewMultipleDepot(statusStore, options, depot)

	//
	pastIfaceReplies := newPastIfaceReplies()
	err := multipleDepot.PersistIfaceReplies(level, &pastIfaceReplies)
	require.NoError(t, err)
	require.Equal(t, 1, len(pastIfaceReplies.DispatchTrans))
	require.Equal(t, 1, len(pastIfaceReplies.InitiateLedger.Events))
	require.Equal(t, 1, len(pastIfaceReplies.TerminateLedger.Events))

	replyCompleteLedger, err := multipleDepot.ImportCompleteLedgerReply(level)
	require.NoError(t, err)

	//
	require.NotNil(t, replyCompleteLedger.TransOutcomes)
	require.NotNil(t, replyCompleteLedger.Events)
	require.NotNil(t, replyCompleteLedger.RatifierRefreshes)
	require.NotNil(t, replyCompleteLedger.AgreementArgumentRefreshes)
	require.Nil(t, replyCompleteLedger.ApplicationDigest)

	//
	require.Equal(t, 1, len(replyCompleteLedger.TransOutcomes))
	require.Equal(t, len(pastIfaceReplies.DispatchTrans), len(replyCompleteLedger.TransOutcomes))
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Code, replyCompleteLedger.TransOutcomes[0].Code)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Data, replyCompleteLedger.TransOutcomes[0].Data)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Log, replyCompleteLedger.TransOutcomes[0].Log)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].FuelDesired, replyCompleteLedger.TransOutcomes[0].FuelDesired)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].FuelApplied, replyCompleteLedger.TransOutcomes[0].FuelApplied)
	require.Equal(t, len(pastIfaceReplies.DispatchTrans[0].Events), len(replyCompleteLedger.TransOutcomes[0].Events))
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Events[0].Kind, replyCompleteLedger.TransOutcomes[0].Events[0].Kind)
	require.Equal(t, len(pastIfaceReplies.DispatchTrans[0].Events[0].Properties), len(replyCompleteLedger.TransOutcomes[0].Events[0].Properties))
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Events[0].Properties[0].Key, replyCompleteLedger.TransOutcomes[0].Events[0].Properties[0].Key)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Events[0].Properties[0].Item, replyCompleteLedger.TransOutcomes[0].Events[0].Properties[0].Item)
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].Codex, replyCompleteLedger.TransOutcomes[0].Codex)

	require.Equal(t, 2, len(replyCompleteLedger.Events))
	require.Equal(t, len(pastIfaceReplies.InitiateLedger.Events)+len(pastIfaceReplies.TerminateLedger.Events), len(replyCompleteLedger.Events))

	require.Equal(t, pastIfaceReplies.InitiateLedger.Events[0].Kind, replyCompleteLedger.Events[0].Kind)
	require.Equal(t, len(pastIfaceReplies.InitiateLedger.Events[0].Properties)+1, len(replyCompleteLedger.Events[0].Properties)) //
	require.Equal(t, pastIfaceReplies.InitiateLedger.Events[0].Properties[0].Key, replyCompleteLedger.Events[0].Properties[0].Key)
	require.Equal(t, pastIfaceReplies.InitiateLedger.Events[0].Properties[0].Item, replyCompleteLedger.Events[0].Properties[0].Item)

	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Ledger.MaximumOctets, replyCompleteLedger.AgreementArgumentRefreshes.Ledger.MaximumOctets)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Ledger.MaximumFuel, replyCompleteLedger.AgreementArgumentRefreshes.Ledger.MaximumFuel)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Proof.MaximumDurationCountLedgers, replyCompleteLedger.AgreementArgumentRefreshes.Proof.MaximumDurationCountLedgers)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Proof.MaximumDurationPeriod, replyCompleteLedger.AgreementArgumentRefreshes.Proof.MaximumDurationPeriod)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Proof.MaximumOctets, replyCompleteLedger.AgreementArgumentRefreshes.Proof.MaximumOctets)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Ratifier.PublicKeyKinds, replyCompleteLedger.AgreementArgumentRefreshes.Ratifier.PublicKeyKinds)
	require.Equal(t, pastIfaceReplies.TerminateLedger.AgreementArgumentRefreshes.Release.App, replyCompleteLedger.AgreementArgumentRefreshes.Release.App)

	require.Nil(t, replyCompleteLedger.AgreementArgumentRefreshes.Iface)
	require.Nil(t, replyCompleteLedger.ApplicationDigest)

	require.Equal(t, len(pastIfaceReplies.TerminateLedger.RatifierRefreshes), len(replyCompleteLedger.RatifierRefreshes))
	require.Equal(t, pastIfaceReplies.TerminateLedger.RatifierRefreshes[0].Energy, replyCompleteLedger.RatifierRefreshes[0].Energy)

	//
	require.Equal(t, pastIfaceReplies.TerminateLedger.RatifierRefreshes[0].PublicKey.FetchEd25519(), replyCompleteLedger.RatifierRefreshes[0].PublicKey.FetchEd25519())

	//
	level = int64(2)
	pastIfaceReplies = newPastIfaceRepliesWithNilAttributes()
	require.Equal(t, 1, len(pastIfaceReplies.DispatchTrans))
	require.Equal(t, 1, len(pastIfaceReplies.InitiateLedger.Events))
	require.Nil(t, pastIfaceReplies.TerminateLedger)
	err = multipleDepot.PersistIfaceReplies(level, &pastIfaceReplies)
	require.NoError(t, err)
	replyCompleteLedger, err = multipleDepot.ImportCompleteLedgerReply(level)
	require.NoError(t, err)

	require.Equal(t, len(pastIfaceReplies.DispatchTrans), len(replyCompleteLedger.TransOutcomes))
	require.Equal(t, pastIfaceReplies.DispatchTrans[0].String(), replyCompleteLedger.TransOutcomes[0].String())
	require.Equal(t, len(pastIfaceReplies.InitiateLedger.Events), len(replyCompleteLedger.Events))
}

//
func newPastIfaceReplies() cometstatus.PastIfaceReplies {
	eventProperty := iface.EventProperty{
		Key:   "REDACTED",
		Item: "REDACTED",
	}

	dispatchTransferEvent := iface.Event{
		Kind:       "REDACTED",
		Properties: []iface.EventProperty{eventProperty},
	}

	terminateLedgerEvent := iface.Event{
		Kind:       "REDACTED",
		Properties: []iface.EventProperty{eventProperty},
	}

	initiateLedgerEvent := iface.Event{
		Kind:       "REDACTED",
		Properties: []iface.EventProperty{eventProperty},
	}

	replyDispatchTransfer := iface.InvokeTransferOutcome{
		Code:   iface.CodeKindSuccess,
		Events: []iface.Event{dispatchTransferEvent},
	}

	ratifierRefreshes := []iface.RatifierModify{{
		PublicKey: cmtcrypto.PublicKey{Sum: &cmtcrypto.Publickey_Ed25519{Ed25519: make([]byte, 1)}},
		Energy:  int64(10),
	}}

	agreementOptions := &engineproto.AgreementOptions{
		Ledger: &engineproto.LedgerOptions{
			MaximumOctets: int64(100000),
			MaximumFuel:   int64(10000),
		},
		Proof: &engineproto.ProofOptions{
			MaximumDurationCountLedgers: int64(10),
			MaximumDurationPeriod:  time.Duration(1000),
			MaximumOctets:        int64(10000),
		},
		Ratifier: &engineproto.RatifierOptions{
			PublicKeyKinds: []string{"REDACTED"},
		},
		Release: &engineproto.ReleaseOptions{
			App: uint64(10),
		},
	}

	//
	pastIfaceReplies := cometstatus.PastIfaceReplies{
		DispatchTrans: []*iface.InvokeTransferOutcome{
			&replyDispatchTransfer,
		},
		TerminateLedger: &cometstatus.AnswerTerminateLedger{
			Events:                []iface.Event{terminateLedgerEvent},
			AgreementArgumentRefreshes: agreementOptions,
			RatifierRefreshes:      ratifierRefreshes,
		},
		InitiateLedger: &cometstatus.AnswerInitiateLedger{
			Events: []iface.Event{initiateLedgerEvent},
		},
	}
	return pastIfaceReplies
}

//
func newPastIfaceRepliesWithNilAttributes() cometstatus.PastIfaceReplies {
	eventProperty := iface.EventProperty{
		Key:   "REDACTED",
		Item: "REDACTED",
	}

	dispatchTransferEvent := iface.Event{
		Kind:       "REDACTED",
		Properties: []iface.EventProperty{eventProperty},
	}

	initiateLedgerEvent := iface.Event{
		Kind:       "REDACTED",
		Properties: []iface.EventProperty{eventProperty},
	}

	replyDispatchTransfer := iface.InvokeTransferOutcome{
		Code:   iface.CodeKindSuccess,
		Events: []iface.Event{dispatchTransferEvent},
	}

	//
	pastIfaceReplies := cometstatus.PastIfaceReplies{
		DispatchTrans: []*iface.InvokeTransferOutcome{
			&replyDispatchTransfer,
		},
		InitiateLedger: &cometstatus.AnswerInitiateLedger{
			Events: []iface.Event{initiateLedgerEvent},
		},
	}
	return pastIfaceReplies
}
