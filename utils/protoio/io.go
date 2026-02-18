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
	"io"

	"github.com/cosmos/gogoproto/proto"
)

type Recorder interface {
	RecordMessage(proto.Message) (int, error)
}

type RecordTerminator interface {
	Recorder
	io.Closer
}

type Scanner interface {
	ScanMessage(msg proto.Message) (int, error)
}

type ScanTerminator interface {
	Scanner
	io.Closer
}

type serializer interface {
	SerializeTo(data []byte) (n int, err error)
}

func fetchVolume(v any) (int, bool) {
	if sz, ok := v.(interface {
		Volume() (n int)
	}); ok {
		return sz.Volume(), true
	} else if sz, ok := v.(interface {
		SchemaVolume() (n int)
	}); ok {
		return sz.SchemaVolume(), true
	}
	return 0, false
}

//
//
//
//
type octetScanner struct {
	scanner    io.Reader
	buf       []byte
	octetsFetch int //
}

func newOctetScanner(r io.Reader) *octetScanner {
	return &octetScanner{
		scanner: r,
		buf:    make([]byte, 1),
	}
}

func (r *octetScanner) ScanOctet() (byte, error) {
	n, err := r.scanner.Read(r.buf)
	r.octetsFetch += n
	if err != nil {
		return 0x00, err
	}
	return r.buf[0], nil
}

func (r *octetScanner) restoreOctetsScan() {
	r.octetsFetch = 0
}
