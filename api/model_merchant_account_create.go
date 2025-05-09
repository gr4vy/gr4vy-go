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

// MerchantAccountCreate A request to create a merchant account.
type MerchantAccountCreate struct {
	// The unique ID for the merchant account.
	Id *string `json:"id,omitempty"`
	// The human-readable name of the merchant account.
	DisplayName *string `json:"display_name,omitempty"`
	// The optional URL where webhooks will be received.
	OutboundWebhookUrl NullableString `json:"outbound_webhook_url,omitempty"`
	// The optional username to use when `outbound_webhook_url` is configured and requires basic authentication.
	OutboundWebhookUsername NullableString `json:"outbound_webhook_username,omitempty"`
	// The optional password to use when `outbound_webhook_url` is configured and requires basic authentication.
	OutboundWebhookPassword NullableString `json:"outbound_webhook_password,omitempty"`
	// Requestor ID provided for Visa after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts.
	VisaNetworkTokensRequestorId NullableString `json:"visa_network_tokens_requestor_id,omitempty"`
	// Application ID provided for Visa after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts.
	VisaNetworkTokensAppId NullableString `json:"visa_network_tokens_app_id,omitempty"`
	// Requestor ID provided for Amex after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts.
	AmexNetworkTokensRequestorId NullableString `json:"amex_network_tokens_requestor_id,omitempty"`
	// Application ID provided for Amex after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts.
	AmexNetworkTokensAppId NullableString `json:"amex_network_tokens_app_id,omitempty"`
	// Requestor ID provided for Mastercard after onboarding to use Network Tokens. The requestor ID must be unique across all schemes and merchant accounts.
	MastercardNetworkTokensRequestorId NullableString `json:"mastercard_network_tokens_requestor_id,omitempty"`
	// Application ID provided for Mastercard after onboarding to use Network Tokens. The application ID must be unique across all schemes and merchant accounts.
	MastercardNetworkTokensAppId NullableString `json:"mastercard_network_tokens_app_id,omitempty"`
	// Client key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it's set to `null`, the Account Updater service doesn't get configured. * If the field is set to `null`, the other `loon_*` fields must be set to `null` as well.
	LoonClientKey NullableString `json:"loon_client_key,omitempty"`
	// Secret key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it's set to `null`, the Account Updater service doesn't get configured. * If the field is set to `null`, the other `loon_*` fields must be set to `null` as well.
	LoonSecretKey NullableString `json:"loon_secret_key,omitempty"`
	// Card schemes accepted when creating jobs using this set of Loon API keys. Loon is the Account Updater service used by Gr4vy.  * If the field is not set or if it's set to `null`, the Account Updater service doesn't get configured. * If the field is set to `null`, the other `loon_*` fields must be set to `null` as well.
	LoonAcceptedSchemes []string `json:"loon_accepted_schemes,omitempty"`
	// The maximum monetary amount allowed for over-capture, in the smallest currency unit, for example `1299` cents to allow for an over-capture of `$12.99`.
	OverCaptureAmount NullableFloat32 `json:"over_capture_amount,omitempty"`
	// The maximum percentage allowed for over-capture, for example `25` to allow for an over-capture of 25% of the original transaction amount.
	OverCapturePercentage NullableFloat32 `json:"over_capture_percentage,omitempty"`
}

// NewMerchantAccountCreate instantiates a new MerchantAccountCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantAccountCreate() *MerchantAccountCreate {
	this := MerchantAccountCreate{}
	return &this
}

// NewMerchantAccountCreateWithDefaults instantiates a new MerchantAccountCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantAccountCreateWithDefaults() *MerchantAccountCreate {
	this := MerchantAccountCreate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MerchantAccountCreate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAccountCreate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MerchantAccountCreate) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MerchantAccountCreate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantAccountCreate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MerchantAccountCreate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetOutboundWebhookUrl returns the OutboundWebhookUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetOutboundWebhookUrl() string {
	if o == nil || o.OutboundWebhookUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutboundWebhookUrl.Get()
}

// GetOutboundWebhookUrlOk returns a tuple with the OutboundWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetOutboundWebhookUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OutboundWebhookUrl.Get(), o.OutboundWebhookUrl.IsSet()
}

// HasOutboundWebhookUrl returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasOutboundWebhookUrl() bool {
	if o != nil && o.OutboundWebhookUrl.IsSet() {
		return true
	}

	return false
}

