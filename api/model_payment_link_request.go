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
	"time"
)

// PaymentLinkRequest Request body for creating a new payment link.
type PaymentLinkRequest struct {
	// The amount to request payment for.
	Amount float32 `json:"amount"`
	// The ISO-4217 currency code for the payment.
	Currency string `json:"currency"`
	// The 2-letter ISO code of the country of the transaction. This is used to filter the payment services that is used to process the transaction. 
	Country string `json:"country"`
	Buyer *TransactionBuyerRequest `json:"buyer,omitempty"`
	// The date and time when this payment link expires. Defaults to 24 hours from creation.
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// Allows for passing optional configuration per connection to take advantage of connection specific features. When provided, the data is only passed to the target connection type to prevent sharing configuration across connections.  Please note that each of the keys this object are in kebab-case, for example `cybersource-anti-fraud` as they represent the ID of the connector. All the other keys will be snake case, for example `merchant_defined_data` or camel case to match an external API that the connector uses.
	ConnectionOptions NullableConnectionOptions `json:"connection_options,omitempty"`
	// A value that can be used to match the payment link against your own records. This will also be applied to any transaction.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	StatementDescriptor NullableStatementDescriptor `json:"statement_descriptor,omitempty"`
	// The locale used to translate text within the payment link.
	Locale NullableString `json:"locale,omitempty"`
	// The name of the merchant to display on the payment link.
	MerchantName NullableString `json:"merchant_name,omitempty"`
	// The URL of the merchant to display on the payment link.
	MerchantUrl NullableString `json:"merchant_url,omitempty"`
	// The URL of the merchant banner to display on the payment link.
	MerchantBannerUrl NullableString `json:"merchant_banner_url,omitempty"`
	// The color code of the merchant to display on the payment link.
	MerchantColor NullableString `json:"merchant_color,omitempty"`
	// The message to display on the payment link.
	MerchantMessage NullableString `json:"merchant_message,omitempty"`
	// The URL of the merchant terms and conditions to display on the payment link.
	MerchantTermsAndConditionsUrl NullableString `json:"merchant_terms_and_conditions_url,omitempty"`
	// The URL of the merchant favicon icon.
	MerchantFaviconUrl NullableString `json:"merchant_favicon_url,omitempty"`
	// The intent of the payment link.
	Intent NullableString `json:"intent,omitempty"`
	// The URL to redirect the buyer to after payment.
	ReturnUrl NullableString `json:"return_url,omitempty"`
	// An array of cart items that represents the line items of a payment link.
	CartItems *[]CartItem `json:"cart_items,omitempty"`
	// Any additional information about the payment link that you would like to store as key-value pairs. This data is passed to payment service providers that support it.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The source of the payment link.
	PaymentSource NullableString `json:"payment_source,omitempty"`
}

// NewPaymentLinkRequest instantiates a new PaymentLinkRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentLinkRequest(amount float32, currency string, country string) *PaymentLinkRequest {
	this := PaymentLinkRequest{}
	this.Amount = amount
	this.Currency = currency
	this.Country = country
	var intent string = "authorize"
	this.Intent = *NewNullableString(&intent)
	return &this
}

// NewPaymentLinkRequestWithDefaults instantiates a new PaymentLinkRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentLinkRequestWithDefaults() *PaymentLinkRequest {
	this := PaymentLinkRequest{}
	var intent string = "authorize"
	this.Intent = *NewNullableString(&intent)
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentLinkRequest) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRequest) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentLinkRequest) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *PaymentLinkRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRequest) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *PaymentLinkRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetCountry returns the Country field value
func (o *PaymentLinkRequest) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PaymentLinkRequest) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PaymentLinkRequest) SetCountry(v string) {
	o.Country = v
}

// GetBuyer returns the Buyer field value if set, zero value otherwise.
func (o *PaymentLinkRequest) GetBuyer() TransactionBuyerRequest {
	if o == nil || o.Buyer == nil {
		var ret TransactionBuyerRequest
		return ret
	}
	return *o.Buyer
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkRequest) GetBuyerOk() (*TransactionBuyerRequest, bool) {
	if o == nil || o.Buyer == nil {
		return nil, false
	}
	return o.Buyer, true
}

