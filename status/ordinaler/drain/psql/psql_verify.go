package psql

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

	iface "github.com/valkyrieworks/iface/kinds"
	tmlog "github.com/valkyrieworks/utils/log"
	"github.com/valkyrieworks/status/transordinal"
	"github.com/valkyrieworks/kinds"

	//
	_ "github.com/lib/pq"
)

var (
	doStallAtQuit = flag.Bool("REDACTED", false,
		"REDACTED")

	//
	//
	verifyStore func() *sql.DB
)

const (
	member     = "REDACTED"
	secret = "REDACTED"
	port     = "REDACTED"
	dsn      = "REDACTED"
	storeLabel   = "REDACTED"
	ledgerUID  = "REDACTED"

	displayLedgerEvents = "REDACTED"
	displayTransferEvents    = "REDACTED"
)

func VerifyMain(m *testing.M) {
	flag.Parse()

	//
	depository, err := dockertest.NewPool(os.Getenv("REDACTED"))
	if err != nil {
		log.Fatalf("REDACTED", err)
	}

	asset, err := depository.RunWithOptions(&dockertest.RunOptions{
		Repository: "REDACTED",
		Tag:        "REDACTED",
		Env: []string{
			"REDACTED" + member,
			"REDACTED" + secret,
			"REDACTED" + storeLabel,
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

	if *doStallAtQuit {
		log.Print("REDACTED")
	} else {
		const obsoleteMoments = 60
		_ = asset.Expire(obsoleteMoments)
		log.Printf("REDACTED", obsoleteMoments)
	}

	//
	//
	link := fmt.Sprintf(dsn, member, secret, asset.GetPort(port+"REDACTED"), storeLabel)
	var db *sql.DB

	if err := depository.Retry(func() error {
		drain, err := NewEventDrain(link, ledgerUID)
		if err != nil {
			return err
		}
		db = drain.DB() //
		return db.Ping()
	}); err != nil {
		log.Fatalf("REDACTED", err)
	}

	if err := restoreDepot(db); err != nil {
		log.Fatalf("REDACTED", err)
	}

	sm, err := fetchBlueprint()
	if err != nil {
		log.Fatalf("REDACTED", err)
	}
	relocator := schema.NewMigrator()
	if err := relocator.Apply(db, sm); err != nil {
		log.Fatalf("REDACTED", err)
	}

	//
	verifyStore = func() *sql.DB { return db }

	//
	code := m.Run()

	//
	if *doStallAtQuit {
		log.Print("REDACTED")
		waitForDisrupt()
		log.Print("REDACTED")
	}
	log.Print("REDACTED")
	if err := depository.Purge(asset); err != nil {
		log.Printf("REDACTED", err)
	}
	if err := db.Close(); err != nil {
		log.Printf("REDACTED", err)
	}

	os.Exit(code)
}

func VerifyCataloging(t *testing.T) {
	t.Run("REDACTED", func(t *testing.T) {
		ordinaler := &EventDrain{depot: verifyStore(), ledgerUID: ledgerUID}
		require.NoError(t, ordinaler.OrdinalLedgerEvents(newVerifyLedgerEvents()))

		validateLedger(t, 1)
		validateLedger(t, 2)

		validateNotExecuted(t, "REDACTED", func() (bool, error) { return ordinaler.HasLedger(1) })
		validateNotExecuted(t, "REDACTED", func() (bool, error) { return ordinaler.HasLedger(2) })

		validateNotExecuted(t, "REDACTED", func() (bool, error) {
			v, err := ordinaler.ScanLedgerEvents(context.Background(), nil)
			return v != nil, err
		})

		require.NoError(t, validateTimeImprint(sheetLedgers))

		//
		require.NoError(t, ordinaler.OrdinalLedgerEvents(newVerifyLedgerEvents()))
	})

	t.Run("REDACTED", func(t *testing.T) {
		ordinaler := &EventDrain{depot: verifyStore(), ledgerUID: ledgerUID}

		transOutcome := transferOutcomeWithEvents([]iface.Event{
			createCatalogedEvent("REDACTED", "REDACTED"),
			createCatalogedEvent("REDACTED", "REDACTED"),
			createCatalogedEvent("REDACTED", "REDACTED"),

			{Kind: "REDACTED", Properties: []iface.EventProperty{
				{
					Key:   "REDACTED",
					Item: "REDACTED",
					Ordinal: true,
				},
			}},
		})
		require.NoError(t, ordinaler.OrdinalTransferEvents([]*iface.TransOutcome{transOutcome}))

		txr, err := importTransferOutcome(kinds.Tx(transOutcome.Tx).Digest())
		require.NoError(t, err)
		assert.Equal(t, transOutcome, txr)

		require.NoError(t, validateTimeImprint(sheetTransferOutcomes))
		require.NoError(t, validateTimeImprint(displayTransferEvents))

		validateNotExecuted(t, "REDACTED", func() (bool, error) {
			txr, err := ordinaler.FetchTransferByDigest(kinds.Tx(transOutcome.Tx).Digest())
			return txr != nil, err
		})
		validateNotExecuted(t, "REDACTED", func() (bool, error) {
			txr, err := ordinaler.ScanTransferEvents(context.Background(), nil)
			return txr != nil, err
		})

		//
		err = ordinaler.OrdinalTransferEvents([]*iface.TransOutcome{transOutcome})
		require.NoError(t, err)
	})

	t.Run("REDACTED", func(t *testing.T) {
		ordinaler := &EventDrain{depot: verifyStore(), ledgerUID: ledgerUID}

		//
		eventBus := kinds.NewEventBus()
		err := eventBus.Begin()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := eventBus.Halt(); err != nil {
				t.Error(err)
			}
		})

		daemon := transordinal.NewOrdinalerDaemon(ordinaler.TransOrdinaler(), ordinaler.LedgerOrdinaler(), eventBus, true)
		daemon.AssignTracer(tmlog.VerifyingTracer())
		err = daemon.Begin()
		require.NoError(t, err)
		t.Cleanup(func() {
			if err := daemon.Halt(); err != nil {
				t.Error(err)
			}
		})

		//
		err = eventBus.BroadcastEventNewLedgerEvents(kinds.EventDataNewLedgerEvents{
			Level: 1,
			CountTrans: 2,
		})
		require.NoError(t, err)
		transferOutcome1 := &iface.TransOutcome{
			Level: 1,
			Ordinal:  uint32(0),
			Tx:     kinds.Tx("REDACTED"),
			Outcome: iface.InvokeTransferOutcome{Code: 0},
		}
		err = eventBus.BroadcastEventTransfer(kinds.EventDataTransfer{TransOutcome: *transferOutcome1})
		require.NoError(t, err)
		transferOutcome2 := &iface.TransOutcome{
			Level: 1,
			Ordinal:  uint32(1),
			Tx:     kinds.Tx("REDACTED"),
			Outcome: iface.InvokeTransferOutcome{Code: 1},
		}
		err = eventBus.BroadcastEventTransfer(kinds.EventDataTransfer{TransOutcome: *transferOutcome2})
		require.NoError(t, err)

		time.Sleep(100 * time.Millisecond)
		require.True(t, daemon.IsActive())
	})
}

func VerifyHalt(t *testing.T) {
	ordinaler := &EventDrain{depot: verifyStore()}
	require.NoError(t, ordinaler.Halt())
}

//
//
func newVerifyLedgerEvents() kinds.EventDataNewLedgerEvents {
	return kinds.EventDataNewLedgerEvents{
		Level: 1,
		Events: []iface.Event{
			createCatalogedEvent("REDACTED", "REDACTED"),
			createCatalogedEvent("REDACTED", "REDACTED"),
			createCatalogedEvent("REDACTED", "REDACTED"),
			createCatalogedEvent("REDACTED", "REDACTED"),
		},
	}
}

//
func fetchBlueprint() ([]*schema.Migration, error) {
	const filename = "REDACTED"
	elements, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", filename, err)
	}

	return []*schema.Migration{{
		ID:     time.Now().Local().String() + "REDACTED",
		Script: string(elements),
	}}, nil
}

