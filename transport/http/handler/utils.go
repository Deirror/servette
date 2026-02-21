package handler

import (
    "net/http"
    "strings"
)

// SafeRedirect redirects to the Referer header if valid, otherwise to a default path.
func SafeRedirect(w http.ResponseWriter, r *http.Request, defaultPath string) {
    referer := r.Header.Get("Referer")
    if referer == "" || !strings.HasPrefix(referer, "/") {
        referer = defaultPath
    }

    http.Redirect(w, r, referer, http.StatusSeeOther)
}
