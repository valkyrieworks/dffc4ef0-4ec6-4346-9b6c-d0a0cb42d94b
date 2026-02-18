package autofile

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

	"github.com/valkyrieworks/utils/daemon"
)

var catalogedEntryTemplate = regexp.MustCompile("REDACTED")

const (
	standardClusterInspectPeriod = 5000 * time.Millisecond
	standardFrontVolumeCeiling      = 10 * 1024 * 1024       //
	standardSumVolumeCeiling     = 1 * 1024 * 1024 * 1024 //
	maximumEntriesToDelete          = 4                      //
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
	daemon.RootDaemon

	ID                 string
	Front               *AutomaticEntry //
	frontImage            *bufio.Writer
	Dir                string //
	timer             *time.Ticker
	mtx                sync.Mutex
	frontVolumeCeiling      int64
	sumVolumeCeiling     int64
	clusterInspectPeriod time.Duration
	minimumOrdinal           int //
	maximumOrdinal           int //

	//
	//
	//
	doneHandleTicks chan struct{}

	//
	//
}

//
//
func AccessCluster(frontRoute string, clusterSettings ...func(*Cluster)) (*Cluster, error) {
	dir, err := filepath.Abs(filepath.Dir(frontRoute))
	if err != nil {
		return nil, err
	}
	front, err := AccessAutomaticEntry(frontRoute)
	if err != nil {
		return nil, err
	}

	g := &Cluster{
		ID:                 "REDACTED" + front.ID,
		Front:               front,
		frontImage:            bufio.NewWriterSize(front, 4096*10),
		Dir:                dir,
		frontVolumeCeiling:      standardFrontVolumeCeiling,
		sumVolumeCeiling:     standardSumVolumeCeiling,
		clusterInspectPeriod: standardClusterInspectPeriod,
		minimumOrdinal:           0,
		maximumOrdinal:           0,
		doneHandleTicks:   make(chan struct{}),
	}

	for _, setting := range clusterSettings {
		setting(g)
	}

	g.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", g)

	gDetails := g.readerClusterDetails()
	g.minimumOrdinal, g.maximumOrdinal = gDetails.MinimumOrdinal, gDetails.MaximumOrdinal
	return g, nil
}

//
func ClusterInspectPeriod(period time.Duration) func(*Cluster) {
	return func(g *Cluster) {
		g.clusterInspectPeriod = period
	}
}

//
func ClusterFrontVolumeCeiling(ceiling int64) func(*Cluster) {
	return func(g *Cluster) {
		g.frontVolumeCeiling = ceiling
	}
}

//
func ClusterSumVolumeCeiling(ceiling int64) func(*Cluster) {
	return func(g *Cluster) {
		g.sumVolumeCeiling = ceiling
	}
}

//
//
func (g *Cluster) OnBegin() error {
	g.timer = time.NewTicker(g.clusterInspectPeriod)
	go g.handleTicks()
	return nil
}

//
//
func (g *Cluster) OnHalt() {
	g.timer.Stop()
	if err := g.PurgeAndAlign(); err != nil {
		g.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}

//
//
func (g *Cluster) Wait() {
	//
	<-g.doneHandleTicks
}

//
func (g *Cluster) End() {
	if err := g.PurgeAndAlign(); err != nil {
		g.Tracer.Fault("REDACTED", "REDACTED", err)
	}

	g.mtx.Lock()
	_ = g.Front.endEntry()
	g.mtx.Unlock()
}

//
func (g *Cluster) FrontVolumeCeiling() int64 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.frontVolumeCeiling
}

//
func (g *Cluster) SumVolumeCeiling() int64 {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.sumVolumeCeiling
}

//
func (g *Cluster) MaximumOrdinal() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.maximumOrdinal
}

//
func (g *Cluster) MinimumOrdinal() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.minimumOrdinal
}

//
//
//
//
//
func (g *Cluster) Record(p []byte) (nn int, err error) {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.frontImage.Write(p)
}

