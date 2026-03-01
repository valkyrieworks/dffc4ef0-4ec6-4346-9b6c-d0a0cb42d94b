package customer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	metrics "github.com/rcrowley/go-metrics"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	commitrand "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/arbitrary"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/facility"
	commitchronize "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/chronize"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

const (
	fallbackMaximumReestablishEndeavors = 25
	fallbackRecordPause            = 0
	fallbackFetchPause             = 0
	fallbackPingSpan           = 0
)

//
//
//
//
type SocketCustomer struct {
	link *websocket.Conn

	Location  string //
	Gateway string //
	Loginname string
	Secret string

	Caller func(string, string) (net.Conn, error)

	//
	//
	RepliesChnl chan kinds.RemoteReply

	//
	uponReestablish func()

	//
	transmit            chan kinds.RemoteSolicit //
	pendinglist         chan kinds.RemoteSolicit //
	reestablishSubsequent  chan error            //
	fetchProcedureExit chan struct{}         //

	//
	maximumReestablishEndeavors int

	//
	scheme string

	wg sync.WaitGroup

	mtx            commitchronize.ReadwriteExclusion
	relayedFinalPingLocated time.Time
	reestablishing   bool
	followingRequestUUID      int
	//

	//
	recordPause time.Duration

	//
	fetchPause time.Duration

	//
	pingSpan time.Duration

	facility.FoundationFacility

	//
	//
	PingPongWaitstateClock metrics.Timer
}

//
//
//
//
func FreshSocket(distantLocation, gateway string, choices ...func(*SocketCustomer)) (*SocketCustomer, error) {
	processedWebroute, err := freshProcessedWebroute(distantLocation)
	if err != nil {
		return nil, err
	}
	//
	if processedWebroute.Scheme == schemaTransportsecsvc {
		processedWebroute.Scheme = schemaTransportsecsocket
	} else if processedWebroute.Scheme != schemaTransportsecsocket {
		processedWebroute.Scheme = schemaSocket
	}

	//
	loginname := "REDACTED"
	secret := "REDACTED"
	if processedWebroute.User.String() != "REDACTED" {
		loginname = processedWebroute.User.Username()
		secret, _ = processedWebroute.User.Password()
	}

	callProc, err := createHttpsvcCaller(distantLocation)
	if err != nil {
		return nil, err
	}

	c := &SocketCustomer{
		Location:  processedWebroute.ObtainShortenedMachineUsingRoute(),
		Loginname: loginname,
		Secret: secret,
		Caller:   callProc,
		Gateway: gateway,

		maximumReestablishEndeavors: fallbackMaximumReestablishEndeavors,
		fetchPause:             fallbackFetchPause,
		recordPause:            fallbackRecordPause,
		pingSpan:           fallbackPingSpan,
		scheme:             processedWebroute.Scheme,

		//
	}
	c.FoundationFacility = *facility.FreshFoundationFacility(nil, "REDACTED", c)
	for _, selection := range choices {
		selection(c)
	}
	return c, nil
}

//
//
func MaximumReestablishEndeavors(max int) func(*SocketCustomer) {
	return func(c *SocketCustomer) {
		c.maximumReestablishEndeavors = max
	}
}

//
//
func ScanPause(fetchPause time.Duration) func(*SocketCustomer) {
	return func(c *SocketCustomer) {
		c.fetchPause = fetchPause
	}
}

//
//
func RecordPause(recordPause time.Duration) func(*SocketCustomer) {
	return func(c *SocketCustomer) {
		c.recordPause = recordPause
	}
}

//
//
func PingSpan(pingSpan time.Duration) func(*SocketCustomer) {
	return func(c *SocketCustomer) {
		c.pingSpan = pingSpan
	}
}

//
//
func UponReestablish(cb func()) func(*SocketCustomer) {
	return func(c *SocketCustomer) {
		c.uponReestablish = cb
	}
}

//
func (c *SocketCustomer) Text() string {
	return fmt.Sprintf("REDACTED", c.Location, c.Gateway)
}

//
//
func (c *SocketCustomer) UponInitiate() error {
	err := c.call()
	if err != nil {
		return err
	}

	c.RepliesChnl = make(chan kinds.RemoteReply)
	c.PingPongWaitstateClock = metrics.NewTimer()

	c.transmit = make(chan kinds.RemoteSolicit)
	//
	//
	c.reestablishSubsequent = make(chan error, 1)
	//
	//
	c.pendinglist = make(chan kinds.RemoteSolicit, 1)

	c.initiateFetchRecordThreads()
	go c.reestablishProcedure()

	return nil
}

