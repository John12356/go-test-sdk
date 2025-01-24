/*
Password Retrieval API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// DefaultAPIService DefaultAPI service
type DefaultAPIService service

type ApiGetPasswordGetRequest struct {
	ctx context.Context
	ApiService *DefaultAPIService
	accountId *int64
	accountName *string
	accountTitle *string
	accountType *string
}

func (r ApiGetPasswordGetRequest) AccountId(accountId int64) ApiGetPasswordGetRequest {
	r.accountId = &accountId
	return r
}

func (r ApiGetPasswordGetRequest) AccountName(accountName string) ApiGetPasswordGetRequest {
	r.accountName = &accountName
	return r
}

func (r ApiGetPasswordGetRequest) AccountTitle(accountTitle string) ApiGetPasswordGetRequest {
	r.accountTitle = &accountTitle
	return r
}

func (r ApiGetPasswordGetRequest) AccountType(accountType string) ApiGetPasswordGetRequest {
	r.accountType = &accountType
	return r
}

func (r ApiGetPasswordGetRequest) Execute() (*GetPasswordGet200Response, *http.Response, error) {
	return r.ApiService.GetPasswordGetExecute(r)
}

/*
GetPasswordGet Retrieves password information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPasswordGetRequest
*/
func (a *DefaultAPIService) GetPasswordGet(ctx context.Context) ApiGetPasswordGetRequest {
	return ApiGetPasswordGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPasswordGet200Response
func (a *DefaultAPIService) GetPasswordGetExecute(r ApiGetPasswordGetRequest) (*GetPasswordGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetPasswordGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetPasswordGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/get_password"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.accountId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_id", r.accountId, "form", "")
	}
	if r.accountName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_name", r.accountName, "form", "")
	}
	if r.accountTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_title", r.accountTitle, "form", "")
	}
	if r.accountType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_type", r.accountType, "form", "")
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
