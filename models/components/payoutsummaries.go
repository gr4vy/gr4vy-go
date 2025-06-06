// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type PayoutSummaries struct {
	// A list of items returned for this request.
	Items []PayoutSummary `json:"items"`
	// The number of items for this page.
	Limit *int64 `default:"20" json:"limit"`
	// The cursor pointing at the next page of items.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor pointing at the previous page of items.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

func (p PayoutSummaries) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PayoutSummaries) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PayoutSummaries) GetItems() []PayoutSummary {
	if o == nil {
		return []PayoutSummary{}
	}
	return o.Items
}

func (o *PayoutSummaries) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *PayoutSummaries) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *PayoutSummaries) GetPreviousCursor() *string {
	if o == nil {
		return nil
	}
	return o.PreviousCursor
}
