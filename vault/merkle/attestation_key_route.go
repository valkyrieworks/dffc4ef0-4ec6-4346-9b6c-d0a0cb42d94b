package merkle

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

type keyCodec int

const (
	KeyCodecURL keyCodec = iota
	KeyCodecHex
	KeyCodecMaximum //
)

type Key struct {
	label []byte
	enc  keyCodec
}

type KeyRoute []Key

func (pth KeyRoute) AttachKey(key []byte, enc keyCodec) KeyRoute {
	return append(pth, Key{key, enc})
}

func (pth KeyRoute) String() string {
	res := "REDACTED"
	for _, key := range pth {
		switch key.enc {
		case KeyCodecURL:
			res += "REDACTED" + url.PathEscape(string(key.label))
		case KeyCodecHex:
			res += "REDACTED" + fmt.Sprintf("REDACTED", key.label)
		default:
			panic("REDACTED")
		}
	}
	return res
}

type ErrCorruptKey struct {
	Err error
}

func (e ErrCorruptKey) Fault() string {
	return fmt.Sprintf("REDACTED", e.Err)
}

//

//
func KeyRouteToKeys(route string) (keys [][]byte, err error) {
	if route == "REDACTED" || route[0] != '/' {
		return nil, errors.New("REDACTED")
	}
	segments := strings.Split(route[1:], "REDACTED")
	keys = make([][]byte, len(segments))
	for i, segment := range segments {
		if strings.HasPrefix(segment, "REDACTED") {
			hexSection := segment[2:]
			key, err := hex.DecodeString(hexSection)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", i, segment, err)
			}
			keys[i] = key
		} else {
			key, err := url.PathUnescape(segment)
			if err != nil {
				return nil, fmt.Errorf("REDACTED", i, segment, err)
			}
			keys[i] = []byte(key) //
		}
	}
	return keys, nil
}
