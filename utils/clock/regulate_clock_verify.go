package clock

import (
	"testing"
	"time"

	//

	asrt "github.com/stretchr/testify/assert"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

type threadTally struct {
	influx chan struct{}
	mtx   commitchronize.Exclusion
	tally int
}

func (c *threadTally) Advance() {
	c.mtx.Lock()
	c.tally++
	c.mtx.Unlock()
}

func (c *threadTally) Tally() int {
	c.mtx.Lock()
	val := c.tally
	c.mtx.Unlock()
	return val
}

//
//
func (c *threadTally) Obtain() {
	for range c.influx {
		c.Advance()
	}
}

func VerifyRegulate(verify *testing.T) {
	affirm := asrt.New(verify)

	ms := 50
	deferral := time.Duration(ms) * time.Millisecond
	extendedpause := time.Duration(2) * deferral
	t := FreshRegulateClock("REDACTED", deferral)

	//
	c := &threadTally{influx: t.Ch}
	affirm.Equal(0, c.Tally())
	go c.Obtain()

	//
	time.Sleep(extendedpause)
	affirm.Equal(0, c.Tally())

	//
	t.Set()
	time.Sleep(extendedpause)
	affirm.Equal(1, c.Tally())

	//
	for i := 0; i < 5; i++ {
		t.Set()
	}
	time.Sleep(extendedpause)
	affirm.Equal(2, c.Tally())

	//
	//
	//
	brief := time.Duration(ms/5) * time.Millisecond
	for i := 0; i < 13; i++ {
		t.Set()
		time.Sleep(brief)
	}
	time.Sleep(extendedpause)
	affirm.LessOrEqual(5, c.Tally())

	close(t.Ch)
}
