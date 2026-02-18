package kv

import (
	"fmt"
	"math/big"

	ordinalutil "github.com/valkyrieworks/intrinsic/ordinaler"
	cmtsyntax "github.com/valkyrieworks/utils/broadcast/inquire/grammar"
	"github.com/valkyrieworks/status/ordinaler"
	"github.com/valkyrieworks/kinds"
	"github.com/google/orderedcode"
)

type LevelDetails struct {
	levelSpan     ordinaler.InquireSpan
	level          int64
	levelEqualOrdinal     int
	solelyLevelSpan bool
	solelyLevelEqual    bool
}

//
func integerInSection(a int, catalog []int) bool {
	for _, b := range catalog {
		if b == a {
			return true
		}
	}
	return false
}

func AnalyzeEventSeqFromEventKey(key []byte) (int64, error) {
	var (
		compoundKey, typ, eventItem string
		level                        int64
		eventSeq                      int64
	)

	pending, err := orderedcode.Parse(string(key), &compoundKey, &eventItem, &level, &typ, &eventSeq)
	if err != nil {
		return 0, fmt.Errorf("REDACTED", err)
	}

	if len(pending) != 0 {
		return 0, fmt.Errorf("REDACTED", pending)
	}

	return eventSeq, nil
}

func deduplicateLevel(states []cmtsyntax.Status) (deduplicateStates []cmtsyntax.Status, levelDetails LevelDetails) {
	levelDetails.levelEqualOrdinal = -1
	levelSpanPresent := false
	located := false
	var levelStatus []cmtsyntax.Status
	levelDetails.solelyLevelEqual = true
	levelDetails.solelyLevelSpan = true
	for _, c := range states {
		if c.Tag == kinds.TransferLevelKey {
			if c.Op == cmtsyntax.TEq {
				if levelSpanPresent || located {
					continue
				}
				hFloat := c.Arg.Amount()
				if hFloat != nil {
					h, _ := hFloat.Int64()
					levelDetails.level = h
					located = true
					levelStatus = append(levelStatus, c)
				}
			} else {
				levelDetails.solelyLevelEqual = false
				levelSpanPresent = true
				deduplicateStates = append(deduplicateStates, c)
			}
		} else {
			levelDetails.solelyLevelSpan = false
			levelDetails.solelyLevelEqual = false
			deduplicateStates = append(deduplicateStates, c)
		}
	}
	if !levelSpanPresent && len(levelStatus) != 0 {
		levelDetails.levelEqualOrdinal = len(deduplicateStates)
		levelDetails.solelyLevelSpan = false
		deduplicateStates = append(deduplicateStates, levelStatus...)
	} else {
		//
		//
		levelDetails.levelEqualOrdinal = -1
		levelDetails.level = 0
		levelDetails.solelyLevelEqual = false
	}
	return deduplicateStates, levelDetails
}

func inspectLevelStates(levelDetails LevelDetails, keyLevel int64) (bool, error) {
	if levelDetails.levelSpan.Key != "REDACTED" {
		insideLimits, err := ordinalutil.InspectLimits(levelDetails.levelSpan, big.NewInt(keyLevel))
		if err != nil || !insideLimits {
			return false, err
		}
	} else if levelDetails.level != 0 && keyLevel != levelDetails.level {
		return false, nil
	}
	return true, nil
}
