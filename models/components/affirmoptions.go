// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AffirmOptions struct {
	// Passes additional discounts to the Affirm widget.
	Discounts map[string]map[string]any `json:"discounts,omitempty"`
	// Passes itinerary data to the Affirm API.
	Itinerary *AffirmItineraryOptions `json:"itinerary,omitempty"`
}

func (o *AffirmOptions) GetDiscounts() map[string]map[string]any {
	if o == nil {
		return nil
	}
	return o.Discounts
}

func (o *AffirmOptions) GetItinerary() *AffirmItineraryOptions {
	if o == nil {
		return nil
	}
	return o.Itinerary
}
