package abciend

import (
	"bufio"
	"container/list"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/valkyrieworks/iface/kinds"
	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/utils/daemon"
	"github.com/valkyrieworks/utils/clock"
)

const (
	requestBufferVolume    = 256 //
	purgeRegulateMillis = 20  //
)

//
//
//
//
//
//
type socketCustomer struct {
	daemon.RootDaemon

	address        string
	shouldLink bool
	link        net.Conn

	requestBuffer   chan *RequestOutput
	purgeClock *clock.RegulateClock

	mtx     sync.Mutex
	err     error
	requestRelayed *list.List                            //
	outputCallbackfn   func(*kinds.Query, *kinds.Reply) //
}

var _ Customer = (*socketCustomer)(nil)

//
//
//
func NewSocketCustomer(address string, shouldLink bool) Customer {
	cli := &socketCustomer{
		requestBuffer:    make(chan *RequestOutput, requestBufferVolume),
		purgeClock:  clock.NewRegulateClock("REDACTED", purgeRegulateMillis),
		shouldLink: shouldLink,

		address:    address,
		requestRelayed: list.New(),
		outputCallbackfn:   nil,
	}
	cli.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", cli)
	return cli
}

//
//
func (cli *socketCustomer) OnBegin() error {
	var (
		err  error
		link net.Conn
	)

	for {
		link, err = cometnet.Link(cli.address)
		if err != nil {
			if cli.shouldLink {
				return err
			}
			cli.Tracer.Fault(fmt.Sprintf("REDACTED",
				cli.address, callReprocessCadenceMoments), "REDACTED", err)
			time.Sleep(time.Second * callReprocessCadenceMoments)
			continue
		}
		cli.link = link

		go cli.transmitQueriesProcedure(link)
		go cli.receiveReplyProcedure(link)

		return nil
	}
}

//
func (cli *socketCustomer) OnHalt() {
	if cli.link != nil {
		cli.link.Close()
	}

	cli.purgeBuffer()
	cli.purgeClock.Halt()
}

//
func (cli *socketCustomer) Fault() error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	return cli.err
}

//

//
//
//
//
func (cli *socketCustomer) CollectionReplyCallback(outputCallbackfn Callback) {
	cli.mtx.Lock()
	cli.outputCallbackfn = outputCallbackfn
	cli.mtx.Unlock()
}

func (cli *socketCustomer) InspectTransferAsync(ctx context.Context, req *kinds.QueryInspectTransfer) (*RequestOutput, error) {
	return cli.bufferQuery(ctx, kinds.ToQueryInspectTransfer(req))
}

//

func (cli *socketCustomer) transmitQueriesProcedure(link io.Writer) {
	w := bufio.NewWriter(link)
	for {
		select {
		case requestresponse := <-cli.requestBuffer:
			//
			//
			//
			cli.monitorQuery(requestresponse)

			err := kinds.RecordSignal(requestresponse.Query, w)
			if err != nil {
				cli.haltForFault(fmt.Errorf("REDACTED", err))
				return
			}

			//
			if _, ok := requestresponse.Query.Item.(*kinds.Query_Purge); ok {
				err = w.Flush()
				if err != nil {
					cli.haltForFault(fmt.Errorf("REDACTED", err))
					return
				}
			}
		case <-cli.purgeClock.Ch: //
			select {
			case cli.requestBuffer <- NewRequestOutput(kinds.ToQueryPurge()):
			default:
				//
			}
		case <-cli.Exit():
			return
		}
	}
}

func (cli *socketCustomer) receiveReplyProcedure(link io.Reader) {
	r := bufio.NewReader(link)
	for {
		if !cli.IsActive() {
			return
		}

		res := &kinds.Reply{}
		err := kinds.FetchSignal(r, res)
		if err != nil {
			cli.haltForFault(fmt.Errorf("REDACTED", err))
			return
		}

		switch r := res.Item.(type) {
		case *kinds.Reply_Exemption: //
			//
			cli.haltForFault(errors.New(r.Exemption.Fault))
			return
		default:
			err := cli.didReceiveReply(res)
			if err != nil {
				cli.haltForFault(err)
				return
			}
		}
	}
}

func (cli *socketCustomer) monitorQuery(requestresponse *RequestOutput) {
	//
	//
	if !cli.IsActive() {
		return
	}

	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	cli.requestRelayed.PushBack(requestresponse)
}

func (cli *socketCustomer) didReceiveReply(res *kinds.Reply) error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()

	//
	following := cli.requestRelayed.Front()
	if following == nil {
		return ErrUnforeseenReply{Reply: *res, Cause: "REDACTED"}
	}

	requestresponse := following.Value.(*RequestOutput)
	if !outputAlignsRequest(requestresponse.Query, res) {
		return ErrUnforeseenReply{Reply: *res, Cause: fmt.Sprintf("REDACTED", requestresponse.Query.Item)}
	}

	requestresponse.Reply = res
	requestresponse.Done()            //
	cli.requestRelayed.Remove(following) //

	//
	if cli.outputCallbackfn != nil {
		cli.outputCallbackfn(requestresponse.Query, res)
	}

	//
	//
	//
	//
	requestresponse.ExecuteCallback()

	return nil
}

//

func (cli *socketCustomer) Purge(ctx context.Context) error {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryPurge())
	if err != nil {
		return err
	}
	requestOutput.Wait()
	return nil
}

func (cli *socketCustomer) Replicate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryReverberate(msg))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchReverberate(), cli.Fault()
}

