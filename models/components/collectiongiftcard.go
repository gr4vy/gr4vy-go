// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/gr4vy/gr4vy-go/internal/utils"
)

type CollectionGiftCard struct {
	// A list of items returned for this request.
	Items []GiftCard `json:"items"`
	// The number of items for this page.
	Limit *int64 `default:"20" json:"limit"`
	// The cursor pointing at the next page of items.
	NextCursor *string `json:"next_cursor,omitempty"`
	// The cursor pointing at the previous page of items.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
}

func (c CollectionGiftCard) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CollectionGiftCard) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CollectionGiftCard) GetItems() []GiftCard {
	if o == nil {
		return []GiftCard{}
	}
	return o.Items
}

func (o *CollectionGiftCard) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *CollectionGiftCard) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *CollectionGiftCard) GetPreviousCursor() *string {
	if o == nil {
		return nil
	}
	return o.PreviousCursor
}
