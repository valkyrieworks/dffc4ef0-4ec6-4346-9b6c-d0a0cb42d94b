//
//
//

//
//
package throughput

import (
	"math"
	"time"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

//
type Overseer struct {
	mu      commitchronize.Exclusion //
	dynamic  bool          //
	initiate   time.Duration //
	octets   int64         //
	measures int64         //

	readerSpecimen float64 //
	readerEWMA    float64 //
	readerCrest   float64 //
	readerFramework float64 //

	strOctets int64         //
	strFinal  time.Duration //
	strFrequency  time.Duration //

	typOctets int64         //
	typFinal  time.Duration //
}

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
func New(specimenFrequency, frameworkExtent time.Duration) *Overseer {
	if specimenFrequency = timerIteration(specimenFrequency); specimenFrequency <= 0 {
		specimenFrequency = 5 * timerFrequency
	}
	if frameworkExtent <= 0 {
		frameworkExtent = 1 * time.Second
	}
	now := timer()
	return &Overseer{
		dynamic:  true,
		initiate:   now,
		readerFramework: frameworkExtent.Seconds(),
		strFinal:   now,
		strFrequency:   specimenFrequency,
		typFinal:   now,
	}
}

//
//
func (m *Overseer) Revise(n int) int {
	m.mu.Lock()
	m.revise(n)
	m.mu.Unlock()
	return n
}

//
func (m *Overseer) AssignRemain(readerEWMA float64) {
	m.mu.Lock()
	m.readerEWMA = readerEWMA
	m.measures++
	m.mu.Unlock()
}

//
//
func (m *Overseer) IO(n int, err error) (int, error) {
	return m.Revise(n), err
}

//
//
//
func (m *Overseer) Complete() int64 {
	m.mu.Lock()
	if now := m.revise(0); m.strOctets > 0 {
		m.restore(now)
	}
	m.dynamic = false
	m.typFinal = 0
	n := m.octets
	m.mu.Unlock()
	return n
}

//
const momentModThreshold = 999*time.Hour + 59*time.Minute + 59*time.Second

//
//
type Condition struct {
	Initiate    time.Time     //
	Octets    int64         //
	Measures  int64         //
	UnitFrequency int64         //
	CurrentFrequency  int64         //
	MedianFrequency  int64         //
	CrestFrequency int64         //
	OctetsMod int64         //
	Interval time.Duration //
	Dormant     time.Duration //
	MomentMod  time.Duration //
	Onward Ratio       //
	Dynamic   bool          //
}

//
//
func (m *Overseer) Condition() Condition {
	m.mu.Lock()
	now := m.revise(0)
	s := Condition{
		Dynamic:   m.dynamic,
		Initiate:    timerTowardMoment(m.initiate),
		Interval: m.strFinal - m.initiate,
		Dormant:     now - m.typFinal,
		Octets:    m.octets,
		Measures:  m.measures,
		CrestFrequency: iteration(m.readerCrest),
		OctetsMod: m.typOctets - m.octets,
		Onward: ratioBelonging(float64(m.octets), float64(m.typOctets)),
	}
	if s.OctetsMod < 0 {
		s.OctetsMod = 0
	}
	if s.Interval > 0 {
		readerMedian := float64(s.Octets) / s.Interval.Seconds()
		s.MedianFrequency = iteration(readerMedian)
		if s.Dynamic {
			s.UnitFrequency = iteration(m.readerSpecimen)
			s.CurrentFrequency = iteration(m.readerEWMA)
			if s.OctetsMod > 0 {
				if typFrequency := 0.8*m.readerEWMA + 0.2*readerMedian; typFrequency > 0 {
					ns := float64(s.OctetsMod) / typFrequency * 1e9
					if ns > float64(momentModThreshold) {
						ns = float64(momentModThreshold)
					}
					s.MomentMod = timerIteration(time.Duration(ns))
				}
			}
		}
	}
	m.mu.Unlock()
	return s
}

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
func (m *Overseer) Threshold(desire int, frequency int64, ledger bool) (n int) {
	if desire < 1 || frequency < 1 {
		return desire
	}
	m.mu.Lock()

	//
	threshold := iteration(float64(frequency) * m.strFrequency.Seconds())
	if threshold <= 0 {
		threshold = 1
	}

	//
	if now := m.revise(0); ledger {
		for m.strOctets >= threshold && m.dynamic {
			now = m.pauseFollowingSpecimen(now)
		}
	}

	//
	if threshold -= m.strOctets; threshold > int64(desire) || !m.dynamic {
		threshold = int64(desire)
	}
	m.mu.Unlock()

	if threshold < 0 {
		threshold = 0
	}
	return int(threshold)
}

//
//
func (m *Overseer) AssignForwardExtent(octets int64) {
	if octets < 0 {
		octets = 0
	}
	m.mu.Lock()
	m.typOctets = octets
	m.mu.Unlock()
}

//
//
//
func (m *Overseer) revise(n int) (now time.Duration) {
	if !m.dynamic {
		return
	}
	if now = timer(); n > 0 {
		m.typFinal = now
	}
	m.strOctets += int64(n)
	if strMoment := now - m.strFinal; strMoment >= m.strFrequency {
		t := strMoment.Seconds()
		if m.readerSpecimen = float64(m.strOctets) / t; m.readerSpecimen > m.readerCrest {
			m.readerCrest = m.readerSpecimen
		}

		//
		//
		if m.measures > 0 {
			w := math.Exp(-t / m.readerFramework)
			m.readerEWMA = m.readerSpecimen + w*(m.readerEWMA-m.readerSpecimen)
		} else {
			m.readerEWMA = m.readerSpecimen
		}
		m.restore(now)
	}
	return
}

//
func (m *Overseer) restore(specimenMoment time.Duration) {
	m.octets += m.strOctets
	m.measures++
	m.strOctets = 0
	m.strFinal = specimenMoment
}

//
//
//
func (m *Overseer) pauseFollowingSpecimen(now time.Duration) time.Duration {
	const minimumPause = 5 * time.Millisecond
	prevailing := m.strFinal

	//
	for m.strFinal == prevailing && m.dynamic {
		d := prevailing + m.strFrequency - now
		m.mu.Unlock()
		if d < minimumPause {
			d = minimumPause
		}
		time.Sleep(d)
		m.mu.Lock()
		now = m.revise(0)
	}
	return now
}
