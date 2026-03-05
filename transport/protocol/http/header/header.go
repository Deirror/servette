package header

import (
	"fmt"
	"net/http"
	"strconv"
)

func Set(w http.ResponseWriter, key string, val string) {
	w.Header().Set(key, val)
}

func SetCacheControl(w http.ResponseWriter, privacy string, age int) {
	Set(w, CacheControlKey, fmt.Sprintf("%s, max-age=%s", privacy, strconv.Itoa(age)))
}
