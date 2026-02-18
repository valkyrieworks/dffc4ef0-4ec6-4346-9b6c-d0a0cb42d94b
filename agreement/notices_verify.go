package agreement

import (
	"encoding/hex"
	"math"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/utils/bits"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	cometconnect "github.com/valkyrieworks/schema/consensuscore/agreement"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

func VerifyMessageToSchema(t *testing.T) {
	psh := kinds.SegmentAssignHeading{
		Sum: 1,
		Digest:  engineseed.Octets(32),
	}
	pbPsh := psh.ToSchema()
	bi := kinds.LedgerUID{
		Digest:          engineseed.Octets(32),
		SegmentAssignHeading: psh,
	}
	pbBi := bi.ToSchema()
	bits := bits.NewBitList(1)
	pbBits := bits.ToSchema()

	segments := kinds.Segment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: merkle.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: engineseed.Octets(32),
			Kin:    [][]byte{},
		},
	}
	pbSegments, err := segments.ToSchema()
	require.NoError(t, err)

	nomination := kinds.Nomination{
		Kind:      engineproto.NominationKind,
		Level:    1,
		Cycle:     1,
		POLDuration:  1,
		LedgerUID:   bi,
		Timestamp: time.Now(),
		Autograph: engineseed.Octets(20),
	}
	pbNomination := nomination.ToSchema()

	ballot := kinds.CreateBallotNoFault(
		t,
		kinds.NewEmulatePV(),
		"REDACTED",
		0,
		1,
		0,
		engineproto.PreendorseKind,
		bi,
		time.Now(),
	)
	pbBallot := ballot.ToSchema()

	verifiesScenarios := []struct {
		verifyLabel string
		msg      Signal
		desire     proto.Message
		desireErr  bool
	}{
		{
			"REDACTED", &NewDurationPhaseSignal{
				Level:                2,
				Cycle:                 1,
				Phase:                  1,
				MomentsSinceBeginTime: 1,
				FinalEndorseDuration:       2,
			}, &cometconnect.NewDurationPhase{
				Level:                2,
				Cycle:                 1,
				Phase:                  1,
				MomentsSinceBeginTime: 1,
				FinalEndorseDuration:       2,
			},

			false,
		},

		{
			"REDACTED", &NewSoundLedgerSignal{
				Level:             1,
				Cycle:              1,
				LedgerSegmentAssignHeading: psh,
				LedgerSegments:         bits,
				IsEndorse:           false,
			}, &cometconnect.NewSoundLedger{
				Level:             1,
				Cycle:              1,
				LedgerSegmentAssignHeading: pbPsh,
				LedgerSegments:         pbBits,
				IsEndorse:           false,
			},

			false,
		},
		{
			"REDACTED", &LedgerSegmentSignal{
				Level: 100,
				Cycle:  1,
				Segment:   &segments,
			}, &cometconnect.LedgerSegment{
				Level: 100,
				Cycle:  1,
				Segment:   *pbSegments,
			},

			false,
		},
		{
			"REDACTED", &NominationPOLSignal{
				Level:           1,
				NominationPOLDuration: 1,
				NominationPOL:      bits,
			}, &cometconnect.NominationPOL{
				Level:           1,
				NominationPolDuration: 1,
				NominationPol:      *pbBits,
			},
			false,
		},
		{
			"REDACTED", &NominationSignal{
				Nomination: &nomination,
			}, &cometconnect.Nomination{
				Nomination: *pbNomination,
			},

			false,
		},
		{
			"REDACTED", &BallotSignal{
				Ballot: ballot,
			}, &cometconnect.Ballot{
				Ballot: pbBallot,
			},

			false,
		},
		{
			"REDACTED", &BallotAssignMaj23signal{
				Level:  1,
				Cycle:   1,
				Kind:    1,
				LedgerUID: bi,
			}, &cometconnect.BallotAssignMaj23{
				Level:  1,
				Cycle:   1,
				Kind:    1,
				LedgerUID: pbBi,
			},

			false,
		},
		{
			"REDACTED", &BallotAssignBitsSignal{
				Level:  1,
				Cycle:   1,
				Kind:    1,
				LedgerUID: bi,
				Ballots:   bits,
			}, &cometconnect.BallotAssignBits{
				Level:  1,
				Cycle:   1,
				Kind:    1,
				LedgerUID: pbBi,
				Ballots:   *pbBits,
			},

			false,
		},
		{"REDACTED", nil, &cometconnect.Signal{}, true},
	}
	for _, tt := range verifiesScenarios {
		t.Run(tt.verifyLabel, func(t *testing.T) {
			pb, err := MessageToSchema(tt.msg)
			if tt.desireErr == true {
				assert.Equal(t, err != nil, tt.desireErr)
				return
			}
			assert.EqualValues(t, tt.desire, pb, tt.verifyLabel)

			msg, err := MessageFromSchema(pb)

			if !tt.desireErr {
				require.NoError(t, err)
				bcm := assert.Equal(t, tt.msg, msg, tt.verifyLabel)
				assert.True(t, bcm, tt.verifyLabel)
			} else {
				require.Error(t, err, tt.verifyLabel)
			}
		})
	}
}

