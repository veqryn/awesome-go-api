# PostReviewInputBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] [readonly] 
**Author** | **string** | Author of the review | 
**Message** | Pointer to **string** | Review message | [optional] 
**Rating** | **int64** | Rating from 1 to 5 | 

## Methods

### NewPostReviewInputBody

`func NewPostReviewInputBody(author string, rating int64, ) *PostReviewInputBody`

NewPostReviewInputBody instantiates a new PostReviewInputBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostReviewInputBodyWithDefaults

`func NewPostReviewInputBodyWithDefaults() *PostReviewInputBody`

NewPostReviewInputBodyWithDefaults instantiates a new PostReviewInputBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *PostReviewInputBody) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *PostReviewInputBody) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *PostReviewInputBody) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *PostReviewInputBody) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetAuthor

`func (o *PostReviewInputBody) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *PostReviewInputBody) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *PostReviewInputBody) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetMessage

`func (o *PostReviewInputBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostReviewInputBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostReviewInputBody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostReviewInputBody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRating

`func (o *PostReviewInputBody) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PostReviewInputBody) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PostReviewInputBody) SetRating(v int64)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


