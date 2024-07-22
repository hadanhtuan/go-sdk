package sdk

import (
	"strings"
	"github.com/labstack/echo"
)

type HTTPRequest struct {
	context echo.Context
}

func newHTTPRequest(e echo.Context) *HTTPRequest {
	return &HTTPRequest{
		context: e,
	}
}

func (req *HTTPRequest) GetPath() string {
	return req.context.Path()
}

func (req *HTTPRequest) GetParam(name string) string {
	return req.context.QueryParam(name)
}

func (req *HTTPRequest) GetAllParams() map[string]string {
	var vals = req.context.QueryParams()
	var m = make(map[string]string)
	for key := range vals {
		m[key] = vals.Get(key)
	}
	return m
}

func (req *HTTPRequest) GetHeader(name string) string {
	return req.context.Request().Header.Get(name)
}

func (req *HTTPRequest) GetAllHeaders() map[string]string {
	var vals = req.context.Request().Header
	var m = make(map[string]string)
	for key := range vals {
		m[key] = vals.Get(key)
	}
	return m
}

func (req *HTTPRequest) GetIP() string {
	// for forwarded case
	forwarded := req.GetHeader("X-Forwarded-For")
	if forwarded == "" {
		httpReq := req.context.Request()
		return strings.Split(httpReq.RemoteAddr, ":")[0]
	}

	splitted := strings.Split(forwarded, ",")
	return splitted[0]
}

func (req *HTTPRequest) GetBody(data interface{}) error {
	if err := req.context.Bind(data); err != nil {
		return err
	}
	return nil
}
