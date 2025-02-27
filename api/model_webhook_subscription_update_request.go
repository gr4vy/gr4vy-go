/*
 * Gr4vy API
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.1.0-beta
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"encoding/json"
)

// WebhookSubscriptionUpdateRequest struct for WebhookSubscriptionUpdateRequest
type WebhookSubscriptionUpdateRequest struct {
	// Defines if this subscription is currently active or not. When deactivated, webhook events are not sent to the endpoint configured for this subscription.
	Active NullableBool `json:"active,omitempty"`
	// The URL of the endpoint to deliver webhook events to.
	Url NullableString `json:"url,omitempty"`
	Authentication NullableWebhookSubscriptionAuthentication `json:"authentication,omitempty"`
}

// NewWebhookSubscriptionUpdateRequest instantiates a new WebhookSubscriptionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookSubscriptionUpdateRequest() *WebhookSubscriptionUpdateRequest {
	this := WebhookSubscriptionUpdateRequest{}
	return &this
}

// NewWebhookSubscriptionUpdateRequestWithDefaults instantiates a new WebhookSubscriptionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookSubscriptionUpdateRequestWithDefaults() *WebhookSubscriptionUpdateRequest {
	this := WebhookSubscriptionUpdateRequest{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSubscriptionUpdateRequest) GetActive() bool {
	if o == nil || o.Active.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Active.Get()
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSubscriptionUpdateRequest) GetActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Active.Get(), o.Active.IsSet()
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookSubscriptionUpdateRequest) HasActive() bool {
	if o != nil && o.Active.IsSet() {
		return true
	}

	return false
}

// SetActive gets a reference to the given NullableBool and assigns it to the Active field.
func (o *WebhookSubscriptionUpdateRequest) SetActive(v bool) {
	o.Active.Set(&v)
}
// SetActiveNil sets the value for Active to be an explicit nil
func (o *WebhookSubscriptionUpdateRequest) SetActiveNil() {
	o.Active.Set(nil)
}

// UnsetActive ensures that no value is present for Active, not even an explicit nil
func (o *WebhookSubscriptionUpdateRequest) UnsetActive() {
	o.Active.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSubscriptionUpdateRequest) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSubscriptionUpdateRequest) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *WebhookSubscriptionUpdateRequest) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *WebhookSubscriptionUpdateRequest) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *WebhookSubscriptionUpdateRequest) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *WebhookSubscriptionUpdateRequest) UnsetUrl() {
	o.Url.Unset()
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookSubscriptionUpdateRequest) GetAuthentication() WebhookSubscriptionAuthentication {
	if o == nil || o.Authentication.Get() == nil {
		var ret WebhookSubscriptionAuthentication
		return ret
	}
	return *o.Authentication.Get()
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookSubscriptionUpdateRequest) GetAuthenticationOk() (*WebhookSubscriptionAuthentication, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Authentication.Get(), o.Authentication.IsSet()
}

// HasAuthentication returns a boolean if a field has been set.
func (o *WebhookSubscriptionUpdateRequest) HasAuthentication() bool {
	if o != nil && o.Authentication.IsSet() {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given NullableWebhookSubscriptionAuthentication and assigns it to the Authentication field.
func (o *WebhookSubscriptionUpdateRequest) SetAuthentication(v WebhookSubscriptionAuthentication) {
	o.Authentication.Set(&v)
}
// SetAuthenticationNil sets the value for Authentication to be an explicit nil
func (o *WebhookSubscriptionUpdateRequest) SetAuthenticationNil() {
	o.Authentication.Set(nil)
}

// UnsetAuthentication ensures that no value is present for Authentication, not even an explicit nil
func (o *WebhookSubscriptionUpdateRequest) UnsetAuthentication() {
	o.Authentication.Unset()
}

func (o WebhookSubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active.IsSet() {
		toSerialize["active"] = o.Active.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.Authentication.IsSet() {
		toSerialize["authentication"] = o.Authentication.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookSubscriptionUpdateRequest struct {
	value *WebhookSubscriptionUpdateRequest
	isSet bool
}

func (v NullableWebhookSubscriptionUpdateRequest) Get() *WebhookSubscriptionUpdateRequest {
	return v.value
}

func (v *NullableWebhookSubscriptionUpdateRequest) Set(val *WebhookSubscriptionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookSubscriptionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookSubscriptionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookSubscriptionUpdateRequest(val *WebhookSubscriptionUpdateRequest) *NullableWebhookSubscriptionUpdateRequest {
	return &NullableWebhookSubscriptionUpdateRequest{value: val, isSet: true}
}

func (v NullableWebhookSubscriptionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookSubscriptionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


