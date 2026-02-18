package directives

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/-db"

	ifacetypes "github.com/valkyrieworks/iface/kinds"
	cometsettings "github.com/valkyrieworks/settings"
	"github.com/valkyrieworks/utils/taskstatus"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/status/ordinaler"
	ledgerordinalkv "github.com/valkyrieworks/status/ordinaler/ledger/kv"
	"github.com/valkyrieworks/status/ordinaler/drain/psql"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/status/transordinal/kv"
	"github.com/valkyrieworks/kinds"
)

const (
	reordinalErrored = "REDACTED"
)

var (
	ErrLevelNotAccessible = errors.New("REDACTED")
	ErrCorruptQuery     = errors.New("REDACTED")
)

//
var ReOrdinalEventCommand = &cobra.Command{
	Use:     "REDACTED",
	Aliases: []string{"REDACTED"},
	Short:   "REDACTED",
	Long: `
REDACTED,
REDACTED 
REDACTED 
REDACTED 
REDACTEDt
REDACTED.

REDACTEDu
REDACTED.
REDACTED`,
	Example: `
REDACTEDt
REDACTED2
REDACTED0
REDACTED0
REDACTED`,
	Run: func(cmd *cobra.Command, args []string) {
		bs, ss, err := importStatusAndLedgerDepot(settings)
		if err != nil {
			fmt.Println(reordinalErrored, err)
			return
		}

		status, err := ss.Import()
		if err != nil {
			fmt.Println(reordinalErrored, err)
			return
		}

		if err := inspectSoundLevel(bs); err != nil {
			fmt.Println(reordinalErrored, err)
			return
		}

		bi, ti, err := importEventDrains(settings, status.LedgerUID)
		if err != nil {
			fmt.Println(reordinalErrored, err)
			return
		}

		riArgs := eventReOrdinalArgs{
			beginLevel:  beginLevel,
			terminateLevel:    terminateLevel,
			ledgerOrdinaler: bi,
			transOrdinaler:    ti,
			ledgerDepot:   bs,
			statusDepot:   ss,
		}
		if err := eventReOrdinal(cmd, riArgs); err != nil {
			panic(fmt.Errorf("REDACTED", reordinalErrored, err))
		}

		fmt.Println("REDACTED")
	},
}

var (
	beginLevel int64
	terminateLevel   int64
)

func init() {
	ReOrdinalEventCommand.Flags().Int64Var(&beginLevel, "REDACTED", 0, "REDACTED")
	ReOrdinalEventCommand.Flags().Int64Var(&terminateLevel, "REDACTED", 0, "REDACTED")
}

func importEventDrains(cfg *cometsettings.Settings, ledgerUID string) (ordinaler.LedgerOrdinaler, transordinal.TransOrdinaler, error) {
	switch strings.ToLower(cfg.TransOrdinal.Ordinaler) {
	case "REDACTED":
		return nil, nil, errors.New("REDACTED")
	case "REDACTED":
		link := cfg.TransOrdinal.PsqlLink
		if link == "REDACTED" {
			return nil, nil, errors.New("REDACTED")
		}
		es, err := psql.NewEventDrain(link, ledgerUID)
		if err != nil {
			return nil, nil, err
		}
		return es.LedgerOrdinaler(), es.TransOrdinaler(), nil
	case "REDACTED":
		depot, err := dbm.NewStore("REDACTED", dbm.OriginKind(cfg.StoreOrigin), cfg.StoreFolder())
		if err != nil {
			return nil, nil, err
		}

		transOrdinaler := kv.NewTransOrdinal(depot)
		ledgerOrdinaler := ledgerordinalkv.New(dbm.NewHeadingStore(depot, []byte("REDACTED")))
		return ledgerOrdinaler, transOrdinaler, nil
	default:
		return nil, nil, fmt.Errorf("REDACTED", cfg.TransOrdinal.Ordinaler)
	}
}

type eventReOrdinalArgs struct {
	beginLevel  int64
	terminateLevel    int64
	ledgerOrdinaler ordinaler.LedgerOrdinaler
	transOrdinaler    transordinal.TransOrdinaler
	ledgerDepot   status.LedgerDepot
	statusDepot   status.Depot
}

func eventReOrdinal(cmd *cobra.Command, args eventReOrdinalArgs) error {
	var bar taskstatus.Bar
	bar.NewSetting(args.beginLevel-1, args.terminateLevel)

	fmt.Println("REDACTED")
	defer bar.Conclude()
	for level := args.beginLevel; level <= args.terminateLevel; level++ {
		select {
		case <-cmd.Context().Done():
			return fmt.Errorf("REDACTED", level, cmd.Context().Err())
		default:
			ledger := args.ledgerDepot.ImportLedger(level)
			if ledger == nil {
				return fmt.Errorf("REDACTED", level)
			}

			reply, err := args.statusDepot.ImportCompleteLedgerReply(level)
			if err != nil {
				return fmt.Errorf("REDACTED", level)
			}

			e := kinds.EventDataNewLedgerEvents{
				Level: level,
				Events: reply.Events,
			}

			countTrans := len(reply.TransOutcomes)

			var group *transordinal.Group
			if countTrans > 0 {
				group = transordinal.NewGroup(int64(countTrans))

				for idx, transOutcome := range reply.TransOutcomes {
					tr := ifacetypes.TransOutcome{
						Level: level,
						Ordinal:  uint32(idx),
						Tx:     ledger.Txs[idx],
						Outcome: *transOutcome,
					}

					if err = group.Add(&tr); err != nil {
						return fmt.Errorf("REDACTED", err)
					}
				}

				if err := args.transOrdinaler.AppendGroup(group); err != nil {
					return fmt.Errorf("REDACTED", level, err)
				}
			}

			if err := args.ledgerOrdinaler.Ordinal(e); err != nil {
				return fmt.Errorf("REDACTED", level, err)
			}
		}

		bar.Simulate(level)
	}

	return nil
}

func inspectSoundLevel(bs status.LedgerDepot) error {
	root := bs.Root()

	if beginLevel == 0 {
		beginLevel = root
		fmt.Printf("REDACTED", root)
	}

	if beginLevel < root {
		return fmt.Errorf("REDACTED",
			ErrLevelNotAccessible, beginLevel, root)
	}

	level := bs.Level()

	if beginLevel > level {
		return fmt.Errorf(
			"REDACTED", ErrLevelNotAccessible, beginLevel, level)
	}

	if terminateLevel == 0 || terminateLevel > level {
		terminateLevel = level
		fmt.Printf("REDACTED", level)
	}

	if terminateLevel < root {
		return fmt.Errorf(
			"REDACTED", ErrLevelNotAccessible, terminateLevel, root)
	}

	if terminateLevel < beginLevel {
		return fmt.Errorf(
			"REDACTED",
			ErrCorruptQuery, beginLevel, terminateLevel)
	}

	return nil
}
