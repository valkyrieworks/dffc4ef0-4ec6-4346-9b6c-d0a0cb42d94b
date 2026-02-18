package agreement

import (
	"encoding/hex"
	"math"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/security/hashing"
	"github.com/valkyrieworks/utils/units"
	ctrng "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	enginecons "github.com/valkyrieworks/schema/consensuscore/agreement"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

func TestMsgToProto(t *testing.T) {
	psh := kinds.PartSetHeader{
		Total: 1,
		Hash:  ctrng.Bytes(32),
	}
	pbPsh := psh.ToProto()
	bi := kinds.BlockID{
		Hash:          ctrng.Bytes(32),
		PartSetHeader: psh,
	}
	pbBi := bi.ToProto()
	bits := units.NewBitArray(1)
	pbBits := units.ToProto()

	parts := kinds.Part{
		Index: 1,
		Bytes: []byte("REDACTED"),
		Proof: hashing.Proof{
			Total:    1,
			Index:    1,
			LeafHash: ctrng.Bytes(32),
			Aunts:    [][]byte{},
		},
	}
	pbParts, err := parts.ToProto()
	require.NoError(t, err)

	proposal := kinds.Proposal{
		Type:      ctschema.ProposalType,
		Height:    1,
		Round:     1,
		POLRound:  1,
		BlockID:   bi,
		Timestamp: time.Now(),
		Signature: ctrng.Bytes(20),
	}
	pbProposal := proposal.ToProto()

	vote := kinds.MakeVoteNoError(
		t,
		kinds.NewMockPV(),
		"REDACTED",
		0,
		1,
		0,
		ctschema.PrecommitType,
		bi,
		time.Now(),
	)
	pbVote := vote.ToProto()

	testsCases := []struct {
		testName string
		msg      Message
		want     proto.Message
		wantErr  bool
	}{
		{
			"REDACTED", &NewRoundStepMessage{
				Height:                2,
				Round:                 1,
				Step:                  1,
				SecondsSinceStartTime: 1,
				LastCommitRound:       2,
			}, &enginecons.NewRoundStep{
				Height:                2,
				Round:                 1,
				Step:                  1,
				SecondsSinceStartTime: 1,
				LastCommitRound:       2,
			},

			false,
		},

		{
			"REDACTED", &NewValidBlockMessage{
				Height:             1,
				Round:              1,
				BlockPartSetHeader: psh,
				BlockParts:         bits,
				IsCommit:           false,
			}, &enginecons.NewValidBlock{
				Height:             1,
				Round:              1,
				BlockPartSetHeader: pbPsh,
				BlockParts:         pbBits,
				IsCommit:           false,
			},

			false,
		},
		{
			"REDACTED", &BlockPartMessage{
				Height: 100,
				Round:  1,
				Part:   &parts,
			}, &enginecons.BlockPart{
				Height: 100,
				Round:  1,
				Part:   *pbParts,
			},

			false,
		},
		{
			"REDACTED", &ProposalPOLMessage{
				Height:           1,
				ProposalPOLRound: 1,
				ProposalPOL:      bits,
			}, &enginecons.ProposalPOL{
				Height:           1,
				ProposalPolRound: 1,
				ProposalPol:      *pbBits,
			},
			false,
		},
		{
			"REDACTED", &ProposalMessage{
				Proposal: &proposal,
			}, &enginecons.Proposal{
				Proposal: *pbProposal,
			},

			false,
		},
		{
			"REDACTED", &VoteMessage{
				Vote: vote,
			}, &enginecons.Vote{
				Vote: pbVote,
			},

			false,
		},
		{
			"REDACTED", &VoteSetMaj23Message{
				Height:  1,
				Round:   1,
				Type:    1,
				BlockID: bi,
			}, &enginecons.VoteSetMaj23{
				Height:  1,
				Round:   1,
				Type:    1,
				BlockID: pbBi,
			},

			false,
		},
		{
			"REDACTED", &VoteSetBitsMessage{
				Height:  1,
				Round:   1,
				Type:    1,
				BlockID: bi,
				Votes:   bits,
			}, &enginecons.VoteSetBits{
				Height:  1,
				Round:   1,
				Type:    1,
				BlockID: pbBi,
				Votes:   *pbBits,
			},

			false,
		},
		{"REDACTED", nil, &enginecons.Message{}, true},
	}
	for _, tt := range testsCases {
		t.Run(tt.testName, func(t *testing.T) {
			pb, err := MsgToProto(tt.msg)
			if tt.wantErr == true {
				assert.Equal(t, err != nil, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, pb, tt.testName)

			msg, err := MsgFromProto(pb)

			if !tt.wantErr {
				require.NoError(t, err)
				bcm := assert.Equal(t, tt.msg, msg, tt.testName)
				assert.True(t, bcm, tt.testName)
			} else {
				require.Error(t, err, tt.testName)
			}
		})
	}
}

func TestWALMsgProto(t *testing.T) {
	parts := kinds.Part{
		Index: 1,
		Bytes: []byte("REDACTED"),
		Proof: hashing.Proof{
			Total:    1,
			Index:    1,
			LeafHash: ctrng.Bytes(32),
			Aunts:    [][]byte{},
		},
	}
	pbParts, err := parts.ToProto()
	require.NoError(t, err)

	testsCases := []struct {
		testName string
		msg      WALMessage
		want     *enginecons.WALMessage
		wantErr  bool
	}{
		{"REDACTED", kinds.EventDataRoundState{
			Height: 2,
			Round:  1,
			Step:   "REDACTED",
		}, &enginecons.WALMessage{
			Sum: &enginecons.WALMessage_EventDataRoundState{
				EventDataRoundState: &ctschema.EventDataRoundState{
					Height: 2,
					Round:  1,
					Step:   "REDACTED",
				},
			},
		}, false},
		{"REDACTED", msgInfo{
			Msg: &BlockPartMessage{
				Height: 100,
				Round:  1,
				Part:   &parts,
			},
			PeerID: p2p.ID("REDACTED"),
		}, &enginecons.WALMessage{
			Sum: &enginecons.WALMessage_MsgInfo{
				MsgInfo: &enginecons.MsgInfo{
					Msg: enginecons.Message{
						Sum: &enginecons.Message_BlockPart{
							BlockPart: &enginecons.BlockPart{
								Height: 100,
								Round:  1,
								Part:   *pbParts,
							},
						},
					},
					PeerID: "REDACTED",
				},
			},
		}, false},
		{"REDACTED", timeoutInfo{
			Duration: time.Duration(100),
			Height:   1,
			Round:    1,
			Step:     1,
		}, &enginecons.WALMessage{
			Sum: &enginecons.WALMessage_TimeoutInfo{
				TimeoutInfo: &enginecons.TimeoutInfo{
					Duration: time.Duration(100),
					Height:   1,
					Round:    1,
					Step:     1,
				},
			},
		}, false},
		{"REDACTED", EndHeightMessage{
			Height: 1,
		}, &enginecons.WALMessage{
			Sum: &enginecons.WALMessage_EndHeight{
				EndHeight: &enginecons.EndHeight{
					Height: 1,
				},
			},
		}, false},
		{"REDACTED", nil, &enginecons.WALMessage{}, true},
	}
	for _, tt := range testsCases {
		t.Run(tt.testName, func(t *testing.T) {
			pb, err := WALToProto(tt.msg)
			if tt.wantErr == true {
				assert.Equal(t, err != nil, tt.wantErr)
				return
			}
			assert.EqualValues(t, tt.want, pb, tt.testName)

			msg, err := WALFromProto(pb)

			if !tt.wantErr {
				require.NoError(t, err)
				assert.Equal(t, tt.msg, msg, tt.testName) //
			} else {
				require.Error(t, err, tt.testName)
			}
		})
	}
}

//
func TestConsMsgsVectors(t *testing.T) {
	date := time.Date(2018, 8, 30, 12, 0, 0, 0, time.UTC)
	psh := kinds.PartSetHeader{
		Total: 1,
		Hash:  []byte("REDACTED"),
	}
	pbPsh := psh.ToProto()

	bi := kinds.BlockID{
		Hash:          []byte("REDACTED"),
		PartSetHeader: psh,
	}
	pbBi := bi.ToProto()
	bits := units.NewBitArray(1)
	pbBits := units.ToProto()

	parts := kinds.Part{
		Index: 1,
		Bytes: []byte("REDACTED"),
		Proof: hashing.Proof{
			Total:    1,
			Index:    1,
			LeafHash: []byte("REDACTED"),
			Aunts:    [][]byte{},
		},
	}
	pbParts, err := parts.ToProto()
	require.NoError(t, err)

	proposal := kinds.Proposal{
		Type:      ctschema.ProposalType,
		Height:    1,
		Round:     1,
		POLRound:  1,
		BlockID:   bi,
		Timestamp: date,
		Signature: []byte("REDACTED"),
	}
	pbProposal := proposal.ToProto()

	v := &kinds.Vote{
		ValidatorAddress: []byte("REDACTED"),
		ValidatorIndex:   1,
		Height:           1,
		Round:            0,
		Timestamp:        date,
		Type:             ctschema.PrecommitType,
		BlockID:          bi,
	}
	vpb := v.ToProto()
	v.Extension = []byte("REDACTED")
	vextPb := v.ToProto()

	testCases := []struct {
		testName string
		cMsg     proto.Message
		expBytes string
	}{
		{"REDACTED", &enginecons.Message{Sum: &enginecons.Message_NewRoundStep{NewRoundStep: &enginecons.NewRoundStep{
			Height:                1,
			Round:                 1,
			Step:                  1,
			SecondsSinceStartTime: 1,
			LastCommitRound:       1,
		}}}, "REDACTED"},
		{"REDACTED", &enginecons.Message{Sum: &enginecons.Message_NewRoundStep{NewRoundStep: &enginecons.NewRoundStep{
			Height:                math.MaxInt64,
			Round:                 math.MaxInt32,
			Step:                  math.MaxUint32,
			SecondsSinceStartTime: math.MaxInt64,
			LastCommitRound:       math.MaxInt32,
		}}}, "REDACTED"},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_NewValidBlock{
				NewValidBlock: &enginecons.NewValidBlock{
					Height: 1, Round: 1, BlockPartSetHeader: pbPsh, BlockParts: pbBits, IsCommit: false,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_Proposal{Proposal: &enginecons.Proposal{Proposal: *pbProposal}}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_ProposalPol{
				ProposalPol: &enginecons.ProposalPOL{Height: 1, ProposalPolRound: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_BlockPart{
				BlockPart: &enginecons.BlockPart{Height: 1, Round: 1, Part: *pbParts},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_Vote{
				Vote: &enginecons.Vote{Vote: vpb},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_Vote{
				Vote: &enginecons.Vote{Vote: vextPb},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_HasVote{
				HasVote: &enginecons.HasVote{Height: 1, Round: 1, Type: ctschema.PrevoteType, Index: 1},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_HasVote{
				HasVote: &enginecons.HasVote{
					Height: math.MaxInt64, Round: math.MaxInt32,
					Type: ctschema.PrevoteType, Index: math.MaxInt32,
				},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_VoteSetMaj23{
				VoteSetMaj23: &enginecons.VoteSetMaj23{Height: 1, Round: 1, Type: ctschema.PrevoteType, BlockID: pbBi},
			}},
			"REDACTED",
		},
		{
			"REDACTED", &enginecons.Message{Sum: &enginecons.Message_VoteSetBits{
				VoteSetBits: &enginecons.VoteSetBits{Height: 1, Round: 1, Type: ctschema.PrevoteType, BlockID: pbBi, Votes: *pbBits},
			}},
			"REDACTED",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			bz, err := proto.Marshal(tc.cMsg)
			require.NoError(t, err)

			require.Equal(t, tc.expBytes, hex.EncodeToString(bz))
		})
	}
}
