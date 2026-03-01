package jsoniface

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

	tendermintoctets "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/octets"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"

	customer "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/customer"
	node "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/node"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
const (
	tcpsocketLocation = "REDACTED"

	posixPort = "REDACTED"
	posixLocation   = "REDACTED" + posixPort

	webterminalGateway = "REDACTED"

	verifyItem = "REDACTED"
)

var ctx = context.Background()

type OutcomeReverberate struct {
	Datum string `json:"datum"`
}

type OutcomeReverberateInteger struct {
	Datum int `json:"datum"`
}

type OutcomeReverberateOctets struct {
	Datum []byte `json:"datum"`
}

type OutcomeReverberateDataOctets struct {
	Datum tendermintoctets.HexadecimalOctets `json:"datum"`
}

type OutcomeReverberateUsingFallback struct {
	Datum int `json:"datum"`
}

//
var Paths = map[string]*node.RemoteMethod{
	"REDACTED":            node.FreshRemoteMethod(ReverberateOutcome, "REDACTED"),
	"REDACTED":         node.FreshSocketifaceMethod(ReverberateSocketOutcome, "REDACTED"),
	"REDACTED":      node.FreshRemoteMethod(ReverberateOctetsOutcome, "REDACTED"),
	"REDACTED": node.FreshRemoteMethod(ReverberateDataOctetsOutcome, "REDACTED"),
	"REDACTED":        node.FreshRemoteMethod(ReverberateIntegerOutcome, "REDACTED"),
	"REDACTED":    node.FreshRemoteMethod(ReverberateUsingFallback, "REDACTED", node.Storable("REDACTED")),
}

func ReverberateOutcome(_ *kinds.Env, v string) (*OutcomeReverberate, error) {
	return &OutcomeReverberate{v}, nil
}

func ReverberateSocketOutcome(_ *kinds.Env, v string) (*OutcomeReverberate, error) {
	return &OutcomeReverberate{v}, nil
}

func ReverberateIntegerOutcome(_ *kinds.Env, v int) (*OutcomeReverberateInteger, error) {
	return &OutcomeReverberateInteger{v}, nil
}

func ReverberateOctetsOutcome(_ *kinds.Env, v []byte) (*OutcomeReverberateOctets, error) {
	return &OutcomeReverberateOctets{v}, nil
}

func ReverberateDataOctetsOutcome(_ *kinds.Env, v tendermintoctets.HexadecimalOctets) (*OutcomeReverberateDataOctets, error) {
	return &OutcomeReverberateDataOctets{v}, nil
}

func ReverberateUsingFallback(_ *kinds.Env, v *int) (*OutcomeReverberateUsingFallback, error) {
	val := -1
	if v != nil {
		val = *v
	}
	return &OutcomeReverberateUsingFallback{val}, nil
}

func VerifyPrimary(m *testing.M) {
	configure()
	cipher := m.Run()
	os.Exit(cipher)
}

