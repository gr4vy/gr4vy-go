// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type DeleteCheckoutSessionGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *DeleteCheckoutSessionGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type DeleteCheckoutSessionRequest struct {
	// The ID of the checkout session.
	SessionID string `pathParam:"style=simple,explode=false,name=session_id"`
	// The ID of the merchant account to use for this request.
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *DeleteCheckoutSessionRequest) GetSessionID() string {
	if o == nil {
		return ""
	}
	return o.SessionID
}

func (o *DeleteCheckoutSessionRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}
