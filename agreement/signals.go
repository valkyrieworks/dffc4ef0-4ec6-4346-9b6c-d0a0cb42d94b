package agreement

import (
	"fmt"

	engineerrors "github.com/valkyrieworks/kinds/faults"
	"github.com/cosmos/gogoproto/proto"

	statetypes "github.com/valkyrieworks/agreement/kinds"
	"github.com/valkyrieworks/utils/units"
	ctalgebra "github.com/valkyrieworks/utils/algebra"
	"github.com/valkyrieworks/p2p"
	enginecons "github.com/valkyrieworks/schema/consensuscore/agreement"
	ctschema "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
//
func MsgToProto(msg Message) (proto.Message, error) {
	if msg == nil {
		return nil, ErrNilMessage
	}
	var pb proto.Message

	switch msg := msg.(type) {
	case *NewRoundStepMessage:
		pb = &enginecons.NewRoundStep{
			Height:                msg.Height,
			Round:                 msg.Round,
			Step:                  uint32(msg.Step),
			SecondsSinceStartTime: msg.SecondsSinceStartTime,
			LastCommitRound:       msg.LastCommitRound,
		}

	case *NewValidBlockMessage:
		pbPartSetHeader := msg.BlockPartSetHeader.ToProto()
		pbBits := msg.BlockParts.ToProto()
		pb = &enginecons.NewValidBlock{
			Height:             msg.Height,
			Round:              msg.Round,
			BlockPartSetHeader: pbPartSetHeader,
			BlockParts:         pbBits,
			IsCommit:           msg.IsCommit,
		}

	case *ProposalMessage:
		pbP := msg.Proposal.ToProto()
		pb = &enginecons.Proposal{
			Proposal: *pbP,
		}

	case *ProposalPOLMessage:
		pbBits := msg.ProposalPOL.ToProto()
		pb = &enginecons.ProposalPOL{
			Height:           msg.Height,
			ProposalPolRound: msg.ProposalPOLRound,
			ProposalPol:      *pbBits,
		}

	case *BlockPartMessage:
		parts, err := msg.Part.ToProto()
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}
		pb = &enginecons.BlockPart{
			Height: msg.Height,
			Round:  msg.Round,
			Part:   *parts,
		}

	case *VoteMessage:
		vote := msg.Vote.ToProto()
		pb = &enginecons.Vote{
			Vote: vote,
		}

	case *HasVoteMessage:
		pb = &enginecons.HasVote{
			Height: msg.Height,
			Round:  msg.Round,
			Type:   msg.Type,
			Index:  msg.Index,
		}

	case *VoteSetMaj23Message:
		bi := msg.BlockID.ToProto()
		pb = &enginecons.VoteSetMaj23{
			Height:  msg.Height,
			Round:   msg.Round,
			Type:    msg.Type,
			BlockID: bi,
		}

	case *VoteSetBitsMessage:
		bi := msg.BlockID.ToProto()
		bits := msg.Votes.ToProto()

		vsb := &enginecons.VoteSetBits{
			Height:  msg.Height,
			Round:   msg.Round,
			Type:    msg.Type,
			BlockID: bi,
		}

		if bits != nil {
			vsb.Votes = *bits
		}

		pb = vsb

	default:
		return nil, ErrConsensusMessageNotRecognized{msg}
	}

	return pb, nil
}