// SetOutboundWebhookUrl gets a reference to the given NullableString and assigns it to the OutboundWebhookUrl field.
func (o *MerchantAccountCreate) SetOutboundWebhookUrl(v string) {
	o.OutboundWebhookUrl.Set(&v)
}
// SetOutboundWebhookUrlNil sets the value for OutboundWebhookUrl to be an explicit nil
func (o *MerchantAccountCreate) SetOutboundWebhookUrlNil() {
	o.OutboundWebhookUrl.Set(nil)
}

// UnsetOutboundWebhookUrl ensures that no value is present for OutboundWebhookUrl, not even an explicit nil
func (o *MerchantAccountCreate) UnsetOutboundWebhookUrl() {
	o.OutboundWebhookUrl.Unset()
}

// GetOutboundWebhookUsername returns the OutboundWebhookUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetOutboundWebhookUsername() string {
	if o == nil || o.OutboundWebhookUsername.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutboundWebhookUsername.Get()
}

// GetOutboundWebhookUsernameOk returns a tuple with the OutboundWebhookUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetOutboundWebhookUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OutboundWebhookUsername.Get(), o.OutboundWebhookUsername.IsSet()
}

// HasOutboundWebhookUsername returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasOutboundWebhookUsername() bool {
	if o != nil && o.OutboundWebhookUsername.IsSet() {
		return true
	}

	return false
}

// SetOutboundWebhookUsername gets a reference to the given NullableString and assigns it to the OutboundWebhookUsername field.
func (o *MerchantAccountCreate) SetOutboundWebhookUsername(v string) {
	o.OutboundWebhookUsername.Set(&v)
}
// SetOutboundWebhookUsernameNil sets the value for OutboundWebhookUsername to be an explicit nil
func (o *MerchantAccountCreate) SetOutboundWebhookUsernameNil() {
	o.OutboundWebhookUsername.Set(nil)
}

// UnsetOutboundWebhookUsername ensures that no value is present for OutboundWebhookUsername, not even an explicit nil
func (o *MerchantAccountCreate) UnsetOutboundWebhookUsername() {
	o.OutboundWebhookUsername.Unset()
}

// GetOutboundWebhookPassword returns the OutboundWebhookPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetOutboundWebhookPassword() string {
	if o == nil || o.OutboundWebhookPassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.OutboundWebhookPassword.Get()
}

// GetOutboundWebhookPasswordOk returns a tuple with the OutboundWebhookPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetOutboundWebhookPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OutboundWebhookPassword.Get(), o.OutboundWebhookPassword.IsSet()
}

// HasOutboundWebhookPassword returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasOutboundWebhookPassword() bool {
	if o != nil && o.OutboundWebhookPassword.IsSet() {
		return true
	}

	return false
}

// SetOutboundWebhookPassword gets a reference to the given NullableString and assigns it to the OutboundWebhookPassword field.
func (o *MerchantAccountCreate) SetOutboundWebhookPassword(v string) {
	o.OutboundWebhookPassword.Set(&v)
}
// SetOutboundWebhookPasswordNil sets the value for OutboundWebhookPassword to be an explicit nil
func (o *MerchantAccountCreate) SetOutboundWebhookPasswordNil() {
	o.OutboundWebhookPassword.Set(nil)
}

// UnsetOutboundWebhookPassword ensures that no value is present for OutboundWebhookPassword, not even an explicit nil
func (o *MerchantAccountCreate) UnsetOutboundWebhookPassword() {
	o.OutboundWebhookPassword.Unset()
}

// GetVisaNetworkTokensRequestorId returns the VisaNetworkTokensRequestorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetVisaNetworkTokensRequestorId() string {
	if o == nil || o.VisaNetworkTokensRequestorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VisaNetworkTokensRequestorId.Get()
}

// GetVisaNetworkTokensRequestorIdOk returns a tuple with the VisaNetworkTokensRequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetVisaNetworkTokensRequestorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VisaNetworkTokensRequestorId.Get(), o.VisaNetworkTokensRequestorId.IsSet()
}

// HasVisaNetworkTokensRequestorId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasVisaNetworkTokensRequestorId() bool {
	if o != nil && o.VisaNetworkTokensRequestorId.IsSet() {
		return true
	}

	return false
}

