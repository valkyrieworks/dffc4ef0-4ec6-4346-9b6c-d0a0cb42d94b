package customer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

const (
	schemaHttpsvc  = "REDACTED"
	schemaTransportsecsvc = "REDACTED"
	schemaTransportsecsocket   = "REDACTED"
	schemaSocket    = "REDACTED"
	schemaTcpsocket   = "REDACTED"
	schemaPosix  = "REDACTED"
)

var terminatesUsingChannelTemplate = regexp.MustCompile("REDACTED")

//

//
type processedWebroute struct {
	url.URL

	equalsPosixPort bool
}

//
func freshProcessedWebroute(distantLocation string) (*processedWebroute, error) {
	u, err := url.Parse(distantLocation)
	if err != nil {
		return nil, err
	}

	//
	if u.Scheme == "REDACTED" {
		u.Scheme = schemaTcpsocket
	}

	pu := &processedWebroute{
		URL:          *u,
		equalsPosixPort: false,
	}

	if u.Scheme == schemaPosix {
		pu.equalsPosixPort = true
	}

	return pu, nil
}

//
func (u *processedWebroute) AssignFallbackProtocolHttpsvc() {
	//
	switch u.Scheme {
	case schemaHttpsvc, schemaTransportsecsvc, schemaSocket, schemaTransportsecsocket:
		//
	default:
		//
		u.Scheme = schemaHttpsvc
	}
}

//
func (u processedWebroute) ObtainMachineUsingRoute() string {
	//
	return u.Host + u.EscapedPath()
}

//
func (u processedWebroute) ObtainShortenedMachineUsingRoute() string {
	//
	if !u.equalsPosixPort {
		return u.ObtainMachineUsingRoute()
	}
	//
	//
	//
	return strings.ReplaceAll(u.ObtainMachineUsingRoute(), "REDACTED", "REDACTED")
}

//
func (u processedWebroute) ObtainCallLocator() string {
	//
	if !u.equalsPosixPort {
		ownsChannel := terminatesUsingChannelTemplate.MatchString(u.Host)
		if !ownsChannel {
			//
			//
			//
			switch u.Scheme {
			case schemaHttpsvc, schemaSocket:
				return u.Host + "REDACTED"
			case schemaTransportsecsvc, schemaTransportsecsocket:
				return u.Host + "REDACTED"
			}
		}
		return u.Host
	}
	//
	return u.ObtainMachineUsingRoute()
}

//
func (u processedWebroute) ObtainShortenedWebroute() string {
	return u.Scheme + "REDACTED" + u.ObtainShortenedMachineUsingRoute()
}

//

//
type HttpsvcCustomer interface {
	//
	Invocation(ctx context.Context, procedure string, parameters map[string]any, outcome any) (any, error)
}

//
type Invoker interface {
	Invocation(ctx context.Context, procedure string, parameters map[string]any, outcome any) (any, error)
}

//

//
//
//
//
type Customer struct {
	location  string
	loginname string
	secret string

	customer *http.Client

	mtx       commitchronize.Exclusion
	followingRequestUUID int
}

var _ HttpsvcCustomer = (*Customer)(nil)

//
//
var (
	_ Invoker = (*Customer)(nil)
	_ Invoker = (*SolicitCluster)(nil)
)

var _ fmt.Stringer = (*Customer)(nil)

//
//
func New(distant string) (*Customer, error) {
	httpsvcCustomer, err := FallbackHttpsvcCustomer(distant)
	if err != nil {
		return nil, err
	}
	return FreshUsingHttpsvcCustomer(distant, httpsvcCustomer)
}

//
//
//
func FreshUsingHttpsvcCustomer(distant string, customer *http.Client) (*Customer, error) {
	if customer == nil {
		panic("REDACTED")
	}

	processedWebroute, err := freshProcessedWebroute(distant)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", distant, err)
	}

	processedWebroute.AssignFallbackProtocolHttpsvc()

	location := processedWebroute.ObtainShortenedWebroute()
	loginname := processedWebroute.User.Username()
	secret, _ := processedWebroute.User.Password()

	ifaceCustomer := &Customer{
		location:  location,
		loginname: loginname,
		secret: secret,
		customer:   customer,
	}

	return ifaceCustomer, nil
}

//
//
func (c *Customer) Invocation(
	ctx context.Context,
	procedure string,
	parameters map[string]any,
	outcome any,
) (any, error) {
	id := c.followingSolicitUUID()

	solicit, err := kinds.IndexTowardSolicit(id, procedure, parameters)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	solicitOctets, err := json.Marshal(solicit)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	solicitArea := bytes.NewBuffer(solicitOctets)
	httpsvcSolicit, err := http.NewRequestWithContext(ctx, http.MethodPost, c.location, solicitArea)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpsvcSolicit.Header.Set("REDACTED", "REDACTED")

	if c.loginname != "REDACTED" || c.secret != "REDACTED" {
		httpsvcSolicit.SetBasicAuth(c.loginname, c.secret)
	}

	httpsvcReply, err := c.customer.Do(httpsvcSolicit)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	defer httpsvcReply.Body.Close()

	replyOctets, err := io.ReadAll(httpsvcReply.Body)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", obtainHttpsvcAnswerFaultHeading(httpsvcReply), err)
	}

	res, err := decodeReplyOctets(replyOctets, id, outcome)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", obtainHttpsvcAnswerFaultHeading(httpsvcReply), err)
	}
	return res, nil
}

