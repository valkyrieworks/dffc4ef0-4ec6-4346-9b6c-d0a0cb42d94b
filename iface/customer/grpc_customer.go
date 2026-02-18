package abciend

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valkyrieworks/iface/kinds"
	cometnet "github.com/valkyrieworks/utils/net"
	"github.com/valkyrieworks/utils/daemon"
)

var _ Customer = (*grpcCustomer)(nil)

//
//
type grpcCustomer struct {
	daemon.RootDaemon
	shouldLink bool

	customer   kinds.IfaceCustomer
	link     *grpc.ClientConn
	chanRequestOutput chan *RequestOutput //

	mtx   sync.Mutex
	address  string
	err   error
	outputCallbackfn func(*kinds.Query, *kinds.Reply) //
}

func NewGRPCCustomer(address string, shouldLink bool) Customer {
	cli := &grpcCustomer{
		address:        address,
		shouldLink: shouldLink,
		//
		//
		//
		//
		//
		//
		chanRequestOutput: make(chan *RequestOutput, 64),
	}
	cli.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", cli)
	return cli
}

func callerFunction(_ context.Context, address string) (net.Conn, error) {
	return cometnet.Link(address)
}

func (cli *grpcCustomer) OnBegin() error {
	if err := cli.RootDaemon.OnBegin(); err != nil {
		return err
	}

	//
	//
	go func() {
		//
		invokeCallbackfn := func(requestresponse *RequestOutput) {
			cli.mtx.Lock()
			defer cli.mtx.Unlock()

			requestresponse.Done()

			//
			if cli.outputCallbackfn != nil {
				cli.outputCallbackfn(requestresponse.Query, requestresponse.Reply)
			}

			//
			requestresponse.ExecuteCallback()
		}
		for requestresponse := range cli.chanRequestOutput {
			if requestresponse != nil {
				invokeCallbackfn(requestresponse)
			} else {
				cli.Tracer.Fault("REDACTED")
			}
		}
	}()

REPROCESS_CYCLE:
	for {
		link, err := grpc.NewClient(cli.address,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(callerFunction),
		)
		if err != nil {
			if cli.shouldLink {
				return err
			}
			cli.Tracer.Fault(fmt.Sprintf("REDACTED", cli.address), "REDACTED", err)
			time.Sleep(time.Second * callReprocessCadenceMoments)
			continue REPROCESS_CYCLE
		}

		cli.Tracer.Details("REDACTED", "REDACTED", cli.address)
		customer := kinds.NewIfaceCustomer(link)
		cli.link = link

	ASSURE_LINKED:
		for {
			_, err := customer.Replicate(context.Background(), &kinds.QueryReverberate{Signal: "REDACTED"}, grpc.WaitForReady(true))
			if err == nil {
				break ASSURE_LINKED
			}
			cli.Tracer.Fault("REDACTED", "REDACTED", err)
			time.Sleep(time.Second * reverberateReprocessCadenceMoments)
		}

		cli.customer = customer
		return nil
	}
}

func (cli *grpcCustomer) OnHalt() {
	cli.RootDaemon.OnHalt()

	if cli.link != nil {
		cli.link.Close()
	}
	close(cli.chanRequestOutput)
}

func (cli *grpcCustomer) HaltForFault(err error) {
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

func (cli *grpcCustomer) Fault() error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	return cli.err
}

//
//
func (cli *grpcCustomer) CollectionReplyCallback(outputCallbackfn Callback) {
	cli.mtx.Lock()
	cli.outputCallbackfn = outputCallbackfn
	cli.mtx.Unlock()
}

//

func (cli *grpcCustomer) InspectTransferAsync(ctx context.Context, req *kinds.QueryInspectTransfer) (*RequestOutput, error) {
	res, err := cli.customer.InspectTransfer(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		cli.HaltForFault(err)
		return nil, err
	}
	return cli.concludeAsyncInvoke(kinds.ToQueryInspectTransfer(req), &kinds.Reply{Item: &kinds.Reply_Transfercheck{InspectTransfer: res}}), nil
}

//
//
func (cli *grpcCustomer) concludeAsyncInvoke(req *kinds.Query, res *kinds.Reply) *RequestOutput {
	requestresponse := NewRequestOutput(req)
	requestresponse.Reply = res
	cli.chanRequestOutput <- requestresponse //
	return requestresponse
}