// SetVisaNetworkTokensRequestorId gets a reference to the given NullableString and assigns it to the VisaNetworkTokensRequestorId field.
func (o *MerchantAccountCreate) SetVisaNetworkTokensRequestorId(v string) {
	o.VisaNetworkTokensRequestorId.Set(&v)
}
// SetVisaNetworkTokensRequestorIdNil sets the value for VisaNetworkTokensRequestorId to be an explicit nil
func (o *MerchantAccountCreate) SetVisaNetworkTokensRequestorIdNil() {
	o.VisaNetworkTokensRequestorId.Set(nil)
}

// UnsetVisaNetworkTokensRequestorId ensures that no value is present for VisaNetworkTokensRequestorId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetVisaNetworkTokensRequestorId() {
	o.VisaNetworkTokensRequestorId.Unset()
}

// GetVisaNetworkTokensAppId returns the VisaNetworkTokensAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetVisaNetworkTokensAppId() string {
	if o == nil || o.VisaNetworkTokensAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.VisaNetworkTokensAppId.Get()
}

// GetVisaNetworkTokensAppIdOk returns a tuple with the VisaNetworkTokensAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetVisaNetworkTokensAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.VisaNetworkTokensAppId.Get(), o.VisaNetworkTokensAppId.IsSet()
}

// HasVisaNetworkTokensAppId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasVisaNetworkTokensAppId() bool {
	if o != nil && o.VisaNetworkTokensAppId.IsSet() {
		return true
	}

	return false
}

// SetVisaNetworkTokensAppId gets a reference to the given NullableString and assigns it to the VisaNetworkTokensAppId field.
func (o *MerchantAccountCreate) SetVisaNetworkTokensAppId(v string) {
	o.VisaNetworkTokensAppId.Set(&v)
}
// SetVisaNetworkTokensAppIdNil sets the value for VisaNetworkTokensAppId to be an explicit nil
func (o *MerchantAccountCreate) SetVisaNetworkTokensAppIdNil() {
	o.VisaNetworkTokensAppId.Set(nil)
}

// UnsetVisaNetworkTokensAppId ensures that no value is present for VisaNetworkTokensAppId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetVisaNetworkTokensAppId() {
	o.VisaNetworkTokensAppId.Unset()
}

// GetAmexNetworkTokensRequestorId returns the AmexNetworkTokensRequestorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetAmexNetworkTokensRequestorId() string {
	if o == nil || o.AmexNetworkTokensRequestorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AmexNetworkTokensRequestorId.Get()
}

// GetAmexNetworkTokensRequestorIdOk returns a tuple with the AmexNetworkTokensRequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetAmexNetworkTokensRequestorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AmexNetworkTokensRequestorId.Get(), o.AmexNetworkTokensRequestorId.IsSet()
}

// HasAmexNetworkTokensRequestorId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasAmexNetworkTokensRequestorId() bool {
	if o != nil && o.AmexNetworkTokensRequestorId.IsSet() {
		return true
	}

	return false
}

// SetAmexNetworkTokensRequestorId gets a reference to the given NullableString and assigns it to the AmexNetworkTokensRequestorId field.
func (o *MerchantAccountCreate) SetAmexNetworkTokensRequestorId(v string) {
	o.AmexNetworkTokensRequestorId.Set(&v)
}
// SetAmexNetworkTokensRequestorIdNil sets the value for AmexNetworkTokensRequestorId to be an explicit nil
func (o *MerchantAccountCreate) SetAmexNetworkTokensRequestorIdNil() {
	o.AmexNetworkTokensRequestorId.Set(nil)
}

// UnsetAmexNetworkTokensRequestorId ensures that no value is present for AmexNetworkTokensRequestorId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetAmexNetworkTokensRequestorId() {
	o.AmexNetworkTokensRequestorId.Unset()
}

// GetAmexNetworkTokensAppId returns the AmexNetworkTokensAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetAmexNetworkTokensAppId() string {
	if o == nil || o.AmexNetworkTokensAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AmexNetworkTokensAppId.Get()
}

// GetAmexNetworkTokensAppIdOk returns a tuple with the AmexNetworkTokensAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetAmexNetworkTokensAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AmexNetworkTokensAppId.Get(), o.AmexNetworkTokensAppId.IsSet()
}

