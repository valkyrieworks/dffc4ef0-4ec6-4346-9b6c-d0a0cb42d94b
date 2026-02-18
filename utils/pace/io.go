//
//
//

package pace

import (
	"errors"
	"io"
)

//
//
var ErrCeiling = errors.New("REDACTED")

//
//
type Regulator interface {
	Done() int64
	Status() Status
	CollectionTransmitVolume(octets int64)
	CollectionCeiling(new int64) (old int64)
	CollectionHalting(new bool) (old bool)
}

//
//
type Scanner struct {
	io.Scanner //
	*Auditor  //

	ceiling int64 //
	ledger bool  //
}

//
func NewScanner(r io.Reader, ceiling int64) *Scanner {
	return &Scanner{r, New(0, 0), ceiling, true}
}

//
//
//
func (r *Scanner) Scan(p []byte) (n int, err error) {
	p = p[:r.Ceiling(len(p), r.ceiling, r.ledger)]
	if len(p) > 0 {
		n, err = r.IO(r.Scanner.Read(p))
	}
	return
}

//
//
func (r *Scanner) CollectionCeiling(new int64) (old int64) {
	old, r.ceiling = r.ceiling, new
	return
}

//
//
//
func (r *Scanner) CollectionHalting(new bool) (old bool) {
	old, r.ledger = r.ledger, new
	return
}

//
func (r *Scanner) End() error {
	defer r.Done()
	if c, ok := r.Scanner.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

//
//
type Recorder struct {
	io.Recorder //
	*Auditor  //

	ceiling int64 //
	ledger bool  //
}

//
//
//
func NewRecorder(w io.Writer, ceiling int64) *Recorder {
	return &Recorder{w, New(0, 0), ceiling, true}
}

//
//
//
func (w *Recorder) Record(p []byte) (n int, err error) {
	var c int
	for len(p) > 0 && err == nil {
		s := p[:w.Ceiling(len(p), w.ceiling, w.ledger)]
		if len(s) > 0 {
			c, err = w.IO(w.Recorder.Write(s))
		} else {
			return n, ErrCeiling
		}
		p = p[c:]
		n += c
	}
	return
}

//
//
func (w *Recorder) CollectionCeiling(new int64) (old int64) {
	old, w.ceiling = w.ceiling, new
	return
}

//
//
//
func (w *Recorder) CollectionHalting(new bool) (old bool) {
	old, w.ledger = w.ledger, new
	return
}

//
func (w *Recorder) End() error {
	defer w.Done()
	if c, ok := w.Recorder.(io.Closer); ok {
		return c.Close()
	}
	return nil
}
