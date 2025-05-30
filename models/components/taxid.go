// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TaxID struct {
	// The tax ID for the buyer.
	Value string    `json:"value"`
	Kind  TaxIDKind `json:"kind"`
}

func (o *TaxID) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *TaxID) GetKind() TaxIDKind {
	if o == nil {
		return TaxIDKind("")
	}
	return o.Kind
}
