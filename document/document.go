package document

// Document is a basic type for any kind of document
type Document struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

/*
// Name returns the name of the document
func (d Document) Name() string {
	return d.Name
}

// Content returns the content of the document
func (d Document) Content() string { `json:name`
	return d.Content
}
*/

// NewDocument creates a new document
func NewDocument(name, content string) Document {
	return Document{Name: name, Content: content}
}
