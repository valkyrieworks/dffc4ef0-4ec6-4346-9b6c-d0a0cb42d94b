package emulate

import (
	"context"

	"github.com/valkyrieworks/rpc/customer"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
)

//
type StateEmulate struct {
	Invoke
}

var (
	_ customer.StateCustomer = (*StateEmulate)(nil)
	_ customer.StateCustomer = (*StateTracer)(nil)
)

func (m *StateEmulate) Status(context.Context) (*ctypes.OutcomeState, error) {
	res, err := m.FetchReply(nil)
	if err != nil {
		return nil, err
	}
	return res.(*ctypes.OutcomeState), nil
}

//
//
type StateTracer struct {
	Customer customer.StateCustomer
	Invocations  []Invoke
}

func NewStateTracer(customer customer.StateCustomer) *StateTracer {
	return &StateTracer{
		Customer: customer,
		Invocations:  []Invoke{},
	}
}

func (r *StateTracer) appendInvoke(invocation Invoke) {
	r.Invocations = append(r.Invocations, invocation)
}

func (r *StateTracer) Status(ctx context.Context) (*ctypes.OutcomeState, error) {
	res, err := r.Customer.Status(ctx)
	r.appendInvoke(Invoke{
		Label:     "REDACTED",
		Reply: res,
		Fault:    err,
	})
	return res, err
}
