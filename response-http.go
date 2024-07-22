package sdk

import (
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/labstack/echo"
)

type HTTPResponse struct {
	context echo.Context
}

func newHTTPResponse(e echo.Context) *HTTPResponse {
	return &HTTPResponse{
		context: e,
	}
}

func (resp *HTTPResponse) Respond(response *common.APIResponse) error {
	var context = resp.context

	// can implement more code for response in here

	if response.Headers != nil {
		header := context.Response().Header()
		for key, value := range response.Headers {
			header.Set(key, value)
		}
		response.Headers = nil
	}

	return context.JSON(int(response.Status), response)
}
