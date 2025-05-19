# ListPaymentServicesRequest


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `Method`                                                | [*components.Method](../../models/components/method.md) | :heavy_minus_sign:                                      | Return any payment service for this method.             |                                                         |
| `Cursor`                                                | **string*                                               | :heavy_minus_sign:                                      | A pointer to the page of results to return.             | ZXhhbXBsZTE                                             |
| `Limit`                                                 | **int64*                                                | :heavy_minus_sign:                                      | The maximum number of items that are at returned.       | 20                                                      |
| `Deleted`                                               | **bool*                                                 | :heavy_minus_sign:                                      | Return any deleted payment service.                     | true                                                    |
| `MerchantAccountID`                                     | **string*                                               | :heavy_minus_sign:                                      | The ID of the merchant account to use for this request. |                                                         |