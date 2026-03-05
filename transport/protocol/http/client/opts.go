package client

import (
	"net/http"
)

type RequestOption func(req *http.Request)

func WithCookies(cookies ...*http.Cookie) RequestOption {
	return func(req *http.Request) {
		for _, c := range cookies {
			req.AddCookie(c)
		}
	}
}

func WithHeader(key, value string) RequestOption {
	return func(req *http.Request) {
		req.Header.Set(key, value)
	}
}
