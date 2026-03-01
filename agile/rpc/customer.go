package rpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/security/hashmap"
	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	strongarithmetic "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arithmetic"
	facility "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	customeriface "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/customer"
	ktypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/base/kinds"
	remoteifacetypes "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/status"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds"
	strongminderrors "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/kinds/faults"
)

var faultNegativeEitherNullAltitude = errors.New("REDACTED")

//
type TokenRouteMethod func(route string, key []byte) (hashmap.TokenRoute, error)

//
//
//
type AgileCustomer interface {
	SuccessionUUID() string
	Revise(ctx context.Context, now time.Time) (*kinds.AgileLedger, error)
	ValidateAgileLedgerLocatedAltitude(ctx context.Context, altitude int64, now time.Time) (*kinds.AgileLedger, error)
	ReliableAgileLedger(altitude int64) (*kinds.AgileLedger, error)
}

var _ customeriface.Customer = (*Customer)(nil)

//
//
//
type Customer struct {
	facility.FoundationFacility

	following customeriface.Customer
	lc   AgileCustomer

	//
	prt       *hashmap.AttestationExecution
	tokenRouteProc TokenRouteMethod
}

var _ customeriface.Customer = (*Customer)(nil)

//
type Selection func(*Customer)

//
//
//
func TokenRouteProc(fn TokenRouteMethod) Selection {
	return func(c *Customer) {
		c.tokenRouteProc = fn
	}
}

//
//
//
func FallbackHashmapTokenRouteProc() TokenRouteMethod {
	//
	depotAliasPattern := regexp.MustCompile("REDACTED")

	return func(route string, key []byte) (hashmap.TokenRoute, error) {
		aligns := depotAliasPattern.FindStringSubmatch(route)
		if len(aligns) != 2 {
			return nil, fmt.Errorf("REDACTED", route, depotAliasPattern)
		}
		depotAlias := aligns[1]

		kp := hashmap.TokenRoute{}
		kp = kp.AttachToken([]byte(depotAlias), hashmap.TokenSerializationWebroute)
		kp = kp.AttachToken(key, hashmap.TokenSerializationWebroute)
		return kp, nil
	}
}

//
func FreshCustomer(following customeriface.Customer, lc AgileCustomer, choices ...Selection) *Customer {
	c := &Customer{
		following: following,
		lc:   lc,
		prt:  hashmap.FallbackAttestationExecution(),
	}
	c.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", c)
	for _, o := range choices {
		o(c)
	}
	return c
}

func (c *Customer) UponInitiate() error {
	if !c.following.EqualsActive() {
		return c.following.Initiate()
	}
	return nil
}

func (c *Customer) UponHalt() {
	if c.following.EqualsActive() {
		if err := c.following.Halt(); err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", err)
		}
	}
}

func (c *Customer) Condition(ctx context.Context) (*ktypes.OutcomeCondition, error) {
	return c.following.Condition(ctx)
}

func (c *Customer) IfaceDetails(ctx context.Context) (*ktypes.OutcomeIfaceDetails, error) {
	return c.following.IfaceDetails(ctx)
}

//
func (c *Customer) IfaceInquire(ctx context.Context, route string, data tendermintoctets.HexadecimalOctets) (*ktypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireUsingChoices(ctx, route, data, customeriface.FallbackIfaceInquireChoices)
}

