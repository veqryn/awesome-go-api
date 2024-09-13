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
