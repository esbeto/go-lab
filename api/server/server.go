package server

import (
	"os"

	"github.com/kataras/iris"
)

// SetUp server
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

// RunServer should start server
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
