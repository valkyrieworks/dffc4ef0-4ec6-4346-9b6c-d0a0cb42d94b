package host

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	cometjson "github.com/valkyrieworks/utils/json"
	"github.com/valkyrieworks/utils/log"
	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

//

var reInteger = regexp.MustCompile("REDACTED")

//
func createHTTPManager(rpcFunction *RPCFunction, tracer log.Tracer) func(http.ResponseWriter, *http.Request) {
	//
	mockUID := kinds.JsonrpcIntegerUID(-1) //

	//
	if rpcFunction.ws {
		return func(w http.ResponseWriter, r *http.Request) {
			res := kinds.RPCProcedureNegateLocatedFault(mockUID)
			if writerErr := RecordRPCReplyHTTPFault(w, http.StatusNotFound, res); writerErr != nil {
				tracer.Fault("REDACTED", "REDACTED", writerErr)
			}
		}
	}

	//
	return func(w http.ResponseWriter, r *http.Request) {
		tracer.Diagnose("REDACTED", "REDACTED", r)

		ctx := &kinds.Context{HTTPRequest: r}
		args := []reflect.Value{reflect.ValueOf(ctx)}

		fnArgs, err := httpOptionsToArgs(rpcFunction, r)
		if err != nil {
			res := kinds.RPCCorruptOptionsFault(mockUID,
				fmt.Errorf("REDACTED", err),
			)
			if writerErr := RecordRPCReplyHTTPFault(w, http.StatusInternalServerError, res); writerErr != nil {
				tracer.Fault("REDACTED", "REDACTED", writerErr)
			}
			return
		}
		args = append(args, fnArgs...)

		yields := rpcFunction.f.Call(args)

		tracer.Diagnose("REDACTED", "REDACTED", r.URL.Path, "REDACTED", args, "REDACTED", yields)
		outcome, err := unmirrorOutcome(yields)
		if err != nil {
			if err := RecordRPCReplyHTTPFault(w, http.StatusInternalServerError,
				kinds.RPCIntrinsicFault(mockUID, err)); err != nil {
				tracer.Fault("REDACTED", "REDACTED", err)
				return
			}
			return
		}

		reply := kinds.NewRPCSuccessReply(mockUID, outcome)
		if rpcFunction.storableWithArgs(args) {
			err = RecordStorableRPCReplyHTTP(w, reply)
		} else {
			err = RecordRPCReplyHTTP(w, reply)
		}
		if err != nil {
			tracer.Fault("REDACTED", "REDACTED", err)
			return
		}
	}
}

//
//
func httpOptionsToArgs(rpcFunction *RPCFunction, r *http.Request) ([]reflect.Value, error) {
	//
	const argsDisplacement = 1

	items := make([]reflect.Value, len(rpcFunction.argumentLabels))

	for i, label := range rpcFunction.argumentLabels {
		argumentKind := rpcFunction.args[i+argsDisplacement]

		items[i] = reflect.Zero(argumentKind) //

		arg := fetchArgument(r, label)
		//

		if arg == "REDACTED" {
			continue
		}

		v, ok, err := notJSONStringToArgument(argumentKind, arg)
		if err != nil {
			return nil, err
		}
		if ok {
			items[i] = v
			continue
		}

		items[i], err = jsonStringToArgument(argumentKind, arg)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func jsonStringToArgument(rt reflect.Type, arg string) (reflect.Value, error) {
	rv := reflect.New(rt)
	err := cometjson.Unserialize([]byte(arg), rv.Interface())
	if err != nil {
		return rv, err
	}
	rv = rv.Elem()
	return rv, nil
}

func notJSONStringToArgument(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	if rt.Kind() == reflect.Ptr {
		rv1, ok, err := notJSONStringToArgument(rt.Elem(), arg)
		switch {
		case err != nil:
			return reflect.Value{}, false, err
		case ok:
			rv := reflect.New(rt.Elem())
			rv.Elem().Set(rv1)
			return rv, true, nil
		default:
			return reflect.Value{}, false, nil
		}
	} else {
		return _notjsonstringtoargument(rt, arg)
	}
}

//
func _notjsonstringtoargument(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	isIntegerString := reInteger.Match([]byte(arg))
	isCitedString := strings.HasPrefix(arg, "REDACTED") && strings.HasSuffix(arg, "REDACTED")
	isHexString := strings.HasPrefix(strings.ToLower(arg), "REDACTED")

	var awaitingString, awaitingOctetSection, awaitingInteger bool
	switch rt.Kind() {
	case reflect.Int,
		reflect.Uint,
		reflect.Int8,
		reflect.Uint8,
		reflect.Int16,
		reflect.Uint16,
		reflect.Int32,
		reflect.Uint32,
		reflect.Int64,
		reflect.Uint64:
		awaitingInteger = true
	case reflect.String:
		awaitingString = true
	case reflect.Slice:
		awaitingOctetSection = rt.Elem().Kind() == reflect.Uint8
	}

	if isIntegerString && awaitingInteger {
		qargument := "REDACTED" + arg + "REDACTED"
		rv, err := jsonStringToArgument(rt, qargument)
		if err != nil {
			return rv, false, err
		}

		return rv, true, nil
	}

	if isHexString {
		if !awaitingString && !awaitingOctetSection {
			err := fmt.Errorf("REDACTED",
				rt.Kind().String())
			return reflect.ValueOf(nil), false, err
		}

		var item []byte
		item, err := hex.DecodeString(arg[2:])
		if err != nil {
			return reflect.ValueOf(nil), false, err
		}
		if rt.Kind() == reflect.String {
			return reflect.ValueOf(string(item)), true, nil
		}
		return reflect.ValueOf(item), true, nil
	}

	if isCitedString && awaitingOctetSection {
		v := reflect.New(reflect.TypeOf("REDACTED"))
		err := cometjson.Unserialize([]byte(arg), v.Interface())
		if err != nil {
			return reflect.ValueOf(nil), false, err
		}
		v = v.Elem()
		return reflect.ValueOf([]byte(v.String())), true, nil
	}

	return reflect.ValueOf(nil), false, nil
}

func fetchArgument(r *http.Request, argument string) string {
	s := r.URL.Query().Get(argument)
	if s == "REDACTED" {
		s = r.FormValue(argument)
	}
	return s
}
