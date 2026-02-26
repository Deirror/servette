package unix

import (
	"io"

	"github.com/Deirror/servette/transport/socket/unix/client"
)

type Transport struct {
	cl *client.Client
}

func NewTransport(cl *client.Client) *Transport {
	return &Transport{
		cl: cl,
	}
}

// Send implements Transport.Send (POST/PUT/any payload)
func (u *Transport) Send(path string, contentType string, body io.Reader) (io.ReadCloser, error) {
	resp, err := u.cl.Post(path, contentType, body)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

// Get implements Transport.Get
func (u *Transport) Get(path string) (io.ReadCloser, error) {
	resp, err := u.cl.Get(path)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
