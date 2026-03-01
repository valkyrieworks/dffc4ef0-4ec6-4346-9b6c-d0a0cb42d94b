package agreement

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifyStatusAbsorbAttestedLedger(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		//
		ts := freshAbsorbVerifyCollection(t)

		//
		ic := ts.CreateAbsorbNominee()

		//
		err := ts.AbsorbAttestedLedger(ic)

		//
		require.NoError(t, err)

		assert.Equal(t, ic.Altitude(), ts.cs.ObtainFinalAltitude())
		assert.NotNil(t, ts.cs.ledgerDepot.FetchLedger(ic.Altitude()))
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		ts := freshAbsorbVerifyCollection(t)

		//
		ic := ts.CreateAbsorbNominee()

		//
		err := ts.AbsorbAttestedLedger(ic)
		require.NoError(t, err)

		//
		//
		err = ts.AbsorbAttestedLedger(ic)

		//
		require.ErrorIs(t, err, FaultEarlierComprised)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		ts := freshAbsorbVerifyCollection(t)

		//
		ic := ts.CreateAbsorbNominee()
		ic.ledger.Altitude++

		//
		err := ts.AbsorbAttestedLedger(ic)

		//
		require.ErrorIs(t, err, FaultAltitudeInterval)
	})

	t.Run("REDACTED", func(t *testing.T) {
		//
		ts := freshAbsorbVerifyCollection(t)

		//
		ic := ts.CreateAbsorbNominee()
		soundLedger := ic.ledger

		//
		ic.attested = false
		ic.ledger = &kinds.Ledger{
			Heading:     soundLedger.Heading,
			Data:       soundLedger.Data,
			Proof:   soundLedger.Proof,
			FinalEndorse: nil, //
		}

		//
		err := ts.AbsorbAttestedLedger(ic)

		//
		require.ErrorIs(t, err, FaultCertification)
		require.ErrorContains(t, err, "REDACTED")
	})
}

func VerifyAbsorbNominee(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		ts := freshAbsorbVerifyCollection(t)

		for _, tt := range []struct {
			alias        string
			alter      func(ic *AbsorbNominee)
			faultIncludes string
		}{
			{
				alias:   "REDACTED",
				alter: nil,
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.ledger = nil
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.ledgerFragments = nil
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.endorse = nil
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.addnEndorse = nil
					ic.endorse.Altitude = ic.ledger.Altitude + 1
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.addnEndorse = nil
					ic.endorse.LedgerUUID = kinds.LedgerUUID{}
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.addnEndorse = &kinds.ExpandedEndorse{
						Altitude: ic.ledger.Altitude + 1,
					}
				},
				faultIncludes: "REDACTED",
			},
			{
				alias: "REDACTED",
				alter: func(ic *AbsorbNominee) {
					ic.addnEndorse = &kinds.ExpandedEndorse{
						Altitude:  ic.ledger.Altitude,
						LedgerUUID: kinds.LedgerUUID{},
					}
				},
				faultIncludes: "REDACTED",
			},
		} {
			t.Run(tt.alias, func(t *testing.T) {
				//
				ic := ts.CreateAbsorbNominee()

				if tt.alter != nil {
					tt.alter(&ic)
				}

				//
				err := ic.CertifyFundamental()

				//
				if tt.faultIncludes == "REDACTED" {
					require.NoError(t, err)
					return
				}

				require.ErrorIs(t, err, FaultCertification)
				require.ErrorContains(t, err, tt.faultIncludes)
			})
		}
	})

	t.Run("REDACTED", func(t *testing.T) {
		for _, tt := range []struct {
			alias           string
			ballotAdditions bool
			alter         func(ic *AbsorbNominee, st *sm.Status)
			faultIncludes    string
		}{
			{
				alias:           "REDACTED",
				ballotAdditions: false,
				alter:         nil,
			},
			{
				alias:           "REDACTED",
				ballotAdditions: true,
				alter:         nil,
			},
			{
				alias:           "REDACTED",
				ballotAdditions: true,
				alter: func(ic *AbsorbNominee, st *sm.Status) {
					st.AgreementSettings.Iface.BallotAdditionsActivateAltitude = 0
				},
				faultIncludes: "REDACTED",
			},
			{
				alias:           "REDACTED",
				ballotAdditions: false,
				alter: func(ic *AbsorbNominee, st *sm.Status) {
					soundLedger := ic.ledger
					ic.ledger = &kinds.Ledger{
						Heading:     soundLedger.Heading,
						Data:       soundLedger.Data,
						Proof:   soundLedger.Proof,
						FinalEndorse: nil,
					}
				},
				faultIncludes: "REDACTED",
			},
			{
				alias:           "REDACTED",
				ballotAdditions: false,
				alter: func(ic *AbsorbNominee, st *sm.Status) {
					ic.endorse.Notations[0].Notation = nil
				},
				faultIncludes: "REDACTED",
			},
			{
				alias:           "REDACTED",
				ballotAdditions: true,
				alter: func(ic *AbsorbNominee, st *sm.Status) {
					ic.addnEndorse.ExpandedNotations[0].AdditionNotation = nil
				},
				faultIncludes: "REDACTED",
			},
			{
				alias:           "REDACTED",
				ballotAdditions: true,
				alter: func(ic *AbsorbNominee, st *sm.Status) {
					ic.addnEndorse.ExpandedNotations[0].Notation = nil
				},
				faultIncludes: "REDACTED",
			},
		} {
			t.Run(tt.alias, func(t *testing.T) {
				//
				ts := freshAbsorbVerifyCollection(t)

				if tt.ballotAdditions {
					ts.cs.status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = 1
				} else {
					ts.cs.status.AgreementSettings.Iface.BallotAdditionsActivateAltitude = 0
				}

				//
				ic := ts.CreateAbsorbNominee()
				if tt.ballotAdditions {
					require.NotNil(t, ic.addnEndorse)
				} else {
					require.Nil(t, ic.addnEndorse)
				}

				//
				ic.attested = false
				validateStatus := ts.cs.status

				if tt.alter != nil {
					tt.alter(&ic, &validateStatus)
				}

				//
				err := ic.Validate(validateStatus)

				//
				if tt.faultIncludes == "REDACTED" {
					require.NoError(t, err)
					require.True(t, ic.attested)
					return
				}

				require.ErrorContains(t, err, tt.faultIncludes)
				require.False(t, ic.attested)
			})
		}
	})
}

