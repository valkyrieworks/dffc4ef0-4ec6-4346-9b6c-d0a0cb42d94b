package host

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/valkyrieworks/utils/log"
)

//
//
//
//
func EnrollRPCRoutines(mux *http.ServeMux, functionIndex map[string]*RPCFunction, tracer log.Tracer) {
	//
	for functionLabel, rpcFunction := range functionIndex {
		mux.HandleFunc("REDACTED"+functionLabel, createHTTPManager(rpcFunction, tracer))
	}

	//
	mux.HandleFunc("REDACTED", processCorruptJsonrpcRoutes(createJsonrpcManager(functionIndex, tracer)))
}

type Setting func(*RPCFunction)

//
//
//
//
//
//
func Storable(noRepositoryDefinitionArgs ...string) Setting {
	return func(r *RPCFunction) {
		r.storable = true
		r.noRepositoryDefinitionArgs = make(map[string]any)
		for _, arg := range noRepositoryDefinitionArgs {
			r.noRepositoryDefinitionArgs[arg] = nil
		}
	}
}

//
func Ws() Setting {
	return func(r *RPCFunction) {
		r.ws = true
	}
}

//
type RPCFunction struct {
	f              reflect.Value  //
	args           []reflect.Type //
	yields        []reflect.Type //
	argumentLabels       []string       //
	storable      bool           //
	ws             bool           //
	noRepositoryDefinitionArgs map[string]any //
}

//
//
func NewRPCFunction(f any, args string, options ...Setting) *RPCFunction {
	return newRPCFunction(f, args, options...)
}

//
func NewWsrpcFunction(f any, args string, options ...Setting) *RPCFunction {
	options = append(options, Ws())
	return newRPCFunction(f, args, options...)
}

//
//
func (f *RPCFunction) storableWithArgs(args []reflect.Value) bool {
	if !f.storable {
		return false
	}
	//
	for i := 1; i < len(f.args); i++ {
		//
		argumentLabel := f.argumentLabels[i-1]
		if _, hasStandard := f.noRepositoryDefinitionArgs[argumentLabel]; hasStandard {
			//
			if i >= len(args) {
				return false
			}
			//
			if args[i].IsZero() {
				return false
			}
		}
	}
	return true
}

func newRPCFunction(f any, args string, options ...Setting) *RPCFunction {
	var argumentLabels []string
	if args != "REDACTED" {
		argumentLabels = strings.Split(args, "REDACTED")
	}

	r := &RPCFunction{
		f:        reflect.ValueOf(f),
		args:     functionArgumentKinds(f),
		yields:  functionYieldKinds(f),
		argumentLabels: argumentLabels,
	}

	for _, opt := range options {
		opt(r)
	}

	return r
}

//
func functionArgumentKinds(f any) []reflect.Type {
	t := reflect.TypeOf(f)
	n := t.NumIn()
	classes := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		classes[i] = t.In(i)
	}
	return classes
}

//
func functionYieldKinds(f any) []reflect.Type {
	t := reflect.TypeOf(f)
	n := t.NumOut()
	classes := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		classes[i] = t.Out(i)
	}
	return classes
}

//

//
func unmirrorOutcome(yields []reflect.Value) (any, error) {
	errV := yields[1]
	if errV.Interface() != nil {
		return nil, fmt.Errorf("REDACTED", errV.Interface())
	}
	rv := yields[0]
	//
	//
	rvp := reflect.New(rv.Type())
	rvp.Elem().Set(rv)
	return rvp.Interface(), nil
}
