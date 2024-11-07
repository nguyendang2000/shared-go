package elastic

import "errors"

// Configuration Errors
var (
	// ErrNoAddresses is returned when no addresses are provided in the configuration.
	ErrNoAddresses = errors.New("no addresses provided in the configuration")
	// ErrOpeningCACert is returned when there is an error opening the CA certificate file.
	ErrOpeningCACert = errors.New("error opening CA certificate file")
	// ErrCreatingElasticClient is returned when there is an error creating the Elasticsearch client.
	ErrCreatingElasticClient = errors.New("error creating Elasticsearch client")
)

// Indexing Errors
var (
	// ErrIndexingDocument is returned when indexing a document fails.
	ErrIndexingDocument = errors.New("failed to index document")
	// ErrIndexingDocuments is returned when indexing multiple documents fails.
	ErrIndexingDocuments = errors.New("failed to index multiple documents")
	// ErrMarshalingDocument is returned when marshaling a document fails.
	ErrMarshalingDocument = errors.New("failed to marshal document")
	// ErrAssigningDocument is returned when assigning a document to the result fails.
	ErrAssigningDocument = errors.New("failed to assign document to result")
	// ErrMarshalingDocuments is returned when marshaling multiple documents fails.
	ErrMarshalingDocuments = errors.New("failed to marshal documents")
	// ErrAssigningDocuments is returned when assigning multiple documents to the result fails.
	ErrAssigningDocuments = errors.New("failed to assign documents to result")
)

// Document Retrieval Errors
var (
	// ErrGettingDocument is returned when retrieving a document fails.
	ErrGettingDocument = errors.New("failed to get document")
	// ErrDocumentNotFound is returned when a document is not found in the specified index.
	ErrDocumentNotFound = errors.New("document not found in specified index")
	// ErrUnmarshalingDocument is returned when unmarshaling a document into the result fails.
	ErrUnmarshalingDocument = errors.New("failed to unmarshal document into result")
	// ErrUnmarshalingDocuments is returned when unmarshaling multiple documents fails.
	ErrUnmarshalingDocuments = errors.New("failed to unmarshal documents")
)

// Document Deletion Errors
var (
	// ErrDeletingDocument is returned when deleting a document fails.
	ErrDeletingDocument = errors.New("failed to delete document")
	// ErrDeletingDocuments is returned when deleting documents by query fails.
	ErrDeletingDocuments = errors.New("failed to delete documents by query")
	// ErrDocumentNotDeleted is returned when a document is not found or could not be deleted in the specified index.
	ErrDocumentNotDeleted = errors.New("document not found or could not be deleted in specified index")
)

// Search and Query Errors
var (
	// ErrSearchingDocuments is returned when a search query fails to execute.
	ErrSearchingDocuments = errors.New("failed to execute search query")
	// ErrDecodingSearchResponse is returned when decoding a search response into the result fails.
	ErrDecodingSearchResponse = errors.New("failed to decode search response into result")
	// ErrCountingDocuments is returned when counting documents fails.
	ErrCountingDocuments = errors.New("failed to count documents")
	// ErrCheckingDocumentExists is returned when checking if a document exists fails.
	ErrCheckingDocumentExists = errors.New("failed to check if document exists")
)

// General Errors
var (
	// ErrMarshalingSource is returned when marshaling a document source fails.
	ErrMarshalingSource = errors.New("failed to marshal document source")
)
