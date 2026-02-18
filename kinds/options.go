package kinds

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/vault/bls12381"
	"github.com/valkyrieworks/vault/ed25519"
	"github.com/valkyrieworks/vault/secp256k1"
	"github.com/valkyrieworks/vault/comethash"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
)

const (
	//
	MaximumLedgerVolumeOctets = 104857600 //

	//
	LedgerSegmentVolumeOctets uint32 = 65536 //

	//
	MaximumLedgerSegmentsTally = (MaximumLedgerVolumeOctets / LedgerSegmentVolumeOctets) + 1

	IfacePublicKeyKindEd25519   = ed25519.KeyKind
	IfacePublicKeyKindSecp256k1 = secp256k1.KeyKind
	IfacePublicKeyKindBls12381  = bls12381.KeyKind
)

var IfacePublicKeyKindsToLabels = map[string]string{
	IfacePublicKeyKindEd25519:   ed25519.PublicKeyLabel,
	IfacePublicKeyKindSecp256k1: secp256k1.PublicKeyLabel,
}

func init() {
	if bls12381.Activated {
		IfacePublicKeyKindsToLabels[IfacePublicKeyKindBls12381] = bls12381.PublicKeyLabel
	}
}

//
//
type AgreementOptions struct {
	Ledger     LedgerOptions     `json:"ledger"`
	Proof  ProofOptions  `json:"proof"`
	Ratifier RatifierOptions `json:"ratifier"`
	Release   ReleaseOptions   `json:"release"`
	Iface      IfaceOptions      `json:"iface"`
}

//
//
type LedgerOptions struct {
	MaximumOctets int64 `json:"maximum_octets"`
	MaximumFuel   int64 `json:"maximum_fuel"`
}

//
type ProofOptions struct {
	MaximumDurationCountLedgers int64         `json:"maximum_duration_count_ledgers"` //
	MaximumDurationPeriod  time.Duration `json:"maximum_duration_period"`
	MaximumOctets        int64         `json:"maximum_octets"`
}

//
//
type RatifierOptions struct {
	PublicKeyKinds []string `json:"public_key_kinds"`
}

type ReleaseOptions struct {
	App uint64 `json:"app"`
}

//
//
type IfaceOptions struct {
	BallotPluginsActivateLevel int64 `json:"ballot_plugins_activate_level"`
}

//
//
func (a IfaceOptions) BallotPluginsActivated(h int64) bool {
	if h < 1 {
		panic(fmt.Errorf("REDACTED", h))
	}
	if a.BallotPluginsActivateLevel == 0 {
		return false
	}
	return a.BallotPluginsActivateLevel <= h
}

//
func StandardAgreementOptions() *AgreementOptions {
	return &AgreementOptions{
		Ledger:     StandardLedgerOptions(),
		Proof:  StandardProofOptions(),
		Ratifier: StandardRatifierOptions(),
		Release:   StandardReleaseOptions(),
		Iface:      StandardIfaceOptions(),
	}
}

//
func StandardLedgerOptions() LedgerOptions {
	return LedgerOptions{
		MaximumOctets: 22020096, //
		MaximumFuel:   -1,
	}
}

//
func StandardProofOptions() ProofOptions {
	return ProofOptions{
		MaximumDurationCountLedgers: 100000, //
		MaximumDurationPeriod:  48 * time.Hour,
		MaximumOctets:        1048576, //
	}
}

//
//
func StandardRatifierOptions() RatifierOptions {
	return RatifierOptions{
		PublicKeyKinds: []string{IfacePublicKeyKindEd25519},
	}
}

func StandardReleaseOptions() ReleaseOptions {
	return ReleaseOptions{
		App: 0,
	}
}

func StandardIfaceOptions() IfaceOptions {
	return IfaceOptions{
		//
		BallotPluginsActivateLevel: 0,
	}
}

func IsSoundPublickeyKind(options RatifierOptions, publickeyKind string) bool {
	for i := 0; i < len(options.PublicKeyKinds); i++ {
		if options.PublicKeyKinds[i] == publickeyKind {
			return true
		}
	}
	return false
}