//
func (c *Customer) IfaceInquireUsingChoices(ctx context.Context, route string, data tendermintoctets.HexadecimalOctets,
	choices customeriface.IfaceInquireChoices,
) (*ktypes.OutcomeIfaceInquire, error) {
	//
	choices.Validate = true

	res, err := c.following.IfaceInquireUsingChoices(ctx, route, data, choices)
	if err != nil {
		return nil, err
	}
	reply := res.Reply

	//
	if reply.EqualsFault() {
		return nil, fmt.Errorf("REDACTED", reply.Cipher)
	}
	if len(reply.Key) == 0 {
		return nil, strongminderrors.FaultMandatoryAttribute{Attribute: "REDACTED"}
	}
	if reply.AttestationActions == nil || len(reply.AttestationActions.Ops) == 0 {
		return nil, errors.New("REDACTED")
	}
	if reply.Altitude <= 0 {
		return nil, faultNegativeEitherNullAltitude
	}

	//
	//
	followingAltitude := reply.Altitude + 1
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &followingAltitude)
	if err != nil {
		return nil, err
	}

	//
	if reply.Datum != nil {
		//
		if c.tokenRouteProc == nil {
			return nil, errors.New("REDACTED")
		}

		kp, err := c.tokenRouteProc(route, reply.Key)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		//
		err = c.prt.ValidateDatum(reply.AttestationActions, l.PlatformDigest, kp.Text(), reply.Datum)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	} else { //
		err = c.prt.ValidateOmission(reply.AttestationActions, l.PlatformDigest, string(reply.Key))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	return &ktypes.OutcomeIfaceInquire{Reply: reply}, nil
}

func (c *Customer) MulticastTransferEndorse(ctx context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransferEndorse, error) {
	return c.following.MulticastTransferEndorse(ctx, tx)
}

func (c *Customer) MulticastTransferAsyncronous(ctx context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.following.MulticastTransferAsyncronous(ctx, tx)
}

func (c *Customer) MulticastTransferChronize(ctx context.Context, tx kinds.Tx) (*ktypes.OutcomeMulticastTransfer, error) {
	return c.following.MulticastTransferChronize(ctx, tx)
}

func (c *Customer) PendingTrans(ctx context.Context, threshold *int) (*ktypes.OutcomePendingTrans, error) {
	return c.following.PendingTrans(ctx, threshold)
}

func (c *Customer) CountPendingTrans(ctx context.Context) (*ktypes.OutcomePendingTrans, error) {
	return c.following.CountPendingTrans(ctx)
}

func (c *Customer) InspectTransfer(ctx context.Context, tx kinds.Tx) (*ktypes.OutcomeInspectTransfer, error) {
	return c.following.InspectTransfer(ctx, tx)
}

func (c *Customer) NetworkDetails(ctx context.Context) (*ktypes.OutcomeNetworkDetails, error) {
	return c.following.NetworkDetails(ctx)
}

func (c *Customer) ExportAgreementStatus(ctx context.Context) (*ktypes.OutcomeExportAgreementStatus, error) {
	return c.following.ExportAgreementStatus(ctx)
}

func (c *Customer) AgreementStatus(ctx context.Context) (*ktypes.OutcomeAgreementStatus, error) {
	return c.following.AgreementStatus(ctx)
}

func (c *Customer) AgreementSettings(ctx context.Context, altitude *int64) (*ktypes.OutcomeAgreementParameters, error) {
	res, err := c.following.AgreementSettings(ctx, altitude)
	if err != nil {
		return nil, err
	}

	//
	if err := res.AgreementSettings.CertifyFundamental(); err != nil {
		return nil, err
	}
	if res.LedgerAltitude <= 0 {
		return nil, faultNegativeEitherNullAltitude
	}

	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &res.LedgerAltitude)
	if err != nil {
		return nil, err
	}

	//
	if cH, tH := res.AgreementSettings.Digest(), l.AgreementDigest; !bytes.Equal(cH, tH) {
		return nil, fmt.Errorf("REDACTED",
			cH, tH)
	}

	return res, nil
}

func (c *Customer) Vitality(ctx context.Context) (*ktypes.OutcomeVitality, error) {
	return c.following.Vitality(ctx)
}