// HasAmexNetworkTokensAppId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasAmexNetworkTokensAppId() bool {
	if o != nil && o.AmexNetworkTokensAppId.IsSet() {
		return true
	}

	return false
}

// SetAmexNetworkTokensAppId gets a reference to the given NullableString and assigns it to the AmexNetworkTokensAppId field.
func (o *MerchantAccountCreate) SetAmexNetworkTokensAppId(v string) {
	o.AmexNetworkTokensAppId.Set(&v)
}
// SetAmexNetworkTokensAppIdNil sets the value for AmexNetworkTokensAppId to be an explicit nil
func (o *MerchantAccountCreate) SetAmexNetworkTokensAppIdNil() {
	o.AmexNetworkTokensAppId.Set(nil)
}

// UnsetAmexNetworkTokensAppId ensures that no value is present for AmexNetworkTokensAppId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetAmexNetworkTokensAppId() {
	o.AmexNetworkTokensAppId.Unset()
}

// GetMastercardNetworkTokensRequestorId returns the MastercardNetworkTokensRequestorId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetMastercardNetworkTokensRequestorId() string {
	if o == nil || o.MastercardNetworkTokensRequestorId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MastercardNetworkTokensRequestorId.Get()
}

// GetMastercardNetworkTokensRequestorIdOk returns a tuple with the MastercardNetworkTokensRequestorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetMastercardNetworkTokensRequestorIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MastercardNetworkTokensRequestorId.Get(), o.MastercardNetworkTokensRequestorId.IsSet()
}

// HasMastercardNetworkTokensRequestorId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasMastercardNetworkTokensRequestorId() bool {
	if o != nil && o.MastercardNetworkTokensRequestorId.IsSet() {
		return true
	}

	return false
}

// SetMastercardNetworkTokensRequestorId gets a reference to the given NullableString and assigns it to the MastercardNetworkTokensRequestorId field.
func (o *MerchantAccountCreate) SetMastercardNetworkTokensRequestorId(v string) {
	o.MastercardNetworkTokensRequestorId.Set(&v)
}
// SetMastercardNetworkTokensRequestorIdNil sets the value for MastercardNetworkTokensRequestorId to be an explicit nil
func (o *MerchantAccountCreate) SetMastercardNetworkTokensRequestorIdNil() {
	o.MastercardNetworkTokensRequestorId.Set(nil)
}

// UnsetMastercardNetworkTokensRequestorId ensures that no value is present for MastercardNetworkTokensRequestorId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetMastercardNetworkTokensRequestorId() {
	o.MastercardNetworkTokensRequestorId.Unset()
}

// GetMastercardNetworkTokensAppId returns the MastercardNetworkTokensAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetMastercardNetworkTokensAppId() string {
	if o == nil || o.MastercardNetworkTokensAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.MastercardNetworkTokensAppId.Get()
}

// GetMastercardNetworkTokensAppIdOk returns a tuple with the MastercardNetworkTokensAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetMastercardNetworkTokensAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MastercardNetworkTokensAppId.Get(), o.MastercardNetworkTokensAppId.IsSet()
}

// HasMastercardNetworkTokensAppId returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasMastercardNetworkTokensAppId() bool {
	if o != nil && o.MastercardNetworkTokensAppId.IsSet() {
		return true
	}

	return false
}

// SetMastercardNetworkTokensAppId gets a reference to the given NullableString and assigns it to the MastercardNetworkTokensAppId field.
func (o *MerchantAccountCreate) SetMastercardNetworkTokensAppId(v string) {
	o.MastercardNetworkTokensAppId.Set(&v)
}
// SetMastercardNetworkTokensAppIdNil sets the value for MastercardNetworkTokensAppId to be an explicit nil
func (o *MerchantAccountCreate) SetMastercardNetworkTokensAppIdNil() {
	o.MastercardNetworkTokensAppId.Set(nil)
}

// UnsetMastercardNetworkTokensAppId ensures that no value is present for MastercardNetworkTokensAppId, not even an explicit nil
func (o *MerchantAccountCreate) UnsetMastercardNetworkTokensAppId() {
	o.MastercardNetworkTokensAppId.Unset()
}

// GetLoonClientKey returns the LoonClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetLoonClientKey() string {
	if o == nil || o.LoonClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoonClientKey.Get()
}

