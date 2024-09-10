# Error

An API error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**title** | **str** | A short, human-readable summary of the problem type. This value should not change between occurrences of the error. | [optional] 
**details** | **str** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**properties** | **Dict[str, object]** | Optional map of properties | [optional] 

## Example

```python
from openapi_client.models.error import Error

# TODO update the JSON string below
json = "{}"
# create an instance of Error from a JSON string
error_instance = Error.from_json(json)
# print the JSON string representation of the object
print(Error.to_json())

# convert the object into a dict
error_dict = error_instance.to_dict()
# create an instance of Error from a dict
error_from_dict = Error.from_dict(error_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


