// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
	"time"
)

type Buyer struct {
	// Always `buyer`.
	type_ *string `const:"buyer" json:"type"`
	// The ID for the buyer.
	ID *string `json:"id,omitempty"`
	// The ID of the merchant account this buyer belongs to.
	MerchantAccountID string `json:"merchant_account_id"`
	// The display name for the buyer.
	DisplayName *string `json:"display_name,omitempty"`
	// The merchant identifier for this buyer.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	// The billing name, address, email, and other fields for this buyer.
	BillingDetails *BillingDetailsOutput `json:"billing_details,omitempty"`
	// The buyer account number
	AccountNumber *string `json:"account_number,omitempty"`
	// The date this buyer was created at.
	CreatedAt time.Time `json:"created_at"`
	// The date this buyer was last updated at.
	UpdatedAt time.Time `json:"updated_at"`
}

func (b Buyer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *Buyer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Buyer) GetType() *string {
	return types.String("buyer")
}

func (o *Buyer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Buyer) GetMerchantAccountID() string {
	if o == nil {
		return ""
	}
	return o.MerchantAccountID
}

func (o *Buyer) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Buyer) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}

func (o *Buyer) GetBillingDetails() *BillingDetailsOutput {
	if o == nil {
		return nil
	}
	return o.BillingDetails
}

func (o *Buyer) GetAccountNumber() *string {
	if o == nil {
		return nil
	}
	return o.AccountNumber
}

func (o *Buyer) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Buyer) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
