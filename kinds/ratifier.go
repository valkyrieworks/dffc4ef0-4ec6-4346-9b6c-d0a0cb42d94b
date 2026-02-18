package kinds

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/valkyrieworks/vault"
	ce "github.com/valkyrieworks/vault/codec"
	engineseed "github.com/valkyrieworks/utils/random"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

//
//
//
type Ratifier struct {
	Location     Location       `json:"location"`
	PublicKey      vault.PublicKey `json:"public_key"`
	PollingEnergy int64         `json:"polling_energy"`

	RecommenderUrgency int64 `json:"recommender_urgency"`
}

//
func NewRatifier(publicKey vault.PublicKey, pollingEnergy int64) *Ratifier {
	return &Ratifier{
		Location:          publicKey.Location(),
		PublicKey:           publicKey,
		PollingEnergy:      pollingEnergy,
		RecommenderUrgency: 0,
	}
}

//
func (v *Ratifier) CertifySimple() error {
	if v == nil {
		return errors.New("REDACTED")
	}
	if v.PublicKey == nil {
		return errors.New("REDACTED")
	}

	if v.PollingEnergy < 0 {
		return errors.New("REDACTED")
	}

	address := v.PublicKey.Location()
	if !bytes.Equal(v.Location, address) {
		return fmt.Errorf("REDACTED", address, v.Location)
	}

	return nil
}

//
//
func (v *Ratifier) Clone() *Ratifier {
	vClone := *v
	return &vClone
}

//
func (v *Ratifier) ContrastRecommenderUrgency(another *Ratifier) *Ratifier {
	if v == nil {
		return another
	}
	switch {
	case v.RecommenderUrgency > another.RecommenderUrgency:
		return v
	case v.RecommenderUrgency < another.RecommenderUrgency:
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
func (v *Ratifier) String() string {
	if v == nil {
		return "REDACTED"
	}
	return fmt.Sprintf("REDACTED",
		v.Location,
		v.PublicKey,
		v.PollingEnergy,
		v.RecommenderUrgency)
}

//
func RatifierCatalogString(values []*Ratifier) string {
	var sb strings.Builder
	for i, val := range values {
		if i > 0 {
			sb.WriteString("REDACTED")
		}
		sb.WriteString(val.Location.String())
		sb.WriteString("REDACTED")
		sb.WriteString(strconv.FormatInt(val.PollingEnergy, 10))
	}
	return sb.String()
}

//
//
//
//
func (v *Ratifier) Octets() []byte {
	pk, err := ce.PublicKeyToSchema(v.PublicKey)
	if err != nil {
		panic(err)
	}

	pbv := engineproto.BasicRatifier{
		PublicKey:      &pk,
		PollingEnergy: v.PollingEnergy,
	}

	bz, err := pbv.Serialize()
	if err != nil {
		panic(err)
	}
	return bz
}

//
func (v *Ratifier) ToSchema() (*engineproto.Ratifier, error) {
	if v == nil {
		return nil, errors.New("REDACTED")
	}

	pk, err := ce.PublicKeyToSchema(v.PublicKey)
	if err != nil {
		return nil, err
	}

	vp := engineproto.Ratifier{
		Location:          v.Location,
		PublicKey:           pk,
		PollingEnergy:      v.PollingEnergy,
		RecommenderUrgency: v.RecommenderUrgency,
	}

	return &vp, nil
}

//
//
func RatifierFromSchema(vp *engineproto.Ratifier) (*Ratifier, error) {
	if vp == nil {
		return nil, errors.New("REDACTED")
	}

	pk, err := ce.PublicKeyFromSchema(vp.PublicKey)
	if err != nil {
		return nil, err
	}
	v := new(Ratifier)
	v.Location = vp.FetchLocation()
	v.PublicKey = pk
	v.PollingEnergy = vp.FetchPollingEnergy()
	v.RecommenderUrgency = vp.FetchRecommenderUrgency()

	return v, nil
}

//
//

//
//
func RandomRatifier(randomEnergy bool, minimumEnergy int64) (*Ratifier, PrivateRatifier) {
	privateValue := NewEmulatePV()
	ballotEnergy := minimumEnergy
	if randomEnergy {
		ballotEnergy += int64(engineseed.Uint32())
	}
	publicKey, err := privateValue.FetchPublicKey()
	if err != nil {
		panic(fmt.Errorf("REDACTED", err))
	}
	val := NewRatifier(publicKey, ballotEnergy)
	return val, privateValue
}
