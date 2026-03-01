package node

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

//
//
//
//
func EnrollRemoteRoutines(mux *http.ServeMux, methodIndex map[string]*RemoteMethod, tracer log.Tracer) {
	//
	for methodAlias, remoteMethod := range methodIndex {
		mux.HandleFunc("REDACTED"+methodAlias, createHttpsvcProcessor(remoteMethod, tracer))
	}

	//
	mux.HandleFunc("REDACTED", processUnfitJsonifaceRoutes(createJsonifaceProcessor(methodIndex, tracer)))
}

type Selection func(*RemoteMethod)

//
//
//
//
//
//
func Storable(negativeStashDefinitionArguments ...string) Selection {
	return func(r *RemoteMethod) {
		r.storable = true
		r.negativeStashDefinitionArguments = make(map[string]any)
		for _, arg := range negativeStashDefinitionArguments {
			r.negativeStashDefinitionArguments[arg] = nil
		}
	}
}

//
func Ws() Selection {
	return func(r *RemoteMethod) {
		r.ws = true
	}
}

//
type RemoteMethod struct {
	f              reflect.Value  //
	arguments           []reflect.Type //
	yields        []reflect.Type //
	argumentIdentifiers       []string       //
	storable      bool           //
	ws             bool           //
	negativeStashDefinitionArguments map[string]any //
}

//
//
func FreshRemoteMethod(f any, arguments string, choices ...Selection) *RemoteMethod {
	return freshRemoteMethod(f, arguments, choices...)
}

//
func FreshSocketifaceMethod(f any, arguments string, choices ...Selection) *RemoteMethod {
	choices = append(choices, Ws())
	return freshRemoteMethod(f, arguments, choices...)
}

//
//
func (f *RemoteMethod) storableUsingArguments(arguments []reflect.Value) bool {
	if !f.storable {
		return false
	}
	//
	for i := 1; i < len(f.arguments); i++ {
		//
		argumentAlias := f.argumentIdentifiers[i-1]
		if _, ownsFallback := f.negativeStashDefinitionArguments[argumentAlias]; ownsFallback {
			//
			if i >= len(arguments) {
				return false
			}
			//
			if arguments[i].IsZero() {
				return false
			}
		}
	}
	return true
}

func freshRemoteMethod(f any, arguments string, choices ...Selection) *RemoteMethod {
	var argumentIdentifiers []string
	if arguments != "REDACTED" {
		argumentIdentifiers = strings.Split(arguments, "REDACTED")
	}

	r := &RemoteMethod{
		f:        reflect.ValueOf(f),
		arguments:     methodArgumentKinds(f),
		yields:  methodYieldKinds(f),
		argumentIdentifiers: argumentIdentifiers,
	}

	for _, opt := range choices {
		opt(r)
	}

	return r
}

//
func methodArgumentKinds(f any) []reflect.Type {
	t := reflect.TypeOf(f)
	n := t.NumIn()
	varieties := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		varieties[i] = t.In(i)
	}
	return varieties
}

//
func methodYieldKinds(f any) []reflect.Type {
	t := reflect.TypeOf(f)
	n := t.NumOut()
	varieties := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		varieties[i] = t.Out(i)
	}
	return varieties
}

//

//
func unmirrorOutcome(yields []reflect.Value) (any, error) {
	faultVER := yields[1]
	if faultVER.Interface() != nil {
		return nil, fmt.Errorf("REDACTED", faultVER.Interface())
	}
	rv := yields[0]
	//
	//
	rvp := reflect.New(rv.Type())
	rvp.Elem().Set(rv)
	return rvp.Interface(), nil
}
