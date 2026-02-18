//
//
//

//
//
package pace

import (
	"math"
	"time"

	engineconnect "github.com/valkyrieworks/utils/align"
)

//
type Auditor struct {
	mu      engineconnect.Lock //
	enabled  bool          //
	begin   time.Duration //
	octets   int64         //
	specimens int64         //

	readerSpecimen float64 //
	readerEMA    float64 //
	readerSummit   float64 //
	readerSpan float64 //

	sOctets int64         //
	sFinal  time.Duration //
	sRatio  time.Duration //

	tOctets int64         //
	tFinal  time.Duration //
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
func New(specimenRatio, spanVolume time.Duration) *Auditor {
	if specimenRatio = timerEpoch(specimenRatio); specimenRatio <= 0 {
		specimenRatio = 5 * timerRatio
	}
	if spanVolume <= 0 {
		spanVolume = 1 * time.Second
	}
	now := timer()
	return &Auditor{
		enabled:  true,
		begin:   now,
		readerSpan: spanVolume.Seconds(),
		sFinal:   now,
		sRatio:   specimenRatio,
		tFinal:   now,
	}
}

//
//
func (m *Auditor) Modify(n int) int {
	m.mu.Lock()
	m.modify(n)
	m.mu.Unlock()
	return n
}

//
func (m *Auditor) AssignRemain(readerEMA float64) {
	m.mu.Lock()
	m.readerEMA = readerEMA
	m.specimens++
	m.mu.Unlock()
}

//
//
func (m *Auditor) IO(n int, err error) (int, error) {
	return m.Modify(n), err
}

//
//
//
func (m *Auditor) Done() int64 {
	m.mu.Lock()
	if now := m.modify(0); m.sOctets > 0 {
		m.restore(now)
	}
	m.enabled = false
	m.tFinal = 0
	n := m.octets
	m.mu.Unlock()
	return n
}

//
const timeModCeiling = 999*time.Hour + 59*time.Minute + 59*time.Second

//
//
type Status struct {
	Begin    time.Time     //
	Octets    int64         //
	Specimens  int64         //
	InstanceRatio int64         //
	CurrentRatio  int64         //
	AverageRatio  int64         //
	SummitRatio int64         //
	OctetsMod int64         //
	Period time.Duration //
	Inactive     time.Duration //
	TimeMod  time.Duration //
	Advancement Fraction       //
	Enabled   bool          //
}

//
//
func (m *Auditor) Status() Status {
	m.mu.Lock()
	now := m.modify(0)
	s := Status{
		Enabled:   m.enabled,
		Begin:    timerToTime(m.begin),
		Period: m.sFinal - m.begin,
		Inactive:     now - m.tFinal,
		Octets:    m.octets,
		Specimens:  m.specimens,
		SummitRatio: epoch(m.readerSummit),
		OctetsMod: m.tOctets - m.octets,
		Advancement: fractionOf(float64(m.octets), float64(m.tOctets)),
	}
	if s.OctetsMod < 0 {
		s.OctetsMod = 0
	}
	if s.Period > 0 {
		readerAverage := float64(s.Octets) / s.Period.Seconds()
		s.AverageRatio = epoch(readerAverage)
		if s.Enabled {
			s.InstanceRatio = epoch(m.readerSpecimen)
			s.CurrentRatio = epoch(m.readerEMA)
			if s.OctetsMod > 0 {
				if tRatio := 0.8*m.readerEMA + 0.2*readerAverage; tRatio > 0 {
					ns := float64(s.OctetsMod) / tRatio * 1e9
					if ns > float64(timeModCeiling) {
						ns = float64(timeModCeiling)
					}
					s.TimeMod = timerEpoch(time.Duration(ns))
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
func (m *Auditor) Ceiling(desire int, ratio int64, ledger bool) (n int) {
	if desire < 1 || ratio < 1 {
		return desire
	}
	m.mu.Lock()

	//
	ceiling := epoch(float64(ratio) * m.sRatio.Seconds())
	if ceiling <= 0 {
		ceiling = 1
	}

	//
	if now := m.modify(0); ledger {
		for m.sOctets >= ceiling && m.enabled {
			now = m.waitFollowingSpecimen(now)
		}
	}

	//
	if ceiling -= m.sOctets; ceiling > int64(desire) || !m.enabled {
		ceiling = int64(desire)
	}
	m.mu.Unlock()

	if ceiling < 0 {
		ceiling = 0
	}
	return int(ceiling)
}

//
//
func (m *Auditor) CollectionTransmitVolume(octets int64) {
	if octets < 0 {
		octets = 0
	}
	m.mu.Lock()
	m.tOctets = octets
	m.mu.Unlock()
}

//
//
//
func (m *Auditor) modify(n int) (now time.Duration) {
	if !m.enabled {
		return
	}
	if now = timer(); n > 0 {
		m.tFinal = now
	}
	m.sOctets += int64(n)
	if sTime := now - m.sFinal; sTime >= m.sRatio {
		t := sTime.Seconds()
		if m.readerSpecimen = float64(m.sOctets) / t; m.readerSpecimen > m.readerSummit {
			m.readerSummit = m.readerSpecimen
		}

		//
		//
		if m.specimens > 0 {
			w := math.Exp(-t / m.readerSpan)
			m.readerEMA = m.readerSpecimen + w*(m.readerEMA-m.readerSpecimen)
		} else {
			m.readerEMA = m.readerSpecimen
		}
		m.restore(now)
	}
	return
}

//
func (m *Auditor) restore(specimenTime time.Duration) {
	m.octets += m.sOctets
	m.specimens++
	m.sOctets = 0
	m.sFinal = specimenTime
}

//
//
//
func (m *Auditor) waitFollowingSpecimen(now time.Duration) time.Duration {
	const minimumWait = 5 * time.Millisecond
	ongoing := m.sFinal

	//
	for m.sFinal == ongoing && m.enabled {
		d := ongoing + m.sRatio - now
		m.mu.Unlock()
		if d < minimumWait {
			d = minimumWait
		}
		time.Sleep(d)
		m.mu.Lock()
		now = m.modify(0)
	}
	return now
}
