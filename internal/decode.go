package internal

import (
	"encoding/json"
	"matryer/pkg/joker"
	"net/http"
)

func Decode(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return joker.WrapErrorf(err, joker.CodeInvalidArgument, "invalid body")
	}
	return nil
}
