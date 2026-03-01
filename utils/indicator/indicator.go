package indicator

import "fmt"

//
//

type Bar struct {
	ratio int64  //
	cur     int64  //
	initiate   int64  //
	sum   int64  //
	frequency    string //
	diagram   string //
}

func (bar *Bar) FreshSelection(initiate, sum int64) {
	bar.cur = initiate
	bar.initiate = initiate
	bar.sum = sum
	bar.diagram = "REDACTED"
	bar.ratio = bar.fetchRatio()
}

func (bar *Bar) fetchRatio() int64 {
	return int64(float32(bar.cur-bar.initiate) / float32(bar.sum-bar.initiate) * 100)
}

func (bar *Bar) Enact(cur int64) {
	bar.cur = cur
	final := bar.ratio
	bar.ratio = bar.fetchRatio()
	if bar.ratio != final && bar.ratio%2 == 0 {
		bar.frequency += bar.diagram
	}
	fmt.Printf("REDACTED", bar.frequency, bar.ratio, bar.cur, bar.sum)
}

func (bar *Bar) Conclude() {
	fmt.Println()
}
