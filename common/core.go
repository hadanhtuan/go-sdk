package common

// Error custom error of sdk
type Error struct {
	Type    string
	Message string
	Data    interface{}
}

func (e *Error) Error() string {
	return e.Type + " : " + e.Message
}

// APIResponse This is  response object with JSON format
type APIResponse struct {
	Status    int32            `json:"status"`
	Data      interface{}       `json:"data,omitempty"`
	Message   string            `json:"message"`
	ErrorCode string            `json:"errorCode,omitempty"`
	Total     int64             `json:"total,omitempty"`
	Headers   map[string]string `json:"headers,omitempty"`
}

