// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// RequiredCheckoutFields - A collection of checkout fields and the conditions under which they are required.
type RequiredCheckoutFields struct {
	// A list of transaction fields that are required to process a payment for this service.
	RequiredFields []string `json:"required_fields"`
	// The conditions under which these fields are required
	Conditions map[string]any `json:"conditions,omitempty"`
}

func (o *RequiredCheckoutFields) GetRequiredFields() []string {
	if o == nil {
		return []string{}
	}
	return o.RequiredFields
}

func (o *RequiredCheckoutFields) GetConditions() map[string]any {
	if o == nil {
		return nil
	}
	return o.Conditions
}
