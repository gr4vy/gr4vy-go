// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type GetPaymentServiceDefinitionRequest struct {
	PaymentServiceDefinitionID string `pathParam:"style=simple,explode=false,name=payment_service_definition_id"`
}

func (o *GetPaymentServiceDefinitionRequest) GetPaymentServiceDefinitionID() string {
	if o == nil {
		return ""
	}
	return o.PaymentServiceDefinitionID
}
