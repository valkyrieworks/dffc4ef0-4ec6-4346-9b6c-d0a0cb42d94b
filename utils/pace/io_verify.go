//
//
//

package pace

import (
	"bytes"
	"testing"
	"time"
)

const (
	_50moment  = 50 * time.Millisecond
	_100moment = 100 * time.Millisecond
	_200moment = 200 * time.Millisecond
	_300moment = 300 * time.Millisecond
	_400moment = 400 * time.Millisecond
	_500moment = 500 * time.Millisecond
)

func followingState(m *Auditor) Status {
	specimens := m.specimens
	for i := 0; i < 30; i++ {
		if s := m.Status(); s.Specimens != specimens {
			return s
		}
		time.Sleep(5 * time.Millisecond)
	}
	return m.Status()
}

func VerifyScanner(t *testing.T) {
	in := make([]byte, 100)
	for i := range in {
		in[i] = byte(i)
	}
	b := make([]byte, 100)
	r := NewScanner(bytes.NewReader(in), 100)
	begin := time.Now()

	//
	_ = Regulator(r)

	//
	if n, err := r.Scan(b); n != 10 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(begin); rt > _50moment {
		t.Fatalf("REDACTED", rt)
	}

	//
	r.CollectionHalting(false)
	if n, err := r.Scan(b); n != 0 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(begin); rt > _50moment {
		t.Fatalf("REDACTED", rt)
	}

	state := [6]Status{0: r.Status()} //

	//
	r.CollectionHalting(true)
	if n, err := r.Scan(b[10:]); n != 10 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(begin); rt < _100moment {
		t.Fatalf("REDACTED", rt)
	}

	state[1] = r.Status()            //
	state[2] = followingState(r.Auditor) //
	state[3] = followingState(r.Auditor) //

	if n := r.Done(); n != 20 {
		t.Fatalf("REDACTED", n)
	}

	state[4] = r.Status()
	state[5] = followingState(r.Auditor) //
	begin = state[0].Begin

	//
	desire := []Status{
		{begin, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, true},
		{begin, 10, 1, 100, 100, 100, 100, 0, _100moment, 0, 0, 0, true},
		{begin, 20, 2, 100, 100, 100, 100, 0, _200moment, _100moment, 0, 0, true},
		{begin, 20, 3, 0, 90, 67, 100, 0, _300moment, _200moment, 0, 0, true},
		{begin, 20, 3, 0, 0, 67, 100, 0, _300moment, 0, 0, 0, false},
		{begin, 20, 3, 0, 0, 67, 100, 0, _300moment, 0, 0, 0, false},
	}
	for i, s := range state {

		if !statesAreEquivalent(&s, &desire[i]) {
			t.Errorf("REDACTED", i, desire[i], s)
		}
	}
	if !bytes.Equal(b[:20], in[:20]) {
		t.Errorf("REDACTED")
	}
}

func VerifyRecorder(t *testing.T) {
	b := make([]byte, 100)
	for i := range b {
		b[i] = byte(i)
	}
	w := NewRecorder(&bytes.Buffer{}, 200)
	begin := time.Now()

	//
	_ = Regulator(w)

	//
	w.CollectionHalting(false)
	if n, err := w.Record(b); n != 20 || err != ErrCeiling {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(begin); rt > _50moment {
		t.Fatalf("REDACTED", rt)
	}

	//
	w.CollectionHalting(true)
	if n, err := w.Record(b[20:]); n != 80 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(begin); rt < _300moment {
		//
		//
		//
		//
		//
		//
		//
		//
		t.Fatalf("REDACTED", rt)
	}

	w.CollectionTransmitVolume(100)
	state := []Status{w.Status(), followingState(w.Auditor)}
	begin = state[0].Begin

	//
	desire := []Status{
		{begin, 80, 4, 200, 200, 200, 200, 20, _400moment, 0, _100moment, 80000, true},
		{begin, 100, 5, 200, 200, 200, 200, 0, _500moment, _100moment, 0, 100000, true},
	}

	for i, s := range state {

		if !statesAreEquivalent(&s, &desire[i]) {
			t.Errorf("REDACTED", i, desire[i], s)
		}
	}
	if !bytes.Equal(b, w.Recorder.(*bytes.Buffer).Bytes()) {
		t.Errorf("REDACTED")
	}
}

const (
	maximumVarianceForPeriod       = 50 * time.Millisecond
	maximumVarianceForRatio     int64 = 50
)

//
//
//
//
func statesAreEquivalent(s1 *Status, s2 *Status) bool {
	if s1.Enabled == s2.Enabled &&
		s1.Begin.Equal(s2.Begin) &&
		periodsAreEquivalent(s1.Period, s2.Period, maximumVarianceForPeriod) &&
		s1.Inactive == s2.Inactive &&
		s1.Octets == s2.Octets &&
		s1.Specimens == s2.Specimens &&
		frequenciesAreEquivalent(s1.InstanceRatio, s2.InstanceRatio, maximumVarianceForRatio) &&
		frequenciesAreEquivalent(s1.CurrentRatio, s2.CurrentRatio, maximumVarianceForRatio) &&
		frequenciesAreEquivalent(s1.AverageRatio, s2.AverageRatio, maximumVarianceForRatio) &&
		frequenciesAreEquivalent(s1.SummitRatio, s2.SummitRatio, maximumVarianceForRatio) &&
		s1.OctetsMod == s2.OctetsMod &&
		periodsAreEquivalent(s1.TimeMod, s2.TimeMod, maximumVarianceForPeriod) &&
		s1.Advancement == s2.Advancement {
		return true
	}
	return false
}

func periodsAreEquivalent(d1 time.Duration, d2 time.Duration, maximumVariance time.Duration) bool {
	return d2-d1 <= maximumVariance
}

func frequenciesAreEquivalent(r1 int64, r2 int64, maximumVariance int64) bool {
	sub := r1 - r2
	if sub < 0 {
		sub = -sub
	}
	if sub <= maximumVariance {
		return true
	}
	return false
}
