package agreement

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

var (
	_ p2p.Encapsulator = &BallotAssignDigits{}
	_ p2p.Encapsulator = &BallotAssignMajor23{}
	_ p2p.Encapsulator = &Ballot{}
	_ p2p.Encapsulator = &NominationPolicy{}
	_ p2p.Encapsulator = &Nomination{}
	_ p2p.Encapsulator = &FreshSoundLedger{}
	_ p2p.Encapsulator = &FreshIterationPhase{}
	_ p2p.Encapsulator = &OwnsBallot{}
	_ p2p.Encapsulator = &LedgerFragment{}
)

func (m *BallotAssignDigits) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballotsetdigits{BallotAssignDigits: m}
	return cm
}

func (m *BallotAssignMajor23) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballotsetmaj23{BallotAssignMajor23: m}
	return cm
}

func (m *OwnsBallot) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Hasballot{OwnsBallot: m}
	return cm
}

func (m *Ballot) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballot{Ballot: m}
	return cm
}

func (m *LedgerFragment) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ledgerfragment{LedgerFragment: m}
	return cm
}

func (m *NominationPolicy) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Proposalpolicy{NominationPolicy: m}
	return cm
}

func (m *Nomination) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Nomination{Nomination: m}
	return cm
}

func (m *FreshSoundLedger) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Newvalidledger{FreshSoundLedger: m}
	return cm
}

func (m *FreshIterationPhase) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Newcyclephase{FreshIterationPhase: m}
	return cm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Newcyclephase:
		return m.ObtainFreshIterationPhase(), nil

	case *Signal_Newvalidledger:
		return m.ObtainFreshSoundLedger(), nil

	case *Signal_Nomination:
		return m.ObtainNomination(), nil

	case *Signal_Proposalpolicy:
		return m.ObtainNominationPolicy(), nil

	case *Signal_Ledgerfragment:
		return m.ObtainLedgerFragment(), nil

	case *Signal_Ballot:
		return m.FetchBallot(), nil

	case *Signal_Hasballot:
		return m.ObtainOwnsBallot(), nil

	case *Signal_Ballotsetmaj23:
		return m.ObtainBallotAssignMajor23(), nil

	case *Signal_Ballotsetdigits:
		return m.ObtainBallotAssignDigits(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
