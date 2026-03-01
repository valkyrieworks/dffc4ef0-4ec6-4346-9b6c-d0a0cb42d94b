package agreement

import (
	"encoding/hex"
	"math"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/digits"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	strongmindcons "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/agreement"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func VerifySignalTowardSchema(t *testing.T) {
	psh := kinds.FragmentAssignHeading{
		Sum: 1,
		Digest:  commitrand.Octets(32),
	}
	bufferPush := psh.TowardSchema()
	bi := kinds.LedgerUUID{
		Digest:          commitrand.Octets(32),
		FragmentAssignHeading: psh,
	}
	bufferBigint := bi.TowardSchema()
	digits := digits.FreshDigitCollection(1)
	bufferDigits := digits.TowardSchema()

	fragments := kinds.Fragment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: hashmap.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: commitrand.Octets(32),
			Kin:    [][]byte{},
		},
	}
	bufferFragments, err := fragments.TowardSchema()
	require.NoError(t, err)

	nomination := kinds.Nomination{
		Kind:      commitchema.NominationKind,
		Altitude:    1,
		Iteration:     1,
		PolicyIteration:  1,
		LedgerUUID:   bi,
		Timestamp: time.Now(),
		Notation: commitrand.Octets(20),
	}
	bufferNomination := nomination.TowardSchema()

	ballot := kinds.CreateBallotNegativeFailure(
		t,
		kinds.FreshSimulatePRV(),
		"REDACTED",
		0,
		1,
		0,
		commitchema.PreendorseKind,
		bi,
		time.Now(),
	)
	bufferBallot := ballot.TowardSchema()

	verifiesScenarios := []struct {
		verifyAlias string
		msg      Signal
		desire     proto.Message
		desireFault  bool
	}{
		{
			"REDACTED", &FreshIterationPhaseSignal{
				Altitude:                2,
				Iteration:                 1,
				Phase:                  1,
				MomentsBecauseInitiateMoment: 1,
				FinalEndorseIteration:       2,
			}, &strongmindcons.FreshIterationPhase{
				Altitude:                2,
				Iteration:                 1,
				Phase:                  1,
				MomentsBecauseInitiateMoment: 1,
				FinalEndorseIteration:       2,
			},

			false,
		},

		{
			"REDACTED", &FreshSoundLedgerSignal{
				Altitude:             1,
				Iteration:              1,
				LedgerFragmentAssignHeading: psh,
				LedgerFragments:         digits,
				EqualsEndorse:           false,
			}, &strongmindcons.FreshSoundLedger{
				Altitude:             1,
				Iteration:              1,
				LedgerFragmentAssignHeading: bufferPush,
				LedgerFragments:         bufferDigits,
				EqualsEndorse:           false,
			},

			false,
		},
		{
			"REDACTED", &LedgerFragmentSignal{
				Altitude: 100,
				Iteration:  1,
				Fragment:   &fragments,
			}, &strongmindcons.LedgerFragment{
				Altitude: 100,
				Iteration:  1,
				Fragment:   *bufferFragments,
			},

			false,
		},
		{
			"REDACTED", &NominationPolicySignal{
				Altitude:           1,
				NominationPolicyIteration: 1,
				NominationPolicy:      digits,
			}, &strongmindcons.NominationPolicy{
				Altitude:           1,
				NominationPolicyIteration: 1,
				NominationPolicy:      *bufferDigits,
			},
			false,
		},
		{
			"REDACTED", &NominationSignal{
				Nomination: &nomination,
			}, &strongmindcons.Nomination{
				Nomination: *bufferNomination,
			},

			false,
		},
		{
			"REDACTED", &BallotSignal{
				Ballot: ballot,
			}, &strongmindcons.Ballot{
				Ballot: bufferBallot,
			},

			false,
		},
		{
			"REDACTED", &BallotAssignMajor23signal{
				Altitude:  1,
				Iteration:   1,
				Kind:    1,
				LedgerUUID: bi,
			}, &strongmindcons.BallotAssignMajor23{
				Altitude:  1,
				Iteration:   1,
				Kind:    1,
				LedgerUUID: bufferBigint,
			},

			false,
		},
		{
			"REDACTED", &BallotAssignDigitsSignal{
				Altitude:  1,
				Iteration:   1,
				Kind:    1,
				LedgerUUID: bi,
				Ballots:   digits,
			}, &strongmindcons.BallotAssignDigits{
				Altitude:  1,
				Iteration:   1,
				Kind:    1,
				LedgerUUID: bufferBigint,
				Ballots:   *bufferDigits,
			},

			false,
		},
		{"REDACTED", nil, &strongmindcons.Signal{}, true},
	}
	for _, tt := range verifiesScenarios {
		t.Run(tt.verifyAlias, func(t *testing.T) {
			pb, err := SignalTowardSchema(tt.msg)
			if tt.desireFault == true {
				assert.Equal(t, err != nil, tt.desireFault)
				return
			}
			assert.EqualValues(t, tt.desire, pb, tt.verifyAlias)

			msg, err := SignalOriginatingSchema(pb)

			if !tt.desireFault {
				require.NoError(t, err)
				bcm := assert.Equal(t, tt.msg, msg, tt.verifyAlias)
				assert.True(t, bcm, tt.verifyAlias)
			} else {
				require.Error(t, err, tt.verifyAlias)
			}
		})
	}
}

