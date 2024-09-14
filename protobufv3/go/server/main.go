package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/veqryn/awesome-go-api/protobufv3/go/gen"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var _ pb.DefaultServer = &app{}

type app struct {
	// pb.UnimplementedDefaultServer // Used to be future proof to api changes
}

func (s *app) Greeting(ctx context.Context, req *pb.GreetingReq) (*pb.GreetingResp, error) {
	resp := &pb.GreetingResp{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}
	return resp, nil
}

func (s *app) Review(ctx context.Context, req *pb.ReviewReq) (*pb.ReviewResp, error) {
	fmt.Println(req.String())
	return &pb.ReviewResp{}, nil
}

func (s *app) Error(ctx context.Context, req *pb.ErrorReq) (*pb.ErrorResp, error) {
	// Error codes: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	// Map to HTTP codes: https://gist.github.com/hamakn/708b9802ca845eb59f3975dbb3ae2a01

	// This is how you return a custom status or error
	st := status.New(codes.InvalidArgument, "This is an example error")

	// Additional details can be appended to the status
	cause, err := anypb.New(wrapperspb.Double(224.92))
	if err != nil {
		panic(err)
	}

	st, err = st.WithDetails(
		cause,
		// errdetails contains well known error details
		&errdetails.LocalizedMessage{
			Locale:  "en-US",
			Message: "Fix your widget",
		},
	)
	if err != nil {
		panic(err)
	}

	return nil, st.Err()
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	// Create a gRPC server, and register our app as the handler/server for the service interface
	s := grpc.NewServer()
	pb.RegisterDefaultServer(s, &app{})
	go func() {
		serveErr := s.Serve(lis)
		if serveErr != nil {
			panic(serveErr)
		}
	}()

	// Create the grpc-gateway for proxying rest/json calls
	gwmux := runtime.NewServeMux()

	// Create a client to proxy the requests through
	clientConn, err := grpc.NewClient("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	err = pb.RegisterDefaultHandler(context.Background(), gwmux, clientConn)
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8080", gwmux)
}