//
//
func (c *SocketCustomer) Halt() error {
	if err := c.FoundationFacility.Halt(); err != nil {
		return err
	}
	//
	c.wg.Wait()
	close(c.RepliesChnl)

	return nil
}

//
func (c *SocketCustomer) EqualsReestablishing() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.reestablishing
}

//
func (c *SocketCustomer) EqualsDynamic() bool {
	return c.EqualsActive() && !c.EqualsReestablishing()
}

//
//
//
func (c *SocketCustomer) Transmit(ctx context.Context, solicit kinds.RemoteSolicit) error {
	select {
	case c.transmit <- solicit:
		c.Tracer.Details("REDACTED", "REDACTED", solicit)
		//
		//
		//
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

//
func (c *SocketCustomer) Invocation(ctx context.Context, procedure string, parameters map[string]any) error {
	solicit, err := kinds.IndexTowardSolicit(c.followingSolicitUUID(), procedure, parameters)
	if err != nil {
		return err
	}
	return c.Transmit(ctx, solicit)
}

//
//
func (c *SocketCustomer) InvocationUsingSeriesParameters(ctx context.Context, procedure string, parameters []any) error {
	solicit, err := kinds.SeriesTowardSolicit(c.followingSolicitUUID(), procedure, parameters)
	if err != nil {
		return err
	}
	return c.Transmit(ctx, solicit)
}

//

func (c *SocketCustomer) followingSolicitUUID() kinds.JsonifaceIntegerUUID {
	c.mtx.Lock()
	id := c.followingRequestUUID
	c.followingRequestUUID++
	c.mtx.Unlock()
	return kinds.JsonifaceIntegerUUID(id)
}

func (c *SocketCustomer) call() error {
	caller := &websocket.Dialer{
		NetDial: c.Caller,
		Proxy:   http.ProxyFromEnvironment,
	}
	readerHeadline := http.Header{}

	//
	if c.Loginname != "REDACTED" && c.Secret != "REDACTED" {
		readerHeadline.Set("REDACTED", "REDACTED"+base64.StdEncoding.EncodeToString([]byte(c.Loginname+"REDACTED"+c.Secret)))
	}

	link, _, err := caller.Dial(c.scheme+"REDACTED"+c.Location+c.Gateway, readerHeadline) //
	if err != nil {
		return err
	}
	c.link = link
	return nil
}

//
//
func (c *SocketCustomer) reestablish() error {
	effort := 0

	c.mtx.Lock()
	c.reestablishing = true
	c.mtx.Unlock()
	defer func() {
		c.mtx.Lock()
		c.reestablishing = false
		c.mtx.Unlock()
	}()

	for {
		variation := time.Duration(commitrand.Float64() * float64(time.Second)) //
		retreatInterval := variation + ((1 << uint(effort)) * time.Second)

		c.Tracer.Details("REDACTED", "REDACTED", effort+1, "REDACTED", retreatInterval)
		time.Sleep(retreatInterval)

		err := c.call()
		if err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", err)
		} else {
			c.Tracer.Details("REDACTED")
			if c.uponReestablish != nil {
				go c.uponReestablish()
			}
			return nil
		}

		effort++

		if effort > c.maximumReestablishEndeavors {
			return fmt.Errorf("REDACTED", err)
		}
	}
}

func (c *SocketCustomer) initiateFetchRecordThreads() {
	c.wg.Add(2)
	c.fetchProcedureExit = make(chan struct{})
	go c.fetchProcedure()
	go c.recordProcedure()
}

func (c *SocketCustomer) handlePendinglist() error {
	select {
	case solicit := <-c.pendinglist:
		if c.recordPause > 0 {
			if err := c.link.SetWriteDeadline(time.Now().Add(c.recordPause)); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		}
		if err := c.link.WriteJSON(solicit); err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", err)
			c.reestablishSubsequent <- err
			//
			c.pendinglist <- solicit
			return err
		}
		c.Tracer.Details("REDACTED", "REDACTED", solicit)
	default:
	}
	return nil
}

