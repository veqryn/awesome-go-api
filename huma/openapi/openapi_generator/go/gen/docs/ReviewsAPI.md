# \ReviewsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostReview**](ReviewsAPI.md#PostReview) | **Post** /reviews | Post a review



## PostReview

> PostReview(ctx).PostReviewInputBody(postReviewInputBody).Execute()

Post a review

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
	postReviewInputBody := *openapiclient.NewPostReviewInputBody("Author_example", int64(123)) // PostReviewInputBody | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReviewsAPI.PostReview(context.Background()).PostReviewInputBody(postReviewInputBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.PostReview``: %v\n", err)
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
- **Accept**: application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

