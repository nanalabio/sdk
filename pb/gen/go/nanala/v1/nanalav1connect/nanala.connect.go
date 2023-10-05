// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: nanala/v1/nanala.proto

package nanalav1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/nanalabio/sdk/pb/gen/go/nanala/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// NanalaServiceName is the fully-qualified name of the NanalaService service.
	NanalaServiceName = "nanala.v1.NanalaService"
)

// NanalaServiceClient is a client for the nanala.v1.NanalaService service.
type NanalaServiceClient interface {
	Synthesize(context.Context, *connect_go.Request[v1.SynthesizeRequest]) (*connect_go.Response[v1.SynthesizeResponse], error)
}

// NewNanalaServiceClient constructs a client for the nanala.v1.NanalaService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewNanalaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) NanalaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &nanalaServiceClient{
		synthesize: connect_go.NewClient[v1.SynthesizeRequest, v1.SynthesizeResponse](
			httpClient,
			baseURL+"/nanala.v1.NanalaService/Synthesize",
			opts...,
		),
	}
}

// nanalaServiceClient implements NanalaServiceClient.
type nanalaServiceClient struct {
	synthesize *connect_go.Client[v1.SynthesizeRequest, v1.SynthesizeResponse]
}

// Synthesize calls nanala.v1.NanalaService.Synthesize.
func (c *nanalaServiceClient) Synthesize(ctx context.Context, req *connect_go.Request[v1.SynthesizeRequest]) (*connect_go.Response[v1.SynthesizeResponse], error) {
	return c.synthesize.CallUnary(ctx, req)
}

// NanalaServiceHandler is an implementation of the nanala.v1.NanalaService service.
type NanalaServiceHandler interface {
	Synthesize(context.Context, *connect_go.Request[v1.SynthesizeRequest]) (*connect_go.Response[v1.SynthesizeResponse], error)
}

// NewNanalaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewNanalaServiceHandler(svc NanalaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/nanala.v1.NanalaService/Synthesize", connect_go.NewUnaryHandler(
		"/nanala.v1.NanalaService/Synthesize",
		svc.Synthesize,
		opts...,
	))
	return "/nanala.v1.NanalaService/", mux
}

// UnimplementedNanalaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedNanalaServiceHandler struct{}

func (UnimplementedNanalaServiceHandler) Synthesize(context.Context, *connect_go.Request[v1.SynthesizeRequest]) (*connect_go.Response[v1.SynthesizeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("nanala.v1.NanalaService.Synthesize is not implemented"))
}
