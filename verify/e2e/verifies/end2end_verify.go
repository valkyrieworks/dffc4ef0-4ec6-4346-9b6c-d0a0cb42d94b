package end2end_typ_test

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	rpchttpsvc "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer/httpsvc"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

func initialize() {
	//
	//
	//
	//
}

var (
	ctx             = context.Background()
	simnetStash    = map[string]e2e.Simnet{}
	simnetStashMutex = sync.Mutex{}
	ledgersStash     = map[string][]*kinds.Ledger{}
	ledgersStashMutex  = sync.Mutex{}
)

//
//
//
//
//
//
//
func verifyPeer(t *testing.T, verifyMethod func(*testing.T, e2e.Peer)) {
	t.Helper()

	simnet := fetchSimnet(t)
	peers := simnet.Peers

	if alias := os.Getenv("REDACTED"); alias != "REDACTED" {
		peer := simnet.SearchPeer(alias)
		require.NotNil(t, peer, "REDACTED", alias, simnet.Alias)
		peers = []*e2e.Peer{peer}
	}

	for _, peer := range peers {
		if peer.Untracked() {
			continue
		}

		peer := *peer
		t.Run(peer.Alias, func(t *testing.T) {
			t.Parallel()
			verifyMethod(t, peer)
		})
	}
}

//
func fetchSimnet(t *testing.T) e2e.Simnet {
	t.Helper()

	declarationRecord := os.Getenv("REDACTED")
	if declarationRecord == "REDACTED" {
		t.Skip("REDACTED")
	}
	if !filepath.IsAbs(declarationRecord) {
		declarationRecord = filepath.Join("REDACTED", declarationRecord)
	}
	identifierdataKind := os.Getenv("REDACTED")
	identifierdataRecord := os.Getenv("REDACTED")
	if identifierdataKind != "REDACTED" && identifierdataRecord == "REDACTED" {
		t.Fatalf("REDACTED")
	}
	simnetStashMutex.Lock()
	defer simnetStashMutex.Unlock()
	if simnet, ok := simnetStash[declarationRecord]; ok {
		return simnet
	}
	m, err := e2e.FetchDeclaration(declarationRecord)
	require.NoError(t, err)

	var ifd e2e.FrameworkData
	switch identifierdataKind {
	case "REDACTED":
		ifd, err = e2e.FreshDockFrameworkData(m)
		require.NoError(t, err)
	case "REDACTED":
		ifd, err = e2e.FrameworkDataOriginatingRecord(identifierdataRecord)
		require.NoError(t, err)
	default:
	}
	require.NoError(t, err)

	simnet, err := e2e.FetchSimnet(declarationRecord, ifd)
	require.NoError(t, err)
	simnetStash[declarationRecord] = *simnet
	return *simnet
}

//
//
func acquireLedgerSuccession(t *testing.T) []*kinds.Ledger {
	t.Helper()

	simnet := fetchSimnet(t)

	//
	var (
		customer *rpchttpsvc.Httpsvc
		condition *remoteifacetypes.OutcomeCondition
	)
	for _, peer := range simnet.RepositoryPeers() {
		c, err := peer.Customer()
		require.NoError(t, err)
		s, err := c.Condition(ctx)
		require.NoError(t, err)
		if condition == nil || s.ChronizeDetails.NewestLedgerAltitude > condition.ChronizeDetails.NewestLedgerAltitude {
			customer = c
			condition = s
		}
	}
	require.NotNil(t, customer, "REDACTED")

	//
	//
	ledgersStashMutex.Lock()
	defer ledgersStashMutex.Unlock()

	originating := condition.ChronizeDetails.InitialLedgerAltitude
	to := condition.ChronizeDetails.NewestLedgerAltitude
	ledgers, ok := ledgersStash[simnet.Alias]
	if !ok {
		ledgers = make([]*kinds.Ledger, 0, to-originating+1)
	}
	if len(ledgers) > 0 {
		originating = ledgers[len(ledgers)-1].Altitude + 1
	}

	for h := originating; h <= to; h++ {
		reply, err := customer.Ledger(ctx, &(h))
		require.NoError(t, err)
		require.NotNil(t, reply.Ledger)
		require.Equal(t, h, reply.Ledger.Altitude, "REDACTED", reply.Ledger.Altitude)
		ledgers = append(ledgers, reply.Ledger)
	}
	require.NotEmpty(t, ledgers, "REDACTED")
	ledgersStash[simnet.Alias] = ledgers

	return ledgers
}
