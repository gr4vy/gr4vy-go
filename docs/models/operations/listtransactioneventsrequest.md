# ListTransactionEventsRequest


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `TransactionID`                                         | *string*                                                | :heavy_check_mark:                                      | N/A                                                     | 7099948d-7286-47e4-aad8-b68f7eb44591                    |
| `Cursor`                                                | **string*                                               | :heavy_minus_sign:                                      | A pointer to the page of results to return.             | ZXhhbXBsZTE                                             |
| `Limit`                                                 | **int64*                                                | :heavy_minus_sign:                                      | The maximum number of items that are at returned.       | 100                                                     |
| `MerchantAccountID`                                     | **string*                                               | :heavy_minus_sign:                                      | The ID of the merchant account to use for this request. |                                                         |