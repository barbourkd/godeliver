package main

import (
	"fmt"
	"time"

	"github.com/barbourkd/docdeliver/document"
	"github.com/barbourkd/docdeliver/scheduler"
)

func main() {
	scheduler := scheduler.Scheduler{}
	document := document.NewDocument("This is a doc!", "This is contents!")

	scheduler.Enqueue(document)

	for scheduler.Length() > 0 {
		fmt.Printf("%d document(s) are currently scheduled\n", scheduler.Length())
		time.Sleep(2 * time.Second)
	}
}
