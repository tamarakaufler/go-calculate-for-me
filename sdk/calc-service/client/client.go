package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Log       *log.Logger

	httpClient *http.Client
}

type Output struct {
	Result uint64 `json:"result"`
}

// NewClient creates a client for access to the calculate service
func NewClient(u *url.URL, ua string, l *log.Logger) *Client {
	return &Client{
		BaseURL:    u,
		UserAgent:  ua,
		Log:        l,
		httpClient: &http.Client{},
	}
}

// Fibonacci makes a request to api-service for a fibonacci calculation
func (c *Client) Fibonacci(ctx context.Context, n uint64) (*Output, error) {
	u := c.prepareServiceURL("/fib", fmt.Sprintf("%d", n))

	output := &Output{}
	err := c.doRequest(ctx, u, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// Factorial makes a request to api-service for a factorial calculation
func (c *Client) Factorial(ctx context.Context, n uint64) (*Output, error) {
	u := c.prepareServiceURL("/fact", fmt.Sprintf("%d", n))

	output := &Output{}
	err := c.doRequest(ctx, u, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// GreatestCommonDenominator makes a request to api-service for a gcd calculation
func (c *Client) GreatestCommonDenominator(ctx context.Context, n, m uint64) (*Output, error) {
	args := fmt.Sprintf("/%d/%d", n, m)
	u := c.prepareServiceURL("/gcd", args)

	output := &Output{}
	err := c.doRequest(ctx, u, output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

// Helper functions and methods --------------------------------

func (c *Client) prepareServiceURL(path, args string) *url.URL {
	ur := &url.URL{
		Path: fmt.Sprintf("%s/%s", path, args),
	}
	fullURL := c.BaseURL.ResolveReference(ur)

	c.Log.Printf("fullURL = %v\n", fullURL)
	return fullURL
}

func (c *Client) doRequest(ctx context.Context, u *url.URL, output *Output) error {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(output)
	return nil
}
