package sqls

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/adlio/schema"
	"github.com/cosmos/gogoproto/proto"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	iface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	engineevent "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status/transferordinal"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"

	//
	_ "github.com/lib/pq"
)

var (
	conductBreakLocatedQuit = flag.Bool("REDACTED", false,
		"REDACTED")

	//
	//
	verifyDatastore func() *sql.DB
)

const (
	consumer     = "REDACTED"
	secret = "REDACTED"
	channel     = "REDACTED"
	dsn      = "REDACTED"
	datastoreAlias   = "REDACTED"
	successionUUID  = "REDACTED"

	displayLedgerIncidents = "REDACTED"
	displayTransferIncidents    = "REDACTED"
)

func VerifyPrimary(m *testing.M) {
	flag.Parse()

	//
	hub, err := dockertest.NewPool(os.Getenv("REDACTED"))
	if err != nil {
		log.Fatalf("REDACTED", err)
	}

	asset, err := hub.RunWithOptions(&dockertest.RunOptions{
		Repository: "REDACTED",
		Tag:        "REDACTED",
		Env: []string{
			"REDACTED" + consumer,
			"REDACTED" + secret,
			"REDACTED" + datastoreAlias,
			"REDACTED",
		},
	}, func(settings *docker.HostConfig) {
		//
		settings.AutoRemove = true
		settings.RestartPolicy = docker.RestartPolicy{
			Name: "REDACTED",
		}
	})
	if err != nil {
		log.Fatalf("REDACTED", err)
	}

	if *conductBreakLocatedQuit {
		log.Print("REDACTED")
	} else {
		const lapseMoments = 60
		_ = asset.Expire(lapseMoments)
		log.Printf("REDACTED", lapseMoments)
	}

	//
	//
	link := fmt.Sprintf(dsn, consumer, secret, asset.GetPort(channel+"REDACTED"), datastoreAlias)
	var db *sql.DB

	if err := hub.Retry(func() error {
		receiver, err := FreshIncidentReceiver(link, successionUUID)
		if err != nil {
			return err
		}
		db = receiver.DB() //
		return db.Ping()
	}); err != nil {
		log.Fatalf("REDACTED", err)
	}

	if err := restoreDatastore(db); err != nil {
		log.Fatalf("REDACTED", err)
	}

	sm, err := retrieveBlueprint()
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	relocator := schema.NewMigrator()
	if err := relocator.Apply(db, sm); err != nil {
		log.Fatalf("REDACTED", err)
	}

	//
	verifyDatastore = func() *sql.DB { return db }

	//
	cipher := m.Run()

	//
	if *conductBreakLocatedQuit {
		log.Print("REDACTED")
		pauseForeachDisrupt()
		log.Print("REDACTED")
	}
	log.Print("REDACTED")
	if err := hub.Purge(asset); err != nil {
		log.Printf("REDACTED", err)
	}
	if err := db.Close(); err != nil {
		log.Printf("REDACTED", err)
	}

	os.Exit(cipher)
}

func VerifyCataloging(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		ordinalizer := &IncidentReceiver{depot: verifyDatastore(), successionUUID: successionUUID}
		require.NoError(t, ordinalizer.PositionLedgerIncidents(freshVerifyLedgerIncidents()))

		validateLedger(t, 1)
		validateLedger(t, 2)

		validateNegationExecuted(t, "REDACTED", func() (bool, error) { return ordinalizer.OwnsLedger(1) })
		validateNegationExecuted(t, "REDACTED", func() (bool, error) { return ordinalizer.OwnsLedger(2) })

		validateNegationExecuted(t, "REDACTED", func() (bool, error) {
			v, err := ordinalizer.LookupLedgerIncidents(context.Background(), nil)
			return v != nil, err
		})

		require.NoError(t, validateMomentImprint(registryLedgers))

		//
		require.NoError(t, ordinalizer.PositionLedgerIncidents(freshVerifyLedgerIncidents()))
	})

	t.Run("REDACTED", func(t *testing.T) {
		ordinalizer := &IncidentReceiver{depot: verifyDatastore(), successionUUID: successionUUID}

		transferOutcome := transferOutcomeUsingIncidents([]iface.Incident{
			createPositionedIncident("REDACTED", "REDACTED"),
			createPositionedIncident("REDACTED", "REDACTED"),
			createPositionedIncident("REDACTED", "REDACTED"),

			{Kind: "REDACTED", Properties: []iface.IncidentProperty{
				{
					Key:   "REDACTED",
					Datum: "REDACTED",
					Ordinal: true,
				},
			}},
		})
		require.NoError(t, ordinalizer.PositionTransferIncidents([]*iface.TransferOutcome{transferOutcome}))

		txr, err := fetchTransferOutcome(kinds.Tx(transferOutcome.Tx).Digest())
		require.NoError(t, err)
		assert.Equal(t, transferOutcome, txr)

		require.NoError(t, validateMomentImprint(registryTransferOutcomes))
		require.NoError(t, validateMomentImprint(displayTransferIncidents))

		validateNegationExecuted(t, "REDACTED", func() (bool, error) {
			txr, err := ordinalizer.ObtainTransferViaDigest(kinds.Tx(transferOutcome.Tx).Digest())
			return txr != nil, err
		})
		validateNegationExecuted(t, "REDACTED", func() (bool, error) {
			txr, err := ordinalizer.LookupTransferIncidents(context.Background(), nil)
			return txr != nil, err
		})

		//
		err = ordinalizer.PositionTransferIncidents([]*iface.TransferOutcome{transferOutcome})
		require.NoError(t, err)
	})

	t.Run("REDACTED", func(t *testing.T) {
		ordinalizer := &IncidentReceiver{depot: verifyDatastore(), successionUUID: successionUUID}

		//
		incidentPipeline := kinds.FreshIncidentPipeline()
		err := incidentPipeline.Initiate()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := incidentPipeline.Halt(); err != nil {
				t.Error(err)
			}
		})

		facility := transferordinal.FreshOrdinalizerFacility(ordinalizer.TransferOrdinalizer(), ordinalizer.LedgerOrdinalizer(), incidentPipeline, true)
		facility.AssignTracer(engineevent.VerifyingTracer())
		err = facility.Initiate()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := facility.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		err = incidentPipeline.BroadcastIncidentFreshLedgerIncidents(kinds.IncidentDataFreshLedgerIncidents{
			Altitude: 1,
			CountTrans: 2,
		})
		require.NoError(t, err)
		transferOutcome1 := &iface.TransferOutcome{
			Altitude: 1,
			Ordinal:  uint32(0),
			Tx:     kinds.Tx("REDACTED"),
			Outcome: iface.InvokeTransferOutcome{Cipher: 0},
		}
		err = incidentPipeline.BroadcastIncidentTransfer(kinds.IncidentDataTransfer{TransferOutcome: *transferOutcome1})
		require.NoError(t, err)
		transferOutcome2 := &iface.TransferOutcome{
			Altitude: 1,
			Ordinal:  uint32(1),
			Tx:     kinds.Tx("REDACTED"),
			Outcome: iface.InvokeTransferOutcome{Cipher: 1},
		}
		err = incidentPipeline.BroadcastIncidentTransfer(kinds.IncidentDataTransfer{TransferOutcome: *transferOutcome2})
		require.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		require.True(t, facility.EqualsActive())
	})
}

