package summary

import (
	"math"
	"sort"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"gonum.org/v1/gonum/stat"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/verify/fetchmoment/content"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

//
//
//
//
type LedgerDepot interface {
	Altitude() int64
	Foundation() int64
	FetchLedger(int64) *kinds.Ledger
}

//
type DataMark struct {
	Interval  time.Duration
	LedgerMoment time.Time
	Digest      []byte
}

//
//
type Summary struct {
	ID                      uuid.UUID
	Frequency, Linkages, Extent uint64
	Max, Min, Avg, StandardDevelop   time.Duration

	//
	//
	//
	//
	//
	AdverseTally int

	//
	//
	//
	All []DataMark

	//
	sum int64
}

type Summaries struct {
	s map[uuid.UUID]Summary
	l []Summary

	//
	//
	//
	failureTally int
}

func (rs *Summaries) Catalog() []Summary {
	return rs.l
}

func (rs *Summaries) FailureTally() int {
	return rs.failureTally
}

func (rs *Summaries) appendDataMark(id uuid.UUID, l time.Duration, bt time.Time, digest []byte, links, frequency, extent uint64) {
	r, ok := rs.s[id]
	if !ok {
		r = Summary{
			Max:         0,
			Min:         math.MaxInt64,
			ID:          id,
			Linkages: links,
			Frequency:        frequency,
			Extent:        extent,
		}
		rs.s[id] = r
	}
	r.All = append(r.All, DataMark{Interval: l, LedgerMoment: bt, Digest: digest})
	if l > r.Max {
		r.Max = l
	}
	if l < r.Min {
		r.Min = l
	}
	if int64(l) < 0 {
		r.AdverseTally++
	}
	//
	//
	//
	r.sum += int64(l)
	rs.s[id] = r
}

func (rs *Summaries) computeEvery() {
	rs.l = make([]Summary, 0, len(rs.s))
	for _, r := range rs.s {
		if len(r.All) == 0 {
			r.Min = 0
			rs.l = append(rs.l, r)
			continue
		}
		r.Avg = time.Duration(r.sum / int64(len(r.All)))
		r.StandardDevelop = time.Duration(int64(stat.StdDev(towardDecimal(r.All), nil)))
		rs.l = append(rs.l, r)
	}
	sort.Slice(rs.l, func(i, j int) bool {
		if rs.l[i].Linkages == rs.l[j].Linkages {
			return rs.l[i].Frequency < rs.l[j].Frequency
		}
		return rs.l[i].Linkages < rs.l[j].Linkages
	})
}

func (rs *Summaries) appendFailure() {
	rs.failureTally++
}

//
//
func ComposeOriginatingLedgerDepot(s LedgerDepot) (*Summaries, error) {
	type contentData struct {
		id                      uuid.UUID
		l                       time.Duration
		bt                      time.Time
		digest                    []byte
		linkages, frequency, extent uint64
		err                     error
	}
	type transferData struct {
		tx kinds.Tx
		bt time.Time
	}
	summaries := &Summaries{
		s: make(map[uuid.UUID]Summary),
	}

	//
	//
	//
	//
	//
	const hubExtent = 16

	txc := make(chan transferData)
	pdc := make(chan contentData, hubExtent)

	wg := &sync.WaitGroup{}
	wg.Add(hubExtent)
	for i := 0; i < hubExtent; i++ {
		go func() {
			defer wg.Done()
			for b := range txc {
				p, err := content.OriginatingOctets(b.tx)
				if err != nil {
					pdc <- contentData{err: err}
					continue
				}

				l := b.bt.Sub(p.Moment.AsTime())
				idb := (*[16]byte)(p.Id)
				pdc <- contentData{
					l:           l,
					bt:          b.bt,
					digest:        b.tx.Digest(),
					id:          uuid.UUID(*idb),
					linkages: p.Linkages,
					frequency:        p.Frequency,
					extent:        p.Extent,
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(pdc)
	}()

	go func() {
		foundation, altitude := s.Foundation(), s.Altitude()
		previous := s.FetchLedger(foundation)
		for i := foundation + 1; i < altitude; i++ {
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
			cur := s.FetchLedger(i)
			for _, tx := range previous.Txs {
				txc <- transferData{tx: tx, bt: cur.Moment}
			}
			previous = cur
		}
		close(txc)
	}()
	for pd := range pdc {
		if pd.err != nil {
			summaries.appendFailure()
			continue
		}
		summaries.appendDataMark(pd.id, pd.l, pd.bt, pd.digest, pd.linkages, pd.frequency, pd.extent)
	}
	summaries.computeEvery()
	return summaries, nil
}

func towardDecimal(in []DataMark) []float64 {
	r := make([]float64, len(in))
	for i, v := range in {
		r[i] = float64(int64(v.Interval))
	}
	return r
}
