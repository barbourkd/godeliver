package devices

import (
	"testing"

	"github.com/barbourkd/docdeliver/document"
)

func TestReturnPrinter(t *testing.T) {
	printer := ReturnPrinter{name: "Test Printer"}
	doc := document.NewDocument("Name", "Content")

	got := printer.Print(doc)
	want := doc.Content()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
