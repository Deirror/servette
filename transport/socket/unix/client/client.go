package client

import (
	"net"
	"net/http"
	"time"

	"context"
)

// New creates an *http.Client that communicates over a UNIX socket.
func New(cfg *Config) *http.Client {
	transport := &http.Transport{
		DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
			return net.DialTimeout("unix", cfg.SocketPath, cfg.IdleTimeout)
		},
		IdleConnTimeout:       cfg.IdleTimeout,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Timeout:   cfg.WriteTimeout,
		Transport: transport,
	}
}
