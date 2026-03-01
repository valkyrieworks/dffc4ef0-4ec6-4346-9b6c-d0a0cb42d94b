package base

import (
	"errors"
	"fmt"
	"sort"

	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/nothing"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
func (env *Context) Tx(_ *remoteifacetypes.Env, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error) {
	//
	if _, ok := env.TransferOrdinalizer.(*nothing.TransferOrdinal); ok {
		return nil, fmt.Errorf("REDACTED")
	}

	r, err := env.TransferOrdinalizer.Get(digest)
	if err != nil {
		return nil, err
	}

	if r == nil {
		return nil, fmt.Errorf("REDACTED", digest)
	}

	var attestation kinds.TransferAttestation
	if ascertain {
		ledger := env.LedgerDepot.FetchLedger(r.Altitude)
		if ledger != nil {
			attestation = ledger.Txs.Attestation(int(r.Ordinal))
		}
	}

	return &ktypes.OutcomeTransfer{
		Digest:     digest,
		Altitude:   r.Altitude,
		Ordinal:    r.Ordinal,
		TransferOutcome: r.Outcome,
		Tx:       r.Tx,
		Attestation:    attestation,
	}, nil
}

//
//
//
func (env *Context) TransferLookup(
	ctx *remoteifacetypes.Env,
	inquire string,
	ascertain bool,
	screenReference, everyScreenReference *int,
	sequenceVia string,
) (*ktypes.OutcomeTransferLookup, error) {
	//
	if _, ok := env.TransferOrdinalizer.(*nothing.TransferOrdinal); ok {
		return nil, errors.New("REDACTED")
	} else if len(inquire) > maximumInquireMagnitude {
		return nil, errors.New("REDACTED")
	}

	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.TransferOrdinalizer.Lookup(ctx.Env(), q)
	if err != nil {
		return nil, err
	}

	//
	switch sequenceVia {
	case "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool {
			if outcomes[i].Altitude == outcomes[j].Altitude {
				return outcomes[i].Ordinal > outcomes[j].Ordinal
			}
			return outcomes[i].Altitude > outcomes[j].Altitude
		})
	case "REDACTED", "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool {
			if outcomes[i].Altitude == outcomes[j].Altitude {
				return outcomes[i].Ordinal < outcomes[j].Ordinal
			}
			return outcomes[i].Altitude < outcomes[j].Altitude
		})
	default:
		return nil, errors.New("REDACTED")
	}

	//
	sumTally := len(outcomes)
	everyScreen := env.certifyEveryScreen(everyScreenReference)

	screen, err := certifyScreen(screenReference, everyScreen, sumTally)
	if err != nil {
		return nil, err
	}

	omitTally := certifyOmitTally(screen, everyScreen)
	screenExtent := strongarithmetic.MinimumInteger(everyScreen, sumTally-omitTally)

	apiOutcomes := make([]*ktypes.OutcomeTransfer, 0, screenExtent)
	for i := omitTally; i < omitTally+screenExtent; i++ {
		r := outcomes[i]

		var attestation kinds.TransferAttestation
		if ascertain {
			ledger := env.LedgerDepot.FetchLedger(r.Altitude)
			if ledger != nil {
				attestation = ledger.Txs.Attestation(int(r.Ordinal))
			}
		}

		apiOutcomes = append(apiOutcomes, &ktypes.OutcomeTransfer{
			Digest:     kinds.Tx(r.Tx).Digest(),
			Altitude:   r.Altitude,
			Ordinal:    r.Ordinal,
			TransferOutcome: r.Outcome,
			Tx:       r.Tx,
			Attestation:    attestation,
		})
	}

	return &ktypes.OutcomeTransferLookup{Txs: apiOutcomes, SumTally: sumTally}, nil
}
