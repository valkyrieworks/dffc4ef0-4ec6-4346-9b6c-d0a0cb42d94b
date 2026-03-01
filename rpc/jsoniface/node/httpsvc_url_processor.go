package node

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//

var againInteger = regexp.MustCompile("REDACTED")

//
func createHttpsvcProcessor(remoteMethod *RemoteMethod, tracer log.Tracer) func(http.ResponseWriter, *http.Request) {
	//
	placeholderUUID := kinds.JsonifaceIntegerUUID(-1) //

	//
	if remoteMethod.ws {
		return func(w http.ResponseWriter, r *http.Request) {
			res := kinds.RemoteProcedureNegationDetectedFailure(placeholderUUID)
			if wrFault := RecordRemoteReplyHttpsvcFailure(w, http.StatusNotFound, res); wrFault != nil {
				tracer.Failure("REDACTED", "REDACTED", wrFault)
			}
		}
	}

	//
	return func(w http.ResponseWriter, r *http.Request) {
		tracer.Diagnose("REDACTED", "REDACTED", r)

		ctx := &kinds.Env{HttpsvcRequest: r}
		arguments := []reflect.Value{reflect.ValueOf(ctx)}

		procArguments, err := httpsvcParametersTowardArguments(remoteMethod, r)
		if err != nil {
			res := kinds.RemoteUnfitParametersFailure(placeholderUUID,
				fmt.Errorf("REDACTED", err),
			)
			if wrFault := RecordRemoteReplyHttpsvcFailure(w, http.StatusInternalServerError, res); wrFault != nil {
				tracer.Failure("REDACTED", "REDACTED", wrFault)
			}
			return
		}
		arguments = append(arguments, procArguments...)

		yields := remoteMethod.f.Call(arguments)

		tracer.Diagnose("REDACTED", "REDACTED", r.URL.Path, "REDACTED", arguments, "REDACTED", yields)
		outcome, err := unmirrorOutcome(yields)
		if err != nil {
			if err := RecordRemoteReplyHttpsvcFailure(w, http.StatusInternalServerError,
				kinds.RemoteIntrinsicFailure(placeholderUUID, err)); err != nil {
				tracer.Failure("REDACTED", "REDACTED", err)
				return
			}
			return
		}

		reply := kinds.FreshRemoteTriumphReply(placeholderUUID, outcome)
		if remoteMethod.storableUsingArguments(arguments) {
			err = RecordStorableRemoteReplyHttpsvc(w, reply)
		} else {
			err = RecordRemoteReplyHttpsvc(w, reply)
		}
		if err != nil {
			tracer.Failure("REDACTED", "REDACTED", err)
			return
		}
	}
}

//
//
func httpsvcParametersTowardArguments(remoteMethod *RemoteMethod, r *http.Request) ([]reflect.Value, error) {
	//
	const argumentsDisplacement = 1

	items := make([]reflect.Value, len(remoteMethod.argumentIdentifiers))

	for i, alias := range remoteMethod.argumentIdentifiers {
		argumentKind := remoteMethod.arguments[i+argumentsDisplacement]

		items[i] = reflect.Zero(argumentKind) //

		arg := obtainArgument(r, alias)
		//

		if arg == "REDACTED" {
			continue
		}

		v, ok, err := unJSNTextTowardArgument(argumentKind, arg)
		if err != nil {
			return nil, err
		}
		if ok {
			items[i] = v
			continue
		}

		items[i], err = jsnTextTowardArgument(argumentKind, arg)
		if err != nil {
			return nil, err
		}
	}

	return items, nil
}

func jsnTextTowardArgument(rt reflect.Type, arg string) (reflect.Value, error) {
	rv := reflect.New(rt)
	err := strongmindjson.Decode([]byte(arg), rv.Interface())
	if err != nil {
		return rv, err
	}
	rv = rv.Elem()
	return rv, nil
}

func unJSNTextTowardArgument(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	if rt.Kind() == reflect.Ptr {
		rv1, ok, err := unJSNTextTowardArgument(rt.Elem(), arg)
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
		return _unjsontexttoarg(rt, arg)
	}
}

//
func _unjsontexttoarg(rt reflect.Type, arg string) (reflect.Value, bool, error) {
	equalsIntegerText := againInteger.Match([]byte(arg))
	equalsCitedText := strings.HasPrefix(arg, "REDACTED") && strings.HasSuffix(arg, "REDACTED")
	equalsHexadecimalText := strings.HasPrefix(strings.ToLower(arg), "REDACTED")

	var awaitingText, awaitingOctetSection, awaitingInteger bool
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
		awaitingText = true
	case reflect.Slice:
		awaitingOctetSection = rt.Elem().Kind() == reflect.Uint8
	}

	if equalsIntegerText && awaitingInteger {
		queryarg := "REDACTED" + arg + "REDACTED"
		rv, err := jsnTextTowardArgument(rt, queryarg)
		if err != nil {
			return rv, false, err
		}

		return rv, true, nil
	}

	if equalsHexadecimalText {
		if !awaitingText && !awaitingOctetSection {
			err := fmt.Errorf("REDACTED",
				rt.Kind().String())
			return reflect.ValueOf(nil), false, err
		}

		var datum []byte
		datum, err := hex.DecodeString(arg[2:])
		if err != nil {
			return reflect.ValueOf(nil), false, err
		}
		if rt.Kind() == reflect.String {
			return reflect.ValueOf(string(datum)), true, nil
		}
		return reflect.ValueOf(datum), true, nil
	}

	if equalsCitedText && awaitingOctetSection {
		v := reflect.New(reflect.TypeOf("REDACTED"))
		err := strongmindjson.Decode([]byte(arg), v.Interface())
		if err != nil {
			return reflect.ValueOf(nil), false, err
		}
		v = v.Elem()
		return reflect.ValueOf([]byte(v.String())), true, nil
	}

	return reflect.ValueOf(nil), false, nil
}

func obtainArgument(r *http.Request, argument string) string {
	s := r.URL.Query().Get(argument)
	if s == "REDACTED" {
		s = r.FormValue(argument)
	}
	return s
}
