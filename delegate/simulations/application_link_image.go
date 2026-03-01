//

package simulations

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
)

//
type PlatformLinkImage struct {
	mock.Simulate
}

//
func (_m *PlatformLinkImage) ExecuteImageSegment(_a0 context.Context, _a1 *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExecuteImageSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExecuteImageSegment) *kinds.ReplyExecuteImageSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExecuteImageSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitExecuteImageSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *PlatformLinkImage) Failure() error {
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
func (_m *PlatformLinkImage) CollectionImages(_a0 context.Context, _a1 *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCatalogImages
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitCollectionImages) *kinds.ReplyCatalogImages); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCatalogImages)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitCollectionImages) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *PlatformLinkImage) FetchImageSegment(_a0 context.Context, _a1 *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyFetchImageSegment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitFetchImageSegment) *kinds.ReplyFetchImageSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyFetchImageSegment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitFetchImageSegment) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *PlatformLinkImage) ExtendImage(_a0 context.Context, _a1 *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExtendImage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.SolicitExtendImage) *kinds.ReplyExtendImage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExtendImage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.SolicitExtendImage) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func FreshApplicationLinkImage(t interface {
	mock.TestingT
	Sanitize(func())
}) *PlatformLinkImage {
	simulate := &PlatformLinkImage{}
	mock.Simulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return simulate
}
