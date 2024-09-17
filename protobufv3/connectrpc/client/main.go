package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"github.com/davecgh/go-spew/spew"
	pb "github.com/veqryn/awesome-go-api/protobufv3/go/gen"
	"github.com/veqryn/awesome-go-api/protobufv3/go/gen/genconnect"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails" // register the error detail types
)

func main() {
	ctx := context.Background()

	c := genconnect.NewDefaultClient(
		http.DefaultClient,
		"http://localhost:8000",
		// connect.WithGRPC(),
	)

	fmt.Println("--- Getting a Greeting:")
	greetingResp, err := c.Greeting(ctx, connect.NewRequest(&pb.GreetingReq{Name: "Bob"}))
	if err != nil {
		panic(err)
	}
	spew.Dump(greetingResp.Header())
	spew.Dump(greetingResp.Trailer())
	spew.Dump(greetingResp.Msg)

	fmt.Println("--- Sending a Review:")
	message := "foobar"
	reviewResp, err := c.Review(ctx, connect.NewRequest(&pb.ReviewReq{
		Author:  "Bob",
		Message: &message,
		Rating:  4,
	}))
	if err != nil {
		panic(err)
	}
	spew.Dump(reviewResp.Msg)

	fmt.Println("--- Getting an Error:")
	errorResp, err := c.Error(ctx, connect.NewRequest(&pb.ErrorReq{}))
	spew.Dump(errorResp)
	spew.Dump(err)
	fmt.Println(connect.CodeOf(err))
	if connectErr := new(connect.Error); errors.As(err, &connectErr) {
		spew.Dump(connectErr.Error())
		spew.Dump(connectErr.Message())
		spew.Dump(connectErr.Details())
	}
}
