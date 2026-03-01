package abcicustomer

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	strongmindnet "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/net"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/clock"
)

const (
	requestStagingExtent    = 256 //
	purgeRegulateMSEC = 20  //
)

//
//
//
//
//
//
type portCustomer struct {
	facility.FoundationFacility

	location        string
	shouldRelate bool
	link        net.Conn

	requestStaging   chan *RequestResult
	purgeClock *clock.RegulateClock

	mtx     sync.Mutex
	err     error
	requestRelayed *list.List                            //
	resultClbk   func(*kinds.Solicit, *kinds.Reply) //
}

var _ Customer = (*portCustomer)(nil)

//
//
//
func FreshPortCustomer(location string, shouldRelate bool) Customer {
	cli := &portCustomer{
		requestStaging:    make(chan *RequestResult, requestStagingExtent),
		purgeClock:  clock.FreshRegulateClock("REDACTED", purgeRegulateMSEC),
		shouldRelate: shouldRelate,

		location:    location,
		requestRelayed: list.New(),
		resultClbk:   nil,
	}
	cli.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", cli)
	return cli
}

//
//
func (cli *portCustomer) UponInitiate() error {
	var (
		err  error
		link net.Conn
	)

	for {
		link, err = strongmindnet.Relate(cli.location)
		if err != nil {
			if cli.shouldRelate {
				return err
			}
			cli.Tracer.Failure(fmt.Sprintf("REDACTED",
				cli.location, callReissueDurationMoments), "REDACTED", err)
			time.Sleep(time.Second * callReissueDurationMoments)
			continue
		}
		cli.link = link

		go cli.transmitSolicitsProcedure(link)
		go cli.obtainReplyProcedure(link)

		return nil
	}
}

//
func (cli *portCustomer) UponHalt() {
	if cli.link != nil {
		cli.link.Close()
	}

	cli.purgeStaging()
	cli.purgeClock.Halt()
}

//
func (cli *portCustomer) Failure() error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	return cli.err
}

//

//
//
//
//
func (cli *portCustomer) AssignReplyClbk(resultClbk Clbk) {
	cli.mtx.Lock()
	cli.resultClbk = resultClbk
	cli.mtx.Unlock()
}

func (cli *portCustomer) InspectTransferAsyncronous(ctx context.Context, req *kinds.SolicitInspectTransfer) (*RequestResult, error) {
	return cli.stagingSolicit(ctx, kinds.TowardSolicitInspectTransfer(req))
}

//

func (cli *portCustomer) transmitSolicitsProcedure(link io.Writer) {
	w := bufio.NewWriter(link)
	for {
		select {
		case requestresponse := <-cli.requestStaging:
			//
			//
			//
			cli.monitorSolicit(requestresponse)

			err := kinds.PersistArtifact(requestresponse.Solicit, w)
			if err != nil {
				cli.haltForeachFailure(fmt.Errorf("REDACTED", err))
				return
			}

			//
			if _, ok := requestresponse.Solicit.Datum.(*kinds.Solicit_Purge); ok {
				err = w.Flush()
				if err != nil {
					cli.haltForeachFailure(fmt.Errorf("REDACTED", err))
					return
				}
			}
		case <-cli.purgeClock.Ch: //
			select {
			case cli.requestStaging <- FreshRequestResult(kinds.TowardSolicitPurge()):
			default:
				//
			}
		case <-cli.Exit():
			return
		}
	}
}

func (cli *portCustomer) obtainReplyProcedure(link io.Reader) {
	r := bufio.NewReader(link)
	for {
		if !cli.EqualsActive() {
			return
		}

		res := &kinds.Reply{}
		err := kinds.FetchArtifact(r, res)
		if err != nil {
			cli.haltForeachFailure(fmt.Errorf("REDACTED", err))
			return
		}

		switch r := res.Datum.(type) {
		case *kinds.Reply_Exemption: //
			//
			cli.haltForeachFailure(errors.New(r.Exemption.Failure))
			return
		default:
			err := cli.actedObtainReply(res)
			if err != nil {
				cli.haltForeachFailure(err)
				return
			}
		}
	}
}

