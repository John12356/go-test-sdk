package openapi

import (
	"context"
	"fmt"
	"log"
	"testing"

	securden "github.com/John12356/go-test-sdk"
)

func Test_openapi_DefaultAPIService(t *testing.T) {
	// Set up the configuration
	BaseURL := "https://ruthra.securden.com:5454"
	cfg := securden.NewConfiguration(BaseURL)
	// cfg.SetAuthToken("6f04c729-2ae4-46e4-8af4-f7820f2796e9") //msp local
	// cfg.SetAuthToken("2c1e3958-ef6b-4855-abb0-7e03ec7852a5") //pam exe
	cfg.SetAuthToken("22d1ed59-33c0-4852-a132-b3312705e386") //vault ca

	// Load the certificate file
	// cfg.SSLCert = `C:\\Users\\RuthraMoorthyK\\Documents\\CA_Certs\\securden-cert.pem`
	// cfg.SSLCert = `C:\\Program Files\\Securden\\Privileged_Account_Manager\\conf\\securden-cert-copy.pem`
	// cfg.InsecureSkipVerify = true

	err := cfg.SetSSLConfig()

	// certFile, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatalf("Failed to read certificateeeee file: %v", err)
	}

	// Create the API client
	client := securden.NewAPIClient(cfg)

	// Define account ID and context
	accountID := int64(2000000002379)
	// accountID := int64(3)
	// accountTitle := "sss"
	// accountCategory := 2

	ctx := context.Background()

	// Make the API request
	resp, httpResp, err := client.DefaultAPI.GetPassword(ctx).
	// AccountTitle(accountTitle).
	// AccountCategory(client.WorkAccountType).
	// AccountType("windows member").
	AccountId(accountID).Execute()
	if err != nil {
		log.Fatalf("Error calling API: %v\nHTTP Response: %v", err, httpResp)
	}

	// Print the raw response
	fmt.Printf("Password response: %+v\n", resp.GetPassword())
}

