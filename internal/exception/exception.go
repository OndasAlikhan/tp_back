package exception

import (
	"errors"
	"net/http"
)

var (
	ErrEmptyRequest  = errors.New("body is empty")
	ErrDecodeRequest = errors.New("body is incorrect format")
	ErrAlreadyExists = errors.New("recording already exists")
	ErrInvalidCred   = errors.New("invalid email or password")
	ErrNotFound      = errors.New("recording not found")
)

func Code(err error) int {
	switch {
	case errors.Is(err, ErrEmptyRequest), errors.Is(err, ErrDecodeRequest), errors.Is(err, ErrAlreadyExists), errors.Is(err, ErrInvalidCred):
		return http.StatusBadRequest
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
