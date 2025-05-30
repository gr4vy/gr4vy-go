// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type CollectionPaymentService struct {
	// A list of items returned for this request.
	Items []PaymentService `json:"items"`
	// The number of items for this page.
	Limit *int64 `default:"20" json:"limit"`
	// The cursor pointing at the next page of items.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor pointing at the previous page of items.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

func (c CollectionPaymentService) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionPaymentService) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionPaymentService) GetItems() []PaymentService {
	if o == nil {
		return []PaymentService{}
	}
	return o.Items
}

func (o *CollectionPaymentService) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *CollectionPaymentService) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *CollectionPaymentService) GetPreviousCursor() *string {
	if o == nil {
		return nil
	}
	return o.PreviousCursor
}
