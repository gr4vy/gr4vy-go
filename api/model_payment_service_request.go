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

// PaymentServiceRequest Request body for activating a payment service.
type PaymentServiceRequest struct {
	// A custom name for the payment service. This will be shown in the Admin UI.
	DisplayName string `json:"display_name"`
	// A list of fields, each containing a key-value pair for each field defined by the definition for this payment service e.g. for stripe-card `secret_key` is required and so must be sent within this field.
	Fields []PaymentServiceUpdateFields `json:"fields"`
	// A list of countries that this payment service needs to support in ISO two-letter code format.
	AcceptedCountries []string `json:"accepted_countries"`
	// A list of currencies that this payment service needs to support in ISO 4217 three-letter code format.
	AcceptedCurrencies []string `json:"accepted_currencies"`
	// Defines if 3-D Secure is enabled for the service (can only be enabled if the payment service definition supports the `three_d_secure_hosted` feature). This does not affect pass through 3-D Secure data.
	ThreeDSecureEnabled *bool `json:"three_d_secure_enabled,omitempty"`
	// Acquiring institution identification code for VISA.
	AcquirerBinVisa NullableString `json:"acquirer_bin_visa,omitempty"`
	// Acquiring institution identification code for Mastercard.
	AcquirerBinMastercard NullableString `json:"acquirer_bin_mastercard,omitempty"`
	// Acquiring institution identification code for Amex.
	AcquirerBinAmex NullableString `json:"acquirer_bin_amex,omitempty"`
	// Acquiring institution identification code for Discover.
	AcquirerBinDiscover NullableString `json:"acquirer_bin_discover,omitempty"`
	// Merchant identifier used in authorisation requests (assigned by the acquirer).
	AcquirerMerchantId NullableString `json:"acquirer_merchant_id,omitempty"`
	// Merchant name (assigned by the acquirer).
	MerchantName NullableString `json:"merchant_name,omitempty"`
	// ISO 3166-1 numeric three-digit country code.
	MerchantCountryCode NullableString `json:"merchant_country_code,omitempty"`
	// Merchant category code that describes the business.
	MerchantCategoryCode NullableString `json:"merchant_category_code,omitempty"`
	// Fully qualified URL of 3-D Secure requestor website or customer care site.
	MerchantUrl NullableString `json:"merchant_url,omitempty"`
	// Defines if this service is currently active or not.
	Active *bool `json:"active,omitempty"`
	// The numeric rank of a payment service. Payment services with a lower position value are processed first. When a payment services is inserted at a position, any payment services with the the same value or higher are shifted down a position accordingly. When left out, the payment service is inserted at the end of the list.
	Position *float32 `json:"position,omitempty"`
	// The ID of the payment service to use.
	PaymentServiceDefinitionId string `json:"payment_service_definition_id"`
}

// NewPaymentServiceRequest instantiates a new PaymentServiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceRequest(displayName string, fields []PaymentServiceUpdateFields, acceptedCountries []string, acceptedCurrencies []string, paymentServiceDefinitionId string) *PaymentServiceRequest {
	this := PaymentServiceRequest{}
	this.DisplayName = displayName
	this.Fields = fields
	this.AcceptedCountries = acceptedCountries
	this.AcceptedCurrencies = acceptedCurrencies
	var threeDSecureEnabled bool = false
	this.ThreeDSecureEnabled = &threeDSecureEnabled
	var active bool = true
	this.Active = &active
	this.PaymentServiceDefinitionId = paymentServiceDefinitionId
	return &this
}

