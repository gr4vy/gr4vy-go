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

// PaymentMethod A generic payment method.
type PaymentMethod struct {
	// `payment-method`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the payment method.
	Id *string `json:"id,omitempty"`
	// Additional schemes of the card. Only applies to card payment methods.
	AdditionalSchemes []string `json:"additional_schemes,omitempty"`
	// The browser target that an approval URL must be opened in. If `any` or `null`, then there is no specific requirement.
	ApprovalTarget NullableString `json:"approval_target,omitempty"`
	// The optional URL that the buyer needs to be redirected to to further authorize their payment.
	ApprovalUrl NullableString `json:"approval_url,omitempty"`
	// The optional buyer for which this payment method has been stored.
	Buyer NullableBuyer `json:"buyer,omitempty"`
	// The 2-letter ISO code of the country this payment method can be used for. If this value is `null` the payment method may be used in multiple countries.
	Country NullableString `json:"country,omitempty"`
	// The date and time when this payment method was first created in our system.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The ISO-4217 currency code that this payment method can be used for. If this value is `null` the payment method may be used for multiple currencies.
	Currency NullableString `json:"currency,omitempty"`
	Details *PaymentMethodDetailsCard `json:"details,omitempty"`
	// The expiration date for the payment method.
	ExpirationDate NullableString `json:"expiration_date,omitempty"`
	// An external identifier that can be used to match the payment method against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// Whether this card has a pending replacement that hasn't been applied yet.  When the Account Updater determines that new card details are available, existing details are not changed immediately, but this field is set to `true`. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * `canceled_payment_method`     * `expired_payment_method`     * `unavailable_payment_method`     * `unknown_payment_method`  When the replacement is applied, this field is set to `false`. For non-card payment methods, the value of this field is always set to `false`.
	HasReplacement *bool `json:"has_replacement,omitempty"`
	// A label for the card or the account. For a `paypal` payment method this is the user's email address. For a card it is the last 4 digits of the card.
	Label NullableString `json:"label,omitempty"`
	// The date and time when this card was last replaced.  When the Account Updater determines that new card details are available, existing details are not changed immediately. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * `canceled_payment_method`     * `expired_payment_method`     * `unavailable_payment_method`     * `unknown_payment_method`  When the replacement is applied, this field is updated. For non-card payment methods, the value of this field is always set to `null`.
	LastReplacedAt NullableTime `json:"last_replaced_at,omitempty"`
	// The unique ID for a merchant account.
	MerchantAccountId *string `json:"merchant_account_id,omitempty"`
	// The type of this payment method.
	Method *string `json:"method,omitempty"`
	// The mode to use with this payment method.
	Mode *string `json:"mode,omitempty"`
	// The scheme of the card. Only applies to card payments.
	Scheme NullableString `json:"scheme,omitempty"`
	// The state of the payment method.  - `processing` - The payment method is stored but is not ready to be    used yet, as we may be waiting for a notification from a connector    to complete the setup. - `buyer_approval_required` - Storing the payment method requires   the buyer to provide approval. Follow the `approval_url` for next steps. - `succeeded` - The payment method is stored and can be used. - `failed` - The payment method could not be stored, or failed verification.
	Status *string `json:"status,omitempty"`
	// The date and time when this payment method was last updated in our system.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The unique hash derived from the payment method identifier (e.g. card number).
	Fingerprint NullableString `json:"fingerprint,omitempty"`
	// The timestamp when this payment method was last used in a transaction.
	LastUsedAt NullableTime `json:"last_used_at,omitempty"`
	// The number of times this payment method has been used in transactions.
	UsageCount *int32 `json:"usage_count,omitempty"`
	// The timestamp when this payment method was last used in a transaction for client initiated transactions.
	CitLastUsedAt NullableTime `json:"cit_last_used_at,omitempty"`
	// The number of times this payment method has been used in transactions for client initiated transactions.
	CitUsageCount *int32 `json:"cit_usage_count,omitempty"`
}

// NewPaymentMethod instantiates a new PaymentMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethod() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodWithDefaults() *PaymentMethod {
	this := PaymentMethod{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethod) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethod) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethod) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentMethod) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethod) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentMethod) SetId(v string) {
	o.Id = &v
}

// GetAdditionalSchemes returns the AdditionalSchemes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetAdditionalSchemes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.AdditionalSchemes
}

// GetAdditionalSchemesOk returns a tuple with the AdditionalSchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetAdditionalSchemesOk() (*[]string, bool) {
	if o == nil || o.AdditionalSchemes == nil {
		return nil, false
	}
	return &o.AdditionalSchemes, true
}

// HasAdditionalSchemes returns a boolean if a field has been set.
func (o *PaymentMethod) HasAdditionalSchemes() bool {
	if o != nil && o.AdditionalSchemes != nil {
		return true
	}

	return false
}