func (cli *portCustomer) monitorSolicit(requestresponse *RequestResult) {
	//
	//
	if !cli.EqualsActive() {
		return
	}

	cli.mtx.Lock()
	defer cli.mtx.Unlock()
	cli.requestRelayed.PushBack(requestresponse)
}

func (cli *portCustomer) actedObtainReply(res *kinds.Reply) error {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()

	//
	following := cli.requestRelayed.Front()
	if following == nil {
		return FaultUnforeseenReply{Reply: *res, Rationale: "REDACTED"}
	}

	requestresponse := following.Value.(*RequestResult)
	if !resultAlignsRequest(requestresponse.Solicit, res) {
		return FaultUnforeseenReply{Reply: *res, Rationale: fmt.Sprintf("REDACTED", requestresponse.Solicit.Datum)}
	}

	requestresponse.Reply = res
	requestresponse.Done()            //
	cli.requestRelayed.Remove(following) //

	//
	if cli.resultClbk != nil {
		cli.resultClbk(requestresponse.Solicit, res)
	}

	//
	//
	//
	//
	requestresponse.ExecuteClbk()

	return nil
}

//

func (cli *portCustomer) Purge(ctx context.Context) error {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitPurge())
	if err != nil {
		return err
	}
	requestResult.Wait()
	return nil
}

func (cli *portCustomer) Reverberate(ctx context.Context, msg string) (*kinds.ReplyReverberate, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitReverberate(msg))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainReverberate(), cli.Failure()
}

