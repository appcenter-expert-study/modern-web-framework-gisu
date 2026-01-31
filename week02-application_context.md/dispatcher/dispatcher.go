package dispatcher

import (
	"fmt"
	"week02-application_context/http"
)

type Dispatcher struct{}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(ctx *http.RequestContext) error {
	fmt.Printf(
		"Dispatcher received %s %s \n",
		ctx.Method,
		ctx.Path,
	)
	return nil
}