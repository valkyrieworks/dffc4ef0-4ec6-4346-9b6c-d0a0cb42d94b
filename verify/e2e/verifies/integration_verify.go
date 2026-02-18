package integration_t_test

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	rpchttp "github.com/valkyrieworks/rpc/customer/http"
	rpctypes "github.com/valkyrieworks/rpc/core/kinds"
	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
)

func init() {
	//
	//
	//
	//
}

var (
	ctx             = context.Background()
	verifychainRepository    = map[string]e2e.Verifychain{}
	verifychainRepositoryMutex = sync.Mutex{}
	ledgersRepository     = map[string][]*kinds.Ledger{}
	ledgersRepositoryMutex  = sync.Mutex{}
)

//
//
//
//
//
//
//
func verifyMember(t *testing.T, verifyFunction func(*testing.T, e2e.Member)) {
	t.Helper()

	verifychain := importVerifychain(t)
	instances := verifychain.Instances

	if label := os.Getenv("REDACTED"); label != "REDACTED" {
		member := verifychain.SearchMember(label)
		require.NotNil(t, member, "REDACTED", label, verifychain.Label)
		instances = []*e2e.Member{member}
	}

	for _, member := range instances {
		if member.Untracked() {
			continue
		}

		member := *member
		t.Run(member.Label, func(t *testing.T) {
			t.Parallel()
			verifyFunction(t, member)
		})
	}
}

//
func importVerifychain(t *testing.T) e2e.Verifychain {
	t.Helper()

	declarationEntry := os.Getenv("REDACTED")
	if declarationEntry == "REDACTED" {
		t.Skip("REDACTED")
	}
	if !filepath.IsAbs(declarationEntry) {
		declarationEntry = filepath.Join("REDACTED", declarationEntry)
	}
	ifdKind := os.Getenv("REDACTED")
	ifdEntry := os.Getenv("REDACTED")
	if ifdKind != "REDACTED" && ifdEntry == "REDACTED" {
		t.Fatalf("REDACTED")
	}
	verifychainRepositoryMutex.Lock()
	defer verifychainRepositoryMutex.Unlock()
	if verifychain, ok := verifychainRepository[declarationEntry]; ok {
		return verifychain
	}
	m, err := e2e.ImportDeclaration(declarationEntry)
	require.NoError(t, err)

	var ifd e2e.PlatformData
	switch ifdKind {
	case "REDACTED":
		ifd, err = e2e.NewDockerPlatformData(m)
		require.NoError(t, err)
	case "REDACTED":
		ifd, err = e2e.PlatformDataFromEntry(ifdEntry)
		require.NoError(t, err)
	default:
	}
	require.NoError(t, err)

	verifychain, err := e2e.ImportVerifychain(declarationEntry, ifd)
	require.NoError(t, err)
	verifychainRepository[declarationEntry] = *verifychain
	return *verifychain
}

//
//
func acquireLedgerSeries(t *testing.T) []*kinds.Ledger {
	t.Helper()

	verifychain := importVerifychain(t)

	//
	var (
		customer *rpchttp.HTTP
		state *rpctypes.OutcomeState
	)
	for _, member := range verifychain.CatalogInstances() {
		c, err := member.Customer()
		require.NoError(t, err)
		s, err := c.Status(ctx)
		require.NoError(t, err)
		if state == nil || s.AlignDetails.NewestLedgerLevel > state.AlignDetails.NewestLedgerLevel {
			customer = c
			state = s
		}
	}
	require.NotNil(t, customer, "REDACTED")

	//
	//
	ledgersRepositoryMutex.Lock()
	defer ledgersRepositoryMutex.Unlock()

	from := state.AlignDetails.OldestLedgerLevel
	to := state.AlignDetails.NewestLedgerLevel
	ledgers, ok := ledgersRepository[verifychain.Label]
	if !ok {
		ledgers = make([]*kinds.Ledger, 0, to-from+1)
	}
	if len(ledgers) > 0 {
		from = ledgers[len(ledgers)-1].Level + 1
	}

	for h := from; h <= to; h++ {
		reply, err := customer.Ledger(ctx, &(h))
		require.NoError(t, err)
		require.NotNil(t, reply.Ledger)
		require.Equal(t, h, reply.Ledger.Level, "REDACTED", reply.Ledger.Level)
		ledgers = append(ledgers, reply.Ledger)
	}
	require.NotEmpty(t, ledgers, "REDACTED")
	ledgersRepository[verifychain.Label] = ledgers

	return ledgers
}
