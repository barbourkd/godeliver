package devices

import (
	"fmt"

	"github.com/barbourkd/docdeliver/document"
)

// ReturnPrinter is a super basic printer that simply returns the doc contents
type ReturnPrinter struct {
	name string
}

// Print simply returns the doc contents for ReturnPrinter
func (p ReturnPrinter) Print(doc document.Document) string {
	fmt.Printf("\nPrinting %q\n%q\n\n", doc.Name(), doc.Content())
	return doc.Content()
}

// NewReturnPrinter returns a new return printer with the given name
func NewReturnPrinter(name string) ReturnPrinter {
	return ReturnPrinter{name: name}
}
