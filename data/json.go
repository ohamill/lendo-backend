package data

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

// DecodeJson will convert r's JSON-encoded content into a value of type T
func DecodeJson[T any](r io.Reader) (T, error) {
	var t T
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	err := d.Decode(&t)
	if err != nil {
		return t, err
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(t)
	return t, err
}