//
//
//
func (g *Cluster) RecordRow(row string) error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	_, err := g.frontImage.Write([]byte(row + "REDACTED"))
	return err
}

//
func (g *Cluster) Cached() int {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.frontImage.Buffered()
}

//
//
func (g *Cluster) PurgeAndAlign() error {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	err := g.frontImage.Flush()
	if err == nil {
		err = g.Front.Align()
	}
	return err
}

func (g *Cluster) handleTicks() {
	defer close(g.doneHandleTicks)
	for {
		select {
		case <-g.timer.C:
			g.inspectFrontVolumeCeiling()
			g.inspectSumVolumeCeiling()
		case <-g.Exit():
			return
		}
	}
}

//
func (g *Cluster) inspectFrontVolumeCeiling() {
	ceiling := g.FrontVolumeCeiling()
	if ceiling == 0 {
		return
	}
	volume, err := g.Front.Volume()
	if err != nil {
		g.Tracer.Fault("REDACTED", "REDACTED", g.Front.Route, "REDACTED", err)
		return
	}
	if volume >= ceiling {
		g.SpinEntry()
	}
}

func (g *Cluster) inspectSumVolumeCeiling() {
	ceiling := g.SumVolumeCeiling()
	if ceiling == 0 {
		return
	}

	gDetails := g.readerClusterDetails()
	sumVolume := gDetails.SumVolume
	for i := 0; i < maximumEntriesToDelete; i++ {
		ordinal := gDetails.MinimumOrdinal + i
		if sumVolume < ceiling {
			return
		}
		if ordinal == gDetails.MaximumOrdinal {
			//
			g.Tracer.Fault("REDACTED", "REDACTED", g.Front.Route)
			return
		}
		routeToDelete := entryRouteForOrdinal(g.Front.Route, ordinal, gDetails.MaximumOrdinal)
		fDetails, err := os.Stat(routeToDelete)
		if err != nil {
			g.Tracer.Fault("REDACTED", "REDACTED", routeToDelete)
			continue
		}
		err = os.Remove(routeToDelete)
		if err != nil {
			g.Tracer.Fault("REDACTED", "REDACTED", routeToDelete)
			return
		}
		sumVolume -= fDetails.Size()
	}
}

//
//
func (g *Cluster) SpinEntry() {
	g.mtx.Lock()
	defer g.mtx.Unlock()

	frontRoute := g.Front.Route

	if err := g.frontImage.Flush(); err != nil {
		panic(err)
	}

	if err := g.Front.Align(); err != nil {
		panic(err)
	}

	if err := g.Front.endEntry(); err != nil {
		panic(err)
	}

	ordinalRoute := entryRouteForOrdinal(frontRoute, g.maximumOrdinal, g.maximumOrdinal+1)
	if err := os.Rename(frontRoute, ordinalRoute); err != nil {
		panic(err)
	}

	g.maximumOrdinal++
}

//
//
func (g *Cluster) NewScanner(ordinal int) (*ClusterScanner, error) {
	r := newClusterScanner(g)
	if err := r.AssignOrdinal(ordinal); err != nil {
		return nil, err
	}
	return r, nil
}

//
type ClusterDetails struct {
	MinimumOrdinal  int   //
	MaximumOrdinal  int   //
	SumVolume int64 //
	FrontVolume  int64 //
}

//
func (g *Cluster) ReaderClusterDetails() ClusterDetails {
	g.mtx.Lock()
	defer g.mtx.Unlock()
	return g.readerClusterDetails()
}

