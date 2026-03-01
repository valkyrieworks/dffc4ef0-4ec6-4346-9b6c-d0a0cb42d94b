//

package simulations

import (
	log "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	link "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p/link"

	mock "github.com/stretchr/testify/mock"

	net "net"

	p2p "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/p2p"
)

//
type Node struct {
	mock.Simulate
}

//
func (_m *Node) ShutdownLink() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) PurgeHalt() {
	_m.Called()
}

//
func (_m *Node) Get(_a0 string) any {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 any
	if rf, ok := ret.Get(0).(func(string) any); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(any)
		}
	}

	return r0
}

//
func (_m *Node) ObtainDeletionUnsuccessful() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Node) ID() p2p.ID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 p2p.ID
	if rf, ok := ret.Get(0).(func() p2p.ID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(p2p.ID)
	}

	return r0
}

//
func (_m *Node) EqualsOutgoing() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Node) EqualsEnduring() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Node) EqualsActive() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Node) PeerDetails() p2p.PeerDetails {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 p2p.PeerDetails
	if rf, ok := ret.Get(0).(func() p2p.PeerDetails); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.PeerDetails)
		}
	}

	return r0
}

//
func (_m *Node) UponRestore() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) UponInitiate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) UponHalt() {
	_m.Called()
}

//
func (_m *Node) Exit() <-chan struct{} {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

//
func (_m *Node) DistantLocation() net.Addr {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 net.Addr
	if rf, ok := ret.Get(0).(func() net.Addr); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.Addr)
		}
	}

	return r0
}

//
func (_m *Node) DistantINET() net.IP {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 net.IP
	if rf, ok := ret.Get(0).(func() net.IP); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(net.IP)
		}
	}

	return r0
}

//
func (_m *Node) Restore() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) Transmit(_a0 p2p.Wrapper) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Wrapper) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Node) Set(_a0 string, _a1 any) {
	_m.Called(_a0, _a1)
}

//
func (_m *Node) AssignTracer(_a0 log.Tracer) {
	_m.Called(_a0)
}

//
func (_m *Node) AssignDeletionUnsuccessful() {
	_m.Called()
}

//
func (_m *Node) PortLocation() *p2p.NetworkLocator {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *p2p.NetworkLocator
	if rf, ok := ret.Get(0).(func() *p2p.NetworkLocator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.NetworkLocator)
		}
	}

	return r0
}

//
func (_m *Node) Initiate() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) Condition() link.LinkageCondition {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 link.LinkageCondition
	if rf, ok := ret.Get(0).(func() link.LinkageCondition); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(link.LinkageCondition)
	}

	return r0
}

//
func (_m *Node) Halt() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Node) Text() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *Node) AttemptTransmit(_a0 p2p.Wrapper) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Wrapper) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
//
func FreshNode(t interface {
	mock.TestingT
	Sanitize(func())
}) *Node {
	simulate := &Node{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
