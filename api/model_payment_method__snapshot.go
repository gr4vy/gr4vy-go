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

// PaymentMethodSnapshot Snapshot of a payment method, as used when embedded inside other resources.
type PaymentMethodSnapshot struct {
	// `payment-method`.
	Type *string `json:"type,omitempty"`
	// The unique ID of the payment method.
	Id NullableString `json:"id,omitempty"`
	// The browser target that an approval URL must be opened in. If `any` or `null`, then there is no specific requirement.
	ApprovalTarget NullableString `json:"approval_target,omitempty"`
	// The optional URL that the buyer needs to be redirected to to further authorize their payment.
	ApprovalUrl NullableString `json:"approval_url,omitempty"`
	// The 2-letter ISO code of the country this payment method can be used for. If this value is `null` the payment method may be used in multiple countries.
	Country NullableString `json:"country,omitempty"`
	// The ISO-4217 currency code that this payment method can be used for. If this value is `null` the payment method may be used for multiple currencies.
	Currency NullableString `json:"currency,omitempty"`
	Details *PaymentMethodDetailsCard `json:"details,omitempty"`
	// The expiration date for this payment method. This is mostly used by cards where the card might have an expiration date.
	ExpirationDate NullableString `json:"expiration_date,omitempty"`
	// An external identifier that can be used to match the payment method against your own records.
	ExternalIdentifier NullableString `json:"external_identifier,omitempty"`
	// A label for the payment method. This can be the last 4 digits for a card, or the email address for an alternative payment method.
	Label *string `json:"label,omitempty"`
	// The date and time when this card was last replaced.  When the Account Updater determines that new card details are available, existing details are not changed immediately. There are three scenarios in which the actual replacement occurs:  1. When this card has expired. 2. When only the expiration date changed. 3. When a transaction using this card is declined with any of the following codes:     * `canceled_payment_method`     * `expired_payment_method`     * `unavailable_payment_method`     * `unknown_payment_method`  When the replacement is applied, this field is updated. For non-card payment methods, the value of this field is always set to `null`.
	LastReplacedAt NullableTime `json:"last_replaced_at,omitempty"`
	// The type of this payment method.
	Method *string `json:"method,omitempty"`
	// The payment account reference (PAR) returned by the card scheme. This is a unique reference to the underlying account that has been used to fund this payment method. This value will be unique if the same underlying account was used, regardless of the actual payment method used. For example, a network token or an Apple Pay device token will return the same PAR when possible.  The uniqueness of this value will depend on the card scheme, please refer to their documentation for further details. The availability of the PAR in our API depends on the availability of its value in the API of the payment service used for the transaction.
	PaymentAccountReference NullableString `json:"payment_account_reference,omitempty"`
	// An additional label used to differentiate different sub-types of a payment method. Most notably this can include the type of card used in a transaction.
	Scheme NullableString `json:"scheme,omitempty"`
}

// NewPaymentMethodSnapshot instantiates a new PaymentMethodSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodSnapshot() *PaymentMethodSnapshot {
	this := PaymentMethodSnapshot{}
	return &this
}

// NewPaymentMethodSnapshotWithDefaults instantiates a new PaymentMethodSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodSnapshotWithDefaults() *PaymentMethodSnapshot {
	this := PaymentMethodSnapshot{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentMethodSnapshot) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSnapshot) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PaymentMethodSnapshot) SetType(v string) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *PaymentMethodSnapshot) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PaymentMethodSnapshot) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetId() {
	o.Id.Unset()
}

// GetApprovalTarget returns the ApprovalTarget field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetApprovalTarget() string {
	if o == nil || o.ApprovalTarget.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalTarget.Get()
}

// GetApprovalTargetOk returns a tuple with the ApprovalTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetApprovalTargetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalTarget.Get(), o.ApprovalTarget.IsSet()
}

// HasApprovalTarget returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasApprovalTarget() bool {
	if o != nil && o.ApprovalTarget.IsSet() {
		return true
	}

	return false
}

// SetApprovalTarget gets a reference to the given NullableString and assigns it to the ApprovalTarget field.
func (o *PaymentMethodSnapshot) SetApprovalTarget(v string) {
	o.ApprovalTarget.Set(&v)
}
// SetApprovalTargetNil sets the value for ApprovalTarget to be an explicit nil
func (o *PaymentMethodSnapshot) SetApprovalTargetNil() {
	o.ApprovalTarget.Set(nil)
}

// UnsetApprovalTarget ensures that no value is present for ApprovalTarget, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetApprovalTarget() {
	o.ApprovalTarget.Unset()
}

