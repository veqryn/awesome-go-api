# openapi_client.DefaultApi

All URIs are relative to *http://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**default_error**](DefaultApi.md#default_error) | **GET** /error | Get an error
[**default_greeting**](DefaultApi.md#default_greeting) | **GET** /greeting/{name} | Say hello to someone
[**default_review**](DefaultApi.md#default_review) | **POST** /reviews | Send a review


# **default_error**
> object default_error()

Get an error

Responds with an error

### Example


```python
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)

    try:
        # Get an error
        api_response = api_instance.default_error()
        print("The response of DefaultApi->default_error:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->default_error: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **default_greeting**
> AwesomeGreetingResp default_greeting(name)

Say hello to someone

Responds with a greeting

### Example


```python
import openapi_client
from openapi_client.models.awesome_greeting_resp import AwesomeGreetingResp
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    name = 'name_example' # str | Name to greet

    try:
        # Say hello to someone
        api_response = api_instance.default_greeting(name)
        print("The response of DefaultApi->default_greeting:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->default_greeting: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **str**| Name to greet | 

### Return type

[**AwesomeGreetingResp**](AwesomeGreetingResp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **default_review**
> object default_review(body)

Send a review

Post a review to be saved

### Example


```python
import openapi_client
from openapi_client.models.awesome_review_req import AwesomeReviewReq
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost:8080
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost:8080"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    body = openapi_client.AwesomeReviewReq() # AwesomeReviewReq | 

    try:
        # Send a review
        api_response = api_instance.default_review(body)
        print("The response of DefaultApi->default_review:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DefaultApi->default_review: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AwesomeReviewReq**](AwesomeReviewReq.md)|  | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A successful response. |  -  |
**0** | An unexpected error response. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