func VerifyJournalMessageSchema(t *testing.T) {
	segments := kinds.Segment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: merkle.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: engineseed.Octets(32),
			Kin:    [][]byte{},
		},
	}
	pbSegments, err := segments.ToSchema()
	require.NoError(t, err)

	verifiesScenarios := []struct {
		verifyLabel string
		msg      JournalSignal
		desire     *cometconnect.JournalSignal
		desireErr  bool
	}{
		{"REDACTED", kinds.EventDataDurationStatus{
			Level: 2,
			Cycle:  1,
			Phase:   "REDACTED",
		}, &cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Signaldatadurationstate{
				EventDataDurationStatus: &engineproto.EventDataDurationStatus{
					Level: 2,
					Cycle:  1,
					Phase:   "REDACTED",
				},
			},
		}, false},
		{"REDACTED", messageDetails{
			Msg: &LedgerSegmentSignal{
				Level: 100,
				Cycle:  1,
				Segment:   &segments,
			},
			NodeUID: p2p.ID("REDACTED"),
		}, &cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Signaldetails{
				MessageDetails: &cometconnect.MessageDetails{
					Msg: cometconnect.Signal{
						Sum: &cometconnect.Signal_Ledgersection{
							LedgerSegment: &cometconnect.LedgerSegment{
								Level: 100,
								Cycle:  1,
								Segment:   *pbSegments,
							},
						},
					},
					NodeUID: "REDACTED",
				},
			},
		}, false},
		{"REDACTED", deadlineDetails{
			Period: time.Duration(100),
			Level:   1,
			Cycle:    1,
			Phase:     1,
		}, &cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Deadlinedetails{
				DeadlineDetails: &cometconnect.DeadlineDetails{
					Period: time.Duration(100),
					Level:   1,
					Cycle:    1,
					Phase:     1,
				},
			},
		}, false},
		{"REDACTED", TerminateLevelSignal{
			Level: 1,
		}, &cometconnect.JournalSignal{
			Sum: &cometconnect.Journalsignal_Finallayer{
				TerminateLevel: &cometconnect.TerminateLevel{
					Level: 1,
				},
			},
		}, false},
		{"REDACTED", nil, &cometconnect.JournalSignal{}, true},
	}
	for _, tt := range verifiesScenarios {
		t.Run(tt.verifyLabel, func(t *testing.T) {
			pb, err := JournalToSchema(tt.msg)
			if tt.desireErr == true {
				assert.Equal(t, err != nil, tt.desireErr)
				return
			}
			assert.EqualValues(t, tt.desire, pb, tt.verifyLabel)

			msg, err := JournalFromSchema(pb)

			if !tt.desireErr {
				require.NoError(t, err)
				assert.Equal(t, tt.msg, msg, tt.verifyLabel) //
			} else {
				require.Error(t, err, tt.verifyLabel)
			}
		})
	}
}

