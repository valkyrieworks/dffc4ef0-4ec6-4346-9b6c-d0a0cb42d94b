//
package node

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

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//
type Settings struct {
	//
	MaximumInitiateLinks int
	//
	FetchDeadline time.Duration
	//
	PersistDeadline time.Duration
	//
	//
	MaximumContentOctets int64
	//
	MaximumHeadingOctets int
	//
	MaximumSolicitClusterExtent int
}

//
func FallbackSettings() *Settings {
	return &Settings{
		MaximumInitiateLinks:  0, //
		FetchDeadline:         10 * time.Second,
		PersistDeadline:        10 * time.Second,
		MaximumContentOctets:        int64(1000000), //
		MaximumHeadingOctets:      1 << 20,        //
		MaximumSolicitClusterExtent: 10,             //
	}
}

//
//
//
//
//
func Attend(observer net.Listener, processor http.Handler, tracer log.Tracer, settings *Settings) error {
	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", observer.Addr()))
	s := &http.Server{
		Handler:           AnteVerificationsProcessor(RestoreAlsoReportProcessor(fallbackProcessor{h: processor}, tracer), settings),
		ReadTimeout:       settings.FetchDeadline,
		ReadHeaderTimeout: settings.FetchDeadline,
		WriteTimeout:      settings.PersistDeadline,
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
func AttendTransportsec(
	observer net.Listener,
	processor http.Handler,
	licenseRecord, tokenRecord string,
	tracer log.Tracer,
	settings *Settings,
) error {
	tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED",
		observer.Addr(), licenseRecord, tokenRecord))
	s := &http.Server{
		Handler:           AnteVerificationsProcessor(RestoreAlsoReportProcessor(fallbackProcessor{h: processor}, tracer), settings),
		ReadTimeout:       settings.FetchDeadline,
		ReadHeaderTimeout: settings.FetchDeadline,
		WriteTimeout:      settings.PersistDeadline,
		MaxHeaderBytes:    settings.MaximumHeadingOctets,
	}
	err := s.ServeTLS(observer, licenseRecord, tokenRecord)

	tracer.Failure("REDACTED", "REDACTED", err)
	return err
}

//
//
//
//
func RecordRemoteReplyHttpsvcFailure(
	w http.ResponseWriter,
	httpsvcCipher int,
	res kinds.RemoteReply,
) error {
	if res.Failure == nil {
		panic("REDACTED")
	}

	jsnOctets, err := json.Marshal(res)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	w.Header().Set("REDACTED", "REDACTED")
	w.WriteHeader(httpsvcCipher)
	_, err = w.Write(jsnOctets)
	return err
}

//
func RecordRemoteReplyHttpsvc(w http.ResponseWriter, res ...kinds.RemoteReply) error {
	return recordRemoteReplyHttpsvc(w, []httpsvcHeadline{}, res...)
}

//
//
//
func RecordStorableRemoteReplyHttpsvc(w http.ResponseWriter, res ...kinds.RemoteReply) error {
	return recordRemoteReplyHttpsvc(w, []httpsvcHeadline{{"REDACTED", "REDACTED"}}, res...)
}

type httpsvcHeadline struct {
	alias  string
	datum string
}

func recordRemoteReplyHttpsvc(w http.ResponseWriter, headings []httpsvcHeadline, res ...kinds.RemoteReply) error {
	var v any
	if len(res) == 1 {
		v = res[0]
	} else {
		v = res
	}

	jsnOctets, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}
	w.Header().Set("REDACTED", "REDACTED")
	for _, heading := range headings {
		w.Header().Set(heading.alias, heading.datum)
	}
	w.WriteHeader(200)
	_, err = w.Write(jsnOctets)
	return err
}

//

//
//
//
func RestoreAlsoReportProcessor(processor http.Handler, tracer log.Tracer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		rww := &replyPersistorEncapsulator{-1, w}
		commence := time.Now()

		rww.Header().Set("REDACTED", fmt.Sprintf("REDACTED", commence.Unix()))

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
				if res, ok := e.(kinds.RemoteReply); ok {
					if wrFault := RecordRemoteReplyHttpsvc(rww, res); wrFault != nil {
						tracer.Failure("REDACTED", "REDACTED", wrFault)
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

					tracer.Failure("REDACTED", "REDACTED", e, "REDACTED", string(debug.Stack()))

					res := kinds.RemoteIntrinsicFailure(kinds.JsonifaceIntegerUUID(-1), err)
					if wrFault := RecordRemoteReplyHttpsvcFailure(rww, http.StatusInternalServerError, res); wrFault != nil {
						tracer.Failure("REDACTED", "REDACTED", wrFault)
					}
				}
			}

			//
			intervalMSEC := time.Since(commence).Nanoseconds() / 1000000
			if rww.Condition == -1 {
				rww.Condition = 200
			}
			tracer.Diagnose("REDACTED",
				"REDACTED", r.Method,
				"REDACTED", r.URL,
				"REDACTED", rww.Condition,
				"REDACTED", intervalMSEC,
				"REDACTED", r.RemoteAddr,
			)
		}()

		processor.ServeHTTP(rww, r)
	})
}

//
type replyPersistorEncapsulator struct {
	Condition int
	http.ReplyPersistor
}

func (w *replyPersistorEncapsulator) RecordHeadline(condition int) {
	w.Condition = condition
	w.ReplyPersistor.WriteHeader(condition)
}

//
func (w *replyPersistorEncapsulator) Divert() (net.Conn, *bufio.ReadWriter, error) {
	return w.ReplyPersistor.(http.Hijacker).Hijack()
}

type fallbackProcessor struct {
	h http.Handler
}

func (h fallbackProcessor) AttendHttpsvc(w http.ResponseWriter, r *http.Request) {
	h.h.ServeHTTP(w, r)
}

//
//
func Overhear(location string, maximumInitiateLinks int) (observer net.Listener, err error) {
	fragments := strings.SplitN(location, "REDACTED", 2)
	if len(fragments) != 2 {
		return nil, fmt.Errorf(
			"REDACTED",
			location,
		)
	}
	schema, location := fragments[0], fragments[1]
	observer, err = net.Listen(schema, location)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", location, err)
	}
	if maximumInitiateLinks > 0 {
		observer = netutil.LimitListener(observer, maximumInitiateLinks)
	}

	return observer, nil
}

//

//
//
//
func AnteVerificationsProcessor(following http.Handler, settings *Settings) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//
		r.Body = http.MaxBytesReader(w, r.Body, settings.MaximumContentOctets)

		//
		//
		//
		if settings.MaximumSolicitClusterExtent > 0 {
			var solicits []kinds.RemoteSolicit
			var replies []kinds.RemoteReply
			var err error

			data, err := io.ReadAll(r.Body)
			if err != nil {
				res := kinds.RemoteUnfitSolicitFailure(nil, fmt.Errorf("REDACTED", err))
				_ = RecordRemoteReplyHttpsvcFailure(w, http.StatusBadRequest, res)
				return
			}

			err = json.Unmarshal(data, &solicits)
			//
			//
			if err == nil {
				//
				if len(solicits) > settings.MaximumSolicitClusterExtent {
					res := kinds.RemoteUnfitSolicitFailure(nil, fmt.Errorf("REDACTED", settings.MaximumSolicitClusterExtent))
					replies = append(replies, res)
					_ = RecordRemoteReplyHttpsvc(w, replies...)
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