// SetAdditionalSchemes gets a reference to the given []string and assigns it to the AdditionalSchemes field.
func (o *PaymentMethod) SetAdditionalSchemes(v []string) {
	o.AdditionalSchemes = v
}

// GetApprovalTarget returns the ApprovalTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetApprovalTarget() string {
	if o == nil || o.ApprovalTarget.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalTarget.Get()
}

// GetApprovalTargetOk returns a tuple with the ApprovalTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetApprovalTargetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalTarget.Get(), o.ApprovalTarget.IsSet()
}

// HasApprovalTarget returns a boolean if a field has been set.
func (o *PaymentMethod) HasApprovalTarget() bool {
	if o != nil && o.ApprovalTarget.IsSet() {
		return true
	}

	return false
}

// SetApprovalTarget gets a reference to the given NullableString and assigns it to the ApprovalTarget field.
func (o *PaymentMethod) SetApprovalTarget(v string) {
	o.ApprovalTarget.Set(&v)
}
// SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil
func (o *PaymentMethod) SetApprovalTargetNil() {
	o.ApprovalTarget.Set(nil)
}

// UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
func (o *PaymentMethod) UnsetApprovalTarget() {
	o.ApprovalTarget.Unset()
}

// GetApprovalUrl returns the ApprovalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetApprovalUrl() string {
	if o == nil || o.ApprovalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalUrl.Get()
}

// GetApprovalUrlOk returns a tuple with the ApprovalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetApprovalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalUrl.Get(), o.ApprovalUrl.IsSet()
}

// HasApprovalUrl returns a boolean if a field has been set.
func (o *PaymentMethod) HasApprovalUrl() bool {
	if o != nil && o.ApprovalUrl.IsSet() {
		return true
	}

	return false
}

// SetApprovalUrl gets a reference to the given NullableString and assigns it to the ApprovalUrl field.
func (o *PaymentMethod) SetApprovalUrl(v string) {
	o.ApprovalUrl.Set(&v)
}
// SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil
func (o *PaymentMethod) SetApprovalUrlNil() {
	o.ApprovalUrl.Set(nil)
}

// UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
func (o *PaymentMethod) UnsetApprovalUrl() {
	o.ApprovalUrl.Unset()
}

// GetBuyer returns the Buyer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetBuyer() Buyer {
	if o == nil || o.Buyer.Get() == nil {
		var ret Buyer
		return ret
	}
	return *o.Buyer.Get()
}

// GetBuyerOk returns a tuple with the Buyer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetBuyerOk() (*Buyer, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Buyer.Get(), o.Buyer.IsSet()
}

// HasBuyer returns a boolean if a field has been set.
func (o *PaymentMethod) HasBuyer() bool {
	if o != nil && o.Buyer.IsSet() {
		return true
	}

	return false
}

// SetBuyer gets a reference to the given NullableBuyer and assigns it to the Buyer field.
func (o *PaymentMethod) SetBuyer(v Buyer) {
	o.Buyer.Set(&v)
}
// SetBuyerNil sets the value for Buyer to be an explicit nil
func (o *PaymentMethod) SetBuyerNil() {
	o.Buyer.Set(nil)
}

// UnsetBuyer ensures that no value is present for Buyer, not even an explicit nil
func (o *PaymentMethod) UnsetBuyer() {
	o.Buyer.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethod) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *PaymentMethod) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *PaymentMethod) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *PaymentMethod) UnsetCountry() {
	o.Country.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PaymentMethod) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentMethod) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *PaymentMethod) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PaymentMethod) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PaymentMethod) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PaymentMethod) GetDetails() PaymentMethodDetailsCard {
	if o == nil || o.Details == nil {
		var ret PaymentMethodDetailsCard
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetDetailsOk() (*PaymentMethodDetailsCard, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PaymentMethod) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given PaymentMethodDetailsCard and assigns it to the Details field.
func (o *PaymentMethod) SetDetails(v PaymentMethodDetailsCard) {
	o.Details = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetExpirationDate() string {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *PaymentMethod) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableString and assigns it to the ExpirationDate field.
func (o *PaymentMethod) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *PaymentMethod) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *PaymentMethod) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PaymentMethod) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PaymentMethod) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PaymentMethod) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PaymentMethod) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetHasReplacement returns the HasReplacement field value if set, zero value otherwise.
func (o *PaymentMethod) GetHasReplacement() bool {
	if o == nil || o.HasReplacement == nil {
		var ret bool
		return ret
	}
	return *o.HasReplacement
}

// GetHasReplacementOk returns a tuple with the HasReplacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetHasReplacementOk() (*bool, bool) {
	if o == nil || o.HasReplacement == nil {
		return nil, false
	}
	return o.HasReplacement, true
}

