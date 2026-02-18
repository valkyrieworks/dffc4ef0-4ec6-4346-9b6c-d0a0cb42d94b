package directives

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/octets"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/p2p"
	"github.com/valkyrieworks/privatekey"
	"github.com/valkyrieworks/kinds"
	engineclock "github.com/valkyrieworks/kinds/moment"
)

var (
	nRatifiers    int
	nNotRatifiers int
	primaryLevel  int64
	settingsEntry     string
	resultFolder      string
	memberFolderHeading  string

	fillDurableNodes bool
	hostlabelHeading          string
	hostlabelExtension          string
	launchingIPLocation       string
	hostlabels               []string
	p2pPort                 int
	arbitraryMonikers          bool
)

const (
	memberFolderMode = 0o755
)

func init() {
	VerifychainEntriesCommand.Flags().IntVar(&nRatifiers, "REDACTED", 4,
		"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&settingsEntry, "REDACTED", "REDACTED",
		"REDACTED")
	VerifychainEntriesCommand.Flags().IntVar(&nNotRatifiers, "REDACTED", 0,
		"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&resultFolder, "REDACTED", "REDACTED",
		"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&memberFolderHeading, "REDACTED", "REDACTED",
		"REDACTED")
	VerifychainEntriesCommand.Flags().Int64Var(&primaryLevel, "REDACTED", 0,
		"REDACTED")

	VerifychainEntriesCommand.Flags().BoolVar(&fillDurableNodes, "REDACTED", true,
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&hostlabelHeading, "REDACTED", "REDACTED",
		"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&hostlabelExtension, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	VerifychainEntriesCommand.Flags().StringVar(&launchingIPLocation, "REDACTED", "REDACTED",
		"REDACTED"+
			"REDACTED"+
			"REDACTED")
	VerifychainEntriesCommand.Flags().StringArrayVar(&hostlabels, "REDACTED", []string{},
		"REDACTED")
	VerifychainEntriesCommand.Flags().IntVar(&p2pPort, "REDACTED", 26656,
		"REDACTED")
	VerifychainEntriesCommand.Flags().BoolVar(&arbitraryMonikers, "REDACTED", false,
		"REDACTED")
}

//
var VerifychainEntriesCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDh
REDACTED.

REDACTED.

REDACTED.

REDACTED:

