package main

import (
	"fmt"

	"github.com/barbourkd/docdeliver/config"
	"github.com/barbourkd/docdeliver/document"
	"github.com/barbourkd/docdeliver/scheduler"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const configFile = "resources/sample.cfg"

func main() {
	// Read config
	config, err := config.ParseConfig(configFile)

	if err != nil {
		fmt.Print("Error reading config, aborting\n")
		fmt.Printf("Error: %q\n", err)
		return
	}

	// Create devices from config - move to web app soon
	deviceMap, err := config.GenerateDevices()

	if err != nil {
		fmt.Printf("Invalid configuration")
		fmt.Printf("Error: %q\n", err)
		return
	}

	// start scheduler and add devices
	scheduler := scheduler.Scheduler{}
	for d := range deviceMap {
		scheduler.AddDevice(deviceMap[d])
	}
	scheduler.Start()

	// Web app config and start
	app := iris.Default()

	app.Get("/api/devices", func(ctx iris.Context) {
		results := scheduler.Devices()

		ctx.JSON(results)
	})

	app.Post("/api/documents", func(ctx iris.Context) {
		param := &document.Document{}
		err := ctx.ReadJSON(param)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			ctx.Writef("Yeah we did it!")
			scheduler.Enqueue(*param)
		}
	})

	app.Run(iris.Addr(":8081"))

	scheduler.Stop()
}
