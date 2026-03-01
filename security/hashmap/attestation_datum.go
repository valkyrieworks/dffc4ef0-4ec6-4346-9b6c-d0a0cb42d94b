package hashmap

import (
	"bytes"
	"fmt"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	cmtsecurity "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

const AttestationActionDatum = "REDACTED"

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
type DatumAction struct {
	//
	key []byte

	//
	Attestation *Attestation `json:"attestation"`
}

var _ AttestationHandler = DatumAction{}

func FreshDatumAction(key []byte, attestation *Attestation) DatumAction {
	return DatumAction{
		key:   key,
		Attestation: attestation,
	}
}

func DatumActionDeserializer(pop cmtsecurity.AttestationAction) (AttestationHandler, error) {
	if pop.Kind != AttestationActionDatum {
		return nil, fmt.Errorf("REDACTED", pop.Kind, AttestationActionDatum)
	}
	var schemaaction cmtsecurity.DatumAction //
	err := schemaaction.Decode(pop.Data)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	sp, err := AttestationOriginatingSchema(schemaaction.Attestation)
	if err != nil {
		return nil, err
	}
	return FreshDatumAction(pop.Key, sp), nil
}

func (op DatumAction) AttestationAction() cmtsecurity.AttestationAction {
	schemadatum := cmtsecurity.DatumAction{
		Key:   op.key,
		Attestation: op.Attestation.TowardSchema(),
	}
	bz, err := schemadatum.Serialize()
	if err != nil {
		panic(err)
	}
	return cmtsecurity.AttestationAction{
		Kind: AttestationActionDatum,
		Key:  op.key,
		Data: bz,
	}
}

func (op DatumAction) Text() string {
	return fmt.Sprintf("REDACTED", op.ObtainToken())
}

func (op DatumAction) Run(arguments [][]byte) ([][]byte, error) {
	if len(arguments) != 1 {
		return nil, fmt.Errorf("REDACTED", len(arguments))
	}
	datum := arguments[0]
	digester := tenderminthash.New()
	digester.Write(datum)
	datadigest := digester.Sum(nil)

	bz := new(bytes.Buffer)
	//
	serializeOctetSegment(bz, op.key) //
	serializeOctetSegment(bz, datadigest)  //
	mapdigest := terminalDigest(bz.Bytes())

	if !bytes.Equal(mapdigest, op.Attestation.NodeDigest) {
		return nil, fmt.Errorf("REDACTED", op.Attestation.NodeDigest, mapdigest)
	}

	originDigest, err := op.Attestation.calculateOriginDigest(tenderminthash.New())
	if err != nil {
		return nil, err
	}
	return [][]byte{
		originDigest,
	}, nil
}

func (op DatumAction) ObtainToken() []byte {
	return op.key
}
