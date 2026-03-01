package primary

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"

	abcicustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/instance/statedepot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/node"
	daemontest "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/verifies/node"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/edition"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
)

//
var (
	customer abcicustomer.Customer
	tracer log.Tracer
)

//
var (
	//
	markerLocator  string
	markerIface     string
	markerDetailed  bool   //
	markerRecordStratum string //

	//
	markerRoute   string
	markerAltitude int
	markerValidate  bool

	//
	markerEndure string
)

var OriginDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	PersistentPreRunE: func(cmd *cobra.Command, arguments []string) error {
		switch cmd.Use {
		case "REDACTED", "REDACTED", "REDACTED":
			return nil
		}

		if tracer == nil {
			permitStratum, err := log.PermitStratum(markerRecordStratum)
			if err != nil {
				return err
			}
			tracer = log.FreshRefine(log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout)), permitStratum)
		}
		if customer == nil {
			var err error
			customer, err = abcicustomer.FreshCustomer(markerLocator, markerIface, false)
			if err != nil {
				return err
			}
			customer.AssignTracer(tracer.Using("REDACTED", "REDACTED"))
			if err := customer.Initiate(); err != nil {
				return err
			}
		}
		return nil
	},
}

//
type reply struct {
	//
	Data   []byte
	Cipher   uint32
	Details   string
	Log    string
	Condition int32

	Inquire *inquireReply
}

type inquireReply struct {
	Key      []byte
	Datum    []byte
	Altitude   int64
	AttestationActions *security.AttestationActions
}

func Perform() error {
	appendUniversalSwitches()
	appendDirectives()
	return OriginDirective.Execute()
}

func appendUniversalSwitches() {
	OriginDirective.PersistentFlags().StringVarP(&markerLocator,
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED")
	OriginDirective.PersistentFlags().StringVarP(&markerIface, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
	OriginDirective.PersistentFlags().BoolVarP(&markerDetailed,
		"REDACTED",
		"REDACTED",
		false,
		"REDACTED")
	OriginDirective.PersistentFlags().StringVarP(&markerRecordStratum, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
}

func appendInquireSwitches() {
	inquireDirective.PersistentFlags().StringVarP(&markerRoute, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
	inquireDirective.PersistentFlags().IntVarP(&markerAltitude, "REDACTED", "REDACTED", 0, "REDACTED")
	inquireDirective.PersistentFlags().BoolVarP(&markerValidate,
		"REDACTED",
		"REDACTED",
		false,
		"REDACTED")
}

func appendTokvalDepotSwitches() {
	statedepotDirective.PersistentFlags().StringVarP(&markerEndure, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
}

func appendDirectives() {
	OriginDirective.AddCommand(clusterDirective)
	OriginDirective.AddCommand(terminalDirective)
	OriginDirective.AddCommand(reverberateDirective)
	OriginDirective.AddCommand(detailsDirective)
	OriginDirective.AddCommand(inspectTransferDirective)
	OriginDirective.AddCommand(endorseDirective)
	OriginDirective.AddCommand(editionDirective)
	OriginDirective.AddCommand(verifyDirective)
	OriginDirective.AddCommand(arrangeNominationDirective)
	OriginDirective.AddCommand(handleNominationDirective)
	appendInquireSwitches()
	OriginDirective.AddCommand(inquireDirective)
	OriginDirective.AddCommand(culminateLedgerDirective)

	//
	appendTokvalDepotSwitches()
	OriginDirective.AddCommand(statedepotDirective)
}

var clusterDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDn

REDACTEDs
REDACTED:

REDACTEDe

REDACTED:

REDACTED0
REDACTEDf
REDACTED0
REDACTED0
REDACTEDf
REDACTEDo
REDACTED`,
	Args: cobra.ExactArgs(0),
	RunE: directiveCluster,
}

var terminalDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDs

REDACTEDs
REDACTEDe
REDACTED`,
	Args:      cobra.ExactArgs(0),
	ValidArgs: []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
	RunE:      directiveTerminal,
}

var reverberateDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  directiveReverberate,
}

var detailsDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  directiveDetails,
}

var culminateLedgerDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(1),
	RunE:  directiveCulminateLedger,
}

var inspectTransferDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  directiveInspectTransfer,
}

var endorseDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  directiveEndorse,
}

var editionDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, arguments []string) error {
		fmt.Println(edition.Edition)
		return nil
	},
}

var arrangeNominationDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(0),
	RunE:  directiveArrangeNomination,
}

var handleNominationDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(0),
	RunE:  directiveHandleNomination,
}

var inquireDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  directiveInquire,
}

var statedepotDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  directiveTokvalDepot,
}

var verifyDirective = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  directiveVerify,
}

//
func enduringArguments(row []byte) []string {
	//
	//
	arguments := os.Args
	arguments = arguments[:len(arguments)-1] //

	if len(row) > 0 { //
		arguments = append(arguments, strings.Split(string(row), "REDACTED")...)
	}
	return arguments
}

//

func construct(fs []func() error) error {
	if len(fs) == 0 {
		return nil
	}

	err := fs[0]()
	if err == nil {
		return construct(fs[1:])
	}

	return err
}

func directiveVerify(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	return construct(
		[]func() error{
			func() error { return daemontest.InitializeSuccession(ctx, customer) },
			func() error { return daemontest.Endorse(ctx, customer) },
			func() error {
				return daemontest.CulminateLedger(ctx, customer, [][]byte{
					[]byte("REDACTED"),
				}, []uint32{
					statedepot.CipherKindUnfitTransferLayout,
				}, nil, nil)
			},
			func() error { return daemontest.Endorse(ctx, customer) },
			func() error {
				return daemontest.CulminateLedger(ctx, customer, [][]byte{
					{0x00},
				}, []uint32{
					statedepot.CipherKindOKAY,
				}, nil, []byte{0, 0, 0, 0, 0, 0, 0, 1})
			},
			func() error { return daemontest.Endorse(ctx, customer) },
			func() error {
				return daemontest.CulminateLedger(ctx, customer, [][]byte{
					{0x00},
					{0x01},
					{0x00, 0x02},
					{0x00, 0x03},
					{0x00, 0x00, 0x04},
					{0x00, 0x00, 0x06},
				}, []uint32{
					statedepot.CipherKindUnfitTransferLayout,
					statedepot.CipherKindOKAY,
					statedepot.CipherKindOKAY,
					statedepot.CipherKindOKAY,
					statedepot.CipherKindOKAY,
					statedepot.CipherKindUnfitTransferLayout,
				}, nil, []byte{0, 0, 0, 0, 0, 0, 0, 5})
			},
			func() error { return daemontest.Endorse(ctx, customer) },
			func() error {
				return daemontest.ArrangeNomination(ctx, customer, [][]byte{
					{0x01},
				}, [][]byte{{0x01}}, nil)
			},
			func() error {
				return daemontest.HandleNomination(ctx, customer, [][]byte{
					{0x01},
				}, kinds.Responseexecuteitem_EMBRACE)
			},
		})
}

func directiveCluster(cmd *cobra.Command, _ []string) error {
	bufferFetcher := bufio.NewReader(os.Stdin)
Cycle:
	for {

		row, extra, err := bufferFetcher.ReadLine()
		switch {
		case extra:
			return errors.New("REDACTED")
		case err == io.EOF:
			break Cycle
		case len(row) == 0:
			continue
		case err != nil:
			return err
		}

		directiveArguments := enduringArguments(row)
		if err := multiplexerUponDirectives(cmd, directiveArguments); err != nil {
			return err
		}
		fmt.Println()
	}
	return nil
}

func directiveTerminal(cmd *cobra.Command, _ []string) error {
	for {
		fmt.Printf("REDACTED")
		bufferFetcher := bufio.NewReader(os.Stdin)
		row, extra, err := bufferFetcher.ReadLine()
		if extra {
			return errors.New("REDACTED")
		} else if err != nil {
			return err
		}

		paramArguments := enduringArguments(row)
		if err := multiplexerUponDirectives(cmd, paramArguments); err != nil {
			return err
		}
	}
}

