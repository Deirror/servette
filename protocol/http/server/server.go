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
	transType string
	addr      string
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
		addr:      cfg.Addr,
	}, nil
}

func (s *Server) Start(ctx context.Context) error {
	s.log.Info("HTTP Server starting",
		slog.String("transport", s.transType),
		slog.String("addr", s.addr),
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
	if s.transType == "unix" {
		_ = os.Remove(s.addr)
	}

	return err
}

// NewHTTPServer creates an http.Server listening on TCP or UNIX.
func NewHTTPServer(cfg *Config, handler http.Handler) (*http.Server, net.Listener, error) {
	network := transport.NetworkFromTransType(cfg.TransType)

	if network == "unix" {
		// Remove old socket file if exists
		if _, err := os.Stat(cfg.Addr); err == nil {
			if err := os.Remove(cfg.Addr); err != nil {
				return nil, nil, err
			}
		}

	}

	listener, err := net.Listen(network, cfg.Addr)
	if err != nil {
		return nil, nil, err
	}

	if network == "unix" {
		if err := os.Chmod(cfg.Addr, 0660); err != nil {
			listener.Close()
			os.Remove(cfg.Addr)
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
