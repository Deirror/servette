// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/Deirror/servette/transport"
)

// Server wraps an HTTP server over TCP or UNIX socket.
type Server struct {
	log *slog.Logger

	srv       *http.Server
	listener  net.Listener
	transType transport.Type
	endpoint  string
}

func New(cfg *Config, logger *slog.Logger, h http.Handler) (*Server, error) {
	srv, ln, err := NewHTTPServer(cfg, h)
	if err != nil {
		return nil, err
	}

	return &Server{
		log:       logger,
		srv:       srv,
		listener:  ln,
		transType: cfg.TransType,
		endpoint:  cfg.Endpoint,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	s.log.Info("HTTP Server starting",
		slog.String(Transport, string(s.transType)),
		slog.String(Endpoint, s.endpoint),
	)

	// Serve TLS if configured
	if s.srv.TLSConfig != nil {
		return s.srv.ServeTLS(s.listener, "", "")
	}

	return s.srv.Serve(s.listener)
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("HTTP Server shutting down")

	err := s.srv.Shutdown(ctx)

	// Cleanup UNIX socket file
	if s.transType.IsUDS() {
		_ = os.Remove(s.endpoint)
	}

	return err
}

// NewHTTPServer creates an http.Server listening on TCP or UNIX.
func NewHTTPServer(cfg *Config, handler http.Handler) (*http.Server, net.Listener, error) {
	network := transport.NetworkFromTransType(cfg.TransType)

	if cfg.TransType.IsUDS() {
		// Remove old socket file if exists
		if _, err := os.Stat(cfg.Endpoint); err == nil {
			if err := os.Remove(cfg.Endpoint); err != nil {
				return nil, nil, err
			}
		}

	}

	listener, err := net.Listen(network, cfg.Endpoint)
	if err != nil {
		return nil, nil, err
	}

	if cfg.TransType.IsUDS() {
		if err := os.Chmod(cfg.Endpoint, 0660); err != nil {
			listener.Close()
			os.Remove(cfg.Endpoint)
			return nil, nil, err
		}
	}

	server := &http.Server{
		Handler:           handler,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		MaxHeaderBytes:    cfg.MaxHeaderBytes,
		TLSConfig:         cfg.TLSConfig,
	}

	return server, listener, nil
}
