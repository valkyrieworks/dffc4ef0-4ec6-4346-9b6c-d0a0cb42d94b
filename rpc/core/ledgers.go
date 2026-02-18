package core

import (
	"errors"
	"fmt"
	"sort"

	"github.com/valkyrieworks/utils/octets"
	cometmath "github.com/valkyrieworks/utils/math"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	blockordinalvoid "github.com/valkyrieworks/status/ordinaler/ledger/void"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
//
//
//
//
//
//
func (env *Context) LedgerchainDetails(
	_ *rpctypes.Context,
	minimumLevel, maximumLevel int64,
) (*ctypes.OutcomeLedgerchainDetails, error) {
	const ceiling int64 = 20
	var err error
	minimumLevel, maximumLevel, err = refineMinimumMaximum(
		env.LedgerDepot.Root(),
		env.LedgerDepot.Level(),
		minimumLevel,
		maximumLevel,
		ceiling)
	if err != nil {
		return nil, err
	}
	env.Tracer.Diagnose("REDACTED", "REDACTED", maximumLevel, "REDACTED", minimumLevel)

	ledgerMetadata := []*kinds.LedgerMeta{}
	for level := maximumLevel; level >= minimumLevel; level-- {
		ledgerMeta := env.LedgerDepot.ImportLedgerMeta(level)
		ledgerMetadata = append(ledgerMetadata, ledgerMeta)
	}

	return &ctypes.OutcomeLedgerchainDetails{
		FinalLevel: env.LedgerDepot.Level(),
		LedgerMetadata: ledgerMetadata,
	}, nil
}

//
//
//
func refineMinimumMaximum(root, level, min, max, ceiling int64) (int64, int64, error) {
	//
	if min < 0 || max < 0 {
		return min, max, fmt.Errorf("REDACTED")
	}

	//
	if min == 0 {
		min = 1
	}
	if max == 0 {
		max = level
	}

	//
	max = cometmath.MinimumInt64(level, max)

	//
	min = cometmath.MaximumInt64(root, min)

	//
	//
	min = cometmath.MaximumInt64(min, max-ceiling+1)

	if min > max {
		return min, max, fmt.Errorf("REDACTED", min, max)
	}
	return min, max, nil
}

//
//
//
func (env *Context) Heading(_ *rpctypes.Context, levelPointer *int64) (*ctypes.OutcomeHeading, error) {
	level, err := env.fetchLevel(env.LedgerDepot.Level(), levelPointer)
	if err != nil {
		return nil, err
	}

	ledgerMeta := env.LedgerDepot.ImportLedgerMeta(level)
	if ledgerMeta == nil {
		return &ctypes.OutcomeHeading{}, nil
	}

	return &ctypes.OutcomeHeading{Heading: &ledgerMeta.Heading}, nil
}

//
//
func (env *Context) HeadingByDigest(_ *rpctypes.Context, digest octets.HexOctets) (*ctypes.OutcomeHeading, error) {
	//
	//
	//

	ledgerMeta := env.LedgerDepot.ImportLedgerMetaByDigest(digest)
	if ledgerMeta == nil {
		return &ctypes.OutcomeHeading{}, nil
	}

	return &ctypes.OutcomeHeading{Heading: &ledgerMeta.Heading}, nil
}

//
//
//
func (env *Context) Ledger(_ *rpctypes.Context, levelPointer *int64) (*ctypes.OutcomeLedger, error) {
	level, err := env.fetchLevel(env.LedgerDepot.Level(), levelPointer)
	if err != nil {
		return nil, err
	}

	ledger := env.LedgerDepot.ImportLedger(level)
	ledgerMeta := env.LedgerDepot.ImportLedgerMeta(level)
	if ledgerMeta == nil {
		return &ctypes.OutcomeLedger{LedgerUID: kinds.LedgerUID{}, Ledger: ledger}, nil
	}
	return &ctypes.OutcomeLedger{LedgerUID: ledgerMeta.LedgerUID, Ledger: ledger}, nil
}

//
//
func (env *Context) LedgerByDigest(_ *rpctypes.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
	ledger := env.LedgerDepot.ImportLedgerByDigest(digest)
	if ledger == nil {
		return &ctypes.OutcomeLedger{LedgerUID: kinds.LedgerUID{}, Ledger: nil}, nil
	}
	//
	ledgerMeta := env.LedgerDepot.ImportLedgerMeta(ledger.Level)
	return &ctypes.OutcomeLedger{LedgerUID: ledgerMeta.LedgerUID, Ledger: ledger}, nil
}

//
//
//
func (env *Context) Endorse(_ *rpctypes.Context, levelPointer *int64) (*ctypes.OutcomeEndorse, error) {
	level, err := env.fetchLevel(env.LedgerDepot.Level(), levelPointer)
	if err != nil {
		return nil, err
	}

	ledgerMeta := env.LedgerDepot.ImportLedgerMeta(level)
	if ledgerMeta == nil {
		return nil, nil
	}
	heading := ledgerMeta.Heading

	//
	//
	if level == env.LedgerDepot.Level() {
		endorse := env.LedgerDepot.ImportViewedEndorse(level)
		return ctypes.NewOutcomeEndorse(&heading, endorse, false), nil
	}

	//
	endorse := env.LedgerDepot.ImportLedgerEndorse(level)
	return ctypes.NewOutcomeEndorse(&heading, endorse, true), nil
}

//
//
//
//
//
//
//
func (env *Context) LedgerOutcomes(_ *rpctypes.Context, levelPointer *int64) (*ctypes.OutcomeLedgerOutcomes, error) {
	level, err := env.fetchLevel(env.LedgerDepot.Level(), levelPointer)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.StatusDepot.ImportCompleteLedgerReply(level)
	if err != nil {
		env.Tracer.Fault("REDACTED", "REDACTED", err)
		return nil, err
	}

	return &ctypes.OutcomeLedgerOutcomes{
		Level:                level,
		TransOutcomes:            outcomes.TransOutcomes,
		CompleteLedgerEvents:   outcomes.Events,
		RatifierRefreshes:      outcomes.RatifierRefreshes,
		AgreementArgumentRefreshes: outcomes.AgreementArgumentRefreshes,
		ApplicationDigest:               outcomes.ApplicationDigest,
	}, nil
}

//
//
func (env *Context) LedgerScan(
	ctx *rpctypes.Context,
	inquire string,
	screenPointer, eachScreenPointer *int,
	arrangeBy string,
) (*ctypes.OutcomeLedgerScan, error) {
	//
	if _, ok := env.LedgerOrdinaler.(*blockordinalvoid.ImpedimentOrdinaler); ok {
		return nil, errors.New("REDACTED")
	}

	q, err := cmtinquire.New(inquire)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.LedgerOrdinaler.Scan(ctx.Context(), q)
	if err != nil {
		return nil, err
	}

	//
	switch arrangeBy {
	case "REDACTED", "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool { return outcomes[i] > outcomes[j] })

	case "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool { return outcomes[i] < outcomes[j] })

	default:
		return nil, errors.New("REDACTED")
	}

	//
	sumNumber := len(outcomes)
	eachScreen := env.certifyEachScreen(eachScreenPointer)

	screen, err := certifyScreen(screenPointer, eachScreen, sumNumber)
	if err != nil {
		return nil, err
	}

	omitNumber := certifyOmitNumber(screen, eachScreen)
	screenVolume := cometmath.MinimumInteger(eachScreen, sumNumber-omitNumber)

	apiOutcomes := make([]*ctypes.OutcomeLedger, 0, screenVolume)
	for i := omitNumber; i < omitNumber+screenVolume; i++ {
		ledger := env.LedgerDepot.ImportLedger(outcomes[i])
		if ledger != nil {
			ledgerMeta := env.LedgerDepot.ImportLedgerMeta(ledger.Level)
			if ledgerMeta != nil {
				apiOutcomes = append(apiOutcomes, &ctypes.OutcomeLedger{
					Ledger:   ledger,
					LedgerUID: ledgerMeta.LedgerUID,
				})
			}
		}
	}

	return &ctypes.OutcomeLedgerScan{Ledgers: apiOutcomes, SumNumber: sumNumber}, nil
}
