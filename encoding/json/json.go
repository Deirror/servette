package json

import (
	"encoding/json"
	"io"
	"net/http"
)

// A wrapper for json data encoding.
func Encode(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}

// A func that can be used in handlers to write JSON.
// Includes status code and structured data to marshall.
func Write(w http.ResponseWriter, s int, v any) error {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(s)
	return Encode(w, v)
}

// Can be called when already having a var of type T.
// Unmarshals data into the target.
func DecodeInto[T any](r io.Reader, target *T) error {
	return json.NewDecoder(r).Decode(target)
}

// Same as func DecodeInto, but initziliazes new var of T.
func Decode[T any](r io.Reader) (T, error) {
	var v T
	err := json.NewDecoder(r).Decode(&v)
	return v, err
}

// Marshal returns JSON bytes from a Go value
func Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal reads JSON bytes into a target
func Unmarshal[T any](data []byte) (T, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return v, err
}

// BodyToJSON reads an HTTP response body into a typed struct
func BodyToJSON[T any](resp *http.Response) (T, error) {
	defer resp.Body.Close()
	return Decode[T](resp.Body)
}
