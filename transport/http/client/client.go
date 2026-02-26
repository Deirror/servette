package client

import (
	"io"
	"net"
	"net/http"
	"time"
	"bytes"

	"github.com/Deirror/servette/encoding/json"
)

type Client struct {
	cl *http.Client

	baseURL string
}

func New(cfg *Config) *Client {
	return &Client{
		cl:      NewStdClient(cfg),
		baseURL: cfg.BaseURL,
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.cl.Get(c.baseURL + path)
}

func (c *Client) Post(path, contentType string, body io.Reader) (*http.Response, error) {
	return c.cl.Post(c.baseURL+path, contentType, body)
}

func (c *Client) PostJSON(path string, v any) (*http.Response, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return c.cl.Post(
		c.baseURL+path,
		"application/json",
		bytes.NewReader(data),
	)
}

func (c *Client) Put(path, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, c.baseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.cl.Do(req)
}

func (c *Client) Delete(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, c.baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	return c.cl.Do(req)
}

// New creates an *http.Client configured with the given Config.
func NewStdClient(cfg *Config) *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   cfg.ReadTimeout,
			KeepAlive: cfg.IdleTimeout,
		}).DialContext,
		MaxIdleConns:          100, // default value
		IdleConnTimeout:       cfg.IdleTimeout,
		TLSHandshakeTimeout:   5 * time.Second, // default value
		ExpectContinueTimeout: 1 * time.Second, // default value
	}

	return &http.Client{
		Timeout:   cfg.WriteTimeout,
		Transport: transport,
	}
}
