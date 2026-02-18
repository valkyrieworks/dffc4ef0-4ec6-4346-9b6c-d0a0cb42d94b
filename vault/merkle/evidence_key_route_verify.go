package merkle

import (
	//
	//
	crand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyKeyRoute(t *testing.T) {
	var route KeyRoute
	keys := make([][]byte, 10)
	alphanum := "REDACTED"

	for d := 0; d < 1e4; d++ {
		route = nil

		for i := range keys {
			enc := keyCodec(rand.Intn(int(KeyCodecMaximum)))
			keys[i] = make([]byte, rand.Uint32()%20)
			switch enc {
			case KeyCodecURL:
				for j := range keys[i] {
					keys[i][j] = alphanum[rand.Intn(len(alphanum))]
				}
			case KeyCodecHex:
				_, _ = crand.Read(keys[i])
			default:
				panic("REDACTED")
			}
			route = route.AttachKey(keys[i], enc)
		}

		res, err := KeyRouteToKeys(route.String())
		require.Nil(t, err)
		require.Equal(t, len(keys), len(res))

		for i, key := range keys {
			require.Equal(t, key, res[i])
		}
	}
}