func (cli *socketCustomer) Details(ctx context.Context, req *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryDetails(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchDetails(), cli.Fault()
}

func (cli *socketCustomer) InspectTransfer(ctx context.Context, req *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryInspectTransfer(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchInspectTransfer(), cli.Fault()
}

func (cli *socketCustomer) EmbedTransfer(ctx context.Context, req *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryEmbedTransfer(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchEmbedTransfer(), cli.Fault()
}

func (cli *socketCustomer) HarvestTrans(ctx context.Context, req *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryHarvestTrans(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchHarvestTrans(), cli.Fault()
}

func (cli *socketCustomer) Inquire(ctx context.Context, req *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryInquire(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchInquire(), cli.Fault()
}

func (cli *socketCustomer) Endorse(ctx context.Context, _ *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryEndorse())
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchEndorse(), cli.Fault()
}

func (cli *socketCustomer) InitSeries(ctx context.Context, req *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryInitSeries(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchInitSeries(), cli.Fault()
}

func (cli *socketCustomer) CatalogMirrors(ctx context.Context, req *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryCatalogMirrors(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchCatalogMirrors(), cli.Fault()
}

func (cli *socketCustomer) ProposalMirror(ctx context.Context, req *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryProposalMirror(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchProposalMirror(), cli.Fault()
}

func (cli *socketCustomer) ImportMirrorSegment(ctx context.Context, req *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryImportMirrorSegment(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchImportMirrorSegment(), cli.Fault()
}

func (cli *socketCustomer) ExecuteMirrorSegment(ctx context.Context, req *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryExecuteMirrorSegment(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchExecuteMirrorSegment(), cli.Fault()
}

func (cli *socketCustomer) ArrangeNomination(ctx context.Context, req *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryArrangeNomination(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchArrangeNomination(), cli.Fault()
}

func (cli *socketCustomer) HandleNomination(ctx context.Context, req *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryHandleNomination(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchHandleNomination(), cli.Fault()
}

func (cli *socketCustomer) ExpandBallot(ctx context.Context, req *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryExpandBallot(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchExpandBallot(), cli.Fault()
}

func (cli *socketCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryValidateBallotAddition(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchValidateBallotAddition(), cli.Fault()
}

func (cli *socketCustomer) CompleteLedger(ctx context.Context, req *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	requestOutput, err := cli.bufferQuery(ctx, kinds.ToQueryCompleteLedger(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestOutput.Reply.FetchCompleteLedger(), cli.Fault()
}

func (cli *socketCustomer) bufferQuery(ctx context.Context, req *kinds.Query) (*RequestOutput, error) {
	requestresponse := NewRequestOutput(req)

	//
	select {
	case cli.requestBuffer <- requestresponse:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	//
	switch req.Item.(type) {
	case *kinds.Query_Purge:
		cli.purgeClock.Clear()
	default:
		cli.purgeClock.Set()
	}

	return requestresponse, nil
}

//
//
func (cli *socketCustomer) purgeBuffer() {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()

	//
	for req := cli.requestRelayed.Front(); req != nil; req = req.Next() {
		requestresponse := req.Value.(*RequestOutput)
		requestresponse.Done()
	}

	//
Cycle:
	for {
		select {
		case requestresponse := <-cli.requestBuffer:
			requestresponse.Done()
		default:
			break Cycle
		}
	}
}

//

func outputAlignsRequest(req *kinds.Query, res *kinds.Reply) (ok bool) {
	switch req.Item.(type) {
	case *kinds.Query_Reverberate:
		_, ok = res.Item.(*kinds.Reply_Reverberate)
	case *kinds.Query_Purge:
		_, ok = res.Item.(*kinds.Reply_Purge)
	case *kinds.Query_Details:
		_, ok = res.Item.(*kinds.Reply_Details)
	case *kinds.Query_Transfercheck:
		_, ok = res.Item.(*kinds.Reply_Transfercheck)
	case *kinds.Query_Endorse:
		_, ok = res.Item.(*kinds.Reply_Endorse)
	case *kinds.Query_Inquire:
		_, ok = res.Item.(*kinds.Reply_Inquire)
	case *kinds.Query_Initiatechain:
		_, ok = res.Item.(*kinds.Reply_Initiatechain)
	case *kinds.Query_Executemirrorsegment:
		_, ok = res.Item.(*kinds.Reply_Executemirrorsegment)
	case *kinds.Query_Loadmirrorsegment:
		_, ok = res.Item.(*kinds.Reply_Loadmirrorsegment)
	case *kinds.Query_Catalogmirrors:
		_, ok = res.Item.(*kinds.Reply_Catalogmirrors)
	case *kinds.Query_Mirrorsnapshot:
		_, ok = res.Item.(*kinds.Reply_Mirrorsnapshot)
	case *kinds.Query_Ballotextend:
		_, ok = res.Item.(*kinds.Reply_Ballotextend)
	case *kinds.Query_Validateballotextension:
		_, ok = res.Item.(*kinds.Reply_Validateballotextension)
	case *kinds.Query_Arrangenomination:
		_, ok = res.Item.(*kinds.Reply_Arrangenomination)
	case *kinds.Query_Processnomination:
		_, ok = res.Item.(*kinds.Reply_Processnomination)
	case *kinds.Query_Terminateblock:
		_, ok = res.Item.(*kinds.Reply_Terminateblock)
	}
	return ok
}

func (cli *socketCustomer) haltForFault(err error) {
	if !cli.IsActive() {
		return
	}

	cli.mtx.Lock()
	if cli.err == nil {
		cli.err = err
	}
	cli.mtx.Unlock()

	cli.Tracer.Fault(fmt.Sprintf("REDACTED", err.Error()))
	if err := cli.Halt(); err != nil {
		cli.Tracer.Fault("REDACTED", "REDACTED", err)
	}
}
