package kinds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/valkyrieworks/vault"
	cometbytes "github.com/valkyrieworks/utils/octets"
	cometjson "github.com/valkyrieworks/utils/json"
	cometos "github.com/valkyrieworks/utils/os"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

const (
	//
	MaximumSeriesUIDSize = 50
)

//
//
//
//
//

//
type OriginRatifier struct {
	Location Location       `json:"location"`
	PublicKey  vault.PublicKey `json:"public_key"`
	Energy   int64         `json:"energy"`
	Label    string        `json:"label"`
}

//
type OriginPaper struct {
	OriginMoment     time.Time          `json:"origin_time"`
	LedgerUID         string             `json:"series_uid"`
	PrimaryLevel   int64              `json:"primary_level"`
	AgreementOptions *AgreementOptions   `json:"agreement_options,omitempty"`
	Ratifiers      []OriginRatifier `json:"ratifiers,omitempty"`
	ApplicationDigest         cometbytes.HexOctets  `json:"application_digest"`
	ApplicationStatus        json.RawMessage    `json:"application_status,omitempty"`
}

//
func (generatePaper *OriginPaper) PersistAs(entry string) error {
	generatePaperOctets, err := cometjson.SerializeIndent(generatePaper, "REDACTED", "REDACTED")
	if err != nil {
		return err
	}
	return cometos.RecordEntry(entry, generatePaperOctets, 0o644)
}

//
func (generatePaper *OriginPaper) RatifierDigest() []byte {
	values := make([]*Ratifier, len(generatePaper.Ratifiers))
	for i, v := range generatePaper.Ratifiers {
		values[i] = NewRatifier(v.PublicKey, v.Energy)
	}
	rset := NewRatifierCollection(values)
	return rset.Digest()
}

//
//
func (generatePaper *OriginPaper) CertifyAndFinished() error {
	if generatePaper.LedgerUID == "REDACTED" {
		return errors.New("REDACTED")
	}
	if len(generatePaper.LedgerUID) > MaximumSeriesUIDSize {
		return fmt.Errorf("REDACTED", MaximumSeriesUIDSize)
	}
	if generatePaper.PrimaryLevel < 0 {
		return fmt.Errorf("REDACTED", generatePaper.PrimaryLevel)
	}
	if generatePaper.PrimaryLevel == 0 {
		generatePaper.PrimaryLevel = 1
	}

	if generatePaper.AgreementOptions == nil {
		generatePaper.AgreementOptions = StandardAgreementOptions()
	} else if err := generatePaper.AgreementOptions.CertifySimple(); err != nil {
		return err
	}

	for i, v := range generatePaper.Ratifiers {
		if v.Energy == 0 {
			return fmt.Errorf("REDACTED", v)
		}
		if len(v.Location) > 0 && !bytes.Equal(v.PublicKey.Location(), v.Location) {
			return fmt.Errorf("REDACTED", v, v.PublicKey.Location())
		}
		if len(v.Location) == 0 {
			generatePaper.Ratifiers[i].Location = v.PublicKey.Location()
		}
	}

	if generatePaper.OriginMoment.IsZero() {
		generatePaper.OriginMoment = engineclock.Now()
	}

	return nil
}

//
//

//
func OriginPaperFromJSON(jsonBinary []byte) (*OriginPaper, error) {
	generatePaper := OriginPaper{}
	err := cometjson.Unserialize(jsonBinary, &generatePaper)
	if err != nil {
		return nil, err
	}

	if err := generatePaper.CertifyAndFinished(); err != nil {
		return nil, err
	}

	return &generatePaper, err
}

//
func OriginPaperFromEntry(generatePaperEntry string) (*OriginPaper, error) {
	jsonBinary, err := os.ReadFile(generatePaperEntry)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	generatePaper, err := OriginPaperFromJSON(jsonBinary)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", generatePaperEntry, err)
	}
	return generatePaper, nil
}
