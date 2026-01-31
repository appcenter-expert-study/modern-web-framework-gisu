package http

import "github.com/labstack/echo/v4"

type Handler interface {
	Dispatch(ctx *RequestContext) error
}

type EchoEngine struct {
	instance *echo.Echo
}

func NewEchoEngine() *EchoEngine {
	e := echo.New()
	return &EchoEngine{instance: e}
}

func (engine *EchoEngine) RegisterDispatcher(dispatcher Handler) {
	engine.instance.Any("/*", func(context echo.Context) error {
		ctx := NewRequestContext(context)
		return dispatcher.Dispatch(ctx)
	})
}

func (engine *EchoEngine) Start(address string) error {
	return engine.instance.Start(address)
}