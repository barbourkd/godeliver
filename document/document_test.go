package document

import "testing"

func TestNewDocument(t *testing.T) {
	t.Run("new doc", func(t *testing.T) {
		name := "Name"
		content := "Content"
		doc := NewDocument(name, content)

		assertString(t, doc.Name(), name)
		assertString(t, doc.Content(), content)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
