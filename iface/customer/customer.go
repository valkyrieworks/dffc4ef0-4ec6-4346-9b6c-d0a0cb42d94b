package abciend

import (
	"context"
	"sync"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
)

const (
	callReprocessCadenceMoments = 3
	reverberateReprocessCadenceMoments = 1
)

//

//
//
//
//
//
type Customer interface {
	daemon.Daemon
	kinds.Software

	//
	Fault() error
	//
	Purge(context.Context) error
	Replicate(context.Context, string) (*kinds.ReplyReverberate, error)

	//
	//
	//
	//
	//
	CollectionReplyCallback(Callback)
	InspectTransferAsync(context.Context, *kinds.QueryInspectTransfer) (*RequestOutput, error)
}

//

//
//
func NewCustomer(address, carrier string, shouldLink bool) (customer Customer, err error) {
	switch carrier {
	case "REDACTED":
		customer = NewSocketCustomer(address, shouldLink)
	case "REDACTED":
		customer = NewGRPCCustomer(address, shouldLink)
	default:
		err = ErrUnclearIfaceCarrier{Carrier: carrier}
	}
	return
}

type Callback func(*kinds.Query, *kinds.Reply)

type RequestOutput struct {
	*kinds.Query
	*sync.WaitCluster
	*kinds.Reply //

	mtx engineconnect.Lock

	//
	//
	//
	//
	//
	callbackExecuted bool
	cb              func(*kinds.Reply) //
}

func NewRequestOutput(req *kinds.Query) *RequestOutput {
	return &RequestOutput{
		Query:   req,
		WaitCluster: waitCluster1(),
		Reply:  nil,

		callbackExecuted: false,
		cb:              nil,
	}
}

//
//
//
func (r *RequestOutput) CollectionCallback(cb func(res *kinds.Reply)) {
	r.mtx.Lock()

	if r.callbackExecuted {
		r.mtx.Unlock()
		cb(r.Reply)
		return
	}

	r.cb = cb
	r.mtx.Unlock()
}

//
//
func (r *RequestOutput) ExecuteCallback() {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	if r.cb != nil {
		r.cb(r.Reply)
	}
	r.callbackExecuted = true
}

//
//
//
//
//
//
func (r *RequestOutput) FetchCallback() func(*kinds.Reply) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	return r.cb
}

func waitCluster1() (wg *sync.WaitGroup) {
	wg = &sync.WaitGroup{}
	wg.Add(1)
	return
}
