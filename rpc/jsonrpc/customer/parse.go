package customer

import (
	"encoding/json"
	"errors"
	"fmt"

	cometjson "github.com/valkyrieworks/utils/json"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

func unserializeReplyOctets(
	replyOctets []byte,
	anticipatedUID kinds.JsonrpcIntegerUID,
	outcome any,
) (any, error) {
	//
	//
	reply := &kinds.RPCAnswer{}
	if err := json.Unmarshal(replyOctets, reply); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if reply.Fault != nil {
		return nil, reply.Fault
	}

	if err := certifyAndValidateUID(reply, anticipatedUID); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	if err := cometjson.Unserialize(reply.Outcome, outcome); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return outcome, nil
}

//
func unserializeDistinctReply(replyOctets []byte) (kinds.RPCAnswer, error) {
	var uniqueReply kinds.RPCAnswer
	err := json.Unmarshal(replyOctets, &uniqueReply)
	return uniqueReply, err
}

func unserializeVariedReplies(replyOctets []byte) ([]kinds.RPCAnswer, error) {
	var replies []kinds.RPCAnswer
	err := json.Unmarshal(replyOctets, &replies)
	return replies, err
}

func unserializeReplyOctetsList(
	replyOctets []byte,
	anticipatedIDXDatastore []kinds.JsonrpcIntegerUID,
	outcomes []any,
) ([]any, error) {
	var replies []kinds.RPCAnswer

	//
	replies, err := unserializeVariedReplies(replyOctets)
	//
	if err == nil {
		//
		//

		if len(outcomes) != len(replies) {
			return nil, fmt.Errorf(
				"REDACTED",
				len(replies),
				len(outcomes),
			)
		}

		//
		ids := make([]kinds.JsonrpcIntegerUID, len(replies))
		var ok bool
		for i, reply := range replies {
			ids[i], ok = reply.ID.(kinds.JsonrpcIntegerUID)
			if !ok {
				return nil, fmt.Errorf("REDACTED", reply.ID)
			}
		}
		if err := certifyReplyIDXDatastore(ids, anticipatedIDXDatastore); err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		for i := 0; i < len(replies); i++ {
			if err := cometjson.Unserialize(replies[i].Outcome, outcomes[i]); err != nil {
				return nil, fmt.Errorf("REDACTED", i, err)
			}
		}

		return outcomes, nil
	}
	//
	uniqueReply, err := unserializeDistinctReply(replyOctets)
	if err != nil {
		//
		//
		return nil, fmt.Errorf("REDACTED", err)
	}
	uniqueOutcome := make([]any, 0)
	if uniqueReply.Fault != nil {
		uniqueOutcome = append(uniqueOutcome, uniqueReply.Fault)
	} else {
		uniqueOutcome = append(uniqueOutcome, uniqueReply.Outcome)
	}
	return uniqueOutcome, nil
}

func certifyReplyIDXDatastore(ids, anticipatedIDXDatastore []kinds.JsonrpcIntegerUID) error {
	m := make(map[kinds.JsonrpcIntegerUID]bool, len(anticipatedIDXDatastore))
	for _, anticipatedUID := range anticipatedIDXDatastore {
		m[anticipatedUID] = true
	}

	for i, id := range ids {
		if m[id] {
			delete(m, id)
		} else {
			return fmt.Errorf("REDACTED", i, id)
		}
	}

	return nil
}

//
//
func certifyAndValidateUID(res *kinds.RPCAnswer, anticipatedUID kinds.JsonrpcIntegerUID) error {
	if err := certifyReplyUID(res.ID); err != nil {
		return err
	}
	if anticipatedUID != res.ID.(kinds.JsonrpcIntegerUID) { //
		return fmt.Errorf("REDACTED", res.ID, anticipatedUID)
	}
	return nil
}

func certifyReplyUID(id any) error {
	if id == nil {
		return errors.New("REDACTED")
	}
	_, ok := id.(kinds.JsonrpcIntegerUID)
	if !ok {
		return fmt.Errorf("REDACTED", id)
	}
	return nil
}
