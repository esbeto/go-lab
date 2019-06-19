package server

import (
	"os"

	"github.com/kataras/iris"
)

// SetUp will initialize the iris application
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	return app
}

// RunServer will run the iris application
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
	)
}