//
//
func (c *Customer) LedgerchainDetails(ctx context.Context, minimumAltitude, maximumAltitude int64) (*ktypes.OutcomeLedgerchainDetails, error) {
	res, err := c.following.LedgerchainDetails(ctx, minimumAltitude, maximumAltitude)
	if err != nil {
		return nil, err
	}

	//
	for i, summary := range res.LedgerMetadata {
		if summary == nil {
			return nil, fmt.Errorf("REDACTED", i)
		}
		if err := summary.CertifyFundamental(); err != nil {
			return nil, fmt.Errorf("REDACTED", i, err)
		}
	}

	//
	if len(res.LedgerMetadata) > 0 {
		finalAltitude := res.LedgerMetadata[len(res.LedgerMetadata)-1].Heading.Altitude
		if _, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &finalAltitude); err != nil {
			return nil, err
		}
	}

	//
	for _, summary := range res.LedgerMetadata {
		h, err := c.lc.ReliableAgileLedger(summary.Heading.Altitude)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", summary.Heading.Altitude, err)
		}
		if bmH, tH := summary.Heading.Digest(), h.Digest(); !bytes.Equal(bmH, tH) {
			return nil, fmt.Errorf("REDACTED",
				bmH, tH)
		}
	}

	return res, nil
}

func (c *Customer) Inauguration(ctx context.Context) (*ktypes.OutcomeInauguration, error) {
	return c.following.Inauguration(ctx)
}

func (c *Customer) InaugurationSegmented(ctx context.Context, id uint) (*ktypes.OutcomeInaugurationSegment, error) {
	return c.following.InaugurationSegmented(ctx, id)
}

//
func (c *Customer) Ledger(ctx context.Context, altitude *int64) (*ktypes.OutcomeLedger, error) {
	res, err := c.following.Ledger(ctx, altitude)
	if err != nil {
		return nil, err
	}

	//
	if err := res.LedgerUUID.CertifyFundamental(); err != nil {
		return nil, err
	}
	if err := res.Ledger.CertifyFundamental(); err != nil {
		return nil, err
	}
	if bmH, bH := res.LedgerUUID.Digest, res.Ledger.Digest(); !bytes.Equal(bmH, bH) {
		return nil, fmt.Errorf("REDACTED",
			bmH, bH)
	}

	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &res.Ledger.Altitude)
	if err != nil {
		return nil, err
	}

	//
	if bH, tH := res.Ledger.Digest(), l.Digest(); !bytes.Equal(bH, tH) {
		return nil, fmt.Errorf("REDACTED",
			bH, tH)
	}

	return res, nil
}

//
func (c *Customer) LedgerViaDigest(ctx context.Context, digest []byte) (*ktypes.OutcomeLedger, error) {
	res, err := c.following.LedgerViaDigest(ctx, digest)
	if err != nil {
		return nil, err
	}

	//
	if err := res.LedgerUUID.CertifyFundamental(); err != nil {
		return nil, err
	}
	if err := res.Ledger.CertifyFundamental(); err != nil {
		return nil, err
	}
	if bmH, bH := res.LedgerUUID.Digest, res.Ledger.Digest(); !bytes.Equal(bmH, bH) {
		return nil, fmt.Errorf("REDACTED",
			bmH, bH)
	}

	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &res.Ledger.Altitude)
	if err != nil {
		return nil, err
	}

	//
	if bH, tH := res.Ledger.Digest(), l.Digest(); !bytes.Equal(bH, tH) {
		return nil, fmt.Errorf("REDACTED",
			bH, tH)
	}

	return res, nil
}

//
//
//
func (c *Customer) LedgerOutcomes(ctx context.Context, altitude *int64) (*ktypes.OutcomeLedgerOutcomes, error) {
	var h int64
	if altitude == nil {
		res, err := c.following.Condition(ctx)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		//
		//
		h = res.ChronizeDetails.NewestLedgerAltitude - 1
	} else {
		h = *altitude
	}

	res, err := c.following.LedgerOutcomes(ctx, &h)
	if err != nil {
		return nil, err
	}

	//
	if res.Altitude <= 0 {
		return nil, faultNegativeEitherNullAltitude
	}

	//
	followingAltitude := h + 1
	reliableLedger, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &followingAltitude)
	if err != nil {
		return nil, err
	}

	//
	rH := status.TransferOutcomesDigest(res.TransOutcomes)

	//
	if !bytes.Equal(rH, reliableLedger.FinalOutcomesDigest) {
		return nil, fmt.Errorf("REDACTED",
			rH, reliableLedger.FinalOutcomesDigest)
	}

	return res, nil
}

