# TransactionEvent


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Type`                                             | **string*                                          | :heavy_minus_sign:                                 | Always `transaction-event`.                        | transaction-event                                  |
| `ID`                                               | *string*                                           | :heavy_check_mark:                                 | The ID for the event.                              | f133a3b7-e67e-4d83-bcd3-3e438fedf348               |
| `Name`                                             | [components.Name](../../models/components/name.md) | :heavy_check_mark:                                 | The specific event name.                           | transaction-api-request                            |
| `CreatedAt`                                        | [time.Time](https://pkg.go.dev/time#Time)          | :heavy_check_mark:                                 | The date this event was created at.                | 2013-07-16T19:23:00.000+00:00                      |
| `Context`                                          | map[string]*any*                                   | :heavy_check_mark:                                 | N/A                                                |                                                    |