package verify

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/edition"
)

const (
	FallbackVerifySuccessionUUID = "REDACTED"
)

var FallbackVerifyMoment = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

func UnpredictableLocator() []byte {
	return security.CHARArbitraryOctets(security.LocatorExtent)
}

func ArbitraryDigest() []byte {
	return security.CHARArbitraryOctets(tenderminthash.Extent)
}

func CreateLedgerUUID() kinds.LedgerUUID {
	return CreateLedgerUUIDUsingDigest(ArbitraryDigest())
}

func CreateLedgerUUIDUsingDigest(digest []byte) kinds.LedgerUUID {
	return kinds.LedgerUUID{
		Digest: digest,
		FragmentAssignHeading: kinds.FragmentAssignHeading{
			Sum: 100,
			Digest:  ArbitraryDigest(),
		},
	}
}

//
//
func CreateHeadline(t *testing.T, h *kinds.Heading) *kinds.Heading {
	t.Helper()
	if h.Edition.Ledger == 0 {
		h.Edition.Ledger = edition.LedgerScheme
	}
	if h.Altitude == 0 {
		h.Altitude = 1
	}
	if h.FinalLedgerUUID.EqualsNull() {
		h.FinalLedgerUUID = CreateLedgerUUID()
	}
	if h.SuccessionUUID == "REDACTED" {
		h.SuccessionUUID = FallbackVerifySuccessionUUID
	}
	if len(h.FinalEndorseDigest) == 0 {
		h.FinalEndorseDigest = ArbitraryDigest()
	}
	if len(h.DataDigest) == 0 {
		h.DataDigest = ArbitraryDigest()
	}
	if len(h.AssessorsDigest) == 0 {
		h.AssessorsDigest = ArbitraryDigest()
	}
	if len(h.FollowingAssessorsDigest) == 0 {
		h.FollowingAssessorsDigest = ArbitraryDigest()
	}
	if len(h.AgreementDigest) == 0 {
		h.AgreementDigest = ArbitraryDigest()
	}
	if len(h.PlatformDigest) == 0 {
		h.PlatformDigest = ArbitraryDigest()
	}
	if len(h.FinalOutcomesDigest) == 0 {
		h.FinalOutcomesDigest = ArbitraryDigest()
	}
	if len(h.ProofDigest) == 0 {
		h.ProofDigest = ArbitraryDigest()
	}
	if len(h.NominatorLocation) == 0 {
		h.NominatorLocation = UnpredictableLocator()
	}

	require.NoError(t, h.CertifyFundamental())

	return h
}
