package httpsvc

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	tendermintpubsub "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	jsoncustomer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/customer"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
)

/**
r
.

.
s
.

.
s
o
d
.

o
e
.

:

)
{
r
}

s
)
{
r
}
)

)
{
r
}

t
*/
type Httpsvc struct {
	distant string
	rpc    *jsoncustomer.Customer

	*foundationRemoteCustomer
	*SocketIncidents
}

//
//
//
//
//
//
//
//
//
//
type ClusterHttpsvc struct {
	remoteCluster *jsoncustomer.SolicitCluster
	*foundationRemoteCustomer
}

//
//
//
type ifaceCustomer interface {
	customeriface.IfaceCustomer
	customeriface.ChronicleCustomer
	customeriface.FabricCustomer
	customeriface.AttestCustomer
	customeriface.ConditionCustomer
}

//
//
type foundationRemoteCustomer struct {
	invoker jsoncustomer.Invoker
}

var (
	_ ifaceCustomer = (*Httpsvc)(nil)
	_ ifaceCustomer = (*ClusterHttpsvc)(nil)
	_ ifaceCustomer = (*foundationRemoteCustomer)(nil)
)

//
//

//
//
//
func New(distant, socketGateway string) (*Httpsvc, error) {
	httpsvcCustomer, err := jsoncustomer.FallbackHttpsvcCustomer(distant)
	if err != nil {
		return nil, err
	}
	return FreshUsingCustomer(distant, socketGateway, httpsvcCustomer)
}

//
func FreshUsingDeadline(distant, socketGateway string, deadline uint) (*Httpsvc, error) {
	httpsvcCustomer, err := jsoncustomer.FallbackHttpsvcCustomer(distant)
	if err != nil {
		return nil, err
	}
	httpsvcCustomer.Timeout = time.Duration(deadline) * time.Second
	return FreshUsingCustomer(distant, socketGateway, httpsvcCustomer)
}

//
//
func FreshUsingCustomer(distant, socketGateway string, customer *http.Client) (*Httpsvc, error) {
	if customer == nil {
		panic("REDACTED")
	}

	rc, err := jsoncustomer.FreshUsingHttpsvcCustomer(distant, customer)
	if err != nil {
		return nil, err
	}

	socketIncidents, err := freshSocketIncidents(distant, socketGateway)
	if err != nil {
		return nil, err
	}

	httpsvcCustomer := &Httpsvc{
		rpc:           rc,
		distant:        distant,
		foundationRemoteCustomer: &foundationRemoteCustomer{invoker: rc},
		SocketIncidents:      socketIncidents,
	}

	return httpsvcCustomer, nil
}

var _ customeriface.Customer = (*Httpsvc)(nil)

//
func (c *Httpsvc) AssignTracer(l log.Tracer) {
	c.SocketIncidents.AssignTracer(l)
}

//
func (c *Httpsvc) Distant() string {
	return c.distant
}

//
func (c *Httpsvc) FreshCluster() *ClusterHttpsvc {
	remoteCluster := c.rpc.FreshSolicitCluster()
	return &ClusterHttpsvc{
		remoteCluster: remoteCluster,
		foundationRemoteCustomer: &foundationRemoteCustomer{
			invoker: remoteCluster,
		},
	}
}

//
//

//
//
//
//
func (b *ClusterHttpsvc) Transmit(ctx context.Context) ([]any, error) {
	return b.remoteCluster.Transmit(ctx)
}

//
//
func (b *ClusterHttpsvc) Flush() int {
	return b.remoteCluster.Flush()
}

//
func (b *ClusterHttpsvc) Tally() int {
	return b.remoteCluster.Tally()
}

//
//

func (c *foundationRemoteCustomer) Condition(ctx context.Context) (*ktypes.OutcomeCondition, error) {
	outcome := new(ktypes.OutcomeCondition)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *foundationRemoteCustomer) IfaceDetails(ctx context.Context) (*ktypes.OutcomeIfaceDetails, error) {
	outcome := new(ktypes.OutcomeIfaceDetails)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *foundationRemoteCustomer) IfaceInquire(
	ctx context.Context,
	route string,
	data octets.HexadecimalOctets,
) (*ktypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireUsingChoices(ctx, route, data, customeriface.FallbackIfaceInquireChoices)
}

