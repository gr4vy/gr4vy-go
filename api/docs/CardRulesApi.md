# \CardRulesApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCardRule**](CardRulesApi.md#AddCardRule) | **Post** /card-rules | Create card rule
[**DeleteCardRule**](CardRulesApi.md#DeleteCardRule) | **Delete** /card-rules/{card_rule_id} | Delete card rule
[**GetCardRule**](CardRulesApi.md#GetCardRule) | **Get** /card-rules/{card_rule_id} | Get card rule
[**ListCardsRules**](CardRulesApi.md#ListCardsRules) | **Get** /card-rules | List card rules
[**UpdateCardRule**](CardRulesApi.md#UpdateCardRule) | **Put** /card-rules/{card_rule_id} | Update card rule



## AddCardRule

> CardRule AddCardRule(ctx).CardRuleRequest(cardRuleRequest).Execute()

Create card rule



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
    cardRuleRequest := *openapiclient.NewCardRuleRequest([]openapiclient.CardRule{*openapiclient.NewCardRule()}, []string{"PaymentServiceIds_example"}) // CardRuleRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardRulesApi.AddCardRule(context.Background()).CardRuleRequest(cardRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardRulesApi.AddCardRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCardRule`: CardRule
    fmt.Fprintf(os.Stdout, "Response from `CardRulesApi.AddCardRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCardRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardRuleRequest** | [**CardRuleRequest**](CardRuleRequest.md) |  | 

### Return type

[**CardRule**](CardRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCardRule

> DeleteCardRule(ctx, cardRuleId).Execute()

Delete card rule



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
    cardRuleId := TODO // string | The unique ID for a card rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardRulesApi.DeleteCardRule(context.Background(), cardRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardRulesApi.DeleteCardRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardRuleId** | [**string**](.md) | The unique ID for a card rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCardRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCardRule

> CardRule GetCardRule(ctx, cardRuleId).Execute()

Get card rule



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
    cardRuleId := TODO // string | The unique ID for a card rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardRulesApi.GetCardRule(context.Background(), cardRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardRulesApi.GetCardRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCardRule`: CardRule
    fmt.Fprintf(os.Stdout, "Response from `CardRulesApi.GetCardRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardRuleId** | [**string**](.md) | The unique ID for a card rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCardRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CardRule**](CardRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCardsRules

> CardRules ListCardsRules(ctx).Limit(limit).Cursor(cursor).Environment(environment).Execute()

List card rules



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
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)
    environment := "staging" // string | Filters the results to only the items available in this environment. (optional) (default to "production")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardRulesApi.ListCardsRules(context.Background()).Limit(limit).Cursor(cursor).Environment(environment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardRulesApi.ListCardsRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCardsRules`: CardRules
    fmt.Fprintf(os.Stdout, "Response from `CardRulesApi.ListCardsRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCardsRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 
 **environment** | **string** | Filters the results to only the items available in this environment. | [default to &quot;production&quot;]

### Return type

[**CardRules**](CardRules.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCardRule

> CardRule UpdateCardRule(ctx, cardRuleId).CardRuleUpdate(cardRuleUpdate).Execute()

Update card rule



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
    cardRuleId := TODO // string | The unique ID for a card rule.
    cardRuleUpdate := *openapiclient.NewCardRuleUpdate() // CardRuleUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CardRulesApi.UpdateCardRule(context.Background(), cardRuleId).CardRuleUpdate(cardRuleUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CardRulesApi.UpdateCardRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCardRule`: CardRule
    fmt.Fprintf(os.Stdout, "Response from `CardRulesApi.UpdateCardRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cardRuleId** | [**string**](.md) | The unique ID for a card rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCardRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cardRuleUpdate** | [**CardRuleUpdate**](CardRuleUpdate.md) |  | 

### Return type

[**CardRule**](CardRule.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

