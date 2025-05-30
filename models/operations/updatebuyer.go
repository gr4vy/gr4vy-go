// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

type UpdateBuyerGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *UpdateBuyerGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type UpdateBuyerRequest struct {
	// The ID of the buyer to edit.
	BuyerID string `pathParam:"style=simple,explode=false,name=buyer_id"`
	// The ID of the merchant account to use for this request.
	MerchantAccountID *string                `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
	BuyerUpdate       components.BuyerUpdate `request:"mediaType=application/json"`
}

func (o *UpdateBuyerRequest) GetBuyerID() string {
	if o == nil {
		return ""
	}
	return o.BuyerID
}

func (o *UpdateBuyerRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

func (o *UpdateBuyerRequest) GetBuyerUpdate() components.BuyerUpdate {
	if o == nil {
		return components.BuyerUpdate{}
	}
	return o.BuyerUpdate
}
