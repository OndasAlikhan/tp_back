package rq

import (
	"errors"
	"io"

	"tp_back/internal/exception"

	"github.com/go-chi/render"
)

func DecodeJSON(body io.ReadCloser, request interface{}) error {
	err := render.DecodeJSON(body, &request)
	if errors.Is(err, io.EOF) {
		return exception.ErrEmptyRequest
	}

	if err != nil {
		return exception.ErrDecodeRequest
	}

	return nil
}
