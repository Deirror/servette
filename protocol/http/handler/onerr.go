package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/Deirror/servette/encoding/json"
	"github.com/Deirror/servette/transport/err"
)

// Can be used in Wrap func as default one
func DefaultErrHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	accept := r.Header.Get("Accept")
	if strings.Contains(accept, "text/html") {
		HTMLErrHandler(ctx, w, r, err)
	} else {
		JSONErrHandler(ctx, w, r, err)
	}
}

// Default func for json error handling.
func JSONErrHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	if err == nil {
		err = errx.New(strconv.Itoa(http.StatusInternalServerError), "internal server error", "err is nil")
	}
	status, _ := strconv.Atoi(err.Code)
	json.Write(w, status, err)
}

// Default func for html error handling, with examplary html code.
func HTMLErrHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err *errx.Err) {
	if err == nil {
		err = errx.New(strconv.Itoa(http.StatusInternalServerError), "internal server error", "err is nil")
	}

	status, _ := strconv.Atoi(err.Code)
	w.WriteHeader(status)
	_, _ = w.Write([]byte(
		"<html><head><title>Error</title></head><body>" +
			"<h1>Error</h1>" +
			"<p>Status: " + http.StatusText(status) + "</p>" +
			"<p>MessageKey: " + err.MsgKey + "</p>" +
			"</body></html>",
	))
}
