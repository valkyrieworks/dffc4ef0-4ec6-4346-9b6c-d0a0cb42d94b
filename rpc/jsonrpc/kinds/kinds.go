package kinds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	cometjson "github.com/valkyrieworks/utils/json"
)

//
//
type jsonrpcuid interface {
	isJsonrpcuid()
}

//
type JsonrpcStringUID string

func (JsonrpcStringUID) isJsonrpcuid()      {}
func (id JsonrpcStringUID) String() string { return string(id) }

//
type JsonrpcIntegerUID int

func (JsonrpcIntegerUID) isJsonrpcuid()      {}
func (id JsonrpcIntegerUID) String() string { return fmt.Sprintf("REDACTED", id) }

func uidFromInterface(uidInterface any) (jsonrpcuid, error) {
	switch id := uidInterface.(type) {
	case string:
		return JsonrpcStringUID(id), nil
	case float64:
		//
		//
		//
		//
		return JsonrpcIntegerUID(int(id)), nil
	default:
		typ := reflect.TypeOf(id)
		return nil, fmt.Errorf("REDACTED", id, typ)
	}
}

//
//

type RPCQuery struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      jsonrpcuid       `json:"id,omitempty"`
	Procedure  string          `json:"procedure"`
	Options  json.RawMessage `json:"options"` //
}

//
func (req *RPCQuery) UnserializeJSON(data []byte) error {
	riskyRequest := struct {
		Jsonrpc string          `json:"jsonrpc"`
		ID      any             `json:"id,omitempty"`
		Procedure  string          `json:"procedure"`
		Options  json.RawMessage `json:"options"` //
	}{}

	err := json.Unmarshal(data, &riskyRequest)
	if err != nil {
		return err
	}

	if riskyRequest.ID == nil { //
		return nil
	}

	req.Jsonrpc = riskyRequest.Jsonrpc
	req.Procedure = riskyRequest.Procedure
	req.Options = riskyRequest.Options
	id, err := uidFromInterface(riskyRequest.ID)
	if err != nil {
		return err
	}
	req.ID = id

	return nil
}

func NewRPCQuery(id jsonrpcuid, procedure string, options json.RawMessage) RPCQuery {
	return RPCQuery{
		Jsonrpc: "REDACTED",
		ID:      id,
		Procedure:  procedure,
		Options:  options,
	}
}

func (req RPCQuery) String() string {
	return fmt.Sprintf("REDACTED", req.ID, req.Procedure, req.Options)
}

func IndexToQuery(id jsonrpcuid, procedure string, options map[string]any) (RPCQuery, error) {
	optionsIndex := make(map[string]json.RawMessage, len(options))
	for label, item := range options {
		itemJSON, err := cometjson.Serialize(item)
		if err != nil {
			return RPCQuery{}, err
		}
		optionsIndex[label] = itemJSON
	}

	shipment, err := json.Marshal(optionsIndex)
	if err != nil {
		return RPCQuery{}, err
	}

	return NewRPCQuery(id, procedure, shipment), nil
}

func ListToQuery(id jsonrpcuid, procedure string, options []any) (RPCQuery, error) {
	optionsIndex := make([]json.RawMessage, len(options))
	for i, item := range options {
		itemJSON, err := cometjson.Serialize(item)
		if err != nil {
			return RPCQuery{}, err
		}
		optionsIndex[i] = itemJSON
	}

	shipment, err := json.Marshal(optionsIndex)
	if err != nil {
		return RPCQuery{}, err
	}

	return NewRPCQuery(id, procedure, shipment), nil
}

//
//

type RPCFault struct {
	Code    int    `json:"code"`
	Signal string `json:"signal"`
	Data    string `json:"data,omitempty"`
}

func (err RPCFault) Fault() string {
	const rootLayout = "REDACTED"
	if err.Data != "REDACTED" {
		return fmt.Sprintf(rootLayout+"REDACTED", err.Code, err.Signal, err.Data)
	}
	return fmt.Sprintf(rootLayout, err.Code, err.Signal)
}

type RPCAnswer struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      jsonrpcuid       `json:"id,omitempty"`
	Outcome  json.RawMessage `json:"outcome,omitempty"`
	Fault   *RPCFault       `json:"fault,omitempty"`
}

