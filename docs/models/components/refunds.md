# Refunds


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `Items`                                                  | [][components.Refund](../../models/components/refund.md) | :heavy_check_mark:                                       | A list of items returned for this request.               |                                                          |
| `Limit`                                                  | **int64*                                                 | :heavy_minus_sign:                                       | The number of items for this page.                       | 20                                                       |
| `NextCursor`                                             | **string*                                                | :heavy_minus_sign:                                       | The cursor pointing at the next page of items.           | ZXhhbXBsZTE                                              |
| `PreviousCursor`                                         | **string*                                                | :heavy_minus_sign:                                       | The cursor pointing at the previous page of items.       | Xkjss7asS                                                |