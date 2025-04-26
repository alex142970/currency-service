package currency_client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const url = "https://latest.currency-api.pages.dev/v1/currencies/rub.json"

type RateResponse struct {
	Rate float64
	Date string
}

type CurrencyClient struct {
	httpClient *http.Client
}

func NewCurrencyClient(httpClient *http.Client) *CurrencyClient {
	return &CurrencyClient{
		httpClient: httpClient,
	}
}

func (c *CurrencyClient) FetchRate(base string, target string) (*RateResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var jsonBody interface{}

	_ = json.Unmarshal(body, &jsonBody)

	mapBody := jsonBody.(map[string]interface{})

	return &RateResponse{
		Rate: mapBody[base].(map[string]interface{})[target].(float64),
		Date: mapBody["date"].(string),
	}, nil
}
