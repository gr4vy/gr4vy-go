// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
	"github.com/gr4vy/gr4vy-go/types"
)

type PayoutPaymentService struct {
	// Always `payment-service`.
	type_ *string `const:"payment-service" json:"type"`
	// The ID for the payout service.
	ID *string `json:"id,omitempty"`
	// Always `card`.
	method *string `const:"card" json:"method"`
	// The ID of the connection used for this payout.
	PaymentServiceDefinitionID string `json:"payment_service_definition_id"`
	// The display name of the connection used for this payout.
	DisplayName *string `json:"display_name,omitempty"`
}

func (p PayoutPaymentService) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PayoutPaymentService) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PayoutPaymentService) GetType() *string {
	return types.String("payment-service")
}

func (o *PayoutPaymentService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PayoutPaymentService) GetMethod() *string {
	return types.String("card")
}

func (o *PayoutPaymentService) GetPaymentServiceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.PaymentServiceDefinitionID
}

func (o *PayoutPaymentService) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}
