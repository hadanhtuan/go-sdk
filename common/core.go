package common

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var BODY_PAYLOAD = "BODY_PAYLOAD"

type BaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Total     int64  `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
}

// APIResponse This is  response object with JSON format
type APIResponse struct {
	Status    int32             `json:"status"`
	Data      interface{}       `json:"data,omitempty"`
	Message   string            `json:"message"`
	Total     int64             `json:"total,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
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
}