package primary

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/app"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

//
type Settings struct {
	SuccessionUUID                    string                      `toml:"succession_uuid"`
	Overhear                     string                      `toml:"overhear"`
	Scheme                   string                      `toml:"scheme"`
	Dir                        string                      `toml:"dir"`
	Style                       string                      `toml:"style"`
	EndureDuration            uint64                      `toml:"endure_duration"`
	ImageDuration           uint64                      `toml:"image_duration"`
	PreserveLedgers               uint64                      `toml:"preserve_ledgers"`
	AssessorRevisions           map[string]map[string]uint8 `toml:"assessor_revise"`
	PrivateItemDaemon              string                      `toml:"privatevalue_daemon"`
	PrivateItemToken                 string                      `toml:"privatevalue_token"`
	PrivateItemStatus               string                      `toml:"privatevalue_status"`
	TokenKind                    string                      `toml:"token_kind"`
	BallotAdditionsActivateAltitude int64                       `toml:"ballot_additions_activate_altitude"`
	BallotAdditionsReviseAltitude int64                       `toml:"ballot_additions_revise_altitude"`
	BallotAdditionExtent          uint                        `toml:"ballot_addition_extent"`
	ApplicationFlankTxpool             bool                        `toml:"application_flank_txpool"`
}

//
func (cfg *Settings) App() *app.Settings {
	return &app.Settings{
		Dir:                        cfg.Dir,
		ImageDuration:           cfg.ImageDuration,
		PreserveLedgers:               cfg.PreserveLedgers,
		TokenKind:                    cfg.TokenKind,
		AssessorRevisions:           cfg.AssessorRevisions,
		EndureDuration:            cfg.EndureDuration,
		BallotAdditionsActivateAltitude: cfg.BallotAdditionsActivateAltitude,
		BallotAdditionsReviseAltitude: cfg.BallotAdditionsReviseAltitude,
		BallotAdditionExtent:          cfg.BallotAdditionExtent,
		ApplicationFlankTxpool:             cfg.ApplicationFlankTxpool,
	}
}

//
func FetchSettings(record string) (*Settings, error) {
	cfg := &Settings{
		Overhear:          "REDACTED",
		Scheme:        "REDACTED",
		EndureDuration: 1,
	}
	_, err := toml.DecodeFile(record, &cfg)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", record, err)
	}
	return cfg, cfg.Certify()
}

//
//
//

func (cfg Settings) Certify() error {
	switch {
	case cfg.SuccessionUUID == "REDACTED":
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	case cfg.Overhear == "REDACTED" && cfg.Scheme != "REDACTED" && cfg.Scheme != "REDACTED":
		return strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	default:
		return nil
	}
}
