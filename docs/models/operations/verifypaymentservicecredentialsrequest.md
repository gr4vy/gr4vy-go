# VerifyPaymentServiceCredentialsRequest


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `TimeoutInSeconds`                                                           | **float64*                                                                   | :heavy_minus_sign:                                                           | N/A                                                                          |
| `MerchantAccountID`                                                          | **string*                                                                    | :heavy_minus_sign:                                                           | The ID of the merchant account to use for this request.                      |
| `VerifyCredentials`                                                          | [components.VerifyCredentials](../../models/components/verifycredentials.md) | :heavy_check_mark:                                                           | N/A                                                                          |