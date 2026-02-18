package rpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/valkyrieworks/vault/merkle"
	cometbytes "github.com/valkyrieworks/utils/octets"
	cometmath "github.com/valkyrieworks/utils/math"
	daemon "github.com/valkyrieworks/utils/daemon"
	rpccustomer "github.com/valkyrieworks/rpc/customer"
	ctypes "github.com/valkyrieworks/rpc/core/kinds"
	rpctypes "github.com/valkyrieworks/rpc/jsonrpc/kinds"
	"github.com/valkyrieworks/status"
	"github.com/valkyrieworks/kinds"
	cometfaults "github.com/valkyrieworks/kinds/faults"
)

var errNegativeOrNilLevel = errors.New("REDACTED")

//
type KeyRouteFunction func(route string, key []byte) (merkle.KeyRoute, error)

//
//
//
type RapidCustomer interface {
	LedgerUID() string
	Modify(ctx context.Context, now time.Time) (*kinds.RapidLedger, error)
	ValidateRapidLedgerAtLevel(ctx context.Context, level int64, now time.Time) (*kinds.RapidLedger, error)
	ValidatedRapidLedger(level int64) (*kinds.RapidLedger, error)
}

var _ rpccustomer.Customer = (*Customer)(nil)

//
//
//
type Customer struct {
	daemon.RootDaemon

	following rpccustomer.Customer
	lc   RapidCustomer

	//
	prt       *merkle.EvidenceRuntime
	keyRouteFn KeyRouteFunction
}

var _ rpccustomer.Customer = (*Customer)(nil)

//
type Setting func(*Customer)

//
//
//
func KeyRouteFn(fn KeyRouteFunction) Setting {
	return func(c *Customer) {
		c.keyRouteFn = fn
	}
}

//
//
//
func StandardMerkleKeyRouteFn() KeyRouteFunction {
	//
	depotLabelPattern := regexp.MustCompile("REDACTED")

	return func(route string, key []byte) (merkle.KeyRoute, error) {
		aligns := depotLabelPattern.FindStringSubmatch(route)
		if len(aligns) != 2 {
			return nil, fmt.Errorf("REDACTED", route, depotLabelPattern)
		}
		depotLabel := aligns[1]

		kp := merkle.KeyRoute{}
		kp = kp.AttachKey([]byte(depotLabel), merkle.KeyCodecURL)
		kp = kp.AttachKey(key, merkle.KeyCodecURL)
		return kp, nil
	}
}

//
func NewCustomer(following rpccustomer.Customer, lc RapidCustomer, opts ...Setting) *Customer {
	c := &Customer{
		following: following,
		lc:   lc,
		prt:  merkle.StandardEvidenceRuntime(),
	}
	c.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", c)
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Customer) OnBegin() error {
	if !c.following.IsActive() {
		return c.following.Begin()
	}
	return nil
}

func (c *Customer) OnHalt() {
	if c.following.IsActive() {
		if err := c.following.Halt(); err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", err)
		}
	}
}

func (c *Customer) Status(ctx context.Context) (*ctypes.OutcomeState, error) {
	return c.following.Status(ctx)
}

func (c *Customer) IfaceDetails(ctx context.Context) (*ctypes.OutcomeIfaceDetails, error) {
	return c.following.IfaceDetails(ctx)
}

//
func (c *Customer) IfaceInquire(ctx context.Context, route string, data cometbytes.HexOctets) (*ctypes.OutcomeIfaceInquire, error) {
	return c.IfaceInquireWithSettings(ctx, route, data, rpccustomer.StandardIfaceInquireSettings)
}

//
func (c *Customer) IfaceInquireWithSettings(ctx context.Context, route string, data cometbytes.HexOctets,
	opts rpccustomer.IfaceInquireSettings,
) (*ctypes.OutcomeIfaceInquire, error) {
	//
	opts.Demonstrate = true

	res, err := c.following.IfaceInquireWithSettings(ctx, route, data, opts)
	if err != nil {
		return nil, err
	}
	reply := res.Reply

	//
	if reply.IsErr() {
		return nil, fmt.Errorf("REDACTED", reply.Code)
	}
	if len(reply.Key) == 0 {
		return nil, cometfaults.ErrMandatoryField{Field: "REDACTED"}
	}
	if reply.EvidenceActions == nil || len(reply.EvidenceActions.Ops) == 0 {
		return nil, errors.New("REDACTED")
	}
	if reply.Level <= 0 {
		return nil, errNegativeOrNilLevel
	}

	//
	//
	followingLevel := reply.Level + 1
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, &followingLevel)
	if err != nil {
		return nil, err
	}

	//
	if reply.Item != nil {
		//
		if c.keyRouteFn == nil {
			return nil, errors.New("REDACTED")
		}

		kp, err := c.keyRouteFn(route, reply.Key)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}

		//
		err = c.prt.ValidateItem(reply.EvidenceActions, l.ApplicationDigest, kp.String(), reply.Item)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	} else { //
		err = c.prt.ValidateOmission(reply.EvidenceActions, l.ApplicationDigest, string(reply.Key))
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
	}

	return &ctypes.OutcomeIfaceInquire{Reply: reply}, nil
}

