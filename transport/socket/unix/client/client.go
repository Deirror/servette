package client

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"time"

	"context"

	"github.com/Deirror/servette/encoding/json"
)

type Client struct {
	cl *http.Client
}

func New(cfg *Config) *Client {
	return &Client{
		cl: NewUnixClient(cfg),
	}
}

func (c *Client) Get(path string) (*http.Response, error) {
	return c.cl.Get(BaseURL + path)
}

func (c *Client) Post(path, contentType string, body io.Reader) (*http.Response, error) {
	return c.cl.Post(BaseURL+path, contentType, body)
}

func (c *Client) PostJSON(path string, v any) (*http.Response, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return c.cl.Post(
		BaseURL+path,
		"application/json",
		bytes.NewReader(data),
	)
}

func (c *Client) Put(path, contentType string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, BaseURL+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.cl.Do(req)
}

func (c *Client) Delete(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, BaseURL+path, nil)
	if err != nil {
		return nil, err
	}
	return c.cl.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req.URL.Host = "unix"
	req.URL.Scheme = "http"
	return c.cl.Do(req)
}

// NewUnixClient creates an *http.Client that communicates over a UNIX socket.
func NewUnixClient(cfg *Config) *http.Client {
	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return net.DialTimeout("unix", cfg.SocketPath, cfg.IdleTimeout)
		},
		IdleConnTimeout:       cfg.IdleTimeout,
		TLSHandshakeTimeout:   5 * time.Second, // default value
		ExpectContinueTimeout: 1 * time.Second, // default value
	}

	return &http.Client{
		Timeout:   cfg.WriteTimeout,
		Transport: transport,
	}
}