// GetApprovalUrl returns the ApprovalUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetApprovalUrl() string {
	if o == nil || o.ApprovalUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalUrl.Get()
}

// GetApprovalUrlOk returns a tuple with the ApprovalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetApprovalUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalUrl.Get(), o.ApprovalUrl.IsSet()
}

// HasApprovalUrl returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasApprovalUrl() bool {
	if o != nil && o.ApprovalUrl.IsSet() {
		return true
	}

	return false
}

// SetApprovalUrl gets a reference to the given NullableString and assigns it to the ApprovalUrl field.
func (o *PaymentMethodSnapshot) SetApprovalUrl(v string) {
	o.ApprovalUrl.Set(&v)
}
// SetApprovalUrlNil sets the value for ApprovalUrl to be an explicit nil
func (o *PaymentMethodSnapshot) SetApprovalUrlNil() {
	o.ApprovalUrl.Set(nil)
}

// UnsetApprovalUrl ensures that no value is present for ApprovalUrl, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetApprovalUrl() {
	o.ApprovalUrl.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetCountry() string {
	if o == nil || o.Country.Get() == nil {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetCountryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *PaymentMethodSnapshot) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *PaymentMethodSnapshot) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetCountry() {
	o.Country.Unset()
}

// GetCurrency returns the Currency field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetCurrency() string {
	if o == nil || o.Currency.Get() == nil {
		var ret string
		return ret
	}
	return *o.Currency.Get()
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetCurrencyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Currency.Get(), o.Currency.IsSet()
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasCurrency() bool {
	if o != nil && o.Currency.IsSet() {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given NullableString and assigns it to the Currency field.
func (o *PaymentMethodSnapshot) SetCurrency(v string) {
	o.Currency.Set(&v)
}
// SetCurrencyNil sets the value for Currency to be an explicit nil
func (o *PaymentMethodSnapshot) SetCurrencyNil() {
	o.Currency.Set(nil)
}

// UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetCurrency() {
	o.Currency.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *PaymentMethodSnapshot) GetDetails() PaymentMethodDetailsCard {
	if o == nil || o.Details == nil {
		var ret PaymentMethodDetailsCard
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSnapshot) GetDetailsOk() (*PaymentMethodDetailsCard, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given PaymentMethodDetailsCard and assigns it to the Details field.
func (o *PaymentMethodSnapshot) SetDetails(v PaymentMethodDetailsCard) {
	o.Details = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetExpirationDate() string {
	if o == nil || o.ExpirationDate.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExpirationDate.Get()
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetExpirationDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExpirationDate.Get(), o.ExpirationDate.IsSet()
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasExpirationDate() bool {
	if o != nil && o.ExpirationDate.IsSet() {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given NullableString and assigns it to the ExpirationDate field.
func (o *PaymentMethodSnapshot) SetExpirationDate(v string) {
	o.ExpirationDate.Set(&v)
}
// SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil
func (o *PaymentMethodSnapshot) SetExpirationDateNil() {
	o.ExpirationDate.Set(nil)
}

// UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetExpirationDate() {
	o.ExpirationDate.Unset()
}

// GetExternalIdentifier returns the ExternalIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetExternalIdentifier() string {
	if o == nil || o.ExternalIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalIdentifier.Get()
}

// GetExternalIdentifierOk returns a tuple with the ExternalIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetExternalIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExternalIdentifier.Get(), o.ExternalIdentifier.IsSet()
}

// HasExternalIdentifier returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasExternalIdentifier() bool {
	if o != nil && o.ExternalIdentifier.IsSet() {
		return true
	}

	return false
}

// SetExternalIdentifier gets a reference to the given NullableString and assigns it to the ExternalIdentifier field.
func (o *PaymentMethodSnapshot) SetExternalIdentifier(v string) {
	o.ExternalIdentifier.Set(&v)
}
// SetExternalIdentifierNil sets the value for ExternalIdentifier to be an explicit nil
func (o *PaymentMethodSnapshot) SetExternalIdentifierNil() {
	o.ExternalIdentifier.Set(nil)
}

// UnsetExternalIdentifier ensures that no value is present for ExternalIdentifier, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetExternalIdentifier() {
	o.ExternalIdentifier.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PaymentMethodSnapshot) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSnapshot) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PaymentMethodSnapshot) SetLabel(v string) {
	o.Label = &v
}

// GetLastReplacedAt returns the LastReplacedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetLastReplacedAt() time.Time {
	if o == nil || o.LastReplacedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastReplacedAt.Get()
}

