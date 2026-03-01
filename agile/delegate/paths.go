package delegate

import (
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	airpc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/agile/rpc"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	rpchandler "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func RemotePaths(c *airpc.Customer) map[string]*rpchandler.RemoteMethod {
	return map[string]*rpchandler.RemoteMethod{
		//
		"REDACTED":       rpchandler.FreshSocketifaceMethod(c.ListenSocket, "REDACTED"),
		"REDACTED":     rpchandler.FreshSocketifaceMethod(c.UnlistenSocket, "REDACTED"),
		"REDACTED": rpchandler.FreshSocketifaceMethod(c.UnlistenEverySocket, "REDACTED"),

		//
		"REDACTED":               rpchandler.FreshRemoteMethod(createVitalityMethod(c), "REDACTED"),
		"REDACTED":               rpchandler.FreshRemoteMethod(createConditionMethod(c), "REDACTED"),
		"REDACTED":             rpchandler.FreshRemoteMethod(createNetworkDetailsMethod(c), "REDACTED"),
		"REDACTED":           rpchandler.FreshRemoteMethod(createLedgerchainDetailsMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":              rpchandler.FreshRemoteMethod(createInaugurationMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":      rpchandler.FreshRemoteMethod(createInaugurationSegmentedMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":                rpchandler.FreshRemoteMethod(createLedgerMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED":               rpchandler.FreshRemoteMethod(createHeadlineMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED":       rpchandler.FreshRemoteMethod(createHeadlineViaDigestMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":        rpchandler.FreshRemoteMethod(createLedgerViaDigestMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":        rpchandler.FreshRemoteMethod(createLedgerOutcomesMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED":               rpchandler.FreshRemoteMethod(createEndorseMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED":                   rpchandler.FreshRemoteMethod(createTransferMethod(c), "REDACTED", rpchandler.Storable()),
		"REDACTED":            rpchandler.FreshRemoteMethod(createTransferLookupMethod(c), "REDACTED"),
		"REDACTED":         rpchandler.FreshRemoteMethod(createLedgerLookupMethod(c), "REDACTED"),
		"REDACTED":           rpchandler.FreshRemoteMethod(createAssessorsMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED": rpchandler.FreshRemoteMethod(createExportAgreementStatusMethod(c), "REDACTED"),
		"REDACTED":      rpchandler.FreshRemoteMethod(createAgreementStatusMethod(c), "REDACTED"),
		"REDACTED":     rpchandler.FreshRemoteMethod(createAgreementParametersMethod(c), "REDACTED", rpchandler.Storable("REDACTED")),
		"REDACTED":      rpchandler.FreshRemoteMethod(createPendingTransMethod(c), "REDACTED"),
		"REDACTED":  rpchandler.FreshRemoteMethod(createCountPendingTransMethod(c), "REDACTED"),

		//
		"REDACTED": rpchandler.FreshRemoteMethod(createMulticastTransferEndorseMethod(c), "REDACTED"),
		"REDACTED":   rpchandler.FreshRemoteMethod(createMulticastTransferChronizeMethod(c), "REDACTED"),
		"REDACTED":  rpchandler.FreshRemoteMethod(createMulticastTransferAsyncronousMethod(c), "REDACTED"),

		//
		"REDACTED": rpchandler.FreshRemoteMethod(createIfaceInquireMethod(c), "REDACTED"),
		"REDACTED":  rpchandler.FreshRemoteMethod(createIfaceDetailsMethod(c), "REDACTED", rpchandler.Storable()),

		//
		"REDACTED": rpchandler.FreshRemoteMethod(createMulticastProofMethod(c), "REDACTED"),
	}
}

type remoteVitalityMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeVitality, error)

func createVitalityMethod(c *airpc.Customer) remoteVitalityMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeVitality, error) {
		return c.Vitality(ctx.Env())
	}
}

type remoteConditionMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeCondition, error)

func createConditionMethod(c *airpc.Customer) remoteConditionMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeCondition, error) {
		return c.Condition(ctx.Env())
	}
}

type remoteNetworkDetailsMethod func(ctx *remoteifacetypes.Env, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeNetworkDetails, error)

func createNetworkDetailsMethod(c *airpc.Customer) remoteNetworkDetailsMethod {
	return func(ctx *remoteifacetypes.Env, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeNetworkDetails, error) {
		return c.NetworkDetails(ctx.Env())
	}
}

type remoteLedgerchainDetailsMethod func(ctx *remoteifacetypes.Env, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error)

func createLedgerchainDetailsMethod(c *airpc.Customer) remoteLedgerchainDetailsMethod {
	return func(ctx *remoteifacetypes.Env, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error) {
		return c.LedgerchainDetails(ctx.Env(), minimumAltitude, maximumAltitude)
	}
}

type remoteInaugurationMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeInauguration, error)

