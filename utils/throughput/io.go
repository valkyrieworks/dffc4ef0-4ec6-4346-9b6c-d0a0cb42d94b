//
//
//

package throughput

import (
	"errors"
	"io"
)

//
//
var FaultThreshold = errors.New("REDACTED")

//
//
type Regulator interface {
	Complete() int64
	Condition() Condition
	AssignForwardExtent(octets int64)
	AssignThreshold(new int64) (old int64)
	AssignHalting(new bool) (old bool)
}

//
//
type Fetcher struct {
	io.Fetcher //
	*Overseer  //

	threshold int64 //
	ledger bool  //
}

//
func FreshFetcher(r io.Reader, threshold int64) *Fetcher {
	return &Fetcher{r, New(0, 0), threshold, true}
}

//
//
//
func (r *Fetcher) Obtain(p []byte) (n int, err error) {
	p = p[:r.Threshold(len(p), r.threshold, r.ledger)]
	if len(p) > 0 {
		n, err = r.IO(r.Fetcher.Read(p))
	}
	return
}

//
//
func (r *Fetcher) AssignThreshold(new int64) (old int64) {
	old, r.threshold = r.threshold, new
	return
}

//
//
//
func (r *Fetcher) AssignHalting(new bool) (old bool) {
	old, r.ledger = r.ledger, new
	return
}

//
func (r *Fetcher) Shutdown() error {
	defer r.Complete()
	if c, ok := r.Fetcher.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

//
//
type Persistor struct {
	io.Persistor //
	*Overseer  //

	threshold int64 //
	ledger bool  //
}

//
//
//
func FreshPersistor(w io.Writer, threshold int64) *Persistor {
	return &Persistor{w, New(0, 0), threshold, true}
}

//
//
//
func (w *Persistor) Record(p []byte) (n int, err error) {
	var c int
	for len(p) > 0 && err == nil {
		s := p[:w.Threshold(len(p), w.threshold, w.ledger)]
		if len(s) > 0 {
			c, err = w.IO(w.Persistor.Write(s))
		} else {
			return n, FaultThreshold
		}
		p = p[c:]
		n += c
	}
	return
}

//
//
func (w *Persistor) AssignThreshold(new int64) (old int64) {
	old, w.threshold = w.threshold, new
	return
}

//
//
//
func (w *Persistor) AssignHalting(new bool) (old bool) {
	old, w.ledger = w.ledger, new
	return
}

//
func (w *Persistor) Shutdown() error {
	defer w.Complete()
	if c, ok := w.Persistor.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
