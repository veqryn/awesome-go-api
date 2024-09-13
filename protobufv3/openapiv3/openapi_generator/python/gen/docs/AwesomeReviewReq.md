# AwesomeReviewReq


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**author** | **str** | Author of the review | 
**message** | **str** | Review message | [optional] 
**rating** | **int** | Rating from 1 to 5 | 

## Example

```python
from openapi_client.models.awesome_review_req import AwesomeReviewReq

# TODO update the JSON string below
json = "{}"
# create an instance of AwesomeReviewReq from a JSON string
awesome_review_req_instance = AwesomeReviewReq.from_json(json)
# print the JSON string representation of the object
print(AwesomeReviewReq.to_json())

# convert the object into a dict
awesome_review_req_dict = awesome_review_req_instance.to_dict()
# create an instance of AwesomeReviewReq from a dict
awesome_review_req_from_dict = AwesomeReviewReq.from_dict(awesome_review_req_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


