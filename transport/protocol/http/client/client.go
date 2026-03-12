// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/Deirror/servette/encoding/json"
	"github.com/Deirror/servette/transport"
	"github.com/Deirror/servette/transport/protocol/http/header"
)

type Doer interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	cl       Doer
	endpoint string
}

func New(cfg *Config) *Client {
	endpoint := transport.EndpointFromTransType(cfg.Endpoint, cfg.TransType)
	return &Client{
		cl:       NewHTTPClient(cfg),
		endpoint: endpoint,
	}
}

func (c *Client) DoRequest(method, path string, body io.Reader, opts ...RequestOpt) (*http.Response, error) {
	req, err := http.NewRequest(method, c.endpoint+path, body)
	if err != nil {
		return nil, err
	}

	// Apply all options
	for _, opt := range opts {
		opt(req)
	}

	return c.cl.Do(req)
}

func (c *Client) Get(path string, opts ...RequestOpt) (*http.Response, error) {
	return c.DoRequest(http.MethodGet, path, nil, opts...)
}

func (c *Client) Post(path string, body io.Reader, contentType string, opts ...RequestOpt) (*http.Response, error) {
	if contentType != "" {
		opts = append(opts, WithHeader(header.ContentType, contentType))
	}
	return c.DoRequest(http.MethodPost, path, body, opts...)
}

func (c *Client) PostJSON(path string, v any, opts ...RequestOpt) (*http.Response, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	opts = append(opts, WithHeader(header.ContentType, header.ApplicationJSON))
	return c.DoRequest(http.MethodPost, path, bytes.NewReader(data), opts...)
}

func (c *Client) Put(path string, body io.Reader, contentType string, opts ...RequestOpt) (*http.Response, error) {
	if contentType != "" {
		opts = append(opts, WithHeader(header.ContentType, contentType))
	}
	return c.DoRequest(http.MethodPut, path, body, opts...)
}

func (c *Client) Delete(path string, opts ...RequestOpt) (*http.Response, error) {
	return c.DoRequest(http.MethodDelete, path, nil, opts...)
}

// NewHTTPClient creates an *http.Client configured with the given Config.
func NewHTTPClient(cfg *Config) *http.Client {
	transport := &http.Transport{
		DialContext:           transport.NewDialContext(transport.NetworkFromTransType(cfg.TransType), cfg.Endpoint, cfg.DialTimeout),
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          cfg.MaxIdleConns,
		MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
		MaxConnsPerHost:       cfg.MaxConnsPerHost,
		IdleConnTimeout:       cfg.IdleConnTimeout,
		TLSHandshakeTimeout:   5 * time.Second, // default
		ExpectContinueTimeout: 1 * time.Second, // default
		ResponseHeaderTimeout: cfg.ResponseHeaderTimeout,
		TLSClientConfig:       cfg.TLSConfig,
		Proxy:                 http.ProxyFromEnvironment,
		DisableCompression:    false,
		DisableKeepAlives:     false,
		WriteBufferSize:       0, // default
		ReadBufferSize:        0, // default
	}

	jar, _ := cookiejar.New(nil)

	return &http.Client{
		Timeout:   cfg.RequestTimeout,
		Transport: transport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if cfg.MaxRedirects > 0 && len(via) >= cfg.MaxRedirects {
				return fmt.Errorf("stopped after %d redirects", cfg.MaxRedirects)
			}
			return nil
		},
		Jar: jar,
	}
}
