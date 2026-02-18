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

package protoio

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
func NewSeparatedScanner(r io.Reader, maximumVolume int) ScanTerminator {
	var terminator io.Closer
	if c, ok := r.(io.Closer); ok {
		terminator = c
	}
	return &variableintScanner{r, newOctetScanner(r), nil, maximumVolume, terminator}
}

type variableintScanner struct {
	r io.Reader
	//
	//
	//
	//
	octetScanner *octetScanner
	buf        []byte
	maximumVolume    int
	terminator     io.Closer
}

func (r *variableintScanner) ScanMessage(msg proto.Message) (int, error) {
	r.octetScanner.restoreOctetsScan()
	l, err := binary.ReadUvarint(r.octetScanner)
	n := r.octetScanner.octetsFetch
	if err != nil {
		return n, err
	}

	//
	//
	extent := int(l)
	if l >= uint64(^uint(0)>>1) || extent < 0 || n+extent < 0 {
		return n, fmt.Errorf("REDACTED", l)
	}
	if extent > r.maximumVolume {
		return n, fmt.Errorf("REDACTED", extent, r.maximumVolume)
	}

	if len(r.buf) < extent {
		r.buf = make([]byte, extent)
	}
	buf := r.buf[:extent]
	nr, err := io.ReadFull(r.r, buf)
	n += nr
	if err != nil {
		return n, err
	}
	return n, proto.Unmarshal(buf, msg)
}

func (r *variableintScanner) End() error {
	if r.terminator != nil {
		return r.terminator.Close()
	}
	return nil
}

func UnserializeSeparated(data []byte, msg proto.Message) error {
	_, err := NewSeparatedScanner(bytes.NewReader(data), len(data)).ScanMessage(msg)
	return err
}
