package securden_sdk

import (
	"encoding/json"
)

// checks if the GetPasswordGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPasswordGet200Response{}

// GetPasswordGet200Response struct for GetPasswordGet200Response
type GetPasswordGet200Response struct {
	Password string `json:"password,omitempty"`
}

// NewGetPasswordGet200Response instantiates a new GetPasswordGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPasswordGet200Response() *GetPasswordGet200Response {
	this := GetPasswordGet200Response{}
	return &this
}

// NewGetPasswordGet200ResponseWithDefaults instantiates a new GetPasswordGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPasswordGet200ResponseWithDefaults() *GetPasswordGet200Response {
	this := GetPasswordGet200Response{}
	return &this
}

// GetPassword returns the Password field value
func (o *GetPasswordGet200Response) GetPassword() string {
	return o.Password
}

// SetPassword sets the Password field value
func (o *GetPasswordGet200Response) SetPassword(v string) {
	o.Password = v
}

func (o GetPasswordGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPasswordGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Password != "" {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableGetPasswordGet200Response struct {
	value *GetPasswordGet200Response
	isSet bool
}

func (v NullableGetPasswordGet200Response) Get() *GetPasswordGet200Response {
	return v.value
}

func (v *NullableGetPasswordGet200Response) Set(val *GetPasswordGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPasswordGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPasswordGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPasswordGet200Response(val *GetPasswordGet200Response) *NullableGetPasswordGet200Response {
	return &NullableGetPasswordGet200Response{value: val, isSet: true}
}

func (v NullableGetPasswordGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPasswordGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
