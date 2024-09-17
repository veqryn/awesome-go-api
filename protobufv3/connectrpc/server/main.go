package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	pb "github.com/veqryn/awesome-go-api/protobufv3/go/gen"
	"github.com/veqryn/awesome-go-api/protobufv3/go/gen/genconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// curl http://localhost:8000/com.github.veqryn.awesome.Default/Greeting --header "Content-Type: application/json" --data '{"name": "Jane"}'
// grpcurl -protoset <(buf build -o -) -plaintext -d '{"name": "Jane"}' localhost:8000 com.github.veqryn.awesome.Default/Greeting
var _ genconnect.DefaultHandler = &app{}

type app struct {
	// pb.UnimplementedDefaultServer // Used to be future proof to api changes
}

func (a app) Greeting(ctx context.Context, c *connect.Request[pb.GreetingReq]) (*connect.Response[pb.GreetingResp], error) {
	resp := connect.NewResponse(&pb.GreetingResp{
		Message: fmt.Sprintf("Hello %s!", c.Msg.Name),
	})
	return resp, nil
}

func (a app) Review(ctx context.Context, c *connect.Request[pb.ReviewReq]) (*connect.Response[pb.ReviewResp], error) {
	fmt.Println(c.HTTPMethod())
	fmt.Println(c.Spec())
	fmt.Println(c.Peer())
	fmt.Println(c.Msg.String())
	return connect.NewResponse(&pb.ReviewResp{}), nil
}

func (a app) Error(ctx context.Context, c *connect.Request[pb.ErrorReq]) (*connect.Response[pb.ErrorResp], error) {
	rpcStatusError := connect.NewError(connect.CodeInvalidArgument, errors.New("This is an example error"))

	cause, err := connect.NewErrorDetail(wrapperspb.Double(224.92))
	if err != nil {
		panic(err)
	}
	rpcStatusError.AddDetail(cause)

	localized, err := connect.NewErrorDetail(&errdetails.LocalizedMessage{
		Locale:  "en-US",
		Message: "Fix your widget",
	})
	if err != nil {
		panic(err)
	}
	rpcStatusError.AddDetail(localized)

	return nil, rpcStatusError
}

func main() {
	path, handler := genconnect.NewDefaultHandler(&app{})

	mux := http.NewServeMux()
	mux.Handle(path, handler)

	http.ListenAndServe(
		"localhost:8000",
		// Wrap with h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
