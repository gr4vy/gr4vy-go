# PaymentMethodUpdate

Request body for updating a stored payment method.


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ExpirationDate`                                                                 | `*string`                                                                        | :heavy_minus_sign:                                                               | The new expiration date for the payment method.                                  | 12/30                                                                            |
| `SchemeTransactionID`                                                            | `*string`                                                                        | :heavy_minus_sign:                                                               | A scheme transaction identifier to associate with this payment method.           | 123456789012345                                                                  |
| `SchemeTransactionIDScheme`                                                      | [*components.CardScheme](../../models/components/cardscheme.md)                  | :heavy_minus_sign:                                                               | The scheme associated with scheme_transaction_id. Only applies to card payments. | visa                                                                             |