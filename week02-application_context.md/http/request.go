package http

import "github.com/labstack/echo/v4"

type RequestContext struct {
	Method string
	Path   string
}

func NewRequestContext(c echo.Context) *RequestContext {
	return &RequestContext{
		Method: c.Request().Method,
		Path:   c.Request().URL.Path,
	}
}