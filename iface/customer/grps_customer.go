package abcicustomer

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
)

var _ Customer = (*grpsCustomer)(nil)

//
//
type grpsCustomer struct {
	facility.FoundationFacility
	shouldRelate bool

	customer   kinds.IfaceCustomer
	link     *grpc.ClientConn
	chnlRequestResult chan *RequestResult //

	mtx   sync.Mutex
	location  string
	err   error
	resultClbk func(*kinds.Solicit, *kinds.Reply) //
}

func FreshGRPSCustomer(location string, shouldRelate bool) Customer {
	cli := &grpsCustomer{
		location:        location,
		shouldRelate: shouldRelate,
		//
		//
		//
		//
		//
		//
		chnlRequestResult: make(chan *RequestResult, 64),
	}
	cli.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", cli)
	return cli
}

func callerMethod(_ context.Context, location string) (net.Conn, error) {
	return strongmindnet.Relate(location)
}

func (cli *grpsCustomer) UponInitiate() error {
	if err := cli.FoundationFacility.UponInitiate(); err != nil {
		return err
	}

	//
	//
	go func() {
		//
		invocationClbk := func(requestresponse *RequestResult) {
			cli.mtx.Lock()
			defer cli.mtx.Unlock()

			requestresponse.Done()

			//
			if cli.resultClbk != nil {
				cli.resultClbk(requestresponse.Solicit, requestresponse.Reply)
			}

			//
			requestresponse.ExecuteClbk()
		}
		for requestresponse := range cli.chnlRequestResult {
			if requestresponse != nil {
				invocationClbk(requestresponse)
			} else {
				cli.Tracer.Failure("REDACTED")
			}
		}
	}()

REISSUE_CYCLE:
	for {
		link, err := grpc.NewClient(cli.location,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithContextDialer(callerMethod),
		)
		if err != nil {
			if cli.shouldRelate {
				return err
			}
			cli.Tracer.Failure(fmt.Sprintf("REDACTED", cli.location), "REDACTED", err)
			time.Sleep(time.Second * callReissueDurationMoments)
			continue REISSUE_CYCLE
		}

		cli.Tracer.Details("REDACTED", "REDACTED", cli.location)
		customer := kinds.FreshIfaceCustomer(link)
		cli.link = link

	ASSURE_ASSOCIATED:
		for {
			_, err := customer.Reverberate(context.Background(), &kinds.SolicitReverberate{Signal: "REDACTED"}, grpc.WaitForReady(true))
			if err == nil {
				break ASSURE_ASSOCIATED
			}
			cli.Tracer.Failure("REDACTED", "REDACTED", err)
			time.Sleep(time.Second * reverberateReissueDurationMoments)
		}

		cli.customer = customer
		return nil
	}
}

func (cli *grpsCustomer) UponHalt() {
	cli.FoundationFacility.UponHalt()

	if cli.link != nil {
		cli.link.Close()
	}
	close(cli.chnlRequestResult)
}

func (cli *grpsCustomer) HaltForeachFailure(err error) {
	if !cli.EqualsActive() {
		return
	}

	cli.mtx.Lock()
	if cli.err == nil {
		cli.err = err
	}
	cli.mtx.Unlock()

	cli.Tracer.Failure(fmt.Sprintf("REDACTED", err.Error()))
	if err := cli.Halt(); err != nil {
		cli.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

func (cli *grpsCustomer) Failure() error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	return cli.err
}

//
//
func (cli *grpsCustomer) AssignReplyClbk(resultClbk Clbk) {
	cli.mtx.Lock()
	cli.resultClbk = resultClbk
	cli.mtx.Unlock()
}

//

func (cli *grpsCustomer) InspectTransferAsyncronous(ctx context.Context, req *kinds.SolicitInspectTransfer) (*RequestResult, error) {
	res, err := cli.customer.InspectTransfer(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		cli.HaltForeachFailure(err)
		return nil, err
	}
	return cli.concludeAsyncronousInvocation(kinds.TowardSolicitInspectTransfer(req), &kinds.Reply{Datum: &kinds.Reply_Inspecttrans{InspectTransfer: res}}), nil
}

//
//
func (cli *grpsCustomer) concludeAsyncronousInvocation(req *kinds.Solicit, res *kinds.Reply) *RequestResult {
	requestresponse := FreshRequestResult(req)
	requestresponse.Reply = res
	cli.chnlRequestResult <- requestresponse //
	return requestresponse
}

//

func (cli *grpsCustomer) Purge(ctx context.Context) error {
	_, err := cli.customer.Purge(ctx, kinds.TowardSolicitPurge().ObtainPurge(), grpc.WaitForReady(true))
	return err
}

func (cli *grpsCustomer) Reverberate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	return cli.customer.Reverberate(ctx, kinds.TowardSolicitReverberate(msg).ObtainReverberate(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) Details(ctx context.Context, req *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	return cli.customer.Details(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpsCustomer) InspectTransfer(ctx context.Context, req *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	return cli.customer.InspectTransfer(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpsCustomer) AppendTransfer(ctx context.Context, req *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
	return cli.customer.AppendTransfer(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpsCustomer) HarvestTrans(ctx context.Context, req *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	return cli.customer.HarvestTrans(ctx, req, grpc.WaitForReady(true))
}

func (cli *grpsCustomer) Inquire(ctx context.Context, req *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	return cli.customer.Inquire(ctx, kinds.TowardSolicitInquire(req).ObtainInquire(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) Endorse(ctx context.Context, _ *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
	return cli.customer.Endorse(ctx, kinds.TowardSolicitEndorse().ObtainEndorse(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) InitializeSuccession(ctx context.Context, req *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	return cli.customer.InitializeSuccession(ctx, kinds.TowardSolicitInitializeSuccession(req).ObtainInitializeSuccession(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) CollectionImages(ctx context.Context, req *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	return cli.customer.CollectionImages(ctx, kinds.TowardSolicitCatalogImages(req).ObtainCatalogImages(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) ExtendImage(ctx context.Context, req *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	return cli.customer.ExtendImage(ctx, kinds.TowardSolicitExtendImage(req).ObtainExtendImage(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) FetchImageSegment(ctx context.Context, req *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
	return cli.customer.FetchImageSegment(ctx, kinds.TowardSolicitFetchImageSegment(req).ObtainFetchImageSegment(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) ExecuteImageSegment(ctx context.Context, req *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
	return cli.customer.ExecuteImageSegment(ctx, kinds.TowardSolicitExecuteImageSegment(req).ObtainExecuteImageSegment(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) ArrangeNomination(ctx context.Context, req *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	return cli.customer.ArrangeNomination(ctx, kinds.TowardSolicitArrangeNomination(req).ObtainArrangeNomination(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) HandleNomination(ctx context.Context, req *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	return cli.customer.HandleNomination(ctx, kinds.TowardSolicitHandleNomination(req).ObtainHandleNomination(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) BroadenBallot(ctx context.Context, req *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
	return cli.customer.BroadenBallot(ctx, kinds.TowardSolicitBroadenBallot(req).ObtainBroadenBallot(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	return cli.customer.ValidateBallotAddition(ctx, kinds.TowardSolicitValidateBallotAddition(req).ObtainValidateBallotAddition(), grpc.WaitForReady(true))
}

func (cli *grpsCustomer) CulminateLedger(ctx context.Context, req *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	return cli.customer.CulminateLedger(ctx, kinds.TowardSolicitCulminateLedger(req).ObtainCulminateLedger(), grpc.WaitForReady(true))
}