func multiplexerUponDirectives(cmd *cobra.Command, paramArguments []string) error {
	if len(paramArguments) < 2 {
		return errors.New("REDACTED")
	}

	//
	arguments := []string{}
	for i := 0; i < len(paramArguments); i++ {
		arg := paramArguments[i]

		//
		if strings.HasPrefix(arg, "REDACTED") {
			//
			if strings.Contains(arg, "REDACTED") {
				continue
			}
			//
			_, err := cmd.Flags().GetBool(strings.TrimLeft(arg, "REDACTED"))
			if err == nil {
				continue
			}

			//
			i++
			continue
		}

		//
		arguments = append(arguments, arg)
	}
	var underDirective string
	var preciseArguments []string
	if len(arguments) > 1 {
		underDirective = arguments[1]
	}
	if len(arguments) > 2 {
		preciseArguments = arguments[2:]
	}
	cmd.Use = underDirective //

	switch strings.ToLower(underDirective) {
	case "REDACTED":
		return directiveInspectTransfer(cmd, preciseArguments)
	case "REDACTED":
		return directiveEndorse(cmd, preciseArguments)
	case "REDACTED":
		return directiveCulminateLedger(cmd, preciseArguments)
	case "REDACTED":
		return directiveReverberate(cmd, preciseArguments)
	case "REDACTED":
		return directiveDetails(cmd, preciseArguments)
	case "REDACTED":
		return directiveInquire(cmd, preciseArguments)
	case "REDACTED":
		return directiveArrangeNomination(cmd, preciseArguments)
	case "REDACTED":
		return directiveHandleNomination(cmd, preciseArguments)
	default:
		return directiveUndeveloped(cmd, paramArguments)
	}
}

func directiveUndeveloped(cmd *cobra.Command, arguments []string) error {
	msg := "REDACTED"

	if len(arguments) > 0 {
		msg += fmt.Sprintf("REDACTED", strings.Join(arguments, "REDACTED"))
	}
	renderReply(cmd, arguments, reply{
		Cipher: cipherFlawed,
		Log:  msg,
	})

	fmt.Println("REDACTED")
	fmt.Printf("REDACTED", reverberateDirective.Use, reverberateDirective.Short)
	fmt.Printf("REDACTED", inspectTransferDirective.Use, inspectTransferDirective.Short)
	fmt.Printf("REDACTED", endorseDirective.Use, endorseDirective.Short)
	fmt.Printf("REDACTED", culminateLedgerDirective.Use, culminateLedgerDirective.Short)
	fmt.Printf("REDACTED", detailsDirective.Use, detailsDirective.Short)
	fmt.Printf("REDACTED", inquireDirective.Use, inquireDirective.Short)
	fmt.Printf("REDACTED", arrangeNominationDirective.Use, arrangeNominationDirective.Short)
	fmt.Printf("REDACTED", handleNominationDirective.Use, handleNominationDirective.Short)

	fmt.Println("REDACTED")

	return nil
}

//
func directiveReverberate(cmd *cobra.Command, arguments []string) error {
	msg := "REDACTED"
	if len(arguments) > 0 {
		msg = arguments[0]
	}
	res, err := customer.Reverberate(cmd.Context(), msg)
	if err != nil {
		return err
	}

	renderReply(cmd, arguments, reply{
		Data: []byte(res.Signal),
	})

	return nil
}

//
func directiveDetails(cmd *cobra.Command, arguments []string) error {
	var edition string
	if len(arguments) == 1 {
		edition = arguments[0]
	}
	res, err := customer.Details(cmd.Context(), &kinds.SolicitDetails{Edition: edition})
	if err != nil {
		return err
	}
	renderReply(cmd, arguments, reply{
		Data: []byte(res.Data),
	})
	return nil
}

