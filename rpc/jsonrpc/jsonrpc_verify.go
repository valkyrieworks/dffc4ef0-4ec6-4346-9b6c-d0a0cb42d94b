package jsonrpc

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/go-kit/log/term"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cometbytes "github.com/valkyrieworks/utils/octets"
	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"

	customer "github.com/valkyrieworks/rpc/jsonrpc/customer"
	host "github.com/valkyrieworks/rpc/jsonrpc/host"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
const (
	tcpAddress = "REDACTED"

	unixSocket = "REDACTED"
	unixAddress   = "REDACTED" + unixSocket

	webchannelTerminus = "REDACTED"

	verifyValue = "REDACTED"
)

var ctx = context.Background()

type OutcomeReverberate struct {
	Item string `json:"item"`
}

type OutcomeReverberateInteger struct {
	Item int `json:"item"`
}

type OutcomeReverberateOctets struct {
	Item []byte `json:"item"`
}

type OutcomeReverberateDataOctets struct {
	Item cometbytes.HexOctets `json:"item"`
}

type OutcomeReverberateWithStandard struct {
	Item int `json:"item"`
}

//
var Paths = map[string]*host.RPCFunction{
	"REDACTED":            host.NewRPCFunction(ReverberateOutcome, "REDACTED"),
	"REDACTED":         host.NewWsrpcFunction(ReverberateWSOutcome, "REDACTED"),
	"REDACTED":      host.NewRPCFunction(ReverberateOctetsOutcome, "REDACTED"),
	"REDACTED": host.NewRPCFunction(ReverberateDataOctetsOutcome, "REDACTED"),
	"REDACTED":        host.NewRPCFunction(ReverberateIntegerOutcome, "REDACTED"),
	"REDACTED":    host.NewRPCFunction(ReverberateWithStandard, "REDACTED", host.Storable("REDACTED")),
}

func ReverberateOutcome(_ *kinds.Context, v string) (*OutcomeReverberate, error) {
	return &OutcomeReverberate{v}, nil
}

func ReverberateWSOutcome(_ *kinds.Context, v string) (*OutcomeReverberate, error) {
	return &OutcomeReverberate{v}, nil
}

func ReverberateIntegerOutcome(_ *kinds.Context, v int) (*OutcomeReverberateInteger, error) {
	return &OutcomeReverberateInteger{v}, nil
}

func ReverberateOctetsOutcome(_ *kinds.Context, v []byte) (*OutcomeReverberateOctets, error) {
	return &OutcomeReverberateOctets{v}, nil
}

func ReverberateDataOctetsOutcome(_ *kinds.Context, v cometbytes.HexOctets) (*OutcomeReverberateDataOctets, error) {
	return &OutcomeReverberateDataOctets{v}, nil
}

func ReverberateWithStandard(_ *kinds.Context, v *int) (*OutcomeReverberateWithStandard, error) {
	val := -1
	if v != nil {
		val = *v
	}
	return &OutcomeReverberateWithStandard{val}, nil
}

func VerifyMain(m *testing.M) {
	configure()
	code := m.Run()
	os.Exit(code)
}

var hueFn = func(keyvalues ...any) term.FgBgColor {
	for i := 0; i < len(keyvalues)-1; i += 2 {
		if keyvalues[i] == "REDACTED" {
			switch keyvalues[i+1] {
			case "REDACTED":
				return term.FgBgColor{Fg: term.DarkBlue}
			case "REDACTED":
				return term.FgBgColor{Fg: term.DarkCyan}
			}
		}
	}
	return term.FgBgColor{}
}

//
func configure() {
	tracer := log.NewTMTracerWithHueFn(log.NewAlignRecorder(os.Stdout), hueFn)

	cmd := exec.Command("REDACTED", "REDACTED", unixSocket)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	if err = cmd.Wait(); err != nil {
		panic(err)
	}

	tcpTracer := tracer.With("REDACTED", "REDACTED")
	mux := http.NewServeMux()
	host.EnrollRPCRoutines(mux, Paths, tcpTracer)
	wm := host.NewWebchannelOverseer(Paths, host.ReadWait(5*time.Second), host.PingDuration(1*time.Second))
	wm.AssignTracer(tcpTracer)
	mux.HandleFunc(webchannelTerminus, wm.WebchannelManager)
	settings := host.StandardSettings()
	observer1, err := host.Observe(tcpAddress, settings.MaximumAccessLinks)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := host.Attend(observer1, mux, tcpTracer, settings); err != nil {
			panic(err)
		}
	}()

	unixTracer := tracer.With("REDACTED", "REDACTED")
	multiplexer2 := http.NewServeMux()
	host.EnrollRPCRoutines(multiplexer2, Paths, unixTracer)
	wm = host.NewWebchannelOverseer(Paths)
	wm.AssignTracer(unixTracer)
	multiplexer2.HandleFunc(webchannelTerminus, wm.WebchannelManager)
	observer2, err := host.Observe(unixAddress, settings.MaximumAccessLinks)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := host.Attend(observer2, multiplexer2, unixTracer, settings); err != nil {
			panic(err)
		}
	}()

	//
	time.Sleep(time.Second * 2)
}

