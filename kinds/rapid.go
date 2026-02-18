package kinds

import (
	"bytes"
	"errors"
	"fmt"

	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//
//
type RapidLedger struct {
	*AttestedHeading `json:"attested_heading"`
	RatifierAssign  *RatifierAssign `json:"ratifier_collection"`
}

//
//
//
func (lb RapidLedger) CertifySimple(ledgerUID string) error {
	if lb.AttestedHeading == nil {
		return errors.New("REDACTED")
	}
	if lb.RatifierAssign == nil {
		return errors.New("REDACTED")
	}

	if err := lb.AttestedHeading.CertifySimple(ledgerUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := lb.RatifierAssign.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if valueCollectionDigest := lb.RatifierAssign.Digest(); !bytes.Equal(lb.RatifiersDigest, valueCollectionDigest) {
		return fmt.Errorf("REDACTED",
			lb.RatifiersDigest, valueCollectionDigest,
		)
	}

	return nil
}

//
func (lb RapidLedger) String() string {
	return lb.StringIndented("REDACTED")
}

//
//
//
//
func (lb RapidLedger) StringIndented(indent string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED`,
		indent, lb.AttestedHeading.StringIndented(indent+"REDACTED"),
		indent, lb.RatifierAssign.StringIndented(indent+"REDACTED"),
		indent)
}

//
func (lb *RapidLedger) ToSchema() (*engineproto.RapidLedger, error) {
	if lb == nil {
		return nil, nil
	}

	lbp := new(engineproto.RapidLedger)
	var err error
	if lb.AttestedHeading != nil {
		lbp.AttestedHeading = lb.AttestedHeading.ToSchema()
	}
	if lb.RatifierAssign != nil {
		lbp.RatifierAssign, err = lb.RatifierAssign.ToSchema()
		if err != nil {
			return nil, err
		}
	}

	return lbp, nil
}

//
//
func RapidLedgerFromSchema(pb *engineproto.RapidLedger) (*RapidLedger, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	lb := new(RapidLedger)

	if pb.AttestedHeading != nil {
		sh, err := AttestedHeadingFromSchema(pb.AttestedHeading)
		if err != nil {
			return nil, err
		}
		lb.AttestedHeading = sh
	}

	if pb.RatifierAssign != nil {
		values, err := RatifierCollectionFromSchema(pb.RatifierAssign)
		if err != nil {
			return nil, err
		}
		lb.RatifierAssign = values
	}

	return lb, nil
}

//

//
type AttestedHeading struct {
	*Heading `json:"heading"`

	Endorse *Endorse `json:"endorse"`
}

//
func (sh AttestedHeading) IsEmpty() bool {
	return sh.Heading == nil && sh.Endorse == nil
}

//
//
//
//
//
//
func (sh AttestedHeading) CertifySimple(ledgerUID string) error {
	if sh.Heading == nil {
		return errors.New("REDACTED")
	}
	if sh.Endorse == nil {
		return errors.New("REDACTED")
	}

	if err := sh.Heading.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := sh.Endorse.CertifySimple(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if sh.LedgerUID != ledgerUID {
		return fmt.Errorf("REDACTED", sh.LedgerUID, ledgerUID)
	}

	//
	if sh.Endorse.Level != sh.Level {
		return fmt.Errorf("REDACTED", sh.Level, sh.Endorse.Level)
	}
	if dhash, chash := sh.Digest(), sh.Endorse.LedgerUID.Digest; !bytes.Equal(dhash, chash) {
		return fmt.Errorf("REDACTED", chash, dhash)
	}

	return nil
}

//
func (sh AttestedHeading) String() string {
	return sh.StringIndented("REDACTED")
}

//
//
//
//
func (sh AttestedHeading) StringIndented(indent string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED`,
		indent, sh.Heading.StringIndented(indent+"REDACTED"),
		indent, sh.Endorse.StringIndented(indent+"REDACTED"),
		indent)
}

//
func (sh *AttestedHeading) ToSchema() *engineproto.AttestedHeading {
	if sh == nil {
		return nil
	}

	psh := new(engineproto.AttestedHeading)
	if sh.Heading != nil {
		psh.Heading = sh.Heading.ToSchema()
	}
	if sh.Endorse != nil {
		psh.Endorse = sh.Endorse.ToSchema()
	}

	return psh
}

//
//
func AttestedHeadingFromSchema(shp *engineproto.AttestedHeading) (*AttestedHeading, error) {
	if shp == nil {
		return nil, errors.New("REDACTED")
	}

	sh := new(AttestedHeading)

	if shp.Heading != nil {
		h, err := HeadingFromSchema(shp.Heading)
		if err != nil {
			return nil, err
		}
		sh.Heading = &h
	}

	if shp.Endorse != nil {
		c, err := EndorseFromSchema(shp.Endorse)
		if err != nil {
			return nil, err
		}
		sh.Endorse = c
	}

	return sh, nil
}
