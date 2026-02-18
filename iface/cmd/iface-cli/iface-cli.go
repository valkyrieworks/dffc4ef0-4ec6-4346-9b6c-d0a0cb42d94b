package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"

	abciend "github.com/valkyrieworks/iface/customer"
	"github.com/valkyrieworks/iface/instance/objectdepot"
	"github.com/valkyrieworks/iface/host"
	hostverify "github.com/valkyrieworks/iface/verifies/host"
	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/iface/release"
	"github.com/valkyrieworks/schema/consensuscore/vault"
)

//
var (
	customer abciend.Customer
	tracer log.Tracer
)

//
var (
	//
	markLocation  string
	markIface     string
	markDetailed  bool   //
	markTraceLayer string //

	//
	markRoute   string
	markLevel int
	markDemonstrate  bool

	//
	markEndure string
)

var OriginCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		switch cmd.Use {
		case "REDACTED", "REDACTED", "REDACTED":
			return nil
		}

		if tracer == nil {
			permitLayer, err := log.PermitLayer(markTraceLayer)
			if err != nil {
				return err
			}
			tracer = log.NewRefine(log.NewTMTracer(log.NewAlignRecorder(os.Stdout)), permitLayer)
		}
		if customer == nil {
			var err error
			customer, err = abciend.NewCustomer(markLocation, markIface, false)
			if err != nil {
				return err
			}
			customer.AssignTracer(tracer.With("REDACTED", "REDACTED"))
			if err := customer.Begin(); err != nil {
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
	Code   uint32
	Details   string
	Log    string
	Status int32

	Inquire *inquireReply
}

type inquireReply struct {
	Key      []byte
	Item    []byte
	Level   int64
	EvidenceActions *vault.EvidenceActions
}

func Perform() error {
	appendUniversalMarks()
	appendDirectives()
	return OriginCommand.Execute()
}

func appendUniversalMarks() {
	OriginCommand.PersistentFlags().StringVarP(&markLocation,
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED")
	OriginCommand.PersistentFlags().StringVarP(&markIface, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
	OriginCommand.PersistentFlags().BoolVarP(&markDetailed,
		"REDACTED",
		"REDACTED",
		false,
		"REDACTED")
	OriginCommand.PersistentFlags().StringVarP(&markTraceLayer, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
}

func appendInquireMarks() {
	inquireCommand.PersistentFlags().StringVarP(&markRoute, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
	inquireCommand.PersistentFlags().IntVarP(&markLevel, "REDACTED", "REDACTED", 0, "REDACTED")
	inquireCommand.PersistentFlags().BoolVarP(&markDemonstrate,
		"REDACTED",
		"REDACTED",
		false,
		"REDACTED")
}

func appendObjectDepotMarks() {
	objectdepotCommand.PersistentFlags().StringVarP(&markEndure, "REDACTED", "REDACTED", "REDACTED", "REDACTED")
}

func appendDirectives() {
	OriginCommand.AddCommand(groupCommand)
	OriginCommand.AddCommand(terminalCommand)
	OriginCommand.AddCommand(replicateCommand)
	OriginCommand.AddCommand(detailsCommand)
	OriginCommand.AddCommand(inspectTransferCommand)
	OriginCommand.AddCommand(endorseCommand)
	OriginCommand.AddCommand(releaseCommand)
	OriginCommand.AddCommand(verifyCommand)
	OriginCommand.AddCommand(arrangeNominationCommand)
	OriginCommand.AddCommand(handleNominationCommand)
	appendInquireMarks()
	OriginCommand.AddCommand(inquireCommand)
	OriginCommand.AddCommand(completeLedgerCommand)

	//
	appendObjectDepotMarks()
	OriginCommand.AddCommand(objectdepotCommand)
}

var groupCommand = &cobra.Command{
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
	RunE: commandGroup,
}

var terminalCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long: `REDACTEDs

REDACTEDs
REDACTEDe
REDACTED`,
	Args:      cobra.ExactArgs(0),
	ValidArgs: []string{"REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED", "REDACTED"},
	RunE:      commandTerminal,
}

var replicateCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  commandReplicate,
}

var detailsCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  commandDetails,
}

var completeLedgerCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(1),
	RunE:  commandCompleteLedger,
}

var inspectTransferCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  commandInspectTransfer,
}

var endorseCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  commandEndorse,
}

var releaseCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(release.Release)
		return nil
	},
}

var arrangeNominationCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(0),
	RunE:  commandArrangeNomination,
}

var handleNominationCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.MinimumNArgs(0),
	RunE:  commandHandleNomination,
}

var inquireCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(1),
	RunE:  commandInquire,
}

var objectdepotCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  commandObjectDepot,
}

var verifyCommand = &cobra.Command{
	Use:   "REDACTED",
	Short: "REDACTED",
	Long:  "REDACTED",
	Args:  cobra.ExactArgs(0),
	RunE:  commandVerify,
}

//
func durableArgs(row []byte) []string {
	//
	//
	args := os.Args
	args = args[:len(args)-1] //

	if len(row) > 0 { //
		args = append(args, strings.Split(string(row), "REDACTED")...)
	}
	return args
}

//

func arrange(fs []func() error) error {
	if len(fs) == 0 {
		return nil
	}

	err := fs[0]()
	if err == nil {
		return arrange(fs[1:])
	}

	return err
}

func commandVerify(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	return arrange(
		[]func() error{
			func() error { return hostverify.InitSeries(ctx, customer) },
			func() error { return hostverify.Endorse(ctx, customer) },
			func() error {
				return hostverify.CompleteLedger(ctx, customer, [][]byte{
					[]byte("REDACTED"),
				}, []uint32{
					objectdepot.CodeKindCorruptTransferLayout,
				}, nil, nil)
			},
			func() error { return hostverify.Endorse(ctx, customer) },
			func() error {
				return hostverify.CompleteLedger(ctx, customer, [][]byte{
					{0x00},
				}, []uint32{
					objectdepot.CodeKindSuccess,
				}, nil, []byte{0, 0, 0, 0, 0, 0, 0, 1})
			},
			func() error { return hostverify.Endorse(ctx, customer) },
			func() error {
				return hostverify.CompleteLedger(ctx, customer, [][]byte{
					{0x00},
					{0x01},
					{0x00, 0x02},
					{0x00, 0x03},
					{0x00, 0x00, 0x04},
					{0x00, 0x00, 0x06},
				}, []uint32{
					objectdepot.CodeKindCorruptTransferLayout,
					objectdepot.CodeKindSuccess,
					objectdepot.CodeKindSuccess,
					objectdepot.CodeKindSuccess,
					objectdepot.CodeKindSuccess,
					objectdepot.CodeKindCorruptTransferLayout,
				}, nil, []byte{0, 0, 0, 0, 0, 0, 0, 5})
			},
			func() error { return hostverify.Endorse(ctx, customer) },
			func() error {
				return hostverify.ArrangeNomination(ctx, customer, [][]byte{
					{0x01},
				}, [][]byte{{0x01}}, nil)
			},
			func() error {
				return hostverify.HandleNomination(ctx, customer, [][]byte{
					{0x01},
				}, kinds.Responseprocessnomination_ALLOW)
			},
		})
}

func commandGroup(cmd *cobra.Command, _ []string) error {
	bufferScanner := bufio.NewReader(os.Stdin)
Cycle:
	for {

		row, additional, err := bufferScanner.ReadLine()
		switch {
		case additional:
			return errors.New("REDACTED")
		case err == io.EOF:
			break Cycle
		case len(row) == 0:
			continue
		case err != nil:
			return err
		}

		commandArgs := durableArgs(row)
		if err := multiplexerOnDirectives(cmd, commandArgs); err != nil {
			return err
		}
		fmt.Println()
	}
	return nil
}

func commandTerminal(cmd *cobra.Command, _ []string) error {
	for {
		fmt.Printf("REDACTED")
		bufferScanner := bufio.NewReader(os.Stdin)
		row, additional, err := bufferScanner.ReadLine()
		if additional {
			return errors.New("REDACTED")
		} else if err != nil {
			return err
		}

		pArgs := durableArgs(row)
		if err := multiplexerOnDirectives(cmd, pArgs); err != nil {
			return err
		}
	}
}

