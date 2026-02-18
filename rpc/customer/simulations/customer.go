//

package simulations

import (
	octets "github.com/valkyrieworks/utils/octets"
	customer "github.com/valkyrieworks/rpc/customer"

	context "context"

	basetypes "github.com/valkyrieworks/rpc/core/kinds"

	log "github.com/valkyrieworks/utils/log"

	mock "github.com/stretchr/testify/mock"

	kinds "github.com/valkyrieworks/kinds"
)

//
type Customer struct {
	mock.Emulate
}

//
func (_m *Customer) IfaceDetails(_a0 context.Context) (*basetypes.OutcomeIfaceDetails, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeIfaceDetails
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeIfaceDetails); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IfaceInquire(ctx context.Context, route string, data octets.HexOctets) (*basetypes.OutcomeIfaceInquire, error) {
	ret := _m.Called(ctx, route, data)

	var r0 *basetypes.OutcomeIfaceInquire
	if rf, ok := ret.Get(0).(func(context.Context, string, octets.HexOctets) *basetypes.OutcomeIfaceInquire); ok {
		r0 = rf(ctx, route, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceInquire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, octets.HexOctets) error); ok {
		r1 = rf(ctx, route, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IfaceInquireWithSettings(ctx context.Context, route string, data octets.HexOctets, opts customer.IfaceInquireSettings) (*basetypes.OutcomeIfaceInquire, error) {
	ret := _m.Called(ctx, route, data, opts)

	var r0 *basetypes.OutcomeIfaceInquire
	if rf, ok := ret.Get(0).(func(context.Context, string, octets.HexOctets, customer.IfaceInquireSettings) *basetypes.OutcomeIfaceInquire); ok {
		r0 = rf(ctx, route, data, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeIfaceInquire)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, octets.HexOctets, customer.IfaceInquireSettings) error); ok {
		r1 = rf(ctx, route, data, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Ledger(ctx context.Context, level *int64) (*basetypes.OutcomeLedger, error) {
	ret := _m.Called(ctx, level)

	var r0 *basetypes.OutcomeLedger
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeLedger); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerByDigest(ctx context.Context, digest []byte) (*basetypes.OutcomeLedger, error) {
	ret := _m.Called(ctx, digest)

	var r0 *basetypes.OutcomeLedger
	if rf, ok := ret.Get(0).(func(context.Context, []byte) *basetypes.OutcomeLedger); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedger)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerOutcomes(ctx context.Context, level *int64) (*basetypes.OutcomeLedgerOutcomes, error) {
	ret := _m.Called(ctx, level)

	var r0 *basetypes.OutcomeLedgerOutcomes
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeLedgerOutcomes); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerOutcomes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerScan(ctx context.Context, inquire string, screen *int, eachScreen *int, arrangeBy string) (*basetypes.OutcomeLedgerScan, error) {
	ret := _m.Called(ctx, inquire, screen, eachScreen, arrangeBy)

	var r0 *basetypes.OutcomeLedgerScan
	if rf, ok := ret.Get(0).(func(context.Context, string, *int, *int, string) *basetypes.OutcomeLedgerScan); ok {
		r0 = rf(ctx, inquire, screen, eachScreen, arrangeBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerScan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *int, *int, string) error); ok {
		r1 = rf(ctx, inquire, screen, eachScreen, arrangeBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) LedgerchainDetails(ctx context.Context, minimumLevel int64, maximumLevel int64) (*basetypes.OutcomeLedgerchainDetails, error) {
	ret := _m.Called(ctx, minimumLevel, maximumLevel)

	var r0 *basetypes.OutcomeLedgerchainDetails
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) *basetypes.OutcomeLedgerchainDetails); ok {
		r0 = rf(ctx, minimumLevel, maximumLevel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeLedgerchainDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, minimumLevel, maximumLevel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastProof(_a0 context.Context, _a1 kinds.Proof) (*basetypes.OutcomeMulticastProof, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastProof
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Proof) *basetypes.OutcomeMulticastProof); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastProof)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Proof) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferAsync(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferEndorse(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransferEndorse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransferEndorse
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransferEndorse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransferEndorse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) MulticastTransferAlign(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeMulticastTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeMulticastTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeMulticastTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeMulticastTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) InspectTransfer(_a0 context.Context, _a1 kinds.Tx) (*basetypes.OutcomeInspectTransfer, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeInspectTransfer
	if rf, ok := ret.Get(0).(func(context.Context, kinds.Tx) *basetypes.OutcomeInspectTransfer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeInspectTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kinds.Tx) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Endorse(ctx context.Context, level *int64) (*basetypes.OutcomeEndorse, error) {
	ret := _m.Called(ctx, level)

	var r0 *basetypes.OutcomeEndorse
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeEndorse); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeEndorse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) AgreementOptions(ctx context.Context, level *int64) (*basetypes.OutcomeAgreementOptions, error) {
	ret := _m.Called(ctx, level)

	var r0 *basetypes.OutcomeAgreementOptions
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeAgreementOptions); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeAgreementOptions)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) AgreementStatus(_a0 context.Context) (*basetypes.OutcomeAgreementStatus, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeAgreementStatus
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeAgreementStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeAgreementStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) ExportAgreementStatus(_a0 context.Context) (*basetypes.OutcomeExportAgreementStatus, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeExportAgreementStatus
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeExportAgreementStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeExportAgreementStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Origin(_a0 context.Context) (*basetypes.OutcomeOrigin, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeOrigin
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeOrigin); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeOrigin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) OriginSegmented(_a0 context.Context, _a1 uint) (*basetypes.OutcomeOriginSegment, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *basetypes.OutcomeOriginSegment
	if rf, ok := ret.Get(0).(func(context.Context, uint) *basetypes.OutcomeOriginSegment); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeOriginSegment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Heading(ctx context.Context, level *int64) (*basetypes.OutcomeHeading, error) {
	ret := _m.Called(ctx, level)

	var r0 *basetypes.OutcomeHeading
	if rf, ok := ret.Get(0).(func(context.Context, *int64) *basetypes.OutcomeHeading); ok {
		r0 = rf(ctx, level)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeHeading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64) error); ok {
		r1 = rf(ctx, level)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) HeadingByDigest(ctx context.Context, digest octets.HexOctets) (*basetypes.OutcomeHeading, error) {
	ret := _m.Called(ctx, digest)

	var r0 *basetypes.OutcomeHeading
	if rf, ok := ret.Get(0).(func(context.Context, octets.HexOctets) *basetypes.OutcomeHeading); ok {
		r0 = rf(ctx, digest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeHeading)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, octets.HexOctets) error); ok {
		r1 = rf(ctx, digest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Vitality(_a0 context.Context) (*basetypes.OutcomeVitality, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeVitality
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeVitality); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeVitality)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) IsActive() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

//
func (_m *Customer) NetDetails(_a0 context.Context) (*basetypes.OutcomeNetDetails, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeNetDetails
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeNetDetails); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeNetDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) CountUnattestedTrans(_a0 context.Context) (*basetypes.OutcomeUnattestedTrans, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeUnattestedTrans
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeUnattestedTrans); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeUnattestedTrans)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) OnRestore() error {
	ret := _m.Called()

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
func (_m *Customer) Exit() <-chan struct{} {
	ret := _m.Called()

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
func (_m *Customer) Restore() error {
	ret := _m.Called()

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
func (_m *Customer) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) Status(_a0 context.Context) (*basetypes.OutcomeState, error) {
	ret := _m.Called(_a0)

	var r0 *basetypes.OutcomeState
	if rf, ok := ret.Get(0).(func(context.Context) *basetypes.OutcomeState); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Halt() error {
	ret := _m.Called()

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

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

//
func (_m *Customer) Enrol(ctx context.Context, enrollee string, inquire string, outVolume ...int) (<-chan basetypes.OutcomeEvent, error) {
	_va := make([]any, len(outVolume))
	for _i := range outVolume {
		_va[_i] = outVolume[_i]
	}
	var _ca []any
	_ca = append(_ca, ctx, enrollee, inquire)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 <-chan basetypes.OutcomeEvent
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...int) <-chan basetypes.OutcomeEvent); ok {
		r0 = rf(ctx, enrollee, inquire, outVolume...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan basetypes.OutcomeEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...int) error); ok {
		r1 = rf(ctx, enrollee, inquire, outVolume...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Tx(ctx context.Context, digest []byte, demonstrate bool) (*basetypes.OutcomeTransfer, error) {
	ret := _m.Called(ctx, digest, demonstrate)

	var r0 *basetypes.OutcomeTransfer
	if rf, ok := ret.Get(0).(func(context.Context, []byte, bool) *basetypes.OutcomeTransfer); ok {
		r0 = rf(ctx, digest, demonstrate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeTransfer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, bool) error); ok {
		r1 = rf(ctx, digest, demonstrate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) TransferScan(ctx context.Context, inquire string, demonstrate bool, screen *int, eachScreen *int, arrangeBy string) (*basetypes.OutcomeTransferScan, error) {
	ret := _m.Called(ctx, inquire, demonstrate, screen, eachScreen, arrangeBy)

	var r0 *basetypes.OutcomeTransferScan
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, *int, *int, string) *basetypes.OutcomeTransferScan); ok {
		r0 = rf(ctx, inquire, demonstrate, screen, eachScreen, arrangeBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeTransferScan)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool, *int, *int, string) error); ok {
		r1 = rf(ctx, inquire, demonstrate, screen, eachScreen, arrangeBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) UnattestedTrans(ctx context.Context, ceiling *int) (*basetypes.OutcomeUnattestedTrans, error) {
	ret := _m.Called(ctx, ceiling)

	var r0 *basetypes.OutcomeUnattestedTrans
	if rf, ok := ret.Get(0).(func(context.Context, *int) *basetypes.OutcomeUnattestedTrans); ok {
		r0 = rf(ctx, ceiling)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeUnattestedTrans)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int) error); ok {
		r1 = rf(ctx, ceiling)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

//
func (_m *Customer) Deenroll(ctx context.Context, enrollee string, inquire string) error {
	ret := _m.Called(ctx, enrollee, inquire)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, enrollee, inquire)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) DeenrollAll(ctx context.Context, enrollee string) error {
	ret := _m.Called(ctx, enrollee)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, enrollee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
func (_m *Customer) Ratifiers(ctx context.Context, level *int64, screen *int, eachScreen *int) (*basetypes.OutcomeRatifiers, error) {
	ret := _m.Called(ctx, level, screen, eachScreen)

	var r0 *basetypes.OutcomeRatifiers
	if rf, ok := ret.Get(0).(func(context.Context, *int64, *int, *int) *basetypes.OutcomeRatifiers); ok {
		r0 = rf(ctx, level, screen, eachScreen)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*basetypes.OutcomeRatifiers)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *int64, *int, *int) error); ok {
		r1 = rf(ctx, level, screen, eachScreen)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