type absorbVerifyCollection struct {
	t          *testing.T
	cs         *Status
	assessors []*assessorMock
}

func freshAbsorbVerifyCollection(t *testing.T) *absorbVerifyCollection {
	cs, assessors := arbitraryStatus(4)

	return &absorbVerifyCollection{
		t:          t,
		cs:         cs,
		assessors: assessors,
	}
}

func (ts *absorbVerifyCollection) AbsorbAttestedLedger(ic AbsorbNominee) error {
	ts.t.Helper()

	ts.cs.mtx.Lock()
	defer ts.cs.mtx.Unlock()

	return ts.cs.absorbLedger(ic)
}

func (ts *absorbVerifyCollection) CreateAbsorbNominee() AbsorbNominee {
	ts.t.Helper()

	ledger, err := ts.cs.generateNominationLedger(context.Background())
	require.NoError(ts.t, err)

	ledgerFragments, err := ledger.CreateFragmentAssign(kinds.LedgerFragmentExtentOctets)
	require.NoError(ts.t, err)

	privateItems := make([]kinds.PrivateAssessor, len(ts.assessors))
	for i, vs := range ts.assessors {
		privateItems[i] = vs.PrivateAssessor
	}

	var (
		additionsActivated = ts.cs.status.AgreementSettings.Iface.BallotAdditionsActivated(ledger.Altitude)
		successionUUID           = ts.cs.status.SuccessionUUID
		ledgerAltitude       = ledger.Altitude
		ledgerUUID           = kinds.LedgerUUID{
			Digest:          ledger.Digest(),
			FragmentAssignHeading: ledgerFragments.Heading(),
		}
	)

	var ballotAssign *kinds.BallotAssign
	if additionsActivated {
		ballotAssign = kinds.FreshExpandedBallotAssign(successionUUID, ledgerAltitude, 0, commitchema.PreendorseKind, ts.cs.Assessors)
	} else {
		ballotAssign = kinds.FreshBallotAssign(successionUUID, ledgerAltitude, 0, commitchema.PreendorseKind, ts.cs.Assessors)
	}

	for i := 0; i < len(privateItems); i++ {
		ts.assessors[i].Altitude = ledgerAltitude
		ts.assessors[i].Iteration = 0
		ballot := attestBallot(ts.assessors[i], commitchema.PreendorseKind, ledgerUUID.Digest, ledgerUUID.FragmentAssignHeading, additionsActivated)
		appended, err := ballotAssign.AppendBallot(ballot)
		require.NoError(ts.t, err)
		require.True(ts.t, appended)
	}

	addnEndorse := ballotAssign.CreateExpandedEndorse(ts.cs.status.AgreementSettings.Iface)
	endorse := addnEndorse.TowardEndorse()
	if !additionsActivated {
		addnEndorse = nil
	}

	ic, err := FreshAbsorbNominee(ledger, ledgerFragments, endorse, addnEndorse)
	require.NoError(ts.t, err, "REDACTED")
	require.NoError(ts.t, ic.Validate(ts.cs.status), "REDACTED")

	return ic
}
