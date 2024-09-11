package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: "https://wmsc.lcsc.com",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type APIResponse struct {
	Code   int             `json:"code"`
	Msg    interface{}     `json:"msg"`
	Result json.RawMessage `json:"result"`
}

func (c *Client) makeRequest(method, path string, body interface{}, headers map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, path)

	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var apiResp APIResponse
	err = json.Unmarshal(respBody, &apiResp)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	if apiResp.Code != 200 {
		return nil, fmt.Errorf("API error - Code: %d, Message: %v", apiResp.Code, apiResp.Msg)
	}

	return apiResp.Result, nil
}