// GetLoonClientKeyOk returns a tuple with the LoonClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetLoonClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoonClientKey.Get(), o.LoonClientKey.IsSet()
}

// HasLoonClientKey returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasLoonClientKey() bool {
	if o != nil && o.LoonClientKey.IsSet() {
		return true
	}

	return false
}

// SetLoonClientKey gets a reference to the given NullableString and assigns it to the LoonClientKey field.
func (o *MerchantAccountCreate) SetLoonClientKey(v string) {
	o.LoonClientKey.Set(&v)
}
// SetLoonClientKeyNil sets the value for LoonClientKey to be an explicit nil
func (o *MerchantAccountCreate) SetLoonClientKeyNil() {
	o.LoonClientKey.Set(nil)
}

// UnsetLoonClientKey ensures that no value is present for LoonClientKey, not even an explicit nil
func (o *MerchantAccountCreate) UnsetLoonClientKey() {
	o.LoonClientKey.Unset()
}

// GetLoonSecretKey returns the LoonSecretKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetLoonSecretKey() string {
	if o == nil || o.LoonSecretKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoonSecretKey.Get()
}

// GetLoonSecretKeyOk returns a tuple with the LoonSecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetLoonSecretKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoonSecretKey.Get(), o.LoonSecretKey.IsSet()
}

// HasLoonSecretKey returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasLoonSecretKey() bool {
	if o != nil && o.LoonSecretKey.IsSet() {
		return true
	}

	return false
}

// SetLoonSecretKey gets a reference to the given NullableString and assigns it to the LoonSecretKey field.
func (o *MerchantAccountCreate) SetLoonSecretKey(v string) {
	o.LoonSecretKey.Set(&v)
}
// SetLoonSecretKeyNil sets the value for LoonSecretKey to be an explicit nil
func (o *MerchantAccountCreate) SetLoonSecretKeyNil() {
	o.LoonSecretKey.Set(nil)
}

// UnsetLoonSecretKey ensures that no value is present for LoonSecretKey, not even an explicit nil
func (o *MerchantAccountCreate) UnsetLoonSecretKey() {
	o.LoonSecretKey.Unset()
}

// GetLoonAcceptedSchemes returns the LoonAcceptedSchemes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetLoonAcceptedSchemes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.LoonAcceptedSchemes
}

// GetLoonAcceptedSchemesOk returns a tuple with the LoonAcceptedSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetLoonAcceptedSchemesOk() (*[]string, bool) {
	if o == nil || o.LoonAcceptedSchemes == nil {
		return nil, false
	}
	return &o.LoonAcceptedSchemes, true
}

// HasLoonAcceptedSchemes returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasLoonAcceptedSchemes() bool {
	if o != nil && o.LoonAcceptedSchemes != nil {
		return true
	}

	return false
}

// SetLoonAcceptedSchemes gets a reference to the given []string and assigns it to the LoonAcceptedSchemes field.
func (o *MerchantAccountCreate) SetLoonAcceptedSchemes(v []string) {
	o.LoonAcceptedSchemes = v
}

// GetOverCaptureAmount returns the OverCaptureAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetOverCaptureAmount() float32 {
	if o == nil || o.OverCaptureAmount.Get() == nil {
		var ret float32
		return ret
	}
	return *o.OverCaptureAmount.Get()
}

// GetOverCaptureAmountOk returns a tuple with the OverCaptureAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetOverCaptureAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverCaptureAmount.Get(), o.OverCaptureAmount.IsSet()
}

// HasOverCaptureAmount returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasOverCaptureAmount() bool {
	if o != nil && o.OverCaptureAmount.IsSet() {
		return true
	}

	return false
}

// SetOverCaptureAmount gets a reference to the given NullableFloat32 and assigns it to the OverCaptureAmount field.
func (o *MerchantAccountCreate) SetOverCaptureAmount(v float32) {
	o.OverCaptureAmount.Set(&v)
}
// SetOverCaptureAmountNil sets the value for OverCaptureAmount to be an explicit nil
func (o *MerchantAccountCreate) SetOverCaptureAmountNil() {
	o.OverCaptureAmount.Set(nil)
}

// UnsetOverCaptureAmount ensures that no value is present for OverCaptureAmount, not even an explicit nil
func (o *MerchantAccountCreate) UnsetOverCaptureAmount() {
	o.OverCaptureAmount.Unset()
}