const cipherFlawed uint32 = 10

//
func directiveCulminateLedger(cmd *cobra.Command, arguments []string) error {
	if len(arguments) == 0 {
		renderReply(cmd, arguments, reply{
			Cipher: cipherFlawed,
			Log:  "REDACTED",
		})
		return nil
	}
	txs := make([][]byte, len(arguments))
	for i, arg := range arguments {
		transferOctets, err := textEitherHexadecimalTowardOctets(arg)
		if err != nil {
			return err
		}
		txs[i] = transferOctets
	}
	res, err := customer.CulminateLedger(cmd.Context(), &kinds.SolicitCulminateLedger{Txs: txs})
	if err != nil {
		return err
	}
	replies := make([]reply, 0, len(res.TransferOutcomes)+1)
	for _, tx := range res.TransferOutcomes {
		replies = append(replies, reply{
			Cipher: tx.Cipher,
			Data: tx.Data,
			Details: tx.Details,
			Log:  tx.Log,
		})
	}
	replies = append(replies, reply{
		Data: res.PlatformDigest,
	})
	renderReply(cmd, arguments, replies...)
	return nil
}

//
func directiveInspectTransfer(cmd *cobra.Command, arguments []string) error {
	if len(arguments) == 0 {
		renderReply(cmd, arguments, reply{
			Cipher: cipherFlawed,
			Details: "REDACTED",
		})
		return nil
	}
	transferOctets, err := textEitherHexadecimalTowardOctets(arguments[0])
	if err != nil {
		return err
	}
	res, err := customer.InspectTransfer(cmd.Context(), &kinds.SolicitInspectTransfer{Tx: transferOctets})
	if err != nil {
		return err
	}
	renderReply(cmd, arguments, reply{
		Cipher: res.Cipher,
		Data: res.Data,
		Details: res.Details,
		Log:  res.Log,
	})
	return nil
}

//
func directiveEndorse(cmd *cobra.Command, arguments []string) error {
	_, err := customer.Endorse(cmd.Context(), &kinds.SolicitEndorse{})
	if err != nil {
		return err
	}
	renderReply(cmd, arguments, reply{})
	return nil
}

//
func directiveInquire(cmd *cobra.Command, arguments []string) error {
	if len(arguments) == 0 {
		renderReply(cmd, arguments, reply{
			Cipher: cipherFlawed,
			Details: "REDACTED",
			Log:  "REDACTED",
		})
		return nil
	}
	inquireOctets, err := textEitherHexadecimalTowardOctets(arguments[0])
	if err != nil {
		return err
	}

	outcomeInquire, err := customer.Inquire(cmd.Context(), &kinds.SolicitInquire{
		Data:   inquireOctets,
		Route:   markerRoute,
		Altitude: int64(markerAltitude),
		Validate:  markerValidate,
	})
	if err != nil {
		return err
	}
	renderReply(cmd, arguments, reply{
		Cipher: outcomeInquire.Cipher,
		Details: outcomeInquire.Details,
		Log:  outcomeInquire.Log,
		Inquire: &inquireReply{
			Key:      outcomeInquire.Key,
			Datum:    outcomeInquire.Datum,
			Altitude:   outcomeInquire.Altitude,
			AttestationActions: outcomeInquire.AttestationActions,
		},
	})
	return nil
}

func directiveArrangeNomination(cmd *cobra.Command, arguments []string) error {
	transOctetsSeries := make([][]byte, len(arguments))

	for i, arg := range arguments {
		transferOctets, err := textEitherHexadecimalTowardOctets(arg)
		if err != nil {
			return err
		}
		transOctetsSeries[i] = transferOctets
	}

	res, err := customer.ArrangeNomination(cmd.Context(), &kinds.SolicitArrangeNomination{
		Txs: transOctetsSeries,
		//
		MaximumTransferOctets: 65536,
	})
	if err != nil {
		return err
	}
	replies := make([]reply, 0, len(res.Txs))
	for _, tx := range res.Txs {
		replies = append(replies, reply{
			Cipher: 0, //
			Log:  "REDACTED" + string(tx),
		})
	}

	renderReply(cmd, arguments, replies...)
	return nil
}

