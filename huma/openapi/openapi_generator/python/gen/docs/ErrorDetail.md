# ErrorDetail


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**location** | **str** | Where the error occurred, e.g. &#39;body.items[3].tags&#39; or &#39;path.thing-id&#39; | [optional] 
**message** | **str** | Error message text | [optional] 
**value** | **object** | The value at the given location | [optional] 

## Example

```python
from openapi_client.models.error_detail import ErrorDetail

# TODO update the JSON string below
json = "{}"
# create an instance of ErrorDetail from a JSON string
error_detail_instance = ErrorDetail.from_json(json)
# print the JSON string representation of the object
print(ErrorDetail.to_json())

# convert the object into a dict
error_detail_dict = error_detail_instance.to_dict()
# create an instance of ErrorDetail from a dict
error_detail_from_dict = ErrorDetail.from_dict(error_detail_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


