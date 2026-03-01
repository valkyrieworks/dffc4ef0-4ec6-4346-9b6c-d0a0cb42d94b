package hashmap

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/tenderminthash"
)

func VerifyRfc6962digester(t *testing.T) {
	_, nodeDigestPath := pathsOriginatingOctetSegments([][]byte{[]byte("REDACTED")})
	terminalDigest := nodeDigestPath.Digest
	_, nodeDigestPath = pathsOriginatingOctetSegments([][]byte{{}})
	blankNodeDigest := nodeDigestPath.Digest
	_, blankDigestPath := pathsOriginatingOctetSegments([][]byte{})
	blankGraphDigest := blankDigestPath.Digest
	for _, tc := range []struct {
		description string
		got  []byte
		desire string
	}{
		//
		//
		{
			description: "REDACTED",
			desire: "REDACTED"[:tenderminthash.Extent*2],
			got:  blankGraphDigest,
		},

		//
		//
		{
			description: "REDACTED",
			desire: "REDACTED"[:tenderminthash.Extent*2],
			got:  blankNodeDigest,
		},
		//
		{
			description: "REDACTED",
			desire: "REDACTED"[:tenderminthash.Extent*2],
			got:  terminalDigest,
		},
		//
		{
			description: "REDACTED",
			desire: "REDACTED"[:tenderminthash.Extent*2],
			got:  internalDigest([]byte("REDACTED"), []byte("REDACTED")),
		},
	} {

		t.Run(tc.description, func(t *testing.T) {
			desireOctets, err := hex.DecodeString(tc.desire)
			if err != nil {
				t.Fatalf("REDACTED", tc.desire, err)
			}
			if got, desire := tc.got, desireOctets; !bytes.Equal(got, desire) {
				t.Errorf("REDACTED", got, desire)
			}
		})
	}
}

func VerifyRfc6962digesterClashes(t *testing.T) {
	//
	node1, node2 := []byte("REDACTED"), []byte("REDACTED")
	_, nodeDigestPath := pathsOriginatingOctetSegments([][]byte{node1})
	digest1 := nodeDigestPath.Digest
	_, nodeDigestPath = pathsOriginatingOctetSegments([][]byte{node2})
	digest2 := nodeDigestPath.Digest
	if bytes.Equal(digest1, digest2) {
		t.Errorf("REDACTED", digest1)
	}
	//
	_, subtractDigest1trail := pathsOriginatingOctetSegments([][]byte{digest1, digest2})
	subtractDigest1 := subtractDigest1trail.Digest
	//
	original := bytes.Join([][]byte{digest1, digest2}, nil)
	_, fabricatedDigestPath := pathsOriginatingOctetSegments([][]byte{original})
	fabricatedDigest := fabricatedDigestPath.Digest
	if bytes.Equal(subtractDigest1, fabricatedDigest) {
		t.Errorf("REDACTED")
	}
	//
	_, subtractDigest2trail := pathsOriginatingOctetSegments([][]byte{digest2, digest1})
	subtractDigest2 := subtractDigest2trail.Digest
	if bytes.Equal(subtractDigest1, subtractDigest2) {
		t.Errorf("REDACTED")
	}
}
