// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type VoidTransactionGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *VoidTransactionGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type VoidTransactionRequest struct {
	// The ID of the transaction
	TransactionID string `pathParam:"style=simple,explode=false,name=transaction_id"`
	// The ID of the merchant account to use for this request.
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *VoidTransactionRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

func (o *VoidTransactionRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}
