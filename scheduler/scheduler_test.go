package scheduler

import (
	"testing"

	"github.com/barbourkd/docdeliver/document"
)

func TestEnqueue(t *testing.T) {
	scheduler := Scheduler{}
	doc := document.NewDocument("Test Doc", "Test Contents")

	scheduler.Enqueue(doc)

	if scheduler.Length() != 1 {
		t.Errorf("want 1, got %d", scheduler.Length())
	}
}

func TestDequeue(t *testing.T) {
	scheduler := Scheduler{}
	firstDoc := document.NewDocument("First Doc", "Test Contents")
	secondDoc := document.NewDocument("Second Doc", "Test Contents")

	scheduler.Enqueue(firstDoc)
	scheduler.Enqueue(secondDoc)

	got := scheduler.Dequeue()

	assertDocument(t, got, firstDoc)

	if scheduler.Length() != 1 {
		t.Errorf("wrong number of docs in queue")
	}
}

func assertDocument(t *testing.T, got, want document.Document) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
