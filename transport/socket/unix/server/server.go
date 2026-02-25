package server

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
)

// Server wraps an HTTP server over a UNIX socket.
type Server struct {
	log *slog.Logger

	srv      *http.Server
	listener net.Listener
	socket   string
}

func New(cfg *Config, log *slog.Logger, h http.Handler) (*Server, error) {
	srv, ln, err := NewUnixServer(cfg, h)
	if err != nil {
		return nil, err
	}

	return &Server{
		log:      log,
		srv:      srv,
		listener: ln,
		socket:   cfg.Port,
	}, nil
}

// Start begins serving HTTP requests on the UNIX socket.
func (s *Server) Start(ctx context.Context) error {
	s.log.Info("UNIX Socket Server starting", slog.String("socket", s.socket))

	return s.srv.Serve(s.listener)
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.log.Info("UNIX Socket Server shutting down")

	defer s.listener.Close()
	defer os.Remove(s.socket)

	return s.srv.Shutdown(ctx)
}

// NewUnixServer creates and returns an http.Server listening on a UNIX socket.
// It expects cfg.Port to contain the socket path.
func NewUnixServer(cfg *Config, handler http.Handler) (*http.Server, net.Listener, error) {
	socketPath := cfg.Port

	// Remove existing socket file if present
	if _, err := os.Stat(socketPath); err == nil {
		if err := os.Remove(socketPath); err != nil {
			return nil, nil, err
		}
	}

	// Create a UNIX socket listener
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, nil, err
	}

	// Set socket file permissions (e.g. rw for owner & group)
	if err := os.Chmod(socketPath, 0660); err != nil {
		listener.Close()
		os.Remove(socketPath)
		return nil, nil, err
	}

	server := &http.Server{
		Handler:      handler,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	return server, listener, nil
}
