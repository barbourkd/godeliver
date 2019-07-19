package devices

import "github.com/barbourkd/docdeliver/document"

// ReturnPrinter is a super basic printer that simply returns the doc contents
type ReturnPrinter struct {
	name string
}

// Print simply returns the doc contents for ReturnPrinter
func (p ReturnPrinter) Print(doc document.Document) string {
	return doc.Content()
}
