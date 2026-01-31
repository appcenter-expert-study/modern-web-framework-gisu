package bootstrap

import (
	"reflect"
	"week02-application_context/context"
	"week02-application_context/dispatcher"
	"week02-application_context/http"
)

func Run(){
	engine := http.NewEchoEngine()

	// Create Application Context (IoC)
	ctx := context.NewApplicationContext()

	// Register Dispatcher on Application Context
	ctx.RegisterBean(
		//reflect - Type -> Factory
		reflect.TypeFor[*dispatcher.Dispatcher](),
		func(ctx *context.ApplicationContext) any{
			return dispatcher.NewDispatcher()
		},
	)

	//Get Dispatcher Type form Application Context
	dispatcher := ctx.GetBean(
		reflect.TypeFor[*dispatcher.Dispatcher](),
	).(*dispatcher.Dispatcher)

	//Register Dispatcher Bean on Engine
	engine.RegisterDispatcher(dispatcher)
	engine.Start(":18080")
}