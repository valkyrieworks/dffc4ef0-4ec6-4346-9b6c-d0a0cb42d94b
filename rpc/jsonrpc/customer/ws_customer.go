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

	"github.com/valkyrieworks/utils/log"
	engineseed "github.com/valkyrieworks/utils/random"
	"github.com/valkyrieworks/utils/daemon"
	engineconnect "github.com/valkyrieworks/utils/align"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

const (
	standardMaximumReestablishTries = 25
	standardRecordWait            = 0
	standardReadWait             = 0
	standardPingDuration           = 0
)

//
//
//
//
type WSCustomer struct {
	link *websocket.Conn

	Location  string //
	Terminus string //
	Username string
	Secret string

	Caller func(string, string) (net.Conn, error)

	//
	//
	RepliesChan chan kinds.RPCAnswer

	//
	onReestablish func()

	//
	transmit            chan kinds.RPCQuery //
	pending         chan kinds.RPCQuery //
	reestablishAfter  chan error            //
	readProcedureExit chan struct{}         //

	//
	maximumReestablishTries int

	//
	protocol string

	wg sync.WaitGroup

	mtx            engineconnect.ReadwriteLock
	relayedFinalPingAt time.Time
	reestablishing   bool
	followingRequestUID      int
	//

	//
	recordWait time.Duration

	//
	readWait time.Duration

	//
	pingDuration time.Duration

	daemon.RootDaemon

	//
	//
	PingPongWaitperiodClock metrics.Timer
}

//
//
//
//
func NewWS(distantAddress, gateway string, options ...func(*WSCustomer)) (*WSCustomer, error) {
	analyzedURL, err := newAnalyzedURL(distantAddress)
	if err != nil {
		return nil, err
	}
	//
	if analyzedURL.Scheme == schemaHttps {
		analyzedURL.Scheme = schemaWSS
	} else if analyzedURL.Scheme != schemaWSS {
		analyzedURL.Scheme = schemaWS
	}

	//
	username := "REDACTED"
	secret := "REDACTED"
	if analyzedURL.User.String() != "REDACTED" {
		username = analyzedURL.User.Username()
		secret, _ = analyzedURL.User.Password()
	}

	callFn, err := createHTTPCaller(distantAddress)
	if err != nil {
		return nil, err
	}

	c := &WSCustomer{
		Location:  analyzedURL.FetchClippedMachineWithRoute(),
		Username: username,
		Secret: secret,
		Caller:   callFn,
		Terminus: gateway,

		maximumReestablishTries: standardMaximumReestablishTries,
		readWait:             standardReadWait,
		recordWait:            standardRecordWait,
		pingDuration:           standardPingDuration,
		protocol:             analyzedURL.Scheme,

		//
	}
	c.RootDaemon = *daemon.NewRootDaemon(nil, "REDACTED", c)
	for _, setting := range options {
		setting(c)
	}
	return c, nil
}

//
//
func MaximumReestablishTries(max int) func(*WSCustomer) {
	return func(c *WSCustomer) {
		c.maximumReestablishTries = max
	}
}

//
//
func ReadWait(readWait time.Duration) func(*WSCustomer) {
	return func(c *WSCustomer) {
		c.readWait = readWait
	}
}

//
//
func RecordWait(recordWait time.Duration) func(*WSCustomer) {
	return func(c *WSCustomer) {
		c.recordWait = recordWait
	}
}

//
//
func PingDuration(pingDuration time.Duration) func(*WSCustomer) {
	return func(c *WSCustomer) {
		c.pingDuration = pingDuration
	}
}

//
//
func OnReestablish(cb func()) func(*WSCustomer) {
	return func(c *WSCustomer) {
		c.onReestablish = cb
	}
}

//
func (c *WSCustomer) String() string {
	return fmt.Sprintf("REDACTED", c.Location, c.Terminus)
}

//
//
func (c *WSCustomer) OnBegin() error {
	err := c.call()
	if err != nil {
		return err
	}

	c.RepliesChan = make(chan kinds.RPCAnswer)
	c.PingPongWaitperiodClock = metrics.NewTimer()

	c.transmit = make(chan kinds.RPCQuery)
	//
	//
	c.reestablishAfter = make(chan error, 1)
	//
	//
	c.pending = make(chan kinds.RPCQuery, 1)

	c.beginReadRecordProcedures()
	go c.reestablishProcedure()

	return nil
}

