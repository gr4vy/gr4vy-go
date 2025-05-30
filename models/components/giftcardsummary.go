// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
	"time"
)

type GiftCardSummary struct {
	// Always `gift-card`.
	type_ *string `const:"gift-card" json:"type"`
	// The ID for the gift card.
	ID *string `json:"id,omitempty"`
	// The ID of the merchant account this buyer belongs to.
	MerchantAccountID string `json:"merchant_account_id"`
	// The first 6 digits of the full gift card number.
	Bin string `json:"bin"`
	// The 3 digits after the `bin` of the full gift card number.
	SubBin string `json:"sub_bin"`
	// The last 4 digits for the gift card.
	Last4 string `json:"last4"`
	// The ISO-4217 currency code that this gift card has a balance for.
	Currency *string `json:"currency,omitempty"`
	//  The date and time when this gift card expires. This is a full date/time and may be more accurate than the actual expiry date received by the gift card service.
	ExpirationDate *time.Time `json:"expiration_date,omitempty"`
	// The amount remaining on the balance for this gift card according to the gift card service. This may be `null` if the balance could not be fetched.
	Balance *int64 `json:"balance,omitempty"`
	// If the last balance update failed, this will contain the internal code for this error.
	BalanceErrorCode *GiftCardErrorCode `json:"balance_error_code,omitempty"`
	// If the last balance update failed, this will contain the the raw error code received from the gift card provider.
	BalanceRawErrorCode *string `json:"balance_raw_error_code,omitempty"`
	// If the last balance update failed, this will contain the the raw error message received from the gift card provider.
	BalanceRawErrorMessage *string `json:"balance_raw_error_message,omitempty"`
}

func (g GiftCardSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GiftCardSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GiftCardSummary) GetType() *string {
	return types.String("gift-card")
}

func (o *GiftCardSummary) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GiftCardSummary) GetMerchantAccountID() string {
	if o == nil {
		return ""
	}
	return o.MerchantAccountID
}

func (o *GiftCardSummary) GetBin() string {
	if o == nil {
		return ""
	}
	return o.Bin
}

func (o *GiftCardSummary) GetSubBin() string {
	if o == nil {
		return ""
	}
	return o.SubBin
}

func (o *GiftCardSummary) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *GiftCardSummary) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *GiftCardSummary) GetExpirationDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpirationDate
}

func (o *GiftCardSummary) GetBalance() *int64 {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *GiftCardSummary) GetBalanceErrorCode() *GiftCardErrorCode {
	if o == nil {
		return nil
	}
	return o.BalanceErrorCode
}

func (o *GiftCardSummary) GetBalanceRawErrorCode() *string {
	if o == nil {
		return nil
	}
	return o.BalanceRawErrorCode
}

func (o *GiftCardSummary) GetBalanceRawErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.BalanceRawErrorMessage
}
