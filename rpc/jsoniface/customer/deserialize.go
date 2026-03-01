package customer

import (
	"encoding/json"
	"errors"
	"fmt"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

func decodeReplyOctets(
	replyOctets []byte,
	anticipatedUUID kinds.JsonifaceIntegerUUID,
	outcome any,
) (any, error) {
	//
	//
	reply := &kinds.RemoteReply{}
	if err := json.Unmarshal(replyOctets, reply); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	if reply.Failure != nil {
		return nil, reply.Failure
	}

	if err := certifyAlsoValidateUUID(reply, anticipatedUUID); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	if err := strongmindjson.Decode(reply.Outcome, outcome); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return outcome, nil
}

//
func decodeDistinctReply(replyOctets []byte) (kinds.RemoteReply, error) {
	var uniqueReply kinds.RemoteReply
	err := json.Unmarshal(replyOctets, &uniqueReply)
	return uniqueReply, err
}

func decodeVariousReplies(replyOctets []byte) ([]kinds.RemoteReply, error) {
	var replies []kinds.RemoteReply
	err := json.Unmarshal(replyOctets, &replies)
	return replies, err
}

func decodeReplyOctetsSeries(
	replyOctets []byte,
	anticipatedIDXDstore []kinds.JsonifaceIntegerUUID,
	outcomes []any,
) ([]any, error) {
	var replies []kinds.RemoteReply

	//
	replies, err := decodeVariousReplies(replyOctets)
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
		ids := make([]kinds.JsonifaceIntegerUUID, len(replies))
		var ok bool
		for i, reply := range replies {
			ids[i], ok = reply.ID.(kinds.JsonifaceIntegerUUID)
			if !ok {
				return nil, fmt.Errorf("REDACTED", reply.ID)
			}
		}
		if err := certifyReplyIDXDstore(ids, anticipatedIDXDstore); err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		for i := 0; i < len(replies); i++ {
			if err := strongmindjson.Decode(replies[i].Outcome, outcomes[i]); err != nil {
				return nil, fmt.Errorf("REDACTED", i, err)
			}
		}

		return outcomes, nil
	}
	//
	uniqueReply, err := decodeDistinctReply(replyOctets)
	if err != nil {
		//
		//
		return nil, fmt.Errorf("REDACTED", err)
	}
	uniqueOutcome := make([]any, 0)
	if uniqueReply.Failure != nil {
		uniqueOutcome = append(uniqueOutcome, uniqueReply.Failure)
	} else {
		uniqueOutcome = append(uniqueOutcome, uniqueReply.Outcome)
	}
	return uniqueOutcome, nil
}

func certifyReplyIDXDstore(ids, anticipatedIDXDstore []kinds.JsonifaceIntegerUUID) error {
	m := make(map[kinds.JsonifaceIntegerUUID]bool, len(anticipatedIDXDstore))
	for _, anticipatedUUID := range anticipatedIDXDstore {
		m[anticipatedUUID] = true
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
func certifyAlsoValidateUUID(res *kinds.RemoteReply, anticipatedUUID kinds.JsonifaceIntegerUUID) error {
	if err := certifyReplyUUID(res.ID); err != nil {
		return err
	}
	if anticipatedUUID != res.ID.(kinds.JsonifaceIntegerUUID) { //
		return fmt.Errorf("REDACTED", res.ID, anticipatedUUID)
	}
	return nil
}

func certifyReplyUUID(id any) error {
	if id == nil {
		return errors.New("REDACTED")
	}
	_, ok := id.(kinds.JsonifaceIntegerUUID)
	if !ok {
		return fmt.Errorf("REDACTED", id)
	}
	return nil
}
