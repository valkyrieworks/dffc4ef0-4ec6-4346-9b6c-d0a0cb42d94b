package status

import (
	"github.com/valkyrieworks/kinds"
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
	Root() int64
	Level() int64
	Volume() int64

	ImportRootMeta() *kinds.LedgerMeta
	ImportLedgerMeta(level int64) *kinds.LedgerMeta
	ImportLedger(level int64) *kinds.Ledger

	PersistLedger(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedEndorse *kinds.Endorse)
	PersistLedgerWithExpandedEndorse(ledger *kinds.Ledger, ledgerSegments *kinds.SegmentCollection, viewedEndorse *kinds.ExpandedEndorse)

	TrimLedgers(level int64, status Status) (uint64, int64, error)

	ImportLedgerByDigest(digest []byte) *kinds.Ledger
	ImportLedgerMetaByDigest(digest []byte) *kinds.LedgerMeta
	ImportLedgerSegment(level int64, ordinal int) *kinds.Segment

	ImportLedgerEndorse(level int64) *kinds.Endorse
	ImportViewedEndorse(level int64) *kinds.Endorse
	ImportLedgerExpandedEndorse(level int64) *kinds.ExpandedEndorse

	RemoveNewestLedger() error

	End() error
}

//
//

//

//
type ProofDepository interface {
	AwaitingProof(maximumOctets int64) (ev []kinds.Proof, volume int64)
	AppendProof(kinds.Proof) error
	Modify(Status, kinds.ProofCatalog)
	InspectProof(kinds.ProofCatalog) error
}

//
//
type EmptyProofDepository struct{}

func (EmptyProofDepository) AwaitingProof(int64) (ev []kinds.Proof, volume int64) {
	return nil, 0
}
func (EmptyProofDepository) AppendProof(kinds.Proof) error                { return nil }
func (EmptyProofDepository) Modify(Status, kinds.ProofCatalog)                {}
func (EmptyProofDepository) InspectProof(kinds.ProofCatalog) error          { return nil }
func (EmptyProofDepository) NotifyClashingBallots(*kinds.Ballot, *kinds.Ballot) {}