// HasBuyer returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasBuyer() bool {
	if o != nil && o.Buyer != nil {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given TransactionBuyerRequest and assigns it to the Buyer field.
func (o *PaymentLinkRequest) SetBuyer(v TransactionBuyerRequest) {
	o.Buyer = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *PaymentLinkRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}
// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *PaymentLinkRequest) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *PaymentLinkRequest) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetConnectionOptions returns the ConnectionOptions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetConnectionOptions() ConnectionOptions {
	if o == nil || o.ConnectionOptions.Get() == nil {
		var ret ConnectionOptions
		return ret
	}
	return *o.ConnectionOptions.Get()
}

// GetConnectionOptionsOk returns a tuple with the ConnectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetConnectionOptionsOk() (*ConnectionOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConnectionOptions.Get(), o.ConnectionOptions.IsSet()
}

// HasConnectionOptions returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasConnectionOptions() bool {
	if o != nil && o.ConnectionOptions.IsSet() {
		return true
	}

	return false
}

// SetConnectionOptions gets a reference to the given NullableConnectionOptions and assigns it to the ConnectionOptions field.
func (o *PaymentLinkRequest) SetConnectionOptions(v ConnectionOptions) {
	o.ConnectionOptions.Set(&v)
}
// SetConnectionOptionsNil sets the value for ConnectionOptions to be an explicit nil
func (o *PaymentLinkRequest) SetConnectionOptionsNil() {
	o.ConnectionOptions.Set(nil)
}

// UnsetConnectionOptions ensures that no value is present for ConnectionOptions, not even an explicit nil
func (o *PaymentLinkRequest) UnsetConnectionOptions() {
	o.ConnectionOptions.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PaymentLinkRequest) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PaymentLinkRequest) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PaymentLinkRequest) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetStatementDescriptor returns the StatementDescriptor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetStatementDescriptor() StatementDescriptor {
	if o == nil || o.StatementDescriptor.Get() == nil {
		var ret StatementDescriptor
		return ret
	}
	return *o.StatementDescriptor.Get()
}

// GetStatementDescriptorOk returns a tuple with the StatementDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetStatementDescriptorOk() (*StatementDescriptor, bool) {
	if o == nil  {
		return nil, false
	}
	return o.StatementDescriptor.Get(), o.StatementDescriptor.IsSet()
}

// HasStatementDescriptor returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasStatementDescriptor() bool {
	if o != nil && o.StatementDescriptor.IsSet() {
		return true
	}

	return false
}

// SetStatementDescriptor gets a reference to the given NullableStatementDescriptor and assigns it to the StatementDescriptor field.
func (o *PaymentLinkRequest) SetStatementDescriptor(v StatementDescriptor) {
	o.StatementDescriptor.Set(&v)
}
// SetStatementDescriptorNil sets the value for StatementDescriptor to be an explicit nil
func (o *PaymentLinkRequest) SetStatementDescriptorNil() {
	o.StatementDescriptor.Set(nil)
}

// UnsetStatementDescriptor ensures that no value is present for StatementDescriptor, not even an explicit nil
func (o *PaymentLinkRequest) UnsetStatementDescriptor() {
	o.StatementDescriptor.Unset()
}

// GetLocale returns the Locale field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetLocale() string {
	if o == nil || o.Locale.Get() == nil {
		var ret string
		return ret
	}
	return *o.Locale.Get()
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetLocaleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locale.Get(), o.Locale.IsSet()
}

// HasLocale returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasLocale() bool {
	if o != nil && o.Locale.IsSet() {
		return true
	}

	return false
}

// SetLocale gets a reference to the given NullableString and assigns it to the Locale field.
func (o *PaymentLinkRequest) SetLocale(v string) {
	o.Locale.Set(&v)
}
// SetLocaleNil sets the value for Locale to be an explicit nil
func (o *PaymentLinkRequest) SetLocaleNil() {
	o.Locale.Set(nil)
}

// UnsetLocale ensures that no value is present for Locale, not even an explicit nil
func (o *PaymentLinkRequest) UnsetLocale() {
	o.Locale.Unset()
}

// GetMerchantName returns the MerchantName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantName() string {
	if o == nil || o.MerchantName.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantName.Get()
}

// GetMerchantNameOk returns a tuple with the MerchantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantName.Get(), o.MerchantName.IsSet()
}

// HasMerchantName returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantName() bool {
	if o != nil && o.MerchantName.IsSet() {
		return true
	}

	return false
}

// SetMerchantName gets a reference to the given NullableString and assigns it to the MerchantName field.
func (o *PaymentLinkRequest) SetMerchantName(v string) {
	o.MerchantName.Set(&v)
}
// SetMerchantNameNil sets the value for MerchantName to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantNameNil() {
	o.MerchantName.Set(nil)
}

