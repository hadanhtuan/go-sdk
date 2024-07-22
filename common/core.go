package common

import (
	"github.com/golang-jwt/jwt/v5"
)

var BODY_PAYLOAD = "BODY_PAYLOAD"
var JWT_PAYLOAD = "JWT_PAYLOAD"

// APIResponse This is  response object with JSON format
type APIResponse struct {
	Status  int32             `json:"status"`
	Data    interface{}       `json:"data,omitempty"`
	Message string            `json:"message"`
	Total   int64             `json:"total,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
}

type JWTPayload struct {
	UserID   string `json:"userId"`
	DeviceID string `json:"deviceId"`
	jwt.RegisteredClaims
}

type JWTToken struct {
	AccessToken      string `json:"accessToken,omitempty"`
	AccessExpiresAt  int64  `json:"accessExpiresAt,omitempty"`
	RefreshToken     string `json:"refreshToken,omitempty"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt,omitempty"`
}

// StatusEnum ...
type StatusEnt struct {
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
var APIStatus = &StatusEnt{
	Ok:           200,
	Created:      201,
	BadRequest:   400,
	Unauthorized: 401,
	Forbidden:    403,
	NotFound:     404,
	Timeout:      408,
	ServerError:  500,
}

type HTTPMethodValue string

type HTTPMethodEnt struct {
	GET     HTTPMethodValue
	POST    HTTPMethodValue
	PUT     HTTPMethodValue
	DELETE  HTTPMethodValue
	OPTIONS HTTPMethodValue
}

// APIStatus Published enum
var HTTPMethod = &HTTPMethodEnt{
	GET:     "GET",
	POST:    "POST",
	PUT:     "PUT",
	DELETE:  "DELETE",
	OPTIONS: "OPTIONS",
}
