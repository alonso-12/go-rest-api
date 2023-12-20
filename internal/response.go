package internal

import (
	"encoding/json"
	"errors"
	"matryer/pkg/joker"
	"net/http"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type response struct {
	Error     string      `json:"error"`
	Status    string      `json:"status"`
	HTTPCode  int         `json:"http_code"`
	Datetime  string      `json:"datetime"`
	Timestamp int64       `json:"timestamp"`
	Details   interface{} `json:"details,omitempty"`
}

func RenderResponse(w http.ResponseWriter, err error) {
	var (
		verr validation.Errors
		werr joker.Error
		resp = response{
			Status:    "fail",
			Datetime:  time.Now().Format("2006-01-02 15:04:05"),
			Timestamp: time.Now().Unix(),
		}
	)

	if !errors.As(err, &werr) {
		resp.HTTPCode = http.StatusInternalServerError
		resp.Error = "internal server error"
	} else {
		switch werr.Code() {
		case joker.CodeInvalidArgument:
			resp.HTTPCode = http.StatusBadRequest
			if errors.As(werr, &verr) {
				resp.Details = verr
			}
		case joker.CodeNoContent:
			resp.HTTPCode = http.StatusNotFound
		case joker.CodeNotFound:
			resp.HTTPCode = http.StatusNotFound
		}
		resp.Error = werr.Message()
	}
	Respond(w, resp, resp.HTTPCode)
}

func Respond(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