// UnsetMerchantName ensures that no value is present for MerchantName, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantName() {
	o.MerchantName.Unset()
}

// GetMerchantUrl returns the MerchantUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantUrl() string {
	if o == nil || o.MerchantUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantUrl.Get()
}

// GetMerchantUrlOk returns a tuple with the MerchantUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantUrl.Get(), o.MerchantUrl.IsSet()
}

// HasMerchantUrl returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantUrl() bool {
	if o != nil && o.MerchantUrl.IsSet() {
		return true
	}

	return false
}

// SetMerchantUrl gets a reference to the given NullableString and assigns it to the MerchantUrl field.
func (o *PaymentLinkRequest) SetMerchantUrl(v string) {
	o.MerchantUrl.Set(&v)
}
// SetMerchantUrlNil sets the value for MerchantUrl to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantUrlNil() {
	o.MerchantUrl.Set(nil)
}

// UnsetMerchantUrl ensures that no value is present for MerchantUrl, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantUrl() {
	o.MerchantUrl.Unset()
}

// GetMerchantBannerUrl returns the MerchantBannerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantBannerUrl() string {
	if o == nil || o.MerchantBannerUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantBannerUrl.Get()
}

// GetMerchantBannerUrlOk returns a tuple with the MerchantBannerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantBannerUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantBannerUrl.Get(), o.MerchantBannerUrl.IsSet()
}

// HasMerchantBannerUrl returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantBannerUrl() bool {
	if o != nil && o.MerchantBannerUrl.IsSet() {
		return true
	}

	return false
}

// SetMerchantBannerUrl gets a reference to the given NullableString and assigns it to the MerchantBannerUrl field.
func (o *PaymentLinkRequest) SetMerchantBannerUrl(v string) {
	o.MerchantBannerUrl.Set(&v)
}
// SetMerchantBannerUrlNil sets the value for MerchantBannerUrl to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantBannerUrlNil() {
	o.MerchantBannerUrl.Set(nil)
}

// UnsetMerchantBannerUrl ensures that no value is present for MerchantBannerUrl, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantBannerUrl() {
	o.MerchantBannerUrl.Unset()
}

// GetMerchantColor returns the MerchantColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantColor() string {
	if o == nil || o.MerchantColor.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantColor.Get()
}

// GetMerchantColorOk returns a tuple with the MerchantColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantColor.Get(), o.MerchantColor.IsSet()
}

// HasMerchantColor returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantColor() bool {
	if o != nil && o.MerchantColor.IsSet() {
		return true
	}

	return false
}

// SetMerchantColor gets a reference to the given NullableString and assigns it to the MerchantColor field.
func (o *PaymentLinkRequest) SetMerchantColor(v string) {
	o.MerchantColor.Set(&v)
}
// SetMerchantColorNil sets the value for MerchantColor to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantColorNil() {
	o.MerchantColor.Set(nil)
}

// UnsetMerchantColor ensures that no value is present for MerchantColor, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantColor() {
	o.MerchantColor.Unset()
}

// GetMerchantMessage returns the MerchantMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantMessage() string {
	if o == nil || o.MerchantMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantMessage.Get()
}

// GetMerchantMessageOk returns a tuple with the MerchantMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantMessage.Get(), o.MerchantMessage.IsSet()
}

// HasMerchantMessage returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantMessage() bool {
	if o != nil && o.MerchantMessage.IsSet() {
		return true
	}

	return false
}

// SetMerchantMessage gets a reference to the given NullableString and assigns it to the MerchantMessage field.
func (o *PaymentLinkRequest) SetMerchantMessage(v string) {
	o.MerchantMessage.Set(&v)
}
// SetMerchantMessageNil sets the value for MerchantMessage to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantMessageNil() {
	o.MerchantMessage.Set(nil)
}

// UnsetMerchantMessage ensures that no value is present for MerchantMessage, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantMessage() {
	o.MerchantMessage.Unset()
}

// GetMerchantTermsAndConditionsUrl returns the MerchantTermsAndConditionsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantTermsAndConditionsUrl() string {
	if o == nil || o.MerchantTermsAndConditionsUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantTermsAndConditionsUrl.Get()
}

