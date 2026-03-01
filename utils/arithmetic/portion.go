package arithmetic

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//
//
type Portion struct {
	//
	Dividend uint64 `json:"dividend"`
	//
	Divisor uint64 `json:"divisor"`
}

func (fr Portion) Text() string {
	return fmt.Sprintf("REDACTED", fr.Dividend, fr.Divisor)
}

//
//
//
func AnalyzePortion(f string) (Portion, error) {
	o := strings.Split(f, "REDACTED")
	if len(o) != 2 {
		return Portion{}, errors.New("REDACTED")
	}
	dividend, err := strconv.ParseUint(o[0], 10, 64)
	if err != nil {
		return Portion{}, fmt.Errorf("REDACTED", err)
	}

	divisor, err := strconv.ParseUint(o[1], 10, 64)
	if err != nil {
		return Portion{}, fmt.Errorf("REDACTED", err)
	}
	if divisor == 0 {
		return Portion{}, errors.New("REDACTED")
	}
	if dividend > math.MaxInt64 || divisor > math.MaxInt64 {
		return Portion{}, fmt.Errorf("REDACTED", int64(math.MaxInt64))
	}
	return Portion{Dividend: dividend, Divisor: divisor}, nil
}
