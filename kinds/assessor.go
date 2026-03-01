package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	ce "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

//
//
//
type Assessor struct {
	Location     Location       `json:"location"`
	PublicToken      security.PublicToken `json:"public_token"`
	BallotingPotency int64         `json:"balloting_potency"`

	NominatorUrgency int64 `json:"nominator_urgency"`
}

//
func FreshAssessor(publicToken security.PublicToken, ballotingPotency int64) *Assessor {
	return &Assessor{
		Location:          publicToken.Location(),
		PublicToken:           publicToken,
		BallotingPotency:      ballotingPotency,
		NominatorUrgency: 0,
	}
}

//
func (v *Assessor) CertifyFundamental() error {
	if v == nil {
		return errors.New("REDACTED")
	}
	if v.PublicToken == nil {
		return errors.New("REDACTED")
	}

	if v.BallotingPotency < 0 {
		return errors.New("REDACTED")
	}

	location := v.PublicToken.Location()
	if !bytes.Equal(v.Location, location) {
		return fmt.Errorf("REDACTED", location, v.Location)
	}

	return nil
}

//
//
func (v *Assessor) Duplicate() *Assessor {
	verDuplicate := *v
	return &verDuplicate
}

//
func (v *Assessor) ContrastNominatorUrgency(another *Assessor) *Assessor {
	if v == nil {
		return another
	}
	switch {
	case v.NominatorUrgency > another.NominatorUrgency:
		return v
	case v.NominatorUrgency < another.NominatorUrgency:
		return another
	default:
		outcome := bytes.Compare(v.Location, another.Location)
		switch {
		case outcome < 0:
			return v
		case outcome > 0:
			return another
		default:
			panic("REDACTED")
		}
	}
}

//
//
//
//
//
//
func (v *Assessor) Text() string {
	if v == nil {
		return "REDACTED"
	}
	return fmt.Sprintf("REDACTED",
		v.Location,
		v.PublicToken,
		v.BallotingPotency,
		v.NominatorUrgency)
}

//
func AssessorCatalogText(values []*Assessor) string {
	var sb strings.Builder
	for i, val := range values {
		if i > 0 {
			sb.WriteString("REDACTED")
		}
		sb.WriteString(val.Location.Text())
		sb.WriteString("REDACTED")
		sb.WriteString(strconv.FormatInt(val.BallotingPotency, 10))
	}
	return sb.String()
}

//
//
//
//
func (v *Assessor) Octets() []byte {
	pk, err := ce.PublicTokenTowardSchema(v.PublicToken)
	if err != nil {
		panic(err)
	}

	pbv := commitchema.PlainAssessor{
		PublicToken:      &pk,
		BallotingPotency: v.BallotingPotency,
	}

	bz, err := pbv.Serialize()
	if err != nil {
		panic(err)
	}
	return bz
}

//
func (v *Assessor) TowardSchema() (*commitchema.Assessor, error) {
	if v == nil {
		return nil, errors.New("REDACTED")
	}

	pk, err := ce.PublicTokenTowardSchema(v.PublicToken)
	if err != nil {
		return nil, err
	}

	vp := commitchema.Assessor{
		Location:          v.Location,
		PublicToken:           pk,
		BallotingPotency:      v.BallotingPotency,
		NominatorUrgency: v.NominatorUrgency,
	}

	return &vp, nil
}

//
//
func AssessorOriginatingSchema(vp *commitchema.Assessor) (*Assessor, error) {
	if vp == nil {
		return nil, errors.New("REDACTED")
	}

	pk, err := ce.PublicTokenOriginatingSchema(vp.PublicToken)
	if err != nil {
		return nil, err
	}
	v := new(Assessor)
	v.Location = vp.ObtainLocator()
	v.PublicToken = pk
	v.BallotingPotency = vp.ObtainBallotingPotency()
	v.NominatorUrgency = vp.ObtainNominatorUrgency()

	return v, nil
}

//
//

//
//
func ArbitraryAssessor(arbitraryPotency bool, minimumPotency int64) (*Assessor, PrivateAssessor) {
	privateItem := FreshSimulatePRV()
	ballotPotency := minimumPotency
	if arbitraryPotency {
		ballotPotency += int64(commitrand.Uint32n())
	}
	publicToken, err := privateItem.ObtainPublicToken()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	val := FreshAssessor(publicToken, ballotPotency)
	return val, privateItem
}
