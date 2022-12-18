package thawani

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// NewClient creates a new FusionAuthClient
// if httpClient is nil then a DefaultClient is used
func NewClient(httpClient *http.Client, baseURL *url.URL, apiKey string, publishableKey string) *ThawaniClient {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 5 * time.Minute,
		}
	}
	c := &ThawaniClient{
		HTTPClient:     httpClient,
		BaseURL:        baseURL,
		APIKey:         apiKey,
		PublishableKey: publishableKey,
	}

	return c
}

func (c *ThawaniClient) CreateCustomer(body CreateCustomerReq) (resp *CreateCustomerResp, err error) {
	req, err := c.newRequest("POST", "/api/v1/customers", body)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &resp)
	return resp, err
}

func (c *ThawaniClient) CreateSession(body CreateSessionReq) (resp *Session, redirectTo string, error error) {
	req, err := c.newRequest("POST", "/api/v1/checkout/session", body)
	if err != nil {
		return nil, redirectTo, err
	}

	_, err = c.do(req, &resp)
	return resp, fmt.Sprintf("%s/pay/%s?key=%s", c.BaseURL.String(), resp.Data.SessionId, c.PublishableKey), err
}

func (c *ThawaniClient) GetSessionByClientReference(clientReference string) (resp *Session, error error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/v1/checkout/reference/%s", clientReference), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &resp)
	return resp, err
}

func (c *ThawaniClient) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("thawani-api-key", c.APIKey)
	return req, nil
}

func (c *ThawaniClient) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
