package agreement

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	dbm "github.com/valkyrieworks/-db"

	cfg "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/log"
	cometos "github.com/valkyrieworks/utils/os"
	"github.com/valkyrieworks/gateway"
	sm "github.com/valkyrieworks/status"
	"github.com/valkyrieworks/depot"
	"github.com/valkyrieworks/kinds"
)

const (
	//
	enrollee = "REDACTED"
)

//
//

//
func ExecuteResimulateEntry(settings cfg.RootSettings, csSettings *cfg.AgreementSettings, terminal bool) {
	agreementStatus := newAgreementStatusForResimulate(settings, csSettings)

	if err := agreementStatus.ResimulateEntry(csSettings.JournalEntry(), terminal); err != nil {
		cometos.Quit(fmt.Sprintf("REDACTED", err))
	}
}

//
func (cs *Status) ResimulateEntry(entry string, terminal bool) error {
	if cs.IsActive() {
		return errors.New("REDACTED")
	}
	if cs.wal != nil {
		return errors.New("REDACTED")
	}

	cs.beginForResimulate()

	//

	ctx := context.Background()
	newPhaseEnrollee, err := cs.eventBus.Enrol(ctx, enrollee, kinds.EventInquireNewEpochPhase)
	if err != nil {
		return fmt.Errorf("REDACTED", enrollee, kinds.EventInquireNewEpochPhase)
	}
	defer func() {
		if err := cs.eventBus.Deenroll(ctx, enrollee, kinds.EventInquireNewEpochPhase); err != nil {
			cs.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}()

	//
	fp, err := os.OpenFile(entry, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}

	pb := newResimulation(entry, fp, cs, cs.status.Clone())
	defer pb.fp.Close()

	var followingN int //
	var msg *ScheduledJournalSignal
	for {
		if followingN == 0 && terminal {
			followingN = pb.resimulateTerminalCycle()
		}

		msg, err = pb.dec.Parse()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := pb.cs.scanResimulateSignal(msg, newPhaseEnrollee); err != nil {
			return err
		}

		if followingN > 0 {
			followingN--
		}
		pb.tally++
	}
}

//
//

type resimulation struct {
	cs *Status

	fp    *os.File
	dec   *JournalParser
	tally int //

	//
	entryLabel     string   //
	originStatus sm.Status //
}

func newResimulation(entryLabel string, fp *os.File, cs *Status, generateStatus sm.Status) *resimulation {
	return &resimulation{
		cs:           cs,
		fp:           fp,
		entryLabel:     entryLabel,
		originStatus: generateStatus,
		dec:          NewJournalParser(fp),
	}
}

//
func (pb *resimulation) resimulateRestore(tally int, newPhaseEnrollee kinds.Enrollment) error {
	if err := pb.cs.Halt(); err != nil {
		return err
	}
	pb.cs.Wait()

	newCS := NewStatus(pb.cs.settings, pb.originStatus.Clone(), pb.cs.ledgerExecute,
		pb.cs.ledgerDepot, pb.cs.transferAlerter, pb.cs.eventpool)
	newCS.AssignEventBus(pb.cs.eventBus)
	newCS.beginForResimulate()

	if err := pb.fp.Close(); err != nil {
		return err
	}
	fp, err := os.OpenFile(pb.entryLabel, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}
	pb.fp = fp
	pb.dec = NewJournalParser(fp)
	tally = pb.tally - tally
	fmt.Printf("REDACTED", pb.tally, tally)
	pb.tally = 0
	pb.cs = newCS
	var msg *ScheduledJournalSignal
	for i := 0; i < tally; i++ {
		msg, err = pb.dec.Parse()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if err := pb.cs.scanResimulateSignal(msg, newPhaseEnrollee); err != nil {
			return err
		}
		pb.tally++
	}
	return nil
}

func (cs *Status) beginForResimulate() {
	cs.Tracer.Fault("REDACTED")
	/*!
s
{
{
{
:
:
n
}
}
*/
}

//
func (pb *resimulation) resimulateTerminalCycle() int {
	for {
		fmt.Printf("REDACTED")
		bufferScanner := bufio.NewReader(os.Stdin)
		row, additional, err := bufferScanner.ReadLine()
		if additional {
			cometos.Quit("REDACTED")
		} else if err != nil {
			cometos.Quit(err.Error())
		}

		symbols := strings.Split(string(row), "REDACTED")
		if len(symbols) == 0 {
			continue
		}

		switch symbols[0] {
		case "REDACTED":
			//
			//

			if len(symbols) == 1 {
				return 0
			}
			i, err := strconv.Atoi(symbols[1])
			if err != nil {
				fmt.Println("REDACTED")
			} else {
				return i
			}

		case "REDACTED":
			//
			//

			//
			//

			ctx := context.Background()
			//

			newPhaseEnrollee, err := pb.cs.eventBus.Enrol(ctx, enrollee, kinds.EventInquireNewEpochPhase)
			if err != nil {
				cometos.Quit(fmt.Sprintf("REDACTED", enrollee, kinds.EventInquireNewEpochPhase))
			}
			defer func() {
				if err := pb.cs.eventBus.Deenroll(ctx, enrollee, kinds.EventInquireNewEpochPhase); err != nil {
					pb.cs.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}()

			if len(symbols) == 1 {
				if err := pb.resimulateRestore(1, newPhaseEnrollee); err != nil {
					pb.cs.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			} else {
				i, err := strconv.Atoi(symbols[1])
				if err != nil {
					fmt.Println("REDACTED")
				} else if i > pb.tally {
					fmt.Printf("REDACTED", pb.tally)
				} else if err := pb.resimulateRestore(i, newPhaseEnrollee); err != nil {
					pb.cs.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}

		case "REDACTED":
			//
			//
			//

			rs := pb.cs.EpochStatus
			if len(symbols) == 1 {
				fmt.Println(rs)
			} else {
				switch symbols[1] {
				case "REDACTED":
					fmt.Printf("REDACTED", rs.Level, rs.Cycle, rs.Phase)
				case "REDACTED":
					fmt.Println(rs.Ratifiers)
				case "REDACTED":
					fmt.Println(rs.Nomination)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.NominationLedgerSegments.StringBrief(), rs.NominationLedger.StringBrief())
				case "REDACTED":
					fmt.Println(rs.LatchedEpoch)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.LatchedLedgerSegments.StringBrief(), rs.LatchedLedger.StringBrief())
				case "REDACTED":
					fmt.Println(rs.Ballots.StringIndented("REDACTED"))

				default:
					fmt.Println("REDACTED", symbols[1])
				}
			}
		case "REDACTED":
			fmt.Println(pb.tally)
		}
	}
}

//

//
func newAgreementStatusForResimulate(settings cfg.RootSettings, csSettings *cfg.AgreementSettings) *Status {
	storeKind := dbm.OriginKind(settings.StoreOrigin)
	//
	ledgerDepotStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	if err != nil {
		cometos.Quit(err.Error())
	}
	ledgerDepot := depot.NewLedgerDepot(ledgerDepotStore)

	//
	statusStore, err := dbm.NewStore("REDACTED", storeKind, settings.StoreFolder())
	if err != nil {
		cometos.Quit(err.Error())
	}
	statusDepot := sm.NewDepot(statusStore, sm.DepotSettings{
		DropIfaceReplies: false,
	})
	gpaper, err := sm.CreateOriginPaperFromEntry(settings.OriginEntry())
	if err != nil {
		cometos.Quit(err.Error())
	}
	status, err := sm.CreateOriginStatus(gpaper)
	if err != nil {
		cometos.Quit(err.Error())
	}

	//
	customerOriginator := gateway.StandardCustomerOriginator(settings.GatewayApplication, settings.Iface, settings.StoreFolder())
	gatewayApplication := gateway.NewApplicationLinks(customerOriginator, gateway.NoopStats())
	err = gatewayApplication.Begin()
	if err != nil {
		cometos.Quit(fmt.Sprintf("REDACTED", err))
	}

	eventBus := kinds.NewEventBus()
	if err := eventBus.Begin(); err != nil {
		cometos.Quit(fmt.Sprintf("REDACTED", err))
	}

	greeter := NewGreeter(statusDepot, status, ledgerDepot, gpaper)
	greeter.AssignEventBus(eventBus)
	err = greeter.Greeting(gatewayApplication)
	if err != nil {
		cometos.Quit(fmt.Sprintf("REDACTED", err))
	}

	txpool, eventpool := emptyTxpool{}, sm.EmptyProofDepository{}
	ledgerExecute := sm.NewLedgerRunner(statusDepot, log.VerifyingTracer(), gatewayApplication.Agreement(), txpool, eventpool, ledgerDepot)

	agreementStatus := NewStatus(csSettings, status.Clone(), ledgerExecute,
		ledgerDepot, txpool, eventpool)

	agreementStatus.AssignEventBus(eventBus)
	return agreementStatus
}