// GetMerchantTermsAndConditionsUrlOk returns a tuple with the MerchantTermsAndConditionsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantTermsAndConditionsUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantTermsAndConditionsUrl.Get(), o.MerchantTermsAndConditionsUrl.IsSet()
}

// HasMerchantTermsAndConditionsUrl returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantTermsAndConditionsUrl() bool {
	if o != nil && o.MerchantTermsAndConditionsUrl.IsSet() {
		return true
	}

	return false
}

// SetMerchantTermsAndConditionsUrl gets a reference to the given NullableString and assigns it to the MerchantTermsAndConditionsUrl field.
func (o *PaymentLinkRequest) SetMerchantTermsAndConditionsUrl(v string) {
	o.MerchantTermsAndConditionsUrl.Set(&v)
}
// SetMerchantTermsAndConditionsUrlNil sets the value for MerchantTermsAndConditionsUrl to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantTermsAndConditionsUrlNil() {
	o.MerchantTermsAndConditionsUrl.Set(nil)
}

// UnsetMerchantTermsAndConditionsUrl ensures that no value is present for MerchantTermsAndConditionsUrl, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantTermsAndConditionsUrl() {
	o.MerchantTermsAndConditionsUrl.Unset()
}

// GetMerchantFaviconUrl returns the MerchantFaviconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMerchantFaviconUrl() string {
	if o == nil || o.MerchantFaviconUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.MerchantFaviconUrl.Get()
}

// GetMerchantFaviconUrlOk returns a tuple with the MerchantFaviconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMerchantFaviconUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MerchantFaviconUrl.Get(), o.MerchantFaviconUrl.IsSet()
}

// HasMerchantFaviconUrl returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMerchantFaviconUrl() bool {
	if o != nil && o.MerchantFaviconUrl.IsSet() {
		return true
	}

	return false
}

// SetMerchantFaviconUrl gets a reference to the given NullableString and assigns it to the MerchantFaviconUrl field.
func (o *PaymentLinkRequest) SetMerchantFaviconUrl(v string) {
	o.MerchantFaviconUrl.Set(&v)
}
// SetMerchantFaviconUrlNil sets the value for MerchantFaviconUrl to be an explicit nil
func (o *PaymentLinkRequest) SetMerchantFaviconUrlNil() {
	o.MerchantFaviconUrl.Set(nil)
}

// UnsetMerchantFaviconUrl ensures that no value is present for MerchantFaviconUrl, not even an explicit nil
func (o *PaymentLinkRequest) UnsetMerchantFaviconUrl() {
	o.MerchantFaviconUrl.Unset()
}

// GetIntent returns the Intent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetIntent() string {
	if o == nil || o.Intent.Get() == nil {
		var ret string
		return ret
	}
	return *o.Intent.Get()
}

// GetIntentOk returns a tuple with the Intent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetIntentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Intent.Get(), o.Intent.IsSet()
}

// HasIntent returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasIntent() bool {
	if o != nil && o.Intent.IsSet() {
		return true
	}

	return false
}

// SetIntent gets a reference to the given NullableString and assigns it to the Intent field.
func (o *PaymentLinkRequest) SetIntent(v string) {
	o.Intent.Set(&v)
}
// SetIntentNil sets the value for Intent to be an explicit nil
func (o *PaymentLinkRequest) SetIntentNil() {
	o.Intent.Set(nil)
}

