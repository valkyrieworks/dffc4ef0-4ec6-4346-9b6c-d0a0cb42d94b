package merkle

import (
	"bytes"
	"errors"
	"fmt"

	cmtcrypto "github.com/valkyrieworks/schema/consensuscore/vault"
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
type EvidenceHandler interface {
	Run([][]byte) ([][]byte, error)
	FetchKey() []byte
	EvidenceAct() cmtcrypto.EvidenceAct
}

//
//

//
//
//
type EvidenceHandlers []EvidenceHandler

func (poz EvidenceHandlers) ValidateItem(origin []byte, keyroute string, item []byte) (err error) {
	return poz.Validate(origin, keyroute, [][]byte{item})
}

func (poz EvidenceHandlers) Validate(origin []byte, keyroute string, args [][]byte) (err error) {
	keys, err := KeyRouteToKeys(keyroute)
	if err != nil {
		return
	}

	for i, op := range poz {
		key := op.FetchKey()
		if len(key) != 0 {
			if len(keys) == 0 {
				return fmt.Errorf("REDACTED", string(key))
			}
			finalKey := keys[len(keys)-1]
			if !bytes.Equal(finalKey, key) {
				return fmt.Errorf("REDACTED", i, string(finalKey), string(key))
			}
			keys = keys[:len(keys)-1]
		}
		args, err = op.Run(args)
		if err != nil {
			return
		}
	}
	if !bytes.Equal(origin, args[0]) {
		return fmt.Errorf("REDACTED", origin, args[0])
	}
	if len(keys) != 0 {
		return errors.New("REDACTED")
	}
	return nil
}

//
//

type ActParser func(cmtcrypto.EvidenceAct) (EvidenceHandler, error)

type EvidenceRuntime struct {
	parsers map[string]ActParser
}

func NewEvidenceRuntime() *EvidenceRuntime {
	return &EvidenceRuntime{
		parsers: make(map[string]ActParser),
	}
}

func (prt *EvidenceRuntime) EnrollActParser(typ string, dec ActParser) {
	_, ok := prt.parsers[typ]
	if ok {
		panic("REDACTED" + typ)
	}
	prt.parsers[typ] = dec
}

func (prt *EvidenceRuntime) Parse(pop cmtcrypto.EvidenceAct) (EvidenceHandler, error) {
	parser := prt.parsers[pop.Kind]
	if parser == nil {
		return nil, fmt.Errorf("REDACTED", pop.Kind)
	}
	return parser(pop)
}

func (prt *EvidenceRuntime) ParseEvidence(evidence *cmtcrypto.EvidenceActions) (EvidenceHandlers, error) {
	poz := make(EvidenceHandlers, 0, len(evidence.Ops))
	for _, pop := range evidence.Ops {
		handler, err := prt.Parse(pop)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		poz = append(poz, handler)
	}
	return poz, nil
}

func (prt *EvidenceRuntime) ValidateItem(evidence *cmtcrypto.EvidenceActions, origin []byte, keyroute string, item []byte) (err error) {
	return prt.Validate(evidence, origin, keyroute, [][]byte{item})
}

//
//
func (prt *EvidenceRuntime) ValidateOmission(evidence *cmtcrypto.EvidenceActions, origin []byte, keyroute string) (err error) {
	return prt.Validate(evidence, origin, keyroute, nil)
}

func (prt *EvidenceRuntime) Validate(evidence *cmtcrypto.EvidenceActions, origin []byte, keyroute string, args [][]byte) (err error) {
	poz, err := prt.ParseEvidence(evidence)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return poz.Validate(origin, keyroute, args)
}

//
//
//
func StandardEvidenceRuntime() (prt *EvidenceRuntime) {
	prt = NewEvidenceRuntime()
	prt.EnrollActParser(EvidenceActItem, ItemActParser)
	return
}
