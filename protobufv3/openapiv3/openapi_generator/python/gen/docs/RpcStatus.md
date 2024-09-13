# RpcStatus


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**code** | **int** |  | [optional] 
**message** | **str** |  | [optional] 
**details** | [**List[ProtobufAny]**](ProtobufAny.md) |  | [optional] 

## Example

```python
from openapi_client.models.rpc_status import RpcStatus

# TODO update the JSON string below
json = "{}"
# create an instance of RpcStatus from a JSON string
rpc_status_instance = RpcStatus.from_json(json)
# print the JSON string representation of the object
print(RpcStatus.to_json())

# convert the object into a dict
rpc_status_dict = rpc_status_instance.to_dict()
# create an instance of RpcStatus from a dict
rpc_status_from_dict = RpcStatus.from_dict(rpc_status_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


