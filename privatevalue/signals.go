package privatevalue

import (
	"fmt"

	"github.com/cosmos/gogoproto/proto"

	privatevalueschema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
)

//

func shouldEncloseSignal(pb proto.Message) privatevalueschema.Signal {
	msg := privatevalueschema.Signal{}

	switch pb := pb.(type) {
	case *privatevalueschema.Signal:
		msg = *pb
	case *privatevalueschema.PublicTokenSolicit:
		msg.Sum = &privatevalueschema.Artifact_Publictokensolicit{PublicTokenSolicit: pb}
	case *privatevalueschema.PublicTokenReply:
		msg.Sum = &privatevalueschema.Artifact_Publictokendata{PublicTokenReply: pb}
	case *privatevalueschema.AttestBallotSolicit:
		msg.Sum = &privatevalueschema.Artifact_Notateballotsolicit{AttestBallotSolicit: pb}
	case *privatevalueschema.NotatedBallotReply:
		msg.Sum = &privatevalueschema.Artifact_Notatedballotreply{NotatedBallotReply: pb}
	case *privatevalueschema.NotatedNominationReply:
		msg.Sum = &privatevalueschema.Artifact_Notatedproposalreply{NotatedNominationReply: pb}
	case *privatevalueschema.AttestNominationSolicit:
		msg.Sum = &privatevalueschema.Artifact_Notateproposalsolicit{AttestNominationSolicit: pb}
	case *privatevalueschema.PingSolicit:
		msg.Sum = &privatevalueschema.Artifact_Pingsolicit{PingSolicit: pb}
	case *privatevalueschema.PingReply:
		msg.Sum = &privatevalueschema.Artifact_Pingreply{PingReply: pb}
	default:
		panic(fmt.Errorf("REDACTED", pb))
	}

	return msg
}
