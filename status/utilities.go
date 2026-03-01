package status

import (
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
type LedgerDepot interface {
	Foundation() int64
	Altitude() int64
	Extent() int64

	FetchFoundationSummary() *kinds.LedgerSummary
	FetchLedgerSummary(altitude int64) *kinds.LedgerSummary
	FetchLedger(altitude int64) *kinds.Ledger

	PersistLedger(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedEndorse *kinds.Endorse)
	PersistLedgerUsingExpandedEndorse(ledger *kinds.Ledger, ledgerFragments *kinds.FragmentAssign, observedEndorse *kinds.ExpandedEndorse)

	TrimLedgers(altitude int64, status Status) (uint64, int64, error)

	FetchLedgerViaDigest(digest []byte) *kinds.Ledger
	FetchLedgerSummaryViaDigest(digest []byte) *kinds.LedgerSummary
	FetchLedgerFragment(altitude int64, ordinal int) *kinds.Fragment

	FetchLedgerEndorse(altitude int64) *kinds.Endorse
	FetchObservedEndorse(altitude int64) *kinds.Endorse
	FetchLedgerExpandedEndorse(altitude int64) *kinds.ExpandedEndorse

	EraseNewestLedger() error

	Shutdown() error
}

//
//

//

//
type ProofHub interface {
	AwaitingProof(maximumOctets int64) (ev []kinds.Proof, extent int64)
	AppendProof(kinds.Proof) error
	Revise(Status, kinds.ProofCatalog)
	InspectProof(kinds.ProofCatalog) error
}

//
//
type VoidProofHub struct{}

func (VoidProofHub) AwaitingProof(int64) (ev []kinds.Proof, extent int64) {
	return nil, 0
}
func (VoidProofHub) AppendProof(kinds.Proof) error                { return nil }
func (VoidProofHub) Revise(Status, kinds.ProofCatalog)                {}
func (VoidProofHub) InspectProof(kinds.ProofCatalog) error          { return nil }
func (VoidProofHub) DiscloseDiscordantBallots(*kinds.Ballot, *kinds.Ballot) {}
