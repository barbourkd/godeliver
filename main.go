package main

import (
	"fmt"
	"time"

	"github.com/barbourkd/docdeliver/config"
	"github.com/barbourkd/docdeliver/document"
	"github.com/barbourkd/docdeliver/scheduler"
)

const configFile = "resources/sample.cfg"

func main() {
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
