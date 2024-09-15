# ErrorModel


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_schema** | **str** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**detail** | **str** | A human-readable explanation specific to this occurrence of the problem. | [optional] 
**errors** | [**List[ErrorDetail]**](ErrorDetail.md) | Optional list of individual error details | [optional] 
**instance** | **str** | A URI reference that identifies the specific occurrence of the problem. | [optional] 
**status** | **int** | HTTP status code | [optional] 
**title** | **str** | A short, human-readable summary of the problem type. This value should not change between occurrences of the error. | [optional] 
**type** | **str** | A URI reference to human-readable documentation for the error. | [optional] [default to 'about:blank']

## Example

```python
from openapi_client.models.error_model import ErrorModel

# TODO update the JSON string below
json = "{}"
# create an instance of ErrorModel from a JSON string
error_model_instance = ErrorModel.from_json(json)
# print the JSON string representation of the object
print(ErrorModel.to_json())

# convert the object into a dict
error_model_dict = error_model_instance.to_dict()
# create an instance of ErrorModel from a dict
error_model_from_dict = ErrorModel.from_dict(error_model_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


