package status_test

import (
	"fmt"
	"testing"
	"time"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"
	cmtsecurity "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongstatus "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/status"
	"github.com/stretchr/testify/require"
)

//

func computeIfaceRepliesToken(altitude int64) []byte {
	return []byte(fmt.Sprintf("REDACTED", altitude))
}

var finalIfaceReplyToken = []byte("REDACTED")

var (
	_ sm.Depot    = (*VariedDepot)(nil)
	_ HeritageDepot = (*VariedDepot)(nil)
)

//
//
type VariedDepot struct {
	sm.Depot
	db dbm.DB
	sm.DepotChoices
}

//
//
func FreshVariedDepot(db dbm.DB, choices sm.DepotChoices, depot sm.Depot) *VariedDepot {
	return &VariedDepot{
		Depot:        depot,
		db:           db,
		DepotChoices: choices,
	}
}

//
type HeritageDepot interface {
	PersistIfaceReplies(altitude int64, ifaceReplies *strongstatus.HeritageIfaceReplies) error
}

//
//
//
//
func (varied VariedDepot) PersistIfaceReplies(altitude int64, ifaceReplies *strongstatus.HeritageIfaceReplies) error {
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
	if !varied.EjectIfaceReplies {
		bz, err := ifaceReplies.Serialize()
		if err != nil {
			return err
		}
		if err := varied.db.Set(computeIfaceRepliesToken(altitude), bz); err != nil {
			return err
		}
	}

	//
	//
	reply := &strongstatus.IfaceRepliesDetails{
		HeritageIfaceReplies: ifaceReplies,
		Altitude:              altitude,
	}
	bz, err := reply.Serialize()
	if err != nil {
		return err
	}

	return varied.db.AssignChronize(finalIfaceReplyToken, bz)
}