//
func restoreDepot(db *sql.DB) error {
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
func transferOutcomeWithEvents(events []iface.Event) *iface.TransOutcome {
	return &iface.TransOutcome{
		Level: 1,
		Ordinal:  0,
		Tx:     kinds.Tx("REDACTED"),
		Outcome: iface.InvokeTransferOutcome{
			Data:   []byte{0},
			Code:   iface.CodeKindSuccess,
			Log:    "REDACTED",
			Events: events,
		},
	}
}

func importTransferOutcome(digest []byte) (*iface.TransOutcome, error) {
	digestString := fmt.Sprintf("REDACTED", digest)
	var outcomeData []byte
	if err := verifyStore().QueryRow(`
REDACTED`+sheetTransferOutcomes+`REDACTED;
REDACTED`, digestString).Scan(&outcomeData); err != nil {
		return nil, fmt.Errorf("REDACTED", digestString, err)
	}

	txr := new(iface.TransOutcome)
	if err := proto.Unmarshal(outcomeData, txr); err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return txr, nil
}

func validateTimeImprint(sheetLabel string) error {
	return verifyStore().QueryRow(fmt.Sprintf(`
REDACTEDt
REDACTEDs
REDACTED;
REDACTED`, sheetLabel), time.Now().Add(-2*time.Second)).Err()
}

func validateLedger(t *testing.T, level int64) {
	//
	if err := verifyStore().QueryRow(`
REDACTED`+sheetLedgers+`REDACTED;
REDACTED`, level).Err(); err == sql.ErrNoRows {
		t.Errorf("REDACTED", level)
	} else if err != nil {
		t.Fatalf("REDACTED", err)
	}

	//
	if err := verifyStore().QueryRow(`
REDACTED`+displayLedgerEvents+`
REDACTED;
REDACTED`, level, eventKindCompleteLedger, ledgerUID).Err(); err == sql.ErrNoRows {
		t.Errorf("REDACTED", eventKindCompleteLedger, level)
	} else if err != nil {
		t.Fatalf("REDACTED", err)
	}
}

//
//
//
func validateNotExecuted(t *testing.T, tag string, f func() (bool, error)) {
	t.Helper()
	t.Logf("REDACTED", tag)

	desire := tag + "REDACTED"
	ok, err := f()
	assert.False(t, ok)
	require.NotNil(t, err)
	assert.Equal(t, desire, err.Error())
}

//
func waitForDisrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
