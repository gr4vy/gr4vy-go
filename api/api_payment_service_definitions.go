/*
 * Gr4vy API (Beta)
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.0
 * Contact: code@gr4vy.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package Openapi

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// PaymentServiceDefinitionsApiService PaymentServiceDefinitionsApi service
type PaymentServiceDefinitionsApiService service

type ApiGetPaymentServiceDefinitionRequest struct {
	ctx _context.Context
	ApiService *PaymentServiceDefinitionsApiService
	paymentServiceDefinitionId string
}


func (r ApiGetPaymentServiceDefinitionRequest) Execute() (PaymentServiceDefinition, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentServiceDefinitionExecute(r)
}

/*
 * GetPaymentServiceDefinition Get payment service definition
 * Gets the definition for a payment service.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param paymentServiceDefinitionId The unique ID of the payment service definition.
 * @return ApiGetPaymentServiceDefinitionRequest
 */
func (a *PaymentServiceDefinitionsApiService) GetPaymentServiceDefinition(ctx _context.Context, paymentServiceDefinitionId string) ApiGetPaymentServiceDefinitionRequest {
	return ApiGetPaymentServiceDefinitionRequest{
		ApiService: a,
		ctx: ctx,
		paymentServiceDefinitionId: paymentServiceDefinitionId,
	}
}

/*
 * Execute executes the request
 * @return PaymentServiceDefinition
 */
func (a *PaymentServiceDefinitionsApiService) GetPaymentServiceDefinitionExecute(r ApiGetPaymentServiceDefinitionRequest) (PaymentServiceDefinition, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentServiceDefinition
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentServiceDefinitionsApiService.GetPaymentServiceDefinition")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-service-definitions/{payment_service_definition_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_service_definition_id"+"}", _neturl.PathEscape(parameterToString(r.paymentServiceDefinitionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404NotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListPaymentServiceDefinitionsRequest struct {
	ctx _context.Context
	ApiService *PaymentServiceDefinitionsApiService
	limit *int32
	cursor *string
}

func (r ApiListPaymentServiceDefinitionsRequest) Limit(limit int32) ApiListPaymentServiceDefinitionsRequest {
	r.limit = &limit
	return r
}
func (r ApiListPaymentServiceDefinitionsRequest) Cursor(cursor string) ApiListPaymentServiceDefinitionsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListPaymentServiceDefinitionsRequest) Execute() (PaymentServiceDefinitions, *_nethttp.Response, error) {
	return r.ApiService.ListPaymentServiceDefinitionsExecute(r)
}

/*
 * ListPaymentServiceDefinitions List payment service definitions
 * Returns a list of all available payment service definitions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListPaymentServiceDefinitionsRequest
 */
func (a *PaymentServiceDefinitionsApiService) ListPaymentServiceDefinitions(ctx _context.Context) ApiListPaymentServiceDefinitionsRequest {
	return ApiListPaymentServiceDefinitionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaymentServiceDefinitions
 */
func (a *PaymentServiceDefinitionsApiService) ListPaymentServiceDefinitionsExecute(r ApiListPaymentServiceDefinitionsRequest) (PaymentServiceDefinitions, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentServiceDefinitions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentServiceDefinitionsApiService.ListPaymentServiceDefinitions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-service-definitions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error401Unauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}