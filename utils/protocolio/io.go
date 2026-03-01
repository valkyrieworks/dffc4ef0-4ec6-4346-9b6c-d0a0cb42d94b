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
	"io"

	"github.com/cosmos/gogoproto/proto"
)

type Persistor interface {
	PersistSignal(proto.Message) (int, error)
}

type PersistTerminator interface {
	Persistor
	io.Closer
}

type Fetcher interface {
	FetchSignal(msg proto.Message) (int, error)
}

type FetchTerminator interface {
	Fetcher
	io.Closer
}

type serializer interface {
	SerializeToward(data []byte) (n int, err error)
}

func obtainExtent(v any) (int, bool) {
	if sz, ok := v.(interface {
		Extent() (n int)
	}); ok {
		return sz.Extent(), true
	} else if sz, ok := v.(interface {
		SchemaExtent() (n int)
	}); ok {
		return sz.SchemaExtent(), true
	}
	return 0, false
}

//
//
//
//
type octetFetcher struct {
	fetcher    io.Reader
	buf       []byte
	octetsFetch int //
}

func freshOctetFetcher(r io.Reader) *octetFetcher {
	return &octetFetcher{
		fetcher: r,
		buf:    make([]byte, 1),
	}
}

func (r *octetFetcher) FetchOctet() (byte, error) {
	n, err := r.fetcher.Read(r.buf)
	r.octetsFetch += n
	if err != nil {
		return 0x00, err
	}
	return r.buf[0], nil
}

func (r *octetFetcher) restoreOctetsFetch() {
	r.octetsFetch = 0
}