//
//
//
//
func VerifyHeritagePersistAlsoFetchCulminateLedger(t *testing.T) {
	dismantleDepressed, statusDatastore, _, depot := configureVerifyInstanceUsingDepot(t)
	defer dismantleDepressed(t)
	choices := sm.DepotChoices{
		EjectIfaceReplies: false,
	}

	altitude := int64(1)
	variedDepot := FreshVariedDepot(statusDatastore, choices, depot)

	//
	heritageIfaceReplies := freshHeritageIfaceReplies()
	err := variedDepot.PersistIfaceReplies(altitude, &heritageIfaceReplies)
	require.NoError(t, err)
	require.Equal(t, 1, len(heritageIfaceReplies.DispatchTrans))
	require.Equal(t, 1, len(heritageIfaceReplies.InitiateLedger.Incidents))
	require.Equal(t, 1, len(heritageIfaceReplies.TerminateLedger.Incidents))

	replyCulminateLedger, err := variedDepot.FetchCulminateLedgerReply(altitude)
	require.NoError(t, err)

	//
	require.NotNil(t, replyCulminateLedger.TransferOutcomes)
	require.NotNil(t, replyCulminateLedger.Incidents)
	require.NotNil(t, replyCulminateLedger.AssessorRevisions)
	require.NotNil(t, replyCulminateLedger.AgreementArgumentRevisions)
	require.Nil(t, replyCulminateLedger.PlatformDigest)

	//
	require.Equal(t, 1, len(replyCulminateLedger.TransferOutcomes))
	require.Equal(t, len(heritageIfaceReplies.DispatchTrans), len(replyCulminateLedger.TransferOutcomes))
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Cipher, replyCulminateLedger.TransferOutcomes[0].Cipher)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Data, replyCulminateLedger.TransferOutcomes[0].Data)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Log, replyCulminateLedger.TransferOutcomes[0].Log)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].FuelDesired, replyCulminateLedger.TransferOutcomes[0].FuelDesired)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].FuelUtilized, replyCulminateLedger.TransferOutcomes[0].FuelUtilized)
	require.Equal(t, len(heritageIfaceReplies.DispatchTrans[0].Incidents), len(replyCulminateLedger.TransferOutcomes[0].Incidents))
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Incidents[0].Kind, replyCulminateLedger.TransferOutcomes[0].Incidents[0].Kind)
	require.Equal(t, len(heritageIfaceReplies.DispatchTrans[0].Incidents[0].Properties), len(replyCulminateLedger.TransferOutcomes[0].Incidents[0].Properties))
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Incidents[0].Properties[0].Key, replyCulminateLedger.TransferOutcomes[0].Incidents[0].Properties[0].Key)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Incidents[0].Properties[0].Datum, replyCulminateLedger.TransferOutcomes[0].Incidents[0].Properties[0].Datum)
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Codeset, replyCulminateLedger.TransferOutcomes[0].Codeset)

	require.Equal(t, 2, len(replyCulminateLedger.Incidents))
	require.Equal(t, len(heritageIfaceReplies.InitiateLedger.Incidents)+len(heritageIfaceReplies.TerminateLedger.Incidents), len(replyCulminateLedger.Incidents))

	require.Equal(t, heritageIfaceReplies.InitiateLedger.Incidents[0].Kind, replyCulminateLedger.Incidents[0].Kind)
	require.Equal(t, len(heritageIfaceReplies.InitiateLedger.Incidents[0].Properties)+1, len(replyCulminateLedger.Incidents[0].Properties)) //
	require.Equal(t, heritageIfaceReplies.InitiateLedger.Incidents[0].Properties[0].Key, replyCulminateLedger.Incidents[0].Properties[0].Key)
	require.Equal(t, heritageIfaceReplies.InitiateLedger.Incidents[0].Properties[0].Datum, replyCulminateLedger.Incidents[0].Properties[0].Datum)

	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Ledger.MaximumOctets, replyCulminateLedger.AgreementArgumentRevisions.Ledger.MaximumOctets)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Ledger.MaximumFuel, replyCulminateLedger.AgreementArgumentRevisions.Ledger.MaximumFuel)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Proof.MaximumLifespanCountLedgers, replyCulminateLedger.AgreementArgumentRevisions.Proof.MaximumLifespanCountLedgers)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Proof.MaximumLifespanInterval, replyCulminateLedger.AgreementArgumentRevisions.Proof.MaximumLifespanInterval)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Proof.MaximumOctets, replyCulminateLedger.AgreementArgumentRevisions.Proof.MaximumOctets)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Assessor.PublicTokenKinds, replyCulminateLedger.AgreementArgumentRevisions.Assessor.PublicTokenKinds)
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AgreementArgumentRevisions.Edition.App, replyCulminateLedger.AgreementArgumentRevisions.Edition.App)

	require.Nil(t, replyCulminateLedger.AgreementArgumentRevisions.Iface)
	require.Nil(t, replyCulminateLedger.PlatformDigest)

	require.Equal(t, len(heritageIfaceReplies.TerminateLedger.AssessorRevisions), len(replyCulminateLedger.AssessorRevisions))
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AssessorRevisions[0].Potency, replyCulminateLedger.AssessorRevisions[0].Potency)

	//
	require.Equal(t, heritageIfaceReplies.TerminateLedger.AssessorRevisions[0].PublicToken.ObtainEdwards25519(), replyCulminateLedger.AssessorRevisions[0].PublicToken.ObtainEdwards25519())

	//
	altitude = int64(2)
	heritageIfaceReplies = freshHeritageIfaceRepliesUsingNothingAreas()
	require.Equal(t, 1, len(heritageIfaceReplies.DispatchTrans))
	require.Equal(t, 1, len(heritageIfaceReplies.InitiateLedger.Incidents))
	require.Nil(t, heritageIfaceReplies.TerminateLedger)
	err = variedDepot.PersistIfaceReplies(altitude, &heritageIfaceReplies)
	require.NoError(t, err)
	replyCulminateLedger, err = variedDepot.FetchCulminateLedgerReply(altitude)
	require.NoError(t, err)

	require.Equal(t, len(heritageIfaceReplies.DispatchTrans), len(replyCulminateLedger.TransferOutcomes))
	require.Equal(t, heritageIfaceReplies.DispatchTrans[0].Text(), replyCulminateLedger.TransferOutcomes[0].Text())
	require.Equal(t, len(heritageIfaceReplies.InitiateLedger.Incidents), len(replyCulminateLedger.Incidents))
}