func (c *Customer) MulticastTransferEndorse(ctx context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransferEndorse, error) {
	return c.following.MulticastTransferEndorse(ctx, tx)
}

func (c *Customer) MulticastTransferAsync(ctx context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.following.MulticastTransferAsync(ctx, tx)
}

func (c *Customer) MulticastTransferAlign(ctx context.Context, tx kinds.Tx) (*ctypes.OutcomeMulticastTransfer, error) {
	return c.following.MulticastTransferAlign(ctx, tx)
}

func (c *Customer) UnattestedTrans(ctx context.Context, ceiling *int) (*ctypes.OutcomeUnattestedTrans, error) {
	return c.following.UnattestedTrans(ctx, ceiling)
}

func (c *Customer) CountUnattestedTrans(ctx context.Context) (*ctypes.OutcomeUnattestedTrans, error) {
	return c.following.CountUnattestedTrans(ctx)
}

func (c *Customer) InspectTransfer(ctx context.Context, tx kinds.Tx) (*ctypes.OutcomeInspectTransfer, error) {
	return c.following.InspectTransfer(ctx, tx)
}

func (c *Customer) NetDetails(ctx context.Context) (*ctypes.OutcomeNetDetails, error) {
	return c.following.NetDetails(ctx)
}

func (c *Customer) ExportAgreementStatus(ctx context.Context) (*ctypes.OutcomeExportAgreementStatus, error) {
	return c.following.ExportAgreementStatus(ctx)
}

func (c *Customer) AgreementStatus(ctx context.Context) (*ctypes.OutcomeAgreementStatus, error) {
	return c.following.AgreementStatus(ctx)
}

func (c *Customer) AgreementOptions(ctx context.Context, level *int64) (*ctypes.OutcomeAgreementOptions, error) {
	res, err := c.following.AgreementOptions(ctx, level)
	if err != nil {
		return nil, err
	}

	//
	if err := res.AgreementOptions.CertifySimple(); err != nil {
		return nil, err
	}
	if res.LedgerLevel <= 0 {
		return nil, errNegativeOrNilLevel
	}

	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, &res.LedgerLevel)
	if err != nil {
		return nil, err
	}

	//
	if cH, tH := res.AgreementOptions.Digest(), l.AgreementDigest; !bytes.Equal(cH, tH) {
		return nil, fmt.Errorf("REDACTED",
			cH, tH)
	}

	return res, nil
}

func (c *Customer) Vitality(ctx context.Context) (*ctypes.OutcomeVitality, error) {
	return c.following.Vitality(ctx)
}

//
//
func (c *Customer) LedgerchainDetails(ctx context.Context, minimumLevel, maximumLevel int64) (*ctypes.OutcomeLedgerchainDetails, error) {
	res, err := c.following.LedgerchainDetails(ctx, minimumLevel, maximumLevel)
	if err != nil {
		return nil, err
	}

	//
	for i, meta := range res.LedgerMetadata {
		if meta == nil {
			return nil, fmt.Errorf("REDACTED", i)
		}
		if err := meta.CertifySimple(); err != nil {
			return nil, fmt.Errorf("REDACTED", i, err)
		}
	}

	//
	if len(res.LedgerMetadata) > 0 {
		finalLevel := res.LedgerMetadata[len(res.LedgerMetadata)-1].Heading.Level
		if _, err := c.modifyRapidCustomerIfRequiredTo(ctx, &finalLevel); err != nil {
			return nil, err
		}
	}

	//
	for _, meta := range res.LedgerMetadata {
		h, err := c.lc.ValidatedRapidLedger(meta.Heading.Level)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", meta.Heading.Level, err)
		}
		if bmH, tH := meta.Heading.Digest(), h.Digest(); !bytes.Equal(bmH, tH) {
			return nil, fmt.Errorf("REDACTED",
				bmH, tH)
		}
	}

	return res, nil
}

func (c *Customer) Origin(ctx context.Context) (*ctypes.OutcomeOrigin, error) {
	return c.following.Origin(ctx)
}

func (c *Customer) OriginSegmented(ctx context.Context, id uint) (*ctypes.OutcomeOriginSegment, error) {
	return c.following.OriginSegmented(ctx, id)
}

