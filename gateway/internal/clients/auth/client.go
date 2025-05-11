package auth

import (
	"context"
	"currency_service/gateway/internal/config"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const (
	endpointGenerate = "/generate"
	endpointValidate = "/validate"
	endpointPing     = "/ping"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func NewAuthClient(conf *config.Config) (*Client, error) {
	parsedURL, err := url.Parse(conf.AuthURL)

	if err != nil {
		return nil, fmt.Errorf("error parsing auth url: %w", err)
	}

	return &Client{
		baseURL: parsedURL,
		httpClient: &http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
	}, nil
}

func (c *Client) resolveUrl(endpoint string) *url.URL {
	return c.baseURL.ResolveReference(&url.URL{Path: endpoint})
}

func (c *Client) GenerateToken(ctx context.Context, login string) (string, error) {

	requestURL := c.resolveUrl(endpointGenerate)

	query := requestURL.Query()
	query.Set("login", login)
	requestURL.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)

	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("error executing request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error executing request, status code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(bodyBytes), nil
}

func (c *Client) ValidateToken(ctx context.Context, token string) error {

	requestURL := c.resolveUrl(endpointValidate)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)

	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("error executing request: %w", err)
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error executing request, status code: %d", resp.StatusCode)
	}

	return nil
}
