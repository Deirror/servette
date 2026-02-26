package transport

import "io"

// Transport defines the minimal interface to send a request and get a response.
type Transport interface {
	// Send sends data to the given path and returns the response body as a reader.
	// ContentType is optional (e.g., "application/json")
	Send(path string, contentType string, body io.Reader) (io.ReadCloser, error)

	// Get sends a GET request to the path and returns the response body.
	Get(path string) (io.ReadCloser, error)
}
