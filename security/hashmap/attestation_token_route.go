package hashmap

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

/**

n
r
.
.
)

s
s
:

)
)
}

e
l
.

)
)
)
)
"

.

e
r
X
e
.

e
n
.

*/

type tokenSerialization int

const (
	TokenSerializationWebroute tokenSerialization = iota
	TokenSerializationHexadecimal
	TokenSerializationMaximum //
)

type Key struct {
	alias []byte
	enc  tokenSerialization
}

type TokenRoute []Key

func (pth TokenRoute) AttachToken(key []byte, enc tokenSerialization) TokenRoute {
	return append(pth, Key{key, enc})
}

func (pth TokenRoute) Text() string {
	res := "REDACTED"
	for _, key := range pth {
		switch key.enc {
		case TokenSerializationWebroute:
			res += "REDACTED" + url.PathEscape(string(key.alias))
		case TokenSerializationHexadecimal:
			res += "REDACTED" + fmt.Sprintf("REDACTED", key.alias)
		default:
			panic("REDACTED")
		}
	}
	return res
}

type FaultUnfitToken struct {
	Err error
}

func (e FaultUnfitToken) Failure() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

//

//
func TokenRouteTowardTokens(route string) (tokens [][]byte, err error) {
	if route == "REDACTED" || route[0] != '/' {
		return nil, errors.New("REDACTED")
	}
	fragments := strings.Split(route[1:], "REDACTED")
	tokens = make([][]byte, len(fragments))
	for i, fragment := range fragments {
		if strings.HasPrefix(fragment, "REDACTED") {
			hexadecimalFragment := fragment[2:]
			key, err := hex.DecodeString(hexadecimalFragment)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", i, fragment, err)
			}
			tokens[i] = key
		} else {
			key, err := url.PathUnescape(fragment)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", i, fragment, err)
			}
			tokens[i] = []byte(key) //
		}
	}
	return tokens, nil
}
