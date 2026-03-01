package base

import (
	"errors"
	"fmt"
	"sort"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	tendermintinquire "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	ledgerpositionnull "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/nothing"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
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
	_ *remoteifacetypes.Env,
	minimumAltitude, maximumAltitude int64,
) (*ktypes.OutcomeLedgerchainDetails, error) {
	const threshold int64 = 20
	var err error
	minimumAltitude, maximumAltitude, err = refineMinimumMaximum(
		env.LedgerDepot.Foundation(),
		env.LedgerDepot.Altitude(),
		minimumAltitude,
		maximumAltitude,
		threshold)
	if err != nil {
		return nil, err
	}
	env.Tracer.Diagnose("REDACTED", "REDACTED", maximumAltitude, "REDACTED", minimumAltitude)

	ledgerMetadata := []*kinds.LedgerSummary{}
	for altitude := maximumAltitude; altitude >= minimumAltitude; altitude-- {
		ledgerSummary := env.LedgerDepot.FetchLedgerSummary(altitude)
		ledgerMetadata = append(ledgerMetadata, ledgerSummary)
	}

	return &ktypes.OutcomeLedgerchainDetails{
		FinalAltitude: env.LedgerDepot.Altitude(),
		LedgerMetadata: ledgerMetadata,
	}, nil
}

//
//
//
func refineMinimumMaximum(foundation, altitude, min, max, threshold int64) (int64, int64, error) {
	//
	if min < 0 || max < 0 {
		return min, max, fmt.Errorf("REDACTED")
	}

	//
	if min == 0 {
		min = 1
	}
	if max == 0 {
		max = altitude
	}

	//
	max = strongarithmetic.MinimumInt64n(altitude, max)

	//
	min = strongarithmetic.MaximumInt64n(foundation, min)

	//
	//
	min = strongarithmetic.MaximumInt64n(min, max-threshold+1)

	if min > max {
		return min, max, fmt.Errorf("REDACTED", min, max)
	}
	return min, max, nil
}

