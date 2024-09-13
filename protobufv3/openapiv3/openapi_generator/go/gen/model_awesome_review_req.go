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

// checks if the AwesomeReviewReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwesomeReviewReq{}

// AwesomeReviewReq struct for AwesomeReviewReq
type AwesomeReviewReq struct {
	// Author of the review
	Author string `json:"author"`
	// Review message
	Message *string `json:"message,omitempty"`
	// Rating from 1 to 5
	Rating int64 `json:"rating"`
}

type _AwesomeReviewReq AwesomeReviewReq

// NewAwesomeReviewReq instantiates a new AwesomeReviewReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwesomeReviewReq(author string, rating int64) *AwesomeReviewReq {
	this := AwesomeReviewReq{}
	this.Author = author
	this.Rating = rating
	return &this
}

// NewAwesomeReviewReqWithDefaults instantiates a new AwesomeReviewReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwesomeReviewReqWithDefaults() *AwesomeReviewReq {
	this := AwesomeReviewReq{}
	return &this
}

// GetAuthor returns the Author field value
func (o *AwesomeReviewReq) GetAuthor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *AwesomeReviewReq) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *AwesomeReviewReq) SetAuthor(v string) {
	o.Author = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AwesomeReviewReq) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwesomeReviewReq) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AwesomeReviewReq) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AwesomeReviewReq) SetMessage(v string) {
	o.Message = &v
}

// GetRating returns the Rating field value
func (o *AwesomeReviewReq) GetRating() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *AwesomeReviewReq) GetRatingOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *AwesomeReviewReq) SetRating(v int64) {
	o.Rating = v
}

func (o AwesomeReviewReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwesomeReviewReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["author"] = o.Author
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["rating"] = o.Rating
	return toSerialize, nil
}

func (o *AwesomeReviewReq) UnmarshalJSON(data []byte) (err error) {
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

	varAwesomeReviewReq := _AwesomeReviewReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAwesomeReviewReq)

	if err != nil {
		return err
	}

	*o = AwesomeReviewReq(varAwesomeReviewReq)

	return err
}

type NullableAwesomeReviewReq struct {
	value *AwesomeReviewReq
	isSet bool
}

func (v NullableAwesomeReviewReq) Get() *AwesomeReviewReq {
	return v.value
}

func (v *NullableAwesomeReviewReq) Set(val *AwesomeReviewReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAwesomeReviewReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAwesomeReviewReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwesomeReviewReq(val *AwesomeReviewReq) *NullableAwesomeReviewReq {
	return &NullableAwesomeReviewReq{value: val, isSet: true}
}

func (v NullableAwesomeReviewReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwesomeReviewReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