func (c *foundationRemoteCustomer) IfaceInquireUsingChoices(
	ctx context.Context,
	route string,
	data octets.HexadecimalOctets,
	choices customeriface.IfaceInquireChoices,
) (*ktypes.OutcomeIfaceInquire, error) {
	outcome := new(ktypes.OutcomeIfaceInquire)
	_, err := c.invoker.Invocation(ctx, "REDACTED",
		map[string]any{"REDACTED": route, "REDACTED": data, "REDACTED": choices.Altitude, "REDACTED": choices.Validate},
		outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *foundationRemoteCustomer) MulticastTransferEndorse(
	ctx context.Context,
	tx kinds.Tx,
) (*ktypes.OutcomeMulticastTransferEndorse, error) {
	outcome := new(ktypes.OutcomeMulticastTransferEndorse)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) MulticastTransferAsyncronous(
	ctx context.Context,
	tx kinds.Tx,
) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.multicastTransfer(ctx, "REDACTED", tx)
}

func (c *foundationRemoteCustomer) MulticastTransferChronize(
	ctx context.Context,
	tx kinds.Tx,
) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.multicastTransfer(ctx, "REDACTED", tx)
}

func (c *foundationRemoteCustomer) multicastTransfer(
	ctx context.Context,
	path string,
	tx kinds.Tx,
) (*ktypes.OutcomeMulticastTransfer, error) {
	outcome := new(ktypes.OutcomeMulticastTransfer)
	_, err := c.invoker.Invocation(ctx, path, map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) PendingTrans(
	ctx context.Context,
	threshold *int,
) (*ktypes.OutcomePendingTrans, error) {
	outcome := new(ktypes.OutcomePendingTrans)
	parameters := make(map[string]any)
	if threshold != nil {
		parameters["REDACTED"] = threshold
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) CountPendingTrans(ctx context.Context) (*ktypes.OutcomePendingTrans, error) {
	outcome := new(ktypes.OutcomePendingTrans)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) InspectTransfer(ctx context.Context, tx kinds.Tx) (*ktypes.OutcomeInspectTransfer, error) {
	outcome := new(ktypes.OutcomeInspectTransfer)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{"REDACTED": tx}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) NetworkDetails(ctx context.Context) (*ktypes.OutcomeNetworkDetails, error) {
	outcome := new(ktypes.OutcomeNetworkDetails)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) ExportAgreementStatus(ctx context.Context) (*ktypes.OutcomeExportAgreementStatus, error) {
	outcome := new(ktypes.OutcomeExportAgreementStatus)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) AgreementStatus(ctx context.Context) (*ktypes.OutcomeAgreementStatus, error) {
	outcome := new(ktypes.OutcomeAgreementStatus)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) AgreementSettings(
	ctx context.Context,
	altitude *int64,
) (*ktypes.OutcomeAgreementParameters, error) {
	outcome := new(ktypes.OutcomeAgreementParameters)
	parameters := make(map[string]any)
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Vitality(ctx context.Context) (*ktypes.OutcomeVitality, error) {
	outcome := new(ktypes.OutcomeVitality)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) LedgerchainDetails(
	ctx context.Context,
	minimumAltitude,
	maximumAltitude int64,
) (*ktypes.OutcomeLedgerchainDetails, error) {
	outcome := new(ktypes.OutcomeLedgerchainDetails)
	_, err := c.invoker.Invocation(ctx, "REDACTED",
		map[string]any{"REDACTED": minimumAltitude, "REDACTED": maximumAltitude},
		outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Inauguration(ctx context.Context) (*ktypes.OutcomeInauguration, error) {
	outcome := new(ktypes.OutcomeInauguration)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) InaugurationSegmented(ctx context.Context, id uint) (*ktypes.OutcomeInaugurationSegment, error) {
	outcome := new(ktypes.OutcomeInaugurationSegment)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{"REDACTED": id}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Ledger(ctx context.Context, altitude *int64) (*ktypes.OutcomeLedger, error) {
	outcome := new(ktypes.OutcomeLedger)
	parameters := make(map[string]any)
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) LedgerViaDigest(ctx context.Context, digest []byte) (*ktypes.OutcomeLedger, error) {
	outcome := new(ktypes.OutcomeLedger)
	parameters := map[string]any{
		"REDACTED": digest,
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) LedgerOutcomes(
	ctx context.Context,
	altitude *int64,
) (*ktypes.OutcomeLedgerOutcomes, error) {
	outcome := new(ktypes.OutcomeLedgerOutcomes)
	parameters := make(map[string]any)
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Heading(ctx context.Context, altitude *int64) (*ktypes.OutcomeHeadline, error) {
	outcome := new(ktypes.OutcomeHeadline)
	parameters := make(map[string]any)
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) HeadingViaDigest(ctx context.Context, digest octets.HexadecimalOctets) (*ktypes.OutcomeHeadline, error) {
	outcome := new(ktypes.OutcomeHeadline)
	parameters := map[string]any{
		"REDACTED": digest,
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Endorse(ctx context.Context, altitude *int64) (*ktypes.OutcomeEndorse, error) {
	outcome := new(ktypes.OutcomeEndorse)
	parameters := make(map[string]any)
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) Tx(ctx context.Context, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error) {
	outcome := new(ktypes.OutcomeTransfer)
	parameters := map[string]any{
		"REDACTED":  digest,
		"REDACTED": ascertain,
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) TransferLookup(
	ctx context.Context,
	inquire string,
	ascertain bool,
	screen,
	everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeTransferLookup, error) {
	outcome := new(ktypes.OutcomeTransferLookup)
	parameters := map[string]any{
		"REDACTED":    inquire,
		"REDACTED":    ascertain,
		"REDACTED": sequenceVia,
	}

	if screen != nil {
		parameters["REDACTED"] = screen
	}
	if everyScreen != nil {
		parameters["REDACTED"] = everyScreen
	}

	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *foundationRemoteCustomer) LedgerLookup(
	ctx context.Context,
	inquire string,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeLedgerLookup, error) {
	outcome := new(ktypes.OutcomeLedgerLookup)
	parameters := map[string]any{
		"REDACTED":    inquire,
		"REDACTED": sequenceVia,
	}

	if screen != nil {
		parameters["REDACTED"] = screen
	}
	if everyScreen != nil {
		parameters["REDACTED"] = everyScreen
	}

	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}

	return outcome, nil
}

func (c *foundationRemoteCustomer) Assessors(
	ctx context.Context,
	altitude *int64,
	screen,
	everyScreen *int,
) (*ktypes.OutcomeAssessors, error) {
	outcome := new(ktypes.OutcomeAssessors)
	parameters := make(map[string]any)
	if screen != nil {
		parameters["REDACTED"] = screen
	}
	if everyScreen != nil {
		parameters["REDACTED"] = everyScreen
	}
	if altitude != nil {
		parameters["REDACTED"] = altitude
	}
	_, err := c.invoker.Invocation(ctx, "REDACTED", parameters, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

func (c *foundationRemoteCustomer) MulticastProof(
	ctx context.Context,
	ev kinds.Proof,
) (*ktypes.OutcomeMulticastProof, error) {
	outcome := new(ktypes.OutcomeMulticastProof)
	_, err := c.invoker.Invocation(ctx, "REDACTED", map[string]any{"REDACTED": ev}, outcome)
	if err != nil {
		return nil, err
	}
	return outcome, nil
}

//
//

var faultNegationActive = errors.New("REDACTED")

//
type SocketIncidents struct {
	facility.FoundationFacility
	distant   string
	gateway string
	ws       *jsoncustomer.SocketCustomer

	mtx           commitchronize.ReadwriteExclusion
	feeds map[string]chan ktypes.OutcomeIncident //
}

func freshSocketIncidents(distant, gateway string) (*SocketIncidents, error) {
	w := &SocketIncidents{
		gateway:      gateway,
		distant:        distant,
		feeds: make(map[string]chan ktypes.OutcomeIncident),
	}
	w.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", w)

	var err error
	w.ws, err = jsoncustomer.FreshSocket(w.distant, w.gateway, jsoncustomer.UponReestablish(func() {
		//
		w.reiterateFeedsSubsequent(0 * time.Second)
	}))
	if err != nil {
		return nil, err
	}
	w.ws.AssignTracer(w.Tracer)

	return w, nil
}

//
func (w *SocketIncidents) UponInitiate() error {
	if err := w.ws.Initiate(); err != nil {
		return err
	}

	go w.incidentObserver()

	return nil
}

//
func (w *SocketIncidents) UponHalt() {
	if err := w.ws.Halt(); err != nil {
		w.Tracer.Failure("REDACTED", "REDACTED", err)
	}
}

//
//
//
//
//
//
//
func (w *SocketIncidents) Listen(ctx context.Context, _, inquire string,
	outputVolume ...int,
) (out <-chan ktypes.OutcomeIncident, err error) {
	if !w.EqualsActive() {
		return nil, faultNegationActive
	}

	if err := w.ws.Listen(ctx, inquire); err != nil {
		return nil, err
	}

	outputCeiling := 1
	if len(outputVolume) > 0 {
		outputCeiling = outputVolume[0]
	}

	resultant := make(chan ktypes.OutcomeIncident, outputCeiling)
	w.mtx.Lock()
	//
	//
	w.feeds[inquire] = resultant
	w.mtx.Unlock()

	return resultant, nil
}

//
//
//
//
func (w *SocketIncidents) Unlisten(ctx context.Context, _, inquire string) error {
	if !w.EqualsActive() {
		return faultNegationActive
	}

	if err := w.ws.Unlisten(ctx, inquire); err != nil {
		return err
	}

	w.mtx.Lock()
	_, ok := w.feeds[inquire]
	if ok {
		delete(w.feeds, inquire)
	}
	w.mtx.Unlock()

	return nil
}

//
//
//
//
func (w *SocketIncidents) UnlistenEvery(ctx context.Context, _ string) error {
	if !w.EqualsActive() {
		return faultNegationActive
	}

	if err := w.ws.UnlistenEvery(ctx); err != nil {
		return err
	}

	w.mtx.Lock()
	w.feeds = make(map[string]chan ktypes.OutcomeIncident)
	w.mtx.Unlock()

	return nil
}

//
//
func (w *SocketIncidents) reiterateFeedsSubsequent(d time.Duration) {
	time.Sleep(d)

	w.mtx.RLock()
	defer w.mtx.RUnlock()
	for q := range w.feeds {
		err := w.ws.Listen(context.Background(), q)
		if err != nil {
			w.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

func equalsFaultEarlierListened(err error) bool {
	return strings.Contains(err.Error(), tendermintpubsub.FaultEarlierListened.Error())
}

func (w *SocketIncidents) incidentObserver() {
	for {
		select {
		case reply, ok := <-w.ws.RepliesChnl:
			if !ok {
				return
			}

			if reply.Failure != nil {
				w.Tracer.Failure("REDACTED", "REDACTED", reply.Failure.Failure())
				//
				//
				//
				//
				if !equalsFaultEarlierListened(reply.Failure) {
					//
					//
					w.reiterateFeedsSubsequent(1 * time.Second)
				}
				continue
			}

			outcome := new(ktypes.OutcomeIncident)
			err := strongmindjson.Decode(reply.Outcome, outcome)
			if err != nil {
				w.Tracer.Failure("REDACTED", "REDACTED", err)
				continue
			}

			w.mtx.RLock()
			if out, ok := w.feeds[outcome.Inquire]; ok {
				if cap(out) == 0 {
					out <- *outcome
				} else {
					select {
					case out <- *outcome:
					default:
						w.Tracer.Failure("REDACTED", "REDACTED", outcome, "REDACTED", outcome.Inquire)
					}
				}
			}
			w.mtx.RUnlock()
		case <-w.Exit():
			return
		}
	}
}