func (c *SocketCustomer) reestablishProcedure() {
	for {
		select {
		case authenticFailure := <-c.reestablishSubsequent:
			//
			c.wg.Wait()
			if err := c.reestablish(); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", authenticFailure)
				if err = c.Halt(); err != nil {
					c.Tracer.Failure("REDACTED", "REDACTED", err)
				}

				return
			}
			//
		Cycle:
			for {
				select {
				case <-c.reestablishSubsequent:
				default:
					break Cycle
				}
			}
			err := c.handlePendinglist()
			if err == nil {
				c.initiateFetchRecordThreads()
			}

		case <-c.Exit():
			return
		}
	}
}

//
//
func (c *SocketCustomer) recordProcedure() {
	var metronome *time.Ticker
	if c.pingSpan > 0 {
		//
		metronome = time.NewTicker(c.pingSpan)
	} else {
		//
		metronome = &time.Ticker{C: make(<-chan time.Time)}
	}

	defer func() {
		metronome.Stop()
		c.link.Close()
		//
		//
		//
		//
		c.wg.Done()
	}()

	for {
		select {
		case solicit := <-c.transmit:
			if c.recordPause > 0 {
				if err := c.link.SetWriteDeadline(time.Now().Add(c.recordPause)); err != nil {
					c.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}
			if err := c.link.WriteJSON(solicit); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
				c.reestablishSubsequent <- err
				//
				c.pendinglist <- solicit
				return
			}
		case <-metronome.C:
			if c.recordPause > 0 {
				if err := c.link.SetWriteDeadline(time.Now().Add(c.recordPause)); err != nil {
					c.Tracer.Failure("REDACTED", "REDACTED", err)
				}
			}
			if err := c.link.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
				c.reestablishSubsequent <- err
				return
			}
			c.mtx.Lock()
			c.relayedFinalPingLocated = time.Now()
			c.mtx.Unlock()
			c.Tracer.Diagnose("REDACTED")
		case <-c.fetchProcedureExit:
			return
		case <-c.Exit():
			if err := c.link.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "REDACTED"),
			); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
			}
			return
		}
	}
}

//
//
func (c *SocketCustomer) fetchProcedure() {
	defer func() {
		c.link.Close()
		//
		//
		//
		//
		c.PingPongWaitstateClock.Stop()
		c.wg.Done()
	}()

	c.link.SetPongHandler(func(string) error {
		//
		c.mtx.RLock()
		t := c.relayedFinalPingLocated
		c.mtx.RUnlock()
		c.PingPongWaitstateClock.UpdateSince(t)

		c.Tracer.Diagnose("REDACTED")
		return nil
	})

	for {
		//
		if c.fetchPause > 0 {
			if err := c.link.SetReadDeadline(time.Now().Add(c.fetchPause)); err != nil {
				c.Tracer.Failure("REDACTED", "REDACTED", err)
			}
		}
		_, data, err := c.link.ReadMessage()
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				return
			}

			c.Tracer.Failure("REDACTED", "REDACTED", err)
			close(c.fetchProcedureExit)
			c.reestablishSubsequent <- err
			return
		}

		var reply kinds.RemoteReply
		err = json.Unmarshal(data, &reply)
		if err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", err, "REDACTED", string(data))
			continue
		}

		if err = certifyReplyUUID(reply.ID); err != nil {
			c.Tracer.Failure("REDACTED", "REDACTED", reply.ID, "REDACTED", err)
			continue
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
		//
		//
		//
		//
		//
		//

		c.Tracer.Details("REDACTED", "REDACTED", reply.ID, "REDACTED", log.FreshIdleFormat("REDACTED", reply.Outcome))

		select {
		case <-c.Exit():
		case c.RepliesChnl <- reply:
		}
	}
}

//

//
//
func (c *SocketCustomer) Listen(ctx context.Context, inquire string) error {
	parameters := map[string]any{"REDACTED": inquire}
	return c.Invocation(ctx, "REDACTED", parameters)
}

//
//
func (c *SocketCustomer) Unlisten(ctx context.Context, inquire string) error {
	parameters := map[string]any{"REDACTED": inquire}
	return c.Invocation(ctx, "REDACTED", parameters)
}

//
//
func (c *SocketCustomer) UnlistenEvery(ctx context.Context) error {
	parameters := map[string]any{}
	return c.Invocation(ctx, "REDACTED", parameters)
}