func reverberateThroughHTTP(cl customer.Invoker, val string) (string, error) {
	options := map[string]any{
		"REDACTED": val,
	}
	outcome := new(OutcomeReverberate)
	if _, err := cl.Invoke(ctx, "REDACTED", options, outcome); err != nil {
		return "REDACTED", err
	}
	return outcome.Item, nil
}

func reverberateIntegerThroughHTTP(cl customer.Invoker, val int) (int, error) {
	options := map[string]any{
		"REDACTED": val,
	}
	outcome := new(OutcomeReverberateInteger)
	if _, err := cl.Invoke(ctx, "REDACTED", options, outcome); err != nil {
		return 0, err
	}
	return outcome.Item, nil
}

func reverberateOctetsThroughHTTP(cl customer.Invoker, octets []byte) ([]byte, error) {
	options := map[string]any{
		"REDACTED": octets,
	}
	outcome := new(OutcomeReverberateOctets)
	if _, err := cl.Invoke(ctx, "REDACTED", options, outcome); err != nil {
		return []byte{}, err
	}
	return outcome.Item, nil
}

func reverberateDataOctetsThroughHTTP(cl customer.Invoker, octets cometbytes.HexOctets) (cometbytes.HexOctets, error) {
	options := map[string]any{
		"REDACTED": octets,
	}
	outcome := new(OutcomeReverberateDataOctets)
	if _, err := cl.Invoke(ctx, "REDACTED", options, outcome); err != nil {
		return []byte{}, err
	}
	return outcome.Item, nil
}

func reverberateWithStandardThroughHTTP(cl customer.Invoker, v *int) (int, error) {
	options := map[string]any{}
	if v != nil {
		options["REDACTED"] = *v
	}
	outcome := new(OutcomeReverberateWithStandard)
	if _, err := cl.Invoke(ctx, "REDACTED", options, outcome); err != nil {
		return 0, err
	}
	return outcome.Item, nil
}

func verifyWithHTTPCustomer(t *testing.T, cl customer.HTTPCustomer) {
	val := verifyValue
	got, err := reverberateThroughHTTP(cl, val)
	require.NoError(t, err)
	assert.Equal(t, got, val)

	value2 := randomOctets(t)
	obtained2, err := reverberateOctetsThroughHTTP(cl, value2)
	require.NoError(t, err)
	assert.Equal(t, obtained2, value2)

	value3 := cometbytes.HexOctets(randomOctets(t))
	obtained3, err := reverberateDataOctetsThroughHTTP(cl, value3)
	require.NoError(t, err)
	assert.Equal(t, obtained3, value3)

	value4 := engineseed.Intn(10000)
	obtained4, err := reverberateIntegerThroughHTTP(cl, value4)
	require.NoError(t, err)
	assert.Equal(t, obtained4, value4)

	obtained5, err := reverberateWithStandardThroughHTTP(cl, nil)
	require.NoError(t, err)
	assert.Equal(t, obtained5, -1)

	value6 := engineseed.Intn(10000)
	obtained6, err := reverberateWithStandardThroughHTTP(cl, &value6)
	require.NoError(t, err)
	assert.Equal(t, obtained6, value6)
}

