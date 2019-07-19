package scheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/barbourkd/docdeliver/devices"
	"github.com/barbourkd/docdeliver/document"
)

// Scheduler holds a bunch of documents
type Scheduler struct {
	// TODO: extract DocumentQueue to its own type
	queue   []document.Document
	devices []devices.Device
	running bool
}

// Start starts the main loop of the scheduler
func (s *Scheduler) Start() {
	s.running = true

	go func() {
		for s.running {
			//fmt.Printf("Scheduler Running:\n%d docs in queue\n\n", s.Length())

			// We have docs to schedule!!!
			if len(s.queue) > 0 && len(s.devices) > 0 {
				doc, _ := s.Dequeue()
				s.devices[0].Print(doc)
			}
			time.Sleep(10 * time.Millisecond)
		}

		fmt.Print("Scheduler shutting down\n")
	}()
}

// Stop stops the main loop of the scheduler
func (s *Scheduler) Stop() {
	s.running = false
}

// Running tells us if the scheduler is currently running
func (s *Scheduler) Running() bool {
	return s.running
}

// Enqueue adds a document to the scheduler
func (s *Scheduler) Enqueue(doc document.Document) {
	s.queue = append(s.queue, doc)
}

// Dequeue pops the first doc off the queue
func (s *Scheduler) Dequeue() (document.Document, error) {
	var doc document.Document

	if s.Length() < 1 {
		return doc, errors.New("attempting to dequeue from empty queue")
	}

	doc = s.queue[0]
	s.queue = s.queue[1:]

	return doc, nil
}

// Peek shows the first doc without dequeuing
func (s *Scheduler) Peek() document.Document {
	var doc document.Document

	if s.Length() > 0 {
		doc = s.queue[0]
	}
	return doc
}

// AddDevice adds a device to the scheduler
func (s *Scheduler) AddDevice(device devices.Device) {
	s.devices = append(s.devices, device)
}

// Length returns the length of the scheduler's queue
func (s *Scheduler) Length() int {
	return len(s.queue)
}

// Devices returns all devices in the scheduler
func (s *Scheduler) Devices() []devices.Device {
	return s.devices
}
