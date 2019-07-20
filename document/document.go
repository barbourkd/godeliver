package document

// Document is a basic type for any kind of document
type Document struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// NewDocument creates a new document
func NewDocument(name, content string) Document {
	return Document{Name: name, Content: content}
}