//
func (c *Customer) Ledger(ctx context.Context, level *int64) (*ctypes.OutcomeLedger, error) {
	res, err := c.following.Ledger(ctx, level)
	if err != nil {
		return nil, err
	}

	//
	if err := res.LedgerUID.CertifySimple(); err != nil {
		return nil, err
	}
	if err := res.Ledger.CertifySimple(); err != nil {
		return nil, err
	}
	if bmH, bH := res.LedgerUID.Digest, res.Ledger.Digest(); !bytes.Equal(bmH, bH) {
		return nil, fmt.Errorf("REDACTED",
			bmH, bH)
	}

	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, &res.Ledger.Level)
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
func (c *Customer) LedgerByDigest(ctx context.Context, digest []byte) (*ctypes.OutcomeLedger, error) {
	res, err := c.following.LedgerByDigest(ctx, digest)
	if err != nil {
		return nil, err
	}

	//
	if err := res.LedgerUID.CertifySimple(); err != nil {
		return nil, err
	}
	if err := res.Ledger.CertifySimple(); err != nil {
		return nil, err
	}
	if bmH, bH := res.LedgerUID.Digest, res.Ledger.Digest(); !bytes.Equal(bmH, bH) {
		return nil, fmt.Errorf("REDACTED",
			bmH, bH)
	}

	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, &res.Ledger.Level)
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
func (c *Customer) LedgerOutcomes(ctx context.Context, level *int64) (*ctypes.OutcomeLedgerOutcomes, error) {
	var h int64
	if level == nil {
		res, err := c.following.Status(ctx)
		if err != nil {
			return nil, fmt.Errorf("REDACTED", err)
		}
		//
		//
		h = res.AlignDetails.NewestLedgerLevel - 1
	} else {
		h = *level
	}

	res, err := c.following.LedgerOutcomes(ctx, &h)
	if err != nil {
		return nil, err
	}

	//
	if res.Level <= 0 {
		return nil, errNegativeOrNilLevel
	}

	//
	followingLevel := h + 1
	validatedLedger, err := c.modifyRapidCustomerIfRequiredTo(ctx, &followingLevel)
	if err != nil {
		return nil, err
	}

	//
	rH := status.TransferOutcomesDigest(res.TransOutcomes)

	//
	if !bytes.Equal(rH, validatedLedger.FinalOutcomesDigest) {
		return nil, fmt.Errorf("REDACTED",
			rH, validatedLedger.FinalOutcomesDigest)
	}

	return res, nil
}

//
func (c *Customer) Heading(ctx context.Context, level *int64) (*ctypes.OutcomeHeading, error) {
	lb, err := c.modifyRapidCustomerIfRequiredTo(ctx, level)
	if err != nil {
		return nil, err
	}

	return &ctypes.OutcomeHeading{Heading: lb.Heading}, nil
}

//
func (c *Customer) HeadingByDigest(ctx context.Context, digest cometbytes.HexOctets) (*ctypes.OutcomeHeading, error) {
	res, err := c.following.HeadingByDigest(ctx, digest)
	if err != nil {
		return nil, err
	}

	if err := res.Heading.CertifySimple(); err != nil {
		return nil, err
	}

	lb, err := c.modifyRapidCustomerIfRequiredTo(ctx, &res.Heading.Level)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(lb.Digest(), res.Heading.Digest()) {
		return nil, fmt.Errorf("REDACTED",
			lb.Digest(), res.Heading.Digest())
	}

	return res, nil
}

func (c *Customer) Endorse(ctx context.Context, level *int64) (*ctypes.OutcomeEndorse, error) {
	//
	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, level)
	if err != nil {
		return nil, err
	}

	return &ctypes.OutcomeEndorse{
		AttestedHeading:    *l.AttestedHeading,
		NormativeEndorse: true,
	}, nil
}

//
//
func (c *Customer) Tx(ctx context.Context, digest []byte, demonstrate bool) (*ctypes.OutcomeTransfer, error) {
	res, err := c.following.Tx(ctx, digest, demonstrate)
	if err != nil || !demonstrate {
		return res, err
	}

	//
	if res.Level <= 0 {
		return nil, errNegativeOrNilLevel
	}

	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, &res.Level)
	if err != nil {
		return nil, err
	}

	//
	return res, res.Attestation.Certify(l.DataDigest)
}

func (c *Customer) TransferScan(
	ctx context.Context,
	inquire string,
	demonstrate bool,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeTransferScan, error) {
	return c.following.TransferScan(ctx, inquire, demonstrate, screen, eachScreen, arrangeBy)
}

func (c *Customer) LedgerScan(
	ctx context.Context,
	inquire string,
	screen, eachScreen *int,
	arrangeBy string,
) (*ctypes.OutcomeLedgerScan, error) {
	return c.following.LedgerScan(ctx, inquire, screen, eachScreen, arrangeBy)
}

