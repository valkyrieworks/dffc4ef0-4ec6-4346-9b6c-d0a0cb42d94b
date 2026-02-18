//

package align

import "sync"

//
type Lock struct {
	sync.Lock
}

//
type ReadwriteLock struct {
	sync.ReadwriteLock
}
