package kinds

import (
	"sort"
)

//

//
type RatifierRefreshes []RatifierModify

var _ sort.Interface = (RatifierRefreshes)(nil)

//
//
//
//
//

func (v RatifierRefreshes) Len() int {
	return len(v)
}

//
func (v RatifierRefreshes) Lower(i, j int) bool {
	return v[i].PublicKey.Contrast(v[j].PublicKey) <= 0
}

func (v RatifierRefreshes) Exchange(i, j int) {
	v[i], v[j] = v[j], v[i]
}
