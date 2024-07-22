package sdk

import (
	"fmt"
	"sync"
	"github.com/hadanhtuan/go-sdk/common"
	"github.com/labstack/echo"
)

type Handler = func(req *HTTPRequest, res *HTTPResponse) error

type HTTPServer struct {
	Echo *echo.Echo
	Port string
	Host string
}

func (app *App) NewHTTPServer(host, port string) *HTTPServer {
	s := &HTTPServer{}
	s.Host = host
	s.Port = port
	s.Echo = echo.New()

	app.HTTPServerList = append(app.HTTPServerList, s)
	return s
}

func (s *HTTPServer) Start(wg *sync.WaitGroup) {
	url := fmt.Sprintf(
		"%s:%s",
		s.Host,
		s.Port,
	)
	err := s.Echo.Start(url)
	if err != nil {
		fmt.Println("Fail to start " + err.Error())
	}
	wg.Done()
}

func (s *HTTPServer) AddHandler(method common.HTTPMethodValue, route string, handler Handler) {
	fn := func(c echo.Context) error {
		// Pass context between HTTPRequest and HTTPResponse
		handler(newHTTPRequest(c), newHTTPResponse(c))
		return nil
	}

	switch method {
	case common.HTTPMethod.GET:
		s.Echo.GET(route, fn)
	case common.HTTPMethod.POST:
		s.Echo.POST(route, fn)
	case common.HTTPMethod.PUT:
		s.Echo.PUT(route, fn)
	case common.HTTPMethod.DELETE:
		s.Echo.DELETE(route, fn)
	}
}
