# PostReviewInputBody


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**author** | **str** | Author of the review | 
**message** | **str** | Review message | [optional] 
**rating** | **int** | Rating from 1 to 5 | 

## Example

```python
from openapi_client.models.post_review_input_body import PostReviewInputBody

# TODO update the JSON string below
json = "{}"
# create an instance of PostReviewInputBody from a JSON string
post_review_input_body_instance = PostReviewInputBody.from_json(json)
# print the JSON string representation of the object
print(PostReviewInputBody.to_json())

# convert the object into a dict
post_review_input_body_dict = post_review_input_body_instance.to_dict()
# create an instance of PostReviewInputBody from a dict
post_review_input_body_from_dict = PostReviewInputBody.from_dict(post_review_input_body_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


