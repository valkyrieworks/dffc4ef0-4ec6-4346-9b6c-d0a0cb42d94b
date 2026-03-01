package kinds

import (
	"errors"
	"fmt"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/signature381"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/ellipticp256"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
)

const (
	//
	MaximumLedgerExtentOctets = 104857600 //

	//
	LedgerFragmentExtentOctets uint32 = 65536 //

	//
	MaximumLedgerFragmentsTally = (MaximumLedgerExtentOctets / LedgerFragmentExtentOctets) + 1

	IfacePublicTokenKindEdwards25519   = edwards25519.TokenKind
	IfacePublicTokenKindEllipticp256 = ellipticp256.TokenKind
	IfacePublicTokenKindSignature381  = signature381.TokenKind
)

var IfacePublicTokenKindsTowardIdentifiers = map[string]string{
	IfacePublicTokenKindEdwards25519:   edwards25519.PublicTokenAlias,
	IfacePublicTokenKindEllipticp256: ellipticp256.PublicTokenAlias,
}

func initialize() {
	if signature381.Activated {
		IfacePublicTokenKindsTowardIdentifiers[IfacePublicTokenKindSignature381] = signature381.PublicTokenAlias
	}
}

//
//
type AgreementSettings struct {
	Ledger     LedgerParameters     `json:"ledger"`
	Proof  ProofParameters  `json:"proof"`
	Assessor AssessorParameters `json:"assessor"`
	Edition   EditionParameters   `json:"edition"`
	Iface      IfaceParameters      `json:"iface"`
}

//
//
type LedgerParameters struct {
	MaximumOctets int64 `json:"maximum_octets"`
	MaximumFuel   int64 `json:"maximum_fuel"`
}

//
type ProofParameters struct {
	MaximumLifespanCountLedgers int64         `json:"maximum_lifespan_count_ledgers"` //
	MaximumLifespanInterval  time.Duration `json:"maximum_lifespan_interval"`
	MaximumOctets        int64         `json:"maximum_octets"`
}

//
//
type AssessorParameters struct {
	PublicTokenKinds []string `json:"public_token_kinds"`
}

type EditionParameters struct {
	App uint64 `json:"app"`
}

//
//
type IfaceParameters struct {
	BallotAdditionsActivateAltitude int64 `json:"ballot_additions_activate_altitude"`
}

//
//
func (a IfaceParameters) BallotAdditionsActivated(h int64) bool {
	if h < 1 {
		panic(fmt.Errorf("REDACTED", h))
	}
	if a.BallotAdditionsActivateAltitude == 0 {
		return false
	}
	return a.BallotAdditionsActivateAltitude <= h
}

//
func FallbackAgreementSettings() *AgreementSettings {
	return &AgreementSettings{
		Ledger:     FallbackLedgerParameters(),
		Proof:  FallbackProofParameters(),
		Assessor: FallbackAssessorParameters(),
		Edition:   FallbackEditionParameters(),
		Iface:      FallbackIfaceParameters(),
	}
}

//
func FallbackLedgerParameters() LedgerParameters {
	return LedgerParameters{
		MaximumOctets: 22020096, //
		MaximumFuel:   -1,
	}
}

//
func FallbackProofParameters() ProofParameters {
	return ProofParameters{
		MaximumLifespanCountLedgers: 100000, //
		MaximumLifespanInterval:  48 * time.Hour,
		MaximumOctets:        1048576, //
	}
}

//
//
func FallbackAssessorParameters() AssessorParameters {
	return AssessorParameters{
		PublicTokenKinds: []string{IfacePublicTokenKindEdwards25519},
	}
}

func FallbackEditionParameters() EditionParameters {
	return EditionParameters{
		App: 0,
	}
}

func FallbackIfaceParameters() IfaceParameters {
	return IfaceParameters{
		//
		BallotAdditionsActivateAltitude: 0,
	}
}

