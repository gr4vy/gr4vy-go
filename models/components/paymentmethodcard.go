// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
)

type PaymentMethodCard struct {
	// Set to `card` to use a new card.
	method *string `const:"card" json:"method"`
	// The 13-19 digit number for this card as it can be found on the front of the card.
	Number string `json:"number"`
	// The expiration date of the card, formatted `MM/YY`.
	ExpirationDate string `json:"expiration_date"`
	// The optional card's network scheme.
	CardScheme *CardScheme `json:"card_scheme,omitempty"`
	// The merchant identifier for this card.
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
}

func (p PaymentMethodCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymentMethodCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *PaymentMethodCard) GetMethod() *string {
	return types.String("card")
}

func (o *PaymentMethodCard) GetNumber() string {
	if o == nil {
		return ""
	}
	return o.Number
}

func (o *PaymentMethodCard) GetExpirationDate() string {
	if o == nil {
		return ""
	}
	return o.ExpirationDate
}

func (o *PaymentMethodCard) GetCardScheme() *CardScheme {
	if o == nil {
		return nil
	}
	return o.CardScheme
}

func (o *PaymentMethodCard) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}
