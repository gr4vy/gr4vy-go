// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type ReportExecutions struct {
	// A list of items returned for this request.
	Items []ReportExecution `json:"items"`
	// The number of items for this page.
	Limit *int64 `default:"20" json:"limit"`
	// The cursor pointing at the next page of items.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor pointing at the previous page of items.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

func (r ReportExecutions) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ReportExecutions) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ReportExecutions) GetItems() []ReportExecution {
	if o == nil {
		return []ReportExecution{}
	}
	return o.Items
}

func (o *ReportExecutions) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ReportExecutions) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ReportExecutions) GetPreviousCursor() *string {
	if o == nil {
		return nil
	}
	return o.PreviousCursor
}
