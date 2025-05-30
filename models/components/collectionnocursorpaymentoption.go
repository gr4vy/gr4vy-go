// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CollectionNoCursorPaymentOption struct {
	// A list of items returned for this request.
	Items []PaymentOption `json:"items"`
}

func (o *CollectionNoCursorPaymentOption) GetItems() []PaymentOption {
	if o == nil {
		return []PaymentOption{}
	}
	return o.Items
}
