package base

import (
	rpc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
)

//

type PathsIndex map[string]*rpc.RemoteMethod

//
func (env *Context) ObtainPaths() PathsIndex {
	return PathsIndex{
		//
		"REDACTED":       rpc.FreshSocketifaceMethod(env.Listen, "REDACTED"),
		"REDACTED":     rpc.FreshSocketifaceMethod(env.Unlisten, "REDACTED"),
		"REDACTED": rpc.FreshSocketifaceMethod(env.UnlistenEvery, "REDACTED"),

		//
		"REDACTED":               rpc.FreshRemoteMethod(env.Vitality, "REDACTED"),
		"REDACTED":               rpc.FreshRemoteMethod(env.Condition, "REDACTED"),
		"REDACTED":             rpc.FreshRemoteMethod(env.NetworkDetails, "REDACTED"),
		"REDACTED":           rpc.FreshRemoteMethod(env.LedgerchainDetails, "REDACTED", rpc.Storable()),
		"REDACTED":              rpc.FreshRemoteMethod(env.Inauguration, "REDACTED", rpc.Storable()),
		"REDACTED":      rpc.FreshRemoteMethod(env.InaugurationSegmented, "REDACTED", rpc.Storable()),
		"REDACTED":                rpc.FreshRemoteMethod(env.Ledger, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":        rpc.FreshRemoteMethod(env.LedgerViaDigest, "REDACTED", rpc.Storable()),
		"REDACTED":        rpc.FreshRemoteMethod(env.LedgerOutcomes, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":               rpc.FreshRemoteMethod(env.Endorse, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":               rpc.FreshRemoteMethod(env.Heading, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":       rpc.FreshRemoteMethod(env.HeadingViaDigest, "REDACTED", rpc.Storable()),
		"REDACTED":             rpc.FreshRemoteMethod(env.InspectTransfer, "REDACTED"),
		"REDACTED":                   rpc.FreshRemoteMethod(env.Tx, "REDACTED", rpc.Storable()),
		"REDACTED":            rpc.FreshRemoteMethod(env.TransferLookup, "REDACTED"),
		"REDACTED":         rpc.FreshRemoteMethod(env.LedgerLookup, "REDACTED"),
		"REDACTED":           rpc.FreshRemoteMethod(env.Assessors, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED": rpc.FreshRemoteMethod(env.ExportAgreementStatus, "REDACTED"),
		"REDACTED":      rpc.FreshRemoteMethod(env.ObtainAgreementStatus, "REDACTED"),
		"REDACTED":     rpc.FreshRemoteMethod(env.AgreementSettings, "REDACTED", rpc.Storable("REDACTED")),
		"REDACTED":      rpc.FreshRemoteMethod(env.PendingTrans, "REDACTED"),
		"REDACTED":  rpc.FreshRemoteMethod(env.CountPendingTrans, "REDACTED"),

		//
		"REDACTED": rpc.FreshRemoteMethod(env.MulticastTransferEndorse, "REDACTED"),
		"REDACTED":   rpc.FreshRemoteMethod(env.MulticastTransferChronize, "REDACTED"),
		"REDACTED":  rpc.FreshRemoteMethod(env.MulticastTransferAsyncronous, "REDACTED"),

		//
		"REDACTED": rpc.FreshRemoteMethod(env.IfaceInquire, "REDACTED"),
		"REDACTED":  rpc.FreshRemoteMethod(env.IfaceDetails, "REDACTED", rpc.Storable()),

		//
		"REDACTED": rpc.FreshRemoteMethod(env.MulticastProof, "REDACTED"),
	}
}

//
func (env *Context) AppendInsecurePaths(paths PathsIndex) {
	//
	paths["REDACTED"] = rpc.FreshRemoteMethod(env.InsecureCallOrigins, "REDACTED")
	paths["REDACTED"] = rpc.FreshRemoteMethod(env.InsecureCallNodes, "REDACTED")
	paths["REDACTED"] = rpc.FreshRemoteMethod(env.InsecurePurgeTxpool, "REDACTED")
}
