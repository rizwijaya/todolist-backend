package http_error

import "errors"

var (
	ErrActivityNotFound = errors.New("Activity not found")
	ErrRecordNotfound   = errors.New("record not found")
)

type Form struct {
	Field   string
	Message string
}