// GetOverCapturePercentage returns the OverCapturePercentage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MerchantAccountCreate) GetOverCapturePercentage() float32 {
	if o == nil || o.OverCapturePercentage.Get() == nil {
		var ret float32
		return ret
	}
	return *o.OverCapturePercentage.Get()
}

// GetOverCapturePercentageOk returns a tuple with the OverCapturePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MerchantAccountCreate) GetOverCapturePercentageOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OverCapturePercentage.Get(), o.OverCapturePercentage.IsSet()
}

// HasOverCapturePercentage returns a boolean if a field has been set.
func (o *MerchantAccountCreate) HasOverCapturePercentage() bool {
	if o != nil && o.OverCapturePercentage.IsSet() {
		return true
	}

	return false
}

// SetOverCapturePercentage gets a reference to the given NullableFloat32 and assigns it to the OverCapturePercentage field.
func (o *MerchantAccountCreate) SetOverCapturePercentage(v float32) {
	o.OverCapturePercentage.Set(&v)
}
// SetOverCapturePercentageNil sets the value for OverCapturePercentage to be an explicit nil
func (o *MerchantAccountCreate) SetOverCapturePercentageNil() {
	o.OverCapturePercentage.Set(nil)
}

// UnsetOverCapturePercentage ensures that no value is present for OverCapturePercentage, not even an explicit nil
func (o *MerchantAccountCreate) UnsetOverCapturePercentage() {
	o.OverCapturePercentage.Unset()
}

func (o MerchantAccountCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.OutboundWebhookUrl.IsSet() {
		toSerialize["outbound_webhook_url"] = o.OutboundWebhookUrl.Get()
	}
	if o.OutboundWebhookUsername.IsSet() {
		toSerialize["outbound_webhook_username"] = o.OutboundWebhookUsername.Get()
	}
	if o.OutboundWebhookPassword.IsSet() {
		toSerialize["outbound_webhook_password"] = o.OutboundWebhookPassword.Get()
	}
	if o.VisaNetworkTokensRequestorId.IsSet() {
		toSerialize["visa_network_tokens_requestor_id"] = o.VisaNetworkTokensRequestorId.Get()
	}
	if o.VisaNetworkTokensAppId.IsSet() {
		toSerialize["visa_network_tokens_app_id"] = o.VisaNetworkTokensAppId.Get()
	}
	if o.AmexNetworkTokensRequestorId.IsSet() {
		toSerialize["amex_network_tokens_requestor_id"] = o.AmexNetworkTokensRequestorId.Get()
	}
	if o.AmexNetworkTokensAppId.IsSet() {
		toSerialize["amex_network_tokens_app_id"] = o.AmexNetworkTokensAppId.Get()
	}
	if o.MastercardNetworkTokensRequestorId.IsSet() {
		toSerialize["mastercard_network_tokens_requestor_id"] = o.MastercardNetworkTokensRequestorId.Get()
	}
	if o.MastercardNetworkTokensAppId.IsSet() {
		toSerialize["mastercard_network_tokens_app_id"] = o.MastercardNetworkTokensAppId.Get()
	}
	if o.LoonClientKey.IsSet() {
		toSerialize["loon_client_key"] = o.LoonClientKey.Get()
	}
	if o.LoonSecretKey.IsSet() {
		toSerialize["loon_secret_key"] = o.LoonSecretKey.Get()
	}
	if o.LoonAcceptedSchemes != nil {
		toSerialize["loon_accepted_schemes"] = o.LoonAcceptedSchemes
	}
	if o.OverCaptureAmount.IsSet() {
		toSerialize["over_capture_amount"] = o.OverCaptureAmount.Get()
	}
	if o.OverCapturePercentage.IsSet() {
		toSerialize["over_capture_percentage"] = o.OverCapturePercentage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMerchantAccountCreate struct {
	value *MerchantAccountCreate
	isSet bool
}

func (v NullableMerchantAccountCreate) Get() *MerchantAccountCreate {
	return v.value
}

func (v *NullableMerchantAccountCreate) Set(val *MerchantAccountCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantAccountCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantAccountCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantAccountCreate(val *MerchantAccountCreate) *NullableMerchantAccountCreate {
	return &NullableMerchantAccountCreate{value: val, isSet: true}
}

func (v NullableMerchantAccountCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantAccountCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


