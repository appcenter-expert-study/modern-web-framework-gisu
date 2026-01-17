// Dispatcher Layer
// Get Request and delegate Execute Responsibility

package dispatcher

import (
	"fmt"
	"week01-dispatcher/http"
)

/*
	type of Dispatcher
		- Empty struct
		- stateless
	For Spring
		- @component class Dispatcher {}
*/
type Dispatcher struct{}


/*
	Generator
	Unique EntryPoint to make Dispatcher Instance
*/
func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

/*
	Static Log Output
	ex) "Dispatcher received {METHOD} {PATH}"

	Dispatcher do not know about
		- Http Implement
		- Echo , net / http / Fiber
	Only Know is abstract Request Information
*/

func (d *Dispatcher) Dispatch(ctx *http.RequestContext) error {
	fmt.Printf(
		"Dispatcher received %s %s \n",
		ctx.Method,
		ctx.Path,
	)
	return nil
}
