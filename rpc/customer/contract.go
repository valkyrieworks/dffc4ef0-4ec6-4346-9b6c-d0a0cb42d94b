package customer

/**
g
.

h
.

n
l
.

r
.

y
,
t
.
*/

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
type Customer interface {
	facility.Facility
	IfaceCustomer
	IncidentsCustomer
	ChronicleCustomer
	FabricCustomer
	AttestCustomer
	ConditionCustomer
	ProofCustomer
	TxpoolCustomer
}

//
//
//
//
//
type IfaceCustomer interface {
	//
	IfaceDetails(context.Context) (*ktypes.OutcomeIfaceDetails, error)
	IfaceInquire(ctx context.Context, route string, data octets.HexadecimalOctets) (*ktypes.OutcomeIfaceInquire, error)
	IfaceInquireUsingChoices(ctx context.Context, route string, data octets.HexadecimalOctets,
		choices IfaceInquireChoices) (*ktypes.OutcomeIfaceInquire, error)

	//
	MulticastTransferEndorse(context.Context, kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error)
	MulticastTransferAsyncronous(context.Context, kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error)
	MulticastTransferChronize(context.Context, kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error)
}

//
//
type AttestCustomer interface {
	Ledger(ctx context.Context, altitude *int64) (*ktypes.OutcomeLedger, error)
	LedgerViaDigest(ctx context.Context, digest []byte) (*ktypes.OutcomeLedger, error)
	LedgerOutcomes(ctx context.Context, altitude *int64) (*ktypes.OutcomeLedgerOutcomes, error)
	Heading(ctx context.Context, altitude *int64) (*ktypes.OutcomeHeadline, error)
	HeadingViaDigest(ctx context.Context, digest octets.HexadecimalOctets) (*ktypes.OutcomeHeadline, error)
	Endorse(ctx context.Context, altitude *int64) (*ktypes.OutcomeEndorse, error)
	Assessors(ctx context.Context, altitude *int64, screen, everyScreen *int) (*ktypes.OutcomeAssessors, error)
	Tx(ctx context.Context, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error)

	//
	//
	TransferLookup(
		ctx context.Context,
		inquire string,
		ascertain bool,
		screen, everyScreen *int,
		sequenceVia string,
	) (*ktypes.OutcomeTransferLookup, error)

	//
	//
	LedgerLookup(
		ctx context.Context,
		inquire string,
		screen, everyScreen *int,
		sequenceVia string,
	) (*ktypes.OutcomeLedgerLookup, error)
}

//
type ChronicleCustomer interface {
	Inauguration(context.Context) (*ktypes.OutcomeInauguration, error)
	InaugurationSegmented(context.Context, uint) (*ktypes.OutcomeInaugurationSegment, error)
	LedgerchainDetails(ctx context.Context, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error)
}

//
type ConditionCustomer interface {
	Condition(context.Context) (*ktypes.OutcomeCondition, error)
}

//
//
type FabricCustomer interface {
	NetworkDetails(context.Context) (*ktypes.OutcomeNetworkDetails, error)
	ExportAgreementStatus(context.Context) (*ktypes.OutcomeExportAgreementStatus, error)
	AgreementStatus(context.Context) (*ktypes.OutcomeAgreementStatus, error)
	AgreementSettings(ctx context.Context, altitude *int64) (*ktypes.OutcomeAgreementParameters, error)
	Vitality(context.Context) (*ktypes.OutcomeVitality, error)
}

//
//
type IncidentsCustomer interface {
	//
	//
	//
	//
	//
	//
	//
	Listen(ctx context.Context, listener, inquire string, outputVolume ...int) (out <-chan ktypes.OutcomeIncident, err error)
	//
	Unlisten(ctx context.Context, listener, inquire string) error
	//
	UnlistenEvery(ctx context.Context, listener string) error
}

//
type TxpoolCustomer interface {
	PendingTrans(ctx context.Context, threshold *int) (*ktypes.OutcomePendingTrans, error)
	CountPendingTrans(context.Context) (*ktypes.OutcomePendingTrans, error)
	InspectTransfer(context.Context, kinds.Tx) (*ktypes.OutcomeInspectTransfer, error)
}

//
//
type ProofCustomer interface {
	MulticastProof(context.Context, kinds.Proof) (*ktypes.OutcomeMulticastProof, error)
}

//
type DistantCustomer interface {
	Customer

	//
	Distant() string
}