//
func VerifyConstNoticesArrays(t *testing.T) {
	date := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	psh := kinds.SegmentAssignHeading{
		Sum: 1,
		Digest:  []byte("REDACTED"),
	}
	pbPsh := psh.ToSchema()

	bi := kinds.LedgerUID{
		Digest:          []byte("REDACTED"),
		SegmentAssignHeading: psh,
	}
	pbBi := bi.ToSchema()
	bits := bits.NewBitList(1)
	pbBits := bits.ToSchema()

	segments := kinds.Segment{
		Ordinal: 1,
		Octets: []byte("REDACTED"),
		Attestation: merkle.Attestation{
			Sum:    1,
			Ordinal:    1,
			NodeDigest: []byte("REDACTED"),
			Kin:    [][]byte{},
		},
	}
	pbSegments, err := segments.ToSchema()
	require.NoError(t, err)

	nomination := kinds.Nomination{
		Kind:      engineproto.NominationKind,
		Level:    1,
		Cycle:     1,
		POLDuration:  1,
		LedgerUID:   bi,
		Timestamp: date,
		Autograph: []byte("REDACTED"),
	}
	pbNomination := nomination.ToSchema()

	v := &kinds.Ballot{
		RatifierLocation: []byte("REDACTED"),
		RatifierOrdinal:   1,
		Level:           1,
		Cycle:            0,
		Timestamp:        date,
		Kind:             engineproto.PreendorseKind,
		LedgerUID:          bi,
	}
	vpb := v.ToSchema()
	v.Addition = []byte("REDACTED")
	vextPb := v.ToSchema()

	verifyScenarios := []struct {
		verifyLabel string
		cMessage     proto.Message
		expirationOctets string
	}{
		{"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Newepochphase{NewDurationPhase: &cometconnect.NewDurationPhase{
			Level:                1,
			Cycle:                 1,
			Phase:                  1,
			MomentsSinceBeginTime: 1,
			FinalEndorseDuration:       1,
		}}}, "REDACTED"},
		{"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Newepochphase{NewDurationPhase: &cometconnect.NewDurationPhase{
			Level:                math.MaxInt64,
			Cycle:                 math.MaxInt32,
			Phase:                  math.MaxUint32,
			MomentsSinceBeginTime: math.MaxInt64,
			FinalEndorseDuration:       math.MaxInt32,
		}}}, "REDACTED"},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Newvalidledger{
				NewSoundLedger: &cometconnect.NewSoundLedger{
					Level: 1, Cycle: 1, LedgerSegmentAssignHeading: pbPsh, LedgerSegments: pbBits, IsEndorse: false,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Nomination{Nomination: &cometconnect.Nomination{Nomination: *pbNomination}}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Nominationpol{
				NominationPol: &cometconnect.NominationPOL{Level: 1, NominationPolDuration: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Ledgersection{
				LedgerSegment: &cometconnect.LedgerSegment{Level: 1, Cycle: 1, Segment: *pbSegments},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Ballot{
				Ballot: &cometconnect.Ballot{Ballot: vpb},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Ballot{
				Ballot: &cometconnect.Ballot{Ballot: vextPb},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Hasballot{
				HasBallot: &cometconnect.HasBallot{Level: 1, Cycle: 1, Kind: engineproto.PreballotKind, Ordinal: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Hasballot{
				HasBallot: &cometconnect.HasBallot{
					Level: math.MaxInt64, Cycle: math.MaxInt32,
					Kind: engineproto.PreballotKind, Ordinal: math.MaxInt32,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Ballotsetmaj23{
				BallotAssignMaj23: &cometconnect.BallotAssignMaj23{Level: 1, Cycle: 1, Kind: engineproto.PreballotKind, LedgerUID: pbBi},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &cometconnect.Signal{Sum: &cometconnect.Signal_Ballotsetbits{
				BallotAssignBits: &cometconnect.BallotAssignBits{Level: 1, Cycle: 1, Kind: engineproto.PreballotKind, LedgerUID: pbBi, Ballots: *pbBits},
			}},
			"REDACTED",
		},
	}

	for _, tc := range verifyScenarios {
		t.Run(tc.verifyLabel, func(t *testing.T) {
			bz, err := proto.Marshal(tc.cMessage)
			require.NoError(t, err)

			require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz))
		})
	}
}