REDACTED2
REDACTED`,
	RunE: verifychainEntries,
}

func verifychainEntries(*cobra.Command, []string) error {
	if len(hostlabels) > 0 && len(hostlabels) != (nRatifiers+nNotRatifiers) {
		return fmt.Errorf(
			"REDACTED",
			nRatifiers+nNotRatifiers,
		)
	}

	settings := cfg.StandardSettings()

	//
	if settingsEntry != "REDACTED" {
		viper.SetConfigFile(settingsEntry)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		if err := viper.Unmarshal(settings); err != nil {
			return err
		}
		if err := settings.CertifySimple(); err != nil {
			return err
		}
	}

	generateValues := make([]kinds.OriginRatifier, nRatifiers)

	for i := 0; i < nRatifiers; i++ {
		memberFolderLabel := fmt.Sprintf("REDACTED", memberFolderHeading, i)
		memberFolder := filepath.Join(resultFolder, memberFolderLabel)
		settings.AssignOrigin(memberFolder)

		err := os.MkdirAll(filepath.Join(memberFolder, "REDACTED"), memberFolderMode)
		if err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}
		err = os.MkdirAll(filepath.Join(memberFolder, "REDACTED"), memberFolderMode)
		if err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}

		if err := initEntriesWithSettings(settings); err != nil {
			return err
		}

		privatekeyKeyEntry := filepath.Join(memberFolder, settings.PrivateRatifierKey)
		privatekeyStatusEntry := filepath.Join(memberFolder, settings.PrivateRatifierStatus)
		pv := privatekey.ImportEntryPrivatekey(privatekeyKeyEntry, privatekeyStatusEntry)

		publicKey, err := pv.FetchPublicKey()
		if err != nil {
			return fmt.Errorf("REDACTED", err)
		}
		generateValues[i] = kinds.OriginRatifier{
			Location: publicKey.Location(),
			PublicKey:  publicKey,
			Energy:   1,
			Label:    memberFolderLabel,
		}
	}

	for i := 0; i < nNotRatifiers; i++ {
		memberFolder := filepath.Join(resultFolder, fmt.Sprintf("REDACTED", memberFolderHeading, i+nRatifiers))
		settings.AssignOrigin(memberFolder)

		err := os.MkdirAll(filepath.Join(memberFolder, "REDACTED"), memberFolderMode)
		if err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}

		err = os.MkdirAll(filepath.Join(memberFolder, "REDACTED"), memberFolderMode)
		if err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}

		if err := initEntriesWithSettings(settings); err != nil {
			return err
		}
	}

	//
	generatePaper := &kinds.OriginPaper{
		LedgerUID:         "REDACTED" + engineseed.Str(6),
		AgreementOptions: kinds.StandardAgreementOptions(),
		OriginMoment:     engineclock.Now(),
		PrimaryLevel:   primaryLevel,
		Ratifiers:      generateValues,
	}

	//
	for i := 0; i < nRatifiers+nNotRatifiers; i++ {
		memberFolder := filepath.Join(resultFolder, fmt.Sprintf("REDACTED", memberFolderHeading, i))
		if err := generatePaper.PersistAs(filepath.Join(memberFolder, settings.Origin)); err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}
	}

	//
	var (
		durableNodes string
		err             error
	)
	if fillDurableNodes {
		durableNodes, err = durableNodesString(settings)
		if err != nil {
			_ = os.RemoveAll(resultFolder)
			return err
		}
	}

	//
	for i := 0; i < nRatifiers+nNotRatifiers; i++ {
		memberFolder := filepath.Join(resultFolder, fmt.Sprintf("REDACTED", memberFolderHeading, i))
		settings.AssignOrigin(memberFolder)
		settings.P2P.AddressLedgerPrecise = false
		settings.P2P.PermitReplicatedIP = true
		if fillDurableNodes {
			settings.P2P.DurableNodes = durableNodes
		}
		settings.Moniker = moniker(i)

		cfg.RecordSettingsEntry(filepath.Join(memberFolder, "REDACTED", "REDACTED"), settings)
	}

	fmt.Printf("REDACTED", nRatifiers+nNotRatifiers)
	return nil
}

func hostlabelOrIP(i int) string {
	if len(hostlabels) > 0 && i < len(hostlabels) {
		return hostlabels[i]
	}
	if launchingIPLocation == "REDACTED" {
		return fmt.Sprintf("REDACTED", hostlabelHeading, i, hostlabelExtension)
	}
	ip := net.ParseIP(launchingIPLocation)
	ip = ip.To4()
	if ip == nil {
		fmt.Printf("REDACTED", launchingIPLocation)
		os.Exit(1)
	}

	for range i {
		increaseIP(ip)
	}
	return ip.String()
}

//
func increaseIP(ip net.IP) {
	//
	for idx := len(ip) - 1; idx >= 0; idx-- {
		ip[idx]++

		if ip[idx] != 0 {
			break
		}
	}
}

func durableNodesString(settings *cfg.Settings) (string, error) {
	durableNodes := make([]string, nRatifiers+nNotRatifiers)
	for i := 0; i < nRatifiers+nNotRatifiers; i++ {
		memberFolder := filepath.Join(resultFolder, fmt.Sprintf("REDACTED", memberFolderHeading, i))
		settings.AssignOrigin(memberFolder)
		memberKey, err := p2p.ImportMemberKey(settings.MemberKeyEntry())
		if err != nil {
			return "REDACTED", err
		}
		durableNodes[i] = p2p.UIDLocationString(memberKey.ID(), fmt.Sprintf("REDACTED", hostlabelOrIP(i), p2pPort))
	}
	return strings.Join(durableNodes, "REDACTED"), nil
}

func moniker(i int) string {
	if arbitraryMonikers {
		return arbitraryMoniker()
	}
	if len(hostlabels) > 0 && i < len(hostlabels) {
		return hostlabels[i]
	}
	if launchingIPLocation == "REDACTED" {
		return fmt.Sprintf("REDACTED", hostlabelHeading, i, hostlabelExtension)
	}
	return arbitraryMoniker()
}

func arbitraryMoniker() string {
	return octets.HexOctets(engineseed.Octets(8)).String()
}
