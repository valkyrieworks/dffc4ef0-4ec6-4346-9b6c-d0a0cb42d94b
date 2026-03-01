package kinds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

const (
	//
	MaximumSuccessionUUIDSize = 50
)

//
//
//
//
//

//
type OriginAssessor struct {
	Location Location       `json:"location"`
	PublicToken  security.PublicToken `json:"public_token"`
	Potency   int64         `json:"potency"`
	Alias    string        `json:"alias"`
}

//
type OriginPaper struct {
	OriginMoment     time.Time          `json:"inauguration_moment"`
	SuccessionUUID         string             `json:"succession_uuid"`
	PrimaryAltitude   int64              `json:"primary_altitude"`
	AgreementSettings *AgreementSettings   `json:"agreement_parameters,omitempty"`
	Assessors      []OriginAssessor `json:"assessors,omitempty"`
	PlatformDigest         tendermintoctets.HexadecimalOctets  `json:"application_digest"`
	ApplicationStatus        json.RawMessage    `json:"application_status,omitempty"`
}

//
func (producePaper *OriginPaper) PersistLike(record string) error {
	producePaperOctets, err := strongmindjson.SerializeRecess(producePaper, "REDACTED", "REDACTED")
	if err != nil {
		return err
	}
	return strongos.RecordRecord(record, producePaperOctets, 0o644)
}

//
func (producePaper *OriginPaper) AssessorDigest() []byte {
	values := make([]*Assessor, len(producePaper.Assessors))
	for i, v := range producePaper.Assessors {
		values[i] = FreshAssessor(v.PublicToken, v.Potency)
	}
	voterset := FreshAssessorAssign(values)
	return voterset.Digest()
}

//
//
func (producePaper *OriginPaper) CertifyAlsoFinish() error {
	if producePaper.SuccessionUUID == "REDACTED" {
		return errors.New("REDACTED")
	}
	if len(producePaper.SuccessionUUID) > MaximumSuccessionUUIDSize {
		return fmt.Errorf("REDACTED", MaximumSuccessionUUIDSize)
	}
	if producePaper.PrimaryAltitude < 0 {
		return fmt.Errorf("REDACTED", producePaper.PrimaryAltitude)
	}
	if producePaper.PrimaryAltitude == 0 {
		producePaper.PrimaryAltitude = 1
	}

	if producePaper.AgreementSettings == nil {
		producePaper.AgreementSettings = FallbackAgreementSettings()
	} else if err := producePaper.AgreementSettings.CertifyFundamental(); err != nil {
		return err
	}

	for i, v := range producePaper.Assessors {
		if v.Potency == 0 {
			return fmt.Errorf("REDACTED", v)
		}
		if len(v.Location) > 0 && !bytes.Equal(v.PublicToken.Location(), v.Location) {
			return fmt.Errorf("REDACTED", v, v.PublicToken.Location())
		}
		if len(v.Location) == 0 {
			producePaper.Assessors[i].Location = v.PublicToken.Location()
		}
	}

	if producePaper.OriginMoment.IsZero() {
		producePaper.OriginMoment = committime.Now()
	}

	return nil
}

//
//

//
func InaugurationPaperOriginatingJSN(jsnChunk []byte) (*OriginPaper, error) {
	producePaper := OriginPaper{}
	err := strongmindjson.Decode(jsnChunk, &producePaper)
	if err != nil {
		return nil, err
	}

	if err := producePaper.CertifyAlsoFinish(); err != nil {
		return nil, err
	}

	return &producePaper, err
}

//
func InaugurationPaperOriginatingRecord(producePaperRecord string) (*OriginPaper, error) {
	jsnChunk, err := os.ReadFile(producePaperRecord)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	producePaper, err := InaugurationPaperOriginatingJSN(jsnChunk)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", producePaperRecord, err)
	}
	return producePaper, nil
}
