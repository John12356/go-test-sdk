package securden_sdk

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	BaseUrl          string             `json:"baseUrl,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	AuthToken        string            `json:"authtoken,omitempty"`
	SSLCert           string            `json:"sslCert,omitempty"`
	InsecureSkipVerify bool             `json:"insecureSkipVerify,omitempty"` 
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

func NewConfiguration(baseURL string) *Configuration {
	cfg := &Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "Securden-SDK/1.0.0/go",
		Debug:         false,
		Servers: ServerConfigurations{
			{
				URL:         fmt.Sprintf("%s/secretsmanagement", strings.TrimSuffix(baseURL, "/")),
				Description: "Updated API server URL",
			},
		},
		OperationServers: map[string]ServerConfigurations{},
	}
	cfg.BaseUrl = baseURL
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) SetAuthToken(token string) {
	c.AuthToken = token
	c.DefaultHeader["authtoken"] = token
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

func (c *Configuration) SetSSLConfig() error {
	// If InsecureSkipVerify is true, configure the HTTP client to skip verification
	if c.InsecureSkipVerify {
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		c.HTTPClient = &http.Client{Transport: transport}
		return nil
	}

	// Parse the BaseURL to extract the host and port
	parsedURL, err := url.Parse(c.BaseUrl)
	if err != nil {
		return fmt.Errorf("invalid BaseURL: %v", err)
	}

	host := parsedURL.Host
	if !strings.Contains(host, ":") {
		// Default to port 443 if no port is specified
		host = fmt.Sprintf("%s:443", host)
	}

	// If SSLCert is empty, fetch the certificate from the server
	if c.SSLCert == "" {
		// Fetch the certificate from the server
		conn, err := tls.Dial("tcp", host, &tls.Config{ InsecureSkipVerify: true })
		if err != nil {
			return fmt.Errorf("failed to fetch certificate ffrom server: %v", err)
		}
		defer conn.Close()

		// Retrieve the peer certificates
		if len(conn.ConnectionState().PeerCertificates) == 0 {
			return fmt.Errorf("no certificates found on the server")
		}

		// Use the first certificate from the server
		serverCert := conn.ConnectionState().PeerCertificates[0]
		// Create a certificate pool and add the server certificate
		certPool := x509.NewCertPool()
		certPool.AddCert(serverCert)

		// Create the HTTP client with the custom transport
		transport := &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: certPool,
			},
		}
		c.HTTPClient = &http.Client{Transport: transport}
		return nil
	}

	// If SSLCert is specified, read the certificate file
	certFile, err := os.ReadFile(c.SSLCert)
	if err != nil {
		return fmt.Errorf("failed to read certificatesss file: %v", err)
	}

	// Create a certificate pool and append the cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(certFile) {
		return fmt.Errorf("failed to add certificate to pool")
	}

	// Create the HTTP client with custom transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: certPool,
		},
	}
	c.HTTPClient = &http.Client{Transport: transport}
	return nil
}
