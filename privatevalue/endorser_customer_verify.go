package privatevalue

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	cryptographyproto "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

type endorserVerifyScenario struct {
	successionUUID      string
	simulatePRV       kinds.PrivateAssessor
	endorserCustomer *EndorserCustomer
	endorserDaemon *EndorserDaemon
}

func obtainEndorserVerifyScenarios(t *testing.T) []endorserVerifyScenario {
	verifyScenarios := make([]endorserVerifyScenario, 0)

	//
	for _, dtc := range obtainCallerVerifyScenarios(t) {
		successionUUID := commitrand.Str(12)
		simulatePRV := kinds.FreshSimulatePRV()

		//
		sl, sd := obtainSimulateTerminals(t, dtc.location, dtc.caller)
		sc, err := FreshEndorserCustomer(sl, successionUUID)
		require.NoError(t, err)
		ss := FreshEndorserDaemon(sd, successionUUID, simulatePRV)

		err = ss.Initiate()
		require.NoError(t, err)

		tc := endorserVerifyScenario{
			successionUUID:      successionUUID,
			simulatePRV:       simulatePRV,
			endorserCustomer: sc,
			endorserDaemon: ss,
		}

		verifyScenarios = append(verifyScenarios, tc)
	}

	return verifyScenarios
}

func VerifyEndorserShutdown(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		err := tc.endorserCustomer.Shutdown()
		assert.NoError(t, err)

		err = tc.endorserDaemon.Halt()
		assert.NoError(t, err)
	}
}

func VerifyEndorserPing(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		err := tc.endorserCustomer.Ping()
		assert.NoError(t, err)
	}
}

func VerifyEndorserObtainPublicToken(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		publicToken, err := tc.endorserCustomer.ObtainPublicToken()
		require.NoError(t, err)
		anticipatedPublicToken, err := tc.simulatePRV.ObtainPublicToken()
		require.NoError(t, err)

		assert.Equal(t, anticipatedPublicToken, publicToken)

		publicToken, err = tc.endorserCustomer.ObtainPublicToken()
		require.NoError(t, err)
		anticipatedpk, err := tc.simulatePRV.ObtainPublicToken()
		require.NoError(t, err)
		anticipatedLocation := anticipatedpk.Location()

		assert.Equal(t, anticipatedLocation, publicToken.Location())
	}
}

func VerifyEndorserNomination(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		possess := &kinds.Nomination{
			Kind:      commitchema.NominationKind,
			Altitude:    1,
			Iteration:     2,
			PolicyIteration:  2,
			LedgerUUID:   kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
		}
		desire := &kinds.Nomination{
			Kind:      commitchema.NominationKind,
			Altitude:    1,
			Iteration:     2,
			PolicyIteration:  2,
			LedgerUUID:   kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
		}

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, tc.simulatePRV.AttestNomination(tc.successionUUID, desire.TowardSchema()))
		require.NoError(t, tc.endorserCustomer.AttestNomination(tc.successionUUID, possess.TowardSchema()))

		assert.Equal(t, desire.Notation, possess.Notation)
	}
}

func VerifyEndorserBallot(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		itemLocation := commitrand.Octets(security.LocatorExtent)
		desire := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, tc.simulatePRV.AttestBallot(tc.successionUUID, desire.TowardSchema()))
		require.NoError(t, tc.endorserCustomer.AttestBallot(tc.successionUUID, possess.TowardSchema()))

		assert.Equal(t, desire.Notation, possess.Notation)
	}
}

func VerifyEndorserBallotRestoreExpiration(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		itemLocation := commitrand.Octets(security.LocatorExtent)
		desire := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		time.Sleep(verifyDeadlineRetrievePersist2o3)

		require.NoError(t, tc.simulatePRV.AttestBallot(tc.successionUUID, desire.TowardSchema()))
		require.NoError(t, tc.endorserCustomer.AttestBallot(tc.successionUUID, possess.TowardSchema()))
		assert.Equal(t, desire.Notation, possess.Notation)

		//

		//
		time.Sleep(verifyDeadlineRetrievePersist2o3)

		require.NoError(t, tc.simulatePRV.AttestBallot(tc.successionUUID, desire.TowardSchema()))
		require.NoError(t, tc.endorserCustomer.AttestBallot(tc.successionUUID, possess.TowardSchema()))
		assert.Equal(t, desire.Notation, possess.Notation)
	}
}

