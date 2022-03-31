package main

import (
	"github.com/Capucinoxx/vibrance/server/service/event"
	"github.com/Capucinoxx/vibrance/server/utils/router"
	"github.com/Capucinoxx/vibrance/server/utils/server"
)

func handleRoutes() {
	r := router.Router()
	r.AddRoutes(event.Handle()...)

	r.AddMiddlewares(router.Logger)

	r.Consumer("")
}

func main() {
	handleRoutes()

	server.Start()
}
