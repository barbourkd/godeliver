package main

import (
	"time"

	"github.com/barbourkd/docdeliver/document"
	"github.com/barbourkd/docdeliver/scheduler"
)

func main() {
	scheduler := scheduler.Scheduler{}
	document := document.NewDocument("This is a doc!", "This is contents!")

	scheduler.Enqueue(document)

	scheduler.Start()

	time.Sleep(5 * time.Second)

	scheduler.Stop()

	time.Sleep(5 * time.Second)
}