//
func freshHeritageIfaceReplies() strongstatus.HeritageIfaceReplies {
	incidentProperty := iface.IncidentProperty{
		Key:   "REDACTED",
		Datum: "REDACTED",
	}

	dispatchTransferIncident := iface.Incident{
		Kind:       "REDACTED",
		Properties: []iface.IncidentProperty{incidentProperty},
	}

	terminateLedgerIncident := iface.Incident{
		Kind:       "REDACTED",
		Properties: []iface.IncidentProperty{incidentProperty},
	}

	commenceLedgerIncident := iface.Incident{
		Kind:       "REDACTED",
		Properties: []iface.IncidentProperty{incidentProperty},
	}

	replyDispatchTransfer := iface.InvokeTransferOutcome{
		Cipher:   iface.CipherKindOKAY,
		Incidents: []iface.Incident{dispatchTransferIncident},
	}

	assessorRevisions := []iface.AssessorRevise{{
		PublicToken: cmtsecurity.CommonToken{Sum: &cmtsecurity.Commonkey_Edwards25519{Edwards25519: make([]byte, 1)}},
		Potency:  int64(10),
	}}

	agreementParameters := &commitchema.AgreementSettings{
		Ledger: &commitchema.LedgerParameters{
			MaximumOctets: int64(100000),
			MaximumFuel:   int64(10000),
		},
		Proof: &commitchema.ProofParameters{
			MaximumLifespanCountLedgers: int64(10),
			MaximumLifespanInterval:  time.Duration(1000),
			MaximumOctets:        int64(10000),
		},
		Assessor: &commitchema.AssessorParameters{
			PublicTokenKinds: []string{"REDACTED"},
		},
		Edition: &commitchema.EditionParameters{
			App: uint64(10),
		},
	}

	//
	heritageIfaceReplies := strongstatus.HeritageIfaceReplies{
		DispatchTrans: []*iface.InvokeTransferOutcome{
			&replyDispatchTransfer,
		},
		TerminateLedger: &strongstatus.ReplyTerminateLedger{
			Incidents:                []iface.Incident{terminateLedgerIncident},
			AgreementArgumentRevisions: agreementParameters,
			AssessorRevisions:      assessorRevisions,
		},
		InitiateLedger: &strongstatus.ReplyInitiateLedger{
			Incidents: []iface.Incident{commenceLedgerIncident},
		},
	}
	return heritageIfaceReplies
}

//
func freshHeritageIfaceRepliesUsingNothingAreas() strongstatus.HeritageIfaceReplies {
	incidentProperty := iface.IncidentProperty{
		Key:   "REDACTED",
		Datum: "REDACTED",
	}

	dispatchTransferIncident := iface.Incident{
		Kind:       "REDACTED",
		Properties: []iface.IncidentProperty{incidentProperty},
	}

	commenceLedgerIncident := iface.Incident{
		Kind:       "REDACTED",
		Properties: []iface.IncidentProperty{incidentProperty},
	}

	replyDispatchTransfer := iface.InvokeTransferOutcome{
		Cipher:   iface.CipherKindOKAY,
		Incidents: []iface.Incident{dispatchTransferIncident},
	}

	//
	heritageIfaceReplies := strongstatus.HeritageIfaceReplies{
		DispatchTrans: []*iface.InvokeTransferOutcome{
			&replyDispatchTransfer,
		},
		InitiateLedger: &strongstatus.ReplyInitiateLedger{
			Incidents: []iface.Incident{commenceLedgerIncident},
		},
	}
	return heritageIfaceReplies
}
