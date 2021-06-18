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
)

// Linger please
var (
	_ _context.Context
)

// PaymentOptionsApiService PaymentOptionsApi service
type PaymentOptionsApiService service

type ApiListPaymentOptionsRequest struct {
	ctx _context.Context
	ApiService *PaymentOptionsApiService
	country *string
	currency *string
	environment *string
	locale *string
}

func (r ApiListPaymentOptionsRequest) Country(country string) ApiListPaymentOptionsRequest {
	r.country = &country
	return r
}
func (r ApiListPaymentOptionsRequest) Currency(currency string) ApiListPaymentOptionsRequest {
	r.currency = &currency
	return r
}
func (r ApiListPaymentOptionsRequest) Environment(environment string) ApiListPaymentOptionsRequest {
	r.environment = &environment
	return r
}
func (r ApiListPaymentOptionsRequest) Locale(locale string) ApiListPaymentOptionsRequest {
	r.locale = &locale
	return r
}

func (r ApiListPaymentOptionsRequest) Execute() (PaymentOptions, *_nethttp.Response, error) {
	return r.ApiService.ListPaymentOptionsExecute(r)
}

/*
 * ListPaymentOptions List payment options
 * Returns a list of available payment method options for a currency
and country.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiListPaymentOptionsRequest
 */
func (a *PaymentOptionsApiService) ListPaymentOptions(ctx _context.Context) ApiListPaymentOptionsRequest {
	return ApiListPaymentOptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaymentOptions
 */
func (a *PaymentOptionsApiService) ListPaymentOptionsExecute(r ApiListPaymentOptionsRequest) (PaymentOptions, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaymentOptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PaymentOptionsApiService.ListPaymentOptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/payment-options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.country != nil {
		localVarQueryParams.Add("country", parameterToString(*r.country, ""))
	}
	if r.currency != nil {
		localVarQueryParams.Add("currency", parameterToString(*r.currency, ""))
	}
	if r.environment != nil {
		localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
	}
	if r.locale != nil {
		localVarQueryParams.Add("locale", parameterToString(*r.locale, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error400BadRequest
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