var hueProc = func(tokvals ...any) term.FgBgColor {
	for i := 0; i < len(tokvals)-1; i += 2 {
		if tokvals[i] == "REDACTED" {
			switch tokvals[i+1] {
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
	tracer := log.FreshTEMPTracerUsingHueProc(log.FreshChronizePersistor(os.Stdout), hueProc)

	cmd := exec.Command("REDACTED", "REDACTED", posixPort)
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	if err = cmd.Wait(); err != nil {
		panic(err)
	}

	tcpsocketTracer := tracer.Using("REDACTED", "REDACTED")
	mux := http.NewServeMux()
	node.EnrollRemoteRoutines(mux, Paths, tcpsocketTracer)
	wm := node.FreshWebterminalAdministrator(Paths, node.ScanPause(5*time.Second), node.PingSpan(1*time.Second))
	wm.AssignTracer(tcpsocketTracer)
	mux.HandleFunc(webterminalGateway, wm.WebterminalProcessor)
	settings := node.FallbackSettings()
	observer1, err := node.Overhear(tcpsocketLocation, settings.MaximumInitiateLinks)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := node.Attend(observer1, mux, tcpsocketTracer, settings); err != nil {
			panic(err)
		}
	}()

	posixTracer := tracer.Using("REDACTED", "REDACTED")
	multiplexer2 := http.NewServeMux()
	node.EnrollRemoteRoutines(multiplexer2, Paths, posixTracer)
	wm = node.FreshWebterminalAdministrator(Paths)
	wm.AssignTracer(posixTracer)
	multiplexer2.HandleFunc(webterminalGateway, wm.WebterminalProcessor)
	observer2, err := node.Overhear(posixLocation, settings.MaximumInitiateLinks)
	if err != nil {
		panic(err)
	}
	go func() {
		if err := node.Attend(observer2, multiplexer2, posixTracer, settings); err != nil {
			panic(err)
		}
	}()

	//
	time.Sleep(time.Second * 2)
}

func reverberateThroughHttpsvc(cl customer.Invoker, val string) (string, error) {
	parameters := map[string]any{
		"REDACTED": val,
	}
	outcome := new(OutcomeReverberate)
	if _, err := cl.Invocation(ctx, "REDACTED", parameters, outcome); err != nil {
		return "REDACTED", err
	}
	return outcome.Datum, nil
}

func reverberateIntegerThroughHttpsvc(cl customer.Invoker, val int) (int, error) {
	parameters := map[string]any{
		"REDACTED": val,
	}
	outcome := new(OutcomeReverberateInteger)
	if _, err := cl.Invocation(ctx, "REDACTED", parameters, outcome); err != nil {
		return 0, err
	}
	return outcome.Datum, nil
}

func reverberateOctetsThroughHttpsvc(cl customer.Invoker, octets []byte) ([]byte, error) {
	parameters := map[string]any{
		"REDACTED": octets,
	}
	outcome := new(OutcomeReverberateOctets)
	if _, err := cl.Invocation(ctx, "REDACTED", parameters, outcome); err != nil {
		return []byte{}, err
	}
	return outcome.Datum, nil
}

func reverberateDataOctetsThroughHttpsvc(cl customer.Invoker, octets tendermintoctets.HexadecimalOctets) (tendermintoctets.HexadecimalOctets, error) {
	parameters := map[string]any{
		"REDACTED": octets,
	}
	outcome := new(OutcomeReverberateDataOctets)
	if _, err := cl.Invocation(ctx, "REDACTED", parameters, outcome); err != nil {
		return []byte{}, err
	}
	return outcome.Datum, nil
}

func reverberateUsingFallbackThroughHttpsvc(cl customer.Invoker, v *int) (int, error) {
	parameters := map[string]any{}
	if v != nil {
		parameters["REDACTED"] = *v
	}
	outcome := new(OutcomeReverberateUsingFallback)
	if _, err := cl.Invocation(ctx, "REDACTED", parameters, outcome); err != nil {
		return 0, err
	}
	return outcome.Datum, nil
}

func verifyUsingHttpsvcCustomer(t *testing.T, cl customer.HttpsvcCustomer) {
	val := verifyItem
	got, err := reverberateThroughHttpsvc(cl, val)
	require.NoError(t, err)
	assert.Equal(t, got, val)

	valid2 := arbitraryOctets(t)
	taken2, err := reverberateOctetsThroughHttpsvc(cl, valid2)
	require.NoError(t, err)
	assert.Equal(t, taken2, valid2)

	item3 := tendermintoctets.HexadecimalOctets(arbitraryOctets(t))
	taken3, err := reverberateDataOctetsThroughHttpsvc(cl, item3)
	require.NoError(t, err)
	assert.Equal(t, taken3, item3)

	item4 := commitrand.Integern(10000)
	taken4, err := reverberateIntegerThroughHttpsvc(cl, item4)
	require.NoError(t, err)
	assert.Equal(t, taken4, item4)

	taken5, err := reverberateUsingFallbackThroughHttpsvc(cl, nil)
	require.NoError(t, err)
	assert.Equal(t, taken5, -1)

	item6 := commitrand.Integern(10000)
	taken6, err := reverberateUsingFallbackThroughHttpsvc(cl, &item6)
	require.NoError(t, err)
	assert.Equal(t, taken6, item6)
}

func reverberateThroughSocket(cl *customer.SocketCustomer, val string) (string, error) {
	parameters := map[string]any{
		"REDACTED": val,
	}
	err := cl.Invocation(context.Background(), "REDACTED", parameters)
	if err != nil {
		return "REDACTED", err
	}

	msg := <-cl.RepliesChnl
	if msg.Failure != nil {
		return "REDACTED", err
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	if err != nil {
		return "REDACTED", nil
	}
	return outcome.Datum, nil
}

func reverberateOctetsThroughSocket(cl *customer.SocketCustomer, octets []byte) ([]byte, error) {
	parameters := map[string]any{
		"REDACTED": octets,
	}
	err := cl.Invocation(context.Background(), "REDACTED", parameters)
	if err != nil {
		return []byte{}, err
	}

	msg := <-cl.RepliesChnl
	if msg.Failure != nil {
		return []byte{}, msg.Failure
	}
	outcome := new(OutcomeReverberateOctets)
	err = json.Unmarshal(msg.Outcome, outcome)
	if err != nil {
		return []byte{}, nil
	}
	return outcome.Datum, nil
}

func verifyUsingSocketCustomer(t *testing.T, cl *customer.SocketCustomer) {
	val := verifyItem
	got, err := reverberateThroughSocket(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)

	valid2 := arbitraryOctets(t)
	taken2, err := reverberateOctetsThroughSocket(cl, valid2)
	require.Nil(t, err)
	assert.Equal(t, taken2, valid2)
}

//

func VerifyHostsAlsoCustomersFundamental(t *testing.T) {
	daemonLocations := [...]string{tcpsocketLocation, posixLocation}
	for _, location := range daemonLocations {
		cl1, err := customer.FreshURL(location)
		require.Nil(t, err)
		fmt.Printf("REDACTED", location)
		verifyUsingHttpsvcCustomer(t, cl1)

		cl2, err := customer.New(location)
		require.Nil(t, err)
		fmt.Printf("REDACTED", location)
		verifyUsingHttpsvcCustomer(t, cl2)

		cl3, err := customer.FreshSocket(location, webterminalGateway)
		require.Nil(t, err)
		cl3.AssignTracer(log.VerifyingTracer())
		err = cl3.Initiate()
		require.Nil(t, err)
		fmt.Printf("REDACTED", location)
		verifyUsingSocketCustomer(t, cl3)
		err = cl3.Halt()
		require.NoError(t, err)
	}
}

func VerifyHexadecimalTextArgument(t *testing.T) {
	cl, err := customer.FreshURL(tcpsocketLocation)
	require.Nil(t, err)
	//
	val := "REDACTED"
	got, err := reverberateThroughHttpsvc(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)
}

func VerifyCitedTextArgument(t *testing.T) {
	cl, err := customer.FreshURL(tcpsocketLocation)
	require.Nil(t, err)
	//
	val := "REDACTED"
	got, err := reverberateThroughHttpsvc(cl, val)
	require.Nil(t, err)
	assert.Equal(t, got, val)
}

func VerifySocketFreshSocketifaceMethod(t *testing.T) {
	cl, err := customer.FreshSocket(tcpsocketLocation, webterminalGateway)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Initiate()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	val := verifyItem
	parameters := map[string]any{
		"REDACTED": val,
	}
	err = cl.Invocation(context.Background(), "REDACTED", parameters)
	require.Nil(t, err)

	msg := <-cl.RepliesChnl
	if msg.Failure != nil {
		t.Fatal(err)
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	require.Nil(t, err)
	got := outcome.Datum
	assert.Equal(t, got, val)
}

func VerifySocketOverseesSeriesParameters(t *testing.T) {
	cl, err := customer.FreshSocket(tcpsocketLocation, webterminalGateway)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Initiate()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	val := verifyItem
	parameters := []any{val}
	err = cl.InvocationUsingSeriesParameters(context.Background(), "REDACTED", parameters)
	require.Nil(t, err)

	msg := <-cl.RepliesChnl
	if msg.Failure != nil {
		t.Fatalf("REDACTED", err)
	}
	outcome := new(OutcomeReverberate)
	err = json.Unmarshal(msg.Outcome, outcome)
	require.Nil(t, err)
	got := outcome.Datum
	assert.Equal(t, got, val)
}

//
//
func VerifySocketCustomerPingPong(t *testing.T) {
	cl, err := customer.FreshSocket(tcpsocketLocation, webterminalGateway)
	require.Nil(t, err)
	cl.AssignTracer(log.VerifyingTracer())
	err = cl.Initiate()
	require.Nil(t, err)
	t.Cleanup(func() {
		if err := cl.Halt(); err != nil {
			t.Error(err)
		}
	})

	time.Sleep(6 * time.Second)
}

func VerifyJsonifaceStashing(t *testing.T) {
	httpsvcLocation := strings.Replace(tcpsocketLocation, "REDACTED", "REDACTED", 1)
	cl, err := customer.FallbackHttpsvcCustomer(httpsvcLocation)
	require.NoError(t, err)

	//
	parameters := make(map[string]any)
	req, err := kinds.IndexTowardSolicit(kinds.JsonifaceIntegerUUID(1000), "REDACTED", parameters)
	require.NoError(t, err)

	ans1, err := crudeJsonifaceSolicit(t, cl, httpsvcLocation, req)
	defer func() { _ = ans1.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", ans1.Header.Get("REDACTED"))

	//
	parameters["REDACTED"] = commitrand.Integern(10000)
	req, err = kinds.IndexTowardSolicit(kinds.JsonifaceIntegerUUID(1001), "REDACTED", parameters)
	require.NoError(t, err)

	ans2, err := crudeJsonifaceSolicit(t, cl, httpsvcLocation, req)
	defer func() { _ = ans2.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", ans2.Header.Get("REDACTED"))
}

func crudeJsonifaceSolicit(t *testing.T, cl *http.Client, url string, req any) (*http.Response, error) {
	requestOctets, err := json.Marshal(req)
	require.NoError(t, err)

	requestArea := bytes.NewBuffer(requestOctets)
	httpsvcRequest, err := http.NewRequest(http.MethodPost, url, requestArea)
	require.NoError(t, err)

	httpsvcRequest.Header.Set("REDACTED", "REDACTED")

	return cl.Do(httpsvcRequest)
}

func VerifyURLStashing(t *testing.T) {
	httpsvcLocation := strings.Replace(tcpsocketLocation, "REDACTED", "REDACTED", 1)
	cl, err := customer.FallbackHttpsvcCustomer(httpsvcLocation)
	require.NoError(t, err)

	//
	arguments := url.Values{}
	ans1, err := crudeURLSolicit(t, cl, httpsvcLocation+"REDACTED", arguments)
	defer func() { _ = ans1.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", ans1.Header.Get("REDACTED"))

	//
	arguments.Set("REDACTED", fmt.Sprintf("REDACTED", commitrand.Integern(10000)))
	ans2, err := crudeURLSolicit(t, cl, httpsvcLocation+"REDACTED", arguments)
	defer func() { _ = ans2.Body.Close() }()
	require.NoError(t, err)
	assert.Equal(t, "REDACTED", ans2.Header.Get("REDACTED"))
}

func crudeURLSolicit(t *testing.T, cl *http.Client, url string, arguments url.Values) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(arguments.Encode()))
	require.NoError(t, err)

	req.Header.Set("REDACTED", "REDACTED")

	return cl.Do(req)
}

func arbitraryOctets(t *testing.T) []byte {
	n := commitrand.Integern(10) + 2
	buf := make([]byte, n)
	_, err := crand.Read(buf)
	require.Nil(t, err)
	return bytes.ReplaceAll(buf, []byte("REDACTED"), []byte{100})
}
