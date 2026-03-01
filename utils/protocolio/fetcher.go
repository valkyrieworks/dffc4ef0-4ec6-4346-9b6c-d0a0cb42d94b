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

package protocolio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/cosmos/gogoproto/proto"
)

//
//
//
//
//
func FreshSeparatedFetcher(r io.Reader, maximumExtent int) FetchTerminator {
	var terminator io.Closer
	if c, ok := r.(io.Closer); ok {
		terminator = c
	}
	return &variableintFetcher{r, freshOctetFetcher(r), nil, maximumExtent, terminator}
}

type variableintFetcher struct {
	r io.Reader
	//
	//
	//
	//
	octetFetcher *octetFetcher
	buf        []byte
	maximumExtent    int
	terminator     io.Closer
}

func (r *variableintFetcher) FetchSignal(msg proto.Message) (int, error) {
	r.octetFetcher.restoreOctetsFetch()
	l, err := binary.ReadUvarint(r.octetFetcher)
	n := r.octetFetcher.octetsFetch
	if err != nil {
		return n, err
	}

	//
	//
	magnitude := int(l)
	if l >= uint64(^uint(0)>>1) || magnitude < 0 || n+magnitude < 0 {
		return n, fmt.Errorf("REDACTED", l)
	}
	if magnitude > r.maximumExtent {
		return n, fmt.Errorf("REDACTED", magnitude, r.maximumExtent)
	}

	if len(r.buf) < magnitude {
		r.buf = make([]byte, magnitude)
	}
	buf := r.buf[:magnitude]
	nr, err := io.ReadFull(r.r, buf)
	n += nr
	if err != nil {
		return n, err
	}
	return n, proto.Unmarshal(buf, msg)
}

func (r *variableintFetcher) Shutdown() error {
	if r.terminator != nil {
		return r.terminator.Close()
	}
	return nil
}

func DecodeSeparated(data []byte, msg proto.Message) error {
	_, err := FreshSeparatedFetcher(bytes.NewReader(data), len(data)).FetchSignal(msg)
	return err
}