//
//
func (options AgreementOptions) CertifySimple() error {
	if options.Ledger.MaximumOctets == 0 {
		return fmt.Errorf("REDACTED")
	}
	if options.Ledger.MaximumOctets < -1 {
		return fmt.Errorf("REDACTED",

			options.Ledger.MaximumOctets)
	}
	if options.Ledger.MaximumOctets > MaximumLedgerVolumeOctets {
		return fmt.Errorf("REDACTED",
			options.Ledger.MaximumOctets, MaximumLedgerVolumeOctets)
	}

	if options.Ledger.MaximumFuel < -1 {
		return fmt.Errorf("REDACTED",
			options.Ledger.MaximumFuel)
	}

	if options.Proof.MaximumDurationCountLedgers <= 0 {
		return fmt.Errorf("REDACTED",
			options.Proof.MaximumDurationCountLedgers)
	}

	if options.Proof.MaximumDurationPeriod <= 0 {
		return fmt.Errorf("REDACTED",
			options.Proof.MaximumDurationPeriod)
	}

	maximumOctets := options.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(MaximumLedgerVolumeOctets)
	}
	if options.Proof.MaximumOctets > maximumOctets {
		return fmt.Errorf("REDACTED",
			options.Proof.MaximumOctets, maximumOctets)
	}

	if options.Proof.MaximumOctets < 0 {
		return fmt.Errorf("REDACTED",
			options.Proof.MaximumOctets)
	}

	if options.Iface.BallotPluginsActivateLevel < 0 {
		return fmt.Errorf("REDACTED", options.Iface.BallotPluginsActivateLevel)
	}

	if len(options.Ratifier.PublicKeyKinds) == 0 {
		return errors.New("REDACTED")
	}

	//
	for i := 0; i < len(options.Ratifier.PublicKeyKinds); i++ {
		keyKind := options.Ratifier.PublicKeyKinds[i]
		if _, ok := IfacePublicKeyKindsToLabels[keyKind]; !ok {
			return fmt.Errorf("REDACTED",
				i, keyKind)
		}
	}

	return nil
}

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
//
//
func (options AgreementOptions) CertifyModify(refreshed *engineproto.AgreementOptions, h int64) error {
	//
	if refreshed == nil || refreshed.Iface == nil {
		return nil
	}
	//
	if refreshed.Iface.BallotPluginsActivateLevel < 0 {
		return errors.New("REDACTED")
	}
	//
	if options.Iface.BallotPluginsActivateLevel <= 0 && refreshed.Iface.BallotPluginsActivateLevel == 0 {
		return nil
	}
	//
	if options.Iface.BallotPluginsActivateLevel == refreshed.Iface.BallotPluginsActivateLevel {
		return nil
	}
	//
	if options.Iface.BallotPluginsActivateLevel > 0 && refreshed.Iface.BallotPluginsActivateLevel == 0 {
		//
		if options.Iface.BallotPluginsActivateLevel <= h {
			return fmt.Errorf("REDACTED"+
				"REDACTED",
				options.Iface.BallotPluginsActivateLevel, h)
		}
		//
		return nil
	}
	//
	if refreshed.Iface.BallotPluginsActivateLevel <= h {
		return fmt.Errorf("REDACTED"+
			"REDACTED",
			refreshed.Iface.BallotPluginsActivateLevel, h)
	}
	//
	if options.Iface.BallotPluginsActivateLevel <= 0 {
		return nil
	}
	//
	if options.Iface.BallotPluginsActivateLevel <= h {
		return fmt.Errorf("REDACTED"+
			"REDACTED",
			options.Iface.BallotPluginsActivateLevel, h)
	}
	//
	return nil
}

//
//
//
//
func (options AgreementOptions) Digest() []byte {
	digester := comethash.New()

	hp := engineproto.DigestedOptions{
		LedgerMaximumOctets: options.Ledger.MaximumOctets,
		LedgerMaximumFuel:   options.Ledger.MaximumFuel,
	}

	bz, err := hp.Serialize()
	if err != nil {
		panic(err)
	}

	_, err = digester.Write(bz)
	if err != nil {
		panic(err)
	}
	return digester.Sum(nil)
}

//
//
func (options AgreementOptions) Modify(options2 *engineproto.AgreementOptions) AgreementOptions {
	res := options //

	if options2 == nil {
		return res
	}

	//
	if options2.Ledger != nil {
		res.Ledger.MaximumOctets = options2.Ledger.MaximumOctets
		res.Ledger.MaximumFuel = options2.Ledger.MaximumFuel
	}
	if options2.Proof != nil {
		res.Proof.MaximumDurationCountLedgers = options2.Proof.MaximumDurationCountLedgers
		res.Proof.MaximumDurationPeriod = options2.Proof.MaximumDurationPeriod
		res.Proof.MaximumOctets = options2.Proof.MaximumOctets
	}
	if options2.Ratifier != nil {
		//
		//
		res.Ratifier.PublicKeyKinds = append([]string{}, options2.Ratifier.PublicKeyKinds...)
	}
	if options2.Release != nil {
		res.Release.App = options2.Release.App
	}
	if options2.Iface != nil {
		res.Iface.BallotPluginsActivateLevel = options2.Iface.FetchBallotPluginsActivateLevel()
	}
	return res
}

func (options *AgreementOptions) ToSchema() engineproto.AgreementOptions {
	return engineproto.AgreementOptions{
		Ledger: &engineproto.LedgerOptions{
			MaximumOctets: options.Ledger.MaximumOctets,
			MaximumFuel:   options.Ledger.MaximumFuel,
		},
		Proof: &engineproto.ProofOptions{
			MaximumDurationCountLedgers: options.Proof.MaximumDurationCountLedgers,
			MaximumDurationPeriod:  options.Proof.MaximumDurationPeriod,
			MaximumOctets:        options.Proof.MaximumOctets,
		},
		Ratifier: &engineproto.RatifierOptions{
			PublicKeyKinds: options.Ratifier.PublicKeyKinds,
		},
		Release: &engineproto.ReleaseOptions{
			App: options.Release.App,
		},
		Iface: &engineproto.IfaceOptions{
			BallotPluginsActivateLevel: options.Iface.BallotPluginsActivateLevel,
		},
	}
}

func AgreementOptionsFromSchema(pbOptions engineproto.AgreementOptions) AgreementOptions {
	c := AgreementOptions{
		Ledger: LedgerOptions{
			MaximumOctets: pbOptions.Ledger.MaximumOctets,
			MaximumFuel:   pbOptions.Ledger.MaximumFuel,
		},
		Proof: ProofOptions{
			MaximumDurationCountLedgers: pbOptions.Proof.MaximumDurationCountLedgers,
			MaximumDurationPeriod:  pbOptions.Proof.MaximumDurationPeriod,
			MaximumOctets:        pbOptions.Proof.MaximumOctets,
		},
		Ratifier: RatifierOptions{
			PublicKeyKinds: pbOptions.Ratifier.PublicKeyKinds,
		},
		Release: ReleaseOptions{
			App: pbOptions.Release.App,
		},
	}
	if pbOptions.Iface != nil {
		c.Iface.BallotPluginsActivateLevel = pbOptions.Iface.FetchBallotPluginsActivateLevel()
	}
	return c
}