// NewPaymentServiceRequestWithDefaults instantiates a new PaymentServiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceRequestWithDefaults() *PaymentServiceRequest {
	this := PaymentServiceRequest{}
	var threeDSecureEnabled bool = false
	this.ThreeDSecureEnabled = &threeDSecureEnabled
	var active bool = true
	this.Active = &active
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *PaymentServiceRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *PaymentServiceRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetFields returns the Fields field value
func (o *PaymentServiceRequest) GetFields() []PaymentServiceUpdateFields {
	if o == nil {
		var ret []PaymentServiceUpdateFields
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetFieldsOk() (*[]PaymentServiceUpdateFields, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value
func (o *PaymentServiceRequest) SetFields(v []PaymentServiceUpdateFields) {
	o.Fields = v
}

// GetAcceptedCountries returns the AcceptedCountries field value
func (o *PaymentServiceRequest) GetAcceptedCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcceptedCountries
}

// GetAcceptedCountriesOk returns a tuple with the AcceptedCountries field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetAcceptedCountriesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcceptedCountries, true
}

// SetAcceptedCountries sets field value
func (o *PaymentServiceRequest) SetAcceptedCountries(v []string) {
	o.AcceptedCountries = v
}

// GetAcceptedCurrencies returns the AcceptedCurrencies field value
func (o *PaymentServiceRequest) GetAcceptedCurrencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcceptedCurrencies
}

// GetAcceptedCurrenciesOk returns a tuple with the AcceptedCurrencies field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetAcceptedCurrenciesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AcceptedCurrencies, true
}

// SetAcceptedCurrencies sets field value
func (o *PaymentServiceRequest) SetAcceptedCurrencies(v []string) {
	o.AcceptedCurrencies = v
}

// GetThreeDSecureEnabled returns the ThreeDSecureEnabled field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetThreeDSecureEnabled() bool {
	if o == nil || o.ThreeDSecureEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ThreeDSecureEnabled
}

// GetThreeDSecureEnabledOk returns a tuple with the ThreeDSecureEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetThreeDSecureEnabledOk() (*bool, bool) {
	if o == nil || o.ThreeDSecureEnabled == nil {
		return nil, false
	}
	return o.ThreeDSecureEnabled, true
}

// HasThreeDSecureEnabled returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasThreeDSecureEnabled() bool {
	if o != nil && o.ThreeDSecureEnabled != nil {
		return true
	}

	return false
}

// SetThreeDSecureEnabled gets a reference to the given bool and assigns it to the ThreeDSecureEnabled field.
func (o *PaymentServiceRequest) SetThreeDSecureEnabled(v bool) {
	o.ThreeDSecureEnabled = &v
}

// GetAcquirerBinVisa returns the AcquirerBinVisa field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetAcquirerBinVisa() string {
	if o == nil || o.AcquirerBinVisa.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcquirerBinVisa.Get()
}

// GetAcquirerBinVisaOk returns a tuple with the AcquirerBinVisa field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetAcquirerBinVisaOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcquirerBinVisa.Get(), o.AcquirerBinVisa.IsSet()
}

// HasAcquirerBinVisa returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasAcquirerBinVisa() bool {
	if o != nil && o.AcquirerBinVisa.IsSet() {
		return true
	}

	return false
}

// SetAcquirerBinVisa gets a reference to the given NullableString and assigns it to the AcquirerBinVisa field.
func (o *PaymentServiceRequest) SetAcquirerBinVisa(v string) {
	o.AcquirerBinVisa.Set(&v)
}
// SetAcquirerBinVisaNil sets the value for AcquirerBinVisa to be an explicit nil
func (o *PaymentServiceRequest) SetAcquirerBinVisaNil() {
	o.AcquirerBinVisa.Set(nil)
}

// UnsetAcquirerBinVisa ensures that no value is present for AcquirerBinVisa, not even an explicit nil
func (o *PaymentServiceRequest) UnsetAcquirerBinVisa() {
	o.AcquirerBinVisa.Unset()
}

// GetAcquirerBinMastercard returns the AcquirerBinMastercard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetAcquirerBinMastercard() string {
	if o == nil || o.AcquirerBinMastercard.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcquirerBinMastercard.Get()
}

// GetAcquirerBinMastercardOk returns a tuple with the AcquirerBinMastercard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetAcquirerBinMastercardOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcquirerBinMastercard.Get(), o.AcquirerBinMastercard.IsSet()
}

// HasAcquirerBinMastercard returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasAcquirerBinMastercard() bool {
	if o != nil && o.AcquirerBinMastercard.IsSet() {
		return true
	}

	return false
}

