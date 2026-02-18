//

package simulations

import (
	log "github.com/valkyrieworks/utils/log"
	link "github.com/valkyrieworks/p2p/link"

	mock "github.com/stretchr/testify/mock"

	net "net"

	p2p "github.com/valkyrieworks/p2p"
)

//
type Node struct {
	mock.Emulate
}

//
func (_m *Node) EndLink() error {
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
func (_m *Node) FetchDeletionErrored() bool {
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
func (_m *Node) IsOutgoing() bool {
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
func (_m *Node) IsDurable() bool {
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
func (_m *Node) IsActive() bool {
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
func (_m *Node) MemberDetails() p2p.MemberDetails {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 p2p.MemberDetails
	if rf, ok := ret.Get(0).(func() p2p.MemberDetails); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(p2p.MemberDetails)
		}
	}

	return r0
}

//
func (_m *Node) OnRestore() error {
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
func (_m *Node) OnBegin() error {
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
func (_m *Node) OnHalt() {
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
func (_m *Node) DistantAddress() net.Addr {
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
func (_m *Node) DistantIP() net.IP {
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
func (_m *Node) Transmit(_a0 p2p.Packet) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Packet) bool); ok {
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
func (_m *Node) CollectionDeletionErrored() {
	_m.Called()
}

//
func (_m *Node) SocketAddress() *p2p.NetLocation {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *p2p.NetLocation
	if rf, ok := ret.Get(0).(func() *p2p.NetLocation); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.NetLocation)
		}
	}

	return r0
}

//
func (_m *Node) Begin() error {
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
func (_m *Node) Status() link.LinkageState {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 link.LinkageState
	if rf, ok := ret.Get(0).(func() link.LinkageState); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(link.LinkageState)
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
func (_m *Node) String() string {
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
func (_m *Node) AttemptTransmit(_a0 p2p.Packet) bool {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(p2p.Packet) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
//
func NewNode(t interface {
	mock.TestingT
	Sanitize(func())
}) *Node {
	emulate := &Node{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
