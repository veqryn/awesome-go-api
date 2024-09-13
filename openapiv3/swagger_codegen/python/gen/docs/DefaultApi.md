# swagger_client.DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_error**](DefaultApi.md#get_error) | **GET** /error | Get an error
[**greeting**](DefaultApi.md#greeting) | **GET** /greeting/{name} | Say hello to someone
[**post_review**](DefaultApi.md#post_review) | **POST** /reviews | Send a review

# **get_error**
> get_error()

Get an error

Responds with an error

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()

try:
    # Get an error
    api_instance.get_error()
except ApiException as e:
    print("Exception when calling DefaultApi->get_error: %s\n" % e)
```

### Parameters
This endpoint does not need any parameter.

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **greeting**
> GetGreetingOutputBody greeting(name)

Say hello to someone

Responds with a greeting

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
name = 'name_example' # str | Name to greet

try:
    # Say hello to someone
    api_response = api_instance.greeting(name)
    pprint(api_response)
except ApiException as e:
    print("Exception when calling DefaultApi->greeting: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Name to greet | 

### Return type

[**GetGreetingOutputBody**](GetGreetingOutputBody.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_review**
> post_review(body)

Send a review

Post a review to be saved

### Example
```python
from __future__ import print_function
import time
import swagger_client
from swagger_client.rest import ApiException
from pprint import pprint

# create an instance of the API class
api_instance = swagger_client.DefaultApi()
body = swagger_client.PostReviewInputBody() # PostReviewInputBody | 

try:
    # Send a review
    api_instance.post_review(body)
except ApiException as e:
    print("Exception when calling DefaultApi->post_review: %s\n" % e)
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PostReviewInputBody**](PostReviewInputBody.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