// SetAcquirerBinMastercard gets a reference to the given NullableString and assigns it to the AcquirerBinMastercard field.
func (o *PaymentServiceRequest) SetAcquirerBinMastercard(v string) {
	o.AcquirerBinMastercard.Set(&v)
}
// SetAcquirerBinMastercardNil sets the value for AcquirerBinMastercard to be an explicit nil
func (o *PaymentServiceRequest) SetAcquirerBinMastercardNil() {
	o.AcquirerBinMastercard.Set(nil)
}

// UnsetAcquirerBinMastercard ensures that no value is present for AcquirerBinMastercard, not even an explicit nil
func (o *PaymentServiceRequest) UnsetAcquirerBinMastercard() {
	o.AcquirerBinMastercard.Unset()
}

// GetAcquirerBinAmex returns the AcquirerBinAmex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetAcquirerBinAmex() string {
	if o == nil || o.AcquirerBinAmex.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcquirerBinAmex.Get()
}

// GetAcquirerBinAmexOk returns a tuple with the AcquirerBinAmex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetAcquirerBinAmexOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcquirerBinAmex.Get(), o.AcquirerBinAmex.IsSet()
}

// HasAcquirerBinAmex returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasAcquirerBinAmex() bool {
	if o != nil && o.AcquirerBinAmex.IsSet() {
		return true
	}

	return false
}

// SetAcquirerBinAmex gets a reference to the given NullableString and assigns it to the AcquirerBinAmex field.
func (o *PaymentServiceRequest) SetAcquirerBinAmex(v string) {
	o.AcquirerBinAmex.Set(&v)
}
// SetAcquirerBinAmexNil sets the value for AcquirerBinAmex to be an explicit nil
func (o *PaymentServiceRequest) SetAcquirerBinAmexNil() {
	o.AcquirerBinAmex.Set(nil)
}

// UnsetAcquirerBinAmex ensures that no value is present for AcquirerBinAmex, not even an explicit nil
func (o *PaymentServiceRequest) UnsetAcquirerBinAmex() {
	o.AcquirerBinAmex.Unset()
}

// GetAcquirerBinDiscover returns the AcquirerBinDiscover field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetAcquirerBinDiscover() string {
	if o == nil || o.AcquirerBinDiscover.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcquirerBinDiscover.Get()
}

// GetAcquirerBinDiscoverOk returns a tuple with the AcquirerBinDiscover field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetAcquirerBinDiscoverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcquirerBinDiscover.Get(), o.AcquirerBinDiscover.IsSet()
}

// HasAcquirerBinDiscover returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasAcquirerBinDiscover() bool {
	if o != nil && o.AcquirerBinDiscover.IsSet() {
		return true
	}

	return false
}

// SetAcquirerBinDiscover gets a reference to the given NullableString and assigns it to the AcquirerBinDiscover field.
func (o *PaymentServiceRequest) SetAcquirerBinDiscover(v string) {
	o.AcquirerBinDiscover.Set(&v)
}
// SetAcquirerBinDiscoverNil sets the value for AcquirerBinDiscover to be an explicit nil
func (o *PaymentServiceRequest) SetAcquirerBinDiscoverNil() {
	o.AcquirerBinDiscover.Set(nil)
}

// UnsetAcquirerBinDiscover ensures that no value is present for AcquirerBinDiscover, not even an explicit nil
func (o *PaymentServiceRequest) UnsetAcquirerBinDiscover() {
	o.AcquirerBinDiscover.Unset()
}

// GetAcquirerMerchantId returns the AcquirerMerchantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetAcquirerMerchantId() string {
	if o == nil || o.AcquirerMerchantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcquirerMerchantId.Get()
}

// GetAcquirerMerchantIdOk returns a tuple with the AcquirerMerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetAcquirerMerchantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcquirerMerchantId.Get(), o.AcquirerMerchantId.IsSet()
}

