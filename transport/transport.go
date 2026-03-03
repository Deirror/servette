package transport

import (
	"context"
	"net"
	"time"
)

type DialContextFunc = func(ctx context.Context, _, addr string) (net.Conn, error)

func NewDialContext(network, addr string, dialTimeout time.Duration) DialContextFunc {
	return func(ctx context.Context, _, _ string) (net.Conn, error) {
		d := net.Dialer{
			Timeout:   dialTimeout,
			KeepAlive: 30 * time.Second, // default
		}
		return d.DialContext(ctx, network, addr)
	}
}

func NetworkFromTransType(t TransportType) string {
	switch t {
	case TCPKey:
		return "tcp"
	case UDSKey:
		return "unix"
	default:
		return "tcp" // default fallback
	}
}

func EndpointFromTransType(endpoint string, t TransportType) string {
	switch t {
	case TCPKey:
		return endpoint
	case UDSKey:
		return "http://unix"
	default:
		return endpoint // default fallback
	}

}