//
//
//
func (env *Context) Heading(_ *remoteifacetypes.Env, altitudeReference *int64) (*ktypes.OutcomeHeadline, error) {
	altitude, err := env.obtainAltitude(env.LedgerDepot.Altitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	ledgerSummary := env.LedgerDepot.FetchLedgerSummary(altitude)
	if ledgerSummary == nil {
		return &ktypes.OutcomeHeadline{}, nil
	}

	return &ktypes.OutcomeHeadline{Heading: &ledgerSummary.Heading}, nil
}

//
//
func (env *Context) HeadingViaDigest(_ *remoteifacetypes.Env, digest octets.HexadecimalOctets) (*ktypes.OutcomeHeadline, error) {
	//
	//
	//

	ledgerSummary := env.LedgerDepot.FetchLedgerSummaryViaDigest(digest)
	if ledgerSummary == nil {
		return &ktypes.OutcomeHeadline{}, nil
	}

	return &ktypes.OutcomeHeadline{Heading: &ledgerSummary.Heading}, nil
}

//
//
//
func (env *Context) Ledger(_ *remoteifacetypes.Env, altitudeReference *int64) (*ktypes.OutcomeLedger, error) {
	altitude, err := env.obtainAltitude(env.LedgerDepot.Altitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	ledger := env.LedgerDepot.FetchLedger(altitude)
	ledgerSummary := env.LedgerDepot.FetchLedgerSummary(altitude)
	if ledgerSummary == nil {
		return &ktypes.OutcomeLedger{LedgerUUID: kinds.LedgerUUID{}, Ledger: ledger}, nil
	}
	return &ktypes.OutcomeLedger{LedgerUUID: ledgerSummary.LedgerUUID, Ledger: ledger}, nil
}

//
//
func (env *Context) LedgerViaDigest(_ *remoteifacetypes.Env, digest []byte) (*ktypes.OutcomeLedger, error) {
	ledger := env.LedgerDepot.FetchLedgerViaDigest(digest)
	if ledger == nil {
		return &ktypes.OutcomeLedger{LedgerUUID: kinds.LedgerUUID{}, Ledger: nil}, nil
	}
	//
	ledgerSummary := env.LedgerDepot.FetchLedgerSummary(ledger.Altitude)
	return &ktypes.OutcomeLedger{LedgerUUID: ledgerSummary.LedgerUUID, Ledger: ledger}, nil
}

//
//
//
func (env *Context) Endorse(_ *remoteifacetypes.Env, altitudeReference *int64) (*ktypes.OutcomeEndorse, error) {
	altitude, err := env.obtainAltitude(env.LedgerDepot.Altitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	ledgerSummary := env.LedgerDepot.FetchLedgerSummary(altitude)
	if ledgerSummary == nil {
		return nil, nil
	}
	heading := ledgerSummary.Heading

	//
	//
	if altitude == env.LedgerDepot.Altitude() {
		endorse := env.LedgerDepot.FetchObservedEndorse(altitude)
		return ktypes.FreshOutcomeEndorse(&heading, endorse, false), nil
	}

	//
	endorse := env.LedgerDepot.FetchLedgerEndorse(altitude)
	return ktypes.FreshOutcomeEndorse(&heading, endorse, true), nil
}

//
//
//
//
//
//
//
func (env *Context) LedgerOutcomes(_ *remoteifacetypes.Env, altitudeReference *int64) (*ktypes.OutcomeLedgerOutcomes, error) {
	altitude, err := env.obtainAltitude(env.LedgerDepot.Altitude(), altitudeReference)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.StatusDepot.FetchCulminateLedgerReply(altitude)
	if err != nil {
		env.Tracer.Failure("REDACTED", "REDACTED", err)
		return nil, err
	}

	return &ktypes.OutcomeLedgerOutcomes{
		Altitude:                altitude,
		TransOutcomes:            outcomes.TransferOutcomes,
		CulminateLedgerIncidents:   outcomes.Incidents,
		AssessorRevisions:      outcomes.AssessorRevisions,
		AgreementArgumentRevisions: outcomes.AgreementArgumentRevisions,
		PlatformDigest:               outcomes.PlatformDigest,
	}, nil
}

//
//
func (env *Context) LedgerLookup(
	ctx *remoteifacetypes.Env,
	inquire string,
	screenReference, everyScreenReference *int,
	sequenceVia string,
) (*ktypes.OutcomeLedgerLookup, error) {
	//
	if _, ok := env.LedgerOrdinalizer.(*ledgerpositionnull.PreventerOrdinalizer); ok {
		return nil, errors.New("REDACTED")
	}

	q, err := tendermintinquire.New(inquire)
	if err != nil {
		return nil, err
	}

	outcomes, err := env.LedgerOrdinalizer.Lookup(ctx.Env(), q)
	if err != nil {
		return nil, err
	}

	//
	switch sequenceVia {
	case "REDACTED", "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool { return outcomes[i] > outcomes[j] })

	case "REDACTED":
		sort.Slice(outcomes, func(i, j int) bool { return outcomes[i] < outcomes[j] })

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

	apiOutcomes := make([]*ktypes.OutcomeLedger, 0, screenExtent)
	for i := omitTally; i < omitTally+screenExtent; i++ {
		ledger := env.LedgerDepot.FetchLedger(outcomes[i])
		if ledger != nil {
			ledgerSummary := env.LedgerDepot.FetchLedgerSummary(ledger.Altitude)
			if ledgerSummary != nil {
				apiOutcomes = append(apiOutcomes, &ktypes.OutcomeLedger{
					Ledger:   ledger,
					LedgerUUID: ledgerSummary.LedgerUUID,
				})
			}
		}
	}

	return &ktypes.OutcomeLedgerLookup{Ledgers: apiOutcomes, SumTally: sumTally}, nil
}