// HasAcquirerMerchantId returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasAcquirerMerchantId() bool {
	if o != nil && o.AcquirerMerchantId.IsSet() {
		return true
	}

	return false
}

// SetAcquirerMerchantId gets a reference to the given NullableString and assigns it to the AcquirerMerchantId field.
func (o *PaymentServiceRequest) SetAcquirerMerchantId(v string) {
	o.AcquirerMerchantId.Set(&v)
}
// SetAcquirerMerchantIdNil sets the value for AcquirerMerchantId to be an explicit nil
func (o *PaymentServiceRequest) SetAcquirerMerchantIdNil() {
	o.AcquirerMerchantId.Set(nil)
}

// UnsetAcquirerMerchantId ensures that no value is present for AcquirerMerchantId, not even an explicit nil
func (o *PaymentServiceRequest) UnsetAcquirerMerchantId() {
	o.AcquirerMerchantId.Unset()
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// HasMerchantName returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasMerchantName() bool {
	if o != nil && o.MerchantName.IsSet() {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given NullableString and assigns it to the MerchantName field.
func (o *PaymentServiceRequest) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}
// SetMerchantNameNil sets the value for MerchantName to be an explicit nil
func (o *PaymentServiceRequest) SetMerchantNameNil() {
	o.MerchantName.Set(nil)
}

// UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
func (o *PaymentServiceRequest) UnsetMerchantName() {
	o.MerchantName.Unset()
}

// GetMerchantCountryCode returns the MerchantCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetMerchantCountryCode() string {
	if o == nil || o.MerchantCountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantCountryCode.Get()
}

// GetMerchantCountryCodeOk returns a tuple with the MerchantCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetMerchantCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantCountryCode.Get(), o.MerchantCountryCode.IsSet()
}

// HasMerchantCountryCode returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasMerchantCountryCode() bool {
	if o != nil && o.MerchantCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMerchantCountryCode gets a reference to the given NullableString and assigns it to the MerchantCountryCode field.
func (o *PaymentServiceRequest) SetMerchantCountryCode(v string) {
	o.MerchantCountryCode.Set(&v)
}
// SetMerchantCountryCodeNil sets the value for MerchantCountryCode to be an explicit nil
func (o *PaymentServiceRequest) SetMerchantCountryCodeNil() {
	o.MerchantCountryCode.Set(nil)
}

// UnsetMerchantCountryCode ensures that no value is present for MerchantCountryCode, not even an explicit nil
func (o *PaymentServiceRequest) UnsetMerchantCountryCode() {
	o.MerchantCountryCode.Unset()
}

// GetMerchantCategoryCode returns the MerchantCategoryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetMerchantCategoryCode() string {
	if o == nil || o.MerchantCategoryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantCategoryCode.Get()
}

// GetMerchantCategoryCodeOk returns a tuple with the MerchantCategoryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetMerchantCategoryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantCategoryCode.Get(), o.MerchantCategoryCode.IsSet()
}

// HasMerchantCategoryCode returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasMerchantCategoryCode() bool {
	if o != nil && o.MerchantCategoryCode.IsSet() {
		return true
	}

	return false
}

// SetMerchantCategoryCode gets a reference to the given NullableString and assigns it to the MerchantCategoryCode field.
func (o *PaymentServiceRequest) SetMerchantCategoryCode(v string) {
	o.MerchantCategoryCode.Set(&v)
}
// SetMerchantCategoryCodeNil sets the value for MerchantCategoryCode to be an explicit nil
func (o *PaymentServiceRequest) SetMerchantCategoryCodeNil() {
	o.MerchantCategoryCode.Set(nil)
}

// UnsetMerchantCategoryCode ensures that no value is present for MerchantCategoryCode, not even an explicit nil
func (o *PaymentServiceRequest) UnsetMerchantCategoryCode() {
	o.MerchantCategoryCode.Unset()
}

// GetMerchantUrl returns the MerchantUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentServiceRequest) GetMerchantUrl() string {
	if o == nil || o.MerchantUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantUrl.Get()
}

