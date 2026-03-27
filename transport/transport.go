// Copyright 2025 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
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
			KeepAlive: 30 * time.Second, // Default timeout
		}
		return d.DialContext(ctx, network, addr)
	}
}

func NetworkFromTransType(t Type) string {
	switch t {
	case TCP:
		return "tcp"
	case UDS:
		return "unix"
	default:
		return "tcp" // Default fallback
	}
}

func EndpointFromTransType(endpoint string, t Type) string {
	switch t {
	case TCP:
		return endpoint
	case UDS:
		return "http://unix"
	default:
		return endpoint // Default fallback
	}

}
