package core

import (
	rpc "github.com/valkyrieworks/rpc/jsonrpc/host"
)

//

type PathsIndex map[string]*rpc.RPCFunction

//
func (env *Context) FetchPaths() PathsIndex {
	return PathsIndex{
		//
		"REDACTED":       rpc.NewWsrpcFunction(env.Enrol, "REDACTED"),
		"REDACTED":     rpc.NewWsrpcFunction(env.Deenroll, "REDACTED"),
		"REDACTED": rpc.NewWsrpcFunction(env.DeenrollAll, "REDACTED"),

		//
		"REDACTED":               rpc.NewRPCFunction(env.Vitality, "REDACTED"),
		"REDACTED":               rpc.NewRPCFunction(env.Status, "REDACTED"),
		"REDACTED":             rpc.NewRPCFunction(env.NetDetails, "REDACTED"),
		"REDACTED":           rpc.NewRPCFunction(env.LedgerchainDetails, "REDACTED", rpc.Storable()),
		"REDACTED":              rpc.NewRPCFunction(env.Origin, "REDACTED", rpc.Storable()),
		"REDACTED":      rpc.NewRPCFunction(env.OriginSegmented, "REDACTED", rpc.Storable()),
		"REDACTED":                rpc.NewRPCFunction(env.Ledger, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":        rpc.NewRPCFunction(env.LedgerByDigest, "REDACTED", rpc.Storable()),
		"REDACTED":        rpc.NewRPCFunction(env.LedgerOutcomes, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":               rpc.NewRPCFunction(env.Endorse, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":               rpc.NewRPCFunction(env.Heading, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":       rpc.NewRPCFunction(env.HeadingByDigest, "REDACTED", rpc.Storable()),
		"REDACTED":             rpc.NewRPCFunction(env.InspectTransfer, "REDACTED"),
		"REDACTED":                   rpc.NewRPCFunction(env.Tx, "REDACTED", rpc.Storable()),
		"REDACTED":            rpc.NewRPCFunction(env.TransferScan, "REDACTED"),
		"REDACTED":         rpc.NewRPCFunction(env.LedgerScan, "REDACTED"),
		"REDACTED":           rpc.NewRPCFunction(env.Ratifiers, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED": rpc.NewRPCFunction(env.ExportAgreementStatus, "REDACTED"),
		"REDACTED":      rpc.NewRPCFunction(env.FetchAgreementStatus, "REDACTED"),
		"REDACTED":     rpc.NewRPCFunction(env.AgreementOptions, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":      rpc.NewRPCFunction(env.UnattestedTrans, "REDACTED"),
		"REDACTED":  rpc.NewRPCFunction(env.CountUnattestedTrans, "REDACTED"),

		//
		"REDACTED": rpc.NewRPCFunction(env.MulticastTransferEndorse, "REDACTED"),
		"REDACTED":   rpc.NewRPCFunction(env.MulticastTransferAlign, "REDACTED"),
		"REDACTED":  rpc.NewRPCFunction(env.MulticastTransferAsync, "REDACTED"),

		//
		"REDACTED": rpc.NewRPCFunction(env.IfaceInquire, "REDACTED"),
		"REDACTED":  rpc.NewRPCFunction(env.IfaceDetails, "REDACTED", rpc.Storable()),

		//
		"REDACTED": rpc.NewRPCFunction(env.MulticastProof, "REDACTED"),
	}
}

//
func (env *Context) AppendRiskyPaths(paths PathsIndex) {
	//
	paths["REDACTED"] = rpc.NewRPCFunction(env.RiskyCallOrigins, "REDACTED")
	paths["REDACTED"] = rpc.NewRPCFunction(env.RiskyCallNodes, "REDACTED")
	paths["REDACTED"] = rpc.NewRPCFunction(env.RiskyPurgeTxpool, "REDACTED")
}
