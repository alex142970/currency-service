package auth

import (
	"io/ioutil"
	"net/http"
)

type TokensClient struct {
	httpClient *http.Client
}

func NewTokensClient(httpClient *http.Client) *TokensClient {
	return &TokensClient{
		httpClient: httpClient,
	}
}

func (c *TokensClient) Generate(login string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8082/generate?login="+login, nil)

	if err != nil {
		return "", err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return "", err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (c *TokensClient) Validate(token string) bool {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8082/validate", nil)

	req.Header.Add("Authorization", "Bearer "+token)

	if err != nil {
		return false
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}
