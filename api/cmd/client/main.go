package main

import (
	barv1 "connect-sandbox/api/gen/protobuf/bar/v1"
	"connect-sandbox/api/gen/protobuf/bar/v1/barv1connect"
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
)

var token string

func init() {
	token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxIiwiaWF0IjoxNjc2OTAyMDYzLCJuYmYiOjE2NzY5MDIwNjMsImp0aSI6IjVkMWM3MzU4LWZhOTEtNDUzMi1hZWI0LTFiNDhhNDQ3NGVkMSIsInR5cGUiOiJhY2Nlc3MiLCJmcmVzaCI6ZmFsc2V9.ddrL2Y9ykue4sOfsvoa8AmNvRllcSeNJ2b_ksqUPxfo"
	log.SetFlags(log.Lshortfile)
}

func main() {
	client := barv1connect.NewBarServiceClient(
		http.DefaultClient,
		"http://localhost:8889",
	)
	req := connect.NewRequest(&barv1.GetRequest{
		Id:   10,
		Name: "foo",
	})
	req.Header().Set("Authorization", "Bearer "+token)
	res, err := client.Get(
		context.Background(),
		req,
	)
	if err != nil {
		var connectErr *connect.Error
		log.Println(err)
		if !errors.As(err, &connectErr) {
			return
		}
		for _, detail := range connectErr.Details() {
			log.Println(detail)
		}
	}
	log.Println(res)
}
