package server

import (
	"fmt"
)

type ResponseError struct {
	Error      string
	Reason     string
	StatusCode int
}

func (e ResponseError) Marshal() []byte {
	return []byte(
		fmt.Sprintf(
			"{\"error\": \"%v\", \"reason\": \"%v\", \"status_code\": %v}",
			e.Error,
			e.Reason,
			e.StatusCode,
		))
}