func VerifyEndorserBallotRetainLive(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		itemLocation := commitrand.Octets(security.LocatorExtent)
		desire := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		//
		//

		//
		//
		tc.endorserDaemon.Tracer.Diagnose("REDACTED")
		time.Sleep(verifyDeadlineRetrievePersist * 3)
		tc.endorserDaemon.Tracer.Diagnose("REDACTED")

		require.NoError(t, tc.simulatePRV.AttestBallot(tc.successionUUID, desire.TowardSchema()))
		require.NoError(t, tc.endorserCustomer.AttestBallot(tc.successionUUID, possess.TowardSchema()))

		assert.Equal(t, desire.Notation, possess.Notation)
	}
}

func VerifyEndorserAttestNominationFaults(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		//
		tc.endorserDaemon.privateItem = kinds.FreshFaultingSimulatePRV()
		tc.simulatePRV = kinds.FreshFaultingSimulatePRV()

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		nomination := &kinds.Nomination{
			Kind:      commitchema.NominationKind,
			Altitude:    1,
			Iteration:     2,
			PolicyIteration:  2,
			LedgerUUID:   kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
			Notation: []byte("REDACTED"),
		}

		err := tc.endorserCustomer.AttestNomination(tc.successionUUID, nomination.TowardSchema())
		require.Equal(t, err.(*RemoteEndorserFailure).Characterization, kinds.FaultingSimulatePRVFault.Error())

		err = tc.simulatePRV.AttestNomination(tc.successionUUID, nomination.TowardSchema())
		require.Error(t, err)

		err = tc.endorserCustomer.AttestNomination(tc.successionUUID, nomination.TowardSchema())
		require.Error(t, err)
	}
}

func VerifyEndorserAttestBallotFaults(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		ts := time.Now()
		digest := commitrand.Octets(tenderminthash.Extent)
		itemLocation := commitrand.Octets(security.LocatorExtent)
		ballot := &kinds.Ballot{
			Kind:             commitchema.PreendorseKind,
			Altitude:           1,
			Iteration:            2,
			LedgerUUID:          kinds.LedgerUUID{Digest: digest, FragmentAssignHeading: kinds.FragmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			AssessorLocation: itemLocation,
			AssessorOrdinal:   1,
			Notation:        []byte("REDACTED"),
		}

		//
		tc.endorserDaemon.privateItem = kinds.FreshFaultingSimulatePRV()
		tc.simulatePRV = kinds.FreshFaultingSimulatePRV()

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		err := tc.endorserCustomer.AttestBallot(tc.successionUUID, ballot.TowardSchema())
		require.Equal(t, err.(*RemoteEndorserFailure).Characterization, kinds.FaultingSimulatePRVFault.Error())

		err = tc.simulatePRV.AttestBallot(tc.successionUUID, ballot.TowardSchema())
		require.Error(t, err)

		err = tc.endorserCustomer.AttestBallot(tc.successionUUID, ballot.TowardSchema())
		require.Error(t, err)
	}
}

func rupturedProcessor(_ kinds.PrivateAssessor, solicit privatevalueschema.Signal, _ string) (privatevalueschema.Signal, error) {
	var res privatevalueschema.Signal
	var err error

	switch r := solicit.Sum.(type) {
	//
	case *privatevalueschema.Artifact_Publictokensolicit:
		res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{PublicToken: cryptographyproto.CommonToken{}, Failure: nil})
	case *privatevalueschema.Artifact_Notateballotsolicit:
		res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{PublicToken: cryptographyproto.CommonToken{}, Failure: nil})
	case *privatevalueschema.Artifact_Notateproposalsolicit:
		res = shouldEncloseSignal(&privatevalueschema.PublicTokenReply{PublicToken: cryptographyproto.CommonToken{}, Failure: nil})
	case *privatevalueschema.Artifact_Pingsolicit:
		err, res = nil, shouldEncloseSignal(&privatevalueschema.PingReply{})
	default:
		err = fmt.Errorf("REDACTED", r)
	}

	return res, err
}

func VerifyEndorserUnforeseenReply(t *testing.T) {
	for _, tc := range obtainEndorserVerifyScenarios(t) {
		tc.endorserDaemon.privateItem = kinds.FreshSimulatePRV()
		tc.simulatePRV = kinds.FreshSimulatePRV()

		tc.endorserDaemon.AssignSolicitProcessor(rupturedProcessor)

		t.Cleanup(func() {
			if err := tc.endorserDaemon.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.endorserCustomer.Shutdown(); err != nil {
				t.Error(err)
			}
		})

		ts := time.Now()
		desire := &kinds.Ballot{Timestamp: ts, Kind: commitchema.PreendorseKind}

		e := tc.endorserCustomer.AttestBallot(tc.successionUUID, desire.TowardSchema())
		assert.ErrorIs(t, e, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"})
	}
}
