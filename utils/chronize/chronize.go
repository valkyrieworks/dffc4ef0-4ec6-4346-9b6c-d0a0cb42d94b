//

package chronize

import "sync"

//
type Exclusion struct {
	sync.Exclusion
}

//
type ReadwriteExclusion struct {
	sync.ReadwriteExclusion
}