func multiplexerOnDirectives(cmd *cobra.Command, pArgs []string) error {
	if len(pArgs) < 2 {
		return errors.New("REDACTED")
	}

	//
	args := []string{}
	for i := 0; i < len(pArgs); i++ {
		arg := pArgs[i]

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
		args = append(args, arg)
	}
	var subtractDirective string
	var factualArgs []string
	if len(args) > 1 {
		subtractDirective = args[1]
	}
	if len(args) > 2 {
		factualArgs = args[2:]
	}
	cmd.Use = subtractDirective //

	switch strings.ToLower(subtractDirective) {
	case "REDACTED":
		return commandInspectTransfer(cmd, factualArgs)
	case "REDACTED":
		return commandEndorse(cmd, factualArgs)
	case "REDACTED":
		return commandCompleteLedger(cmd, factualArgs)
	case "REDACTED":
		return commandReplicate(cmd, factualArgs)
	case "REDACTED":
		return commandDetails(cmd, factualArgs)
	case "REDACTED":
		return commandInquire(cmd, factualArgs)
	case "REDACTED":
		return commandArrangeNomination(cmd, factualArgs)
	case "REDACTED":
		return commandHandleNomination(cmd, factualArgs)
	default:
		return commandUnexecuted(cmd, pArgs)
	}
}

func commandUnexecuted(cmd *cobra.Command, args []string) error {
	msg := "REDACTED"

	if len(args) > 0 {
		msg += fmt.Sprintf("REDACTED", strings.Join(args, "REDACTED"))
	}
	renderReply(cmd, args, reply{
		Code: codeFlawed,
		Log:  msg,
	})

	fmt.Println("REDACTED")
	fmt.Printf("REDACTED", replicateCommand.Use, replicateCommand.Short)
	fmt.Printf("REDACTED", inspectTransferCommand.Use, inspectTransferCommand.Short)
	fmt.Printf("REDACTED", endorseCommand.Use, endorseCommand.Short)
	fmt.Printf("REDACTED", completeLedgerCommand.Use, completeLedgerCommand.Short)
	fmt.Printf("REDACTED", detailsCommand.Use, detailsCommand.Short)
	fmt.Printf("REDACTED", inquireCommand.Use, inquireCommand.Short)
	fmt.Printf("REDACTED", arrangeNominationCommand.Use, arrangeNominationCommand.Short)
	fmt.Printf("REDACTED", handleNominationCommand.Use, handleNominationCommand.Short)

	fmt.Println("REDACTED")

	return nil
}

//
func commandReplicate(cmd *cobra.Command, args []string) error {
	msg := "REDACTED"
	if len(args) > 0 {
		msg = args[0]
	}
	res, err := customer.Replicate(cmd.Context(), msg)
	if err != nil {
		return err
	}

	renderReply(cmd, args, reply{
		Data: []byte(res.Signal),
	})

	return nil
}

//
func commandDetails(cmd *cobra.Command, args []string) error {
	var release string
	if len(args) == 1 {
		release = args[0]
	}
	res, err := customer.Details(cmd.Context(), &kinds.QueryDetails{Release: release})
	if err != nil {
		return err
	}
	renderReply(cmd, args, reply{
		Data: []byte(res.Data),
	})
	return nil
}

const codeFlawed uint32 = 10

//
func commandCompleteLedger(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		renderReply(cmd, args, reply{
			Code: codeFlawed,
			Log:  "REDACTED",
		})
		return nil
	}
	txs := make([][]byte, len(args))
	for i, arg := range args {
		transferOctets, err := stringOrHexToOctets(arg)
		if err != nil {
			return err
		}
		txs[i] = transferOctets
	}
	res, err := customer.CompleteLedger(cmd.Context(), &kinds.QueryCompleteLedger{Txs: txs})
	if err != nil {
		return err
	}
	replies := make([]reply, 0, len(res.TransOutcomes)+1)
	for _, tx := range res.TransOutcomes {
		replies = append(replies, reply{
			Code: tx.Code,
			Data: tx.Data,
			Details: tx.Details,
			Log:  tx.Log,
		})
	}
	replies = append(replies, reply{
		Data: res.ApplicationDigest,
	})
	renderReply(cmd, args, replies...)
	return nil
}

//
func commandInspectTransfer(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		renderReply(cmd, args, reply{
			Code: codeFlawed,
			Details: "REDACTED",
		})
		return nil
	}
	transferOctets, err := stringOrHexToOctets(args[0])
	if err != nil {
		return err
	}
	res, err := customer.InspectTransfer(cmd.Context(), &kinds.QueryInspectTransfer{Tx: transferOctets})
	if err != nil {
		return err
	}
	renderReply(cmd, args, reply{
		Code: res.Code,
		Data: res.Data,
		Details: res.Details,
		Log:  res.Log,
	})
	return nil
}

//
func commandEndorse(cmd *cobra.Command, args []string) error {
	_, err := customer.Endorse(cmd.Context(), &kinds.QueryEndorse{})
	if err != nil {
		return err
	}
	renderReply(cmd, args, reply{})
	return nil
}

