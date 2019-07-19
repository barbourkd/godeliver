package main

import (
	"fmt"

	"github.com/barbourkd/docdeliver/config"
	"github.com/barbourkd/docdeliver/devices"
	"github.com/barbourkd/docdeliver/document"
	"github.com/barbourkd/docdeliver/scheduler"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const configFile = "resources/sample.cfg"

func main() {
	app := iris.Default()

	config, err := config.ParseConfig(configFile)

	if err != nil {
		fmt.Print("Error reading config, aborting\n")
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Printf("Running with config %q\n", config.Title)

	deviceMap, err := config.GenerateDevices()

	if err != nil {
		fmt.Printf("Invalid configuration")
		fmt.Printf("Error: %q\n", err)
		return
	}

	scheduler := scheduler.Scheduler{}

	for d := range deviceMap {
		scheduler.AddDevice(deviceMap[d])
	}

	doc := document.NewDocument("This is a doc!", "This is contents!")
	scheduler.Enqueue(doc)

	scheduler.Start()

	oneDevice := devices.NewReturnPrinter("this is")
	oneDevice.Status = "Active"
	twoDevice := devices.NewReturnPrinter("from the back-end!")
	twoDevice.Status = "Inactive"

	app.Get("/devices", func(ctx iris.Context) {
		results := []devices.Device{}
		results = append(results, oneDevice)
		results = append(results, twoDevice)

		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON(results)
	})

	app.Post("/documents", func(ctx iris.Context) {
		param := &document.Document{}
		err := ctx.ReadJSON(param)
		if err != nil {
			ctx.JSON(context.Map{"response": err.Error()})
		} else {
			ctx.Writef("Yeah we did it!")
			scheduler.Enqueue(*param)
		}
	})

	//	app.Post("/somePost", posting)

	app.Run(iris.Addr(":8081"))
}

/*
	config, err := config.ParseConfig(configFile)

	if err != nil {
		fmt.Print("Error reading config, aborting\n")
		fmt.Printf("Error: %q\n", err)
		return
	}

	fmt.Printf("Running with config %q\n", config.Title)

	deviceMap, err := config.GenerateDevices()

	if err != nil {
		fmt.Printf("Invalid configuration")
		fmt.Printf("Error: %q\n", err)
		return
	}

	scheduler := scheduler.Scheduler{}

	for d := range deviceMap {
		scheduler.AddDevice(deviceMap[d])
	}

	document := document.NewDocument("This is a doc!", "This is contents!")
	scheduler.Enqueue(document)

	scheduler.Start()

	time.Sleep(5 * time.Second)

	scheduler.Stop()

	time.Sleep(5 * time.Second)
}
*/
