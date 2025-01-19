package elastic

// Document defines the interface for documents stored in an Elasticsearch index.
// It provides methods to retrieve and assign a document's unique identifier.
type Document interface {
	// GetID returns the unique identifier of the document.
	GetID() string

	// SetID assigns a unique identifier to the document.
	SetID(id string)
}
