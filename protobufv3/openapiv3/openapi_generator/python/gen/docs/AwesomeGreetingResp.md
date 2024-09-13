# AwesomeGreetingResp


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** | Greeting message | 

## Example

```python
from openapi_client.models.awesome_greeting_resp import AwesomeGreetingResp

# TODO update the JSON string below
json = "{}"
# create an instance of AwesomeGreetingResp from a JSON string
awesome_greeting_resp_instance = AwesomeGreetingResp.from_json(json)
# print the JSON string representation of the object
print(AwesomeGreetingResp.to_json())

# convert the object into a dict
awesome_greeting_resp_dict = awesome_greeting_resp_instance.to_dict()
# create an instance of AwesomeGreetingResp from a dict
awesome_greeting_resp_from_dict = AwesomeGreetingResp.from_dict(awesome_greeting_resp_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


