//

package align

import (
	deadlock "github.com/sasha-s/go-deadlock"
)

//
type Mutex struct {
	deadlock.Mutex
}

//
type RWMutex struct {
	deadlock.RWMutex
}
