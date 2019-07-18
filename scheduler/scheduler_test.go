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
	t.Run("multiple documents in queue", func(t *testing.T) {
		scheduler := Scheduler{}
		firstDoc := document.NewDocument("First Doc", "Test Contents")
		secondDoc := document.NewDocument("Second Doc", "Test Contents")

		scheduler.Enqueue(firstDoc)
		scheduler.Enqueue(secondDoc)

		got, _ := scheduler.Dequeue()

		assertDocument(t, got, firstDoc)

		if scheduler.Length() != 1 {
			t.Errorf("wrong number of docs in queue")
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		scheduler := Scheduler{}

		_, err := scheduler.Dequeue()

		if err == nil {
			t.Errorf("expected an error dequeuing an empty queue")
		}
	})
}

func TestPeek(t *testing.T) {
	t.Run("single document in queue", func(t *testing.T) {
		scheduler := Scheduler{}
		doc := document.NewDocument("First Doc", "Test Contents")

		scheduler.Enqueue(doc)

		got := scheduler.Peek()
		assertDocument(t, got, doc)

		if scheduler.Length() != 1 {
			t.Error("wrong number of docs in queue")
		}
	})

	t.Run("empty queue", func(t *testing.T) {
		scheduler := Scheduler{}

		// for now just don't want to error out, no big deal
		scheduler.Peek()
	})
}

func assertDocument(t *testing.T, got, want document.Document) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
