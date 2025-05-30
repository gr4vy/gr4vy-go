// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ThreeDSecureV2 struct {
	Version                string  `json:"version"`
	AuthenticationResponse *string `json:"authentication_response,omitempty"`
	DirectoryResponse      *string `json:"directory_response,omitempty"`
	DirectoryTransactionID *string `json:"directory_transaction_id,omitempty"`
	TransactionReason      *string `json:"transaction_reason,omitempty"`
	Cavv                   *string `json:"cavv,omitempty"`
	Eci                    *string `json:"eci,omitempty"`
	CardholderInfo         *string `json:"cardholder_info,omitempty"`
}

func (o *ThreeDSecureV2) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *ThreeDSecureV2) GetAuthenticationResponse() *string {
	if o == nil {
		return nil
	}
	return o.AuthenticationResponse
}

func (o *ThreeDSecureV2) GetDirectoryResponse() *string {
	if o == nil {
		return nil
	}
	return o.DirectoryResponse
}

func (o *ThreeDSecureV2) GetDirectoryTransactionID() *string {
	if o == nil {
		return nil
	}
	return o.DirectoryTransactionID
}

func (o *ThreeDSecureV2) GetTransactionReason() *string {
	if o == nil {
		return nil
	}
	return o.TransactionReason
}

func (o *ThreeDSecureV2) GetCavv() *string {
	if o == nil {
		return nil
	}
	return o.Cavv
}

func (o *ThreeDSecureV2) GetEci() *string {
	if o == nil {
		return nil
	}
	return o.Eci
}

func (o *ThreeDSecureV2) GetCardholderInfo() *string {
	if o == nil {
		return nil
	}
	return o.CardholderInfo
}