func createInaugurationMethod(c *airpc.Customer) remoteInaugurationMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeInauguration, error) {
		return c.Inauguration(ctx.Env())
	}
}

type remoteInaugurationSegmentedMethod func(ctx *remoteifacetypes.Env, segment uint) (*ktypes.OutcomeInaugurationSegment, error)

func createInaugurationSegmentedMethod(c *airpc.Customer) remoteInaugurationSegmentedMethod {
	return func(ctx *remoteifacetypes.Env, segment uint) (*ktypes.OutcomeInaugurationSegment, error) {
		return c.InaugurationSegmented(ctx.Env(), segment)
	}
}

type remoteLedgerMethod func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeLedger, error)

func createLedgerMethod(c *airpc.Customer) remoteLedgerMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeLedger, error) {
		return c.Ledger(ctx.Env(), altitude)
	}
}

type remoteHeadlineMethod func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeHeadline, error)

func createHeadlineMethod(c *airpc.Customer) remoteHeadlineMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeHeadline, error) {
		return c.Heading(ctx.Env(), altitude)
	}
}

type remoteHeadlineViaDigestMethod func(ctx *remoteifacetypes.Env, digest []byte) (*ktypes.OutcomeHeadline, error)

func createHeadlineViaDigestMethod(c *airpc.Customer) remoteHeadlineViaDigestMethod {
	return func(ctx *remoteifacetypes.Env, digest []byte) (*ktypes.OutcomeHeadline, error) {
		return c.HeadingViaDigest(ctx.Env(), digest)
	}
}

type remoteLedgerViaDigestMethod func(ctx *remoteifacetypes.Env, digest []byte) (*ktypes.OutcomeLedger, error)

func createLedgerViaDigestMethod(c *airpc.Customer) remoteLedgerViaDigestMethod {
	return func(ctx *remoteifacetypes.Env, digest []byte) (*ktypes.OutcomeLedger, error) {
		return c.LedgerViaDigest(ctx.Env(), digest)
	}
}

type remoteLedgerOutcomesMethod func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeLedgerOutcomes, error)

func createLedgerOutcomesMethod(c *airpc.Customer) remoteLedgerOutcomesMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeLedgerOutcomes, error) {
		return c.LedgerOutcomes(ctx.Env(), altitude)
	}
}

type remoteEndorseMethod func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeEndorse, error)

func createEndorseMethod(c *airpc.Customer) remoteEndorseMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeEndorse, error) {
		return c.Endorse(ctx.Env(), altitude)
	}
}

type remoteTransferMethod func(ctx *remoteifacetypes.Env, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error)

func createTransferMethod(c *airpc.Customer) remoteTransferMethod {
	return func(ctx *remoteifacetypes.Env, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error) {
		return c.Tx(ctx.Env(), digest, ascertain)
	}
}

type remoteTransferLookupMethod func(
	ctx *remoteifacetypes.Env,
	inquire string,
	ascertain bool,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeTransferLookup, error)

func createTransferLookupMethod(c *airpc.Customer) remoteTransferLookupMethod {
	return func(
		ctx *remoteifacetypes.Env,
		inquire string,
		ascertain bool,
		screen, everyScreen *int,
		sequenceVia string,
	) (*ktypes.OutcomeTransferLookup, error) {
		return c.TransferLookup(ctx.Env(), inquire, ascertain, screen, everyScreen, sequenceVia)
	}
}

type remoteLedgerLookupMethod func(
	ctx *remoteifacetypes.Env,
	inquire string,
	ascertain bool,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeLedgerLookup, error)

func createLedgerLookupMethod(c *airpc.Customer) remoteLedgerLookupMethod {
	return func(
		ctx *remoteifacetypes.Env,
		inquire string,
		ascertain bool,
		screen, everyScreen *int,
		sequenceVia string,
	) (*ktypes.OutcomeLedgerLookup, error) {
		return c.LedgerLookup(ctx.Env(), inquire, screen, everyScreen, sequenceVia)
	}
}

