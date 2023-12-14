package domain

import (
	"net/http"

	"github.com/go-chi/render"
)

type Error struct {
	Code       string `json:"code"`
	HttpStatus int    `json:"status"`
}

func NewError(err error) *Error {
	status := http.StatusInternalServerError
	switch err.Error() {
	case ErrCodeNotFound:
		status = http.StatusNotFound
	case ErrCodeValidation:
		status = http.StatusBadRequest
	case ErrCodeDuplicate:
		status = http.StatusConflict
	case ErrCodeParsing:
		status = http.StatusBadRequest
	}

	return &Error{
		Code:       err.Error(),
		HttpStatus: status,
	}
}

func (e *Error) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatus)
	return nil
}

func (e *Error) Error() string {
	return string(e.Code)
}

const (
	ErrCodeNotFound   string = "err_not_found"
	ErrCodeValidation string = "err_validation"
	ErrCodeDuplicate  string = "err_duplicate"
	ErrCodeParsing    string = "err_parsing"
)
