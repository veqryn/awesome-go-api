# AwesomeReviewReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | **string** | Author of the review | 
**Message** | Pointer to **string** | Review message | [optional] 
**Rating** | **int64** | Rating from 1 to 5 | 

## Methods

### NewAwesomeReviewReq

`func NewAwesomeReviewReq(author string, rating int64, ) *AwesomeReviewReq`

NewAwesomeReviewReq instantiates a new AwesomeReviewReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwesomeReviewReqWithDefaults

`func NewAwesomeReviewReqWithDefaults() *AwesomeReviewReq`

NewAwesomeReviewReqWithDefaults instantiates a new AwesomeReviewReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AwesomeReviewReq) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AwesomeReviewReq) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AwesomeReviewReq) SetAuthor(v string)`

SetAuthor sets Author field to given value.


### GetMessage

`func (o *AwesomeReviewReq) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AwesomeReviewReq) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AwesomeReviewReq) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AwesomeReviewReq) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRating

`func (o *AwesomeReviewReq) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *AwesomeReviewReq) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *AwesomeReviewReq) SetRating(v int64)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


