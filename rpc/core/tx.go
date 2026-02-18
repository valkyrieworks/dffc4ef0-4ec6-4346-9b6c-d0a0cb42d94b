package core

import (
	"errors"
	"fmt"
	"sort"

	cometmath "github.com/valkyrieworks/utils/math"
	cmtinquire "github.com/valkyrieworks/utils/broadcast/inquire"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/status/transordinal/void"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
func (env *Context) Tx(_ *rpctypes.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error) {
	//
	if _, ok := env.TransOrdinaler.(*void.TransOrdinal); ok {
		return nil, fmt.Errorf("REDACTED")
	}

	r, err := env.TransOrdinaler.Get(digest)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return nil, fmt.Errorf("REDACTED", digest)
	}

	var evidence kinds.TransferEvidence
	if demonstrate {
		ledger := env.LedgerDepot.ImportLedger(r.Level)
		if ledger != nil {
			evidence = ledger.Txs.Attestation(int(r.Ordinal))
		}
	}

	return &ctypes.OutcomeTransfer{
		Digest:     digest,
		Level:   r.Level,
		Ordinal:    r.Ordinal,
		TransOutcome: r.Outcome,
		Tx:       r.Tx,
		Attestation:    evidence,
	}, nil
}

//
//
//
func (env *Context) TransferScan(
	ctx *rpctypes.Context,
	inquire string,
	demonstrate bool,
	screenPointer, eachScreenPointer *int,
	arrangeBy string,
) (*ctypes.OutcomeTransferScan, error) {
	//
	if _, ok := env.TransOrdinaler.(*void.TransOrdinal); ok {
		return nil, errors.New("REDACTED")
	} else if len(inquire) > maximumInquireExtent {
		return nil, errors.New("REDACTED")
	}

	q, err := cmtinquire.New(inquire)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.TransOrdinaler.Scan(ctx.Context(), q)
	if err != nil {
		return nil, err
	}

	//
	switch arrangeBy {
	case "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool {
			if outcomes[i].Level == outcomes[j].Level {
				return outcomes[i].Ordinal > outcomes[j].Ordinal
			}
			return outcomes[i].Level > outcomes[j].Level
		})
	case "REDACTED", "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool {
			if outcomes[i].Level == outcomes[j].Level {
				return outcomes[i].Ordinal < outcomes[j].Ordinal
			}
			return outcomes[i].Level < outcomes[j].Level
		})
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

	apiOutcomes := make([]*ctypes.OutcomeTransfer, 0, screenVolume)
	for i := omitNumber; i < omitNumber+screenVolume; i++ {
		r := outcomes[i]

		var evidence kinds.TransferEvidence
		if demonstrate {
			ledger := env.LedgerDepot.ImportLedger(r.Level)
			if ledger != nil {
				evidence = ledger.Txs.Attestation(int(r.Ordinal))
			}
		}

		apiOutcomes = append(apiOutcomes, &ctypes.OutcomeTransfer{
			Digest:     kinds.Tx(r.Tx).Digest(),
			Level:   r.Level,
			Ordinal:    r.Ordinal,
			TransOutcome: r.Outcome,
			Tx:       r.Tx,
			Attestation:    evidence,
		})
	}

	return &ctypes.OutcomeTransferScan{Txs: apiOutcomes, SumNumber: sumNumber}, nil
}