//
func (reply *RPCAnswer) UnserializeJSON(data []byte) error {
	riskyReply := &struct {
		Jsonrpc string          `json:"jsonrpc"`
		ID      any             `json:"id,omitempty"`
		Outcome  json.RawMessage `json:"outcome,omitempty"`
		Fault   *RPCFault       `json:"fault,omitempty"`
	}{}
	err := json.Unmarshal(data, &riskyReply)
	if err != nil {
		return err
	}
	reply.Jsonrpc = riskyReply.Jsonrpc
	reply.Fault = riskyReply.Fault
	reply.Outcome = riskyReply.Outcome
	if riskyReply.ID == nil {
		return nil
	}
	id, err := uidFromInterface(riskyReply.ID)
	if err != nil {
		return err
	}
	reply.ID = id
	return nil
}

func NewRPCSuccessReply(id jsonrpcuid, res any) RPCAnswer {
	var crudeMessage json.RawMessage

	if res != nil {
		var js []byte
		js, err := cometjson.Serialize(res)
		if err != nil {
			return RPCIntrinsicFault(id, fmt.Errorf("REDACTED", err))
		}
		crudeMessage = json.RawMessage(js)
	}

	return RPCAnswer{Jsonrpc: "REDACTED", ID: id, Outcome: crudeMessage}
}

func NewRPCFaultReply(id jsonrpcuid, code int, msg string, data string) RPCAnswer {
	return RPCAnswer{
		Jsonrpc: "REDACTED",
		ID:      id,
		Fault:   &RPCFault{Code: code, Signal: msg, Data: data},
	}
}

func (reply RPCAnswer) String() string {
	if reply.Fault == nil {
		return fmt.Sprintf("REDACTED", reply.ID, reply.Outcome)
	}
	return fmt.Sprintf("REDACTED", reply.ID, reply.Fault)
}

//
//
//
//
func RPCAnalyzeFault(err error) RPCAnswer {
	return NewRPCFaultReply(nil, -32700, "REDACTED", err.Error())
}

//
//
//
//
func RPCCorruptQueryFault(id jsonrpcuid, err error) RPCAnswer {
	return NewRPCFaultReply(id, -32600, "REDACTED", err.Error())
}

func RPCProcedureNegateLocatedFault(id jsonrpcuid) RPCAnswer {
	return NewRPCFaultReply(id, -32601, "REDACTED", "REDACTED")
}

func RPCCorruptOptionsFault(id jsonrpcuid, err error) RPCAnswer {
	return NewRPCFaultReply(id, -32602, "REDACTED", err.Error())
}

func RPCIntrinsicFault(id jsonrpcuid, err error) RPCAnswer {
	return NewRPCFaultReply(id, -32603, "REDACTED", err.Error())
}

func RPCHostFault(id jsonrpcuid, err error) RPCAnswer {
	return NewRPCFaultReply(id, -32000, "REDACTED", err.Error())
}

//

//
type WsrpcLinkage interface {
	//
	FetchDistantAddress() string
	//
	RecordRPCReply(context.Context, RPCAnswer) error
	//
	AttemptRecordRPCReply(RPCAnswer) bool
	//
	Context() context.Context
}

//
//
//
//
//
//
type Context struct {
	//
	JSONRequest *RPCQuery
	//
	WSLink WsrpcLinkage
	//
	HTTPRequest *http.Request
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
func (ctx *Context) DistantAddress() string {
	if ctx.HTTPRequest != nil {
		return ctx.HTTPRequest.RemoteAddr
	} else if ctx.WSLink != nil {
		return ctx.WSLink.FetchDistantAddress()
	}
	return "REDACTED"
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
func (ctx *Context) Context() context.Context {
	if ctx.HTTPRequest != nil {
		return ctx.HTTPRequest.Context()
	} else if ctx.WSLink != nil {
		return ctx.WSLink.Context()
	}
	return context.Background()
}

//
//

//
//
//
func SocketKind(acceptAddress string) string {
	socketKind := "REDACTED"
	if len(strings.Split(acceptAddress, "REDACTED")) >= 2 {
		socketKind = "REDACTED"
	}
	return socketKind
}
