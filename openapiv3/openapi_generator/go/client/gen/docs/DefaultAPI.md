# \DefaultAPI

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetError**](DefaultAPI.md#GetError) | **Get** /error | Get an error
[**Greeting**](DefaultAPI.md#Greeting) | **Get** /greeting/{name} | Say hello to someone
[**PostReview**](DefaultAPI.md#PostReview) | **Post** /reviews | Send a review



## GetError

> GetError(ctx).Execute()

Get an error



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/huma/openapi/openapi_generator/go/client/gen"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.GetError(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetError``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetErrorRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Greeting

> GetGreetingOutputBody Greeting(ctx, name).Execute()

Say hello to someone



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/huma/openapi/openapi_generator/go/client/gen"
)

func main() {
	name := "world" // string | Name to greet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Greeting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Greeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Greeting`: GetGreetingOutputBody
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Greeting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name to greet | 

### Other Parameters

Other parameters are passed through a pointer to a apiGreetingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGreetingOutputBody**](GetGreetingOutputBody.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostReview

> PostReview(ctx).PostReviewInputBody(postReviewInputBody).Execute()

Send a review



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/veqryn/awesome-go-api/huma/openapi/openapi_generator/go/client/gen"
)

func main() {
	postReviewInputBody := *openapiclient.NewPostReviewInputBody("Author_example", int64(123)) // PostReviewInputBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.PostReview(context.Background()).PostReviewInputBody(postReviewInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postReviewInputBody** | [**PostReviewInputBody**](PostReviewInputBody.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

