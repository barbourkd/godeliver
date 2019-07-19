package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()

	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	// Method: GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to DocDeliver</h1>")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

/* - NON WEB STUFF - TBF
	scheduler := scheduler.Scheduler{}
	document := document.NewDocument("This is a doc!", "This is contents!")

	scheduler.Enqueue(document)

	scheduler.Start()

	time.Sleep(5 * time.Second)

	scheduler.Stop()

	time.Sleep(5 * time.Second)
}

*/
