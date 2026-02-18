package merkle

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

	"github.com/valkyrieworks/vault/comethash"
)

func VerifyRfc6962digester(t *testing.T) {
	_, elementDigestPath := pathsFromOctetSegments([][]byte{[]byte("REDACTED")})
	elementDigest := elementDigestPath.Digest
	_, elementDigestPath = pathsFromOctetSegments([][]byte{{}})
	emptyElementDigest := elementDigestPath.Digest
	_, emptyDigestPath := pathsFromOctetSegments([][]byte{})
	emptyGraphDigest := emptyDigestPath.Digest
	for _, tc := range []struct {
		note string
		got  []byte
		desire string
	}{
		//
		//
		{
			note: "REDACTED",
			desire: "REDACTED"[:comethash.Volume*2],
			got:  emptyGraphDigest,
		},

		//
		//
		{
			note: "REDACTED",
			desire: "REDACTED"[:comethash.Volume*2],
			got:  emptyElementDigest,
		},
		//
		{
			note: "REDACTED",
			desire: "REDACTED"[:comethash.Volume*2],
			got:  elementDigest,
		},
		//
		{
			note: "REDACTED",
			desire: "REDACTED"[:comethash.Volume*2],
			got:  deeperDigest([]byte("REDACTED"), []byte("REDACTED")),
		},
	} {

		t.Run(tc.note, func(t *testing.T) {
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
	element1, element2 := []byte("REDACTED"), []byte("REDACTED")
	_, elementDigestPath := pathsFromOctetSegments([][]byte{element1})
	digest1 := elementDigestPath.Digest
	_, elementDigestPath = pathsFromOctetSegments([][]byte{element2})
	digest2 := elementDigestPath.Digest
	if bytes.Equal(digest1, digest2) {
		t.Errorf("REDACTED", digest1)
	}
	//
	_, subtractDigest1trail := pathsFromOctetSegments([][]byte{digest1, digest2})
	subtractDigest1 := subtractDigest1trail.Digest
	//
	preimage := bytes.Join([][]byte{digest1, digest2}, nil)
	_, falsifiedDigestPath := pathsFromOctetSegments([][]byte{preimage})
	falsifiedDigest := falsifiedDigestPath.Digest
	if bytes.Equal(subtractDigest1, falsifiedDigest) {
		t.Errorf("REDACTED")
	}
	//
	_, subtractDigest2trail := pathsFromOctetSegments([][]byte{digest2, digest1})
	subtractDigest2 := subtractDigest2trail.Digest
	if bytes.Equal(subtractDigest1, subtractDigest2) {
		t.Errorf("REDACTED")
	}
}
