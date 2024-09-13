# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DefaultError**](DefaultAPI.md#DefaultError) | **Get** /error | Get an error
[**DefaultGreeting**](DefaultAPI.md#DefaultGreeting) | **Get** /greeting/{name} | Say hello to someone
[**DefaultReview**](DefaultAPI.md#DefaultReview) | **Post** /reviews | Send a review



## DefaultError

> map[string]interface{} DefaultError(ctx).Execute()

Get an error



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/protobufv3/openapiv3/openapi_generator/go/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DefaultError(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DefaultError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DefaultError`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DefaultError`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDefaultErrorRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DefaultGreeting

> AwesomeGreetingResp DefaultGreeting(ctx, name).Execute()

Say hello to someone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/protobufv3/openapiv3/openapi_generator/go/gen"
)

func main() {
	name := "name_example" // string | Name to greet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DefaultGreeting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DefaultGreeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DefaultGreeting`: AwesomeGreetingResp
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DefaultGreeting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name to greet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDefaultGreetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwesomeGreetingResp**](AwesomeGreetingResp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DefaultReview

> map[string]interface{} DefaultReview(ctx).Body(body).Execute()

Send a review



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/protobufv3/openapiv3/openapi_generator/go/gen"
)

func main() {
	body := *openapiclient.NewAwesomeReviewReq("Author_example", int64(123)) // AwesomeReviewReq | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DefaultReview(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DefaultReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DefaultReview`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DefaultReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDefaultReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwesomeReviewReq**](AwesomeReviewReq.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

