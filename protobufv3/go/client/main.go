/*
grpc generates both clients and servers from a proto spec.
Every call includes an argument struct with the fields required.
Every call returns a struct and any error.
Any errors or non-OK status codes are returned in the error.
The returned error must be converted from a golang error to a protobuf status,
after which you can access the status code, message, and a free-form slice of any's.
The returned status codes are different http, and there is only 1 OK status code.
*/
package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/veqryn/awesome-go-api/protobufv3/go/gen"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails" // register the error detail types
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewDefaultClient(conn)

	fmt.Println("--- Getting a Greeting:")
	greetingResp, err := c.Greeting(ctx, &pb.GreetingReq{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	spew.Dump(greetingResp)

	fmt.Println("--- Sending a Review:")
	message := "foobar"
	reviewResp, err := c.Review(ctx, &pb.ReviewReq{
		Author:  "Bob",
		Message: &message,
		Rating:  4,
	})
	if err != nil {
		panic(err)
	}
	spew.Dump(reviewResp)

	fmt.Println("--- Getting an Error:")
	errorResp, err := c.Error(ctx, &pb.ErrorReq{})
	spew.Dump(errorResp)
	spew.Dump(err)
	st := status.Convert(err)
	spew.Dump(st)
	spew.Dump(st.Details())
}
