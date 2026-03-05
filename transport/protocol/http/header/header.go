package header

import (
	"net/http"
)

func Set(h *http.Header, key string, val string) {
	h.Set(key, val)
}

func SetCacheControl(h *http.Header, val string) {
	Set(h, CacheControlKey, val)
}
