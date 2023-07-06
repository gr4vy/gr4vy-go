# \RolesApi

All URIs are relative to *https://api.plantly.gr4vy.app*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRoleAssignment**](RolesApi.md#DeleteRoleAssignment) | **Delete** /roles/assignments/{role_assignment_id} | Delete role assignment
[**ListRoleAssignments**](RolesApi.md#ListRoleAssignments) | **Get** /roles/assignments | List role assignments
[**ListRoles**](RolesApi.md#ListRoles) | **Get** /roles | List roles
[**NewRoleAssignment**](RolesApi.md#NewRoleAssignment) | **Post** /roles/assignments | New role assignment



## DeleteRoleAssignment

> DeleteRoleAssignment(ctx, roleAssignmentId).Execute()

Delete role assignment



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
    roleAssignmentId := TODO // string | The unique ID for the role assignment.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.DeleteRoleAssignment(context.Background(), roleAssignmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.DeleteRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleAssignmentId** | [**string**](.md) | The unique ID for the role assignment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoleAssignmentRequest struct via the builder pattern


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


## ListRoleAssignments

> RoleAssignments ListRoleAssignments(ctx).RoleId(roleId).AssigneeType(assigneeType).AssigneeId(assigneeId).Limit(limit).Cursor(cursor).Execute()

List role assignments



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
    roleId := TODO // string | Filters for role assignments for the role that has a matching `id` value. (optional)
    assigneeType := "user" // string | Filters for role assignments for the assignee of the given type. (optional)
    assigneeId := TODO // string | Filters for role assignments for the assignee that has a matching `id` value. The `assignee_type` must also be specified. (optional)
    limit := int32(1) // int32 | Defines the maximum number of items to return for this request. (optional) (default to 20)
    cursor := "ZXhhbXBsZTE" // string | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the `next_cursor` field. Similarly the `previous_cursor` can be used to reverse backwards in the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.ListRoleAssignments(context.Background()).RoleId(roleId).AssigneeType(assigneeType).AssigneeId(assigneeId).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoleAssignments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleAssignments`: RoleAssignments
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoleAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | [**string**](string.md) | Filters for role assignments for the role that has a matching &#x60;id&#x60; value. | 
 **assigneeType** | **string** | Filters for role assignments for the assignee of the given type. | 
 **assigneeId** | [**string**](string.md) | Filters for role assignments for the assignee that has a matching &#x60;id&#x60; value. The &#x60;assignee_type&#x60; must also be specified. | 
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**RoleAssignments**](RoleAssignments.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoles

> Roles ListRoles(ctx).Limit(limit).Cursor(cursor).Execute()

List roles



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.ListRoles(context.Background()).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.ListRoles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoles`: Roles
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.ListRoles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Defines the maximum number of items to return for this request. | [default to 20]
 **cursor** | **string** | A cursor that identifies the page of results to return. This is used to paginate the results of this API.  For the first page of results, this parameter can be left out. For additional pages, use the value returned by the API in the &#x60;next_cursor&#x60; field. Similarly the &#x60;previous_cursor&#x60; can be used to reverse backwards in the list. | 

### Return type

[**Roles**](Roles.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NewRoleAssignment

> RoleAssignment NewRoleAssignment(ctx).RoleAssignmentRequest(roleAssignmentRequest).Execute()

New role assignment



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
    roleAssignmentRequest := *openapiclient.NewRoleAssignmentRequest(*openapiclient.NewRoleAssignmentRequestRole("462ab2e2-3e29-44bd-b39f-e4d1293affbb"), *openapiclient.NewRoleAssignmentRequestAssignee("user", "42aae896-8ce2-4a60-b80a-5f6ae1dfbbd4")) // RoleAssignmentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.RolesApi.NewRoleAssignment(context.Background()).RoleAssignmentRequest(roleAssignmentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RolesApi.NewRoleAssignment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewRoleAssignment`: RoleAssignment
    fmt.Fprintf(os.Stdout, "Response from `RolesApi.NewRoleAssignment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleAssignmentRequest** | [**RoleAssignmentRequest**](RoleAssignmentRequest.md) |  | 

### Return type

[**RoleAssignment**](RoleAssignment.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