//
//
func (c *WSCustomer) Halt() error {
	if err := c.RootDaemon.Halt(); err != nil {
		return err
	}
	//
	c.wg.Wait()
	close(c.RepliesChan)

	return nil
}

//
func (c *WSCustomer) IsReestablishing() bool {
	c.mtx.RLock()
	defer c.mtx.RUnlock()
	return c.reestablishing
}

//
func (c *WSCustomer) IsEnabled() bool {
	return c.IsActive() && !c.IsReestablishing()
}

//
//
//
func (c *WSCustomer) Transmit(ctx context.Context, query kinds.RPCQuery) error {
	select {
	case c.transmit <- query:
		c.Tracer.Details("REDACTED", "REDACTED", query)
		//
		//
		//
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

//
func (c *WSCustomer) Invoke(ctx context.Context, procedure string, options map[string]any) error {
	query, err := kinds.IndexToQuery(c.followingQueryUID(), procedure, options)
	if err != nil {
		return err
	}
	return c.Transmit(ctx, query)
}

//
//
func (c *WSCustomer) InvokeWithListOptions(ctx context.Context, procedure string, options []any) error {
	query, err := kinds.ListToQuery(c.followingQueryUID(), procedure, options)
	if err != nil {
		return err
	}
	return c.Transmit(ctx, query)
}

//

func (c *WSCustomer) followingQueryUID() kinds.JsonrpcIntegerUID {
	c.mtx.Lock()
	id := c.followingRequestUID
	c.followingRequestUID++
	c.mtx.Unlock()
	return kinds.JsonrpcIntegerUID(id)
}

func (c *WSCustomer) call() error {
	caller := &websocket.Dialer{
		NetDial: c.Caller,
		Proxy:   http.ProxyFromEnvironment,
	}
	readerHeading := http.Header{}

	//
	if c.Username != "REDACTED" && c.Secret != "REDACTED" {
		readerHeading.Set("REDACTED", "REDACTED"+base64.StdEncoding.EncodeToString([]byte(c.Username+"REDACTED"+c.Secret)))
	}

	link, _, err := caller.Dial(c.protocol+"REDACTED"+c.Location+c.Terminus, readerHeading) //
	if err != nil {
		return err
	}
	c.link = link
	return nil
}

//
//
func (c *WSCustomer) reestablish() error {
	endeavor := 0

	c.mtx.Lock()
	c.reestablishing = true
	c.mtx.Unlock()
	defer func() {
		c.mtx.Lock()
		c.reestablishing = false
		c.mtx.Unlock()
	}()

	for {
		variance := time.Duration(engineseed.Float64() * float64(time.Second)) //
		retreatPeriod := variance + ((1 << uint(endeavor)) * time.Second)

		c.Tracer.Details("REDACTED", "REDACTED", endeavor+1, "REDACTED", retreatPeriod)
		time.Sleep(retreatPeriod)

		err := c.call()
		if err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", err)
		} else {
			c.Tracer.Details("REDACTED")
			if c.onReestablish != nil {
				go c.onReestablish()
			}
			return nil
		}

		endeavor++

		if endeavor > c.maximumReestablishTries {
			return fmt.Errorf("REDACTED", err)
		}
	}
}

func (c *WSCustomer) beginReadRecordProcedures() {
	c.wg.Add(2)
	c.readProcedureExit = make(chan struct{})
	go c.readProcedure()
	go c.recordProcedure()
}

func (c *WSCustomer) handlePending() error {
	select {
	case query := <-c.pending:
		if c.recordWait > 0 {
			if err := c.link.SetWriteDeadline(time.Now().Add(c.recordWait)); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		}
		if err := c.link.WriteJSON(query); err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", err)
			c.reestablishAfter <- err
			//
			c.pending <- query
			return err
		}
		c.Tracer.Details("REDACTED", "REDACTED", query)
	default:
	}
	return nil
}