//
func (c *Customer) Ratifiers(
	ctx context.Context,
	level *int64,
	screenPointer, eachScreenPointer *int,
) (*ctypes.OutcomeRatifiers, error) {
	//
	//
	l, err := c.modifyRapidCustomerIfRequiredTo(ctx, level)
	if err != nil {
		return nil, err
	}

	sumNumber := len(l.RatifierAssign.Ratifiers)
	eachScreen := certifyEachScreen(eachScreenPointer)
	screen, err := certifyScreen(screenPointer, eachScreen, sumNumber)
	if err != nil {
		return nil, err
	}

	omitNumber := certifyOmitNumber(screen, eachScreen)
	v := l.RatifierAssign.Ratifiers[omitNumber : omitNumber+cometmath.MinimumInteger(eachScreen, sumNumber-omitNumber)]

	return &ctypes.OutcomeRatifiers{
		LedgerLevel: l.Level,
		Ratifiers:  v,
		Number:       len(v),
		Sum:       sumNumber,
	}, nil
}

func (c *Customer) MulticastProof(ctx context.Context, ev kinds.Proof) (*ctypes.OutcomeMulticastProof, error) {
	return c.following.MulticastProof(ctx, ev)
}

func (c *Customer) Enrol(ctx context.Context, enrollee, inquire string,
	outVolume ...int,
) (out <-chan ctypes.OutcomeEvent, err error) {
	return c.following.Enrol(ctx, enrollee, inquire, outVolume...)
}

func (c *Customer) Deenroll(ctx context.Context, enrollee, inquire string) error {
	return c.following.Deenroll(ctx, enrollee, inquire)
}

func (c *Customer) DeenrollAll(ctx context.Context, enrollee string) error {
	return c.following.DeenrollAll(ctx, enrollee)
}

func (c *Customer) modifyRapidCustomerIfRequiredTo(ctx context.Context, level *int64) (*kinds.RapidLedger, error) {
	var (
		l   *kinds.RapidLedger
		err error
	)
	if level == nil {
		l, err = c.lc.Modify(ctx, time.Now())
	} else {
		l, err = c.lc.ValidateRapidLedgerAtLevel(ctx, *level, time.Now())
	}
	if err != nil {
		return nil, fmt.Errorf("REDACTED", *level, err)
	}
	return l, nil
}

func (c *Customer) EnrollActParser(typ string, dec merkle.ActParser) {
	c.prt.EnrollActParser(typ, dec)
}

//
//
//
func (c *Customer) EnrolWS(ctx *rpctypes.Context, inquire string) (*ctypes.OutcomeEnrol, error) {
	out, err := c.following.Enrol(context.Background(), ctx.DistantAddress(), inquire)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			select {
			case outcomeEvent := <-out:
				//
				//
				ctx.WSLink.AttemptRecordRPCReply(
					rpctypes.NewRPCSuccessReply(
						rpctypes.JsonrpcStringUID(fmt.Sprintf("REDACTED", ctx.JSONRequest.ID)),
						outcomeEvent,
					))
			case <-c.Exit():
				return
			}
		}
	}()

	return &ctypes.OutcomeEnrol{}, nil
}

//
//
func (c *Customer) DeenrollWS(ctx *rpctypes.Context, inquire string) (*ctypes.OutcomeDeenroll, error) {
	err := c.following.Deenroll(context.Background(), ctx.DistantAddress(), inquire)
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeDeenroll{}, nil
}

//
//
func (c *Customer) DeenrollAllWS(ctx *rpctypes.Context) (*ctypes.OutcomeDeenroll, error) {
	err := c.following.DeenrollAll(context.Background(), ctx.DistantAddress())
	if err != nil {
		return nil, err
	}
	return &ctypes.OutcomeDeenroll{}, nil
}

//
const (
	//
	standardEachScreen = 30
	maximumEachScreen     = 100
)

func certifyScreen(screenPointer *int, eachScreen, sumNumber int) (int, error) {
	if eachScreen < 1 {
		panic(fmt.Sprintf("REDACTED", eachScreen))
	}

	if screenPointer == nil { //
		return 1, nil
	}

	sections := ((sumNumber - 1) / eachScreen) + 1
	if sections == 0 {
		sections = 1 //
	}
	screen := *screenPointer
	if screen <= 0 || screen > sections {
		return 1, fmt.Errorf("REDACTED", sections, screen)
	}

	return screen, nil
}

func certifyEachScreen(eachScreenPointer *int) int {
	if eachScreenPointer == nil { //
		return standardEachScreen
	}

	eachScreen := *eachScreenPointer
	if eachScreen < 1 {
		return standardEachScreen
	} else if eachScreen > maximumEachScreen {
		return maximumEachScreen
	}
	return eachScreen
}

func certifyOmitNumber(screen, eachScreen int) int {
	omitNumber := (screen - 1) * eachScreen
	if omitNumber < 0 {
		return 0
	}

	return omitNumber
}