//
func commandInquire(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		renderReply(cmd, args, reply{
			Code: codeFlawed,
			Details: "REDACTED",
			Log:  "REDACTED",
		})
		return nil
	}
	inquireOctets, err := stringOrHexToOctets(args[0])
	if err != nil {
		return err
	}

	outcomeInquire, err := customer.Inquire(cmd.Context(), &kinds.QueryInquire{
		Data:   inquireOctets,
		Route:   markRoute,
		Level: int64(markLevel),
		Demonstrate:  markDemonstrate,
	})
	if err != nil {
		return err
	}
	renderReply(cmd, args, reply{
		Code: outcomeInquire.Code,
		Details: outcomeInquire.Details,
		Log:  outcomeInquire.Log,
		Inquire: &inquireReply{
			Key:      outcomeInquire.Key,
			Item:    outcomeInquire.Item,
			Level:   outcomeInquire.Level,
			EvidenceActions: outcomeInquire.EvidenceActions,
		},
	})
	return nil
}

func commandArrangeNomination(cmd *cobra.Command, args []string) error {
	transOctetsList := make([][]byte, len(args))

	for i, arg := range args {
		transferOctets, err := stringOrHexToOctets(arg)
		if err != nil {
			return err
		}
		transOctetsList[i] = transferOctets
	}

	res, err := customer.ArrangeNomination(cmd.Context(), &kinds.QueryArrangeNomination{
		Txs: transOctetsList,
		//
		MaximumTransferOctets: 65536,
	})
	if err != nil {
		return err
	}
	replies := make([]reply, 0, len(res.Txs))
	for _, tx := range res.Txs {
		replies = append(replies, reply{
			Code: 0, //
			Log:  "REDACTED" + string(tx),
		})
	}

	renderReply(cmd, args, replies...)
	return nil
}

func commandHandleNomination(cmd *cobra.Command, args []string) error {
	transOctetsList := make([][]byte, len(args))

	for i, arg := range args {
		transferOctets, err := stringOrHexToOctets(arg)
		if err != nil {
			return err
		}
		transOctetsList[i] = transferOctets
	}

	res, err := customer.HandleNomination(cmd.Context(), &kinds.QueryHandleNomination{
		Txs: transOctetsList,
	})
	if err != nil {
		return err
	}

	renderReply(cmd, args, reply{
		Status: int32(res.Status),
	})
	return nil
}

func commandObjectDepot(*cobra.Command, []string) error {
	tracer := log.NewTMTracer(log.NewAlignRecorder(os.Stdout))

	//
	var app kinds.Software
	if markEndure == "REDACTED" {
		var err error
		markEndure, err = os.MkdirTemp("REDACTED", "REDACTED")
		if err != nil {
			return err
		}
	}
	app = objectdepot.NewDurableSoftware(markEndure)

	//
	srv, err := host.NewHost(markLocation, markIface, app)
	if err != nil {
		return err
	}
	srv.AssignTracer(tracer.With("REDACTED", "REDACTED"))
	if err := srv.Begin(); err != nil {
		return err
	}

	//
	cometos.InterceptAlert(tracer, func() {
		//
		if err := srv.Halt(); err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
		}
	})

	//
	select {}
}

//

func renderReply(cmd *cobra.Command, args []string, rsps ...reply) {
	if markDetailed {
		fmt.Println("REDACTED", cmd.Use, strings.Join(args, "REDACTED"))
	}

	for _, rsp := range rsps {
		//
		if rsp.Code == kinds.CodeKindSuccess {
			fmt.Printf("REDACTED")
		} else {
			fmt.Printf("REDACTED", rsp.Code)
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
			fmt.Printf("REDACTED", kinds.Responseprocessnomination_Nominationstate_label[rsp.Status])
		}

		if rsp.Inquire != nil {
			fmt.Printf("REDACTED", rsp.Inquire.Level)
			if rsp.Inquire.Key != nil {
				fmt.Printf("REDACTED", rsp.Inquire.Key)
				fmt.Printf("REDACTED", rsp.Inquire.Key)
			}
			if rsp.Inquire.Item != nil {
				fmt.Printf("REDACTED", rsp.Inquire.Item)
				fmt.Printf("REDACTED", rsp.Inquire.Item)
			}
			if rsp.Inquire.EvidenceActions != nil {
				fmt.Printf("REDACTED", rsp.Inquire.EvidenceActions)
			}
		}
	}
}

//
func stringOrHexToOctets(s string) ([]byte, error) {
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
