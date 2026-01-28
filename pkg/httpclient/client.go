package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	RequestIDKey     = "X-Request-ID"
	DefaultTimeout   = 30 * time.Second
	ContentTypeJSON  = "application/json"
)

// Client is an HTTP client that propagates X-Request-ID
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// Option configures the Client
type Option func(*Client)

// WithTimeout sets the client timeout
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// WithBaseURL sets the base URL for all requests
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// New creates a new HTTP client
func New(opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// RequestOption configures individual requests
type RequestOption func(*http.Request)

// WithHeader adds a header to the request
func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set(key, value)
	}
}

// WithAuth adds Authorization header
func WithAuth(token string) RequestOption {
	return func(r *http.Request) {
		r.Header.Set("Authorization", "Bearer "+token)
	}
}

// Do executes an HTTP request with request_id propagation
func (c *Client) Do(ctx context.Context, method, path string, body interface{}, opts ...RequestOption) (*http.Response, error) {
	url := path
	if c.baseURL != "" {
		url = c.baseURL + path
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	// Set default headers
	req.Header.Set("Content-Type", ContentTypeJSON)

	// Propagate X-Request-ID from context
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		req.Header.Set(RequestIDKey, requestID.(string))
	}

	// Apply custom options
	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, opts ...RequestOption) (*http.Response, error) {
	return c.Do(ctx, http.MethodGet, path, nil, opts...)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*http.Response, error) {
	return c.Do(ctx, http.MethodPost, path, body, opts...)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, body interface{}, opts ...RequestOption) (*http.Response, error) {
	return c.Do(ctx, http.MethodPut, path, body, opts...)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string, opts ...RequestOption) (*http.Response, error) {
	return c.Do(ctx, http.MethodDelete, path, nil, opts...)
}
