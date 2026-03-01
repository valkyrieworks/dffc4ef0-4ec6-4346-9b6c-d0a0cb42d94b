package kinds

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	strongmindjson "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/jsn"
)

//
//
type jsonrpcuuid interface {
	equalsJsonrpcuuid()
}

//
type JsonifaceTextUUID string

func (JsonifaceTextUUID) equalsJsonrpcuuid()      {}
func (id JsonifaceTextUUID) Text() string { return string(id) }

//
type JsonifaceIntegerUUID int

func (JsonifaceIntegerUUID) equalsJsonrpcuuid()      {}
func (id JsonifaceIntegerUUID) Text() string { return fmt.Sprintf("REDACTED", id) }

func uuidOriginatingContract(uuidContract any) (jsonrpcuuid, error) {
	switch id := uuidContract.(type) {
	case string:
		return JsonifaceTextUUID(id), nil
	case float64:
		//
		//
		//
		//
		return JsonifaceIntegerUUID(int(id)), nil
	default:
		typ := reflect.TypeOf(id)
		return nil, fmt.Errorf("REDACTED", id, typ)
	}
}

//
//

type RemoteSolicit struct {
	Jsoniface string          `json:"jsoniface"`
	ID      jsonrpcuuid       `json:"id,omitempty"`
	Procedure  string          `json:"procedure"`
	Parameters  json.RawMessage `json:"parameters"` //
}

//
func (req *RemoteSolicit) DecodeJSN(data []byte) error {
	insecureRequest := struct {
		Jsoniface string          `json:"jsoniface"`
		ID      any             `json:"id,omitempty"`
		Procedure  string          `json:"procedure"`
		Parameters  json.RawMessage `json:"parameters"` //
	}{}

	err := json.Unmarshal(data, &insecureRequest)
	if err != nil {
		return err
	}

	if insecureRequest.ID == nil { //
		return nil
	}

	req.Jsoniface = insecureRequest.Jsoniface
	req.Procedure = insecureRequest.Procedure
	req.Parameters = insecureRequest.Parameters
	id, err := uuidOriginatingContract(insecureRequest.ID)
	if err != nil {
		return err
	}
	req.ID = id

	return nil
}

func FreshRemoteSolicit(id jsonrpcuuid, procedure string, parameters json.RawMessage) RemoteSolicit {
	return RemoteSolicit{
		Jsoniface: "REDACTED",
		ID:      id,
		Procedure:  procedure,
		Parameters:  parameters,
	}
}

func (req RemoteSolicit) Text() string {
	return fmt.Sprintf("REDACTED", req.ID, req.Procedure, req.Parameters)
}

func IndexTowardSolicit(id jsonrpcuuid, procedure string, parameters map[string]any) (RemoteSolicit, error) {
	parametersIndex := make(map[string]json.RawMessage, len(parameters))
	for alias, datum := range parameters {
		datumJSN, err := strongmindjson.Serialize(datum)
		if err != nil {
			return RemoteSolicit{}, err
		}
		parametersIndex[alias] = datumJSN
	}

	content, err := json.Marshal(parametersIndex)
	if err != nil {
		return RemoteSolicit{}, err
	}

	return FreshRemoteSolicit(id, procedure, content), nil
}

func SeriesTowardSolicit(id jsonrpcuuid, procedure string, parameters []any) (RemoteSolicit, error) {
	parametersIndex := make([]json.RawMessage, len(parameters))
	for i, datum := range parameters {
		datumJSN, err := strongmindjson.Serialize(datum)
		if err != nil {
			return RemoteSolicit{}, err
		}
		parametersIndex[i] = datumJSN
	}

	content, err := json.Marshal(parametersIndex)
	if err != nil {
		return RemoteSolicit{}, err
	}

	return FreshRemoteSolicit(id, procedure, content), nil
}

//
//

type RemoteFailure struct {
	Cipher    int    `json:"cipher"`
	Signal string `json:"signal"`
	Data    string `json:"data,omitempty"`
}

func (err RemoteFailure) Failure() string {
	const foundationLayout = "REDACTED"
	if err.Data != "REDACTED" {
		return fmt.Sprintf(foundationLayout+"REDACTED", err.Cipher, err.Signal, err.Data)
	}
	return fmt.Sprintf(foundationLayout, err.Cipher, err.Signal)
}

