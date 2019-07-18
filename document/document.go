package document

// Document is a basic type for any kind of document
type Document struct {
	name    string
	content string
}

// Name returns the name of the document
func (d Document) Name() string {
	return d.name
}

// Content returns the content of the document
func (d Document) Content() string {
	return d.content
}

// NewDocument creates a new document
func NewDocument(name, content string) Document {
	return Document{name: name, content: content}
}
