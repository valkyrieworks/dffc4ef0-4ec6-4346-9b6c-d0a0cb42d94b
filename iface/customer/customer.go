package abcicustomer

import (
	"context"
	"sync"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
)

const (
	callReissueDurationMoments = 3
	reverberateReissueDurationMoments = 1
)

//

//
//
//
//
//
type Customer interface {
	facility.Facility
	kinds.Platform

	//
	Failure() error
	//
	Purge(context.Context) error
	Reverberate(context.Context, string) (*kinds.ReplyReverberate, error)

	//
	//
	//
	//
	//
	AssignReplyClbk(Clbk)
	InspectTransferAsyncronous(context.Context, *kinds.SolicitInspectTransfer) (*RequestResult, error)
}

//

//
//
func FreshCustomer(location, carrier string, shouldRelate bool) (customer Customer, err error) {
	switch carrier {
	case "REDACTED":
		customer = FreshPortCustomer(location, shouldRelate)
	case "REDACTED":
		customer = FreshGRPSCustomer(location, shouldRelate)
	default:
		err = FaultUnfamiliarIfaceCarrier{Carrier: carrier}
	}
	return
}

type Clbk func(*kinds.Solicit, *kinds.Reply)

type RequestResult struct {
	*kinds.Solicit
	*sync.PauseCluster
	*kinds.Reply //

	mtx commitchronize.Exclusion

	//
	//
	//
	//
	//
	clbkExecuted bool
	cb              func(*kinds.Reply) //
}

func FreshRequestResult(req *kinds.Solicit) *RequestResult {
	return &RequestResult{
		Solicit:   req,
		PauseCluster: pauseCluster1(),
		Reply:  nil,

		clbkExecuted: false,
		cb:              nil,
	}
}

//
//
//
func (r *RequestResult) AssignClbk(cb func(res *kinds.Reply)) {
	r.mtx.Lock()

	if r.clbkExecuted {
		r.mtx.Unlock()
		cb(r.Reply)
		return
	}

	r.cb = cb
	r.mtx.Unlock()
}

//
//
func (r *RequestResult) ExecuteClbk() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if r.cb != nil {
		r.cb(r.Reply)
	}
	r.clbkExecuted = true
}

//
//
//
//
//
//
func (r *RequestResult) ObtainClbk() func(*kinds.Reply) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return r.cb
}

func pauseCluster1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
