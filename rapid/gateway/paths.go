package gateway

import (
	"github.com/valkyrieworks/utils/octets"
	lrpc "github.com/valkyrieworks/rapid/rpc"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpchost "github.com/valkyrieworks/rpc/jsonrpc/host"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/kinds"
)

func RPCPaths(c *lrpc.Customer) map[string]*rpchost.RPCFunction {
	return map[string]*rpchost.RPCFunction{
		//
		"REDACTED":       rpchost.NewWsrpcFunction(c.EnrolWS, "REDACTED"),
		"REDACTED":     rpchost.NewWsrpcFunction(c.DeenrollWS, "REDACTED"),
		"REDACTED": rpchost.NewWsrpcFunction(c.DeenrollAllWS, "REDACTED"),

		//
		"REDACTED":               rpchost.NewRPCFunction(createVitalityFunction(c), "REDACTED"),
		"REDACTED":               rpchost.NewRPCFunction(createStateFunction(c), "REDACTED"),
		"REDACTED":             rpchost.NewRPCFunction(createNetDetailsFunction(c), "REDACTED"),
		"REDACTED":           rpchost.NewRPCFunction(createLedgerchainDetailsFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":              rpchost.NewRPCFunction(createOriginFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":      rpchost.NewRPCFunction(createOriginSegmentedFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":                rpchost.NewRPCFunction(createLedgerFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED":               rpchost.NewRPCFunction(createHeadingFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED":       rpchost.NewRPCFunction(createHeadingByDigestFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":        rpchost.NewRPCFunction(createLedgerByDigestFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":        rpchost.NewRPCFunction(createLedgerOutcomesFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED":               rpchost.NewRPCFunction(createEndorseFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED":                   rpchost.NewRPCFunction(createTransferFunction(c), "REDACTED", rpchost.Storable()),
		"REDACTED":            rpchost.NewRPCFunction(createTransferScanFunction(c), "REDACTED"),
		"REDACTED":         rpchost.NewRPCFunction(createLedgerScanFunction(c), "REDACTED"),
		"REDACTED":           rpchost.NewRPCFunction(createRatifiersFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED": rpchost.NewRPCFunction(createExportAgreementStatusFunction(c), "REDACTED"),
		"REDACTED":      rpchost.NewRPCFunction(createAgreementStatusFunction(c), "REDACTED"),
		"REDACTED":     rpchost.NewRPCFunction(createAgreementOptionsFunction(c), "REDACTED", rpchost.Storable("REDACTED")),
		"REDACTED":      rpchost.NewRPCFunction(createUnattestedTransFunction(c), "REDACTED"),
		"REDACTED":  rpchost.NewRPCFunction(createCountUnattestedTransFunction(c), "REDACTED"),

		//
		"REDACTED": rpchost.NewRPCFunction(createMulticastTransferEndorseFunction(c), "REDACTED"),
		"REDACTED":   rpchost.NewRPCFunction(createMulticastTransferAlignFunction(c), "REDACTED"),
		"REDACTED":  rpchost.NewRPCFunction(createMulticastTransferAsyncFunction(c), "REDACTED"),

		//
		"REDACTED": rpchost.NewRPCFunction(createIfaceInquireFunction(c), "REDACTED"),
		"REDACTED":  rpchost.NewRPCFunction(createIfaceDetailsFunction(c), "REDACTED", rpchost.Storable()),

		//
		"REDACTED": rpchost.NewRPCFunction(createMulticastProofFunction(c), "REDACTED"),
	}
}

type rpcVitalityFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeVitality, error)

func createVitalityFunction(c *lrpc.Customer) rpcVitalityFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeVitality, error) {
		return c.Vitality(ctx.Context())
	}
}

type rpcStateFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeState, error)

func createStateFunction(c *lrpc.Customer) rpcStateFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeState, error) {
		return c.Status(ctx.Context())
	}
}

type rpcNetDetailsFunction func(ctx *rpctypes.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeNetDetails, error)

func createNetDetailsFunction(c *lrpc.Customer) rpcNetDetailsFunction {
	return func(ctx *rpctypes.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeNetDetails, error) {
		return c.NetDetails(ctx.Context())
	}
}

type rpcLedgerchainDetailsFunction func(ctx *rpctypes.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error)

