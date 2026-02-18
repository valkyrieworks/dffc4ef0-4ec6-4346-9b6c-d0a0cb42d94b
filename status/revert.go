package status

import (
	"errors"
	"fmt"

	cometstatus "github.com/valkyrieworks/schema/consensuscore/status"
	cometrelease "github.com/valkyrieworks/schema/consensuscore/release"
	"github.com/valkyrieworks/release"
)

//
//
//
func Revert(bs LedgerDepot, ss Depot, deleteLedger bool) (int64, []byte, error) {
	corruptStatus, err := ss.Import()
	if err != nil {
		return -1, nil, err
	}
	if corruptStatus.IsEmpty() {
		return -1, nil, errors.New("REDACTED")
	}

	level := bs.Level()

	//
	//
	//
	if level == corruptStatus.FinalLedgerLevel+1 {
		if deleteLedger {
			if err := bs.RemoveNewestLedger(); err != nil {
				return -1, nil, fmt.Errorf("REDACTED", err)
			}
		}
		return corruptStatus.FinalLedgerLevel, corruptStatus.ApplicationDigest, nil
	}

	//
	//
	if level != corruptStatus.FinalLedgerLevel {
		return -1, nil, fmt.Errorf("REDACTED",
			corruptStatus.FinalLedgerLevel, level)
	}

	//
	revertLevel := corruptStatus.FinalLedgerLevel - 1
	revertLedger := bs.ImportLedgerMeta(revertLevel)
	if revertLedger == nil {
		return -1, nil, fmt.Errorf("REDACTED", revertLevel)
	}
	//
	//
	newestLedger := bs.ImportLedgerMeta(corruptStatus.FinalLedgerLevel)
	if newestLedger == nil {
		return -1, nil, fmt.Errorf("REDACTED", corruptStatus.FinalLedgerLevel)
	}

	precedingFinalRatifierCollection, err := ss.ImportRatifiers(revertLevel)
	if err != nil {
		return -1, nil, err
	}

	precedingOptions, err := ss.ImportAgreementOptions(revertLevel + 1)
	if err != nil {
		return -1, nil, err
	}

	followingLevel := revertLevel + 1
	valueAlterLevel := corruptStatus.FinalLevelRatifiersModified
	//
	if valueAlterLevel > followingLevel+1 {
		valueAlterLevel = followingLevel + 1
	}

	optionsAlterLevel := corruptStatus.FinalLevelAgreementOptionsModified
	//
	if optionsAlterLevel > revertLevel {
		optionsAlterLevel = revertLevel + 1
	}

	//
	invertedRearStatus := Status{
		Release: cometstatus.Release{
			Agreement: cometrelease.Agreement{
				Ledger: release.LedgerProtocol,
				App:   precedingOptions.Release.App,
			},
			Software: release.TMCoreSemaphoreRev,
		},
		//
		LedgerUID:       corruptStatus.LedgerUID,
		PrimaryLevel: corruptStatus.PrimaryLevel,

		FinalLedgerLevel: revertLedger.Heading.Level,
		FinalLedgerUID:     revertLedger.LedgerUID,
		FinalLedgerTime:   revertLedger.Heading.Time,

		FollowingRatifiers:              corruptStatus.Ratifiers,
		Ratifiers:                  corruptStatus.FinalRatifiers,
		FinalRatifiers:              precedingFinalRatifierCollection,
		FinalLevelRatifiersModified: valueAlterLevel,

		AgreementOptions:                  precedingOptions,
		FinalLevelAgreementOptionsModified: optionsAlterLevel,

		FinalOutcomesDigest: newestLedger.Heading.FinalOutcomesDigest,
		ApplicationDigest:         newestLedger.Heading.ApplicationDigest,
	}

	//
	//
	//
	if err := ss.Persist(invertedRearStatus); err != nil {
		return -1, nil, fmt.Errorf("REDACTED", err)
	}

	//
	//
	if deleteLedger {
		if err := bs.RemoveNewestLedger(); err != nil {
			return -1, nil, fmt.Errorf("REDACTED", err)
		}
	}

	return invertedRearStatus.FinalLedgerLevel, invertedRearStatus.ApplicationDigest, nil
}
