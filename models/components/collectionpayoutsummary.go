// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type CollectionPayoutSummary struct {
	// A list of items returned for this request.
	Items []PayoutSummary `json:"items"`
	// The number of items for this page.
	Limit *int64 `default:"20" json:"limit"`
	// The cursor pointing at the next page of items.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor pointing at the previous page of items.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

func (c CollectionPayoutSummary) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionPayoutSummary) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionPayoutSummary) GetItems() []PayoutSummary {
	if o == nil {
		return []PayoutSummary{}
	}
	return o.Items
}

func (o *CollectionPayoutSummary) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *CollectionPayoutSummary) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *CollectionPayoutSummary) GetPreviousCursor() *string {
	if o == nil {
		return nil
	}
	return o.PreviousCursor
}
