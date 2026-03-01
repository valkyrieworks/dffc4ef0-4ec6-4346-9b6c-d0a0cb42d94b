package agile

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
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
//
type RelianceChoices struct {
	//
	//
	//
	//
	//
	//
	//
	//
	Cycle time.Duration

	//
	//
	Altitude int64
	Digest   []byte
}

//
func (choices RelianceChoices) CertifyFundamental() error {
	if choices.Cycle <= 0 {
		return errors.New("REDACTED")
	}
	if choices.Altitude <= 0 {
		return errors.New("REDACTED")
	}
	if len(choices.Digest) != tenderminthash.Extent {
		return fmt.Errorf("REDACTED",
			tenderminthash.Extent,
			len(choices.Digest),
		)
	}
	return nil
}
