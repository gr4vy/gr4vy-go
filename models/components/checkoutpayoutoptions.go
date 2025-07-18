// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type CheckoutPayoutOptions struct {
	// The processing channel to be used for the payment.
	ProcessingChannelID string `json:"processing_channel_id"`
	// The ID of the currency account that will fund the payout.
	SourceID string `json:"source_id"`
}

func (o *CheckoutPayoutOptions) GetProcessingChannelID() string {
	if o == nil {
		return ""
	}
	return o.ProcessingChannelID
}

func (o *CheckoutPayoutOptions) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}
