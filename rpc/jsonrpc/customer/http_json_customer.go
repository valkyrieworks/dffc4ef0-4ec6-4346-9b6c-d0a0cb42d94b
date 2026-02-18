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

	engineconnect "github.com/valkyrieworks/utils/align"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

const (
	schemaHTTP  = "REDACTED"
	schemaHttps = "REDACTED"
	schemaWSS   = "REDACTED"
	schemaWS    = "REDACTED"
	schemaTCP   = "REDACTED"
	schemaUNIX  = "REDACTED"
)

var terminatesWithPortTemplate = regexp.MustCompile("REDACTED")

//

//
type analyzedURL struct {
	url.URL

	isUnixSocket bool
}

//
func newAnalyzedURL(distantAddress string) (*analyzedURL, error) {
	u, err := url.Parse(distantAddress)
	if err != nil {
		return nil, err
	}

	//
	if u.Scheme == "REDACTED" {
		u.Scheme = schemaTCP
	}

	pu := &analyzedURL{
		URL:          *u,
		isUnixSocket: false,
	}

	if u.Scheme == schemaUNIX {
		pu.isUnixSocket = true
	}

	return pu, nil
}

//
func (u *analyzedURL) CollectionStandardProtocolHTTP() {
	//
	switch u.Scheme {
	case schemaHTTP, schemaHttps, schemaWS, schemaWSS:
		//
	default:
		//
		u.Scheme = schemaHTTP
	}
}

//
func (u analyzedURL) FetchMachineWithRoute() string {
	//
	return u.Host + u.EscapedPath()
}

//
func (u analyzedURL) FetchClippedMachineWithRoute() string {
	//
	if !u.isUnixSocket {
		return u.FetchMachineWithRoute()
	}
	//
	//
	//
	return strings.ReplaceAll(u.FetchMachineWithRoute(), "REDACTED", "REDACTED")
}

//
func (u analyzedURL) FetchCallLocation() string {
	//
	if !u.isUnixSocket {
		hasPort := terminatesWithPortTemplate.MatchString(u.Host)
		if !hasPort {
			//
			//
			//
			switch u.Scheme {
			case schemaHTTP, schemaWS:
				return u.Host + "REDACTED"
			case schemaHttps, schemaWSS:
				return u.Host + "REDACTED"
			}
		}
		return u.Host
	}
	//
	return u.FetchMachineWithRoute()
}

//
func (u analyzedURL) FetchClippedURL() string {
	return u.Scheme + "REDACTED" + u.FetchClippedMachineWithRoute()
}

//

//
type HTTPCustomer interface {
	//
	Invoke(ctx context.Context, procedure string, options map[string]any, outcome any) (any, error)
}

//
type Invoker interface {
	Invoke(ctx context.Context, procedure string, options map[string]any, outcome any) (any, error)
}

//

//
//
//
//
type Customer struct {
	location  string
	username string
	secret string

	customer *http.Client

	mtx       engineconnect.Lock
	followingRequestUID int
}

var _ HTTPCustomer = (*Customer)(nil)

//
//
var (
	_ Invoker = (*Customer)(nil)
	_ Invoker = (*QueryGroup)(nil)
)

var _ fmt.Stringer = (*Customer)(nil)

//
//
func New(external string) (*Customer, error) {
	httpCustomer, err := StandardHTTPCustomer(external)
	if err != nil {
		return nil, err
	}
	return NewWithHTTPCustomer(external, httpCustomer)
}

//
//
//
func NewWithHTTPCustomer(external string, customer *http.Client) (*Customer, error) {
	if customer == nil {
		panic("REDACTED")
	}

	analyzedURL, err := newAnalyzedURL(external)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", external, err)
	}

	analyzedURL.CollectionStandardProtocolHTTP()

	location := analyzedURL.FetchClippedURL()
	username := analyzedURL.User.Username()
	secret, _ := analyzedURL.User.Password()

	rpcCustomer := &Customer{
		location:  location,
		username: username,
		secret: secret,
		customer:   customer,
	}

	return rpcCustomer, nil
}

//
//
func (c *Customer) Invoke(
	ctx context.Context,
	procedure string,
	options map[string]any,
	outcome any,
) (any, error) {
	id := c.followingQueryUID()

	query, err := kinds.IndexToQuery(id, procedure, options)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	queryOctets, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	queryImage := bytes.NewBuffer(queryOctets)
	httpQuery, err := http.NewRequestWithContext(ctx, http.MethodPost, c.location, queryImage)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpQuery.Header.Set("REDACTED", "REDACTED")

	if c.username != "REDACTED" || c.secret != "REDACTED" {
		httpQuery.SetBasicAuth(c.username, c.secret)
	}

	httpReply, err := c.customer.Do(httpQuery)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	defer httpReply.Body.Close()

	replyOctets, err := io.ReadAll(httpReply.Body)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", fetchHTTPReplyErrPrefix(httpReply), err)
	}

	res, err := unserializeReplyOctets(replyOctets, id, outcome)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", fetchHTTPReplyErrPrefix(httpReply), err)
	}
	return res, nil
}