//
//
func (g *Cluster) readerClusterDetails() ClusterDetails {
	clusterFolder := filepath.Dir(g.Front.Route)
	frontRoot := filepath.Base(g.Front.Route)
	minimumOrdinal, maximumOrdinal := -1, -1
	var sumVolume, frontVolume int64 = 0, 0

	dir, err := os.Open(clusterFolder)
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	fiz, err := dir.Readdir(0)
	if err != nil {
		panic(err)
	}

	//
	for _, entryDetails := range fiz {
		entryLabel := entryDetails.Name()
		entryVolume := entryDetails.Size()
		sumVolume += entryVolume

		if entryLabel == frontRoot {
			frontVolume = entryVolume
			continue
		}

		if !strings.HasPrefix(entryLabel, frontRoot) {
			continue
		}

		subsection := catalogedEntryTemplate.FindStringSubmatch(entryLabel)
		if len(subsection) == 2 {
			entryOrdinal, err := strconv.Atoi(subsection[1])
			if err != nil {
				panic(err)
			}
			if entryOrdinal > maximumOrdinal {
				maximumOrdinal = entryOrdinal
			}
			if minimumOrdinal == -1 || entryOrdinal < minimumOrdinal {
				minimumOrdinal = entryOrdinal
			}
		}
	}
	//
	if minimumOrdinal == -1 {
		//
		//
		minimumOrdinal, maximumOrdinal = 0, 0
	} else {
		//
		maximumOrdinal++
	}
	return ClusterDetails{minimumOrdinal, maximumOrdinal, sumVolume, frontVolume}
}

func entryRouteForOrdinal(frontRoute string, ordinal int, maximumOrdinal int) string {
	if ordinal == maximumOrdinal {
		return frontRoute
	}
	return fmt.Sprintf("REDACTED", frontRoute, ordinal)
}

//

//
type ClusterScanner struct {
	*Cluster
	mtx       sync.Mutex
	currentOrdinal  int
	currentEntry   *os.File
	currentScanner *bufio.Reader
	currentRow   []byte
}

func newClusterScanner(g *Cluster) *ClusterScanner {
	return &ClusterScanner{
		Cluster:     g,
		currentOrdinal:  0,
		currentEntry:   nil,
		currentScanner: nil,
		currentRow:   nil,
	}
}

//
func (gr *ClusterScanner) End() error {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()

	if gr.currentScanner != nil {
		err := gr.currentEntry.Close()
		gr.currentOrdinal = 0
		gr.currentScanner = nil
		gr.currentEntry = nil
		gr.currentRow = nil
		return err
	}
	return nil
}

//
//
func (gr *ClusterScanner) Scan(p []byte) (n int, err error) {
	sizeP := len(p)
	if sizeP == 0 {
		return 0, errors.New("REDACTED")
	}

	gr.mtx.Lock()
	defer gr.mtx.Unlock()

	//
	if gr.currentScanner == nil {
		if err = gr.accessEntry(gr.currentOrdinal); err != nil {
			return 0, err
		}
	}

	//
	var nn int
	for {
		nn, err = gr.currentScanner.Read(p[n:])
		n += nn
		switch {
		case err == io.EOF:
			if n >= sizeP {
				return n, nil
			}
			//
			if fault1 := gr.accessEntry(gr.currentOrdinal + 1); fault1 != nil {
				return n, fault1
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
func (gr *ClusterScanner) accessEntry(ordinal int) error {
	//
	gr.Cluster.mtx.Lock()
	defer gr.Cluster.mtx.Unlock()

	if ordinal > gr.maximumOrdinal {
		return io.EOF
	}

	currentEntryRoute := entryRouteForOrdinal(gr.Front.Route, ordinal, gr.maximumOrdinal)
	currentEntry, err := os.OpenFile(currentEntryRoute, os.O_RDONLY|os.O_CREATE, automaticEntryModes)
	if err != nil {
		return err
	}
	currentScanner := bufio.NewReader(currentEntry)

	//
	if gr.currentEntry != nil {
		gr.currentEntry.Close() //
	}
	gr.currentOrdinal = ordinal
	gr.currentEntry = currentEntry
	gr.currentScanner = currentScanner
	gr.currentRow = nil
	return nil
}

//
func (gr *ClusterScanner) CurrentOrdinal() int {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()
	return gr.currentOrdinal
}

//
//
func (gr *ClusterScanner) AssignOrdinal(ordinal int) error {
	gr.mtx.Lock()
	defer gr.mtx.Unlock()
	return gr.accessEntry(ordinal)
}
