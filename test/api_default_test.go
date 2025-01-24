package openapi

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"testing"

	securden "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_DefaultAPIService(t *testing.T) {
	// Set up the configuration
	cfg := securden.NewConfiguration()
	cfg.Host = "localhost:8000"
	cfg.Scheme = "https"
	cfg.Servers = securden.ServerConfigurations{
		{
			URL:         fmt.Sprintf("%s://%s/api", cfg.Scheme, cfg.Host),
			Description: "Updated API server URL",
		},
	}
	cfg.AddDefaultHeader("authtoken", "6f04c729-2ae4-46e4-8af4-f7820f2796e9")

	// Load the certificate file
	// certPath := `C:\\Program Files\\Securden\\Privileged_Account_Manager\\conf\\securden-cert-copy.pem`
	// certFile, err := os.ReadFile(certPath)
	// if err != nil {
	// 	log.Fatalf("Failed to read certificate file: %v", err)
	// }

	// // Create a certificate pool and append the cert
	// certPool := x509.NewCertPool()
	// if !certPool.AppendCertsFromPEM(certFile) {
	// 	log.Fatalf("Failed to add certificate to pool")
	// }

	// Create the HTTP client with custom transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			// RootCAs: certPool,
			InsecureSkipVerify: true,
		},
	}

	// Assign the custom client to the configuration
	cfg.HTTPClient = &http.Client{Transport: transport}

	// Create the API client
	client := securden.NewAPIClient(cfg)

	// Define account ID and context
	accountID := int64(3)
	ctx := context.Background()

	// Make the API request
	resp, httpResp, err := client.DefaultAPI.GetPasswordGet(ctx).AccountId(accountID).Execute()
	if err != nil {
		log.Fatalf("Error calling API: %v\nHTTP Response: %v", err, httpResp)
	}

	// Print the raw response
	fmt.Printf("Password response: %+v\n", *resp.Password)
}

