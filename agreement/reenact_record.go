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

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	cfg "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	strongos "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/os"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/delegate"
	sm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/depot"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	//
	listener = "REDACTED"
)

//
//

//
func ExecuteReenactRecord(settings cfg.FoundationSettings, controlSettings *cfg.AgreementSettings, terminal bool) {
	agreementStatus := freshAgreementStatusForeachReenact(settings, controlSettings)

	if err := agreementStatus.ReenactRecord(controlSettings.JournalRecord(), terminal); err != nil {
		strongos.Quit(fmt.Sprintf("REDACTED", err))
	}
}

//
func (cs *Status) ReenactRecord(record string, terminal bool) error {
	if cs.EqualsActive() {
		return errors.New("REDACTED")
	}
	if cs.wal != nil {
		return errors.New("REDACTED")
	}

	cs.initiateForeachReenact()

	//

	ctx := context.Background()
	freshPhaseListening, err := cs.incidentChannel.Listen(ctx, listener, kinds.IncidentInquireFreshIterationPhase)
	if err != nil {
		return fmt.Errorf("REDACTED", listener, kinds.IncidentInquireFreshIterationPhase)
	}
	defer func() {
		if err := cs.incidentChannel.Unlisten(ctx, listener, kinds.IncidentInquireFreshIterationPhase); err != nil {
			cs.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}()

	//
	fp, err := os.OpenFile(record, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}

	pb := freshReenactment(record, fp, cs, cs.status.Duplicate())
	defer pb.fp.Close()

	var followingNTH int //
	var msg *ScheduledJournalSignal
	for {
		if followingNTH == 0 && terminal {
			followingNTH = pb.reenactTerminalCycle()
		}

		msg, err = pb.dec.Deserialize()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		if err := pb.cs.fetchReenactSignal(msg, freshPhaseListening); err != nil {
			return err
		}

		if followingNTH > 0 {
			followingNTH--
		}
		pb.tally++
	}
}

//
//

type reenactment struct {
	cs *Status

	fp    *os.File
	dec   *JournalDeserializer
	tally int //

	//
	recordAlias     string   //
	inaugurationStatus sm.Status //
}

func freshReenactment(recordAlias string, fp *os.File, cs *Status, produceStatus sm.Status) *reenactment {
	return &reenactment{
		cs:           cs,
		fp:           fp,
		recordAlias:     recordAlias,
		inaugurationStatus: produceStatus,
		dec:          FreshJournalDeserializer(fp),
	}
}

//
func (pb *reenactment) reenactRestore(tally int, freshPhaseListening kinds.Listening) error {
	if err := pb.cs.Halt(); err != nil {
		return err
	}
	pb.cs.Await()

	freshControl := FreshStatus(pb.cs.settings, pb.inaugurationStatus.Duplicate(), pb.cs.ledgerExecute,
		pb.cs.ledgerDepot, pb.cs.transferObserver, pb.cs.incidentpool)
	freshControl.AssignIncidentChannel(pb.cs.incidentChannel)
	freshControl.initiateForeachReenact()

	if err := pb.fp.Close(); err != nil {
		return err
	}
	fp, err := os.OpenFile(pb.recordAlias, os.O_RDONLY, 0o600)
	if err != nil {
		return err
	}
	pb.fp = fp
	pb.dec = FreshJournalDeserializer(fp)
	tally = pb.tally - tally
	fmt.Printf("REDACTED", pb.tally, tally)
	pb.tally = 0
	pb.cs = freshControl
	var msg *ScheduledJournalSignal
	for i := 0; i < tally; i++ {
		msg, err = pb.dec.Deserialize()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		if err := pb.cs.fetchReenactSignal(msg, freshPhaseListening); err != nil {
			return err
		}
		pb.tally++
	}
	return nil
}

func (cs *Status) initiateForeachReenact() {
	cs.Tracer.Failure("REDACTED")
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
func (pb *reenactment) reenactTerminalCycle() int {
	for {
		fmt.Printf("REDACTED")
		bufferFetcher := bufio.NewReader(os.Stdin)
		row, extra, err := bufferFetcher.ReadLine()
		if extra {
			strongos.Quit("REDACTED")
		} else if err != nil {
			strongos.Quit(err.Error())
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

			freshPhaseListening, err := pb.cs.incidentChannel.Listen(ctx, listener, kinds.IncidentInquireFreshIterationPhase)
			if err != nil {
				strongos.Quit(fmt.Sprintf("REDACTED", listener, kinds.IncidentInquireFreshIterationPhase))
			}
			defer func() {
				if err := pb.cs.incidentChannel.Unlisten(ctx, listener, kinds.IncidentInquireFreshIterationPhase); err != nil {
					pb.cs.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}()

			if len(symbols) == 1 {
				if err := pb.reenactRestore(1, freshPhaseListening); err != nil {
					pb.cs.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			} else {
				i, err := strconv.Atoi(symbols[1])
				if err != nil {
					fmt.Println("REDACTED")
				} else if i > pb.tally {
					fmt.Printf("REDACTED", pb.tally)
				} else if err := pb.reenactRestore(i, freshPhaseListening); err != nil {
					pb.cs.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}

		case "REDACTED":
			//
			//
			//

			rs := pb.cs.IterationStatus
			if len(symbols) == 1 {
				fmt.Println(rs)
			} else {
				switch symbols[1] {
				case "REDACTED":
					fmt.Printf("REDACTED", rs.Altitude, rs.Iteration, rs.Phase)
				case "REDACTED":
					fmt.Println(rs.Assessors)
				case "REDACTED":
					fmt.Println(rs.Nomination)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.NominationLedgerFragments.TextBrief(), rs.NominationLedger.TextBrief())
				case "REDACTED":
					fmt.Println(rs.SecuredIteration)
				case "REDACTED":
					fmt.Printf("REDACTED", rs.SecuredLedgerFragments.TextBrief(), rs.SecuredLedger.TextBrief())
				case "REDACTED":
					fmt.Println(rs.Ballots.TextFormatted("REDACTED"))

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
func freshAgreementStatusForeachReenact(settings cfg.FoundationSettings, controlSettings *cfg.AgreementSettings) *Status {
	datastoreKind := dbm.OriginKind(settings.DatastoreRepository)
	//
	ledgerDepotDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	if err != nil {
		strongos.Quit(err.Error())
	}
	ledgerDepot := depot.FreshLedgerDepot(ledgerDepotDatastore)

	//
	statusDatastore, err := dbm.FreshDatastore("REDACTED", datastoreKind, settings.DatastorePath())
	if err != nil {
		strongos.Quit(err.Error())
	}
	statusDepot := sm.FreshDepot(statusDatastore, sm.DepotChoices{
		EjectIfaceReplies: false,
	})
	apidoc, err := sm.CreateInaugurationPaperOriginatingRecord(settings.InaugurationRecord())
	if err != nil {
		strongos.Quit(err.Error())
	}
	status, err := sm.CreateInaugurationStatus(apidoc)
	if err != nil {
		strongos.Quit(err.Error())
	}

	//
	customerOriginator := delegate.FallbackCustomerOriginator(settings.DelegateApplication, settings.Iface, settings.DatastorePath())
	delegatePlatform := delegate.FreshPlatformLinks(customerOriginator, delegate.NooperationTelemetry())
	err = delegatePlatform.Initiate()
	if err != nil {
		strongos.Quit(fmt.Sprintf("REDACTED", err))
	}

	incidentChannel := kinds.FreshIncidentPipeline()
	if err := incidentChannel.Initiate(); err != nil {
		strongos.Quit(fmt.Sprintf("REDACTED", err))
	}

	negotiator := FreshNegotiator(statusDepot, status, ledgerDepot, apidoc)
	negotiator.AssignIncidentChannel(incidentChannel)
	err = negotiator.Negotiation(delegatePlatform)
	if err != nil {
		strongos.Quit(fmt.Sprintf("REDACTED", err))
	}

	txpool, incidentpool := blankTxpool{}, sm.VoidProofHub{}
	ledgerExecute := sm.FreshLedgerHandler(statusDepot, log.VerifyingTracer(), delegatePlatform.Agreement(), txpool, incidentpool, ledgerDepot)

	agreementStatus := FreshStatus(controlSettings, status.Duplicate(), ledgerExecute,
		ledgerDepot, txpool, incidentpool)

	agreementStatus.AssignIncidentChannel(incidentChannel)
	return agreementStatus
}
