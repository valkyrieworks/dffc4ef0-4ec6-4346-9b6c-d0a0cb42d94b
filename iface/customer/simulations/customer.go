//

package simulations

import (
	context "context"

	abciend "github.com/valkyrieworks/iface/customer"

	log "github.com/valkyrieworks/utils/log"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/iface/kinds"
)

//
type Customer struct {
	mock.Emulate
}

//
func (_m *Customer) ExecuteMirrorSegment(_a0 context.Context, _a1 *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
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
func (_m *Customer) InspectTransfer(_a0 context.Context, _a1 *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInspectTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) *kinds.ReplyInspectTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInspectTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InspectTransferAsync(_a0 context.Context, _a1 *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *abciend.RequestOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) (*abciend.RequestOutput, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInspectTransfer) *abciend.RequestOutput); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abciend.RequestOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInspectTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Endorse(_a0 context.Context, _a1 *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEndorse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEndorse) (*kinds.ReplyEndorse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEndorse) *kinds.ReplyEndorse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEndorse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryEndorse) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Replicate(_a0 context.Context, _a1 string) (*kinds.ReplyReverberate, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyReverberate
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*kinds.ReplyReverberate, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *kinds.ReplyReverberate); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyReverberate)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Fault() error {
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
func (_m *Customer) ExpandBallot(_a0 context.Context, _a1 *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyExpandBallot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryExpandBallot) *kinds.ReplyExpandBallot); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyExpandBallot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryExpandBallot) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) CompleteLedger(_a0 context.Context, _a1 *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyCompleteLedger
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryCompleteLedger) *kinds.ReplyCompleteLedger); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyCompleteLedger)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryCompleteLedger) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Purge(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) Details(_a0 context.Context, _a1 *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) (*kinds.ReplyDetails, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryDetails) *kinds.ReplyDetails); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyDetails)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryDetails) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InitSeries(_a0 context.Context, _a1 *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInitSeries
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInitSeries) *kinds.ReplyInitSeries); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInitSeries)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInitSeries) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) EmbedTransfer(_a0 context.Context, _a1 *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyEmbedTransfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryEmbedTransfer) *kinds.ReplyEmbedTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyEmbedTransfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryEmbedTransfer) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IsActive() bool {
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
func (_m *Customer) CatalogMirrors(_a0 context.Context, _a1 *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
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
func (_m *Customer) ImportMirrorSegment(_a0 context.Context, _a1 *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
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
func (_m *Customer) ProposalMirror(_a0 context.Context, _a1 *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
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
func (_m *Customer) OnRestore() error {
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
func (_m *Customer) OnBegin() error {
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
func (_m *Customer) OnHalt() {
	_m.Called()
}

//
func (_m *Customer) ArrangeNomination(_a0 context.Context, _a1 *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyArrangeNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryArrangeNomination) *kinds.ReplyArrangeNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyArrangeNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryArrangeNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) HandleNomination(_a0 context.Context, _a1 *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHandleNomination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHandleNomination) *kinds.ReplyHandleNomination); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHandleNomination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryHandleNomination) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Inquire(_a0 context.Context, _a1 *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyInquire
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) (*kinds.ReplyInquire, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryInquire) *kinds.ReplyInquire); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyInquire)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryInquire) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Exit() <-chan struct{} {
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
func (_m *Customer) HarvestTrans(_a0 context.Context, _a1 *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyHarvestTrans
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryHarvestTrans) *kinds.ReplyHarvestTrans); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyHarvestTrans)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryHarvestTrans) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Restore() error {
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
func (_m *Customer) AssignTracer(_a0 log.Tracer) {
	_m.Called(_a0)
}

//
func (_m *Customer) CollectionReplyCallback(_a0 abciend.Callback) {
	_m.Called(_a0)
}

//
func (_m *Customer) Begin() error {
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
func (_m *Customer) Halt() error {
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
func (_m *Customer) String() string {
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
func (_m *Customer) ValidateBallotAddition(_a0 context.Context, _a1 *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("REDACTED")
	}

	var r0 *kinds.ReplyValidateBallotAddition
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *kinds.QueryValidateBallotAddition) *kinds.ReplyValidateBallotAddition); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kinds.ReplyValidateBallotAddition)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *kinds.QueryValidateBallotAddition) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
//
func NewCustomer(t interface {
	mock.TestingT
	Sanitize(func())
}) *Customer {
	emulate := &Customer{}
	mock.Emulate.Test(t)

	t.Sanitize(func() { mock.AssertExpectations(t) })

	return emulate
}
