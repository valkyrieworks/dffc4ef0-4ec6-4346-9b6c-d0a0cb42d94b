package agreement

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	"github.com/valkyrieworks/p2p"
)

var (
	_ p2p.Adapter = &BallotAssignBits{}
	_ p2p.Adapter = &BallotAssignMaj23{}
	_ p2p.Adapter = &Ballot{}
	_ p2p.Adapter = &NominationPOL{}
	_ p2p.Adapter = &Nomination{}
	_ p2p.Adapter = &NewSoundLedger{}
	_ p2p.Adapter = &NewDurationPhase{}
	_ p2p.Adapter = &HasBallot{}
	_ p2p.Adapter = &LedgerSegment{}
)

func (m *BallotAssignBits) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballotsetbits{BallotAssignBits: m}
	return cm
}

func (m *BallotAssignMaj23) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballotsetmaj23{BallotAssignMaj23: m}
	return cm
}

func (m *HasBallot) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Hasballot{HasBallot: m}
	return cm
}

func (m *Ballot) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ballot{Ballot: m}
	return cm
}

func (m *LedgerSegment) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Ledgersection{LedgerSegment: m}
	return cm
}

func (m *NominationPOL) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Nominationpol{NominationPol: m}
	return cm
}

func (m *Nomination) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Nomination{Nomination: m}
	return cm
}

func (m *NewSoundLedger) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Newvalidledger{NewSoundLedger: m}
	return cm
}

func (m *NewDurationPhase) Enclose() proto.Message {
	cm := &Signal{}
	cm.Sum = &Signal_Newepochphase{NewDurationPhase: m}
	return cm
}

//
//
func (m *Signal) Disclose() (proto.Message, error) {
	switch msg := m.Sum.(type) {
	case *Signal_Newepochphase:
		return m.FetchNewEpochPhase(), nil

	case *Signal_Newvalidledger:
		return m.FetchNewSoundLedger(), nil

	case *Signal_Nomination:
		return m.FetchNomination(), nil

	case *Signal_Nominationpol:
		return m.FetchNominationPol(), nil

	case *Signal_Ledgersection:
		return m.FetchLedgerSection(), nil

	case *Signal_Ballot:
		return m.FetchBallot(), nil

	case *Signal_Hasballot:
		return m.FetchHasBallot(), nil

	case *Signal_Ballotsetmaj23:
		return m.FetchBallotCollectionMaj23(), nil

	case *Signal_Ballotsetbits:
		return m.FetchBallotCollectionBits(), nil

	default:
		return nil, fmt.Errorf("REDACTED", msg)
	}
}
