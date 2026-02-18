//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type ApplicationLinkMirror struct {
	mock.Emulate
}

//
func (_m *ApplicationLinkMirror) ExecuteMirrorSegment(_a0 context.Context, _a1 *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExecuteMirrorSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExecuteMirrorSegment) *kinds.ReplyExecuteMirrorSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExecuteMirrorSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryExecuteMirrorSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkMirror) Fault() error {
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
func (_m *ApplicationLinkMirror) CatalogMirrors(_a0 context.Context, _a1 *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCatalogMirrors
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCatalogMirrors) *kinds.ReplyCatalogMirrors); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCatalogMirrors)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryCatalogMirrors) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkMirror) ImportMirrorSegment(_a0 context.Context, _a1 *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyImportMirrorSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryImportMirrorSegment) *kinds.ReplyImportMirrorSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyImportMirrorSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryImportMirrorSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *ApplicationLinkMirror) ProposalMirror(_a0 context.Context, _a1 *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyProposalMirror
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryProposalMirror) *kinds.ReplyProposalMirror); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyProposalMirror)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryProposalMirror) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewApplicationLinkMirror(t interface {
	mock.TestingT
	Sanitize(func())
}) *ApplicationLinkMirror {
	emulate := &ApplicationLinkMirror{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