// HasHasReplacement returns a boolean if a field has been set.
func (o *PaymentMethod) HasHasReplacement() bool {
	if o != nil && o.HasReplacement != nil {
		return true
	}

	return false
}

// SetHasReplacement gets a reference to the given bool and assigns it to the HasReplacement field.
func (o *PaymentMethod) SetHasReplacement(v bool) {
	o.HasReplacement = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetLabel() string {
	if o == nil || o.Label.Get() == nil {
		var ret string
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *PaymentMethod) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableString and assigns it to the Label field.
func (o *PaymentMethod) SetLabel(v string) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *PaymentMethod) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *PaymentMethod) UnsetLabel() {
	o.Label.Unset()
}

// GetLastReplacedAt returns the LastReplacedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetLastReplacedAt() time.Time {
	if o == nil || o.LastReplacedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReplacedAt.Get()
}

// GetLastReplacedAtOk returns a tuple with the LastReplacedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetLastReplacedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastReplacedAt.Get(), o.LastReplacedAt.IsSet()
}

// HasLastReplacedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasLastReplacedAt() bool {
	if o != nil && o.LastReplacedAt.IsSet() {
		return true
	}

	return false
}

// SetLastReplacedAt gets a reference to the given NullableTime and assigns it to the LastReplacedAt field.
func (o *PaymentMethod) SetLastReplacedAt(v time.Time) {
	o.LastReplacedAt.Set(&v)
}
// SetLastReplacedAtNil sets the value for LastReplacedAt to be an explicit nil
func (o *PaymentMethod) SetLastReplacedAtNil() {
	o.LastReplacedAt.Set(nil)
}

// UnsetLastReplacedAt ensures that no value is present for LastReplacedAt, not even an explicit nil
func (o *PaymentMethod) UnsetLastReplacedAt() {
	o.LastReplacedAt.Unset()
}

// GetMerchantAccountId returns the MerchantAccountId field value if set, zero value otherwise.
func (o *PaymentMethod) GetMerchantAccountId() string {
	if o == nil || o.MerchantAccountId == nil {
		var ret string
		return ret
	}
	return *o.MerchantAccountId
}

// GetMerchantAccountIdOk returns a tuple with the MerchantAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetMerchantAccountIdOk() (*string, bool) {
	if o == nil || o.MerchantAccountId == nil {
		return nil, false
	}
	return o.MerchantAccountId, true
}

// HasMerchantAccountId returns a boolean if a field has been set.
func (o *PaymentMethod) HasMerchantAccountId() bool {
	if o != nil && o.MerchantAccountId != nil {
		return true
	}

	return false
}

// SetMerchantAccountId gets a reference to the given string and assigns it to the MerchantAccountId field.
func (o *PaymentMethod) SetMerchantAccountId(v string) {
	o.MerchantAccountId = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentMethod) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentMethod) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentMethod) SetMethod(v string) {
	o.Method = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PaymentMethod) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PaymentMethod) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PaymentMethod) SetMode(v string) {
	o.Mode = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetScheme() string {
	if o == nil || o.Scheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *PaymentMethod) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *PaymentMethod) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *PaymentMethod) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *PaymentMethod) UnsetScheme() {
	o.Scheme.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PaymentMethod) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PaymentMethod) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PaymentMethod) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentMethod) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PaymentMethod) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetFingerprint() string {
	if o == nil || o.Fingerprint.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint.Get()
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetFingerprintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fingerprint.Get(), o.Fingerprint.IsSet()
}

// HasFingerprint returns a boolean if a field has been set.
func (o *PaymentMethod) HasFingerprint() bool {
	if o != nil && o.Fingerprint.IsSet() {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given NullableString and assigns it to the Fingerprint field.
func (o *PaymentMethod) SetFingerprint(v string) {
	o.Fingerprint.Set(&v)
}
// SetFingerprintNil sets the value for Fingerprint to be an explicit nil
func (o *PaymentMethod) SetFingerprintNil() {
	o.Fingerprint.Set(nil)
}

// UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
func (o *PaymentMethod) UnsetFingerprint() {
	o.Fingerprint.Unset()
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt.IsSet() {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given NullableTime and assigns it to the LastUsedAt field.
func (o *PaymentMethod) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}
// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil
func (o *PaymentMethod) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
func (o *PaymentMethod) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetUsageCount returns the UsageCount field value if set, zero value otherwise.
func (o *PaymentMethod) GetUsageCount() int32 {
	if o == nil || o.UsageCount == nil {
		var ret int32
		return ret
	}
	return *o.UsageCount
}

// GetUsageCountOk returns a tuple with the UsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetUsageCountOk() (*int32, bool) {
	if o == nil || o.UsageCount == nil {
		return nil, false
	}
	return o.UsageCount, true
}

// HasUsageCount returns a boolean if a field has been set.
func (o *PaymentMethod) HasUsageCount() bool {
	if o != nil && o.UsageCount != nil {
		return true
	}

	return false
}

// SetUsageCount gets a reference to the given int32 and assigns it to the UsageCount field.
func (o *PaymentMethod) SetUsageCount(v int32) {
	o.UsageCount = &v
}

// GetCitLastUsedAt returns the CitLastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethod) GetCitLastUsedAt() time.Time {
	if o == nil || o.CitLastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CitLastUsedAt.Get()
}

// GetCitLastUsedAtOk returns a tuple with the CitLastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethod) GetCitLastUsedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CitLastUsedAt.Get(), o.CitLastUsedAt.IsSet()
}