func directiveHandleNomination(cmd *cobra.Command, arguments []string) error {
	transOctetsSeries := make([][]byte, len(arguments))

	for i, arg := range arguments {
		transferOctets, err := textEitherHexadecimalTowardOctets(arg)
		if err != nil {
			return err
		}
		transOctetsSeries[i] = transferOctets
	}

	res, err := customer.HandleNomination(cmd.Context(), &kinds.SolicitHandleNomination{
		Txs: transOctetsSeries,
	})
	if err != nil {
		return err
	}

	renderReply(cmd, arguments, reply{
		Condition: int32(res.Condition),
	})
	return nil
}

func directiveTokvalDepot(*cobra.Command, []string) error {
	tracer := log.FreshTEMPTracer(log.FreshChronizePersistor(os.Stdout))

	//
	var app kinds.Platform
	if markerEndure == "REDACTED" {
		var err error
		markerEndure, err = os.MkdirTemp("REDACTED", "REDACTED")
		if err != nil {
			return err
		}
	}
	app = statedepot.FreshEnduringPlatform(markerEndure)

	//
	srv, err := node.FreshDaemon(markerLocator, markerIface, app)
	if err != nil {
		return err
	}
	srv.AssignTracer(tracer.Using("REDACTED", "REDACTED"))
	if err := srv.Initiate(); err != nil {
		return err
	}

	//
	strongos.EnsnareGesture(tracer, func() {
		//
		if err := srv.Halt(); err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
		}
	})

	//
	select {}
}

//

func renderReply(cmd *cobra.Command, arguments []string, replays ...reply) {
	if markerDetailed {
		fmt.Println("REDACTED", cmd.Use, strings.Join(arguments, "REDACTED"))
	}

	for _, rsp := range replays {
		//
		if rsp.Cipher == kinds.CipherKindOKAY {
			fmt.Printf("REDACTED")
		} else {
			fmt.Printf("REDACTED", rsp.Cipher)
		}

		if len(rsp.Data) != 0 {
			//
			//
			if cmd.Use != "REDACTED" {
				fmt.Printf("REDACTED", rsp.Data)
			}
			fmt.Printf("REDACTED", rsp.Data)
		}
		if rsp.Log != "REDACTED" {
			fmt.Printf("REDACTED", rsp.Log)
		}
		if cmd.Use == "REDACTED" {
			fmt.Printf("REDACTED", kinds.Responseexecuteitem_Itemstatus_alias[rsp.Condition])
		}

		if rsp.Inquire != nil {
			fmt.Printf("REDACTED", rsp.Inquire.Altitude)
			if rsp.Inquire.Key != nil {
				fmt.Printf("REDACTED", rsp.Inquire.Key)
				fmt.Printf("REDACTED", rsp.Inquire.Key)
			}
			if rsp.Inquire.Datum != nil {
				fmt.Printf("REDACTED", rsp.Inquire.Datum)
				fmt.Printf("REDACTED", rsp.Inquire.Datum)
			}
			if rsp.Inquire.AttestationActions != nil {
				fmt.Printf("REDACTED", rsp.Inquire.AttestationActions)
			}
		}
	}
}

//
func textEitherHexadecimalTowardOctets(s string) ([]byte, error) {
	if len(s) > 2 && strings.ToLower(s[:2]) == "REDACTED" {
		b, err := hex.DecodeString(s[2:])
		if err != nil {
			err = fmt.Errorf("REDACTED", err.Error())
			return nil, err
		}
		return b, nil
	}

	if !strings.HasPrefix(s, "REDACTED") || !strings.HasSuffix(s, "REDACTED") {
		err := fmt.Errorf("REDACTED", s)
		return nil, err
	}

	return []byte(s[1 : len(s)-1]), nil
}
