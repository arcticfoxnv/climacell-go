package climacell

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	UserAgentString = fmt.Sprintf("climacell-go/%s (https://github.com/arcticfoxnv/climacell-go)", Version)

	DefaultRequestDefaults = RequestDefaults{
		UnitSystem: SI,
	}
)

type Client struct {
	ApiKey    string
	Defaults  RequestDefaults
	UserAgent string

	httpClient *http.Client
}

type RequestDefaults struct {
	UnitSystem UnitSystem
}

type Option func(*Client)

func SetHTTPClient(httpClient *http.Client) func(*Client) {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func NewClient(apiKey string, options ...Option) *Client {
	c := &Client{
		ApiKey:    apiKey,
		Defaults:  DefaultRequestDefaults,
		UserAgent: UserAgentString,
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
	}

	for _, option := range options {
		option(c)
	}

	return c
}

func (c *Client) getEndpoint(version, endpoint string) string {
	var base string

	switch version {
	case "v3":
		base = ClimacellV3
	default:
		base = ClimacellV3
	}

	return fmt.Sprintf("%s/%s", base, endpoint)
}

func (c *Client) appendQueryParam(req *http.Request, k, v string) {
	q := req.URL.Query()
	q.Set(k, v)
	req.URL.RawQuery = q.Encode()
}

func (c *Client) newGetRequest(version, endpoint string, params Request) (*http.Request, error) {
	url := c.getEndpoint(version, endpoint)
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.URL.RawQuery = params.ToQuery(&c.Defaults).Encode()

	return req, nil
}

func (c *Client) newPostRequest(version, endpoint string, body interface{}) (*http.Request, error) {
	url := c.getEndpoint(version, endpoint)
	return c.newRequest("POST", url, body)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, _ := http.NewRequest(method, path, buf)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) setAuth(req *http.Request) {
	c.appendQueryParam(req, "apikey", c.ApiKey)
}

func (c *Client) do(req *http.Request, data interface{}) error {
	c.setAuth(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}