//
func MsgFromProto(p proto.Message) (Message, error) {
	if p == nil {
		return nil, ErrNilMessage
	}
	var pb Message

	switch msg := p.(type) {
	case *enginecons.NewRoundStep:
		rs, err := ctalgebra.SafeConvertUint8(int64(msg.Step))
		//
		if err != nil {
			return nil, ErrDenyMessageOverflow{err}
		}
		pb = &NewRoundStepMessage{
			Height:                msg.Height,
			Round:                 msg.Round,
			Step:                  statetypes.RoundStepType(rs),
			SecondsSinceStartTime: msg.SecondsSinceStartTime,
			LastCommitRound:       msg.LastCommitRound,
		}
	case *enginecons.NewValidBlock:
		pbPartSetHeader, err := kinds.PartSetHeaderFromProto(&msg.BlockPartSetHeader)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}

		pbBits := new(units.BitArray)
		pbBits.FromProto(msg.BlockParts)

		pb = &NewValidBlockMessage{
			Height:             msg.Height,
			Round:              msg.Round,
			BlockPartSetHeader: *pbPartSetHeader,
			BlockParts:         pbBits,
			IsCommit:           msg.IsCommit,
		}
	case *enginecons.Proposal:
		pbP, err := kinds.ProposalFromProto(&msg.Proposal)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}

		pb = &ProposalMessage{
			Proposal: pbP,
		}
	case *enginecons.ProposalPOL:
		pbBits := new(units.BitArray)
		pbBits.FromProto(&msg.ProposalPol)
		pb = &ProposalPOLMessage{
			Height:           msg.Height,
			ProposalPOLRound: msg.ProposalPolRound,
			ProposalPOL:      pbBits,
		}
	case *enginecons.BlockPart:
		parts, err := kinds.PartFromProto(&msg.Part)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}
		pb = &BlockPartMessage{
			Height: msg.Height,
			Round:  msg.Round,
			Part:   parts,
		}
	case *enginecons.Vote:
		//
		//
		vote, err := kinds.VoteFromProto(msg.Vote)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}

		pb = &VoteMessage{
			Vote: vote,
		}
	case *enginecons.HasVote:
		pb = &HasVoteMessage{
			Height: msg.Height,
			Round:  msg.Round,
			Type:   msg.Type,
			Index:  msg.Index,
		}
	case *enginecons.VoteSetMaj23:
		bi, err := kinds.BlockIDFromProto(&msg.BlockID)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}
		pb = &VoteSetMaj23Message{
			Height:  msg.Height,
			Round:   msg.Round,
			Type:    msg.Type,
			BlockID: *bi,
		}
	case *enginecons.VoteSetBits:
		bi, err := kinds.BlockIDFromProto(&msg.BlockID)
		if err != nil {
			return nil, engineerrors.ErrMsgToProto{MessageName: "REDACTED", Err: err}
		}
		bits := new(units.BitArray)
		units.FromProto(&msg.Votes)

		pb = &VoteSetBitsMessage{
			Height:  msg.Height,
			Round:   msg.Round,
			Type:    msg.Type,
			BlockID: *bi,
			Votes:   bits,
		}
	default:
		return nil, ErrConsensusMessageNotRecognized{msg}
	}

	if err := pb.ValidateBasic(); err != nil {
		return nil, err
	}

	return pb, nil
}

//
func WALToProto(msg WALMessage) (*enginecons.WALMessage, error) {
	var pb enginecons.WALMessage

	switch msg := msg.(type) {
	case kinds.EventDataRoundState:
		pb = enginecons.WALMessage{
			Sum: &enginecons.WALMessage_EventDataRoundState{
				EventDataRoundState: &ctschema.EventDataRoundState{
					Height: msg.Height,
					Round:  msg.Round,
					Step:   msg.Step,
				},
			},
		}
	case msgInfo:
		consMsg, err := MsgToProto(msg.Msg)
		if err != nil {
			return nil, err
		}
		if w, ok := consMsg.(p2p.Wrapper); ok {
			consMsg = w.Wrap()
		}
		cm := consMsg.(*enginecons.Message)
		pb = enginecons.WALMessage{
			Sum: &enginecons.WALMessage_MsgInfo{
				MsgInfo: &enginecons.MsgInfo{
					Msg:    *cm,
					PeerID: string(msg.PeerID),
				},
			},
		}
	case timeoutInfo:
		pb = enginecons.WALMessage{
			Sum: &enginecons.WALMessage_TimeoutInfo{
				TimeoutInfo: &enginecons.TimeoutInfo{
					Duration: msg.Duration,
					Height:   msg.Height,
					Round:    msg.Round,
					Step:     uint32(msg.Step),
				},
			},
		}
	case EndHeightMessage:
		pb = enginecons.WALMessage{
			Sum: &enginecons.WALMessage_EndHeight{
				EndHeight: &enginecons.EndHeight{
					Height: msg.Height,
				},
			},
		}
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}

	return &pb, nil
}

//
func WALFromProto(msg *enginecons.WALMessage) (WALMessage, error) {
	if msg == nil {
		return nil, ErrNilMessage
	}
	var pb WALMessage

	switch msg := msg.Sum.(type) {
	case *enginecons.WALMessage_EventDataRoundState:
		pb = kinds.EventDataRoundState{
			Height: msg.EventDataRoundState.Height,
			Round:  msg.EventDataRoundState.Round,
			Step:   msg.EventDataRoundState.Step,
		}
	case *enginecons.WALMessage_MsgInfo:
		um, err := msg.MsgInfo.Msg.Unwrap()
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		walMsg, err := MsgFromProto(um)
		if err != nil {
			return nil, engineerrors.ErrMsgFromProto{MessageName: "REDACTED", Err: err}
		}
		pb = msgInfo{
			Msg:    walMsg,
			PeerID: p2p.ID(msg.MsgInfo.PeerID),
		}

	case *enginecons.WALMessage_TimeoutInfo:
		tis, err := ctalgebra.SafeConvertUint8(int64(msg.TimeoutInfo.Step))
		//
		if err != nil {
			return nil, ErrDenyMessageOverflow{err}
		}
		pb = timeoutInfo{
			Duration: msg.TimeoutInfo.Duration,
			Height:   msg.TimeoutInfo.Height,
			Round:    msg.TimeoutInfo.Round,
			Step:     statetypes.RoundStepType(tis),
		}
		return pb, nil
	case *enginecons.WALMessage_EndHeight:
		pb := EndHeightMessage{
			Height: msg.EndHeight.Height,
		}
		return pb, nil
	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
	return pb, nil
}
