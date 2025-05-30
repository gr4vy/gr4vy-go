// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type GooglePayPaymentOptionContext struct {
	MerchantName      string   `json:"merchant_name"`
	SupportedSchemes  []string `json:"supported_schemes"`
	Gateway           string   `json:"gateway"`
	GatewayMerchantID string   `json:"gateway_merchant_id"`
}

func (o *GooglePayPaymentOptionContext) GetMerchantName() string {
	if o == nil {
		return ""
	}
	return o.MerchantName
}

func (o *GooglePayPaymentOptionContext) GetSupportedSchemes() []string {
	if o == nil {
		return []string{}
	}
	return o.SupportedSchemes
}

func (o *GooglePayPaymentOptionContext) GetGateway() string {
	if o == nil {
		return ""
	}
	return o.Gateway
}

func (o *GooglePayPaymentOptionContext) GetGatewayMerchantID() string {
	if o == nil {
		return ""
	}
	return o.GatewayMerchantID
}
