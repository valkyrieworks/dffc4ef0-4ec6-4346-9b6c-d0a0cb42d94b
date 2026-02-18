package kv

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"

	"github.com/google/orderedcode"

	ordinalutil "github.com/valkyrieworks/intrinsic/ordinaler"
	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/kinds"
)

type LevelDetails struct {
	levelScope     ordinaler.InquireSpan
	level          int64
	levelEqualIndex     int
	solelyLevelScope bool
	solelyLevelEqual    bool
}

func integerInSection(a int, catalog []int) bool {
	for _, b := range catalog {
		if b == a {
			return true
		}
	}

	return false
}

func int64fromOctets(bz []byte) int64 {
	v, _ := binary.Varint(bz)
	return v
}

func int64toOctets(i int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, i)
	return buf[:n]
}

func levelKey(level int64) ([]byte, error) {
	return orderedcode.Append(
		nil,
		kinds.LedgerLevelKey,
		level,
	)
}

func eventKey(compoundKey, eventItem string, level int64, eventSeq int64) ([]byte, error) {
	return orderedcode.Append(
		nil,
		compoundKey,
		eventItem,
		level,
		eventSeq,
	)
}

func analyzeItemFromLeadingKey(key []byte) (string, error) {
	var (
		compoundKey string
		level       int64
	)

	pending, err := orderedcode.Parse(string(key), &compoundKey, &level)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	if len(pending) != 0 {
		return "REDACTED", fmt.Errorf("REDACTED", pending)
	}

	return strconv.FormatInt(level, 10), nil
}

func analyzeItemFromEventKey(key []byte) (string, error) {
	var (
		compoundKey, eventItem string
		level                   int64
	)

	_, err := orderedcode.Parse(string(key), &compoundKey, &eventItem, &level)
	if err != nil {
		return "REDACTED", fmt.Errorf("REDACTED", err)
	}

	return eventItem, nil
}

func analyzeLevelFromEventKey(key []byte) (int64, error) {
	var (
		compoundKey, eventItem string
		level                   int64
	)

	_, err := orderedcode.Parse(string(key), &compoundKey, &eventItem, &level)
	if err != nil {
		return -1, fmt.Errorf("REDACTED", err)
	}

	return level, nil
}

func analyzeEventSeqFromEventKey(key []byte) (int64, error) {
	var (
		compoundKey, eventItem string
		level                   int64
		eventSeq                 int64
	)

	pending, err := orderedcode.Parse(string(key), &compoundKey, &eventItem, &level)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	//
	//
	//
	//
	//
	//

	if len(pending) == 0 { //
		return 0, fmt.Errorf("REDACTED")
	}
	var typ string
	pending2, err := orderedcode.Parse(pending, &typ) //
	if err != nil {                                       //
		pending, err2 := orderedcode.Parse(string(key), &compoundKey, &eventItem, &level, &eventSeq)
		if err2 != nil || len(pending) != 0 { //
			return 0, fmt.Errorf("REDACTED", err, err2)
		}
	} else if len(pending2) != 0 { //
		pending, err2 := orderedcode.Parse(pending2, &eventSeq) //
		//
		//
		if err2 != nil || len(pending) != 0 { //
			return 0, fmt.Errorf("REDACTED", err2)
		}
	}
	return eventSeq, nil
}

//
//
//
//
func deduplicateLevel(states []grammar.State) (deduplicateStates []grammar.State, levelDetails LevelDetails, located bool) {
	levelDetails.levelEqualIndex = -1
	levelScopePresent := false
	var levelState []grammar.State
	levelDetails.solelyLevelEqual = true
	levelDetails.solelyLevelScope = true
	for _, c := range states {
		if c.Tag == kinds.LedgerLevelKey {
			if c.Op == grammar.TEq {
				if located || levelScopePresent {
					continue
				}
				hFloat := c.Arg.Amount()
				if hFloat != nil {
					h, _ := hFloat.Int64()
					levelDetails.level = h
					levelState = append(levelState, c)
					located = true
				}
			} else {
				levelDetails.solelyLevelEqual = false
				levelScopePresent = true
				deduplicateStates = append(deduplicateStates, c)
			}
		} else {
			levelDetails.solelyLevelScope = false
			levelDetails.solelyLevelEqual = false
			deduplicateStates = append(deduplicateStates, c)
		}
	}
	if !levelScopePresent && len(levelState) != 0 {
		levelDetails.levelEqualIndex = len(deduplicateStates)
		levelDetails.solelyLevelScope = false
		deduplicateStates = append(deduplicateStates, levelState...)
	} else {
		//
		//
		levelDetails.levelEqualIndex = -1
		levelDetails.level = 0
		levelDetails.solelyLevelEqual = false
		located = false
	}
	return deduplicateStates, levelDetails, located
}

func inspectLevelStates(levelDetails LevelDetails, keyLevel int64) (bool, error) {
	if levelDetails.levelScope.Key != "REDACTED" {
		insideLimits, err := ordinalutil.InspectLimits(levelDetails.levelScope, big.NewInt(keyLevel))
		if err != nil || !insideLimits {
			return false, err
		}
	} else if levelDetails.level != 0 && keyLevel != levelDetails.level {
		return false, nil
	}

	return true, nil
}
