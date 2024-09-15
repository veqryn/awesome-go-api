# \GreetingsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Greeting**](GreetingsAPI.md#Greeting) | **Get** /greeting/{name} | Say hello to someone



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
	openapiclient "github.com/veqryn/awesome-go-api/huma/openapi/openapi_generator/go/gen"
)

func main() {
	name := "world" // string | Name to greet

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GreetingsAPI.Greeting(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GreetingsAPI.Greeting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Greeting`: GetGreetingOutputBody
	fmt.Fprintf(os.Stdout, "Response from `GreetingsAPI.Greeting`: %v\n", resp)
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
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

