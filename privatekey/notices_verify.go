package privatekey

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/vault"
	"github.com/valkyrieworks/vault/ed25519"
	cryptocode "github.com/valkyrieworks/vault/codec"
	"github.com/valkyrieworks/vault/comethash"
	cryptography "github.com/valkyrieworks/schema/consensuscore/vault"
	privateproto "github.com/valkyrieworks/schema/consensuscore/privatekey"
	engineproto "github.com/valkyrieworks/schema/consensuscore/kinds"
	"github.com/valkyrieworks/kinds"
)

var imprint = time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC)

func instanceBallot() *kinds.Ballot {
	return &kinds.Ballot{
		Kind:             engineproto.PreendorseKind,
		Level:           3,
		Cycle:            2,
		LedgerUID:          kinds.LedgerUID{Digest: comethash.Sum([]byte("REDACTED")), SegmentAssignHeading: kinds.SegmentAssignHeading{Sum: 1000000, Digest: comethash.Sum([]byte("REDACTED"))}},
		Timestamp:        imprint,
		RatifierLocation: vault.LocationDigest([]byte("REDACTED")),
		RatifierOrdinal:   56789,
		Addition:        []byte("REDACTED"),
	}
}

func instanceNomination() *kinds.Nomination {
	return &kinds.Nomination{
		Kind:      engineproto.AttestedMessageKind(1),
		Level:    3,
		Cycle:     2,
		Timestamp: imprint,
		POLDuration:  2,
		Autograph: []byte("REDACTED"),
		LedgerUID: kinds.LedgerUID{
			Digest: comethash.Sum([]byte("REDACTED")),
			SegmentAssignHeading: kinds.SegmentAssignHeading{
				Sum: 1000000,
				Digest:  comethash.Sum([]byte("REDACTED")),
			},
		},
	}
}

//
func VerifyPrivatekeyArrays(t *testing.T) {
	pk := ed25519.GeneratePrivateKeyFromPrivatekey([]byte("REDACTED")).PublicKey()
	ppk, err := cryptocode.PublicKeyToSchema(pk)
	require.NoError(t, err)

	//
	ballot := instanceBallot()
	ballotpb := ballot.ToSchema()

	//
	nomination := instanceNomination()
	nominationpb := nomination.ToSchema()

	//
	distantFault := &privateproto.DistantNotaryFault{Code: 1, Summary: "REDACTED"}

	verifyScenarios := []struct {
		verifyLabel string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &privateproto.PingQuery{}, "REDACTED"},
		{"REDACTED", &privateproto.PingAnswer{}, "REDACTED"},
		{"REDACTED", &privateproto.PublicKeyQuery{}, "REDACTED"},
		{"REDACTED", &privateproto.PublicKeyAnswer{PublicKey: ppk, Fault: nil}, "REDACTED"},
		{"REDACTED", &privateproto.PublicKeyAnswer{PublicKey: cryptography.PublicKey{}, Fault: distantFault}, "REDACTED"},
		{"REDACTED", &privateproto.AttestBallotQuery{Ballot: ballotpb}, "REDACTED"},
		{"REDACTED", &privateproto.AttestedBallotAnswer{Ballot: *ballotpb, Fault: nil}, "REDACTED"},
		{"REDACTED", &privateproto.AttestedBallotAnswer{Ballot: engineproto.Ballot{}, Fault: distantFault}, "REDACTED"},
		{"REDACTED", &privateproto.AttestNominationQuery{Nomination: nominationpb}, "REDACTED"},
		{"REDACTED", &privateproto.AttestedNominationAnswer{Nomination: *nominationpb, Fault: nil}, "REDACTED"},
		{"REDACTED", &privateproto.AttestedNominationAnswer{Nomination: engineproto.Nomination{}, Fault: distantFault}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		pm := shouldEncloseMessage(tc.msg)
		bz, err := pm.Serialize()
		require.NoError(t, err, tc.verifyLabel)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyLabel)
	}
}
