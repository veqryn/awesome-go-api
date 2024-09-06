# openapi_client.ReviewsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**post_review**](ReviewsApi.md#post_review) | **POST** /reviews | Post a review


# **post_review**
> post_review(post_review_input_body)

Post a review

### Example


```python
import openapi_client
from openapi_client.models.post_review_input_body import PostReviewInputBody
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ReviewsApi(api_client)
    post_review_input_body = openapi_client.PostReviewInputBody() # PostReviewInputBody | 

    try:
        # Post a review
        api_instance.post_review(post_review_input_body)
    except Exception as e:
        print("Exception when calling ReviewsApi->post_review: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **post_review_input_body** | [**PostReviewInputBody**](PostReviewInputBody.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created |  -  |
**0** | Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

