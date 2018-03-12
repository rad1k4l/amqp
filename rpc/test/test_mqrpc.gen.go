// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test.proto

/*
	Package test is a generated protocol buffer package.

	It is generated from these files:
		test.proto

	It has these top-level messages:
		Empty
		SimpleTypes
*/
package test

import fmt "fmt"
import errors "errors"
import rpc "github.com/gavrilaf/amqp/rpc"
import proto "github.com/gogo/protobuf/proto"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Server API
type TestServer interface {
	CopySimple(arg *SimpleTypes) (*SimpleTypes, error)
	GenErr(arg *Empty) (*Empty, error)
}

// Run server API with this call
func RunServer(srv rpc.Server, handler TestServer) {
	srv.Serve(func(funcID int32, arg []byte) ([]byte, error) {
		switch funcID {
		case Functions_CopySimple:
			return _Handle_CopySimple(handler, arg)
		case Functions_GenErr:
			return _Handle_GenErr(handler, arg)
		default:
			return nil, errors.New(fmt.Sprintf("unknown function with code: %d", funcID))
		}
	})
}

// Client API
type TestClient interface {
	Close()
	CopySimple(arg *SimpleTypes) (*SimpleTypes, error)
	GenErr(arg *Empty) (*Empty, error)
}

func NewTestClient(cc rpc.Client) TestClient {
	return &testClient{cc}
}

type testClient struct {
	cc rpc.Client
}

// Functions enum
const (
	Functions_CopySimple int32 = 0
	Functions_GenErr     int32 = 1
)

// Server API handlers
func _Handle_CopySimple(handler interface{}, arg []byte) ([]byte, error) {
	var req SimpleTypes
	err := req.Unmarshal(arg)
	if err != nil {
		return nil, err
	}
	resp, err := handler.(TestServer).CopySimple(&req)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}
func _Handle_GenErr(handler interface{}, arg []byte) ([]byte, error) {
	var req Empty
	err := req.Unmarshal(arg)
	if err != nil {
		return nil, err
	}
	resp, err := handler.(TestServer).GenErr(&req)
	if err != nil {
		return nil, err
	}
	return resp.Marshal()
}

// Client API handlers
func (this *testClient) Close() {
	this.cc.Close()
}
func (this *testClient) CopySimple(arg *SimpleTypes) (*SimpleTypes, error) {
	request, err := arg.Marshal()
	if err != nil {
		return nil, err
	}
	respData, err := this.cc.RemoteCall(rpc.Request{FuncID: Functions_CopySimple, Body: request})
	if err != nil {
		return nil, err
	}
	var resp SimpleTypes
	err = resp.Unmarshal(respData)
	return &resp, err
}
func (this *testClient) GenErr(arg *Empty) (*Empty, error) {
	request, err := arg.Marshal()
	if err != nil {
		return nil, err
	}
	respData, err := this.cc.RemoteCall(rpc.Request{FuncID: Functions_GenErr, Body: request})
	if err != nil {
		return nil, err
	}
	var resp Empty
	err = resp.Unmarshal(respData)
	return &resp, err
}
