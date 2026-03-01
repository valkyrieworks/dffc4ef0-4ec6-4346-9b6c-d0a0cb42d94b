package facility

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

var (
	//
	//
	FaultEarlierInitiated = errors.New("REDACTED")
	//
	//
	FaultEarlierHalted = errors.New("REDACTED")
	//
	//
	FaultNegationInitiated = errors.New("REDACTED")
)

//
type Facility interface {
	//
	//
	//
	Initiate() error
	UponInitiate() error

	//
	//
	//
	Halt() error
	UponHalt()

	//
	//
	Restore() error
	UponRestore() error

	//
	EqualsActive() bool

	//
	Exit() <-chan struct{}

	//
	Text() string

	//
	AssignTracer(log.Tracer)
}

/**
n
.

e
,
.

g
.

.

.

:

{
e
s
}

{
{
t
}
)
s
}

{
.
s
.
}

{
.
s
.
}
*/
type FoundationFacility struct {
	Tracer  log.Tracer
	alias    string
	initiated uint32 //
	halted uint32 //
	exit    chan struct{}

	//
	implementation Facility
}

//
func FreshFoundationFacility(tracer log.Tracer, alias string, implementation Facility) *FoundationFacility {
	if tracer == nil {
		tracer = log.FreshNooperationTracer()
	}

	return &FoundationFacility{
		Tracer: tracer,
		alias:   alias,
		exit:   make(chan struct{}),
		implementation:   implementation,
	}
}

//
func (bs *FoundationFacility) AssignTracer(l log.Tracer) {
	bs.Tracer = l
}

//
//
//
func (bs *FoundationFacility) Initiate() error {
	if atomic.CompareAndSwapUint32(&bs.initiated, 0, 1) {
		if atomic.LoadUint32(&bs.halted) == 1 {
			bs.Tracer.Failure(fmt.Sprintf("REDACTED", bs.alias),
				"REDACTED", bs.implementation)
			//
			atomic.StoreUint32(&bs.initiated, 0)
			return FaultEarlierHalted
		}
		bs.Tracer.Details("REDACTED",
			"REDACTED",
			log.FreshIdleFormat("REDACTED", bs.alias),
			"REDACTED",
			bs.implementation.Text())
		err := bs.implementation.UponInitiate()
		if err != nil {
			//
			atomic.StoreUint32(&bs.initiated, 0)
			return err
		}
		return nil
	}
	bs.Tracer.Diagnose("REDACTED",
		"REDACTED",
		log.FreshIdleFormat("REDACTED", bs.alias),
		"REDACTED",
		bs.implementation)
	return FaultEarlierInitiated
}

//
//
//
func (bs *FoundationFacility) UponInitiate() error { return nil }

//
//
func (bs *FoundationFacility) Halt() error {
	if atomic.CompareAndSwapUint32(&bs.halted, 0, 1) {
		if atomic.LoadUint32(&bs.initiated) == 0 {
			bs.Tracer.Failure(fmt.Sprintf("REDACTED", bs.alias),
				"REDACTED", bs.implementation)
			//
			atomic.StoreUint32(&bs.halted, 0)
			return FaultNegationInitiated
		}
		bs.Tracer.Details("REDACTED",
			"REDACTED",
			log.FreshIdleFormat("REDACTED", bs.alias),
			"REDACTED",
			bs.implementation)
		bs.implementation.UponHalt()
		close(bs.exit)
		return nil
	}
	bs.Tracer.Diagnose("REDACTED",
		"REDACTED",
		log.FreshIdleFormat("REDACTED", bs.alias),
		"REDACTED",
		bs.implementation)
	return FaultEarlierHalted
}

//
//
//
func (bs *FoundationFacility) UponHalt() {}

//
//
func (bs *FoundationFacility) Restore() error {
	if !atomic.CompareAndSwapUint32(&bs.halted, 1, 0) {
		bs.Tracer.Diagnose("REDACTED",
			"REDACTED",
			log.FreshIdleFormat("REDACTED", bs.alias),
			"REDACTED",
			bs.implementation)
		return fmt.Errorf("REDACTED", bs.alias)
	}

	//
	atomic.CompareAndSwapUint32(&bs.initiated, 1, 0)

	bs.exit = make(chan struct{})
	return bs.implementation.UponRestore()
}

//
func (bs *FoundationFacility) UponRestore() error {
	panic("REDACTED")
}

//
//
func (bs *FoundationFacility) EqualsActive() bool {
	return atomic.LoadUint32(&bs.initiated) == 1 && atomic.LoadUint32(&bs.halted) == 0
}

//
func (bs *FoundationFacility) Pause() {
	<-bs.exit
}

//
func (bs *FoundationFacility) Text() string {
	return bs.alias
}

//
func (bs *FoundationFacility) Exit() <-chan struct{} {
	return bs.exit
}
