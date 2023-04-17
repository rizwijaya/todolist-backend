package http_error

import "errors"

var (
	ErrActivityNotFound = errors.New("activity not found")
	ErrRecordNotfound   = errors.New("record not found")
	ErrIsActiveNull     = "is_active cannot be null"
	ErrPriorityNull     = "priority cannot be null"
	ErrTitleRequired    = "title is required"
)

type Form struct {
	Field   string
	Message string
}
