# ReportExecutionURL


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `URL`                                                    | *string*                                                 | :heavy_check_mark:                                       | A signed URL to download the report execution file.      | https://example.com/download/report.csv?signature=abc123 |
| `ExpiresAt`                                              | [time.Time](https://pkg.go.dev/time#Time)                | :heavy_check_mark:                                       | The date and time when the download URL expires.         | 2024-06-01T00:00:00.000Z                                 |