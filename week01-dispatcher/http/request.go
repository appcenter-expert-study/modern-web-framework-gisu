package http

/*
	Direct Dependency for Echo Framework
	Echo Dependency is only allowed in HTTP Adapter Layer
*/
import "github.com/labstack/echo/v4"

/*
	Request Model in Our Framework
	deliver minimal request data (Method, Path)
*/
type RequestContext struct {
	Method string
	Path string
}

/*
	Factory Method
	Echo -> Our Framework Type (Adapter) : ACL (Anti-Corruption Layer)

	do not expose Echo Context.
	Minimal data included (Method, Path)
*/
func NewRequestContext(c echo.Context) *RequestContext {
	return &RequestContext{
		Method: c.Request().Method,
		Path:	c.Request().URL.Path,
	}
}
