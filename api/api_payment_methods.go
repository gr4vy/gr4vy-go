/*
 * Gr4vy API
 *
 * Welcome to the Gr4vy API reference documentation. Our API is still very much a work in product and subject to change.
 *
 * API version: 1.1.0-beta
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

// PaymentMethodsApiService PaymentMethodsApi service
type PaymentMethodsApiService service

type ApiDeletePaymentMethodRequest struct {
	ctx _context.Context
	ApiService *PaymentMethodsApiService
	paymentMethodId string
}


func (r ApiDeletePaymentMethodRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeletePaymentMethodExecute(r)
}

/*
 * DeletePaymentMethod Delete payment method
 * Removes a stored payment method.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param paymentMethodId The ID of the payment method.
 * @return ApiDeletePaymentMethodRequest
 */
func (a *PaymentMethodsApiService) DeletePaymentMethod(ctx _context.Context, paymentMethodId string) ApiDeletePaymentMethodRequest {
	return ApiDeletePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

/*
 * Execute executes the request
 */
func (a *PaymentMethodsApiService) DeletePaymentMethodExecute(r ApiDeletePaymentMethodRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.DeletePaymentMethod")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-methods/{payment_method_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_method_id"+"}", _neturl.PathEscape(parameterToString(r.paymentMethodId, "")), -1)

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error404NotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetPaymentMethodRequest struct {
	ctx _context.Context
	ApiService *PaymentMethodsApiService
	paymentMethodId string
}


func (r ApiGetPaymentMethodRequest) Execute() (PaymentMethod, *_nethttp.Response, error) {
	return r.ApiService.GetPaymentMethodExecute(r)
}

/*
 * GetPaymentMethod Get stored payment method
 * Gets the details for a stored payment method.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param paymentMethodId The ID of the payment method.
 * @return ApiGetPaymentMethodRequest
 */
func (a *PaymentMethodsApiService) GetPaymentMethod(ctx _context.Context, paymentMethodId string) ApiGetPaymentMethodRequest {
	return ApiGetPaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
		paymentMethodId: paymentMethodId,
	}
}

/*
 * Execute executes the request
 * @return PaymentMethod
 */
