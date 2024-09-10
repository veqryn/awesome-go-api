/*
Awesome GO API

Actual example use cases for a curated list of golang api generator libraries

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostReviewInputBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostReviewInputBody{}

// PostReviewInputBody struct for PostReviewInputBody
type PostReviewInputBody struct {
	// Author of the review
	Author string `json:"author"`
	// Review message
	Message *string `json:"message,omitempty"`
	// Rating from 1 to 5
	Rating int64 `json:"rating"`
}

type _PostReviewInputBody PostReviewInputBody

// NewPostReviewInputBody instantiates a new PostReviewInputBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostReviewInputBody(author string, rating int64) *PostReviewInputBody {
	this := PostReviewInputBody{}
	this.Author = author
	this.Rating = rating
	return &this
}

// NewPostReviewInputBodyWithDefaults instantiates a new PostReviewInputBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostReviewInputBodyWithDefaults() *PostReviewInputBody {
	this := PostReviewInputBody{}
	return &this
}

// GetAuthor returns the Author field value
func (o *PostReviewInputBody) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *PostReviewInputBody) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *PostReviewInputBody) SetAuthor(v string) {
	o.Author = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PostReviewInputBody) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostReviewInputBody) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PostReviewInputBody) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PostReviewInputBody) SetMessage(v string) {
	o.Message = &v
}

// GetRating returns the Rating field value
func (o *PostReviewInputBody) GetRating() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *PostReviewInputBody) GetRatingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *PostReviewInputBody) SetRating(v int64) {
	o.Rating = v
}

func (o PostReviewInputBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostReviewInputBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["author"] = o.Author
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["rating"] = o.Rating
	return toSerialize, nil
}

func (o *PostReviewInputBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"author",
		"rating",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostReviewInputBody := _PostReviewInputBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostReviewInputBody)

	if err != nil {
		return err
	}

	*o = PostReviewInputBody(varPostReviewInputBody)

	return err
}

type NullablePostReviewInputBody struct {
	value *PostReviewInputBody
	isSet bool
}

func (v NullablePostReviewInputBody) Get() *PostReviewInputBody {
	return v.value
}

func (v *NullablePostReviewInputBody) Set(val *PostReviewInputBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePostReviewInputBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePostReviewInputBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostReviewInputBody(val *PostReviewInputBody) *NullablePostReviewInputBody {
	return &NullablePostReviewInputBody{value: val, isSet: true}
}

func (v NullablePostReviewInputBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostReviewInputBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


