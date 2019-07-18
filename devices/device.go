package devices

import "github.com/barbourkd/docdeliver/document"

// Device is a generic interface for all devices
type Device interface {
	Print(document.Document) string
}
