package privatekey

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	privatekeyproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
)

//

func shouldEncloseMessage(pb proto.Message) privatekeyproto.Signal {
	msg := privatekeyproto.Signal{}

	switch pb := pb.(type) {
	case *privatekeyproto.Signal:
		msg = *pb
	case *privatekeyproto.PublicKeyQuery:
		msg.Sum = &privatekeyproto.Signal_Publickeyquery{PublicKeyQuery: pb}
	case *privatekeyproto.PublicKeyAnswer:
		msg.Sum = &privatekeyproto.Signal_Publickeyoutcome{PublicKeyAnswer: pb}
	case *privatekeyproto.AttestBallotQuery:
		msg.Sum = &privatekeyproto.Signal_Attestballotquery{AttestBallotQuery: pb}
	case *privatekeyproto.AttestedBallotAnswer:
		msg.Sum = &privatekeyproto.Signal_Attestedballotoutcome{AttestedBallotAnswer: pb}
	case *privatekeyproto.AttestedNominationAnswer:
		msg.Sum = &privatekeyproto.Signal_Attestedproposaloutcome{AttestedNominationAnswer: pb}
	case *privatekeyproto.AttestNominationQuery:
		msg.Sum = &privatekeyproto.Signal_Attestproposalquery{AttestNominationQuery: pb}
	case *privatekeyproto.PingQuery:
		msg.Sum = &privatekeyproto.Signal_Pingquery{PingQuery: pb}
	case *privatekeyproto.PingAnswer:
		msg.Sum = &privatekeyproto.Signal_Pingoutcome{PingAnswer: pb}
	default:
		panic(fmt.Errorf("REDACTED", pb))
	}

	return msg
}
