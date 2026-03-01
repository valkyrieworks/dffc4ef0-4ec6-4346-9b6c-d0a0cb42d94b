package directives

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/privatevalue"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	committime "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/moment"
)

var (
	nthAssessors    int
	nthUnAssessors int
	primaryAltitude  int64
	settingsRecord     string
	emissionPath      string
	peerPathHeading  string

	inhabitEnduringNodes bool
	machinenameHeading          string
	machinenameEnding          string
	launchingINETLocation       string
	machinenames               []string
	peer2peerChannel                 int
	arbitraryPseudonyms          bool
)

const (
	peerPathMode = 0o755
)

func initialize() {
	SimnetRecordsDirective.Flags().IntVar(&nthAssessors, "REDACTED", 4,
		"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&settingsRecord, "REDACTED", "REDACTED",
		"REDACTED")
	SimnetRecordsDirective.Flags().IntVar(&nthUnAssessors, "REDACTED", 0,
		"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&emissionPath, "REDACTED", "REDACTED",
		"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&peerPathHeading, "REDACTED", "REDACTED",
		"REDACTED")
	SimnetRecordsDirective.Flags().Int64Var(&primaryAltitude, "REDACTED", 0,
		"REDACTED")

	SimnetRecordsDirective.Flags().BoolVar(&inhabitEnduringNodes, "REDACTED", true,
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&machinenameHeading, "REDACTED", "REDACTED",
		"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&machinenameEnding, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	SimnetRecordsDirective.Flags().StringVar(&launchingINETLocation, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	SimnetRecordsDirective.Flags().StringArrayVar(&machinenames, "REDACTED", []string{},
		"REDACTED")
	SimnetRecordsDirective.Flags().IntVar(&peer2peerChannel, "REDACTED", 26656,
		"REDACTED")
	SimnetRecordsDirective.Flags().BoolVar(&arbitraryPseudonyms, "REDACTED", false,
		"REDACTED")
}

//
var SimnetRecordsDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDh
REDACTED.

REDACTED.

REDACTED.

REDACTED:

