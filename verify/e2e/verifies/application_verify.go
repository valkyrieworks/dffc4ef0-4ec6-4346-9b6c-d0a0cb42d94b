package end2end_typ_test

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	e2e "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/e2e/pkg"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
func Testcase_Primarystatus(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		if len(peer.Simnet.PrimaryStatus) == 0 {
			return
		}

		customer, err := peer.Customer()
		require.NoError(t, err)
		for k, v := range peer.Simnet.PrimaryStatus {
			reply, err := customer.IfaceInquire(ctx, "REDACTED", []byte(k))
			require.NoError(t, err)
			assert.Equal(t, k, string(reply.Reply.Key))
			assert.Equal(t, v, string(reply.Reply.Datum))
		}
	})
}

//
//
func Testcase_Digest(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		customer, err := peer.Customer()
		require.NoError(t, err)

		details, err := customer.IfaceDetails(ctx)
		require.NoError(t, err)
		require.NotEmpty(t, details.Reply.FinalLedgerPlatformDigest, "REDACTED")

		//
		solicitedAltitude := details.Reply.FinalLedgerAltitude + 1

		require.Eventually(t, func() bool {
			condition, err := customer.Condition(ctx)
			require.NoError(t, err)
			require.NotZero(t, condition.ChronizeDetails.NewestLedgerAltitude)
			return condition.ChronizeDetails.NewestLedgerAltitude >= solicitedAltitude
		}, 5*time.Second, 500*time.Millisecond)

		ledger, err := customer.Ledger(ctx, &solicitedAltitude)
		require.NoError(t, err)
		require.Equal(t,
			fmt.Sprintf("REDACTED", details.Reply.FinalLedgerPlatformDigest),
			fmt.Sprintf("REDACTED", ledger.Ledger.PlatformDigest.Octets()),
			"REDACTED")
	})
}

//
func Testcase_Transfer(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		customer, err := peer.Customer()
		require.NoError(t, err)

		//
		//
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bz := make([]byte, 32)
		_, err = r.Read(bz)
		require.NoError(t, err)

		key := fmt.Sprintf("REDACTED", peer.Alias)
		datum := fmt.Sprintf("REDACTED", bz)
		tx := kinds.Tx(fmt.Sprintf("REDACTED", key, datum))

		_, err = customer.MulticastTransferChronize(ctx, tx)
		require.NoError(t, err)

		digest := tx.Digest()
		awaitMoment := 30 * time.Second
		require.Eventuallyf(t, func() bool {
			transferAnswer, err := customer.Tx(ctx, digest, false)
			return err == nil && bytes.Equal(transferAnswer.Tx, tx)
		}, awaitMoment, time.Second,
			"REDACTED", awaitMoment,
		)

		//
		if peer.Style == e2e.StyleAgile {
			return
		}

		ifaceReply, err := customer.IfaceInquire(ctx, "REDACTED", []byte(key))
		require.NoError(t, err)
		assert.Equal(t, key, string(ifaceReply.Reply.Key))
		assert.Equal(t, datum, string(ifaceReply.Reply.Datum))
	})
}

func Testcase_Voteadditions(t *testing.T) {
	verifyPeer(t, func(t *testing.T, peer e2e.Peer) {
		customer, err := peer.Customer()
		require.NoError(t, err)
		details, err := customer.IfaceDetails(ctx)
		require.NoError(t, err)

		//
		reply, err := customer.IfaceInquire(ctx, "REDACTED", []byte("REDACTED"))
		require.NoError(t, err)

		//
		//
		if peer.Simnet.BallotAdditionsActivateAltitude != 0 &&
			details.Reply.FinalLedgerAltitude > peer.Simnet.BallotAdditionsActivateAltitude {

			fragments := bytes.Split(reply.Reply.Datum, []byte("REDACTED"))
			require.Len(t, fragments, 2)
			addnTotal, err := strconv.Atoi(string(fragments[0]))
			require.NoError(t, err)
			require.GreaterOrEqual(t, addnTotal, 0)
		}
	})
}
