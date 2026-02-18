//
package host

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"golang.org/x/net/netutil"

	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//
type Settings struct {
	//
	MaximumAccessLinks int
	//
	ReadDeadline time.Duration
	//
	RecordDeadline time.Duration
	//
	//
	MaximumContentOctets int64
	//
	MaximumHeadingOctets int
	//
	MaximumQueryClusterVolume int
}

//
func StandardSettings() *Settings {
	return &Settings{
		MaximumAccessLinks:  0, //
		ReadDeadline:         10 * time.Second,
		RecordDeadline:        10 * time.Second,
		MaximumContentOctets:        int64(1000000), //
		MaximumHeadingOctets:      1 << 20,        //
		MaximumQueryClusterVolume: 10,             //
	}
}

//
//
//
//
//
func Attend(observer net.Listener, manager http.Handler, tracer log.Tracer, settings *Settings) error {
	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", observer.Addr()))
	s := &http.Server{
		Handler:           PreValidationsManager(RecoupAndTraceManager(standardManager{h: manager}, tracer), settings),
		ReadTimeout:       settings.ReadDeadline,
		ReadHeaderTimeout: settings.ReadDeadline,
		WriteTimeout:      settings.RecordDeadline,
		MaxHeaderBytes:    settings.MaximumHeadingOctets,
	}
	err := s.Serve(observer)
	tracer.Details("REDACTED", "REDACTED", err)
	return err
}

//
//
//
//
//
func AttendTLS(
	observer net.Listener,
	manager http.Handler,
	tokenEntry, keyEntry string,
	tracer log.Tracer,
	settings *Settings,
) error {
	tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED",
		observer.Addr(), tokenEntry, keyEntry))
	s := &http.Server{
		Handler:           PreValidationsManager(RecoupAndTraceManager(standardManager{h: manager}, tracer), settings),
		ReadTimeout:       settings.ReadDeadline,
		ReadHeaderTimeout: settings.ReadDeadline,
		WriteTimeout:      settings.RecordDeadline,
		MaxHeaderBytes:    settings.MaximumHeadingOctets,
	}
	err := s.ServeTLS(observer, tokenEntry, keyEntry)

	tracer.Fault("REDACTED", "REDACTED", err)
	return err
}

//
//
//
//
func RecordRPCReplyHTTPFault(
	w http.ResponseWriter,
	httpCode int,
	res kinds.RPCAnswer,
) error {
	if res.Fault == nil {
		panic("REDACTED")
	}

	jsonOctets, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	w.Header().Set("REDACTED", "REDACTED")
	w.WriteHeader(httpCode)
	_, err = w.Write(jsonOctets)
	return err
}

//
func RecordRPCReplyHTTP(w http.ResponseWriter, res ...kinds.RPCAnswer) error {
	return recordRPCReplyHTTP(w, []httpHeading{}, res...)
}

//
//
//
func RecordStorableRPCReplyHTTP(w http.ResponseWriter, res ...kinds.RPCAnswer) error {
	return recordRPCReplyHTTP(w, []httpHeading{{"REDACTED", "REDACTED"}}, res...)
}

type httpHeading struct {
	label  string
	item string
}

func recordRPCReplyHTTP(w http.ResponseWriter, headings []httpHeading, res ...kinds.RPCAnswer) error {
	var v any
	if len(res) == 1 {
		v = res[0]
	} else {
		v = res
	}

	jsonOctets, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	w.Header().Set("REDACTED", "REDACTED")
	for _, heading := range headings {
		w.Header().Set(heading.label, heading.item)
	}
	w.WriteHeader(200)
	_, err = w.Write(jsonOctets)
	return err
}

//