func EqualsSoundPublickeyKind(parameters AssessorParameters, publickeyKind string) bool {
	for i := 0; i < len(parameters.PublicTokenKinds); i++ {
		if parameters.PublicTokenKinds[i] == publickeyKind {
			return true
		}
	}
	return false
}

//
//
func (parameters AgreementSettings) CertifyFundamental() error {
	if parameters.Ledger.MaximumOctets == 0 {
		return fmt.Errorf("REDACTED")
	}
	if parameters.Ledger.MaximumOctets < -1 {
		return fmt.Errorf("REDACTED",

			parameters.Ledger.MaximumOctets)
	}
	if parameters.Ledger.MaximumOctets > MaximumLedgerExtentOctets {
		return fmt.Errorf("REDACTED",
			parameters.Ledger.MaximumOctets, MaximumLedgerExtentOctets)
	}

	if parameters.Ledger.MaximumFuel < -1 {
		return fmt.Errorf("REDACTED",
			parameters.Ledger.MaximumFuel)
	}

	if parameters.Proof.MaximumLifespanCountLedgers <= 0 {
		return fmt.Errorf("REDACTED",
			parameters.Proof.MaximumLifespanCountLedgers)
	}

	if parameters.Proof.MaximumLifespanInterval <= 0 {
		return fmt.Errorf("REDACTED",
			parameters.Proof.MaximumLifespanInterval)
	}

	maximumOctets := parameters.Ledger.MaximumOctets
	if maximumOctets == -1 {
		maximumOctets = int64(MaximumLedgerExtentOctets)
	}
	if parameters.Proof.MaximumOctets > maximumOctets {
		return fmt.Errorf("REDACTED",
			parameters.Proof.MaximumOctets, maximumOctets)
	}

	if parameters.Proof.MaximumOctets < 0 {
		return fmt.Errorf("REDACTED",
			parameters.Proof.MaximumOctets)
	}

	if parameters.Iface.BallotAdditionsActivateAltitude < 0 {
		return fmt.Errorf("REDACTED", parameters.Iface.BallotAdditionsActivateAltitude)
	}

	if len(parameters.Assessor.PublicTokenKinds) == 0 {
		return errors.New("REDACTED")
	}

	//
	for i := 0; i < len(parameters.Assessor.PublicTokenKinds); i++ {
		tokenKind := parameters.Assessor.PublicTokenKinds[i]
		if _, ok := IfacePublicTokenKindsTowardIdentifiers[tokenKind]; !ok {
			return fmt.Errorf("REDACTED",
				i, tokenKind)
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
func (parameters AgreementSettings) CertifyRevise(modified *commitchema.AgreementSettings, h int64) error {
	//
	if modified == nil || modified.Iface == nil {
		return nil
	}
	//
	if modified.Iface.BallotAdditionsActivateAltitude < 0 {
		return errors.New("REDACTED")
	}
	//
	if parameters.Iface.BallotAdditionsActivateAltitude <= 0 && modified.Iface.BallotAdditionsActivateAltitude == 0 {
		return nil
	}
	//
	if parameters.Iface.BallotAdditionsActivateAltitude == modified.Iface.BallotAdditionsActivateAltitude {
		return nil
	}
	//
	if parameters.Iface.BallotAdditionsActivateAltitude > 0 && modified.Iface.BallotAdditionsActivateAltitude == 0 {
		//
		if parameters.Iface.BallotAdditionsActivateAltitude <= h {
			return fmt.Errorf("REDACTED"+
				"REDACTED",
				parameters.Iface.BallotAdditionsActivateAltitude, h)
		}
		//
		return nil
	}
	//
	if modified.Iface.BallotAdditionsActivateAltitude <= h {
		return fmt.Errorf("REDACTED"+
			"REDACTED",
			modified.Iface.BallotAdditionsActivateAltitude, h)
	}
	//
	if parameters.Iface.BallotAdditionsActivateAltitude <= 0 {
		return nil
	}
	//
	if parameters.Iface.BallotAdditionsActivateAltitude <= h {
		return fmt.Errorf("REDACTED"+
			"REDACTED",
			parameters.Iface.BallotAdditionsActivateAltitude, h)
	}
	//
	return nil
}

//
//
//
//
func (parameters AgreementSettings) Digest() []byte {
	digester := tenderminthash.New()

	hp := commitchema.DigestedParameters{
		LedgerMaximumOctets: parameters.Ledger.MaximumOctets,
		LedgerMaximumFuel:   parameters.Ledger.MaximumFuel,
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
func (parameters AgreementSettings) Revise(parameters2 *commitchema.AgreementSettings) AgreementSettings {
	res := parameters //

	if parameters2 == nil {
		return res
	}

	//
	if parameters2.Ledger != nil {
		res.Ledger.MaximumOctets = parameters2.Ledger.MaximumOctets
		res.Ledger.MaximumFuel = parameters2.Ledger.MaximumFuel
	}
	if parameters2.Proof != nil {
		res.Proof.MaximumLifespanCountLedgers = parameters2.Proof.MaximumLifespanCountLedgers
		res.Proof.MaximumLifespanInterval = parameters2.Proof.MaximumLifespanInterval
		res.Proof.MaximumOctets = parameters2.Proof.MaximumOctets
	}
	if parameters2.Assessor != nil {
		//
		//
		res.Assessor.PublicTokenKinds = append([]string{}, parameters2.Assessor.PublicTokenKinds...)
	}
	if parameters2.Edition != nil {
		res.Edition.App = parameters2.Edition.App
	}
	if parameters2.Iface != nil {
		res.Iface.BallotAdditionsActivateAltitude = parameters2.Iface.ObtainBallotAdditionsActivateAltitude()
	}
	return res
}

func (parameters *AgreementSettings) TowardSchema() commitchema.AgreementSettings {
	return commitchema.AgreementSettings{
		Ledger: &commitchema.LedgerParameters{
			MaximumOctets: parameters.Ledger.MaximumOctets,
			MaximumFuel:   parameters.Ledger.MaximumFuel,
		},
		Proof: &commitchema.ProofParameters{
			MaximumLifespanCountLedgers: parameters.Proof.MaximumLifespanCountLedgers,
			MaximumLifespanInterval:  parameters.Proof.MaximumLifespanInterval,
			MaximumOctets:        parameters.Proof.MaximumOctets,
		},
		Assessor: &commitchema.AssessorParameters{
			PublicTokenKinds: parameters.Assessor.PublicTokenKinds,
		},
		Edition: &commitchema.EditionParameters{
			App: parameters.Edition.App,
		},
		Iface: &commitchema.IfaceParameters{
			BallotAdditionsActivateAltitude: parameters.Iface.BallotAdditionsActivateAltitude,
		},
	}
}

func AgreementParametersOriginatingSchema(bufferParameters commitchema.AgreementSettings) AgreementSettings {
	c := AgreementSettings{
		Ledger: LedgerParameters{
			MaximumOctets: bufferParameters.Ledger.MaximumOctets,
			MaximumFuel:   bufferParameters.Ledger.MaximumFuel,
		},
		Proof: ProofParameters{
			MaximumLifespanCountLedgers: bufferParameters.Proof.MaximumLifespanCountLedgers,
			MaximumLifespanInterval:  bufferParameters.Proof.MaximumLifespanInterval,
			MaximumOctets:        bufferParameters.Proof.MaximumOctets,
		},
		Assessor: AssessorParameters{
			PublicTokenKinds: bufferParameters.Assessor.PublicTokenKinds,
		},
		Edition: EditionParameters{
			App: bufferParameters.Edition.App,
		},
	}
	if bufferParameters.Iface != nil {
		c.Iface.BallotAdditionsActivateAltitude = bufferParameters.Iface.ObtainBallotAdditionsActivateAltitude()
	}
	return c
}
