package autosave

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

var positionedRecordTemplate = regexp.MustCompile("REDACTED")

const (
	fallbackCohortInspectInterval = 5000 * time.Millisecond
	fallbackHeaderExtentThreshold      = 10 * 1024 * 1024       //
	fallbackSumExtentThreshold     = 1 * 1024 * 1024 * 1024 //
	maximumRecordsTowardDiscard          = 4                      //
)

/**
e
s
.

.

/
>

.

/
e
.
.

.

/
e
e
.
h

,
.
*/
type Cluster struct {
	facility.FoundationFacility

	ID                 string
	Leading               *AutomaticRecord //
	headerArea            *bufio.Writer
	Dir                string //
	metronome             *time.Ticker
	mtx                sync.Mutex
	headerExtentThreshold      int64
	sumExtentThreshold     int64
	cohortInspectInterval time.Duration
	minimumPosition           int //
	maximumPosition           int //

	//
	//
	//
	completeHandleCounts chan struct{}

	//
	//
}

//
//
func InitiateCluster(headerRoute string, clusterChoices ...func(*Cluster)) (*Cluster, error) {
	dir, err := filepath.Abs(filepath.Dir(headerRoute))
	if err != nil {
		return nil, err
	}
	header, err := UnlockAutomaticRecord(headerRoute)
	if err != nil {
		return nil, err
	}

	g := &Cluster{
		ID:                 "REDACTED" + header.ID,
		Leading:               header,
		headerArea:            bufio.NewWriterSize(header, 4096*10),
		Dir:                dir,
		headerExtentThreshold:      fallbackHeaderExtentThreshold,
		sumExtentThreshold:     fallbackSumExtentThreshold,
		cohortInspectInterval: fallbackCohortInspectInterval,
		minimumPosition:           0,
		maximumPosition:           0,
		completeHandleCounts:   make(chan struct{}),
	}

	for _, selection := range clusterChoices {
		selection(g)
	}

	g.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", g)

	gDetails := g.fetchCohortDetails()
	g.minimumPosition, g.maximumPosition = gDetails.MinimumOrdinal, gDetails.MaximumOrdinal
	return g, nil
}

//
func ClusterInspectInterval(interval time.Duration) func(*Cluster) {
	return func(g *Cluster) {
		g.cohortInspectInterval = interval
	}
}

//
func ClusterLeadingExtentThreshold(threshold int64) func(*Cluster) {
	return func(g *Cluster) {
		g.headerExtentThreshold = threshold
	}
}

//
func CohortSumExtentThreshold(threshold int64) func(*Cluster) {
	return func(g *Cluster) {
		g.sumExtentThreshold = threshold
	}
}

//
//
func (g *Cluster) UponInitiate() error {
	g.metronome = time.NewTicker(g.cohortInspectInterval)
	go g.handleCounts()
	return nil
}

//
//
func (g *Cluster) UponHalt() {
	g.metronome.Stop()
	if err := g.PurgeAlsoChronize(); err != nil {
		g.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
//
func (g *Cluster) Pause() {
	//
	<-g.completeHandleCounts
}

//
func (g *Cluster) Shutdown() {
	if err := g.PurgeAlsoChronize(); err != nil {
		g.Tracer.Failure("REDACTED", "REDACTED", err)
	}

	g.mtx.Lock()
	_ = g.Leading.shutdownRecord()
	g.mtx.Unlock()
}

//
func (g *Cluster) HeaderExtentThreshold() int64 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.headerExtentThreshold
}

//
func (g *Cluster) SumExtentThreshold() int64 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.sumExtentThreshold
}

//
func (g *Cluster) MaximumOrdinal() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.maximumPosition
}

//
func (g *Cluster) MinimumOrdinal() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.minimumPosition
}

//
//
//
//
//
func (g *Cluster) Record(p []byte) (nn int, err error) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.headerArea.Write(p)
}

