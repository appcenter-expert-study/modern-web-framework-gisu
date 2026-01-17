package http

import "github.com/labstack/echo/v4"

/*
	Mutual Contract Between Engine & Dispatcher

	Http Engine(Echo) know only this contract.
		- EchoEngine does not know Dispatcher implementation
		- DIP
*/
type Handler interface {
	// be compatible to Handler type, implement Dispatch Method
	Dispatch(ctx *RequestContext) error
}

/*
	Engine Wrapper
*/
type EchoEngine struct {
	instance *echo.Echo
}

/*
	Constructor
	Echo Init Responsibility
		- middleware
		- logger
		- CORS
*/
func NewEchoEngine() *EchoEngine {
	// new Echo Instance
	e := echo.New()

	// abstract - return EchoEngine struct
	return &EchoEngine{instance: e}
}

/*
	Connect only Once : Echo -> Dispatcher
*/
func (engine *EchoEngine) RegisterDispatcher(dispatcher Handler) {
	/* Register only One Route ("/*")
		disregard HTTP METHOD ( GET POST PUT PATCH ,,)
		allow all path
		Echo does not have Routing Obligation

		So Further, Echo is no more Web Framework. it is just simple HTTP ENGINE
	*/
	engine.instance.Any("/*", func(context echo.Context) error {
		// Echo Context -> RequestContext transform
		// Echo Type Finished From this Layer
		ctx := NewRequestContext(context)
		//delegate to Dispatcher
		return dispatcher.Dispatch(ctx)
	})
}

/*
	Server Execute Responsibility : EchoEngine

	From main() or bootstrap : know just engine.start(":8080")
*/
func (engine *EchoEngine) Start(address string) error {
	return engine.instance.Start(address) 
}
