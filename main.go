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

	scheduler := scheduler.Scheduler{}
	document := document.NewDocument("This is a doc!", "This is contents!")

	scheduler.Enqueue(document)

	scheduler.Start()

	time.Sleep(5 * time.Second)

	scheduler.Stop()

	time.Sleep(5 * time.Second)
}
