package notify

import (
	"math"
	"sort"
	"sync"
	"time"

	"github.com/gofrs/uuid"
	"gonum.org/v1/gonum/stat"

	"github.com/valkyrieworks/verify/starttime/shipment"
	"github.com/valkyrieworks/kinds"
)

//
//
//
//
type LedgerDepot interface {
	Level() int64
	Root() int64
	ImportLedger(int64) *kinds.Ledger
}

//
type DataSpot struct {
	Period  time.Duration
	LedgerTime time.Time
	Digest      []byte
}

//
//
type Notify struct {
	ID                      uuid.UUID
	Ratio, Linkages, Volume uint64
	Max, Min, Avg, StandardDevelop   time.Duration

	//
	//
	//
	//
	//
	AdverseNumber int

	//
	//
	//
	All []DataSpot

	//
	sum int64
}

type Notifies struct {
	s map[uuid.UUID]Notify
	l []Notify

	//
	//
	//
	faultNumber int
}

func (rs *Notifies) Catalog() []Notify {
	return rs.l
}

func (rs *Notifies) FaultNumber() int {
	return rs.faultNumber
}

func (rs *Notifies) appendDataSpot(id uuid.UUID, l time.Duration, bt time.Time, digest []byte, links, ratio, volume uint64) {
	r, ok := rs.s[id]
	if !ok {
		r = Notify{
			Max:         0,
			Min:         math.MaxInt64,
			ID:          id,
			Linkages: links,
			Ratio:        ratio,
			Volume:        volume,
		}
		rs.s[id] = r
	}
	r.All = append(r.All, DataSpot{Period: l, LedgerTime: bt, Digest: digest})
	if l > r.Max {
		r.Max = l
	}
	if l < r.Min {
		r.Min = l
	}
	if int64(l) < 0 {
		r.AdverseNumber++
	}
	//
	//
	//
	r.sum += int64(l)
	rs.s[id] = r
}

func (rs *Notifies) computeAll() {
	rs.l = make([]Notify, 0, len(rs.s))
	for _, r := range rs.s {
		if len(r.All) == 0 {
			r.Min = 0
			rs.l = append(rs.l, r)
			continue
		}
		r.Avg = time.Duration(r.sum / int64(len(r.All)))
		r.StandardDevelop = time.Duration(int64(stat.StdDev(toFloat(r.All), nil)))
		rs.l = append(rs.l, r)
	}
	sort.Slice(rs.l, func(i, j int) bool {
		if rs.l[i].Linkages == rs.l[j].Linkages {
			return rs.l[i].Ratio < rs.l[j].Ratio
		}
		return rs.l[i].Linkages < rs.l[j].Linkages
	})
}

func (rs *Notifies) appendFault() {
	rs.faultNumber++
}

//
//
func ComposeFromLedgerDepot(s LedgerDepot) (*Notifies, error) {
	type shipmentData struct {
		id                      uuid.UUID
		l                       time.Duration
		bt                      time.Time
		digest                    []byte
		linkages, ratio, volume uint64
		err                     error
	}
	type transferData struct {
		tx kinds.Tx
		bt time.Time
	}
	notifies := &Notifies{
		s: make(map[uuid.UUID]Notify),
	}

	//
	//
	//
	//
	//
	const depositoryVolume = 16

	txc := make(chan transferData)
	pdc := make(chan shipmentData, depositoryVolume)

	wg := &sync.WaitGroup{}
	wg.Add(depositoryVolume)
	for i := 0; i < depositoryVolume; i++ {
		go func() {
			defer wg.Done()
			for b := range txc {
				p, err := shipment.FromOctets(b.tx)
				if err != nil {
					pdc <- shipmentData{err: err}
					continue
				}

				l := b.bt.Sub(p.Time.AsTime())
				idb := (*[16]byte)(p.Id)
				pdc <- shipmentData{
					l:           l,
					bt:          b.bt,
					digest:        b.tx.Digest(),
					id:          uuid.UUID(*idb),
					linkages: p.Linkages,
					ratio:        p.Ratio,
					volume:        p.Volume,
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(pdc)
	}()

	go func() {
		root, level := s.Root(), s.Level()
		previous := s.ImportLedger(root)
		for i := root + 1; i < level; i++ {
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
			cur := s.ImportLedger(i)
			for _, tx := range previous.Txs {
				txc <- transferData{tx: tx, bt: cur.Time}
			}
			previous = cur
		}
		close(txc)
	}()
	for pd := range pdc {
		if pd.err != nil {
			notifies.appendFault()
			continue
		}
		notifies.appendDataSpot(pd.id, pd.l, pd.bt, pd.digest, pd.linkages, pd.ratio, pd.volume)
	}
	notifies.computeAll()
	return notifies, nil
}

func toFloat(in []DataSpot) []float64 {
	r := make([]float64, len(in))
	for i, v := range in {
		r[i] = float64(int64(v.Period))
	}
	return r
}
