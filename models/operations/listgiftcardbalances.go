// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/gr4vy/gr4vy-go/models/components"
)

type ListGiftCardBalancesGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *ListGiftCardBalancesGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type ListGiftCardBalancesRequest struct {
	// The ID of the merchant account to use for this request.
	MerchantAccountID      *string                           `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
	GiftCardBalanceRequest components.GiftCardBalanceRequest `request:"mediaType=application/json"`
}

func (o *ListGiftCardBalancesRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

func (o *ListGiftCardBalancesRequest) GetGiftCardBalanceRequest() components.GiftCardBalanceRequest {
	if o == nil {
		return components.GiftCardBalanceRequest{}
	}
	return o.GiftCardBalanceRequest
}
