// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

type ListTransactionSettlementsGlobals struct {
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *ListTransactionSettlementsGlobals) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}

type ListTransactionSettlementsRequest struct {
	// The unique identifier of the transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transaction_id"`
	// The ID of the merchant account to use for this request.
	MerchantAccountID *string `header:"style=simple,explode=false,name=x-gr4vy-merchant-account-id"`
}

func (o *ListTransactionSettlementsRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

func (o *ListTransactionSettlementsRequest) GetMerchantAccountID() *string {
	if o == nil {
		return nil
	}
	return o.MerchantAccountID
}
