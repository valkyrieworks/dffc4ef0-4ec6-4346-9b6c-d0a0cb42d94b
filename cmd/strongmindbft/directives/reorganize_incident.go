package directives

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	dbm "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/-db"

	ifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindsettings "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/settings"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/indicator"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer"
	ledgeridxkv "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/ledger/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/ordinalizer/receiver/sqls"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal/kv"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

const (
	reorganizeUnsuccessful = "REDACTED"
)

var (
	FaultAltitudeUnAccessible = errors.New("REDACTED")
	FaultUnfitSolicit     = errors.New("REDACTED")
)

//
var AgainOrdinalIncidentDirective = &cobra.Command{
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
	Run: func(cmd *cobra.Command, arguments []string) {
		bs, ss, err := fetchStatusAlsoLedgerDepot(settings)
		if err != nil {
			fmt.Println(reorganizeUnsuccessful, err)
			return
		}

		status, err := ss.Fetch()
		if err != nil {
			fmt.Println(reorganizeUnsuccessful, err)
			return
		}

		if err := inspectSoundAltitude(bs); err != nil {
			fmt.Println(reorganizeUnsuccessful, err)
			return
		}

		bi, ti, err := fetchIncidentReceivers(settings, status.SuccessionUUID)
		if err != nil {
			fmt.Println(reorganizeUnsuccessful, err)
			return
		}

		randidxArguments := incidentAgainOrdinalArguments{
			initiateAltitude:  initiateAltitude,
			terminateAltitude:    terminateAltitude,
			ledgerOrdinalizer: bi,
			transferOrdinalizer:    ti,
			ledgerDepot:   bs,
			statusDepot:   ss,
		}
		if err := incidentAgainOrdinal(cmd, randidxArguments); err != nil {
			panic(fmt.Errorf("REDACTED", reorganizeUnsuccessful, err))
		}

		fmt.Println("REDACTED")
	},
}

var (
	initiateAltitude int64
	terminateAltitude   int64
)

func initialize() {
	AgainOrdinalIncidentDirective.Flags().Int64Var(&initiateAltitude, "REDACTED", 0, "REDACTED")
	AgainOrdinalIncidentDirective.Flags().Int64Var(&terminateAltitude, "REDACTED", 0, "REDACTED")
}

func fetchIncidentReceivers(cfg *strongmindsettings.Settings, successionUUID string) (ordinalizer.LedgerOrdinalizer, transferordinal.TransferOrdinalizer, error) {
	switch strings.ToLower(cfg.TransferOrdinal.Ordinalizer) {
	case "REDACTED":
		return nil, nil, errors.New("REDACTED")
	case "REDACTED":
		link := cfg.TransferOrdinal.SqlsLink
		if link == "REDACTED" {
			return nil, nil, errors.New("REDACTED")
		}
		es, err := sqls.FreshIncidentReceiver(link, successionUUID)
		if err != nil {
			return nil, nil, err
		}
		return es.LedgerOrdinalizer(), es.TransferOrdinalizer(), nil
	case "REDACTED":
		depot, err := dbm.FreshDatastore("REDACTED", dbm.OriginKind(cfg.DatastoreOrigin), cfg.DatastorePath())
		if err != nil {
			return nil, nil, err
		}

		transferOrdinalizer := kv.FreshTransferOrdinal(depot)
		ledgerOrdinalizer := ledgeridxkv.New(dbm.FreshHeadingDatastore(depot, []byte("REDACTED")))
		return ledgerOrdinalizer, transferOrdinalizer, nil
	default:
		return nil, nil, fmt.Errorf("REDACTED", cfg.TransferOrdinal.Ordinalizer)
	}
}

type incidentAgainOrdinalArguments struct {
	initiateAltitude  int64
	terminateAltitude    int64
	ledgerOrdinalizer ordinalizer.LedgerOrdinalizer
	transferOrdinalizer    transferordinal.TransferOrdinalizer
	ledgerDepot   status.LedgerDepot
	statusDepot   status.Depot
}

func incidentAgainOrdinal(cmd *cobra.Command, arguments incidentAgainOrdinalArguments) error {
	var bar indicator.Bar
	bar.FreshSelection(arguments.initiateAltitude-1, arguments.terminateAltitude)

	fmt.Println("REDACTED")
	defer bar.Conclude()
	for altitude := arguments.initiateAltitude; altitude <= arguments.terminateAltitude; altitude++ {
		select {
		case <-cmd.Context().Done():
			return fmt.Errorf("REDACTED", altitude, cmd.Context().Err())
		default:
			ledger := arguments.ledgerDepot.FetchLedger(altitude)
			if ledger == nil {
				return fmt.Errorf("REDACTED", altitude)
			}

			reply, err := arguments.statusDepot.FetchCulminateLedgerReply(altitude)
			if err != nil {
				return fmt.Errorf("REDACTED", altitude)
			}

			e := kinds.IncidentDataFreshLedgerIncidents{
				Altitude: altitude,
				Incidents: reply.Incidents,
			}

			countTrans := len(reply.TransferOutcomes)

			var cluster *transferordinal.Cluster
			if countTrans > 0 {
				cluster = transferordinal.FreshCluster(int64(countTrans))

				for idx, transferOutcome := range reply.TransferOutcomes {
					tr := ifacetypes.TransferOutcome{
						Altitude: altitude,
						Ordinal:  uint32(idx),
						Tx:     ledger.Txs[idx],
						Outcome: *transferOutcome,
					}

					if err = cluster.Add(&tr); err != nil {
						return fmt.Errorf("REDACTED", err)
					}
				}

				if err := arguments.transferOrdinalizer.AppendCluster(cluster); err != nil {
					return fmt.Errorf("REDACTED", altitude, err)
				}
			}

			if err := arguments.ledgerOrdinalizer.Ordinal(e); err != nil {
				return fmt.Errorf("REDACTED", altitude, err)
			}
		}

		bar.Enact(altitude)
	}

	return nil
}

func inspectSoundAltitude(bs status.LedgerDepot) error {
	foundation := bs.Foundation()

	if initiateAltitude == 0 {
		initiateAltitude = foundation
		fmt.Printf("REDACTED", foundation)
	}

	if initiateAltitude < foundation {
		return fmt.Errorf("REDACTED",
			FaultAltitudeUnAccessible, initiateAltitude, foundation)
	}

	altitude := bs.Altitude()

	if initiateAltitude > altitude {
		return fmt.Errorf(
			"REDACTED", FaultAltitudeUnAccessible, initiateAltitude, altitude)
	}

	if terminateAltitude == 0 || terminateAltitude > altitude {
		terminateAltitude = altitude
		fmt.Printf("REDACTED", altitude)
	}

	if terminateAltitude < foundation {
		return fmt.Errorf(
			"REDACTED", FaultAltitudeUnAccessible, terminateAltitude, foundation)
	}

	if terminateAltitude < initiateAltitude {
		return fmt.Errorf(
			"REDACTED",
			FaultUnfitSolicit, initiateAltitude, terminateAltitude)
	}

	return nil
}