func VerifyHalt(t *testing.T) {
	ordinalizer := &IncidentReceiver{depot: verifyDatastore()}
	require.NoError(t, ordinalizer.Halt())
}

//
//
func freshVerifyLedgerIncidents() kinds.IncidentDataFreshLedgerIncidents {
	return kinds.IncidentDataFreshLedgerIncidents{
		Altitude: 1,
		Incidents: []iface.Incident{
			createPositionedIncident("REDACTED", "REDACTED"),
			createPositionedIncident("REDACTED", "REDACTED"),
			createPositionedIncident("REDACTED", "REDACTED"),
			createPositionedIncident("REDACTED", "REDACTED"),
		},
	}
}

//
func retrieveBlueprint() ([]*schema.Migration, error) {
	const recordname = "REDACTED"
	material, err := os.ReadFile(recordname)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", recordname, err)
	}

	return []*schema.Migration{{
		ID:     time.Now().Local().String() + "REDACTED",
		Script: string(material),
	}}, nil
}

//
func restoreDatastore(db *sql.DB) error {
	_, err := db.Exec("REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	_, err = db.Exec("REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	return nil
}

//
//
func transferOutcomeUsingIncidents(incidents []iface.Incident) *iface.TransferOutcome {
	return &iface.TransferOutcome{
		Altitude: 1,
		Ordinal:  0,
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{
			Data:   []byte{0},
			Cipher:   iface.CipherKindOKAY,
			Log:    "REDACTED",
			Incidents: incidents,
		},
	}
}

func fetchTransferOutcome(digest []byte) (*iface.TransferOutcome, error) {
	digestText := fmt.Sprintf("REDACTED", digest)
	var outcomeData []byte
	if err := verifyDatastore().QueryRow(`
REDACTED`+registryTransferOutcomes+`REDACTED;
REDACTED`, digestText).Scan(&outcomeData); err != nil {
		return nil, fmt.Errorf("REDACTED", digestText, err)
	}

	txr := new(iface.TransferOutcome)
	if err := proto.Unmarshal(outcomeData, txr); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return txr, nil
}

func validateMomentImprint(registryAlias string) error {
	return verifyDatastore().QueryRow(fmt.Sprintf(`
REDACTEDt
REDACTEDs
REDACTED;
REDACTED`, registryAlias), time.Now().Add(-2*time.Second)).Err()
}

func validateLedger(t *testing.T, altitude int64) {
	//
	if err := verifyDatastore().QueryRow(`
REDACTED`+registryLedgers+`REDACTED;
REDACTED`, altitude).Err(); err == sql.ErrNoRows {
		t.Errorf("REDACTED", altitude)
	} else if err != nil {
		t.Fatalf("REDACTED", err)
	}

	//
	if err := verifyDatastore().QueryRow(`
REDACTED`+displayLedgerIncidents+`
REDACTED;
REDACTED`, altitude, incidentKindCulminateLedger, successionUUID).Err(); err == sql.ErrNoRows {
		t.Errorf("REDACTED", incidentKindCulminateLedger, altitude)
	} else if err != nil {
		t.Fatalf("REDACTED", err)
	}
}

//
//
//
func validateNegationExecuted(t *testing.T, tag string, f func() (bool, error)) {
	t.Helper()
	t.Logf("REDACTED", tag)

	desire := tag + "REDACTED"
	ok, err := f()
	assert.False(t, ok)
	require.NotNil(t, err)
	assert.Equal(t, desire, err.Error())
}

//
func pauseForeachDisrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
