package kinds

import (
	"bytes"
	"errors"
	"fmt"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//
type LedgerMeta struct {
	LedgerUID   LedgerUID `json:"ledger_uid"`
	LedgerVolume int     `json:"ledger_volume"`
	Heading    Heading  `json:"heading"`
	CountTrans    int     `json:"count_trans"`
}

//
func NewLedgerMeta(ledger *Ledger, ledgerSegments *SegmentCollection) *LedgerMeta {
	return &LedgerMeta{
		LedgerUID:   LedgerUID{ledger.Digest(), ledgerSegments.Heading()},
		LedgerVolume: ledger.Volume(),
		Heading:    ledger.Heading,
		CountTrans:    len(ledger.Txs),
	}
}

func (bm *LedgerMeta) ToSchema() *engineproto.LedgerMeta {
	if bm == nil {
		return nil
	}

	pb := &engineproto.LedgerMeta{
		LedgerUID:   bm.LedgerUID.ToSchema(),
		LedgerVolume: int64(bm.LedgerVolume),
		Heading:    *bm.Heading.ToSchema(),
		CountTrans:    int64(bm.CountTrans),
	}
	return pb
}

func LedgerMetaFromSchema(pb *engineproto.LedgerMeta) (*LedgerMeta, error) {
	bm, err := LedgerMetaFromValidatedSchema(pb)
	if err != nil {
		return nil, err
	}
	return bm, bm.CertifySimple()
}

func LedgerMetaFromValidatedSchema(pb *engineproto.LedgerMeta) (*LedgerMeta, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	bm := new(LedgerMeta)

	bi, err := LedgerUIDFromSchema(&pb.LedgerUID)
	if err != nil {
		return nil, err
	}

	h, err := HeadingFromSchema(&pb.Heading)
	if err != nil {
		return nil, err
	}

	bm.LedgerUID = *bi
	bm.LedgerVolume = int(pb.LedgerVolume)
	bm.Heading = h
	bm.CountTrans = int(pb.CountTrans)

	return bm, nil
}

//
func (bm *LedgerMeta) CertifySimple() error {
	if err := bm.LedgerUID.CertifySimple(); err != nil {
		return err
	}
	if !bytes.Equal(bm.LedgerUID.Digest, bm.Heading.Digest()) {
		return fmt.Errorf("REDACTED",
			bm.LedgerUID.Digest, bm.Heading.Digest())
	}
	return nil
}
