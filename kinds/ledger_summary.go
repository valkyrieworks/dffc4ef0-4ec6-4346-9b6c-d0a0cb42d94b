package kinds

import (
	"bytes"
	"errors"
	"fmt"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//
type LedgerSummary struct {
	LedgerUUID   LedgerUUID `json:"ledger_uuid"`
	LedgerExtent int     `json:"ledger_extent"`
	Heading    Heading  `json:"heading"`
	CountTrans    int     `json:"count_trans"`
}

//
func FreshLedgerSummary(ledger *Ledger, ledgerFragments *FragmentAssign) *LedgerSummary {
	return &LedgerSummary{
		LedgerUUID:   LedgerUUID{ledger.Digest(), ledgerFragments.Heading()},
		LedgerExtent: ledger.Extent(),
		Heading:    ledger.Heading,
		CountTrans:    len(ledger.Txs),
	}
}

func (bm *LedgerSummary) TowardSchema() *commitchema.LedgerSummary {
	if bm == nil {
		return nil
	}

	pb := &commitchema.LedgerSummary{
		LedgerUUID:   bm.LedgerUUID.TowardSchema(),
		LedgerExtent: int64(bm.LedgerExtent),
		Heading:    *bm.Heading.TowardSchema(),
		CountTrans:    int64(bm.CountTrans),
	}
	return pb
}

func LedgerSummaryOriginatingSchema(pb *commitchema.LedgerSummary) (*LedgerSummary, error) {
	bm, err := LedgerSummaryOriginatingReliableSchema(pb)
	if err != nil {
		return nil, err
	}
	return bm, bm.CertifyFundamental()
}

func LedgerSummaryOriginatingReliableSchema(pb *commitchema.LedgerSummary) (*LedgerSummary, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	bm := new(LedgerSummary)

	bi, err := LedgerUUIDOriginatingSchema(&pb.LedgerUUID)
	if err != nil {
		return nil, err
	}

	h, err := HeadingOriginatingSchema(&pb.Heading)
	if err != nil {
		return nil, err
	}

	bm.LedgerUUID = *bi
	bm.LedgerExtent = int(pb.LedgerExtent)
	bm.Heading = h
	bm.CountTrans = int(pb.CountTrans)

	return bm, nil
}

//
func (bm *LedgerSummary) CertifyFundamental() error {
	if err := bm.LedgerUUID.CertifyFundamental(); err != nil {
		return err
	}
	if !bytes.Equal(bm.LedgerUUID.Digest, bm.Heading.Digest()) {
		return fmt.Errorf("REDACTED",
			bm.LedgerUUID.Digest, bm.Heading.Digest())
	}
	return nil
}
