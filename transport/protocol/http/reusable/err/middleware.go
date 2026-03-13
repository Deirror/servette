// Copyright 2026 Deirror. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package errx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Deirror/servette/encoding/json"
	"github.com/Deirror/servette/logger"
	"github.com/Deirror/servette/transport"
	respx "github.com/Deirror/servette/transport/dtos/resp"
	errx "github.com/Deirror/servette/transport/err"
	httpresp "github.com/Deirror/servette/transport/protocol/http/dtos/resp"
	"github.com/Deirror/servette/transport/protocol/http/htmx"
)

type WriteType int

const (
	JSON  WriteType = iota
	Templ WriteType = 1
)

type Middleware struct {
	log *slog.Logger

	wt WriteType
}

func NewMiddleware(log *slog.Logger, t WriteType) *Middleware {
	return &Middleware{
		log: log,
		wt:  t,
	}
}

func (m *Middleware) ErrMiddleware(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	if err == nil {
		return
	}

	if err.InternalMsg != "" {
		logger.LogFunc(ctx, m.log, r.URL.Path, errors.New(err.InternalMsg))
	}

	switch m.wt {
	case JSON:
		m.WriteErrJSON(ctx, w, err)
	case Templ:
		m.RenderErr(ctx, w, r, err)
	default:
		return
	}
}

func (m *Middleware) WriteErrJSON(ctx context.Context, w http.ResponseWriter, err *errx.Err) {
	if err.MsgKey == transport.JSONFail || err.MsgKey == transport.HeadersWriteFail {
		// Headers are writter when writing json
		return
	}

	status, e := strconv.Atoi(err.Code)
	if e != nil {
		// Generic bad status
		status = http.StatusBadRequest
	}

	resp := respx.New(err.Code, err.MsgKey, nil)
	if e = json.Write(w, status, &resp); e != nil {
		logger.LogFunc(ctx, m.log, "ErrMiddleware", fmt.Errorf("cannot write json: %v", e))
	}
}

func (m *Middleware) RenderErr(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	if htmx.IsHXRequest(r) {
		// Render toast

	} else {
		// Render full page

	}
}

func (m *Middleware) NotFoundMiddleware(w http.ResponseWriter, r *http.Request) {
	switch m.wt {
	case JSON:
		resp := httpresp.New(http.StatusNotFound, transport.URLNotFound, nil)
		if err := json.Write(w, http.StatusNotFound, &resp); err != nil {
			logger.LogFunc(context.Background(), m.log, "NotFoundMiddleware", fmt.Errorf("cannot write json: %v", err))
		}
	case Templ:
		// TODO: Impl.
	default:
		return
	}
}
