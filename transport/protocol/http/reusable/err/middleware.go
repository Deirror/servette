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
	"github.com/Deirror/servette/transport/dtos/resp"
	"github.com/Deirror/servette/transport/err"
	httpresp "github.com/Deirror/servette/transport/protocol/http/dtos/resp"
	"github.com/Deirror/servette/transport/protocol/http/htmx"
)

type Middleware struct {
	log *slog.Logger
}

func NewMiddleware(log *slog.Logger) *Middleware {
	return &Middleware{
		log: log,
	}
}

func (m *Middleware) ErrMiddleware(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	if err == nil {
		return
	}

	if err.InternalMsg != "" {
		logger.LogFunc(ctx, m.log, r.URL.Path, errors.New(err.InternalMsg))
	}

	if err.MsgKey == transport.JSONFailKey || err.MsgKey == transport.HeadersWriteFailKey {
		// headers are writter when writing json
		return
	}

	if err.MsgKey == transport.TemplFailKey {
		// render generic err indicator
		if htmx.IsHXRequest(r) {

		} else {

		}
	}

	status, e := strconv.Atoi(err.Code)
	if e != nil {
		// generic bad status
		status = http.StatusBadRequest
	}

	resp := respx.New(err.Code, err.MsgKey, nil)
	if e = json.Write(w, status, &resp); e != nil {
		logger.LogFunc(ctx, m.log, "ErrMiddleware", fmt.Errorf("cannot write json: %v", e))
	}
}

func (m *Middleware) NotFoundMiddleware(w http.ResponseWriter, r *http.Request) {
	resp := httpresp.New(http.StatusNotFound, transport.URLNotFoundKey, nil)
	if err := json.Write(w, http.StatusNotFound, &resp); err != nil {
		logger.LogFunc(context.Background(), m.log, "NotFoundMiddleware", fmt.Errorf("cannot write json: %v", err))
	}
}