type RemoteReply struct {
	Jsoniface string          `json:"jsoniface"`
	ID      jsonrpcuuid       `json:"id,omitempty"`
	Outcome  json.RawMessage `json:"outcome,omitempty"`
	Failure   *RemoteFailure       `json:"failure,omitempty"`
}

//
func (reply *RemoteReply) DecodeJSN(data []byte) error {
	insecureAnswer := &struct {
		Jsoniface string          `json:"jsoniface"`
		ID      any             `json:"id,omitempty"`
		Outcome  json.RawMessage `json:"outcome,omitempty"`
		Failure   *RemoteFailure       `json:"failure,omitempty"`
	}{}
	err := json.Unmarshal(data, &insecureAnswer)
	if err != nil {
		return err
	}
	reply.Jsoniface = insecureAnswer.Jsoniface
	reply.Failure = insecureAnswer.Failure
	reply.Outcome = insecureAnswer.Outcome
	if insecureAnswer.ID == nil {
		return nil
	}
	id, err := uuidOriginatingContract(insecureAnswer.ID)
	if err != nil {
		return err
	}
	reply.ID = id
	return nil
}

func FreshRemoteTriumphReply(id jsonrpcuuid, res any) RemoteReply {
	var crudeSignal json.RawMessage

	if res != nil {
		var js []byte
		js, err := strongmindjson.Serialize(res)
		if err != nil {
			return RemoteIntrinsicFailure(id, fmt.Errorf("REDACTED", err))
		}
		crudeSignal = json.RawMessage(js)
	}

	return RemoteReply{Jsoniface: "REDACTED", ID: id, Outcome: crudeSignal}
}

func FreshRemoteFailureReply(id jsonrpcuuid, cipher int, msg string, data string) RemoteReply {
	return RemoteReply{
		Jsoniface: "REDACTED",
		ID:      id,
		Failure:   &RemoteFailure{Cipher: cipher, Signal: msg, Data: data},
	}
}

func (reply RemoteReply) Text() string {
	if reply.Failure == nil {
		return fmt.Sprintf("REDACTED", reply.ID, reply.Outcome)
	}
	return fmt.Sprintf("REDACTED", reply.ID, reply.Failure)
}

//
//
//
//
func RemoteAnalyzeFailure(err error) RemoteReply {
	return FreshRemoteFailureReply(nil, -32700, "REDACTED", err.Error())
}

//
//
//
//
func RemoteUnfitSolicitFailure(id jsonrpcuuid, err error) RemoteReply {
	return FreshRemoteFailureReply(id, -32600, "REDACTED", err.Error())
}

func RemoteProcedureNegationDetectedFailure(id jsonrpcuuid) RemoteReply {
	return FreshRemoteFailureReply(id, -32601, "REDACTED", "REDACTED")
}

func RemoteUnfitParametersFailure(id jsonrpcuuid, err error) RemoteReply {
	return FreshRemoteFailureReply(id, -32602, "REDACTED", err.Error())
}

func RemoteIntrinsicFailure(id jsonrpcuuid, err error) RemoteReply {
	return FreshRemoteFailureReply(id, -32603, "REDACTED", err.Error())
}

func RemoteDaemonFailure(id jsonrpcuuid, err error) RemoteReply {
	return FreshRemoteFailureReply(id, -32000, "REDACTED", err.Error())
}

//

//
type SocketifaceLinkage interface {
	//
	ObtainDistantLocation() string
	//
	PersistRemoteReply(context.Context, RemoteReply) error
	//
	AttemptPersistRemoteReply(RemoteReply) bool
	//
	Env() context.Context
}

//
//
//
//
//
//
type Env struct {
	//
	JSNRequest *RemoteSolicit
	//
	SocketLink SocketifaceLinkage
	//
	HttpsvcRequest *http.Request
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
func (ctx *Env) DistantLocation() string {
	if ctx.HttpsvcRequest != nil {
		return ctx.HttpsvcRequest.RemoteAddr
	} else if ctx.SocketLink != nil {
		return ctx.SocketLink.ObtainDistantLocation()
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
func (ctx *Env) Env() context.Context {
	if ctx.HttpsvcRequest != nil {
		return ctx.HttpsvcRequest.Context()
	} else if ctx.SocketLink != nil {
		return ctx.SocketLink.Env()
	}
	return context.Background()
}

//
//

//
//
//
func PortKind(overhearLocation string) string {
	portKind := "REDACTED"
	if len(strings.Split(overhearLocation, "REDACTED")) >= 2 {
		portKind = "REDACTED"
	}
	return portKind
}
