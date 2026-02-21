package htmx

import (
	"net/http"
)

func IsHXRequest(r *http.Request) bool {
	return r.Header.Get(HXRequestKey) == "true"
}