// GetLastReplacedAtOk returns a tuple with the LastReplacedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetLastReplacedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastReplacedAt.Get(), o.LastReplacedAt.IsSet()
}

// HasLastReplacedAt returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasLastReplacedAt() bool {
	if o != nil && o.LastReplacedAt.IsSet() {
		return true
	}

	return false
}

// SetLastReplacedAt gets a reference to the given NullableTime and assigns it to the LastReplacedAt field.
func (o *PaymentMethodSnapshot) SetLastReplacedAt(v time.Time) {
	o.LastReplacedAt.Set(&v)
}
// SetLastReplacedAtNil sets the value for LastReplacedAt to be an explicit nil
func (o *PaymentMethodSnapshot) SetLastReplacedAtNil() {
	o.LastReplacedAt.Set(nil)
}

// UnsetLastReplacedAt ensures that no value is present for LastReplacedAt, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetLastReplacedAt() {
	o.LastReplacedAt.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *PaymentMethodSnapshot) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodSnapshot) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *PaymentMethodSnapshot) SetMethod(v string) {
	o.Method = &v
}

// GetPaymentAccountReference returns the PaymentAccountReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetPaymentAccountReference() string {
	if o == nil || o.PaymentAccountReference.Get() == nil {
		var ret string
		return ret
	}
	return *o.PaymentAccountReference.Get()
}

// GetPaymentAccountReferenceOk returns a tuple with the PaymentAccountReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetPaymentAccountReferenceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentAccountReference.Get(), o.PaymentAccountReference.IsSet()
}

// HasPaymentAccountReference returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasPaymentAccountReference() bool {
	if o != nil && o.PaymentAccountReference.IsSet() {
		return true
	}

	return false
}

// SetPaymentAccountReference gets a reference to the given NullableString and assigns it to the PaymentAccountReference field.
func (o *PaymentMethodSnapshot) SetPaymentAccountReference(v string) {
	o.PaymentAccountReference.Set(&v)
}
// SetPaymentAccountReferenceNil sets the value for PaymentAccountReference to be an explicit nil
func (o *PaymentMethodSnapshot) SetPaymentAccountReferenceNil() {
	o.PaymentAccountReference.Set(nil)
}

// UnsetPaymentAccountReference ensures that no value is present for PaymentAccountReference, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetPaymentAccountReference() {
	o.PaymentAccountReference.Unset()
}

// GetScheme returns the Scheme field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentMethodSnapshot) GetScheme() string {
	if o == nil || o.Scheme.Get() == nil {
		var ret string
		return ret
	}
	return *o.Scheme.Get()
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentMethodSnapshot) GetSchemeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Scheme.Get(), o.Scheme.IsSet()
}

// HasScheme returns a boolean if a field has been set.
func (o *PaymentMethodSnapshot) HasScheme() bool {
	if o != nil && o.Scheme.IsSet() {
		return true
	}

	return false
}

// SetScheme gets a reference to the given NullableString and assigns it to the Scheme field.
func (o *PaymentMethodSnapshot) SetScheme(v string) {
	o.Scheme.Set(&v)
}
// SetSchemeNil sets the value for Scheme to be an explicit nil
func (o *PaymentMethodSnapshot) SetSchemeNil() {
	o.Scheme.Set(nil)
}

// UnsetScheme ensures that no value is present for Scheme, not even an explicit nil
func (o *PaymentMethodSnapshot) UnsetScheme() {
	o.Scheme.Unset()
}

func (o PaymentMethodSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ApprovalTarget.IsSet() {
		toSerialize["approval_target"] = o.ApprovalTarget.Get()
	}
	if o.ApprovalUrl.IsSet() {
		toSerialize["approval_url"] = o.ApprovalUrl.Get()
	}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
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
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.LastReplacedAt.IsSet() {
		toSerialize["last_replaced_at"] = o.LastReplacedAt.Get()
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.PaymentAccountReference.IsSet() {
		toSerialize["payment_account_reference"] = o.PaymentAccountReference.Get()
	}
	if o.Scheme.IsSet() {
		toSerialize["scheme"] = o.Scheme.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentMethodSnapshot struct {
	value *PaymentMethodSnapshot
	isSet bool
}

func (v NullablePaymentMethodSnapshot) Get() *PaymentMethodSnapshot {
	return v.value
}

func (v *NullablePaymentMethodSnapshot) Set(val *PaymentMethodSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodSnapshot(val *PaymentMethodSnapshot) *NullablePaymentMethodSnapshot {
	return &NullablePaymentMethodSnapshot{value: val, isSet: true}
}

func (v NullablePaymentMethodSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


