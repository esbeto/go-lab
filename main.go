package main

import (
	handlers "github.com/esbeto/go-lab/api/handlers"
	server "github.com/esbeto/go-lab/api/server"
	utils "github.com/esbeto/go-lab/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.Handleredirection(app)
	server.RunServer(app)
}
