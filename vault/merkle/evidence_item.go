package merkle

import (
	"bytes"
	"fmt"

	"github.com/valkyrieworks/vault/comethash"
	cmtcrypto "github.com/valkyrieworks/schema/consensuscore/vault"
)

const EvidenceActItem = "REDACTED"

//
//
//
//
//
//
//
//
//
//
type ItemAct struct {
	//
	key []byte

	//
	Attestation *Attestation `json:"evidence"`
}

var _ EvidenceHandler = ItemAct{}

func NewItemAct(key []byte, evidence *Attestation) ItemAct {
	return ItemAct{
		key:   key,
		Attestation: evidence,
	}
}

func ItemActParser(pop cmtcrypto.EvidenceAct) (EvidenceHandler, error) {
	if pop.Kind != EvidenceActItem {
		return nil, fmt.Errorf("REDACTED", pop.Kind, EvidenceActItem)
	}
	var pbop cmtcrypto.ItemAct //
	err := pbop.Unserialize(pop.Data)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	sp, err := EvidenceFromSchema(pbop.Attestation)
	if err != nil {
		return nil, err
	}
	return NewItemAct(pop.Key, sp), nil
}

func (op ItemAct) EvidenceAct() cmtcrypto.EvidenceAct {
	pbitem := cmtcrypto.ItemAct{
		Key:   op.key,
		Attestation: op.Attestation.ToSchema(),
	}
	bz, err := pbitem.Serialize()
	if err != nil {
		panic(err)
	}
	return cmtcrypto.EvidenceAct{
		Kind: EvidenceActItem,
		Key:  op.key,
		Data: bz,
	}
}

func (op ItemAct) String() string {
	return fmt.Sprintf("REDACTED", op.FetchKey())
}

func (op ItemAct) Run(args [][]byte) ([][]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("REDACTED", len(args))
	}
	item := args[0]
	digester := comethash.New()
	digester.Write(item)
	vdigest := digester.Sum(nil)

	bz := new(bytes.Buffer)
	//
	encodeOctetSegment(bz, op.key) //
	encodeOctetSegment(bz, vdigest)  //
	kvdigest := elementDigest(bz.Bytes())

	if !bytes.Equal(kvdigest, op.Attestation.NodeDigest) {
		return nil, fmt.Errorf("REDACTED", op.Attestation.NodeDigest, kvdigest)
	}

	originDigest, err := op.Attestation.calculateOriginDigest(comethash.New())
	if err != nil {
		return nil, err
	}
	return [][]byte{
		originDigest,
	}, nil
}

func (op ItemAct) FetchKey() []byte {
	return op.key
}
