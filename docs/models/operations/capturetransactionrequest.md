# CaptureTransactionRequest


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `TransactionID`                                                                | *string*                                                                       | :heavy_check_mark:                                                             | N/A                                                                            |
| `MerchantAccountID`                                                            | **string*                                                                      | :heavy_minus_sign:                                                             | The ID of the merchant account to use for this request.                        |
| `TransactionCapture`                                                           | [components.TransactionCapture](../../models/components/transactioncapture.md) | :heavy_check_mark:                                                             | N/A                                                                            |