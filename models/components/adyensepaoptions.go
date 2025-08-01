// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AdyenSepaOptions struct {
	// Set to `true` to enable Auto Rescue for a transaction. Use the `maxDaysToRescue` to specify a rescue window.
	AutoRescue *bool `json:"autoRescue,omitempty"`
	// The rescue window for a transaction, in days, when `autoRescue` is set to `true`. You can specify a value between 1 and 48. For cards, the default is one calendar month. For SEPA, the default is 42 days.
	MaxDaysToRescue *int64 `json:"maxDaysToRescue,omitempty"`
	// Passes additional data to the Adyen API when creating a transaction.
	AdditionalData map[string]string `json:"additionalData,omitempty"`
	// The rescue scenario to simulate for a transaction, when `autoRescue` is set to `true`.
	AutoRescueSepaScenario *AdyenAutoRescueSepaScenariosEnum `json:"autoRescueSepaScenario,omitempty"`
	// The name on the SEPA bank account.
	OwnerName *string `json:"ownerName,omitempty"`
}

func (o *AdyenSepaOptions) GetAutoRescue() *bool {
	if o == nil {
		return nil
	}
	return o.AutoRescue
}

func (o *AdyenSepaOptions) GetMaxDaysToRescue() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxDaysToRescue
}

func (o *AdyenSepaOptions) GetAdditionalData() map[string]string {
	if o == nil {
		return nil
	}
	return o.AdditionalData
}

func (o *AdyenSepaOptions) GetAutoRescueSepaScenario() *AdyenAutoRescueSepaScenariosEnum {
	if o == nil {
		return nil
	}
	return o.AutoRescueSepaScenario
}

func (o *AdyenSepaOptions) GetOwnerName() *string {
	if o == nil {
		return nil
	}
	return o.OwnerName
}
