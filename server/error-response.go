package server

import (
	"encoding/json"
	"fmt"
)

type ResponseError struct {
	Error      string `json:"error"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"status_code"`
}

func (e ResponseError) Marshal() []byte {
	responseErrorData, responseErrorMarshalError := json.Marshal(e)
	if responseErrorMarshalError != nil {
		return []byte(fmt.Sprintf("{\"error\": \"%v\", \"reason\": \"%v\", \"status_code\": %v}", "Problem to marshal response error struct", responseErrorMarshalError.Error(), 500))
	}

	return responseErrorData
}