func (a *PaymentMethodsApiService) GetPaymentMethodExecute(r ApiGetPaymentMethodRequest) (PaymentMethod, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.GetPaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-methods/{payment_method_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"payment_method_id"+"}", _neturl.PathEscape(parameterToString(r.paymentMethodId, "")), -1)

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

type ApiListBuyerPaymentMethodsRequest struct {
	ctx _context.Context
	ApiService *PaymentMethodsApiService
	buyerId *string
	buyerExternalIdentifier *string
	country *string
	currency *string
}

func (r ApiListBuyerPaymentMethodsRequest) BuyerId(buyerId string) ApiListBuyerPaymentMethodsRequest {
	r.buyerId = &buyerId
	return r
}
func (r ApiListBuyerPaymentMethodsRequest) BuyerExternalIdentifier(buyerExternalIdentifier string) ApiListBuyerPaymentMethodsRequest {
	r.buyerExternalIdentifier = &buyerExternalIdentifier
	return r
}
func (r ApiListBuyerPaymentMethodsRequest) Country(country string) ApiListBuyerPaymentMethodsRequest {
	r.country = &country
	return r
}
func (r ApiListBuyerPaymentMethodsRequest) Currency(currency string) ApiListBuyerPaymentMethodsRequest {
	r.currency = &currency
	return r
}

func (r ApiListBuyerPaymentMethodsRequest) Execute() (PaymentMethodsTokenized, *_nethttp.Response, error) {
	return r.ApiService.ListBuyerPaymentMethodsExecute(r)
}

/*
 * ListBuyerPaymentMethods List stored payment methods for a buyer
 * Returns a list of stored payment methods for a buyer in a summarized format.
Only payment methods that are compatible with at least one active payment
service in that region are shown.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListBuyerPaymentMethodsRequest
 */
func (a *PaymentMethodsApiService) ListBuyerPaymentMethods(ctx _context.Context) ApiListBuyerPaymentMethodsRequest {
	return ApiListBuyerPaymentMethodsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaymentMethodsTokenized
 */
func (a *PaymentMethodsApiService) ListBuyerPaymentMethodsExecute(r ApiListBuyerPaymentMethodsRequest) (PaymentMethodsTokenized, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentMethodsTokenized
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.ListBuyerPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/buyers/payment-methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.buyerId != nil {
		localVarQueryParams.Add("buyer_id", parameterToString(*r.buyerId, ""))
	}
	if r.buyerExternalIdentifier != nil {
		localVarQueryParams.Add("buyer_external_identifier", parameterToString(*r.buyerExternalIdentifier, ""))
	}
	if r.country != nil {
		localVarQueryParams.Add("country", parameterToString(*r.country, ""))
	}
	if r.currency != nil {
		localVarQueryParams.Add("currency", parameterToString(*r.currency, ""))
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

type ApiListPaymentMethodsRequest struct {
	ctx _context.Context
	ApiService *PaymentMethodsApiService
	buyerId *string
	buyerExternalIdentifier *string
	status *string
	limit *int32
	cursor *string
}

func (r ApiListPaymentMethodsRequest) BuyerId(buyerId string) ApiListPaymentMethodsRequest {
	r.buyerId = &buyerId
	return r
}
func (r ApiListPaymentMethodsRequest) BuyerExternalIdentifier(buyerExternalIdentifier string) ApiListPaymentMethodsRequest {
	r.buyerExternalIdentifier = &buyerExternalIdentifier
	return r
}
func (r ApiListPaymentMethodsRequest) Status(status string) ApiListPaymentMethodsRequest {
	r.status = &status
	return r
}
func (r ApiListPaymentMethodsRequest) Limit(limit int32) ApiListPaymentMethodsRequest {
	r.limit = &limit
	return r
}
func (r ApiListPaymentMethodsRequest) Cursor(cursor string) ApiListPaymentMethodsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiListPaymentMethodsRequest) Execute() (PaymentMethods, *_nethttp.Response, error) {
	return r.ApiService.ListPaymentMethodsExecute(r)
}

/*
 * ListPaymentMethods List payment methods
 * Returns a list of stored payment methods.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListPaymentMethodsRequest
 */
func (a *PaymentMethodsApiService) ListPaymentMethods(ctx _context.Context) ApiListPaymentMethodsRequest {
	return ApiListPaymentMethodsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaymentMethods
 */
func (a *PaymentMethodsApiService) ListPaymentMethodsExecute(r ApiListPaymentMethodsRequest) (PaymentMethods, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentMethods
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.ListPaymentMethods")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.buyerId != nil {
		localVarQueryParams.Add("buyer_id", parameterToString(*r.buyerId, ""))
	}
	if r.buyerExternalIdentifier != nil {
		localVarQueryParams.Add("buyer_external_identifier", parameterToString(*r.buyerExternalIdentifier, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
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

type ApiStorePaymentMethodRequest struct {
	ctx _context.Context
	ApiService *PaymentMethodsApiService
	paymentMethodRequest *PaymentMethodRequest
}

func (r ApiStorePaymentMethodRequest) PaymentMethodRequest(paymentMethodRequest PaymentMethodRequest) ApiStorePaymentMethodRequest {
	r.paymentMethodRequest = &paymentMethodRequest
	return r
}

func (r ApiStorePaymentMethodRequest) Execute() (PaymentMethod, *_nethttp.Response, error) {
	return r.ApiService.StorePaymentMethodExecute(r)
}

/*
 * StorePaymentMethod New payment method
 * Stores and vaults a new payment method.

Vaulting a card only stores its information but doesn't validate it against any
PSP. In order to do so, a CIT (Customer Initiated Transaction) must be done,
even if it's a zero-value one.

 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiStorePaymentMethodRequest
 */
func (a *PaymentMethodsApiService) StorePaymentMethod(ctx _context.Context) ApiStorePaymentMethodRequest {
	return ApiStorePaymentMethodRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaymentMethod
 */
func (a *PaymentMethodsApiService) StorePaymentMethodExecute(r ApiStorePaymentMethodRequest) (PaymentMethod, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentMethod
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentMethodsApiService.StorePaymentMethod")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-methods"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.paymentMethodRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorGeneric
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v Error409DuplicateRecord
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