func VerifyJournalSignalSchema(t *testing.T) {
	fragments := kinds.Fragment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: hashmap.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: commitrand.Octets(32),
			Kin:    [][]byte{},
		},
	}
	bufferFragments, err := fragments.TowardSchema()
	require.NoError(t, err)

	verifiesScenarios := []struct {
		verifyAlias string
		msg      JournalSignal
		desire     *strongmindcons.JournalSignal
		desireFault  bool
	}{
		{"REDACTED", kinds.IncidentDataIterationStatus{
			Altitude: 2,
			Iteration:  1,
			Phase:   "REDACTED",
		}, &strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Incidentiterationstate{
				IncidentDataIterationStatus: &commitchema.IncidentDataIterationStatus{
					Altitude: 2,
					Iteration:  1,
					Phase:   "REDACTED",
				},
			},
		}, false},
		{"REDACTED", signalDetails{
			Msg: &LedgerFragmentSignal{
				Altitude: 100,
				Iteration:  1,
				Fragment:   &fragments,
			},
			NodeUUID: p2p.ID("REDACTED"),
		}, &strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Signalinfo{
				SignalDetails: &strongmindcons.SignalDetails{
					Msg: strongmindcons.Signal{
						Sum: &strongmindcons.Signal_Ledgerfragment{
							LedgerFragment: &strongmindcons.LedgerFragment{
								Altitude: 100,
								Iteration:  1,
								Fragment:   *bufferFragments,
							},
						},
					},
					NodeUUID: "REDACTED",
				},
			},
		}, false},
		{"REDACTED", deadlineDetails{
			Interval: time.Duration(100),
			Altitude:   1,
			Iteration:    1,
			Phase:     1,
		}, &strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Alarminfo{
				DeadlineDetails: &strongmindcons.DeadlineDetails{
					Interval: time.Duration(100),
					Altitude:   1,
					Iteration:    1,
					Phase:     1,
				},
			},
		}, false},
		{"REDACTED", TerminateAltitudeSignal{
			Altitude: 1,
		}, &strongmindcons.JournalSignal{
			Sum: &strongmindcons.Walrecord_Finalheight{
				TerminateAltitude: &strongmindcons.TerminateAltitude{
					Altitude: 1,
				},
			},
		}, false},
		{"REDACTED", nil, &strongmindcons.JournalSignal{}, true},
	}
	for _, tt := range verifiesScenarios {
		t.Run(tt.verifyAlias, func(t *testing.T) {
			pb, err := JournalTowardSchema(tt.msg)
			if tt.desireFault == true {
				assert.Equal(t, err != nil, tt.desireFault)
				return
			}
			assert.EqualValues(t, tt.desire, pb, tt.verifyAlias)

			msg, err := JournalOriginatingSchema(pb)

			if !tt.desireFault {
				require.NoError(t, err)
				assert.Equal(t, tt.msg, msg, tt.verifyAlias) //
			} else {
				require.Error(t, err, tt.verifyAlias)
			}
		})
	}
}

