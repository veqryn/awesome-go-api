// Package client provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// AwesomeErrorResp defines model for awesomeErrorResp.
type AwesomeErrorResp = map[string]interface{}

// AwesomeGreetingResp defines model for awesomeGreetingResp.
type AwesomeGreetingResp struct {
	// Message Greeting message
	Message string `json:"message"`
}

// AwesomeReviewReq defines model for awesomeReviewReq.
type AwesomeReviewReq struct {
	// Author Author of the review
	Author string `json:"author"`

	// Message Review message
	Message *string `json:"message,omitempty"`

	// Rating Rating from 1 to 5
	Rating int64 `json:"rating"`
}

// AwesomeReviewResp defines model for awesomeReviewResp.
type AwesomeReviewResp = map[string]interface{}

// ProtobufAny `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//	Foo foo = ...;
//	Any any;
//	any.PackFrom(foo);
//	...
//	if (any.UnpackTo(&foo)) {
//	  ...
//	}
//
// Example 2: Pack and unpack a message in Java.
//
//	   Foo foo = ...;
//	   Any any = Any.pack(foo);
//	   ...
//	   if (any.is(Foo.class)) {
//	     foo = any.unpack(Foo.class);
//	   }
//	   // or ...
//	   if (any.isSameTypeAs(Foo.getDefaultInstance())) {
//	     foo = any.unpack(Foo.getDefaultInstance());
//	   }
//
//	Example 3: Pack and unpack a message in Python.
//
//	   foo = Foo(...)
//	   any = Any()
//	   any.Pack(foo)
//	   ...
//	   if any.Is(Foo.DESCRIPTOR):
//	     any.Unpack(foo)
//	     ...
//
//	Example 4: Pack and unpack a message in Go
//
//	    foo := &pb.Foo{...}
//	    any, err := anypb.New(foo)
//	    if err != nil {
//	      ...
//	    }
//	    ...
//	    foo := &pb.Foo{}
//	    if err := any.UnmarshalTo(foo); err != nil {
//	      ...
//	    }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//	package google.profile;
//	message Person {
//	  string first_name = 1;
//	  string last_name = 2;
//	}
//
//	{
//	  "@type": "type.googleapis.com/google.profile.Person",
//	  "firstName": <string>,
//	  "lastName": <string>
//	}
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//	{
//	  "@type": "type.googleapis.com/google.protobuf.Duration",
//	  "value": "1.212s"
//	}
type ProtobufAny struct {
	// Type A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com. As of May 2023, there are no widely used type server
	// implementations and no plans to implement one.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	Type                 *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// RpcStatus defines model for rpcStatus.
type RpcStatus struct {
	Code    *int32         `json:"code,omitempty"`
	Details *[]ProtobufAny `json:"details,omitempty"`
	Message *string        `json:"message,omitempty"`
}

// DefaultReviewJSONRequestBody defines body for DefaultReview for application/json ContentType.
type DefaultReviewJSONRequestBody = AwesomeReviewReq

// Getter for additional properties for ProtobufAny. Returns the specified
// element and whether it was found
func (a ProtobufAny) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ProtobufAny
func (a *ProtobufAny) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ProtobufAny to handle AdditionalProperties
func (a *ProtobufAny) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["@type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading '@type': %w", err)
		}
		delete(object, "@type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ProtobufAny to handle AdditionalProperties
func (a ProtobufAny) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Type != nil {
		object["@type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '@type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// DefaultError request
	DefaultError(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DefaultGreeting request
	DefaultGreeting(ctx context.Context, name string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DefaultReviewWithBody request with any body
	DefaultReviewWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	DefaultReview(ctx context.Context, body DefaultReviewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) DefaultError(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDefaultErrorRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DefaultGreeting(ctx context.Context, name string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDefaultGreetingRequest(c.Server, name)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DefaultReviewWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDefaultReviewRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DefaultReview(ctx context.Context, body DefaultReviewJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDefaultReviewRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewDefaultErrorRequest generates requests for DefaultError
func NewDefaultErrorRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/error")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDefaultGreetingRequest generates requests for DefaultGreeting
func NewDefaultGreetingRequest(server string, name string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "name", runtime.ParamLocationPath, name)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/greeting/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDefaultReviewRequest calls the generic DefaultReview builder with application/json body
func NewDefaultReviewRequest(server string, body DefaultReviewJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewDefaultReviewRequestWithBody(server, "application/json", bodyReader)
}

// NewDefaultReviewRequestWithBody generates requests for DefaultReview with any type of body
func NewDefaultReviewRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/reviews")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// DefaultErrorWithResponse request
	DefaultErrorWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DefaultErrorResponse, error)

	// DefaultGreetingWithResponse request
	DefaultGreetingWithResponse(ctx context.Context, name string, reqEditors ...RequestEditorFn) (*DefaultGreetingResponse, error)

	// DefaultReviewWithBodyWithResponse request with any body
	DefaultReviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DefaultReviewResponse, error)

	DefaultReviewWithResponse(ctx context.Context, body DefaultReviewJSONRequestBody, reqEditors ...RequestEditorFn) (*DefaultReviewResponse, error)
}

type DefaultErrorResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AwesomeErrorResp
	JSONDefault  *RpcStatus
}

// Status returns HTTPResponse.Status
func (r DefaultErrorResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DefaultErrorResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DefaultGreetingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AwesomeGreetingResp
	JSONDefault  *RpcStatus
}

// Status returns HTTPResponse.Status
func (r DefaultGreetingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DefaultGreetingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DefaultReviewResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *AwesomeReviewResp
	JSONDefault  *RpcStatus
}

// Status returns HTTPResponse.Status
func (r DefaultReviewResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DefaultReviewResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// DefaultErrorWithResponse request returning *DefaultErrorResponse
func (c *ClientWithResponses) DefaultErrorWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DefaultErrorResponse, error) {
	rsp, err := c.DefaultError(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDefaultErrorResponse(rsp)
}

// DefaultGreetingWithResponse request returning *DefaultGreetingResponse
func (c *ClientWithResponses) DefaultGreetingWithResponse(ctx context.Context, name string, reqEditors ...RequestEditorFn) (*DefaultGreetingResponse, error) {
	rsp, err := c.DefaultGreeting(ctx, name, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDefaultGreetingResponse(rsp)
}

// DefaultReviewWithBodyWithResponse request with arbitrary body returning *DefaultReviewResponse
func (c *ClientWithResponses) DefaultReviewWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*DefaultReviewResponse, error) {
	rsp, err := c.DefaultReviewWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDefaultReviewResponse(rsp)
}

func (c *ClientWithResponses) DefaultReviewWithResponse(ctx context.Context, body DefaultReviewJSONRequestBody, reqEditors ...RequestEditorFn) (*DefaultReviewResponse, error) {
	rsp, err := c.DefaultReview(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDefaultReviewResponse(rsp)
}

// ParseDefaultErrorResponse parses an HTTP response from a DefaultErrorWithResponse call
func ParseDefaultErrorResponse(rsp *http.Response) (*DefaultErrorResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DefaultErrorResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AwesomeErrorResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RpcStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDefaultGreetingResponse parses an HTTP response from a DefaultGreetingWithResponse call
func ParseDefaultGreetingResponse(rsp *http.Response) (*DefaultGreetingResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DefaultGreetingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AwesomeGreetingResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RpcStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseDefaultReviewResponse parses an HTTP response from a DefaultReviewWithResponse call
func ParseDefaultReviewResponse(rsp *http.Response) (*DefaultReviewResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DefaultReviewResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest AwesomeReviewResp
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest RpcStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