type remoteAssessorsMethod func(ctx *remoteifacetypes.Env, altitude *int64,
	screen, everyScreen *int) (*ktypes.OutcomeAssessors, error)

func createAssessorsMethod(c *airpc.Customer) remoteAssessorsMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64, screen, everyScreen *int) (*ktypes.OutcomeAssessors, error) {
		return c.Assessors(ctx.Env(), altitude, screen, everyScreen)
	}
}

type remoteExportAgreementStatusMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeExportAgreementStatus, error)

func createExportAgreementStatusMethod(c *airpc.Customer) remoteExportAgreementStatusMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeExportAgreementStatus, error) {
		return c.ExportAgreementStatus(ctx.Env())
	}
}

type remoteAgreementStatusMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeAgreementStatus, error)

func createAgreementStatusMethod(c *airpc.Customer) remoteAgreementStatusMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeAgreementStatus, error) {
		return c.AgreementStatus(ctx.Env())
	}
}

type remoteAgreementParametersMethod func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeAgreementParameters, error)

func createAgreementParametersMethod(c *airpc.Customer) remoteAgreementParametersMethod {
	return func(ctx *remoteifacetypes.Env, altitude *int64) (*ktypes.OutcomeAgreementParameters, error) {
		return c.AgreementSettings(ctx.Env(), altitude)
	}
}

type remotePendingTransMethod func(ctx *remoteifacetypes.Env, threshold *int) (*ktypes.OutcomePendingTrans, error)

func createPendingTransMethod(c *airpc.Customer) remotePendingTransMethod {
	return func(ctx *remoteifacetypes.Env, threshold *int) (*ktypes.OutcomePendingTrans, error) {
		return c.PendingTrans(ctx.Env(), threshold)
	}
}

type remoteCountPendingTransMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomePendingTrans, error)

func createCountPendingTransMethod(c *airpc.Customer) remoteCountPendingTransMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomePendingTrans, error) {
		return c.CountPendingTrans(ctx.Env())
	}
}

type remoteMulticastTransferEndorseMethod func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error)

func createMulticastTransferEndorseMethod(c *airpc.Customer) remoteMulticastTransferEndorseMethod {
	return func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error) {
		return c.MulticastTransferEndorse(ctx.Env(), tx)
	}
}

type remoteMulticastTransferChronizeMethod func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error)

func createMulticastTransferChronizeMethod(c *airpc.Customer) remoteMulticastTransferChronizeMethod {
	return func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
		return c.MulticastTransferChronize(ctx.Env(), tx)
	}
}

type remoteMulticastTransferAsyncronousMethod func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error)

func createMulticastTransferAsyncronousMethod(c *airpc.Customer) remoteMulticastTransferAsyncronousMethod {
	return func(ctx *remoteifacetypes.Env, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
		return c.MulticastTransferAsyncronous(ctx.Env(), tx)
	}
}

type remoteIfaceInquireMethod func(ctx *remoteifacetypes.Env, route string,
	data octets.HexadecimalOctets, altitude int64, ascertain bool) (*ktypes.OutcomeIfaceInquire, error)

func createIfaceInquireMethod(c *airpc.Customer) remoteIfaceInquireMethod {
	return func(ctx *remoteifacetypes.Env, route string, data octets.HexadecimalOctets,
		altitude int64, ascertain bool,
	) (*ktypes.OutcomeIfaceInquire, error) {
		return c.IfaceInquireUsingChoices(ctx.Env(), route, data, customeriface.IfaceInquireChoices{
			Altitude: altitude,
			Validate:  ascertain,
		})
	}
}

type remoteIfaceDetailsMethod func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeIfaceDetails, error)

func createIfaceDetailsMethod(c *airpc.Customer) remoteIfaceDetailsMethod {
	return func(ctx *remoteifacetypes.Env) (*ktypes.OutcomeIfaceDetails, error) {
		return c.IfaceDetails(ctx.Env())
	}
}

type remoteMulticastProofMethod func(ctx *remoteifacetypes.Env, ev kinds.Proof) (*ktypes.OutcomeMulticastProof, error)

func createMulticastProofMethod(c *airpc.Customer) remoteMulticastProofMethod {
	return func(ctx *remoteifacetypes.Env, ev kinds.Proof) (*ktypes.OutcomeMulticastProof, error) {
		return c.MulticastProof(ctx.Env(), ev)
	}
}