//
//
//
func (g *Cluster) PersistRow(row string) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	_, err := g.headerArea.Write([]byte(row + "REDACTED"))
	return err
}

//
func (g *Cluster) Cached() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.headerArea.Buffered()
}

//
//
func (g *Cluster) PurgeAlsoChronize() error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	err := g.headerArea.Flush()
	if err == nil {
		err = g.Leading.Chronize()
	}
	return err
}

func (g *Cluster) handleCounts() {
	defer close(g.completeHandleCounts)
	for {
		select {
		case <-g.metronome.C:
			g.inspectHeaderExtentThreshold()
			g.inspectSumExtentThreshold()
		case <-g.Exit():
			return
		}
	}
}

//
func (g *Cluster) inspectHeaderExtentThreshold() {
	threshold := g.HeaderExtentThreshold()
	if threshold == 0 {
		return
	}
	extent, err := g.Leading.Extent()
	if err != nil {
		g.Tracer.Failure("REDACTED", "REDACTED", g.Leading.Route, "REDACTED", err)
		return
	}
	if extent >= threshold {
		g.PivotRecord()
	}
}

func (g *Cluster) inspectSumExtentThreshold() {
	threshold := g.SumExtentThreshold()
	if threshold == 0 {
		return
	}

	gDetails := g.fetchCohortDetails()
	sumExtent := gDetails.SumExtent
	for i := 0; i < maximumRecordsTowardDiscard; i++ {
		ordinal := gDetails.MinimumOrdinal + i
		if sumExtent < threshold {
			return
		}
		if ordinal == gDetails.MaximumOrdinal {
			//
			g.Tracer.Failure("REDACTED", "REDACTED", g.Leading.Route)
			return
		}
		routeTowardDiscard := recordRouteForeachPosition(g.Leading.Route, ordinal, gDetails.MaximumOrdinal)
		funcDetails, err := os.Stat(routeTowardDiscard)
		if err != nil {
			g.Tracer.Failure("REDACTED", "REDACTED", routeTowardDiscard)
			continue
		}
		err = os.Remove(routeTowardDiscard)
		if err != nil {
			g.Tracer.Failure("REDACTED", "REDACTED", routeTowardDiscard)
			return
		}
		sumExtent -= funcDetails.Size()
	}
}

//
//
func (g *Cluster) PivotRecord() {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	headerRoute := g.Leading.Route

	if err := g.headerArea.Flush(); err != nil {
		panic(err)
	}

	if err := g.Leading.Chronize(); err != nil {
		panic(err)
	}

	if err := g.Leading.shutdownRecord(); err != nil {
		panic(err)
	}

	positionRoute := recordRouteForeachPosition(headerRoute, g.maximumPosition, g.maximumPosition+1)
	if err := os.Rename(headerRoute, positionRoute); err != nil {
		panic(err)
	}

	g.maximumPosition++
}

//
//
func (g *Cluster) FreshFetcher(ordinal int) (*ClusterFetcher, error) {
	r := freshCohortFetcher(g)
	if err := r.AssignOrdinal(ordinal); err != nil {
		return nil, err
	}
	return r, nil
}

//
type CohortDetails struct {
	MinimumOrdinal  int   //
	MaximumOrdinal  int   //
	SumExtent int64 //
	HeaderExtent  int64 //
}

//
func (g *Cluster) FetchCohortDetails() CohortDetails {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.fetchCohortDetails()
}

