package daemon

import (
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/valkyrieworks/utils/log"
)

var (
	//
	//
	ErrYetLaunched = errors.New("REDACTED")
	//
	//
	ErrYetCeased = errors.New("REDACTED")
	//
	//
	ErrNegateLaunched = errors.New("REDACTED")
)

//
type Daemon interface {
	//
	//
	//
	Begin() error
	OnBegin() error

	//
	//
	//
	Halt() error
	OnHalt()

	//
	//
	Restore() error
	OnRestore() error

	//
	IsActive() bool

	//
	Exit() <-chan struct{}

	//
	String() string

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
type RootDaemon struct {
	Tracer  log.Tracer
	label    string
	launched uint32 //
	ceased uint32 //
	exit    chan struct{}

	//
	impl Daemon
}

//
func NewRootDaemon(tracer log.Tracer, label string, impl Daemon) *RootDaemon {
	if tracer == nil {
		tracer = log.NewNoopTracer()
	}

	return &RootDaemon{
		Tracer: tracer,
		label:   label,
		exit:   make(chan struct{}),
		impl:   impl,
	}
}

//
func (bs *RootDaemon) AssignTracer(l log.Tracer) {
	bs.Tracer = l
}

//
//
//
func (bs *RootDaemon) Begin() error {
	if atomic.CompareAndSwapUint32(&bs.launched, 0, 1) {
		if atomic.LoadUint32(&bs.ceased) == 1 {
			bs.Tracer.Fault(fmt.Sprintf("REDACTED", bs.label),
				"REDACTED", bs.impl)
			//
			atomic.StoreUint32(&bs.launched, 0)
			return ErrYetCeased
		}
		bs.Tracer.Details("REDACTED",
			"REDACTED",
			log.NewIdleFormat("REDACTED", bs.label),
			"REDACTED",
			bs.impl.String())
		err := bs.impl.OnBegin()
		if err != nil {
			//
			atomic.StoreUint32(&bs.launched, 0)
			return err
		}
		return nil
	}
	bs.Tracer.Diagnose("REDACTED",
		"REDACTED",
		log.NewIdleFormat("REDACTED", bs.label),
		"REDACTED",
		bs.impl)
	return ErrYetLaunched
}

//
//
//
func (bs *RootDaemon) OnBegin() error { return nil }

//
//
func (bs *RootDaemon) Halt() error {
	if atomic.CompareAndSwapUint32(&bs.ceased, 0, 1) {
		if atomic.LoadUint32(&bs.launched) == 0 {
			bs.Tracer.Fault(fmt.Sprintf("REDACTED", bs.label),
				"REDACTED", bs.impl)
			//
			atomic.StoreUint32(&bs.ceased, 0)
			return ErrNegateLaunched
		}
		bs.Tracer.Details("REDACTED",
			"REDACTED",
			log.NewIdleFormat("REDACTED", bs.label),
			"REDACTED",
			bs.impl)
		bs.impl.OnHalt()
		close(bs.exit)
		return nil
	}
	bs.Tracer.Diagnose("REDACTED",
		"REDACTED",
		log.NewIdleFormat("REDACTED", bs.label),
		"REDACTED",
		bs.impl)
	return ErrYetCeased
}

//
//
//
func (bs *RootDaemon) OnHalt() {}

//
//
func (bs *RootDaemon) Restore() error {
	if !atomic.CompareAndSwapUint32(&bs.ceased, 1, 0) {
		bs.Tracer.Diagnose("REDACTED",
			"REDACTED",
			log.NewIdleFormat("REDACTED", bs.label),
			"REDACTED",
			bs.impl)
		return fmt.Errorf("REDACTED", bs.label)
	}

	//
	atomic.CompareAndSwapUint32(&bs.launched, 1, 0)

	bs.exit = make(chan struct{})
	return bs.impl.OnRestore()
}

//
func (bs *RootDaemon) OnRestore() error {
	panic("REDACTED")
}

//
//
func (bs *RootDaemon) IsActive() bool {
	return atomic.LoadUint32(&bs.launched) == 1 && atomic.LoadUint32(&bs.ceased) == 0
}

//
func (bs *RootDaemon) Wait() {
	<-bs.exit
}

//
func (bs *RootDaemon) String() string {
	return bs.label
}

//
func (bs *RootDaemon) Exit() <-chan struct{} {
	return bs.exit
}
