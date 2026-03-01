//
//
//

package throughput

import (
	"bytes"
	"testing"
	"time"
)

const (
	_period50ms  = 50 * time.Millisecond
	_period100ms = 100 * time.Millisecond
	_period200ms = 200 * time.Millisecond
	_period300ms = 300 * time.Millisecond
	_period400ms = 400 * time.Millisecond
	_period500ms = 500 * time.Millisecond
)

func followingCondition(m *Overseer) Condition {
	measures := m.measures
	for i := 0; i < 30; i++ {
		if s := m.Condition(); s.Measures != measures {
			return s
		}
		time.Sleep(5 * time.Millisecond)
	}
	return m.Condition()
}

func VerifyFetcher(t *testing.T) {
	in := make([]byte, 100)
	for i := range in {
		in[i] = byte(i)
	}
	b := make([]byte, 100)
	r := FreshFetcher(bytes.NewReader(in), 100)
	initiate := time.Now()

	//
	_ = Regulator(r)

	//
	if n, err := r.Obtain(b); n != 10 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(initiate); rt > _period50ms {
		t.Fatalf("REDACTED", rt)
	}

	//
	r.AssignHalting(false)
	if n, err := r.Obtain(b); n != 0 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(initiate); rt > _period50ms {
		t.Fatalf("REDACTED", rt)
	}

	condition := [6]Condition{0: r.Condition()} //

	//
	r.AssignHalting(true)
	if n, err := r.Obtain(b[10:]); n != 10 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(initiate); rt < _period100ms {
		t.Fatalf("REDACTED", rt)
	}

	condition[1] = r.Condition()            //
	condition[2] = followingCondition(r.Overseer) //
	condition[3] = followingCondition(r.Overseer) //

	if n := r.Complete(); n != 20 {
		t.Fatalf("REDACTED", n)
	}

	condition[4] = r.Condition()
	condition[5] = followingCondition(r.Overseer) //
	initiate = condition[0].Initiate

	//
	desire := []Condition{
		{initiate, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, true},
		{initiate, 10, 1, 100, 100, 100, 100, 0, _period100ms, 0, 0, 0, true},
		{initiate, 20, 2, 100, 100, 100, 100, 0, _period200ms, _period100ms, 0, 0, true},
		{initiate, 20, 3, 0, 90, 67, 100, 0, _period300ms, _period200ms, 0, 0, true},
		{initiate, 20, 3, 0, 0, 67, 100, 0, _period300ms, 0, 0, 0, false},
		{initiate, 20, 3, 0, 0, 67, 100, 0, _period300ms, 0, 0, 0, false},
	}
	for i, s := range condition {

		if !statesExistEquivalent(&s, &desire[i]) {
			t.Errorf("REDACTED", i, desire[i], s)
		}
	}
	if !bytes.Equal(b[:20], in[:20]) {
		t.Errorf("REDACTED")
	}
}

func VerifyPersistor(t *testing.T) {
	b := make([]byte, 100)
	for i := range b {
		b[i] = byte(i)
	}
	w := FreshPersistor(&bytes.Buffer{}, 200)
	initiate := time.Now()

	//
	_ = Regulator(w)

	//
	w.AssignHalting(false)
	if n, err := w.Record(b); n != 20 || err != FaultThreshold {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(initiate); rt > _period50ms {
		t.Fatalf("REDACTED", rt)
	}

	//
	w.AssignHalting(true)
	if n, err := w.Record(b[20:]); n != 80 || err != nil {
		t.Fatalf("REDACTED", n, err)
	} else if rt := time.Since(initiate); rt < _period300ms {
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

	w.AssignForwardExtent(100)
	condition := []Condition{w.Condition(), followingCondition(w.Overseer)}
	initiate = condition[0].Initiate

	//
	desire := []Condition{
		{initiate, 80, 4, 200, 200, 200, 200, 20, _period400ms, 0, _period100ms, 80000, true},
		{initiate, 100, 5, 200, 200, 200, 200, 0, _period500ms, _period100ms, 0, 100000, true},
	}

	for i, s := range condition {

		if !statesExistEquivalent(&s, &desire[i]) {
			t.Errorf("REDACTED", i, desire[i], s)
		}
	}
	if !bytes.Equal(b, w.Persistor.(*bytes.Buffer).Bytes()) {
		t.Errorf("REDACTED")
	}
}

const (
	maximumVarianceForeachInterval       = 50 * time.Millisecond
	maximumVarianceForeachFrequency     int64 = 50
)

//
//
//
//
func statesExistEquivalent(s1 *Condition, s2 *Condition) bool {
	if s1.Dynamic == s2.Dynamic &&
		s1.Initiate.Equal(s2.Initiate) &&
		intervalsExistEquivalent(s1.Interval, s2.Interval, maximumVarianceForeachInterval) &&
		s1.Dormant == s2.Dormant &&
		s1.Octets == s2.Octets &&
		s1.Measures == s2.Measures &&
		metricsExistEquivalent(s1.UnitFrequency, s2.UnitFrequency, maximumVarianceForeachFrequency) &&
		metricsExistEquivalent(s1.CurrentFrequency, s2.CurrentFrequency, maximumVarianceForeachFrequency) &&
		metricsExistEquivalent(s1.MedianFrequency, s2.MedianFrequency, maximumVarianceForeachFrequency) &&
		metricsExistEquivalent(s1.CrestFrequency, s2.CrestFrequency, maximumVarianceForeachFrequency) &&
		s1.OctetsMod == s2.OctetsMod &&
		intervalsExistEquivalent(s1.MomentMod, s2.MomentMod, maximumVarianceForeachInterval) &&
		s1.Onward == s2.Onward {
		return true
	}
	return false
}

func intervalsExistEquivalent(d1 time.Duration, d2 time.Duration, maximumVariance time.Duration) bool {
	return d2-d1 <= maximumVariance
}

func metricsExistEquivalent(r1 int64, r2 int64, maximumVariance int64) bool {
	sub := r1 - r2
	if sub < 0 {
		sub = -sub
	}
	if sub <= maximumVariance {
		return true
	}
	return false
}