// UnsetIntent ensures that no value is present for Intent, not even an explicit nil
func (o *PaymentLinkRequest) UnsetIntent() {
	o.Intent.Unset()
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetReturnUrl() string {
	if o == nil || o.ReturnUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ReturnUrl.Get()
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetReturnUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReturnUrl.Get(), o.ReturnUrl.IsSet()
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasReturnUrl() bool {
	if o != nil && o.ReturnUrl.IsSet() {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given NullableString and assigns it to the ReturnUrl field.
func (o *PaymentLinkRequest) SetReturnUrl(v string) {
	o.ReturnUrl.Set(&v)
}
// SetReturnUrlNil sets the value for ReturnUrl to be an explicit nil
func (o *PaymentLinkRequest) SetReturnUrlNil() {
	o.ReturnUrl.Set(nil)
}

// UnsetReturnUrl ensures that no value is present for ReturnUrl, not even an explicit nil
func (o *PaymentLinkRequest) UnsetReturnUrl() {
	o.ReturnUrl.Unset()
}

// GetCartItems returns the CartItems field value if set, zero value otherwise.
func (o *PaymentLinkRequest) GetCartItems() []CartItem {
	if o == nil || o.CartItems == nil {
		var ret []CartItem
		return ret
	}
	return *o.CartItems
}

// GetCartItemsOk returns a tuple with the CartItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentLinkRequest) GetCartItemsOk() (*[]CartItem, bool) {
	if o == nil || o.CartItems == nil {
		return nil, false
	}
	return o.CartItems, true
}

// HasCartItems returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasCartItems() bool {
	if o != nil && o.CartItems != nil {
		return true
	}

	return false
}

// SetCartItems gets a reference to the given []CartItem and assigns it to the CartItems field.
func (o *PaymentLinkRequest) SetCartItems(v []CartItem) {
	o.CartItems = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetMetadata() map[string]string {
	if o == nil  {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *PaymentLinkRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetPaymentSource returns the PaymentSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentLinkRequest) GetPaymentSource() string {
	if o == nil || o.PaymentSource.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentSource.Get()
}

// GetPaymentSourceOk returns a tuple with the PaymentSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentLinkRequest) GetPaymentSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentSource.Get(), o.PaymentSource.IsSet()
}

// HasPaymentSource returns a boolean if a field has been set.
func (o *PaymentLinkRequest) HasPaymentSource() bool {
	if o != nil && o.PaymentSource.IsSet() {
		return true
	}

	return false
}

// SetPaymentSource gets a reference to the given NullableString and assigns it to the PaymentSource field.
func (o *PaymentLinkRequest) SetPaymentSource(v string) {
	o.PaymentSource.Set(&v)
}
// SetPaymentSourceNil sets the value for PaymentSource to be an explicit nil
func (o *PaymentLinkRequest) SetPaymentSourceNil() {
	o.PaymentSource.Set(nil)
}

// UnsetPaymentSource ensures that no value is present for PaymentSource, not even an explicit nil
func (o *PaymentLinkRequest) UnsetPaymentSource() {
	o.PaymentSource.Unset()
}

func (o PaymentLinkRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["country"] = o.Country
	}
	if o.Buyer != nil {
		toSerialize["buyer"] = o.Buyer
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.ConnectionOptions.IsSet() {
		toSerialize["connection_options"] = o.ConnectionOptions.Get()
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.StatementDescriptor.IsSet() {
		toSerialize["statement_descriptor"] = o.StatementDescriptor.Get()
	}
	if o.Locale.IsSet() {
		toSerialize["locale"] = o.Locale.Get()
	}
	if o.MerchantName.IsSet() {
		toSerialize["merchant_name"] = o.MerchantName.Get()
	}
	if o.MerchantUrl.IsSet() {
		toSerialize["merchant_url"] = o.MerchantUrl.Get()
	}
	if o.MerchantBannerUrl.IsSet() {
		toSerialize["merchant_banner_url"] = o.MerchantBannerUrl.Get()
	}
	if o.MerchantColor.IsSet() {
		toSerialize["merchant_color"] = o.MerchantColor.Get()
	}
	if o.MerchantMessage.IsSet() {
		toSerialize["merchant_message"] = o.MerchantMessage.Get()
	}
	if o.MerchantTermsAndConditionsUrl.IsSet() {
		toSerialize["merchant_terms_and_conditions_url"] = o.MerchantTermsAndConditionsUrl.Get()
	}
	if o.MerchantFaviconUrl.IsSet() {
		toSerialize["merchant_favicon_url"] = o.MerchantFaviconUrl.Get()
	}
	if o.Intent.IsSet() {
		toSerialize["intent"] = o.Intent.Get()
	}
	if o.ReturnUrl.IsSet() {
		toSerialize["return_url"] = o.ReturnUrl.Get()
	}
	if o.CartItems != nil {
		toSerialize["cart_items"] = o.CartItems
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PaymentSource.IsSet() {
		toSerialize["payment_source"] = o.PaymentSource.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentLinkRequest struct {
	value *PaymentLinkRequest
	isSet bool
}

func (v NullablePaymentLinkRequest) Get() *PaymentLinkRequest {
	return v.value
}

func (v *NullablePaymentLinkRequest) Set(val *PaymentLinkRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentLinkRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentLinkRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentLinkRequest(val *PaymentLinkRequest) *NullablePaymentLinkRequest {
	return &NullablePaymentLinkRequest{value: val, isSet: true}
}

func (v NullablePaymentLinkRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentLinkRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