//
//
func (g *Cluster) fetchCohortDetails() CohortDetails {
	cohortPath := filepath.Dir(g.Leading.Route)
	headerFoundation := filepath.Base(g.Leading.Route)
	minimumPosition, maximumPosition := -1, -1
	var sumExtent, headerExtent int64 = 0, 0

	dir, err := os.Open(cohortPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fiz, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	//
	for _, recordDetails := range fiz {
		recordAlias := recordDetails.Name()
		recordExtent := recordDetails.Size()
		sumExtent += recordExtent

		if recordAlias == headerFoundation {
			headerExtent = recordExtent
			continue
		}

		if !strings.HasPrefix(recordAlias, headerFoundation) {
			continue
		}

		undermatch := positionedRecordTemplate.FindStringSubmatch(recordAlias)
		if len(undermatch) == 2 {
			recordPosition, err := strconv.Atoi(undermatch[1])
			if err != nil {
				panic(err)
			}
			if recordPosition > maximumPosition {
				maximumPosition = recordPosition
			}
			if minimumPosition == -1 || recordPosition < minimumPosition {
				minimumPosition = recordPosition
			}
		}
	}
	//
	if minimumPosition == -1 {
		//
		//
		minimumPosition, maximumPosition = 0, 0
	} else {
		//
		maximumPosition++
	}
	return CohortDetails{minimumPosition, maximumPosition, sumExtent, headerExtent}
}

func recordRouteForeachPosition(headerRoute string, ordinal int, maximumPosition int) string {
	if ordinal == maximumPosition {
		return headerRoute
	}
	return fmt.Sprintf("REDACTED", headerRoute, ordinal)
}

//

//
type ClusterFetcher struct {
	*Cluster
	mtx       sync.Mutex
	currentPosition  int
	currentRecord   *os.File
	currentFetcher *bufio.Reader
	currentRow   []byte
}

func freshCohortFetcher(g *Cluster) *ClusterFetcher {
	return &ClusterFetcher{
		Cluster:     g,
		currentPosition:  0,
		currentRecord:   nil,
		currentFetcher: nil,
		currentRow:   nil,
	}
}

//
func (gr *ClusterFetcher) Shutdown() error {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()

	if gr.currentFetcher != nil {
		err := gr.currentRecord.Close()
		gr.currentPosition = 0
		gr.currentFetcher = nil
		gr.currentRecord = nil
		gr.currentRow = nil
		return err
	}
	return nil
}

//
//
func (gr *ClusterFetcher) Obtain(p []byte) (n int, err error) {
	lengthParam := len(p)
	if lengthParam == 0 {
		return 0, errors.New("REDACTED")
	}

	gr.mtx.Lock()
	defer gr.mtx.Unlock()

	//
	if gr.currentFetcher == nil {
		if err = gr.unlockRecord(gr.currentPosition); err != nil {
			return 0, err
		}
	}

	//
	var nn int
	for {
		nn, err = gr.currentFetcher.Read(p[n:])
		n += nn
		switch {
		case err == io.EOF:
			if n >= lengthParam {
				return n, nil
			}
			//
			if faultone := gr.unlockRecord(gr.currentPosition + 1); faultone != nil {
				return n, faultone
			}
		case err != nil:
			return n, err
		case nn == 0: //
			return n, err
		}
	}
}

//
//
func (gr *ClusterFetcher) unlockRecord(ordinal int) error {
	//
	gr.Cluster.mtx.Lock()
	defer gr.Cluster.mtx.Unlock()

	if ordinal > gr.maximumPosition {
		return io.EOF
	}

	currentRecordRoute := recordRouteForeachPosition(gr.Leading.Route, ordinal, gr.maximumPosition)
	currentRecord, err := os.OpenFile(currentRecordRoute, os.O_RDONLY|os.O_CREATE, automaticRecordModes)
	if err != nil {
		return err
	}
	currentFetcher := bufio.NewReader(currentRecord)

	//
	if gr.currentRecord != nil {
		gr.currentRecord.Close() //
	}
	gr.currentPosition = ordinal
	gr.currentRecord = currentRecord
	gr.currentFetcher = currentFetcher
	gr.currentRow = nil
	return nil
}

//
func (gr *ClusterFetcher) CurrentPosition() int {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()
	return gr.currentPosition
}

//
//
func (gr *ClusterFetcher) AssignOrdinal(ordinal int) error {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()
	return gr.unlockRecord(ordinal)
}
