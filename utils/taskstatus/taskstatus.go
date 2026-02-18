package taskstatus

import "fmt"

//
//

type Bar struct {
	fraction int64  //
	cur     int64  //
	begin   int64  //
	sum   int64  //
	ratio    string //
	chart   string //
}

func (bar *Bar) NewSetting(begin, sum int64) {
	bar.cur = begin
	bar.begin = begin
	bar.sum = sum
	bar.chart = "REDACTED"
	bar.fraction = bar.fetchFraction()
}

func (bar *Bar) fetchFraction() int64 {
	return int64(float32(bar.cur-bar.begin) / float32(bar.sum-bar.begin) * 100)
}

func (bar *Bar) Simulate(cur int64) {
	bar.cur = cur
	final := bar.fraction
	bar.fraction = bar.fetchFraction()
	if bar.fraction != final && bar.fraction%2 == 0 {
		bar.ratio += bar.chart
	}
	fmt.Printf("REDACTED", bar.ratio, bar.fraction, bar.cur, bar.sum)
}

func (bar *Bar) Conclude() {
	fmt.Println()
}