func fetchHTTPReplyErrPrefix(reply *http.Response) string {
	return fmt.Sprintf("REDACTED", reply.Status, reply.Proto)
}

func (c *Customer) String() string {
	return fmt.Sprintf("REDACTED", c.username, c.location, c.customer, c.followingRequestUID)
}

//
func (c *Customer) NewQueryGroup() *QueryGroup {
	return &QueryGroup{
		queries: make([]*jsonRPCCachedQuery, 0),
		customer:   c,
	}
}

func (c *Customer) transmitGroup(ctx context.Context, queries []*jsonRPCCachedQuery) ([]any, error) {
	queries := make([]kinds.RPCQuery, 0, len(queries))
	outcomes := make([]any, 0, len(queries))
	for _, req := range queries {
		queries = append(queries, req.query)
		outcomes = append(outcomes, req.outcome)
	}

	//
	queryOctets, err := json.Marshal(queries)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpQuery, err := http.NewRequestWithContext(ctx, http.MethodPost, c.location, bytes.NewBuffer(queryOctets))
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	httpQuery.Header.Set("REDACTED", "REDACTED")

	if c.username != "REDACTED" || c.secret != "REDACTED" {
		httpQuery.SetBasicAuth(c.username, c.secret)
	}

	httpReply, err := c.customer.Do(httpQuery)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	defer httpReply.Body.Close()

	replyOctets, err := io.ReadAll(httpReply.Body)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	//
	ids := make([]kinds.JsonrpcIntegerUID, len(queries))
	for i, req := range queries {
		ids[i] = req.query.ID.(kinds.JsonrpcIntegerUID)
	}

	return unserializeReplyOctetsList(replyOctets, ids, outcomes)
}

func (c *Customer) followingQueryUID() kinds.JsonrpcIntegerUID {
	c.mtx.Lock()
	id := c.followingRequestUID
	c.followingRequestUID++
	c.mtx.Unlock()
	return kinds.JsonrpcIntegerUID(id)
}

//

//
//
type jsonRPCCachedQuery struct {
	query kinds.RPCQuery
	outcome  any //
}

//
//
//
type QueryGroup struct {
	customer *Customer

	mtx      engineconnect.Lock
	queries []*jsonRPCCachedQuery
}

//
func (b *QueryGroup) Number() int {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return len(b.queries)
}

func (b *QueryGroup) queue(req *jsonRPCCachedQuery) {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	b.queries = append(b.queries, req)
}

//
func (b *QueryGroup) Flush() int {
	b.mtx.Lock()
	defer b.mtx.Unlock()
	return b.flush()
}

func (b *QueryGroup) flush() int {
	tally := len(b.queries)
	b.queries = make([]*jsonRPCCachedQuery, 0)
	return tally
}

//
//
//
func (b *QueryGroup) Transmit(ctx context.Context) ([]any, error) {
	b.mtx.Lock()
	defer func() {
		b.flush()
		b.mtx.Unlock()
	}()
	return b.customer.transmitGroup(ctx, b.queries)
}

//
//
func (b *QueryGroup) Invoke(
	_ context.Context,
	procedure string,
	options map[string]any,
	outcome any,
) (any, error) {
	id := b.customer.followingQueryUID()
	query, err := kinds.IndexToQuery(id, procedure, options)
	if err != nil {
		return nil, err
	}
	b.queue(&jsonRPCCachedQuery{query: query, outcome: outcome})
	return outcome, nil
}

//

func createHTTPCaller(distantAddress string) (func(string, string) (net.Conn, error), error) {
	u, err := newAnalyzedURL(distantAddress)
	if err != nil {
		return nil, err
	}

	protocol := u.Scheme

	//
	switch protocol {
	case schemaHTTP, schemaHttps:
		protocol = schemaTCP
	}

	callFn := func(schema, address string) (net.Conn, error) {
		return net.Dial(protocol, u.FetchCallLocation())
	}

	return callFn, nil
}

//
//
//
//
func StandardHTTPCustomer(distantAddress string) (*http.Client, error) {
	callFn, err := createHTTPCaller(distantAddress)
	if err != nil {
		return nil, err
	}

	customer := &http.Client{
		Transport: &http.Transport{
			//
			DisableCompression: true,
			Dial:               callFn,
			Proxy:              http.ProxyFromEnvironment,
		},
	}

	return customer, nil
}
