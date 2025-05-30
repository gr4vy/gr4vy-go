// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
)

type TransactionBuyer struct {
	// Always `buyer`.
	type_ *string `const:"buyer" json:"type"`
	// The ID for the buyer.
	ID *string `json:"id,omitempty"`
	// The display name for the buyer.
	DisplayName *string `json:"display_name,omitempty"`
	// The merchant identifier for this buyer.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	// The billing name, address, email, and other fields for this buyer.
	BillingDetails *BillingDetailsOutput `json:"billing_details,omitempty"`
	// The buyer account number.
	AccountNumber *string `json:"account_number,omitempty"`
}

func (t TransactionBuyer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransactionBuyer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransactionBuyer) GetType() *string {
	return types.String("buyer")
}

func (o *TransactionBuyer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TransactionBuyer) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *TransactionBuyer) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}

func (o *TransactionBuyer) GetBillingDetails() *BillingDetailsOutput {
	if o == nil {
		return nil
	}
	return o.BillingDetails
}

func (o *TransactionBuyer) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}
