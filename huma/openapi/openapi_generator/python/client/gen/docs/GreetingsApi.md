# openapi_client.GreetingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**greeting**](GreetingsApi.md#greeting) | **GET** /greeting/{name} | Say hello to someone


# **greeting**
> GetGreetingOutputBody greeting(name)

Say hello to someone

Respond with a greeting

### Example


```python
import openapi_client
from openapi_client.models.get_greeting_output_body import GetGreetingOutputBody
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
    api_instance = openapi_client.GreetingsApi(api_client)
    name = 'world' # str | Name to greet

    try:
        # Say hello to someone
        api_response = api_instance.greeting(name)
        print("The response of GreetingsApi->greeting:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling GreetingsApi->greeting: %s\n" % e)
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
 - **Accept**: application/json, application/problem+json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**0** | Error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