REDACTED2
REDACTED`,
	RunE: simnetRecords,
}

func simnetRecords(*cobra.Command, []string) error {
	if len(machinenames) > 0 && len(machinenames) != (nthAssessors+nthUnAssessors) {
		return fmt.Errorf(
			"REDACTED",
			nthAssessors+nthUnAssessors,
		)
	}

	settings := cfg.FallbackSettings()

	//
	if settingsRecord != "REDACTED" {
		viper.SetConfigFile(settingsRecord)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		if err := viper.Unmarshal(settings); err != nil {
			return err
		}
		if err := settings.CertifyFundamental(); err != nil {
			return err
		}
	}

	produceValues := make([]kinds.OriginAssessor, nthAssessors)

	for i := 0; i < nthAssessors; i++ {
		peerPathAlias := fmt.Sprintf("REDACTED", peerPathHeading, i)
		peerPath := filepath.Join(emissionPath, peerPathAlias)
		settings.AssignOrigin(peerPath)

		err := os.MkdirAll(filepath.Join(peerPath, "REDACTED"), peerPathMode)
		if err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}
		err = os.MkdirAll(filepath.Join(peerPath, "REDACTED"), peerPathMode)
		if err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}

		if err := initializeRecordsUsingSettings(settings); err != nil {
			return err
		}

		prvTokenRecord := filepath.Join(peerPath, settings.PrivateAssessorToken)
		prvStatusRecord := filepath.Join(peerPath, settings.PrivateAssessorStatus)
		pv := privatevalue.FetchRecordPRV(prvTokenRecord, prvStatusRecord)

		publicToken, err := pv.ObtainPublicToken()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		produceValues[i] = kinds.OriginAssessor{
			Location: publicToken.Location(),
			PublicToken:  publicToken,
			Potency:   1,
			Alias:    peerPathAlias,
		}
	}

	for i := 0; i < nthUnAssessors; i++ {
		peerPath := filepath.Join(emissionPath, fmt.Sprintf("REDACTED", peerPathHeading, i+nthAssessors))
		settings.AssignOrigin(peerPath)

		err := os.MkdirAll(filepath.Join(peerPath, "REDACTED"), peerPathMode)
		if err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}

		err = os.MkdirAll(filepath.Join(peerPath, "REDACTED"), peerPathMode)
		if err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}

		if err := initializeRecordsUsingSettings(settings); err != nil {
			return err
		}
	}

	//
	producePaper := &kinds.OriginPaper{
		SuccessionUUID:         "REDACTED" + commitrand.Str(6),
		AgreementSettings: kinds.FallbackAgreementSettings(),
		OriginMoment:     committime.Now(),
		PrimaryAltitude:   primaryAltitude,
		Assessors:      produceValues,
	}

	//
	for i := 0; i < nthAssessors+nthUnAssessors; i++ {
		peerPath := filepath.Join(emissionPath, fmt.Sprintf("REDACTED", peerPathHeading, i))
		if err := producePaper.PersistLike(filepath.Join(peerPath, settings.Inauguration)); err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}
	}

	//
	var (
		enduringNodes string
		err             error
	)
	if inhabitEnduringNodes {
		enduringNodes, err = enduringNodesText(settings)
		if err != nil {
			_ = os.RemoveAll(emissionPath)
			return err
		}
	}

	//
	for i := 0; i < nthAssessors+nthUnAssessors; i++ {
		peerPath := filepath.Join(emissionPath, fmt.Sprintf("REDACTED", peerPathHeading, i))
		settings.AssignOrigin(peerPath)
		settings.P2P.LocationRegisterPrecise = false
		settings.P2P.PermitReplicatedINET = true
		if inhabitEnduringNodes {
			settings.P2P.EnduringNodes = enduringNodes
		}
		settings.Pseudonym = pseudonym(i)

		cfg.PersistSettingsRecord(filepath.Join(peerPath, "REDACTED", "REDACTED"), settings)
	}

	fmt.Printf("REDACTED", nthAssessors+nthUnAssessors)
	return nil
}

func machinenameEitherINET(i int) string {
	if len(machinenames) > 0 && i < len(machinenames) {
		return machinenames[i]
	}
	if launchingINETLocation == "REDACTED" {
		return fmt.Sprintf("REDACTED", machinenameHeading, i, machinenameEnding)
	}
	ip := net.ParseIP(launchingINETLocation)
	ip = ip.To4()
	if ip == nil {
		fmt.Printf("REDACTED", launchingINETLocation)
		os.Exit(1)
	}

	for range i {
		increaseINET(ip)
	}
	return ip.String()
}

//
func increaseINET(ip net.IP) {
	//
	for idx := len(ip) - 1; idx >= 0; idx-- {
		ip[idx]++

		if ip[idx] != 0 {
			break
		}
	}
}

func enduringNodesText(settings *cfg.Settings) (string, error) {
	enduringNodes := make([]string, nthAssessors+nthUnAssessors)
	for i := 0; i < nthAssessors+nthUnAssessors; i++ {
		peerPath := filepath.Join(emissionPath, fmt.Sprintf("REDACTED", peerPathHeading, i))
		settings.AssignOrigin(peerPath)
		peerToken, err := p2p.FetchPeerToken(settings.PeerTokenRecord())
		if err != nil {
			return "REDACTED", err
		}
		enduringNodes[i] = p2p.UUIDLocationText(peerToken.ID(), fmt.Sprintf("REDACTED", machinenameEitherINET(i), peer2peerChannel))
	}
	return strings.Join(enduringNodes, "REDACTED"), nil
}

func pseudonym(i int) string {
	if arbitraryPseudonyms {
		return arbitraryPseudonym()
	}
	if len(machinenames) > 0 && i < len(machinenames) {
		return machinenames[i]
	}
	if launchingINETLocation == "REDACTED" {
		return fmt.Sprintf("REDACTED", machinenameHeading, i, machinenameEnding)
	}
	return arbitraryPseudonym()
}

func arbitraryPseudonym() string {
	return octets.HexadecimalOctets(commitrand.Octets(8)).Text()
}
