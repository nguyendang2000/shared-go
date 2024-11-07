package elastic

// Document represents an interface for documents that can be stored in an Elasticsearch index.
// It includes methods for retrieving and setting the document's ID.
type Document interface {
	// GetID retrieves the document's ID.
	GetID() string

	// SetID sets the document's ID.
	SetID(id string)
}
