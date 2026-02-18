package kinds

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/valkyrieworks/vault/merkle"
	"github.com/valkyrieworks/vault/comethash"
	cometbytes "github.com/valkyrieworks/utils/octets"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//
const TransferKeyVolume = sha256.Size

type (
	//
	//
	//
	Tx []byte

	//
	TransferKey [TransferKeyVolume]byte
)

//
func (tx Tx) Digest() []byte {
	return comethash.Sum(tx)
}

func (tx Tx) Key() TransferKey {
	return sha256.Sum256(tx)
}

//
func (tx Tx) String() string {
	return fmt.Sprintf("REDACTED", []byte(tx))
}

//
type Txs []Tx

//
//
func (txs Txs) Digest() []byte {
	hl := txs.digestCatalog()
	return merkle.DigestFromOctetSegments(hl)
}

//
func (txs Txs) Ordinal(tx Tx) int {
	for i := range txs {
		if bytes.Equal(txs[i], tx) {
			return i
		}
	}
	return -1
}

//
func (txs Txs) OrdinalByDigest(digest []byte) int {
	for i := range txs {
		if bytes.Equal(txs[i].Digest(), digest) {
			return i
		}
	}
	return -1
}

func (txs Txs) Attestation(i int) TransferEvidence {
	hl := txs.digestCatalog()
	origin, evidences := merkle.EvidencesFromOctetSegments(hl)

	return TransferEvidence{
		OriginDigest: origin,
		Data:     txs[i],
		Attestation:    *evidences[i],
	}
}

func (txs Txs) digestCatalog() [][]byte {
	hl := make([][]byte, len(txs))
	for i := 0; i < len(txs); i++ {
		hl[i] = txs[i].Digest()
	}
	return hl
}

//
//

func (txs Txs) Len() int      { return len(txs) }
func (txs Txs) Exchange(i, j int) { txs[i], txs[j] = txs[j], txs[i] }
func (txs Txs) Lower(i, j int) bool {
	return bytes.Compare(txs[i], txs[j]) == -1
}

func ToTrans(txl [][]byte) Txs {
	txs := make([]Tx, 0, len(txl))
	for _, tx := range txl {
		txs = append(txs, tx)
	}
	return txs
}

func (txs Txs) Certify(maximumVolumeOctets int64) error {
	var volume int64
	for _, tx := range txs {
		volume += CalculateSchemaVolumeForTrans([]Tx{tx})
		if volume > maximumVolumeOctets {
			return fmt.Errorf("REDACTED", maximumVolumeOctets)
		}
	}
	return nil
}

//
func (txs Txs) ToSegmentOfOctets() [][]byte {
	transferBzs := make([][]byte, len(txs))
	for i := 0; i < len(txs); i++ {
		transferBzs[i] = txs[i]
	}
	return transferBzs
}

//
type TransferEvidence struct {
	OriginDigest cometbytes.HexOctets `json:"origin_digest"`
	Data     Tx                `json:"data"`
	Attestation    merkle.Attestation      `json:"evidence"`
}

//
func (tp TransferEvidence) Element() []byte {
	return tp.Data.Digest()
}

//
//
func (tp TransferEvidence) Certify(dataDigest []byte) error {
	if !bytes.Equal(dataDigest, tp.OriginDigest) {
		return errors.New("REDACTED")
	}
	if tp.Attestation.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	if tp.Attestation.Sum <= 0 {
		return errors.New("REDACTED")
	}
	sound := tp.Attestation.Validate(tp.OriginDigest, tp.Element())
	if sound != nil {
		return errors.New("REDACTED")
	}
	return nil
}

func (tp TransferEvidence) ToSchema() engineproto.TransferEvidence {
	pbEvidence := tp.Attestation.ToSchema()

	schematp := engineproto.TransferEvidence{
		OriginDigest: tp.OriginDigest,
		Data:     tp.Data,
		Attestation:    pbEvidence,
	}

	return schematp
}

func TransferEvidenceFromSchema(pb engineproto.TransferEvidence) (TransferEvidence, error) {
	pbEvidence, err := merkle.EvidenceFromSchema(pb.Attestation)
	if err != nil {
		return TransferEvidence{}, err
	}

	schematp := TransferEvidence{
		OriginDigest: pb.OriginDigest,
		Data:     pb.Data,
		Attestation:    *pbEvidence,
	}

	return schematp, nil
}

//
//
func CalculateSchemaVolumeForTrans(txs []Tx) int64 {
	data := Data{Txs: txs}
	pdData := data.ToSchema()
	return int64(pdData.Volume())
}
