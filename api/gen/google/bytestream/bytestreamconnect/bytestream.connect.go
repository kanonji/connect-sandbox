// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: google/bytestream/bytestream.proto

package bytestreamconnect

import (
	bytestream "connect-sandbox/api/gen/google/bytestream"
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
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
	// ByteStreamName is the fully-qualified name of the ByteStream service.
	ByteStreamName = "google.bytestream.ByteStream"
)

// ByteStreamClient is a client for the google.bytestream.ByteStream service.
type ByteStreamClient interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(context.Context, *connect_go.Request[bytestream.ReadRequest]) (*connect_go.ServerStreamForClient[bytestream.ReadResponse], error)
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(context.Context) *connect_go.ClientStreamForClient[bytestream.WriteRequest, bytestream.WriteResponse]
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(context.Context, *connect_go.Request[bytestream.QueryWriteStatusRequest]) (*connect_go.Response[bytestream.QueryWriteStatusResponse], error)
}

// NewByteStreamClient constructs a client for the google.bytestream.ByteStream service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewByteStreamClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ByteStreamClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &byteStreamClient{
		read: connect_go.NewClient[bytestream.ReadRequest, bytestream.ReadResponse](
			httpClient,
			baseURL+"/google.bytestream.ByteStream/Read",
			opts...,
		),
		write: connect_go.NewClient[bytestream.WriteRequest, bytestream.WriteResponse](
			httpClient,
			baseURL+"/google.bytestream.ByteStream/Write",
			opts...,
		),
		queryWriteStatus: connect_go.NewClient[bytestream.QueryWriteStatusRequest, bytestream.QueryWriteStatusResponse](
			httpClient,
			baseURL+"/google.bytestream.ByteStream/QueryWriteStatus",
			opts...,
		),
	}
}

// byteStreamClient implements ByteStreamClient.
type byteStreamClient struct {
	read             *connect_go.Client[bytestream.ReadRequest, bytestream.ReadResponse]
	write            *connect_go.Client[bytestream.WriteRequest, bytestream.WriteResponse]
	queryWriteStatus *connect_go.Client[bytestream.QueryWriteStatusRequest, bytestream.QueryWriteStatusResponse]
}

// Read calls google.bytestream.ByteStream.Read.
func (c *byteStreamClient) Read(ctx context.Context, req *connect_go.Request[bytestream.ReadRequest]) (*connect_go.ServerStreamForClient[bytestream.ReadResponse], error) {
	return c.read.CallServerStream(ctx, req)
}

// Write calls google.bytestream.ByteStream.Write.
func (c *byteStreamClient) Write(ctx context.Context) *connect_go.ClientStreamForClient[bytestream.WriteRequest, bytestream.WriteResponse] {
	return c.write.CallClientStream(ctx)
}

// QueryWriteStatus calls google.bytestream.ByteStream.QueryWriteStatus.
func (c *byteStreamClient) QueryWriteStatus(ctx context.Context, req *connect_go.Request[bytestream.QueryWriteStatusRequest]) (*connect_go.Response[bytestream.QueryWriteStatusResponse], error) {
	return c.queryWriteStatus.CallUnary(ctx, req)
}

// ByteStreamHandler is an implementation of the google.bytestream.ByteStream service.
type ByteStreamHandler interface {
	// `Read()` is used to retrieve the contents of a resource as a sequence
	// of bytes. The bytes are returned in a sequence of responses, and the
	// responses are delivered as the results of a server-side streaming RPC.
	Read(context.Context, *connect_go.Request[bytestream.ReadRequest], *connect_go.ServerStream[bytestream.ReadResponse]) error
	// `Write()` is used to send the contents of a resource as a sequence of
	// bytes. The bytes are sent in a sequence of request protos of a client-side
	// streaming RPC.
	//
	// A `Write()` action is resumable. If there is an error or the connection is
	// broken during the `Write()`, the client should check the status of the
	// `Write()` by calling `QueryWriteStatus()` and continue writing from the
	// returned `committed_size`. This may be less than the amount of data the
	// client previously sent.
	//
	// Calling `Write()` on a resource name that was previously written and
	// finalized could cause an error, depending on whether the underlying service
	// allows over-writing of previously written resources.
	//
	// When the client closes the request channel, the service will respond with
	// a `WriteResponse`. The service will not view the resource as `complete`
	// until the client has sent a `WriteRequest` with `finish_write` set to
	// `true`. Sending any requests on a stream after sending a request with
	// `finish_write` set to `true` will cause an error. The client **should**
	// check the `WriteResponse` it receives to determine how much data the
	// service was able to commit and whether the service views the resource as
	// `complete` or not.
	Write(context.Context, *connect_go.ClientStream[bytestream.WriteRequest]) (*connect_go.Response[bytestream.WriteResponse], error)
	// `QueryWriteStatus()` is used to find the `committed_size` for a resource
	// that is being written, which can then be used as the `write_offset` for
	// the next `Write()` call.
	//
	// If the resource does not exist (i.e., the resource has been deleted, or the
	// first `Write()` has not yet reached the service), this method returns the
	// error `NOT_FOUND`.
	//
	// The client **may** call `QueryWriteStatus()` at any time to determine how
	// much data has been processed for this resource. This is useful if the
	// client is buffering data and needs to know which data can be safely
	// evicted. For any sequence of `QueryWriteStatus()` calls for a given
	// resource name, the sequence of returned `committed_size` values will be
	// non-decreasing.
	QueryWriteStatus(context.Context, *connect_go.Request[bytestream.QueryWriteStatusRequest]) (*connect_go.Response[bytestream.QueryWriteStatusResponse], error)
}

// NewByteStreamHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewByteStreamHandler(svc ByteStreamHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/google.bytestream.ByteStream/Read", connect_go.NewServerStreamHandler(
		"/google.bytestream.ByteStream/Read",
		svc.Read,
		opts...,
	))
	mux.Handle("/google.bytestream.ByteStream/Write", connect_go.NewClientStreamHandler(
		"/google.bytestream.ByteStream/Write",
		svc.Write,
		opts...,
	))
	mux.Handle("/google.bytestream.ByteStream/QueryWriteStatus", connect_go.NewUnaryHandler(
		"/google.bytestream.ByteStream/QueryWriteStatus",
		svc.QueryWriteStatus,
		opts...,
	))
	return "/google.bytestream.ByteStream/", mux
}

// UnimplementedByteStreamHandler returns CodeUnimplemented from all methods.
type UnimplementedByteStreamHandler struct{}

func (UnimplementedByteStreamHandler) Read(context.Context, *connect_go.Request[bytestream.ReadRequest], *connect_go.ServerStream[bytestream.ReadResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("google.bytestream.ByteStream.Read is not implemented"))
}

func (UnimplementedByteStreamHandler) Write(context.Context, *connect_go.ClientStream[bytestream.WriteRequest]) (*connect_go.Response[bytestream.WriteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("google.bytestream.ByteStream.Write is not implemented"))
}

func (UnimplementedByteStreamHandler) QueryWriteStatus(context.Context, *connect_go.Request[bytestream.QueryWriteStatusRequest]) (*connect_go.Response[bytestream.QueryWriteStatusResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("google.bytestream.ByteStream.QueryWriteStatus is not implemented"))
}