//
func (c *Customer) Heading(ctx context.Context, altitude *int64) (*ktypes.OutcomeHeadline, error) {
	lb, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, altitude)
	if err != nil {
		return nil, err
	}

	return &ktypes.OutcomeHeadline{Heading: lb.Heading}, nil
}

//
func (c *Customer) HeadingViaDigest(ctx context.Context, digest tendermintoctets.HexadecimalOctets) (*ktypes.OutcomeHeadline, error) {
	res, err := c.following.HeadingViaDigest(ctx, digest)
	if err != nil {
		return nil, err
	}

	if err := res.Heading.CertifyFundamental(); err != nil {
		return nil, err
	}

	lb, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &res.Heading.Altitude)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(lb.Digest(), res.Heading.Digest()) {
		return nil, fmt.Errorf("REDACTED",
			lb.Digest(), res.Heading.Digest())
	}

	return res, nil
}

func (c *Customer) Endorse(ctx context.Context, altitude *int64) (*ktypes.OutcomeEndorse, error) {
	//
	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, altitude)
	if err != nil {
		return nil, err
	}

	return &ktypes.OutcomeEndorse{
		NotatedHeading:    *l.NotatedHeading,
		StandardEndorse: true,
	}, nil
}

//
//
func (c *Customer) Tx(ctx context.Context, digest []byte, ascertain bool) (*ktypes.OutcomeTransfer, error) {
	res, err := c.following.Tx(ctx, digest, ascertain)
	if err != nil || !ascertain {
		return res, err
	}

	//
	if res.Altitude <= 0 {
		return nil, faultNegativeEitherNullAltitude
	}

	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, &res.Altitude)
	if err != nil {
		return nil, err
	}

	//
	return res, res.Attestation.Certify(l.DataDigest)
}

func (c *Customer) TransferLookup(
	ctx context.Context,
	inquire string,
	ascertain bool,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeTransferLookup, error) {
	return c.following.TransferLookup(ctx, inquire, ascertain, screen, everyScreen, sequenceVia)
}

func (c *Customer) LedgerLookup(
	ctx context.Context,
	inquire string,
	screen, everyScreen *int,
	sequenceVia string,
) (*ktypes.OutcomeLedgerLookup, error) {
	return c.following.LedgerLookup(ctx, inquire, screen, everyScreen, sequenceVia)
}

//
func (c *Customer) Assessors(
	ctx context.Context,
	altitude *int64,
	screenReference, everyScreenReference *int,
) (*ktypes.OutcomeAssessors, error) {
	//
	//
	l, err := c.reviseAgileCustomerConditionalRequiredToward(ctx, altitude)
	if err != nil {
		return nil, err
	}

	sumTally := len(l.AssessorAssign.Assessors)
	everyScreen := certifyEveryScreen(everyScreenReference)
	screen, err := certifyScreen(screenReference, everyScreen, sumTally)
	if err != nil {
		return nil, err
	}

	omitTally := certifyOmitTally(screen, everyScreen)
	v := l.AssessorAssign.Assessors[omitTally : omitTally+strongarithmetic.MinimumInteger(everyScreen, sumTally-omitTally)]

	return &ktypes.OutcomeAssessors{
		LedgerAltitude: l.Altitude,
		Assessors:  v,
		Tally:       len(v),
		Sum:       sumTally,
	}, nil
}

func (c *Customer) MulticastProof(ctx context.Context, ev kinds.Proof) (*ktypes.OutcomeMulticastProof, error) {
	return c.following.MulticastProof(ctx, ev)
}