//

func (cli *grpcCustomer) Purge(ctx context.Context) error {
	_, err := cli.customer.Purge(ctx, kinds.ToQueryPurge().FetchPurge(), grpc.WaitForReady(true))
	return err
}

func (cli *grpcCustomer) Replicate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	return cli.customer.Replicate(ctx, kinds.ToQueryReverberate(msg).FetchReverberate(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) Details(ctx context.Context, req *kinds.QueryDetails) (*kinds.ReplyDetails, error) {
	return cli.customer.Details(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpcCustomer) InspectTransfer(ctx context.Context, req *kinds.QueryInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	return cli.customer.InspectTransfer(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpcCustomer) EmbedTransfer(ctx context.Context, req *kinds.QueryEmbedTransfer) (*kinds.ReplyEmbedTransfer, error) {
	return cli.customer.EmbedTransfer(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpcCustomer) HarvestTrans(ctx context.Context, req *kinds.QueryHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	return cli.customer.HarvestTrans(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpcCustomer) Inquire(ctx context.Context, req *kinds.QueryInquire) (*kinds.ReplyInquire, error) {
	return cli.customer.Inquire(ctx, kinds.ToQueryInquire(req).FetchInquire(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) Endorse(ctx context.Context, _ *kinds.QueryEndorse) (*kinds.ReplyEndorse, error) {
	return cli.customer.Endorse(ctx, kinds.ToQueryEndorse().FetchEndorse(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) InitSeries(ctx context.Context, req *kinds.QueryInitSeries) (*kinds.ReplyInitSeries, error) {
	return cli.customer.InitSeries(ctx, kinds.ToQueryInitSeries(req).FetchInitSeries(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) CatalogMirrors(ctx context.Context, req *kinds.QueryCatalogMirrors) (*kinds.ReplyCatalogMirrors, error) {
	return cli.customer.CatalogMirrors(ctx, kinds.ToQueryCatalogMirrors(req).FetchCatalogMirrors(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ProposalMirror(ctx context.Context, req *kinds.QueryProposalMirror) (*kinds.ReplyProposalMirror, error) {
	return cli.customer.ProposalMirror(ctx, kinds.ToQueryProposalMirror(req).FetchProposalMirror(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ImportMirrorSegment(ctx context.Context, req *kinds.QueryImportMirrorSegment) (*kinds.ReplyImportMirrorSegment, error) {
	return cli.customer.ImportMirrorSegment(ctx, kinds.ToQueryImportMirrorSegment(req).FetchImportMirrorSegment(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ExecuteMirrorSegment(ctx context.Context, req *kinds.QueryExecuteMirrorSegment) (*kinds.ReplyExecuteMirrorSegment, error) {
	return cli.customer.ExecuteMirrorSegment(ctx, kinds.ToQueryExecuteMirrorSegment(req).FetchExecuteMirrorSegment(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ArrangeNomination(ctx context.Context, req *kinds.QueryArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	return cli.customer.ArrangeNomination(ctx, kinds.ToQueryArrangeNomination(req).FetchArrangeNomination(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) HandleNomination(ctx context.Context, req *kinds.QueryHandleNomination) (*kinds.ReplyHandleNomination, error) {
	return cli.customer.HandleNomination(ctx, kinds.ToQueryHandleNomination(req).FetchHandleNomination(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ExpandBallot(ctx context.Context, req *kinds.QueryExpandBallot) (*kinds.ReplyExpandBallot, error) {
	return cli.customer.ExpandBallot(ctx, kinds.ToQueryExpandBallot(req).FetchExpandBallot(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.QueryValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	return cli.customer.ValidateBallotAddition(ctx, kinds.ToQueryValidateBallotAddition(req).FetchValidateBallotAddition(), grpc.WaitForReady(true))
}

func (cli *grpcCustomer) CompleteLedger(ctx context.Context, req *kinds.QueryCompleteLedger) (*kinds.ReplyCompleteLedger, error) {
	return cli.customer.CompleteLedger(ctx, kinds.ToQueryCompleteLedger(req).FetchCompleteLedger(), grpc.WaitForReady(true))
}
