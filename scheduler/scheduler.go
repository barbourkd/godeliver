package scheduler

import (
	"errors"

	"github.com/barbourkd/docdeliver/document"
)

// Scheduler holds a bunch of documents
type Scheduler struct {
	queue []document.Document
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

// Length returns the length of the scheduler's queue
func (s *Scheduler) Length() int {
	return len(s.queue)
}
