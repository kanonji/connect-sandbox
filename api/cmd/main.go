package main

import (
	barv1 "connect-sandbox/api/gen/protobuf/bar/v1"
	"connect-sandbox/api/gen/protobuf/bar/v1/barv1connect"
	bazv1 "connect-sandbox/api/gen/protobuf/baz/v1"
	"connect-sandbox/api/gen/protobuf/baz/v1/bazv1connect"
	context "context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	connect_go "github.com/bufbuild/connect-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BarHandlers struct{}

// Get implements barv1connect.BarServiceHandler
func (*BarHandlers) Get(ctx context.Context, req *connect_go.Request[barv1.GetRequest]) (*connect_go.Response[barv1.GetResponse], error) {
	log.Println("Bar.Get()")
	err := req.Msg.ValidateAll()
	var validationErrors barv1.GetRequestMultiError
	if errors.As(err, &validationErrors) {
		cerr := connect.NewError(connect.CodeInvalidArgument, err)
		badRequest := &errdetails.BadRequest{}
		for _, err := range validationErrors.AllErrors() {
			log.Println(err)
			var verr barv1.GetRequestValidationError
			if errors.As(err, &verr) {
				badRequest.FieldViolations = append(badRequest.FieldViolations, &errdetails.BadRequest_FieldViolation{
					Field:       verr.Field(),
					Description: verr.Reason(),
				})
			} else {
				panic(err)
			}
		}
		errorDetail, errInErr := connect.NewErrorDetail(badRequest)
		if errInErr != nil {
			panic(errInErr)
		}
		cerr.AddDetail(errorDetail)
		return nil, cerr
	}
	if err != nil {
		log.Println(err)
	}
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

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Mount(barv1connect.NewBarServiceHandler(&BarHandlers{}))
	})
	r.Mount(bazv1connect.NewBazServiceHandler(&BazHandlers{}))
	debugDumpRoutingChi(r)

	addr := "0.0.0.0:8889"
	handler := h2c.NewHandler(r, &http2.Server{})
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Println(err)
	}
}

func debugDumpRoutingChi(r *chi.Mux) {
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}
}
