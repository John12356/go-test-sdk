package securden_sdk

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
	accountCategory *int
	ticketId *int64
	reason *string
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

func (r ApiGetPasswordGetRequest) AccountCategory(accountCategory int) ApiGetPasswordGetRequest {
	r.accountCategory = &accountCategory
	return r
}

func (r ApiGetPasswordGetRequest) TicketId(ticketId int64) ApiGetPasswordGetRequest {
	r.ticketId = &ticketId
	return r
}

func (r ApiGetPasswordGetRequest) Reason(reason string) ApiGetPasswordGetRequest {
	r.reason = &reason
	return r
}

func (r ApiGetPasswordGetRequest) Execute() (GetPasswordGet200Response, string, error) {
	return r.ApiService.GetPasswordGetExecute(r)
}

/*
GetPasswordGet Retrieves password information

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetPasswordGetRequest
*/
func (a *DefaultAPIService) GetPassword(ctx context.Context) ApiGetPasswordGetRequest {
	return ApiGetPasswordGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetPasswordGet200Response
func (a *DefaultAPIService) GetPasswordGetExecute(r ApiGetPasswordGetRequest) (GetPasswordGet200Response, string, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  GetPasswordGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetPasswordGet")
	if err != nil {
		return localVarReturnValue, "", &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/get_password_via_tools"

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
	if r.accountCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "account_category", r.accountCategory, "form", "")
	}
	if r.ticketId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ticket_id", r.ticketId, "form", "")
	}
	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reason", r.reason, "form", "")
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
		return localVarReturnValue, "", err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, "", err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, "", err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		return localVarReturnValue, string(localVarBody), &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		return localVarReturnValue, "", &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
	}

	return localVarReturnValue, string(localVarBody), nil
}
