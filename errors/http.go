package errors

import (
	"strings"
)

// HTTPError is a trivial implementation of error.
type HTTPError struct {
	Name    string      `json:"name"`
	Status  int         `json:"status"`
	Message interface{} `json:"stack,omitempty"`
}

func (e *HTTPError) Error() string {
	if message, ok := e.Message.(string); ok {
		return message
	}
	return "-"
}

// Detail Explain
func (e *HTTPError) Detail() interface{} {
	if detail := strings.Split(e.Message.(string), "\n"); len(detail) > 1 {
		return detail
	}

	return e.Message
}
