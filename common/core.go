package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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

type JWTPayload struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"username"`
	DeviceID string    `json:"deviceId"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}

type JWTToken struct {
	Token     string    `json:"token,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"`
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