func createLedgerchainDetailsFunction(c *lrpc.Customer) rpcLedgerchainDetailsFunction {
	return func(ctx *rpctypes.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error) {
		return c.LedgerchainDetails(ctx.Context(), minimumLevel, maximumLevel)
	}
}

type rpcOriginFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeOrigin, error)

func createOriginFunction(c *lrpc.Customer) rpcOriginFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeOrigin, error) {
		return c.Origin(ctx.Context())
	}
}

type rpcOriginSegmentedFunction func(ctx *rpctypes.Context, segment uint) (*ctypes.OutcomeOriginSegment, error)

func createOriginSegmentedFunction(c *lrpc.Customer) rpcOriginSegmentedFunction {
	return func(ctx *rpctypes.Context, segment uint) (*ctypes.OutcomeOriginSegment, error) {
		return c.OriginSegmented(ctx.Context(), segment)
	}
}

type rpcLedgerFunction func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeLedger, error)

func createLedgerFunction(c *lrpc.Customer) rpcLedgerFunction {
	return func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeLedger, error) {
		return c.Ledger(ctx.Context(), level)
	}
}

type rpcHeadingFunction func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeHeading, error)

func createHeadingFunction(c *lrpc.Customer) rpcHeadingFunction {
	return func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeHeading, error) {
		return c.Heading(ctx.Context(), level)
	}
}

type rpcHeadingByDigestFunction func(ctx *rpctypes.Context, digest []byte) (*ctypes.OutcomeHeading, error)

func createHeadingByDigestFunction(c *lrpc.Customer) rpcHeadingByDigestFunction {
	return func(ctx *rpctypes.Context, digest []byte) (*ctypes.OutcomeHeading, error) {
		return c.HeadingByDigest(ctx.Context(), digest)
	}
}

type rpcLedgerByDigestFunction func(ctx *rpctypes.Context, digest []byte) (*ctypes.OutcomeLedger, error)

func createLedgerByDigestFunction(c *lrpc.Customer) rpcLedgerByDigestFunction {
	return func(ctx *rpctypes.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
		return c.LedgerByDigest(ctx.Context(), digest)
	}
}

type rpcLedgerOutcomesFunction func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeLedgerOutcomes, error)

func createLedgerOutcomesFunction(c *lrpc.Customer) rpcLedgerOutcomesFunction {
	return func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeLedgerOutcomes, error) {
		return c.LedgerOutcomes(ctx.Context(), level)
	}
}

type rpcEndorseFunction func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeEndorse, error)

func createEndorseFunction(c *lrpc.Customer) rpcEndorseFunction {
	return func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeEndorse, error) {
		return c.Endorse(ctx.Context(), level)
	}
}

type rpcTransferFunction func(ctx *rpctypes.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error)

func createTransferFunction(c *lrpc.Customer) rpcTransferFunction {
	return func(ctx *rpctypes.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error) {
		return c.Tx(ctx.Context(), digest, demonstrate)
	}
}

type rpcTransferScanFunction func(
	ctx *rpctypes.Context,
	inquire string,
	demonstrate bool,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeTransferScan, error)

func createTransferScanFunction(c *lrpc.Customer) rpcTransferScanFunction {
	return func(
		ctx *rpctypes.Context,
		inquire string,
		demonstrate bool,
		screen, eachScreen *int,
		arrangeBy string,
	) (*ctypes.OutcomeTransferScan, error) {
		return c.TransferScan(ctx.Context(), inquire, demonstrate, screen, eachScreen, arrangeBy)
	}
}

type rpcLedgerScanFunction func(
	ctx *rpctypes.Context,
	inquire string,
	demonstrate bool,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeLedgerScan, error)

func createLedgerScanFunction(c *lrpc.Customer) rpcLedgerScanFunction {
	return func(
		ctx *rpctypes.Context,
		inquire string,
		demonstrate bool,
		screen, eachScreen *int,
		arrangeBy string,
	) (*ctypes.OutcomeLedgerScan, error) {
		return c.LedgerScan(ctx.Context(), inquire, screen, eachScreen, arrangeBy)
	}
}

type rpcRatifiersFunction func(ctx *rpctypes.Context, level *int64,
	screen, eachScreen *int) (*ctypes.OutcomeRatifiers, error)

