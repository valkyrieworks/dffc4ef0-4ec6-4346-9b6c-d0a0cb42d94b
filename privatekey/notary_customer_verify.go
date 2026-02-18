package privatekey

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	engineseed "github.com/valkyrieworks/utils/random"
	cryptography "github.com/valkyrieworks/schema/consensuscore/vault"
	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

type notaryVerifyScenario struct {
	ledgerUID      string
	emulatePV       kinds.PrivateRatifier
	notaryCustomer *NotaryCustomer
	notaryHost *NotaryHost
}

func fetchNotaryVerifyScenarios(t *testing.T) []notaryVerifyScenario {
	verifyScenarios := make([]notaryVerifyScenario, 0)

	//
	for _, dtc := range fetchCallerVerifyScenarios(t) {
		ledgerUID := engineseed.Str(12)
		emulatePV := kinds.NewEmulatePV()

		//
		sl, sd := fetchEmulateTermini(t, dtc.address, dtc.caller)
		sc, err := NewNotaryCustomer(sl, ledgerUID)
		require.NoError(t, err)
		ss := NewNotaryHost(sd, ledgerUID, emulatePV)

		err = ss.Begin()
		require.NoError(t, err)

		tc := notaryVerifyScenario{
			ledgerUID:      ledgerUID,
			emulatePV:       emulatePV,
			notaryCustomer: sc,
			notaryHost: ss,
		}

		verifyScenarios = append(verifyScenarios, tc)
	}

	return verifyScenarios
}

func VerifyNotaryEnd(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		err := tc.notaryCustomer.End()
		assert.NoError(t, err)

		err = tc.notaryHost.Halt()
		assert.NoError(t, err)
	}
}

func VerifyNotaryPing(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		err := tc.notaryCustomer.Ping()
		assert.NoError(t, err)
	}
}

func VerifyNotaryFetchPublicKey(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		publicKey, err := tc.notaryCustomer.FetchPublicKey()
		require.NoError(t, err)
		anticipatedPublicKey, err := tc.emulatePV.FetchPublicKey()
		require.NoError(t, err)

		assert.Equal(t, anticipatedPublicKey, publicKey)

		publicKey, err = tc.notaryCustomer.FetchPublicKey()
		require.NoError(t, err)
		anticipatedkey, err := tc.emulatePV.FetchPublicKey()
		require.NoError(t, err)
		anticipatedAddress := anticipatedkey.Location()

		assert.Equal(t, anticipatedAddress, publicKey.Location())
	}
}

func VerifyNotaryNomination(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		possess := &kinds.Nomination{
			Kind:      engineproto.NominationKind,
			Level:    1,
			Cycle:     2,
			POLDuration:  2,
			LedgerUID:   kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
		}
		desire := &kinds.Nomination{
			Kind:      engineproto.NominationKind,
			Level:    1,
			Cycle:     2,
			POLDuration:  2,
			LedgerUID:   kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
		}

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, tc.emulatePV.AttestNomination(tc.ledgerUID, desire.ToSchema()))
		require.NoError(t, tc.notaryCustomer.AttestNomination(tc.ledgerUID, possess.ToSchema()))

		assert.Equal(t, desire.Autograph, possess.Autograph)
	}
}

func VerifyNotaryBallot(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		valueAddress := engineseed.Octets(vault.LocationVolume)
		desire := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		require.NoError(t, tc.emulatePV.AttestBallot(tc.ledgerUID, desire.ToSchema()))
		require.NoError(t, tc.notaryCustomer.AttestBallot(tc.ledgerUID, possess.ToSchema()))

		assert.Equal(t, desire.Autograph, possess.Autograph)
	}
}

func VerifyNotaryBallotRestoreDeadline(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		valueAddress := engineseed.Octets(vault.LocationVolume)
		desire := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		time.Sleep(verifyDeadlineScanRecord2o3)

		require.NoError(t, tc.emulatePV.AttestBallot(tc.ledgerUID, desire.ToSchema()))
		require.NoError(t, tc.notaryCustomer.AttestBallot(tc.ledgerUID, possess.ToSchema()))
		assert.Equal(t, desire.Autograph, possess.Autograph)

		//

		//
		time.Sleep(verifyDeadlineScanRecord2o3)

		require.NoError(t, tc.emulatePV.AttestBallot(tc.ledgerUID, desire.ToSchema()))
		require.NoError(t, tc.notaryCustomer.AttestBallot(tc.ledgerUID, possess.ToSchema()))
		assert.Equal(t, desire.Autograph, possess.Autograph)
	}
}

