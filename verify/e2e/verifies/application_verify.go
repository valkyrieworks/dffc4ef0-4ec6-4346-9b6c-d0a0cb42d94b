package integration_t_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/verify/e2e/pkg"
	"github.com/valkyrieworks/kinds"
)

//
func Verifyapp_Primarystatus(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		if len(member.Verifychain.PrimaryStatus) == 0 {
			return
		}

		customer, err := member.Customer()
		require.NoError(t, err)
		for k, v := range member.Verifychain.PrimaryStatus {
			reply, err := customer.IfaceInquire(ctx, "REDACTED", []byte(k))
			require.NoError(t, err)
			assert.Equal(t, k, string(reply.Reply.Key))
			assert.Equal(t, v, string(reply.Reply.Item))
		}
	})
}

//
//
func Verifyapp_Digest(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		customer, err := member.Customer()
		require.NoError(t, err)

		details, err := customer.IfaceDetails(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, details.Reply.FinalLedgerApplicationDigest, "REDACTED")

		//
		queriedLevel := details.Reply.FinalLedgerLevel + 1

		require.Eventually(t, func() bool {
			state, err := customer.Status(ctx)
			require.NoError(t, err)
			require.NotZero(t, state.AlignDetails.NewestLedgerLevel)
			return state.AlignDetails.NewestLedgerLevel >= queriedLevel
		}, 5*time.Second, 500*time.Millisecond)

		ledger, err := customer.Ledger(ctx, &queriedLevel)
		require.NoError(t, err)
		require.Equal(t,
			fmt.Sprintf("REDACTED", details.Reply.FinalLedgerApplicationDigest),
			fmt.Sprintf("REDACTED", ledger.Ledger.ApplicationDigest.Octets()),
			"REDACTED")
	})
}

//
func Verifyapp_Transfer(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		customer, err := member.Customer()
		require.NoError(t, err)

		//
		//
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bz := make([]byte, 32)
		_, err = r.Read(bz)
		require.NoError(t, err)

		key := fmt.Sprintf("REDACTED", member.Label)
		item := fmt.Sprintf("REDACTED", bz)
		tx := kinds.Tx(fmt.Sprintf("REDACTED", key, item))

		_, err = customer.MulticastTransferAlign(ctx, tx)
		require.NoError(t, err)

		digest := tx.Digest()
		waitTime := 30 * time.Second
		require.Eventuallyf(t, func() bool {
			transferReply, err := customer.Tx(ctx, digest, false)
			return err == nil && bytes.Equal(transferReply.Tx, tx)
		}, waitTime, time.Second,
			"REDACTED", waitTime,
		)

		//
		if member.Style == e2e.StyleRapid {
			return
		}

		ifaceReply, err := customer.IfaceInquire(ctx, "REDACTED", []byte(key))
		require.NoError(t, err)
		assert.Equal(t, key, string(ifaceReply.Reply.Key))
		assert.Equal(t, item, string(ifaceReply.Reply.Item))
	})
}

func Verifyapp_Ballotadditions(t *testing.T) {
	verifyMember(t, func(t *testing.T, member e2e.Member) {
		customer, err := member.Customer()
		require.NoError(t, err)
		details, err := customer.IfaceDetails(ctx)
		require.NoError(t, err)

		//
		reply, err := customer.IfaceInquire(ctx, "REDACTED", []byte("REDACTED"))
		require.NoError(t, err)

		//
		//
		if member.Verifychain.BallotPluginsActivateLevel != 0 &&
			details.Reply.FinalLedgerLevel > member.Verifychain.BallotPluginsActivateLevel {

			segments := bytes.Split(reply.Reply.Item, []byte("REDACTED"))
			require.Len(t, segments, 2)
			extensionTotal, err := strconv.Atoi(string(segments[0]))
			require.NoError(t, err)
			require.GreaterOrEqual(t, extensionTotal, 0)
		}
	})
}
