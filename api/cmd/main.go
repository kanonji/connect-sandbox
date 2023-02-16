package main

import (
	barv1 "connect-sandbox/api/gen/protobuf_foo/bar/v1"
	"connect-sandbox/api/gen/protobuf_foo/bar/v1/barv1connect"
	bazv1 "connect-sandbox/api/gen/protobuf_foo/baz/v1"
	"connect-sandbox/api/gen/protobuf_foo/baz/v1/bazv1connect"
	context "context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BarHandlers struct{}

// Get implements barv1connect.BarServiceHandler
func (*BarHandlers) Get(ctx context.Context, req *connect_go.Request[barv1.GetRequest]) (*connect_go.Response[barv1.GetResponse], error) {
	log.Println("Bar.Get()")
	res := connect.NewResponse(&barv1.GetResponse{
		Id:    req.Msg.Id,
		Month: 2,
	})
	return res, nil
}

// Invoke implements barv1connect.BarServiceHandler
func (*BarHandlers) Invoke(ctx context.Context, _ *connect_go.Request[emptypb.Empty]) (*connect_go.Response[barv1.InvokeResponse], error) {
	log.Println("Bar.Invoke()")
	res := connect.NewResponse(&barv1.InvokeResponse{
		Id: 1,
		InvokedAt: &timestamppb.Timestamp{
			Seconds: 10,
			Nanos:   1,
		},
	})
	return res, nil
}

type BazHandlers struct{}

// Do implements bazv1connect.BazServiceHandler
func (*BazHandlers) Do(ctx context.Context, _ *connect_go.Request[emptypb.Empty]) (*connect_go.Response[bazv1.DoResponse], error) {
	log.Println("Baz.Do()")
	res := connect.NewResponse(&bazv1.DoResponse{
		Id: 1,
		InvokedAt: &timestamppb.Timestamp{
			Seconds: 100,
			Nanos:   0,
		},
	})
	return res, nil
}

func main() {
	mux := http.NewServeMux()
	mux.Handle(barv1connect.NewBarServiceHandler(&BarHandlers{}))
	mux.Handle(bazv1connect.NewBazServiceHandler(&BazHandlers{}))

	debugDumpRouting(mux)

	addr := "0.0.0.0:8888"
	err := http.ListenAndServe(addr, h2c.NewHandler(mux, &http2.Server{}))
	if err != nil {
		log.Println(err)
	}
}

func debugDumpRouting(mux *http.ServeMux) {
	v := reflect.ValueOf(mux).Elem()
	fmt.Printf("routes: %v\n", v.FieldByName("m"))
}
