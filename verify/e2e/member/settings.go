package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/valkyrieworks/verify/e2e/app"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

//
type Settings struct {
	LedgerUID                    string                      `toml:"series_uid"`
	Observe                     string                      `toml:"observe"`
	Protocol                   string                      `toml:"protocol"`
	Dir                        string                      `toml:"dir"`
	Style                       string                      `toml:"style"`
	EndureCadence            uint64                      `toml:"endure_cadence"`
	MirrorCadence           uint64                      `toml:"mirror_cadence"`
	PreserveLedgers               uint64                      `toml:"preserve_ledgers"`
	RatifierRefreshes           map[string]map[string]uint8 `toml:"ratifier_modify"`
	PrivateValueHost              string                      `toml:"privatekey_host"`
	PrivateValueKey                 string                      `toml:"privatekey_key"`
	PrivateValueStatus               string                      `toml:"privatekey_status"`
	KeyKind                    string                      `toml:"key_kind"`
	BallotPluginsActivateLevel int64                       `toml:"ballot_plugins_activate_level"`
	BallotPluginsModifyLevel int64                       `toml:"ballot_plugins_modify_level"`
	BallotAdditionVolume          uint                        `toml:"ballot_addition_volume"`
	ApplicationFlankTxpool             bool                        `toml:"application_flank_txpool"`
}

//
func (cfg *Settings) App() *app.Settings {
	return &app.Settings{
		Dir:                        cfg.Dir,
		MirrorCadence:           cfg.MirrorCadence,
		PreserveLedgers:               cfg.PreserveLedgers,
		KeyKind:                    cfg.KeyKind,
		RatifierRefreshes:           cfg.RatifierRefreshes,
		EndureCadence:            cfg.EndureCadence,
		BallotPluginsActivateLevel: cfg.BallotPluginsActivateLevel,
		BallotPluginsModifyLevel: cfg.BallotPluginsModifyLevel,
		BallotAdditionVolume:          cfg.BallotAdditionVolume,
		ApplicationFlankTxpool:             cfg.ApplicationFlankTxpool,
	}
}

//
func ImportSettings(entry string) (*Settings, error) {
	cfg := &Settings{
		Observe:          "REDACTED",
		Protocol:        "REDACTED",
		EndureCadence: 1,
	}
	_, err := toml.DecodeFile(entry, &cfg)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", entry, err)
	}
	return cfg, cfg.Certify()
}

//
//
//

func (cfg Settings) Certify() error {
	switch {
	case cfg.LedgerUID == "REDACTED":
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	case cfg.Observe == "REDACTED" && cfg.Protocol != "REDACTED" && cfg.Protocol != "REDACTED":
		return cometfaults.ErrMandatoryField{Field: "REDACTED"}
	default:
		return nil
	}
}