// GetMerchantUrlOk returns a tuple with the MerchantUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentServiceRequest) GetMerchantUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantUrl.Get(), o.MerchantUrl.IsSet()
}

// HasMerchantUrl returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasMerchantUrl() bool {
	if o != nil && o.MerchantUrl.IsSet() {
		return true
	}

	return false
}

// SetMerchantUrl gets a reference to the given NullableString and assigns it to the MerchantUrl field.
func (o *PaymentServiceRequest) SetMerchantUrl(v string) {
	o.MerchantUrl.Set(&v)
}
// SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil
func (o *PaymentServiceRequest) SetMerchantUrlNil() {
	o.MerchantUrl.Set(nil)
}

// UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
func (o *PaymentServiceRequest) UnsetMerchantUrl() {
	o.MerchantUrl.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PaymentServiceRequest) SetActive(v bool) {
	o.Active = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *PaymentServiceRequest) GetPosition() float32 {
	if o == nil || o.Position == nil {
		var ret float32
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetPositionOk() (*float32, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *PaymentServiceRequest) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given float32 and assigns it to the Position field.
func (o *PaymentServiceRequest) SetPosition(v float32) {
	o.Position = &v
}

// GetPaymentServiceDefinitionId returns the PaymentServiceDefinitionId field value
func (o *PaymentServiceRequest) GetPaymentServiceDefinitionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentServiceDefinitionId
}

// GetPaymentServiceDefinitionIdOk returns a tuple with the PaymentServiceDefinitionId field value
// and a boolean to check if the value has been set.
func (o *PaymentServiceRequest) GetPaymentServiceDefinitionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentServiceDefinitionId, true
}

// SetPaymentServiceDefinitionId sets field value
func (o *PaymentServiceRequest) SetPaymentServiceDefinitionId(v string) {
	o.PaymentServiceDefinitionId = v
}

func (o PaymentServiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["fields"] = o.Fields
	}
	if true {
		toSerialize["accepted_countries"] = o.AcceptedCountries
	}
	if true {
		toSerialize["accepted_currencies"] = o.AcceptedCurrencies
	}
	if o.ThreeDSecureEnabled != nil {
		toSerialize["three_d_secure_enabled"] = o.ThreeDSecureEnabled
	}
	if o.AcquirerBinVisa.IsSet() {
		toSerialize["acquirer_bin_visa"] = o.AcquirerBinVisa.Get()
	}
	if o.AcquirerBinMastercard.IsSet() {
		toSerialize["acquirer_bin_mastercard"] = o.AcquirerBinMastercard.Get()
	}
	if o.AcquirerBinAmex.IsSet() {
		toSerialize["acquirer_bin_amex"] = o.AcquirerBinAmex.Get()
	}
	if o.AcquirerBinDiscover.IsSet() {
		toSerialize["acquirer_bin_discover"] = o.AcquirerBinDiscover.Get()
	}
	if o.AcquirerMerchantId.IsSet() {
		toSerialize["acquirer_merchant_id"] = o.AcquirerMerchantId.Get()
	}
	if o.MerchantName.IsSet() {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if o.MerchantCountryCode.IsSet() {
		toSerialize["merchant_country_code"] = o.MerchantCountryCode.Get()
	}
	if o.MerchantCategoryCode.IsSet() {
		toSerialize["merchant_category_code"] = o.MerchantCategoryCode.Get()
	}
	if o.MerchantUrl.IsSet() {
		toSerialize["merchant_url"] = o.MerchantUrl.Get()
	}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}
	if true {
		toSerialize["payment_service_definition_id"] = o.PaymentServiceDefinitionId
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentServiceRequest struct {
	value *PaymentServiceRequest
	isSet bool
}

func (v NullablePaymentServiceRequest) Get() *PaymentServiceRequest {
	return v.value
}

func (v *NullablePaymentServiceRequest) Set(val *PaymentServiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceRequest(val *PaymentServiceRequest) *NullablePaymentServiceRequest {
	return &NullablePaymentServiceRequest{value: val, isSet: true}
}

func (v NullablePaymentServiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


