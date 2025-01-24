# \DefaultAPI

All URIs are relative to *http://https:/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordGet**](DefaultAPI.md#GetPasswordGet) | **Get** /get_password | Retrieves password information



## GetPasswordGet

> GetPasswordGet200Response GetPasswordGet(ctx).AccountId(accountId).AccountName(accountName).AccountTitle(accountTitle).AccountType(accountType).Execute()

Retrieves password information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountId := "accountId_example" // string |  (optional)
	accountName := "accountName_example" // string |  (optional)
	accountTitle := "accountTitle_example" // string |  (optional)
	accountType := "accountType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetPasswordGet(context.Background()).AccountId(accountId).AccountName(accountName).AccountTitle(accountTitle).AccountType(accountType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetPasswordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPasswordGet`: GetPasswordGet200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetPasswordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** |  | 
 **accountName** | **string** |  | 
 **accountTitle** | **string** |  | 
 **accountType** | **string** |  | 

### Return type

[**GetPasswordGet200Response**](GetPasswordGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