func reverberateThroughWS(cl *customer.WSCustomer, val string) (string, error) {
	options := map[string]any{
		"REDACTED": val,
	}
	err := cl.Invoke(context.Background(), "REDACTED", options)
	if err != nil {
		return "REDACTED", err
	}

	msg := <-cl.RepliesChan
	if msg.Fault != nil {
		return "REDACTED", err
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	if err != nil {
		return "REDACTED", nil
	}
	return outcome.Item, nil
}

func reverberateOctetsThroughWS(cl *customer.WSCustomer, octets []byte) ([]byte, error) {
	options := map[string]any{
		"REDACTED": octets,
	}
	err := cl.Invoke(context.Background(), "REDACTED", options)
	if err != nil {
		return []byte{}, err
	}

	msg := <-cl.RepliesChan
	if msg.Fault != nil {
		return []byte{}, msg.Fault
	}
	outcome := new(OutcomeReverberateOctets)
	err = json.Unmarshal(msg.Outcome, outcome)
	if err != nil {
		return []byte{}, nil
	}
	return outcome.Item, nil
}

func verifyWithWSCustomer(t *testing.T, cl *customer.WSCustomer) {
	val := verifyValue
	got, err := reverberateThroughWS(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)

	value2 := randomOctets(t)
	obtained2, err := reverberateOctetsThroughWS(cl, value2)
	require.Nil(t, err)
	assert.Equal(t, obtained2, value2)
}

//

func VerifyHostsAndAgentsSimple(t *testing.T) {
	hostLocations := [...]string{tcpAddress, unixAddress}
	for _, address := range hostLocations {
		cl1, err := customer.NewLocator(address)
		require.Nil(t, err)
		fmt.Printf("REDACTED", address)
		verifyWithHTTPCustomer(t, cl1)

		cl2, err := customer.New(address)
		require.Nil(t, err)
		fmt.Printf("REDACTED", address)
		verifyWithHTTPCustomer(t, cl2)

		cl3, err := customer.NewWS(address, webchannelTerminus)
		require.Nil(t, err)
		cl3.AssignTracer(log.VerifyingTracer())
		err = cl3.Begin()
		require.Nil(t, err)
		fmt.Printf("REDACTED", address)
		verifyWithWSCustomer(t, cl3)
		err = cl3.Halt()
		require.NoError(t, err)
	}
}

func VerifyHexStringArgument(t *testing.T) {
	cl, err := customer.NewLocator(tcpAddress)
	require.Nil(t, err)
	//
	val := "REDACTED"
	got, err := reverberateThroughHTTP(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)
}

func VerifyCitedStringArgument(t *testing.T) {
	cl, err := customer.NewLocator(tcpAddress)
	require.Nil(t, err)
	//
	val := "REDACTED"
	got, err := reverberateThroughHTTP(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)
}

func VerifyWSNewWsrpcFunction(t *testing.T) {
	cl, err := customer.NewWS(tcpAddress, webchannelTerminus)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Begin()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	val := verifyValue
	options := map[string]any{
		"REDACTED": val,
	}
	err = cl.Invoke(context.Background(), "REDACTED", options)
	require.Nil(t, err)

	msg := <-cl.RepliesChan
	if msg.Fault != nil {
		t.Fatal(err)
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	require.Nil(t, err)
	got := outcome.Item
	assert.Equal(t, got, val)
}

func VerifyWSManagersListOptions(t *testing.T) {
	cl, err := customer.NewWS(tcpAddress, webchannelTerminus)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Begin()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	val := verifyValue
	options := []any{val}
	err = cl.InvokeWithListOptions(context.Background(), "REDACTED", options)
	require.Nil(t, err)

	msg := <-cl.RepliesChan
	if msg.Fault != nil {
		t.Fatalf("REDACTED", err)
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	require.Nil(t, err)
	got := outcome.Item
	assert.Equal(t, got, val)
}

//
//
func VerifyWSCustomerPingPong(t *testing.T) {
	cl, err := customer.NewWS(tcpAddress, webchannelTerminus)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Begin()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	time.Sleep(6 * time.Second)
}

func VerifyJsonrpcStoring(t *testing.T) {
	httpAddress := strings.Replace(tcpAddress, "REDACTED", "REDACTED", 1)
	cl, err := customer.StandardHTTPCustomer(httpAddress)
	require.NoError(t, err)

	//
	options := make(map[string]any)
	req, err := kinds.IndexToQuery(kinds.JsonrpcIntegerUID(1000), "REDACTED", options)
	require.NoError(t, err)

	output1, err := crudeJsonrpcQuery(t, cl, httpAddress, req)
	defer func() { _ = output1.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", output1.Header.Get("REDACTED"))

	//
	options["REDACTED"] = engineseed.Intn(10000)
	req, err = kinds.IndexToQuery(kinds.JsonrpcIntegerUID(1001), "REDACTED", options)
	require.NoError(t, err)

	output2, err := crudeJsonrpcQuery(t, cl, httpAddress, req)
	defer func() { _ = output2.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", output2.Header.Get("REDACTED"))
}

func crudeJsonrpcQuery(t *testing.T, cl *http.Client, url string, req any) (*http.Response, error) {
	requestOctets, err := json.Marshal(req)
	require.NoError(t, err)

	requestImage := bytes.NewBuffer(requestOctets)
	httpRequest, err := http.NewRequest(http.MethodPost, url, requestImage)
	require.NoError(t, err)

	httpRequest.Header.Set("REDACTED", "REDACTED")

	return cl.Do(httpRequest)
}

func VerifyLocatorStoring(t *testing.T) {
	httpAddress := strings.Replace(tcpAddress, "REDACTED", "REDACTED", 1)
	cl, err := customer.StandardHTTPCustomer(httpAddress)
	require.NoError(t, err)

	//
	args := url.Values{}
	output1, err := crudeLocatorQuery(t, cl, httpAddress+"REDACTED", args)
	defer func() { _ = output1.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", output1.Header.Get("REDACTED"))

	//
	args.Set("REDACTED", fmt.Sprintf("REDACTED", engineseed.Intn(10000)))
	output2, err := crudeLocatorQuery(t, cl, httpAddress+"REDACTED", args)
	defer func() { _ = output2.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", output2.Header.Get("REDACTED"))
}

func crudeLocatorQuery(t *testing.T, cl *http.Client, url string, args url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(args.Encode()))
	require.NoError(t, err)

	req.Header.Set("REDACTED", "REDACTED")

	return cl.Do(req)
}

func randomOctets(t *testing.T) []byte {
	n := engineseed.Intn(10) + 2
	buf := make([]byte, n)
	_, err := crand.Read(buf)
	require.Nil(t, err)
	return bytes.ReplaceAll(buf, []byte("REDACTED"), []byte{100})
}
