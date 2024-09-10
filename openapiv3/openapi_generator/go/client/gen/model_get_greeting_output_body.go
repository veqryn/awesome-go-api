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

// checks if the GetGreetingOutputBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGreetingOutputBody{}

// GetGreetingOutputBody struct for GetGreetingOutputBody
type GetGreetingOutputBody struct {
	// Greeting message
	Message string `json:"message"`
}

type _GetGreetingOutputBody GetGreetingOutputBody

// NewGetGreetingOutputBody instantiates a new GetGreetingOutputBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGreetingOutputBody(message string) *GetGreetingOutputBody {
	this := GetGreetingOutputBody{}
	this.Message = message
	return &this
}

// NewGetGreetingOutputBodyWithDefaults instantiates a new GetGreetingOutputBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGreetingOutputBodyWithDefaults() *GetGreetingOutputBody {
	this := GetGreetingOutputBody{}
	return &this
}

// GetMessage returns the Message field value
func (o *GetGreetingOutputBody) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GetGreetingOutputBody) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GetGreetingOutputBody) SetMessage(v string) {
	o.Message = v
}

func (o GetGreetingOutputBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGreetingOutputBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *GetGreetingOutputBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
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

	varGetGreetingOutputBody := _GetGreetingOutputBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetGreetingOutputBody)

	if err != nil {
		return err
	}

	*o = GetGreetingOutputBody(varGetGreetingOutputBody)

	return err
}

type NullableGetGreetingOutputBody struct {
	value *GetGreetingOutputBody
	isSet bool
}

func (v NullableGetGreetingOutputBody) Get() *GetGreetingOutputBody {
	return v.value
}

func (v *NullableGetGreetingOutputBody) Set(val *GetGreetingOutputBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGreetingOutputBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGreetingOutputBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGreetingOutputBody(val *GetGreetingOutputBody) *NullableGetGreetingOutputBody {
	return &NullableGetGreetingOutputBody{value: val, isSet: true}
}

func (v NullableGetGreetingOutputBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGreetingOutputBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