func (c *WSCustomer) reestablishProcedure() {
	for {
		select {
		case sourceFault := <-c.reestablishAfter:
			//
			c.wg.Wait()
			if err := c.reestablish(); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", sourceFault)
				if err = c.Halt(); err != nil {
					c.Tracer.Fault("REDACTED", "REDACTED", err)
				}

				return
			}
			//
		Cycle:
			for {
				select {
				case <-c.reestablishAfter:
				default:
					break Cycle
				}
			}
			err := c.handlePending()
			if err == nil {
				c.beginReadRecordProcedures()
			}

		case <-c.Exit():
			return
		}
	}
}

//
//
func (c *WSCustomer) recordProcedure() {
	var timer *time.Ticker
	if c.pingDuration > 0 {
		//
		timer = time.NewTicker(c.pingDuration)
	} else {
		//
		timer = &time.Ticker{C: make(<-chan time.Time)}
	}

	defer func() {
		timer.Stop()
		c.link.Close()
		//
		//
		//
		//
		c.wg.Done()
	}()

	for {
		select {
		case query := <-c.transmit:
			if c.recordWait > 0 {
				if err := c.link.SetWriteDeadline(time.Now().Add(c.recordWait)); err != nil {
					c.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}
			if err := c.link.WriteJSON(query); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
				c.reestablishAfter <- err
				//
				c.pending <- query
				return
			}
		case <-timer.C:
			if c.recordWait > 0 {
				if err := c.link.SetWriteDeadline(time.Now().Add(c.recordWait)); err != nil {
					c.Tracer.Fault("REDACTED", "REDACTED", err)
				}
			}
			if err := c.link.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
				c.reestablishAfter <- err
				return
			}
			c.mtx.Lock()
			c.relayedFinalPingAt = time.Now()
			c.mtx.Unlock()
			c.Tracer.Diagnose("REDACTED")
		case <-c.readProcedureExit:
			return
		case <-c.Exit():
			if err := c.link.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(websocket.CloseNormalClosure, "REDACTED"),
			); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
			}
			return
		}
	}
}

//
//
func (c *WSCustomer) readProcedure() {
	defer func() {
		c.link.Close()
		//
		//
		//
		//
		c.PingPongWaitperiodClock.Stop()
		c.wg.Done()
	}()

	c.link.SetPongHandler(func(string) error {
		//
		c.mtx.RLock()
		t := c.relayedFinalPingAt
		c.mtx.RUnlock()
		c.PingPongWaitperiodClock.UpdateSince(t)

		c.Tracer.Diagnose("REDACTED")
		return nil
	})

	for {
		//
		if c.readWait > 0 {
			if err := c.link.SetReadDeadline(time.Now().Add(c.readWait)); err != nil {
				c.Tracer.Fault("REDACTED", "REDACTED", err)
			}
		}
		_, data, err := c.link.ReadMessage()
		if err != nil {
			if !websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure) {
				return
			}

			c.Tracer.Fault("REDACTED", "REDACTED", err)
			close(c.readProcedureExit)
			c.reestablishAfter <- err
			return
		}

		var reply kinds.RPCAnswer
		err = json.Unmarshal(data, &reply)
		if err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", err, "REDACTED", string(data))
			continue
		}

		if err = certifyReplyUID(reply.ID); err != nil {
			c.Tracer.Fault("REDACTED", "REDACTED", reply.ID, "REDACTED", err)
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

		c.Tracer.Details("REDACTED", "REDACTED", reply.ID, "REDACTED", log.NewIdleFormat("REDACTED", reply.Outcome))

		select {
		case <-c.Exit():
		case c.RepliesChan <- reply:
		}
	}
}

//

//
//
func (c *WSCustomer) Enrol(ctx context.Context, inquire string) error {
	options := map[string]any{"REDACTED": inquire}
	return c.Invoke(ctx, "REDACTED", options)
}

//
//
func (c *WSCustomer) Deenroll(ctx context.Context, inquire string) error {
	options := map[string]any{"REDACTED": inquire}
	return c.Invoke(ctx, "REDACTED", options)
}

//
//
func (c *WSCustomer) DeenrollAll(ctx context.Context) error {
	options := map[string]any{}
	return c.Invoke(ctx, "REDACTED", options)
}
