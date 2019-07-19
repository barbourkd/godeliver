package scheduler

import (
	"testing"
	"time"

	"github.com/barbourkd/docdeliver/devices"
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

func TestDevices(t *testing.T) {
	t.Run("add return printer", func(t *testing.T) {
		scheduler := Scheduler{}
		printer := devices.ReturnPrinter{}
		scheduler.AddDevice(printer)

		if len(scheduler.devices) != 1 {
			t.Error("expected one device on scheduler")
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

func TestRun(t *testing.T) {
	t.Run("start", func(t *testing.T) {
		scheduler := Scheduler{}
		scheduler.Start()

		if !scheduler.Running() {
			t.Error("Started scheduler but it isn't running!")
		}

	})

	t.Run("stop", func(t *testing.T) {
		scheduler := Scheduler{}
		scheduler.Start()
		scheduler.Stop()

		if scheduler.Running() {
			t.Error("Stopped scheduler but it is still running")
		}
	})
}

func TestBasicScheduling(t *testing.T) {
	t.Run("print queued doc to spy", func(t *testing.T) {
		timeout := time.After(2 * time.Second)
		done := make(chan bool)

		scheduler := Scheduler{}
		docName := "Test Doc"
		docContent := "This should print!"
		doc := document.NewDocument(docName, docContent)
		device := SpyDevice{}

		// Is passing this as a pointer the best way to do this?
		scheduler.AddDevice(&device)
		scheduler.Enqueue(doc)
		scheduler.Start()

		go func() {
			for scheduler.Length() > 0 {
			}
			done <- true
			scheduler.Stop()
		}()

		select {
		case <-timeout:
			t.Fatal("test timed out")
		case <-done:
			got := device.output
			if got != docContent {
				t.Errorf("got %q want %q", got, docContent)
			}
		}

	})
}

func assertDocument(t *testing.T, got, want document.Document) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

type SpyDevice struct {
	output string
}

func (s *SpyDevice) Print(doc document.Document) string {
	s.output = doc.Content
	return ""
}
