package privatevalue

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/edwards25519"
	cryptocode "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/serialization"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
	cryptographyproto "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/security"
	privatechema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/privatevalue"
	commitchema "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/schema/strongmind/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

var imprint = time.Date(2019, 10, 13, 16, 14, 44, 0, time.UTC)

func instanceBallot() *kinds.Ballot {
	return &kinds.Ballot{
		Kind:             commitchema.PreendorseKind,
		Altitude:           3,
		Iteration:            2,
		LedgerUUID:          kinds.LedgerUUID{Digest: tenderminthash.Sum([]byte("REDACTED")), FragmentAssignHeading: kinds.FragmentAssignHeading{Sum: 1000000, Digest: tenderminthash.Sum([]byte("REDACTED"))}},
		Timestamp:        imprint,
		AssessorLocation: security.LocatorDigest([]byte("REDACTED")),
		AssessorOrdinal:   56789,
		Addition:        []byte("REDACTED"),
	}
}

func instanceNomination() *kinds.Nomination {
	return &kinds.Nomination{
		Kind:      commitchema.AttestedSignalKind(1),
		Altitude:    3,
		Iteration:     2,
		Timestamp: imprint,
		PolicyIteration:  2,
		Notation: []byte("REDACTED"),
		LedgerUUID: kinds.LedgerUUID{
			Digest: tenderminthash.Sum([]byte("REDACTED")),
			FragmentAssignHeading: kinds.FragmentAssignHeading{
				Sum: 1000000,
				Digest:  tenderminthash.Sum([]byte("REDACTED")),
			},
		},
	}
}

//
func VerifyPrivatevalueArrays(t *testing.T) {
	pk := edwards25519.ProducePrivateTokenOriginatingCredential([]byte("REDACTED")).PublicToken()
	ppk, err := cryptocode.PublicTokenTowardSchema(pk)
	require.NoError(t, err)

	//
	ballot := instanceBallot()
	ballotschema := ballot.TowardSchema()

	//
	nomination := instanceNomination()
	nominationschema := nomination.TowardSchema()

	//
	distantFailure := &privatechema.RemoteEndorserFailure{Cipher: 1, Characterization: "REDACTED"}

	verifyScenarios := []struct {
		verifyAlias string
		msg      proto.Message
		expirationOctets string
	}{
		{"REDACTED", &privatechema.PingSolicit{}, "REDACTED"},
		{"REDACTED", &privatechema.PingReply{}, "REDACTED"},
		{"REDACTED", &privatechema.PublicTokenSolicit{}, "REDACTED"},
		{"REDACTED", &privatechema.PublicTokenReply{PublicToken: ppk, Failure: nil}, "REDACTED"},
		{"REDACTED", &privatechema.PublicTokenReply{PublicToken: cryptographyproto.CommonToken{}, Failure: distantFailure}, "REDACTED"},
		{"REDACTED", &privatechema.AttestBallotSolicit{Ballot: ballotschema}, "REDACTED"},
		{"REDACTED", &privatechema.NotatedBallotReply{Ballot: *ballotschema, Failure: nil}, "REDACTED"},
		{"REDACTED", &privatechema.NotatedBallotReply{Ballot: commitchema.Ballot{}, Failure: distantFailure}, "REDACTED"},
		{"REDACTED", &privatechema.AttestNominationSolicit{Nomination: nominationschema}, "REDACTED"},
		{"REDACTED", &privatechema.NotatedNominationReply{Nomination: *nominationschema, Failure: nil}, "REDACTED"},
		{"REDACTED", &privatechema.NotatedNominationReply{Nomination: commitchema.Nomination{}, Failure: distantFailure}, "REDACTED"},
	}

	for _, tc := range verifyScenarios {

		pm := shouldEncloseSignal(tc.msg)
		bz, err := pm.Serialize()
		require.NoError(t, err, tc.verifyAlias)

		require.Equal(t, tc.expirationOctets, hex.EncodeToString(bz), tc.verifyAlias)
	}
}
