package octets

import (
	"encoding/hex"
	"fmt"
	"strings"
)

//
type HexOctets []byte

//
func (bz HexOctets) Serialize() ([]byte, error) {
	return bz, nil
}

//
func (bz *HexOctets) Unserialize(data []byte) error {
	*bz = data
	return nil
}

//
func (bz HexOctets) SerializeJSON() ([]byte, error) {
	s := strings.ToUpper(hex.EncodeToString(bz))
	jbz := make([]byte, len(s)+2)
	jbz[0] = '"'
	copy(jbz[1:], s)
	jbz[len(jbz)-1] = '"'
	return jbz, nil
}

//
func (bz *HexOctets) UnserializeJSON(data []byte) error {
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return fmt.Errorf("REDACTED", data)
	}
	bz2, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	*bz = bz2
	return nil
}

//
func (bz HexOctets) Octets() []byte {
	return bz
}

func (bz HexOctets) String() string {
	return strings.ToUpper(hex.EncodeToString(bz))
}

//
//
//
func (bz HexOctets) Layout(s fmt.State, command rune) {
	switch command {
	case 'p':
		fmt.Fprintf(s, "REDACTED", bz)
	default:
		fmt.Fprintf(s, "REDACTED", []byte(bz))
	}
}
