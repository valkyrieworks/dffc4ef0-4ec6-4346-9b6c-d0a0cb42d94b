package kinds

import (
	"sort"
)

//

//
type AssessorRevisions []AssessorRevise

var _ sort.Interface = (AssessorRevisions)(nil)

//
//
//
//
//

func (v AssessorRevisions) Len() int {
	return len(v)
}

//
func (v AssessorRevisions) Inferior(i, j int) bool {
	return v[i].PublicToken.Contrast(v[j].PublicToken) <= 0
}

func (v AssessorRevisions) Exchange(i, j int) {
	v[i], v[j] = v[j], v[i]
}
