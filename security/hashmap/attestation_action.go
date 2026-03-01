package hashmap

import (
	"bytes"
	"errors"
	"fmt"

	cmtsecurity "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

//
//

//
//
//
//
//
//
//
type AttestationHandler interface {
	Run([][]byte) ([][]byte, error)
	ObtainToken() []byte
	AttestationAction() cmtsecurity.AttestationAction
}

//
//

//
//
//
type AttestationHandlers []AttestationHandler

func (poz AttestationHandlers) ValidateDatum(origin []byte, tokenpath string, datum []byte) (err error) {
	return poz.Validate(origin, tokenpath, [][]byte{datum})
}

func (poz AttestationHandlers) Validate(origin []byte, tokenpath string, arguments [][]byte) (err error) {
	tokens, err := TokenRouteTowardTokens(tokenpath)
	if err != nil {
		return
	}

	for i, op := range poz {
		key := op.ObtainToken()
		if len(key) != 0 {
			if len(tokens) == 0 {
				return fmt.Errorf("REDACTED", string(key))
			}
			finalToken := tokens[len(tokens)-1]
			if !bytes.Equal(finalToken, key) {
				return fmt.Errorf("REDACTED", i, string(finalToken), string(key))
			}
			tokens = tokens[:len(tokens)-1]
		}
		arguments, err = op.Run(arguments)
		if err != nil {
			return
		}
	}
	if !bytes.Equal(origin, arguments[0]) {
		return fmt.Errorf("REDACTED", origin, arguments[0])
	}
	if len(tokens) != 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
//

type ActionDeserializer func(cmtsecurity.AttestationAction) (AttestationHandler, error)

type AttestationExecution struct {
	deserializers map[string]ActionDeserializer
}

func FreshAttestationExecution() *AttestationExecution {
	return &AttestationExecution{
		deserializers: make(map[string]ActionDeserializer),
	}
}

func (prt *AttestationExecution) EnrollActionDeserializer(typ string, dec ActionDeserializer) {
	_, ok := prt.deserializers[typ]
	if ok {
		panic("REDACTED" + typ)
	}
	prt.deserializers[typ] = dec
}

func (prt *AttestationExecution) Deserialize(pop cmtsecurity.AttestationAction) (AttestationHandler, error) {
	deserializer := prt.deserializers[pop.Kind]
	if deserializer == nil {
		return nil, fmt.Errorf("REDACTED", pop.Kind)
	}
	return deserializer(pop)
}

func (prt *AttestationExecution) DeserializeAttestation(attestation *cmtsecurity.AttestationActions) (AttestationHandlers, error) {
	poz := make(AttestationHandlers, 0, len(attestation.Ops))
	for _, pop := range attestation.Ops {
		handler, err := prt.Deserialize(pop)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		poz = append(poz, handler)
	}
	return poz, nil
}

func (prt *AttestationExecution) ValidateDatum(attestation *cmtsecurity.AttestationActions, origin []byte, tokenpath string, datum []byte) (err error) {
	return prt.Validate(attestation, origin, tokenpath, [][]byte{datum})
}

//
//
func (prt *AttestationExecution) ValidateOmission(attestation *cmtsecurity.AttestationActions, origin []byte, tokenpath string) (err error) {
	return prt.Validate(attestation, origin, tokenpath, nil)
}

func (prt *AttestationExecution) Validate(attestation *cmtsecurity.AttestationActions, origin []byte, tokenpath string, arguments [][]byte) (err error) {
	poz, err := prt.DeserializeAttestation(attestation)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return poz.Validate(origin, tokenpath, arguments)
}

//
//
//
func FallbackAttestationExecution() (prt *AttestationExecution) {
	prt = FreshAttestationExecution()
	prt.EnrollActionDeserializer(AttestationActionDatum, DatumActionDeserializer)
	return
}
