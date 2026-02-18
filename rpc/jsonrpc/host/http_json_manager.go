package host

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"

	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//

//
func createJsonrpcManager(functionIndex map[string]*RPCFunction, tracer log.Tracer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			res := kinds.RPCCorruptQueryFault(nil,
				fmt.Errorf("REDACTED", err),
			)
			if writerErr := RecordRPCReplyHTTPFault(w, http.StatusBadRequest, res); writerErr != nil {
				tracer.Fault("REDACTED", "REDACTED", writerErr)
			}
			return
		}

		//
		//
		if len(b) == 0 {
			recordCatalogOfTermini(w, r, functionIndex)
			return
		}

		//
		var (
			queries  []kinds.RPCQuery
			replies []kinds.RPCAnswer
		)
		if err := json.Unmarshal(b, &queries); err != nil {
			//
			var query kinds.RPCQuery
			if err := json.Unmarshal(b, &query); err != nil {
				res := kinds.RPCAnalyzeFault(fmt.Errorf("REDACTED", err))
				if writerErr := RecordRPCReplyHTTPFault(w, http.StatusInternalServerError, res); writerErr != nil {
					tracer.Fault("REDACTED", "REDACTED", writerErr)
				}
				return
			}
			queries = []kinds.RPCQuery{query}
		}

		//
		//
		//
		//
		repository := true
		for _, query := range queries {

			//
			//
			if query.ID == nil {
				tracer.Diagnose(
					"REDACTED",
					"REDACTED", query,
				)
				continue
			}
			if len(r.URL.Path) > 1 {
				replies = append(
					replies,
					kinds.RPCCorruptQueryFault(query.ID, fmt.Errorf("REDACTED", r.URL.Path)),
				)
				repository = false
				continue
			}
			rpcFunction, ok := functionIndex[query.Procedure]
			if !ok || (rpcFunction.ws) {
				replies = append(replies, kinds.RPCProcedureNegateLocatedFault(query.ID))
				repository = false
				continue
			}
			ctx := &kinds.Context{JSONRequest: &query, HTTPRequest: r}
			args := []reflect.Value{reflect.ValueOf(ctx)}
			if len(query.Options) > 0 {
				fnArgs, err := jsonOptionsToArgs(rpcFunction, query.Options)
				if err != nil {
					replies = append(
						replies,
						kinds.RPCCorruptOptionsFault(query.ID, fmt.Errorf("REDACTED", err)),
					)
					repository = false
					continue
				}
				args = append(args, fnArgs...)
			}

			if repository && !rpcFunction.storableWithArgs(args) {
				repository = false
			}

			yields := rpcFunction.f.Call(args)
			outcome, err := unmirrorOutcome(yields)
			if err != nil {
				replies = append(replies, kinds.RPCIntrinsicFault(query.ID, err))
				continue
			}
			replies = append(replies, kinds.NewRPCSuccessReply(query.ID, outcome))
		}

		if len(replies) > 0 {
			var writerErr error
			if repository {
				writerErr = RecordStorableRPCReplyHTTP(w, replies...)
			} else {
				writerErr = RecordRPCReplyHTTP(w, replies...)
			}
			if writerErr != nil {
				tracer.Fault("REDACTED", "REDACTED", writerErr)
			}
		}
	}
}

func processCorruptJsonrpcRoutes(following http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//
		//
		if r.URL.Path != "REDACTED" {
			http.NotFound(w, r)
			return
		}

		following(w, r)
	}
}

func indexOptionsToArgs(
	rpcFunction *RPCFunction,
	options map[string]json.RawMessage,
	argsDisplacement int,
) ([]reflect.Value, error) {
	items := make([]reflect.Value, len(rpcFunction.argumentLabels))
	for i, argumentLabel := range rpcFunction.argumentLabels {
		argumentKind := rpcFunction.args[i+argsDisplacement]

		if p, ok := options[argumentLabel]; ok && p != nil && len(p) > 0 {
			val := reflect.New(argumentKind)
			err := cometjson.Unserialize(p, val.Interface())
			if err != nil {
				return nil, err
			}
			items[i] = val.Elem()
		} else { //
			items[i] = reflect.Zero(argumentKind)
		}
	}

	return items, nil
}

func listOptionsToArgs(
	rpcFunction *RPCFunction,
	options []json.RawMessage,
	argsDisplacement int,
) ([]reflect.Value, error) {
	if len(rpcFunction.argumentLabels) != len(options) {
		return nil, fmt.Errorf("REDACTED",
			len(rpcFunction.argumentLabels), rpcFunction.argumentLabels, len(options), options)
	}

	items := make([]reflect.Value, len(options))
	for i, p := range options {
		argumentKind := rpcFunction.args[i+argsDisplacement]
		val := reflect.New(argumentKind)
		err := cometjson.Unserialize(p, val.Interface())
		if err != nil {
			return nil, err
		}
		items[i] = val.Elem()
	}
	return items, nil
}

//
//
//
//
//
//
//
func jsonOptionsToArgs(rpcFunction *RPCFunction, raw []byte) ([]reflect.Value, error) {
	const argsDisplacement = 1

	//
	//
	var m map[string]json.RawMessage
	err := json.Unmarshal(raw, &m)
	if err == nil {
		return indexOptionsToArgs(rpcFunction, m, argsDisplacement)
	}

	//
	var a []json.RawMessage
	err = json.Unmarshal(raw, &a)
	if err == nil {
		return listOptionsToArgs(rpcFunction, a, argsDisplacement)
	}

	//
	return nil, fmt.Errorf("REDACTED", err)
}

//
func recordCatalogOfTermini(w http.ResponseWriter, r *http.Request, functionIndex map[string]*RPCFunction) {
	noArgumentLabels := []string{}
	argumentLabels := []string{}
	for label, functionData := range functionIndex {
		if len(functionData.args) == 0 {
			noArgumentLabels = append(noArgumentLabels, label)
		} else {
			argumentLabels = append(argumentLabels, label)
		}
	}
	sort.Strings(noArgumentLabels)
	sort.Strings(argumentLabels)
	buf := new(bytes.Buffer)
	buf.WriteString("REDACTED")
	buf.WriteString("REDACTED")

	for _, label := range noArgumentLabels {
		linkage := fmt.Sprintf("REDACTED", r.Host, label)
		fmt.Fprintf(buf, "REDACTED", linkage, linkage)
	}

	buf.WriteString("REDACTED")
	for _, label := range argumentLabels {
		linkage := fmt.Sprintf("REDACTED", r.Host, label)
		functionData := functionIndex[label]
		for i, argumentLabel := range functionData.argumentLabels {
			linkage += argumentLabel + "REDACTED"
			if i < len(functionData.argumentLabels)-1 {
				linkage += "REDACTED"
			}
		}
		fmt.Fprintf(buf, "REDACTED", linkage, linkage)
	}
	buf.WriteString("REDACTED")
	w.Header().Set("REDACTED", "REDACTED")
	w.WriteHeader(200)
	w.Write(buf.Bytes()) //
}
