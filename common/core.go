package common

import (
	"encoding/json"
	"github.com/hadanhtuan/go-sdk/common/proto"
)

var BODY_PAYLOAD = "BODY_PAYLOAD"


// APIResponse This is  response object with JSON format
type APIResponse struct {
	Status  int32             `json:"status"`
	Data    interface{}       `json:"data,omitempty"`
	Message string            `json:"message"`
	Total   int64             `json:"total,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

// StatusEnum ...
type StatusEnum struct {
	Ok           int32
	Created      int32
	BadRequest   int32
	Unauthorized int32
	Forbidden    int32
	NotFound     int32
	Timeout      int32
	ServerError  int32
}

// APIStatus Published enum
var APIStatus = &StatusEnum{
	Ok:           200,
	Created:      201,
	BadRequest:   400,
	Unauthorized: 401,
	Forbidden:    403,
	NotFound:     404,
	Timeout:      408,
	ServerError:  500,
}

func ConvertResult(payload *sdkProto.BaseResponse) (result *APIResponse) {
	var data interface{}

	if payload == nil {
		result.Message = "Internal Server Error"
		result.Status = APIStatus.ServerError
		return
	}
	err := json.Unmarshal([]byte(payload.Data), &data)

	if err != nil {
		result.Message = "Error marshall payload data. Error detail: " + err.Error()
		result.Status = APIStatus.ServerError
		return
	}

	result.Status = payload.Status
	result.Message = payload.Message
	result.Total = payload.Total
	result.Data = data
	return
}
