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

	"github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/daemon"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	"github.com/valkyrieworks/kinds"
)

//
//
type Customer interface {
	daemon.Daemon
	IfaceCustomer
	EventsCustomer
	LogbookCustomer
	FabricCustomer
	AttestCustomer
	StateCustomer
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
	IfaceDetails(context.Context) (*ctypes.OutcomeIfaceDetails, error)
	IfaceInquire(ctx context.Context, route string, data octets.HexOctets) (*ctypes.OutcomeIfaceInquire, error)
	IfaceInquireWithSettings(ctx context.Context, route string, data octets.HexOctets,
		opts IfaceInquireSettings) (*ctypes.OutcomeIfaceInquire, error)

	//
	MulticastTransferEndorse(context.Context, kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error)
	MulticastTransferAsync(context.Context, kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error)
	MulticastTransferAlign(context.Context, kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error)
}

//
//
type AttestCustomer interface {
	Ledger(ctx context.Context, level *int64) (*ctypes.OutcomeLedger, error)
	LedgerByDigest(ctx context.Context, digest []byte) (*ctypes.OutcomeLedger, error)
	LedgerOutcomes(ctx context.Context, level *int64) (*ctypes.OutcomeLedgerOutcomes, error)
	Heading(ctx context.Context, level *int64) (*ctypes.OutcomeHeading, error)
	HeadingByDigest(ctx context.Context, digest octets.HexOctets) (*ctypes.OutcomeHeading, error)
	Endorse(ctx context.Context, level *int64) (*ctypes.OutcomeEndorse, error)
	Ratifiers(ctx context.Context, level *int64, screen, eachScreen *int) (*ctypes.OutcomeRatifiers, error)
	Tx(ctx context.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error)

	//
	//
	TransferScan(
		ctx context.Context,
		inquire string,
		demonstrate bool,
		screen, eachScreen *int,
		arrangeBy string,
	) (*ctypes.OutcomeTransferScan, error)

	//
	//
	LedgerScan(
		ctx context.Context,
		inquire string,
		screen, eachScreen *int,
		arrangeBy string,
	) (*ctypes.OutcomeLedgerScan, error)
}

//
type LogbookCustomer interface {
	Origin(context.Context) (*ctypes.OutcomeOrigin, error)
	OriginSegmented(context.Context, uint) (*ctypes.OutcomeOriginSegment, error)
	LedgerchainDetails(ctx context.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error)
}

//
type StateCustomer interface {
	Status(context.Context) (*ctypes.OutcomeState, error)
}

//
//
type FabricCustomer interface {
	NetDetails(context.Context) (*ctypes.OutcomeNetDetails, error)
	ExportAgreementStatus(context.Context) (*ctypes.OutcomeExportAgreementStatus, error)
	AgreementStatus(context.Context) (*ctypes.OutcomeAgreementStatus, error)
	AgreementOptions(ctx context.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error)
	Vitality(context.Context) (*ctypes.OutcomeVitality, error)
}

//
//
type EventsCustomer interface {
	//
	//
	//
	//
	//
	//
	//
	Enrol(ctx context.Context, enrollee, inquire string, outVolume ...int) (out <-chan ctypes.OutcomeEvent, err error)
	//
	Deenroll(ctx context.Context, enrollee, inquire string) error
	//
	DeenrollAll(ctx context.Context, enrollee string) error
}

//
type TxpoolCustomer interface {
	UnattestedTrans(ctx context.Context, ceiling *int) (*ctypes.OutcomeUnattestedTrans, error)
	CountUnattestedTrans(context.Context) (*ctypes.OutcomeUnattestedTrans, error)
	InspectTransfer(context.Context, kinds.Tx) (*ctypes.OutcomeInspectTransfer, error)
}

//
//
type ProofCustomer interface {
	MulticastProof(context.Context, kinds.Proof) (*ctypes.OutcomeMulticastProof, error)
}

//
type ExternalCustomer interface {
	Customer

	//
	External() string
}
