package clock

import (
	"testing"
	"time"

	//

	asrt "github.com/stretchr/testify/assert"

	engineconnect "github.com/valkyrieworks/utils/align"
)

type thTally struct {
	influx chan struct{}
	mtx   engineconnect.Lock
	tally int
}

func (c *thTally) Augment() {
	c.mtx.Lock()
	c.tally++
	c.mtx.Unlock()
}

func (c *thTally) Number() int {
	c.mtx.Lock()
	val := c.tally
	c.mtx.Unlock()
	return val
}

//
//
func (c *thTally) Scan() {
	for range c.influx {
		c.Augment()
	}
}

func VerifyRegulate(verify *testing.T) {
	affirm := asrt.New(verify)

	ms := 50
	deferral := time.Duration(ms) * time.Millisecond
	longwait := time.Duration(2) * deferral
	t := NewRegulateClock("REDACTED", deferral)

	//
	c := &thTally{influx: t.Ch}
	affirm.Equal(0, c.Number())
	go c.Scan()

	//
	time.Sleep(longwait)
	affirm.Equal(0, c.Number())

	//
	t.Set()
	time.Sleep(longwait)
	affirm.Equal(1, c.Number())

	//
	for i := 0; i < 5; i++ {
		t.Set()
	}
	time.Sleep(longwait)
	affirm.Equal(2, c.Number())

	//
	//
	//
	brief := time.Duration(ms/5) * time.Millisecond
	for i := 0; i < 13; i++ {
		t.Set()
		time.Sleep(brief)
	}
	time.Sleep(longwait)
	affirm.LessOrEqual(5, c.Number())

	close(t.Ch)
}
