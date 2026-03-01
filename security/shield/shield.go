package shield

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/crypto/openpgp/armor" //
)

func SerializeShield(ledgerKind string, headings map[string]string, data []byte) string {
	buf := new(bytes.Buffer)
	w, err := armor.Encode(buf, ledgerKind, headings)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	_, err = w.Write(data)
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	err = w.Close()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	return buf.String()
}

func DeserializeShield(shieldTxt string) (ledgerKind string, headings map[string]string, data []byte, err error) {
	buf := bytes.NewBufferString(shieldTxt)
	ledger, err := armor.Decode(buf)
	if err != nil {
		return "REDACTED", nil, nil, err
	}
	data, err = io.ReadAll(ledger.Body)
	if err != nil {
		return "REDACTED", nil, nil, err
	}
	return ledger.Type, ledger.Header, data, nil
}
