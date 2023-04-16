package http_error

import "errors"

var (
	ErrActivityNotFound = errors.New("activity not found")
	ErrRecordNotfound   = errors.New("record not found")
	ErrIsActiveNull     = "Is_active cannot be null"
)

type Form struct {
	Field   string
	Message string
}
