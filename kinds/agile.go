package kinds

import (
	"bytes"
	"errors"
	"fmt"

	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//
//
type AgileLedger struct {
	*NotatedHeading `json:"notated_heading"`
	AssessorAssign  *AssessorAssign `json:"assessor_assign"`
}

//
//
//
func (lb AgileLedger) CertifyFundamental(successionUUID string) error {
	if lb.NotatedHeading == nil {
		return errors.New("REDACTED")
	}
	if lb.AssessorAssign == nil {
		return errors.New("REDACTED")
	}

	if err := lb.NotatedHeading.CertifyFundamental(successionUUID); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := lb.AssessorAssign.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	//
	if itemAssignDigest := lb.AssessorAssign.Digest(); !bytes.Equal(lb.AssessorsDigest, itemAssignDigest) {
		return fmt.Errorf("REDACTED",
			lb.AssessorsDigest, itemAssignDigest,
		)
	}

	return nil
}

//
func (lb AgileLedger) Text() string {
	return lb.TextFormatted("REDACTED")
}

//
//
//
//
func (lb AgileLedger) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED`,
		format, lb.NotatedHeading.TextFormatted(format+"REDACTED"),
		format, lb.AssessorAssign.TextFormatted(format+"REDACTED"),
		format)
}

//
func (lb *AgileLedger) TowardSchema() (*commitchema.AgileLedger, error) {
	if lb == nil {
		return nil, nil
	}

	lbp := new(commitchema.AgileLedger)
	var err error
	if lb.NotatedHeading != nil {
		lbp.NotatedHeading = lb.NotatedHeading.TowardSchema()
	}
	if lb.AssessorAssign != nil {
		lbp.AssessorAssign, err = lb.AssessorAssign.TowardSchema()
		if err != nil {
			return nil, err
		}
	}

	return lbp, nil
}

//
//
func AgileLedgerOriginatingSchema(pb *commitchema.AgileLedger) (*AgileLedger, error) {
	if pb == nil {
		return nil, errors.New("REDACTED")
	}

	lb := new(AgileLedger)

	if pb.NotatedHeading != nil {
		sh, err := NotatedHeadingOriginatingSchema(pb.NotatedHeading)
		if err != nil {
			return nil, err
		}
		lb.NotatedHeading = sh
	}

	if pb.AssessorAssign != nil {
		values, err := AssessorAssignOriginatingSchema(pb.AssessorAssign)
		if err != nil {
			return nil, err
		}
		lb.AssessorAssign = values
	}

	return lb, nil
}

//

//
type NotatedHeading struct {
	*Heading `json:"heading"`

	Endorse *Endorse `json:"endorse"`
}

//
func (sh NotatedHeading) EqualsBlank() bool {
	return sh.Heading == nil && sh.Endorse == nil
}

//
//
//
//
//
//
func (sh NotatedHeading) CertifyFundamental(successionUUID string) error {
	if sh.Heading == nil {
		return errors.New("REDACTED")
	}
	if sh.Endorse == nil {
		return errors.New("REDACTED")
	}

	if err := sh.Heading.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	if err := sh.Endorse.CertifyFundamental(); err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	if sh.SuccessionUUID != successionUUID {
		return fmt.Errorf("REDACTED", sh.SuccessionUUID, successionUUID)
	}

	//
	if sh.Endorse.Altitude != sh.Altitude {
		return fmt.Errorf("REDACTED", sh.Altitude, sh.Endorse.Altitude)
	}
	if headerhash, chnlhash := sh.Digest(), sh.Endorse.LedgerUUID.Digest; !bytes.Equal(headerhash, chnlhash) {
		return fmt.Errorf("REDACTED", chnlhash, headerhash)
	}

	return nil
}

//
func (sh NotatedHeading) Text() string {
	return sh.TextFormatted("REDACTED")
}

//
//
//
//
func (sh NotatedHeading) TextFormatted(format string) string {
	return fmt.Sprintf(`REDACTED{
REDACTEDv
REDACTEDv
REDACTED`,
		format, sh.Heading.TextFormatted(format+"REDACTED"),
		format, sh.Endorse.TextFormatted(format+"REDACTED"),
		format)
}

//
func (sh *NotatedHeading) TowardSchema() *commitchema.NotatedHeading {
	if sh == nil {
		return nil
	}

	psh := new(commitchema.NotatedHeading)
	if sh.Heading != nil {
		psh.Heading = sh.Heading.TowardSchema()
	}
	if sh.Endorse != nil {
		psh.Endorse = sh.Endorse.TowardSchema()
	}

	return psh
}

//
//
func NotatedHeadingOriginatingSchema(shp *commitchema.NotatedHeading) (*NotatedHeading, error) {
	if shp == nil {
		return nil, errors.New("REDACTED")
	}

	sh := new(NotatedHeading)

	if shp.Heading != nil {
		h, err := HeadingOriginatingSchema(shp.Heading)
		if err != nil {
			return nil, err
		}
		sh.Heading = &h
	}

	if shp.Endorse != nil {
		c, err := EndorseOriginatingSchema(shp.Endorse)
		if err != nil {
			return nil, err
		}
		sh.Endorse = c
	}

	return sh, nil
}
