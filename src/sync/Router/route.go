package Router

import (
	"github.com/kataras/iris"
	"sync/Handler"
)

func Init() {

	app := iris.New()

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerErrorHandler)
	app.Get("/donate", Handler.DonateHandler, Handler.DonateFinishHandler)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML(`Click <a href="/my500">here</a> to fire the 500 status code`)
	})

	app.Get("/my500", func(ctx iris.Context) {
		ctx.Values().Set("message", "this is the error message")
		ctx.StatusCode(500)
	})
	app.Run(iris.Addr(":8080"))

}

func internalServerErrorHandler(ctx iris.Context) {
	ctx.HTML("Message: <b>" + ctx.Values().GetString("message") + "</b>")

}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("Custom route for 404 not found http code, here you can render a view, html, json <b>any valid response</b>.")
}
