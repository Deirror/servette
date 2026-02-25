package client

import (
	"net"
	"net/http"
	"time"
)

// New creates an http.Client configured with the given Config.
func New(cfg *Config) *http.Client {
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
