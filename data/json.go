package data

import (
	"encoding/json"
	"io"
)

func UnmarshalJson[T any](r io.Reader) (T, error) {
	var t T
	content, err := io.ReadAll(r)
	if err != nil {
		return t, err
	}
	err = json.Unmarshal(content, &t)
	return t, err
}
