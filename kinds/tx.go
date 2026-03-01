package kinds

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//
const TransferTokenExtent = sha256.Size

type (
	//
	//
	//
	Tx []byte

	//
	TransferToken [TransferTokenExtent]byte
)

//
func (tx Tx) Digest() []byte {
	return tenderminthash.Sum(tx)
}

func (tx Tx) Key() TransferToken {
	return sha256.Sum256(tx)
}

//
func (tx Tx) Text() string {
	return fmt.Sprintf("REDACTED", []byte(tx))
}

//
type Txs []Tx

//
//
func (txs Txs) Digest() []byte {
	hl := txs.digestCatalog()
	return hashmap.DigestOriginatingOctetSegments(hl)
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
func (txs Txs) PositionViaDigest(digest []byte) int {
	for i := range txs {
		if bytes.Equal(txs[i].Digest(), digest) {
			return i
		}
	}
	return -1
}

func (txs Txs) Attestation(i int) TransferAttestation {
	hl := txs.digestCatalog()
	origin, attestations := hashmap.AttestationsOriginatingOctetSegments(hl)

	return TransferAttestation{
		OriginDigest: origin,
		Data:     txs[i],
		Attestation:    *attestations[i],
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
func (txs Txs) Inferior(i, j int) bool {
	return bytes.Compare(txs[i], txs[j]) == -1
}

func TowardTrans(txl [][]byte) Txs {
	txs := make([]Tx, 0, len(txl))
	for _, tx := range txl {
		txs = append(txs, tx)
	}
	return txs
}

func (txs Txs) Certify(maximumExtentOctets int64) error {
	var extent int64
	for _, tx := range txs {
		extent += CalculateSchemaExtentForeachTrans([]Tx{tx})
		if extent > maximumExtentOctets {
			return fmt.Errorf("REDACTED", maximumExtentOctets)
		}
	}
	return nil
}

//
func (txs Txs) TowardSegmentBelongingOctets() [][]byte {
	transferByteslices := make([][]byte, len(txs))
	for i := 0; i < len(txs); i++ {
		transferByteslices[i] = txs[i]
	}
	return transferByteslices
}

//
type TransferAttestation struct {
	OriginDigest tendermintoctets.HexadecimalOctets `json:"origin_digest"`
	Data     Tx                `json:"data"`
	Attestation    hashmap.Attestation      `json:"attestation"`
}

//
func (tp TransferAttestation) Node() []byte {
	return tp.Data.Digest()
}

//
//
func (tp TransferAttestation) Certify(dataDigest []byte) error {
	if !bytes.Equal(dataDigest, tp.OriginDigest) {
		return errors.New("REDACTED")
	}
	if tp.Attestation.Ordinal < 0 {
		return errors.New("REDACTED")
	}
	if tp.Attestation.Sum <= 0 {
		return errors.New("REDACTED")
	}
	sound := tp.Attestation.Validate(tp.OriginDigest, tp.Node())
	if sound != nil {
		return errors.New("REDACTED")
	}
	return nil
}

func (tp TransferAttestation) TowardSchema() commitchema.TransferAttestation {
	bufferAttestation := tp.Attestation.TowardSchema()

	schemap := commitchema.TransferAttestation{
		OriginDigest: tp.OriginDigest,
		Data:     tp.Data,
		Attestation:    bufferAttestation,
	}

	return schemap
}

func TransferAttestationOriginatingSchema(pb commitchema.TransferAttestation) (TransferAttestation, error) {
	bufferAttestation, err := hashmap.AttestationOriginatingSchema(pb.Attestation)
	if err != nil {
		return TransferAttestation{}, err
	}

	schemap := TransferAttestation{
		OriginDigest: pb.OriginDigest,
		Data:     pb.Data,
		Attestation:    *bufferAttestation,
	}

	return schemap, nil
}

//
//
func CalculateSchemaExtentForeachTrans(txs []Tx) int64 {
	data := Data{Txs: txs}
	pieceData := data.TowardSchema()
	return int64(pieceData.Extent())
}
