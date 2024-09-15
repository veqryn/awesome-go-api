# GetGreetingOutputBody


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**var_schema** | **str** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**message** | **str** | Greeting message | 

## Example

```python
from openapi_client.models.get_greeting_output_body import GetGreetingOutputBody

# TODO update the JSON string below
json = "{}"
# create an instance of GetGreetingOutputBody from a JSON string
get_greeting_output_body_instance = GetGreetingOutputBody.from_json(json)
# print the JSON string representation of the object
print(GetGreetingOutputBody.to_json())

# convert the object into a dict
get_greeting_output_body_dict = get_greeting_output_body_instance.to_dict()
# create an instance of GetGreetingOutputBody from a dict
get_greeting_output_body_from_dict = GetGreetingOutputBody.from_dict(get_greeting_output_body_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


