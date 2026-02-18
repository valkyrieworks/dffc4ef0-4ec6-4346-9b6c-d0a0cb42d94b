package rapid

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/vault/comethash"
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
type ValidateOptions struct {
	//
	//
	//
	//
	//
	//
	//
	//
	Duration time.Duration

	//
	//
	Level int64
	Digest   []byte
}

//
func (opts ValidateOptions) CertifySimple() error {
	if opts.Duration <= 0 {
		return errors.New("REDACTED")
	}
	if opts.Level <= 0 {
		return errors.New("REDACTED")
	}
	if len(opts.Digest) != comethash.Volume {
		return fmt.Errorf("REDACTED",
			comethash.Volume,
			len(opts.Digest),
		)
	}
	return nil
}