func (c *Customer) Listen(ctx context.Context, listener, inquire string,
	outputVolume ...int,
) (out <-chan ktypes.OutcomeIncident, err error) {
	return c.following.Listen(ctx, listener, inquire, outputVolume...)
}

func (c *Customer) Unlisten(ctx context.Context, listener, inquire string) error {
	return c.following.Unlisten(ctx, listener, inquire)
}

func (c *Customer) UnlistenEvery(ctx context.Context, listener string) error {
	return c.following.UnlistenEvery(ctx, listener)
}

func (c *Customer) reviseAgileCustomerConditionalRequiredToward(ctx context.Context, altitude *int64) (*kinds.AgileLedger, error) {
	var (
		l   *kinds.AgileLedger
		err error
	)
	if altitude == nil {
		l, err = c.lc.Revise(ctx, time.Now())
	} else {
		l, err = c.lc.ValidateAgileLedgerLocatedAltitude(ctx, *altitude, time.Now())
	}
	if err != nil {
		return nil, fmt.Errorf("REDACTED", *altitude, err)
	}
	return l, nil
}

func (c *Customer) EnrollActionDeserializer(typ string, dec hashmap.ActionDeserializer) {
	c.prt.EnrollActionDeserializer(typ, dec)
}

//
//
//
func (c *Customer) ListenSocket(ctx *remoteifacetypes.Env, inquire string) (*ktypes.OutcomeListen, error) {
	out, err := c.following.Listen(context.Background(), ctx.DistantLocation(), inquire)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case outcomeIncident := <-out:
				//
				//
				ctx.SocketLink.AttemptPersistRemoteReply(
					remoteifacetypes.FreshRemoteTriumphReply(
						remoteifacetypes.JsonifaceTextUUID(fmt.Sprintf("REDACTED", ctx.JSNRequest.ID)),
						outcomeIncident,
					))
			case <-c.Exit():
				return
			}
		}
	}()

	return &ktypes.OutcomeListen{}, nil
}

//
//
func (c *Customer) UnlistenSocket(ctx *remoteifacetypes.Env, inquire string) (*ktypes.OutcomeUnlisten, error) {
	err := c.following.Unlisten(context.Background(), ctx.DistantLocation(), inquire)
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeUnlisten{}, nil
}

//
//
func (c *Customer) UnlistenEverySocket(ctx *remoteifacetypes.Env) (*ktypes.OutcomeUnlisten, error) {
	err := c.following.UnlistenEvery(context.Background(), ctx.DistantLocation())
	if err != nil {
		return nil, err
	}
	return &ktypes.OutcomeUnlisten{}, nil
}

//
const (
	//
	fallbackEveryScreen = 30
	maximumEveryScreen     = 100
)

func certifyScreen(screenReference *int, everyScreen, sumTally int) (int, error) {
	if everyScreen < 1 {
		panic(fmt.Sprintf("REDACTED", everyScreen))
	}

	if screenReference == nil { //
		return 1, nil
	}

	displays := ((sumTally - 1) / everyScreen) + 1
	if displays == 0 {
		displays = 1 //
	}
	screen := *screenReference
	if screen <= 0 || screen > displays {
		return 1, fmt.Errorf("REDACTED", displays, screen)
	}

	return screen, nil
}

func certifyEveryScreen(everyScreenReference *int) int {
	if everyScreenReference == nil { //
		return fallbackEveryScreen
	}

	everyScreen := *everyScreenReference
	if everyScreen < 1 {
		return fallbackEveryScreen
	} else if everyScreen > maximumEveryScreen {
		return maximumEveryScreen
	}
	return everyScreen
}

func certifyOmitTally(screen, everyScreen int) int {
	omitTally := (screen - 1) * everyScreen
	if omitTally < 0 {
		return 0
	}

	return omitTally
}
