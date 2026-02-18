package verify

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/comethash"
	"github.com/valkyrieworks/kinds"
	"github.com/valkyrieworks/release"
)

const (
	StandardVerifyLedgerUID = "REDACTED"
)

var StandardVerifyTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

func ArbitraryLocation() []byte {
	return vault.CRandomOctets(vault.LocationVolume)
}

func ArbitraryDigest() []byte {
	return vault.CRandomOctets(comethash.Volume)
}

func CreateLedgerUID() kinds.LedgerUID {
	return CreateLedgerUIDWithDigest(ArbitraryDigest())
}

func CreateLedgerUIDWithDigest(digest []byte) kinds.LedgerUID {
	return kinds.LedgerUID{
		Digest: digest,
		SegmentAssignHeading: kinds.SegmentAssignHeading{
			Sum: 100,
			Digest:  ArbitraryDigest(),
		},
	}
}

//
//
func CreateHeading(t *testing.T, h *kinds.Heading) *kinds.Heading {
	t.Helper()
	if h.Release.Ledger == 0 {
		h.Release.Ledger = release.LedgerProtocol
	}
	if h.Level == 0 {
		h.Level = 1
	}
	if h.FinalLedgerUID.IsNil() {
		h.FinalLedgerUID = CreateLedgerUID()
	}
	if h.LedgerUID == "REDACTED" {
		h.LedgerUID = StandardVerifyLedgerUID
	}
	if len(h.FinalEndorseDigest) == 0 {
		h.FinalEndorseDigest = ArbitraryDigest()
	}
	if len(h.DataDigest) == 0 {
		h.DataDigest = ArbitraryDigest()
	}
	if len(h.RatifiersDigest) == 0 {
		h.RatifiersDigest = ArbitraryDigest()
	}
	if len(h.FollowingRatifiersDigest) == 0 {
		h.FollowingRatifiersDigest = ArbitraryDigest()
	}
	if len(h.AgreementDigest) == 0 {
		h.AgreementDigest = ArbitraryDigest()
	}
	if len(h.ApplicationDigest) == 0 {
		h.ApplicationDigest = ArbitraryDigest()
	}
	if len(h.FinalOutcomesDigest) == 0 {
		h.FinalOutcomesDigest = ArbitraryDigest()
	}
	if len(h.ProofDigest) == 0 {
		h.ProofDigest = ArbitraryDigest()
	}
	if len(h.RecommenderLocation) == 0 {
		h.RecommenderLocation = ArbitraryLocation()
	}

	require.NoError(t, h.CertifySimple())

	return h
}
