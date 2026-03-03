package handler

import (
	"net/http"
	"strings"
)

// SafeRedirect redirects to the Referer header if valid, otherwise to a default path.
func SafeRedirect(w http.ResponseWriter, r *http.Request, defaultPath ...string) {
	// General default path, if not other default is passed.
	path := "/"
	if len(defaultPath) > 0 {
		path = defaultPath[0]
	}
	referer := r.Header.Get("Referer")
	if referer == "" || !strings.HasPrefix(referer, "/") {
		referer = path
	}

	http.Redirect(w, r, referer, http.StatusSeeOther)
}
