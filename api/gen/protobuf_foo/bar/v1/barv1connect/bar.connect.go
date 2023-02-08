// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: protobuf_foo/bar/v1/bar.proto

package barv1connect

import (
	v1 "connect-sandbox/api/gen/protobuf_foo/bar/v1"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// BarServiceName is the fully-qualified name of the BarService service.
	BarServiceName = "protobuf_foo.bar.v1.BarService"
)

// BarServiceClient is a client for the protobuf_foo.bar.v1.BarService service.
type BarServiceClient interface {
	Invoke(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.InvokeResponse], error)
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
}

// NewBarServiceClient constructs a client for the protobuf_foo.bar.v1.BarService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBarServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BarServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &barServiceClient{
		invoke: connect_go.NewClient[emptypb.Empty, v1.InvokeResponse](
			httpClient,
			baseURL+"/protobuf_foo.bar.v1.BarService/Invoke",
			opts...,
		),
		get: connect_go.NewClient[v1.GetRequest, v1.GetResponse](
			httpClient,
			baseURL+"/protobuf_foo.bar.v1.BarService/Get",
			opts...,
		),
	}
}

// barServiceClient implements BarServiceClient.
type barServiceClient struct {
	invoke *connect_go.Client[emptypb.Empty, v1.InvokeResponse]
	get    *connect_go.Client[v1.GetRequest, v1.GetResponse]
}

// Invoke calls protobuf_foo.bar.v1.BarService.Invoke.
func (c *barServiceClient) Invoke(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.InvokeResponse], error) {
	return c.invoke.CallUnary(ctx, req)
}

// Get calls protobuf_foo.bar.v1.BarService.Get.
func (c *barServiceClient) Get(ctx context.Context, req *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// BarServiceHandler is an implementation of the protobuf_foo.bar.v1.BarService service.
type BarServiceHandler interface {
	Invoke(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.InvokeResponse], error)
	Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error)
}

// NewBarServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBarServiceHandler(svc BarServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/protobuf_foo.bar.v1.BarService/Invoke", connect_go.NewUnaryHandler(
		"/protobuf_foo.bar.v1.BarService/Invoke",
		svc.Invoke,
		opts...,
	))
	mux.Handle("/protobuf_foo.bar.v1.BarService/Get", connect_go.NewUnaryHandler(
		"/protobuf_foo.bar.v1.BarService/Get",
		svc.Get,
		opts...,
	))
	return "/protobuf_foo.bar.v1.BarService/", mux
}

// UnimplementedBarServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBarServiceHandler struct{}

func (UnimplementedBarServiceHandler) Invoke(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[v1.InvokeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("protobuf_foo.bar.v1.BarService.Invoke is not implemented"))
}

func (UnimplementedBarServiceHandler) Get(context.Context, *connect_go.Request[v1.GetRequest]) (*connect_go.Response[v1.GetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("protobuf_foo.bar.v1.BarService.Get is not implemented"))
}
