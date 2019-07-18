package scheduler

import (
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
func (s *Scheduler) Dequeue() document.Document {
	doc := s.queue[0]
	s.queue = s.queue[1:]

	return doc
}

// Length returns the length of the scheduler's queue
func (s *Scheduler) Length() int {
	return len(s.queue)
}