// HasCitLastUsedAt returns a boolean if a field has been set.
func (o *PaymentMethod) HasCitLastUsedAt() bool {
	if o != nil && o.CitLastUsedAt.IsSet() {
		return true
	}

	return false
}

// SetCitLastUsedAt gets a reference to the given NullableTime and assigns it to the CitLastUsedAt field.
func (o *PaymentMethod) SetCitLastUsedAt(v time.Time) {
	o.CitLastUsedAt.Set(&v)
}
// SetCitLastUsedAtNil sets the value for CitLastUsedAt to be an explicit nil
func (o *PaymentMethod) SetCitLastUsedAtNil() {
	o.CitLastUsedAt.Set(nil)
}

// UnsetCitLastUsedAt ensures that no value is present for CitLastUsedAt, not even an explicit nil
func (o *PaymentMethod) UnsetCitLastUsedAt() {
	o.CitLastUsedAt.Unset()
}

// GetCitUsageCount returns the CitUsageCount field value if set, zero value otherwise.
func (o *PaymentMethod) GetCitUsageCount() int32 {
	if o == nil || o.CitUsageCount == nil {
		var ret int32
		return ret
	}
	return *o.CitUsageCount
}

// GetCitUsageCountOk returns a tuple with the CitUsageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethod) GetCitUsageCountOk() (*int32, bool) {
	if o == nil || o.CitUsageCount == nil {
		return nil, false
	}
	return o.CitUsageCount, true
}

// HasCitUsageCount returns a boolean if a field has been set.
func (o *PaymentMethod) HasCitUsageCount() bool {
	if o != nil && o.CitUsageCount != nil {
		return true
	}

	return false
}

// SetCitUsageCount gets a reference to the given int32 and assigns it to the CitUsageCount field.
func (o *PaymentMethod) SetCitUsageCount(v int32) {
	o.CitUsageCount = &v
}

func (o PaymentMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AdditionalSchemes != nil {
		toSerialize["additional_schemes"] = o.AdditionalSchemes
	}
	if o.ApprovalTarget.IsSet() {
		toSerialize["approval_target"] = o.ApprovalTarget.Get()
	}
	if o.ApprovalUrl.IsSet() {
		toSerialize["approval_url"] = o.ApprovalUrl.Get()
	}
	if o.Buyer.IsSet() {
		toSerialize["buyer"] = o.Buyer.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Currency.IsSet() {
		toSerialize["currency"] = o.Currency.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ExpirationDate.IsSet() {
		toSerialize["expiration_date"] = o.ExpirationDate.Get()
	}
	if o.ExternalIdentifier.IsSet() {
		toSerialize["external_identifier"] = o.ExternalIdentifier.Get()
	}
	if o.HasReplacement != nil {
		toSerialize["has_replacement"] = o.HasReplacement
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if o.LastReplacedAt.IsSet() {
		toSerialize["last_replaced_at"] = o.LastReplacedAt.Get()
	}
	if o.MerchantAccountId != nil {
		toSerialize["merchant_account_id"] = o.MerchantAccountId
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Fingerprint.IsSet() {
		toSerialize["fingerprint"] = o.Fingerprint.Get()
	}
	if o.LastUsedAt.IsSet() {
		toSerialize["last_used_at"] = o.LastUsedAt.Get()
	}
	if o.UsageCount != nil {
		toSerialize["usage_count"] = o.UsageCount
	}
	if o.CitLastUsedAt.IsSet() {
		toSerialize["cit_last_used_at"] = o.CitLastUsedAt.Get()
	}
	if o.CitUsageCount != nil {
		toSerialize["cit_usage_count"] = o.CitUsageCount
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethod struct {
	value *PaymentMethod
	isSet bool
}

func (v NullablePaymentMethod) Get() *PaymentMethod {
	return v.value
}

func (v *NullablePaymentMethod) Set(val *PaymentMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethod(val *PaymentMethod) *NullablePaymentMethod {
	return &NullablePaymentMethod{value: val, isSet: true}
}

func (v NullablePaymentMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