func obtainHttpsvcAnswerFaultHeading(reply *http.Response) string {
	return fmt.Sprintf("REDACTED", reply.Status, reply.Proto)
}

func (c *Customer) Text() string {
	return fmt.Sprintf("REDACTED", c.loginname, c.location, c.customer, c.followingRequestUUID)
}

//
func (c *Customer) FreshSolicitCluster() *SolicitCluster {
	return &SolicitCluster{
		solicits: make([]*jsnRemoteCachedSolicit, 0),
		customer:   c,
	}
}

func (c *Customer) transmitCluster(ctx context.Context, solicits []*jsnRemoteCachedSolicit) ([]any, error) {
	solicitations := make([]kinds.RemoteSolicit, 0, len(solicits))
	outcomes := make([]any, 0, len(solicits))
	for _, req := range solicits {
		solicitations = append(solicitations, req.solicit)
		outcomes = append(outcomes, req.outcome)
	}

	//
	solicitOctets, err := json.Marshal(solicitations)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpsvcSolicit, err := http.NewRequestWithContext(ctx, http.MethodPost, c.location, bytes.NewBuffer(solicitOctets))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpsvcSolicit.Header.Set("REDACTED", "REDACTED")

	if c.loginname != "REDACTED" || c.secret != "REDACTED" {
		httpsvcSolicit.SetBasicAuth(c.loginname, c.secret)
	}

	httpsvcReply, err := c.customer.Do(httpsvcSolicit)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	defer httpsvcReply.Body.Close()

	replyOctets, err := io.ReadAll(httpsvcReply.Body)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	ids := make([]kinds.JsonifaceIntegerUUID, len(solicits))
	for i, req := range solicits {
		ids[i] = req.solicit.ID.(kinds.JsonifaceIntegerUUID)
	}

	return decodeReplyOctetsSeries(replyOctets, ids, outcomes)
}

func (c *Customer) followingSolicitUUID() kinds.JsonifaceIntegerUUID {
	c.mtx.Lock()
	id := c.followingRequestUUID
	c.followingRequestUUID++
	c.mtx.Unlock()
	return kinds.JsonifaceIntegerUUID(id)
}

//

//
//
type jsnRemoteCachedSolicit struct {
	solicit kinds.RemoteSolicit
	outcome  any //
}

//
//
//
type SolicitCluster struct {
	customer *Customer

	mtx      commitchronize.Exclusion
	solicits []*jsnRemoteCachedSolicit
}

//
func (b *SolicitCluster) Tally() int {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return len(b.solicits)
}

func (b *SolicitCluster) queue(req *jsnRemoteCachedSolicit) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.solicits = append(b.solicits, req)
}

//
func (b *SolicitCluster) Flush() int {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.flush()
}

func (b *SolicitCluster) flush() int {
	tally := len(b.solicits)
	b.solicits = make([]*jsnRemoteCachedSolicit, 0)
	return tally
}

//
//
//
func (b *SolicitCluster) Transmit(ctx context.Context) ([]any, error) {
	b.mtx.Lock()
	defer func() {
		b.flush()
		b.mtx.Unlock()
	}()
	return b.customer.transmitCluster(ctx, b.solicits)
}

//
//
func (b *SolicitCluster) Invocation(
	_ context.Context,
	procedure string,
	parameters map[string]any,
	outcome any,
) (any, error) {
	id := b.customer.followingSolicitUUID()
	solicit, err := kinds.IndexTowardSolicit(id, procedure, parameters)
	if err != nil {
		return nil, err
	}
	b.queue(&jsnRemoteCachedSolicit{solicit: solicit, outcome: outcome})
	return outcome, nil
}

//

func createHttpsvcCaller(distantLocation string) (func(string, string) (net.Conn, error), error) {
	u, err := freshProcessedWebroute(distantLocation)
	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	//
	switch scheme {
	case schemaHttpsvc, schemaTransportsecsvc:
		scheme = schemaTcpsocket
	}

	callProc := func(schema, location string) (net.Conn, error) {
		return net.Dial(scheme, u.ObtainCallLocator())
	}

	return callProc, nil
}

//
//
//
//
func FallbackHttpsvcCustomer(distantLocation string) (*http.Client, error) {
	callProc, err := createHttpsvcCaller(distantLocation)
	if err != nil {
		return nil, err
	}

	customer := &http.Client{
		Transport: &http.Transport{
			//
			DisableCompression: true,
			Dial:               callProc,
			Proxy:              http.ProxyFromEnvironment,
		},
	}

	return customer, nil
}
