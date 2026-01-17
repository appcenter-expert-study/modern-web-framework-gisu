package bootstrap

import (
	"week01-dispatcher/dispatcher"
	"week01-dispatcher/http"
)

/*
	only do assemble in bootstrap
	do not include Request Processing Logic
*/
func Run() {
	println("bootstrap running")
	//app start Logic

	//Create HTTP Engine : abstracted EchoEngine
	engine := http.NewEchoEngine()

	//Dispatcher
	dispatcher := dispatcher.NewDispatcher()

	//Register Dispatcher to Engine
	engine.RegisterDispatcher(dispatcher)

	//Run Server
	engine.Start(":8080")
}
