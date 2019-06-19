package handlers

import (
	"github.com/kataras/iris"
)

// Handleredirection handler
func Handleredirection(app *iris.Application) {
	app.Get("/", func(c iris.Context) {
		c.JSON(iris.Map{
			"result": "ok",
		})
	})
}
