# \CardSchemeDefinitionsApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCardSchemeDefinitions**](CardSchemeDefinitionsApi.md#ListCardSchemeDefinitions) | **Get** /card-scheme-definitions | List card scheme definitions



## ListCardSchemeDefinitions

> CardSchemeDefinitions ListCardSchemeDefinitions(ctx).Execute()

List card scheme definitions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardSchemeDefinitionsApi.ListCardSchemeDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardSchemeDefinitionsApi.ListCardSchemeDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCardSchemeDefinitions`: CardSchemeDefinitions
    fmt.Fprintf(os.Stdout, "Response from `CardSchemeDefinitionsApi.ListCardSchemeDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCardSchemeDefinitionsRequest struct via the builder pattern


### Return type

[**CardSchemeDefinitions**](CardSchemeDefinitions.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

