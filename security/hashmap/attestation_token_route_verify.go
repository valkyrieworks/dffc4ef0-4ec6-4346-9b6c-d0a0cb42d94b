package hashmap

import (
	//
	//
	crand "crypto/rand"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyTokenRoute(t *testing.T) {
	var route TokenRoute
	tokens := make([][]byte, 10)
	alphacount := "REDACTED"

	for d := 0; d < 1e4; d++ {
		route = nil

		for i := range tokens {
			enc := tokenSerialization(rand.Intn(int(TokenSerializationMaximum)))
			tokens[i] = make([]byte, rand.Uint32()%20)
			switch enc {
			case TokenSerializationWebroute:
				for j := range tokens[i] {
					tokens[i][j] = alphacount[rand.Intn(len(alphacount))]
				}
			case TokenSerializationHexadecimal:
				_, _ = crand.Read(tokens[i])
			default:
				panic("REDACTED")
			}
			route = route.AttachToken(tokens[i], enc)
		}

		res, err := TokenRouteTowardTokens(route.Text())
		require.Nil(t, err)
		require.Equal(t, len(tokens), len(res))

		for i, key := range tokens {
			require.Equal(t, key, res[i])
		}
	}
}