func createRatifiersFunction(c *lrpc.Customer) rpcRatifiersFunction {
	return func(ctx *rpctypes.Context, level *int64, screen, eachScreen *int) (*ctypes.OutcomeRatifiers, error) {
		return c.Ratifiers(ctx.Context(), level, screen, eachScreen)
	}
}

type rpcExportAgreementStatusFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeExportAgreementStatus, error)

func createExportAgreementStatusFunction(c *lrpc.Customer) rpcExportAgreementStatusFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
		return c.ExportAgreementStatus(ctx.Context())
	}
}

type rpcAgreementStatusFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeAgreementStatus, error)

func createAgreementStatusFunction(c *lrpc.Customer) rpcAgreementStatusFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeAgreementStatus, error) {
		return c.AgreementStatus(ctx.Context())
	}
}

type rpcAgreementOptionsFunction func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error)

func createAgreementOptionsFunction(c *lrpc.Customer) rpcAgreementOptionsFunction {
	return func(ctx *rpctypes.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error) {
		return c.AgreementOptions(ctx.Context(), level)
	}
}

type rpcUnattestedTransFunction func(ctx *rpctypes.Context, ceiling *int) (*ctypes.OutcomeUnattestedTrans, error)

func createUnattestedTransFunction(c *lrpc.Customer) rpcUnattestedTransFunction {
	return func(ctx *rpctypes.Context, ceiling *int) (*ctypes.OutcomeUnattestedTrans, error) {
		return c.UnattestedTrans(ctx.Context(), ceiling)
	}
}

type rpcCountUnattestedTransFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeUnattestedTrans, error)

func createCountUnattestedTransFunction(c *lrpc.Customer) rpcCountUnattestedTransFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeUnattestedTrans, error) {
		return c.CountUnattestedTrans(ctx.Context())
	}
}

type rpcMulticastTransferEndorseFunction func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error)

func createMulticastTransferEndorseFunction(c *lrpc.Customer) rpcMulticastTransferEndorseFunction {
	return func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error) {
		return c.MulticastTransferEndorse(ctx.Context(), tx)
	}
}

type rpcMulticastTransferAlignFunction func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error)

func createMulticastTransferAlignFunction(c *lrpc.Customer) rpcMulticastTransferAlignFunction {
	return func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
		return c.MulticastTransferAlign(ctx.Context(), tx)
	}
}

type rpcMulticastTransferAsyncFunction func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error)

func createMulticastTransferAsyncFunction(c *lrpc.Customer) rpcMulticastTransferAsyncFunction {
	return func(ctx *rpctypes.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
		return c.MulticastTransferAsync(ctx.Context(), tx)
	}
}

type rpcIfaceInquireFunction func(ctx *rpctypes.Context, route string,
	data octets.HexOctets, level int64, demonstrate bool) (*ctypes.OutcomeIfaceInquire, error)

func createIfaceInquireFunction(c *lrpc.Customer) rpcIfaceInquireFunction {
	return func(ctx *rpctypes.Context, route string, data octets.HexOctets,
		level int64, demonstrate bool,
	) (*ctypes.OutcomeIfaceInquire, error) {
		return c.IfaceInquireWithSettings(ctx.Context(), route, data, rpccustomer.IfaceInquireSettings{
			Level: level,
			Demonstrate:  demonstrate,
		})
	}
}

type rpcIfaceDetailsFunction func(ctx *rpctypes.Context) (*ctypes.OutcomeIfaceDetails, error)

func createIfaceDetailsFunction(c *lrpc.Customer) rpcIfaceDetailsFunction {
	return func(ctx *rpctypes.Context) (*ctypes.OutcomeIfaceDetails, error) {
		return c.IfaceDetails(ctx.Context())
	}
}

type rpcMulticastProofFunction func(ctx *rpctypes.Context, ev kinds.Proof) (*ctypes.OutcomeMulticastProof, error)

func createMulticastProofFunction(c *lrpc.Customer) rpcMulticastProofFunction {
	return func(ctx *rpctypes.Context, ev kinds.Proof) (*ctypes.OutcomeMulticastProof, error) {
		return c.MulticastProof(ctx.Context(), ev)
	}
}
