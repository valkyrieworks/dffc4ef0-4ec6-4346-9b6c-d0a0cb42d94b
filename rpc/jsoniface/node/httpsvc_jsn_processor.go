package node

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

//

//
func createJsonifaceProcessor(methodIndex map[string]*RemoteMethod, tracer log.Tracer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			res := kinds.RemoteUnfitSolicitFailure(nil,
				fmt.Errorf("REDACTED", err),
			)
			if wrFault := RecordRemoteReplyHttpsvcFailure(w, http.StatusBadRequest, res); wrFault != nil {
				tracer.Failure("REDACTED", "REDACTED", wrFault)
			}
			return
		}

		//
		//
		if len(b) == 0 {
			recordCatalogBelongingTerminals(w, r, methodIndex)
			return
		}

		//
		var (
			solicits  []kinds.RemoteSolicit
			replies []kinds.RemoteReply
		)
		if err := json.Unmarshal(b, &solicits); err != nil {
			//
			var solicit kinds.RemoteSolicit
			if err := json.Unmarshal(b, &solicit); err != nil {
				res := kinds.RemoteAnalyzeFailure(fmt.Errorf("REDACTED", err))
				if wrFault := RecordRemoteReplyHttpsvcFailure(w, http.StatusInternalServerError, res); wrFault != nil {
					tracer.Failure("REDACTED", "REDACTED", wrFault)
				}
				return
			}
			solicits = []kinds.RemoteSolicit{solicit}
		}

		//
		//
		//
		//
		stash := true
		for _, solicit := range solicits {

			//
			//
			if solicit.ID == nil {
				tracer.Diagnose(
					"REDACTED",
					"REDACTED", solicit,
				)
				continue
			}
			if len(r.URL.Path) > 1 {
				replies = append(
					replies,
					kinds.RemoteUnfitSolicitFailure(solicit.ID, fmt.Errorf("REDACTED", r.URL.Path)),
				)
				stash = false
				continue
			}
			remoteMethod, ok := methodIndex[solicit.Procedure]
			if !ok || (remoteMethod.ws) {
				replies = append(replies, kinds.RemoteProcedureNegationDetectedFailure(solicit.ID))
				stash = false
				continue
			}
			ctx := &kinds.Env{JSNRequest: &solicit, HttpsvcRequest: r}
			arguments := []reflect.Value{reflect.ValueOf(ctx)}
			if len(solicit.Parameters) > 0 {
				procArguments, err := jsnParametersTowardArguments(remoteMethod, solicit.Parameters)
				if err != nil {
					replies = append(
						replies,
						kinds.RemoteUnfitParametersFailure(solicit.ID, fmt.Errorf("REDACTED", err)),
					)
					stash = false
					continue
				}
				arguments = append(arguments, procArguments...)
			}

			if stash && !remoteMethod.storableUsingArguments(arguments) {
				stash = false
			}

			yields := remoteMethod.f.Call(arguments)
			outcome, err := unmirrorOutcome(yields)
			if err != nil {
				replies = append(replies, kinds.RemoteIntrinsicFailure(solicit.ID, err))
				continue
			}
			replies = append(replies, kinds.FreshRemoteTriumphReply(solicit.ID, outcome))
		}

		if len(replies) > 0 {
			var wrFault error
			if stash {
				wrFault = RecordStorableRemoteReplyHttpsvc(w, replies...)
			} else {
				wrFault = RecordRemoteReplyHttpsvc(w, replies...)
			}
			if wrFault != nil {
				tracer.Failure("REDACTED", "REDACTED", wrFault)
			}
		}
	}
}

func processUnfitJsonifaceRoutes(following http.HandlerFunc) http.HandlerFunc {
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

func indexParametersTowardArguments(
	remoteMethod *RemoteMethod,
	parameters map[string]json.RawMessage,
	argumentsDisplacement int,
) ([]reflect.Value, error) {
	items := make([]reflect.Value, len(remoteMethod.argumentIdentifiers))
	for i, argumentAlias := range remoteMethod.argumentIdentifiers {
		argumentKind := remoteMethod.arguments[i+argumentsDisplacement]

		if p, ok := parameters[argumentAlias]; ok && p != nil && len(p) > 0 {
			val := reflect.New(argumentKind)
			err := strongmindjson.Decode(p, val.Interface())
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

func seriesParametersTowardArguments(
	remoteMethod *RemoteMethod,
	parameters []json.RawMessage,
	argumentsDisplacement int,
) ([]reflect.Value, error) {
	if len(remoteMethod.argumentIdentifiers) != len(parameters) {
		return nil, fmt.Errorf("REDACTED",
			len(remoteMethod.argumentIdentifiers), remoteMethod.argumentIdentifiers, len(parameters), parameters)
	}

	items := make([]reflect.Value, len(parameters))
	for i, p := range parameters {
		argumentKind := remoteMethod.arguments[i+argumentsDisplacement]
		val := reflect.New(argumentKind)
		err := strongmindjson.Decode(p, val.Interface())
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
func jsnParametersTowardArguments(remoteMethod *RemoteMethod, raw []byte) ([]reflect.Value, error) {
	const argumentsDisplacement = 1

	//
	//
	var m map[string]json.RawMessage
	err := json.Unmarshal(raw, &m)
	if err == nil {
		return indexParametersTowardArguments(remoteMethod, m, argumentsDisplacement)
	}

	//
	var a []json.RawMessage
	err = json.Unmarshal(raw, &a)
	if err == nil {
		return seriesParametersTowardArguments(remoteMethod, a, argumentsDisplacement)
	}

	//
	return nil, fmt.Errorf("REDACTED", err)
}

//
func recordCatalogBelongingTerminals(w http.ResponseWriter, r *http.Request, methodIndex map[string]*RemoteMethod) {
	negativeArgumentIdentifiers := []string{}
	argumentIdentifiers := []string{}
	for alias, methodData := range methodIndex {
		if len(methodData.arguments) == 0 {
			negativeArgumentIdentifiers = append(negativeArgumentIdentifiers, alias)
		} else {
			argumentIdentifiers = append(argumentIdentifiers, alias)
		}
	}
	sort.Strings(negativeArgumentIdentifiers)
	sort.Strings(argumentIdentifiers)
	buf := new(bytes.Buffer)
	buf.WriteString("REDACTED")
	buf.WriteString("REDACTED")

	for _, alias := range negativeArgumentIdentifiers {
		connection := fmt.Sprintf("REDACTED", r.Host, alias)
		fmt.Fprintf(buf, "REDACTED", connection, connection)
	}

	buf.WriteString("REDACTED")
	for _, alias := range argumentIdentifiers {
		connection := fmt.Sprintf("REDACTED", r.Host, alias)
		methodData := methodIndex[alias]
		for i, argumentAlias := range methodData.argumentIdentifiers {
			connection += argumentAlias + "REDACTED"
			if i < len(methodData.argumentIdentifiers)-1 {
				connection += "REDACTED"
			}
		}
		fmt.Fprintf(buf, "REDACTED", connection, connection)
	}
	buf.WriteString("REDACTED")
	w.Header().Set("REDACTED", "REDACTED")
	w.WriteHeader(200)
	w.Write(buf.Bytes()) //
}
