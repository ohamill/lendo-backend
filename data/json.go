package data

import (
	"encoding/json"
	"io"
)

// DecodeJson will convert r's JSON-encoded content into a value of type T
func DecodeJson[T any](r io.Reader) (T, error) {
	var t T
	err := json.NewDecoder(r).Decode(&t)
	return t, err
}
