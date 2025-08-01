// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type MerchantAccountCreate struct {
	// Whether the Real-Time Account Updater service is enabled for this merchant account. The Account Updater service is used to update card details when cards are lost, stolen or expired. If the field is not set or if it's set to `false`, the Account Updater service doesn't get called if a payment fails with expired or invalid card details. If the field is set to `true`, the service is called. Please note that for this to work the other `account_updater_* fields` must be set as well.
	AccountUpdaterEnabled *bool `default:"false" json:"account_updater_enabled"`
	// The public key used to encrypt the request to the Real-Time Account Updater service. The Account Updater service is used to update card details when cards are lost, stolen or expired. If the field is not set or if it's set to `null`, the Account Updater service doesn't get called. If the field is set, the other `account_updater_*` fields must be set as well.
	AccountUpdaterRequestEncryptionKey *string `json:"account_updater_request_encryption_key,omitempty"`
	// The ID of the key used to encrypt the request to the Real-Time Account Updater service. The Account Updater service is used to update card details when cards are lost, stolen or expired. If the field is not set or if it's set to `null`, the Account Updater service doesn't get called. If the field is set, the other `account_updater_*` fields must be set as well.
	AccountUpdaterRequestEncryptionKeyID *string `json:"account_updater_request_encryption_key_id,omitempty"`
	// The key used to decrypt the response from the Real-Time Account Updater service. The Account Updater service is used to update card details when cards are lost, stolen or expired. If the field is not set or if it's set to `null`, the Account Updater service doesn't get called. If the field is set, the other `account_updater_*` fields must be set as well.
	AccountUpdaterResponseDecryptionKey *string `json:"account_updater_response_decryption_key,omitempty"`
	// The ID of the key used to decrypt the request from the Real-Time Account Updater service. The Account Updater service is used to update card details when cards are lost, stolen or expired. If the field is not set or if it's set to `null`, the Account Updater service doesn't get called. If the field is set, the other `account_updater_*` fields must be set as well.
	AccountUpdaterResponseDecryptionKeyID *string `json:"account_updater_response_decryption_key_id,omitempty"`
	// The maximum monetary amount allowed for over-capture, in the smallest currency unit, for example `1299` cents to allow for an over-capture of `$12.99`.
	OverCaptureAmount *int64 `json:"over_capture_amount,omitempty"`
	// The maximum percentage allowed for over-capture, for example `25` to allow for an over-capture of `25%` of the original transaction amount.
	OverCapturePercentage *int64 `json:"over_capture_percentage,omitempty"`
	// Client key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service we use and if the field is not set or if it's set to null, the Account Updater service doesn't get configured. If the field is set to `null`, the other `loon_*` fields must be set to null as well.
	LoonClientKey *string `json:"loon_client_key,omitempty"`
	// Secret key provided by Pagos to authenticate to the Loon API. Loon is the Account Updater service we use and if the field is not set or if it's set to null, the Account Updater service doesn't get configured. If the field is set to `null`, the other `loon_*` fields must be set to null as well.
	LoonSecretKey *string `json:"loon_secret_key,omitempty"`
	// Card schemes accepted when creating jobs using this set of Loon API keys. Loon is the Account Updater service we use and if the field is not set or if it's set to null, the Account Updater service doesn't get configured. If the field is set to `null`, the other `loon_*` fields must be set to null as well.
	LoonAcceptedSchemes []CardScheme `json:"loon_accepted_schemes,omitempty"`
	// Requestor ID provided for Visa after onboarding to use Network Tokens.
	VisaNetworkTokensRequestorID *string `json:"visa_network_tokens_requestor_id,omitempty"`
	// Application ID provided for Visa after onboarding to use Network Tokens.
	VisaNetworkTokensAppID *string `json:"visa_network_tokens_app_id,omitempty"`
	// Requestor ID provided for American Express after onboarding to use Network Tokens.
	AmexNetworkTokensRequestorID *string `json:"amex_network_tokens_requestor_id,omitempty"`
	// Application ID provided for American Express after onboarding to use Network Tokens.
	AmexNetworkTokensAppID *string `json:"amex_network_tokens_app_id,omitempty"`
	// Requestor ID provided for Mastercard after onboarding to use Network Tokens.
	MastercardNetworkTokensRequestorID *string `json:"mastercard_network_tokens_requestor_id,omitempty"`
	// Application ID provided for Mastercard after onboarding to use Network Tokens.
	MastercardNetworkTokensAppID *string `json:"mastercard_network_tokens_app_id,omitempty"`
	// The ID for the merchant account.
	ID string `json:"id"`
	// The display name for the merchant account.
	DisplayName string `json:"display_name"`
}

func (m MerchantAccountCreate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MerchantAccountCreate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MerchantAccountCreate) GetAccountUpdaterEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.AccountUpdaterEnabled
}

func (o *MerchantAccountCreate) GetAccountUpdaterRequestEncryptionKey() *string {
	if o == nil {
		return nil
	}
	return o.AccountUpdaterRequestEncryptionKey
}

func (o *MerchantAccountCreate) GetAccountUpdaterRequestEncryptionKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AccountUpdaterRequestEncryptionKeyID
}

func (o *MerchantAccountCreate) GetAccountUpdaterResponseDecryptionKey() *string {
	if o == nil {
		return nil
	}
	return o.AccountUpdaterResponseDecryptionKey
}

func (o *MerchantAccountCreate) GetAccountUpdaterResponseDecryptionKeyID() *string {
	if o == nil {
		return nil
	}
	return o.AccountUpdaterResponseDecryptionKeyID
}

func (o *MerchantAccountCreate) GetOverCaptureAmount() *int64 {
	if o == nil {
		return nil
	}
	return o.OverCaptureAmount
}

func (o *MerchantAccountCreate) GetOverCapturePercentage() *int64 {
	if o == nil {
		return nil
	}
	return o.OverCapturePercentage
}

func (o *MerchantAccountCreate) GetLoonClientKey() *string {
	if o == nil {
		return nil
	}
	return o.LoonClientKey
}

func (o *MerchantAccountCreate) GetLoonSecretKey() *string {
	if o == nil {
		return nil
	}
	return o.LoonSecretKey
}

func (o *MerchantAccountCreate) GetLoonAcceptedSchemes() []CardScheme {
	if o == nil {
		return nil
	}
	return o.LoonAcceptedSchemes
}

func (o *MerchantAccountCreate) GetVisaNetworkTokensRequestorID() *string {
	if o == nil {
		return nil
	}
	return o.VisaNetworkTokensRequestorID
}

func (o *MerchantAccountCreate) GetVisaNetworkTokensAppID() *string {
	if o == nil {
		return nil
	}
	return o.VisaNetworkTokensAppID
}

func (o *MerchantAccountCreate) GetAmexNetworkTokensRequestorID() *string {
	if o == nil {
		return nil
	}
	return o.AmexNetworkTokensRequestorID
}

func (o *MerchantAccountCreate) GetAmexNetworkTokensAppID() *string {
	if o == nil {
		return nil
	}
	return o.AmexNetworkTokensAppID
}

func (o *MerchantAccountCreate) GetMastercardNetworkTokensRequestorID() *string {
	if o == nil {
		return nil
	}
	return o.MastercardNetworkTokensRequestorID
}

func (o *MerchantAccountCreate) GetMastercardNetworkTokensAppID() *string {
	if o == nil {
		return nil
	}
	return o.MastercardNetworkTokensAppID
}

func (o *MerchantAccountCreate) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *MerchantAccountCreate) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}