//
//
//
func RecoupAndTraceManager(manager http.Handler, tracer log.Tracer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		rww := &replyRecorderAdapter{-1, w}
		initiate := time.Now()

		rww.Header().Set("REDACTED", fmt.Sprintf("REDACTED", initiate.Unix()))

		defer func() {
			//
			//
			//
			//
			//
			//
			//
			if e := recover(); e != nil {
				fmt.Fprintf(os.Stderr, "REDACTED", e, string(debug.Stack()))
				w.WriteHeader(500)
			}
		}()

		defer func() {
			//
			//
			//
			if e := recover(); e != nil {
				//
				if res, ok := e.(kinds.RPCAnswer); ok {
					if writerErr := RecordRPCReplyHTTP(rww, res); writerErr != nil {
						tracer.Fault("REDACTED", "REDACTED", writerErr)
					}
				} else {
					//
					var err error
					switch e := e.(type) {
					case error:
						err = e
					case string:
						err = errors.New(e)
					case fmt.Stringer:
						err = errors.New(e.String())
					default:
					}

					tracer.Fault("REDACTED", "REDACTED", e, "REDACTED", string(debug.Stack()))

					res := kinds.RPCIntrinsicFault(kinds.JsonrpcIntegerUID(-1), err)
					if writerErr := RecordRPCReplyHTTPFault(rww, http.StatusInternalServerError, res); writerErr != nil {
						tracer.Fault("REDACTED", "REDACTED", writerErr)
					}
				}
			}

			//
			periodMillis := time.Since(initiate).Nanoseconds() / 1000000
			if rww.Status == -1 {
				rww.Status = 200
			}
			tracer.Diagnose("REDACTED",
				"REDACTED", r.Method,
				"REDACTED", r.URL,
				"REDACTED", rww.Status,
				"REDACTED", periodMillis,
				"REDACTED", r.RemoteAddr,
			)
		}()

		manager.ServeHTTP(rww, r)
	})
}

//
type replyRecorderAdapter struct {
	Status int
	http.ReplyRecorder
}

func (w *replyRecorderAdapter) RecordHeading(state int) {
	w.Status = state
	w.ReplyRecorder.WriteHeader(state)
}

//
func (w *replyRecorderAdapter) Divert() (net.Conn, *bufio.ReadWriter, error) {
	return w.ReplyRecorder.(http.Hijacker).Hijack()
}

type standardManager struct {
	h http.Handler
}

func (h standardManager) AttendHTTP(w http.ResponseWriter, r *http.Request) {
	h.h.ServeHTTP(w, r)
}

//
//
func Observe(address string, maximumAccessLinks int) (observer net.Listener, err error) {
	segments := strings.SplitN(address, "REDACTED", 2)
	if len(segments) != 2 {
		return nil, fmt.Errorf(
			"REDACTED",
			address,
		)
	}
	schema, address := segments[0], segments[1]
	observer, err = net.Listen(schema, address)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", address, err)
	}
	if maximumAccessLinks > 0 {
		observer = netutil.LimitListener(observer, maximumAccessLinks)
	}

	return observer, nil
}

//

//
//
//
func PreValidationsManager(following http.Handler, settings *Settings) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		r.Body = http.MaxBytesReader(w, r.Body, settings.MaximumContentOctets)

		//
		//
		//
		if settings.MaximumQueryClusterVolume > 0 {
			var queries []kinds.RPCQuery
			var replies []kinds.RPCAnswer
			var err error

			data, err := io.ReadAll(r.Body)
			if err != nil {
				res := kinds.RPCCorruptQueryFault(nil, fmt.Errorf("REDACTED", err))
				_ = RecordRPCReplyHTTPFault(w, http.StatusBadRequest, res)
				return
			}

			err = json.Unmarshal(data, &queries)
			//
			//
			if err == nil {
				//
				if len(queries) > settings.MaximumQueryClusterVolume {
					res := kinds.RPCCorruptQueryFault(nil, fmt.Errorf("REDACTED", settings.MaximumQueryClusterVolume))
					replies = append(replies, res)
					_ = RecordRPCReplyHTTP(w, replies...)
					return
				}
			}

			//
			r.Body = io.NopCloser(bytes.NewBuffer(data))
		}

		//
		following.ServeHTTP(w, r)
	})
}