func VerifyNotaryBallotRetainLive(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		valueAddress := engineseed.Octets(vault.LocationVolume)
		desire := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		possess := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
		}

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		//
		//

		//
		//
		tc.notaryHost.Tracer.Diagnose("REDACTED")
		time.Sleep(verifyDeadlineScanRecord * 3)
		tc.notaryHost.Tracer.Diagnose("REDACTED")

		require.NoError(t, tc.emulatePV.AttestBallot(tc.ledgerUID, desire.ToSchema()))
		require.NoError(t, tc.notaryCustomer.AttestBallot(tc.ledgerUID, possess.ToSchema()))

		assert.Equal(t, desire.Autograph, possess.Autograph)
	}
}

func VerifyNotaryAttestNominationFaults(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		//
		tc.notaryHost.privateValue = kinds.NewFaultingEmulatePV()
		tc.emulatePV = kinds.NewFaultingEmulatePV()

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		nomination := &kinds.Nomination{
			Kind:      engineproto.NominationKind,
			Level:    1,
			Cycle:     2,
			POLDuration:  2,
			LedgerUID:   kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp: ts,
			Autograph: []byte("REDACTED"),
		}

		err := tc.notaryCustomer.AttestNomination(tc.ledgerUID, nomination.ToSchema())
		require.Equal(t, err.(*DistantNotaryFault).Summary, kinds.FaultingEmulatePVErr.Error())

		err = tc.emulatePV.AttestNomination(tc.ledgerUID, nomination.ToSchema())
		require.Error(t, err)

		err = tc.notaryCustomer.AttestNomination(tc.ledgerUID, nomination.ToSchema())
		require.Error(t, err)
	}
}

func VerifyNotaryAttestBallotFaults(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		ts := time.Now()
		digest := engineseed.Octets(comethash.Volume)
		valueAddress := engineseed.Octets(vault.LocationVolume)
		ballot := &kinds.Ballot{
			Kind:             engineproto.PreendorseKind,
			Level:           1,
			Cycle:            2,
			LedgerUID:          kinds.LedgerUID{Digest: digest, SegmentAssignHeading: kinds.SegmentAssignHeading{Digest: digest, Sum: 2}},
			Timestamp:        ts,
			RatifierLocation: valueAddress,
			RatifierOrdinal:   1,
			Autograph:        []byte("REDACTED"),
		}

		//
		tc.notaryHost.privateValue = kinds.NewFaultingEmulatePV()
		tc.emulatePV = kinds.NewFaultingEmulatePV()

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		err := tc.notaryCustomer.AttestBallot(tc.ledgerUID, ballot.ToSchema())
		require.Equal(t, err.(*DistantNotaryFault).Summary, kinds.FaultingEmulatePVErr.Error())

		err = tc.emulatePV.AttestBallot(tc.ledgerUID, ballot.ToSchema())
		require.Error(t, err)

		err = tc.notaryCustomer.AttestBallot(tc.ledgerUID, ballot.ToSchema())
		require.Error(t, err)
	}
}

func faultyManager(_ kinds.PrivateRatifier, query privatekeyproto.Signal, _ string) (privatekeyproto.Signal, error) {
	var res privatekeyproto.Signal
	var err error

	switch r := query.Sum.(type) {
	//
	case *privatekeyproto.Signal_Publickeyquery:
		res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{PublicKey: cryptography.PublicKey{}, Fault: nil})
	case *privatekeyproto.Signal_Attestballotquery:
		res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{PublicKey: cryptography.PublicKey{}, Fault: nil})
	case *privatekeyproto.Signal_Attestproposalquery:
		res = shouldEncloseMessage(&privatekeyproto.PublicKeyAnswer{PublicKey: cryptography.PublicKey{}, Fault: nil})
	case *privatekeyproto.Signal_Pingquery:
		err, res = nil, shouldEncloseMessage(&privatekeyproto.PingAnswer{})
	default:
		err = fmt.Errorf("REDACTED", r)
	}

	return res, err
}

func VerifyNotaryUnforeseenReply(t *testing.T) {
	for _, tc := range fetchNotaryVerifyScenarios(t) {
		tc.notaryHost.privateValue = kinds.NewEmulatePV()
		tc.emulatePV = kinds.NewEmulatePV()

		tc.notaryHost.CollectionQueryManager(faultyManager)

		t.Cleanup(func() {
			if err := tc.notaryHost.Halt(); err != nil {
				t.Error(err)
			}
		})
		t.Cleanup(func() {
			if err := tc.notaryCustomer.End(); err != nil {
				t.Error(err)
			}
		})

		ts := time.Now()
		desire := &kinds.Ballot{Timestamp: ts, Kind: engineproto.PreendorseKind}

		e := tc.notaryCustomer.AttestBallot(tc.ledgerUID, desire.ToSchema())
		assert.ErrorIs(t, e, cometfaults.ErrMandatoryField{Field: "REDACTED"})
	}
}