func (cli *portCustomer) Details(ctx context.Context, req *kinds.SolicitDetails) (*kinds.ReplyDetails, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitDetails(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainDetails(), cli.Failure()
}

func (cli *portCustomer) InspectTransfer(ctx context.Context, req *kinds.SolicitInspectTransfer) (*kinds.ReplyInspectTransfer, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitInspectTransfer(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainInspectTransfer(), cli.Failure()
}

func (cli *portCustomer) AppendTransfer(ctx context.Context, req *kinds.SolicitAppendTransfer) (*kinds.ReplyAppendTransfer, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitAppendTransfer(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainAppendTransfer(), cli.Failure()
}

func (cli *portCustomer) HarvestTrans(ctx context.Context, req *kinds.SolicitHarvestTrans) (*kinds.ReplyHarvestTrans, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitHarvestTrans(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainHarvestTrans(), cli.Failure()
}

func (cli *portCustomer) Inquire(ctx context.Context, req *kinds.SolicitInquire) (*kinds.ReplyInquire, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitInquire(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainInquire(), cli.Failure()
}

func (cli *portCustomer) Endorse(ctx context.Context, _ *kinds.SolicitEndorse) (*kinds.ReplyEndorse, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitEndorse())
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainEndorse(), cli.Failure()
}

func (cli *portCustomer) InitializeSuccession(ctx context.Context, req *kinds.SolicitInitializeSuccession) (*kinds.ReplyInitializeSuccession, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitInitializeSuccession(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainInitializeSuccession(), cli.Failure()
}

func (cli *portCustomer) CollectionImages(ctx context.Context, req *kinds.SolicitCollectionImages) (*kinds.ReplyCatalogImages, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitCatalogImages(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainCatalogImages(), cli.Failure()
}

func (cli *portCustomer) ExtendImage(ctx context.Context, req *kinds.SolicitExtendImage) (*kinds.ReplyExtendImage, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitExtendImage(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainExtendImage(), cli.Failure()
}

func (cli *portCustomer) FetchImageSegment(ctx context.Context, req *kinds.SolicitFetchImageSegment) (*kinds.ReplyFetchImageSegment, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitFetchImageSegment(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainFetchImageSegment(), cli.Failure()
}

func (cli *portCustomer) ExecuteImageSegment(ctx context.Context, req *kinds.SolicitExecuteImageSegment) (*kinds.ReplyExecuteImageSegment, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitExecuteImageSegment(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainExecuteImageSegment(), cli.Failure()
}

func (cli *portCustomer) ArrangeNomination(ctx context.Context, req *kinds.SolicitArrangeNomination) (*kinds.ReplyArrangeNomination, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitArrangeNomination(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainArrangeNomination(), cli.Failure()
}

func (cli *portCustomer) HandleNomination(ctx context.Context, req *kinds.SolicitHandleNomination) (*kinds.ReplyHandleNomination, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitHandleNomination(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainHandleNomination(), cli.Failure()
}

func (cli *portCustomer) BroadenBallot(ctx context.Context, req *kinds.SolicitBroadenBallot) (*kinds.ReplyBroadenBallot, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitBroadenBallot(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainBroadenBallot(), cli.Failure()
}

func (cli *portCustomer) ValidateBallotAddition(ctx context.Context, req *kinds.SolicitValidateBallotAddition) (*kinds.ReplyValidateBallotAddition, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitValidateBallotAddition(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainValidateBallotAddition(), cli.Failure()
}

func (cli *portCustomer) CulminateLedger(ctx context.Context, req *kinds.SolicitCulminateLedger) (*kinds.ReplyCulminateLedger, error) {
	requestResult, err := cli.stagingSolicit(ctx, kinds.TowardSolicitCulminateLedger(req))
	if err != nil {
		return nil, err
	}
	if err := cli.Purge(ctx); err != nil {
		return nil, err
	}
	return requestResult.Reply.ObtainCulminateLedger(), cli.Failure()
}

func (cli *portCustomer) stagingSolicit(ctx context.Context, req *kinds.Solicit) (*RequestResult, error) {
	requestresponse := FreshRequestResult(req)

	//
	select {
	case cli.requestStaging <- requestresponse:
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	//
	switch req.Datum.(type) {
	case *kinds.Solicit_Purge:
		cli.purgeClock.Deassign()
	default:
		cli.purgeClock.Set()
	}

	return requestresponse, nil
}

//
//
func (cli *portCustomer) purgeStaging() {
	cli.mtx.Lock()
	defer cli.mtx.Unlock()

	//
	for req := cli.requestRelayed.Front(); req != nil; req = req.Next() {
		requestresponse := req.Value.(*RequestResult)
		requestresponse.Done()
	}

	//
Cycle:
	for {
		select {
		case requestresponse := <-cli.requestStaging:
			requestresponse.Done()
		default:
			break Cycle
		}
	}
}

//

func resultAlignsRequest(req *kinds.Solicit, res *kinds.Reply) (ok bool) {
	switch req.Datum.(type) {
	case *kinds.Solicit_Reverberate:
		_, ok = res.Datum.(*kinds.Reply_Reverberate)
	case *kinds.Solicit_Purge:
		_, ok = res.Datum.(*kinds.Reply_Purge)
	case *kinds.Solicit_Details:
		_, ok = res.Datum.(*kinds.Reply_Details)
	case *kinds.Solicit_Inspecttrans:
		_, ok = res.Datum.(*kinds.Reply_Inspecttrans)
	case *kinds.Solicit_Endorse:
		_, ok = res.Datum.(*kinds.Reply_Endorse)
	case *kinds.Solicit_Inquire:
		_, ok = res.Datum.(*kinds.Reply_Inquire)
	case *kinds.Solicit_Initiatechain:
		_, ok = res.Datum.(*kinds.Reply_Initiatechain)
	case *kinds.Solicit_Executeimagefragment:
		_, ok = res.Datum.(*kinds.Reply_Executeimagefragment)
	case *kinds.Solicit_Loadimagefragment:
		_, ok = res.Datum.(*kinds.Reply_Loadimagefragment)
	case *kinds.Solicit_Catalogimages:
		_, ok = res.Datum.(*kinds.Reply_Catalogimages)
	case *kinds.Solicit_Extendimage:
		_, ok = res.Datum.(*kinds.Reply_Extendimage)
	case *kinds.Solicit_Extendballot:
		_, ok = res.Datum.(*kinds.Reply_Extendballot)
	case *kinds.Solicit_Verifyballotaddition:
		_, ok = res.Datum.(*kinds.Reply_Verifyballotaddition)
	case *kinds.Solicit_Prepareitem:
		_, ok = res.Datum.(*kinds.Reply_Prepareitem)
	case *kinds.Solicit_Executeitem:
		_, ok = res.Datum.(*kinds.Reply_Executeitem)
	case *kinds.Solicit_Finalizeledger:
		_, ok = res.Datum.(*kinds.Reply_Finalizeledger)
	}
	return ok
}

func (cli *portCustomer) haltForeachFailure(err error) {
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