//
func VerifyConsensusSignalsArrays(t *testing.T) {
	time := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	psh := kinds.FragmentAssignHeading{
		Sum: 1,
		Digest:  []byte("REDACTED"),
	}
	bufferPush := psh.TowardSchema()

	bi := kinds.LedgerUUID{
		Digest:          []byte("REDACTED"),
		FragmentAssignHeading: psh,
	}
	bufferBigint := bi.TowardSchema()
	digits := digits.FreshDigitCollection(1)
	bufferDigits := digits.TowardSchema()

	fragments := kinds.Fragment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: hashmap.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: []byte("REDACTED"),
			Kin:    [][]byte{},
		},
	}
	bufferFragments, err := fragments.TowardSchema()
	require.NoError(t, err)

	nomination := kinds.Nomination{
		Kind:      commitchema.NominationKind,
		Altitude:    1,
		Iteration:     1,
		PolicyIteration:  1,
		LedgerUUID:   bi,
		Timestamp: time,
		Notation: []byte("REDACTED"),
	}
	bufferNomination := nomination.TowardSchema()

	v := &kinds.Ballot{
		AssessorLocation: []byte("REDACTED"),
		AssessorOrdinal:   1,
		Altitude:           1,
		Iteration:            0,
		Timestamp:        time,
		Kind:             commitchema.PreendorseKind,
		LedgerUUID:          bi,
	}
	vpb := v.TowardSchema()
	v.Addition = []byte("REDACTED")
	vadditionBuffer := v.TowardSchema()

	verifyScenarios := []struct {
		verifyAlias string
		cnSignal     proto.Message
		expirationOctets string
	}{
		{"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Newcyclephase{FreshIterationPhase: &strongmindcons.FreshIterationPhase{
			Altitude:                1,
			Iteration:                 1,
			Phase:                  1,
			MomentsBecauseInitiateMoment: 1,
			FinalEndorseIteration:       1,
		}}}, "REDACTED"},
		{"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Newcyclephase{FreshIterationPhase: &strongmindcons.FreshIterationPhase{
			Altitude:                math.MaxInt64,
			Iteration:                 math.MaxInt32,
			Phase:                  math.MaxUint32,
			MomentsBecauseInitiateMoment: math.MaxInt64,
			FinalEndorseIteration:       math.MaxInt32,
		}}}, "REDACTED"},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Newvalidledger{
				FreshSoundLedger: &strongmindcons.FreshSoundLedger{
					Altitude: 1, Iteration: 1, LedgerFragmentAssignHeading: bufferPush, LedgerFragments: bufferDigits, EqualsEndorse: false,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Nomination{Nomination: &strongmindcons.Nomination{Nomination: *bufferNomination}}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Proposalpolicy{
				NominationPolicy: &strongmindcons.NominationPolicy{Altitude: 1, NominationPolicyIteration: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Ledgerfragment{
				LedgerFragment: &strongmindcons.LedgerFragment{Altitude: 1, Iteration: 1, Fragment: *bufferFragments},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Ballot{
				Ballot: &strongmindcons.Ballot{Ballot: vpb},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Ballot{
				Ballot: &strongmindcons.Ballot{Ballot: vadditionBuffer},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Hasballot{
				OwnsBallot: &strongmindcons.OwnsBallot{Altitude: 1, Iteration: 1, Kind: commitchema.PreballotKind, Ordinal: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Hasballot{
				OwnsBallot: &strongmindcons.OwnsBallot{
					Altitude: math.MaxInt64, Iteration: math.MaxInt32,
					Kind: commitchema.PreballotKind, Ordinal: math.MaxInt32,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Ballotsetmaj23{
				BallotAssignMajor23: &strongmindcons.BallotAssignMajor23{Altitude: 1, Iteration: 1, Kind: commitchema.PreballotKind, LedgerUUID: bufferBigint},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &strongmindcons.Signal{Sum: &strongmindcons.Signal_Ballotsetdigits{
				BallotAssignDigits: &strongmindcons.BallotAssignDigits{Altitude: 1, Iteration: 1, Kind: commitchema.PreballotKind, LedgerUUID: bufferBigint, Ballots: *bufferDigits},
			}},
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyAlias, func(t *testing.T) {
			bz, err := proto.Marshal(tc.cnSignal)
			require.NoError(t, err)

			require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz))
		})
	}
}
