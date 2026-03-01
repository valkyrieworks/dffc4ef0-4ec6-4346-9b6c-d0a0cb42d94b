package simulate

import (
	"context"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
)

//
type ConditionSimulate struct {
	Invocation
}

var (
	_ customer.ConditionCustomer = (*ConditionSimulate)(nil)
	_ customer.ConditionCustomer = (*ConditionScribe)(nil)
)

func (m *ConditionSimulate) Condition(context.Context) (*ktypes.OutcomeCondition, error) {
	res, err := m.ObtainReply(nil)
	if err != nil {
		return nil, err
	}
	return res.(*ktypes.OutcomeCondition), nil
}

//
//
type ConditionScribe struct {
	Customer customer.ConditionCustomer
	Invocations  []Invocation
}

func FreshConditionScribe(customer customer.ConditionCustomer) *ConditionScribe {
	return &ConditionScribe{
		Customer: customer,
		Invocations:  []Invocation{},
	}
}

func (r *ConditionScribe) appendInvocation(invocation Invocation) {
	r.Invocations = append(r.Invocations, invocation)
}

func (r *ConditionScribe) Condition(ctx context.Context) (*ktypes.OutcomeCondition, error) {
	res, err := r.Customer.Condition(ctx)
	r.appendInvocation(Invocation{
		Alias:     "REDACTED",
		Reply: res,
		Failure:    err,
	})
	return res, err
}
