// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
)

type TransactionGiftCard struct {
	// Always `gift-card`.
	type_ *string `const:"gift-card" json:"type"`
	// The ID for the gift card.
	ID *string `json:"id,omitempty"`
	// The first 6 digits of the full gift card number.
	Bin string `json:"bin"`
	// The 3 digits after the `bin` of the full gift card number.
	SubBin string `json:"sub_bin"`
	// The last 4 digits for the gift card.
	Last4 string `json:"last4"`
}

func (t TransactionGiftCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TransactionGiftCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TransactionGiftCard) GetType() *string {
	return types.String("gift-card")
}

func (o *TransactionGiftCard) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TransactionGiftCard) GetBin() string {
	if o == nil {
		return ""
	}
	return o.Bin
}

func (o *TransactionGiftCard) GetSubBin() string {
	if o == nil {
		return ""
	}
	return o.SubBin
}

func (o *TransactionGiftCard) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}
