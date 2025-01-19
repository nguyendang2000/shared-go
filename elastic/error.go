package elastic

// Configuration Errors
var (
	// ErrNoAddresses represents an error when no addresses are provided in the configuration.
	ErrNoAddresses = "no addresses provided in the configuration"
	// ErrOpeningCACert represents an error when there is an issue opening the CA certificate file.
	ErrOpeningCACert = "error opening CA certificate file: %w"
	// ErrCreatingElasticClient represents an error when creating the Elasticsearch client fails.
	ErrCreatingElasticClient = "error creating Elasticsearch client: %w"
)

// Indexing Errors
var (
	// ErrIndexingDocument represents an error when indexing a document fails.
	ErrIndexingDocument = "failed to index document: %w"
	// ErrIndexingDocuments represents an error when indexing multiple documents fails.
	ErrIndexingDocuments = "failed to index multiple documents: %w"
	// ErrMarshalingDocument represents an error when marshaling a document fails.
	ErrMarshalingDocument = "failed to marshal document: %w"
	// ErrAssigningDocument represents an error when assigning a document to the result fails.
	ErrAssigningDocument = "failed to assign document to result: %w"
	// ErrMarshalingDocuments represents an error when marshaling multiple documents fails.
	ErrMarshalingDocuments = "failed to marshal documents: %w"
	// ErrAssigningDocuments represents an error when assigning multiple documents to the result fails.
	ErrAssigningDocuments = "failed to assign documents to result: %w"
)

// Document Retrieval Errors
var (
	// ErrGettingDocument represents an error when retrieving a document fails.
	ErrGettingDocument = "failed to get document: %w"
	// ErrDocumentNotFound represents an error when a document is not found in the specified index.
	ErrDocumentNotFound = "document with ID %s not found in index %s"
	// ErrUnmarshalingDocument represents an error when unmarshaling a document into the result fails.
	ErrUnmarshalingDocument = "failed to unmarshal document into result: %w"
	// ErrUnmarshalingDocuments represents an error when unmarshaling multiple documents fails.
	ErrUnmarshalingDocuments = "failed to unmarshal documents: %w"
)

// Document Deletion Errors
var (
	// ErrDeletingDocument represents an error when deleting a document fails.
	ErrDeletingDocument = "failed to delete document: %w"
	// ErrDeletingDocuments represents an error when deleting documents by query fails.
	ErrDeletingDocuments = "failed to delete documents by query: %w"
	// ErrDocumentNotDeleted represents an error when a document is not found or could not be deleted in the specified index.
	ErrDocumentNotDeleted = "document with ID %s not found or could not be deleted in index %s"
)

// Search and Query Errors
var (
	// ErrSearchingDocuments represents an error when a search query fails to execute.
	ErrSearchingDocuments = "failed to execute search query: %w"
	// ErrDecodingSearchResponse represents an error when decoding a search response into the result fails.
	ErrDecodingSearchResponse = "failed to decode search response into result: %w"
	// ErrCountingDocuments represents an error when counting documents fails.
	ErrCountingDocuments = "failed to count documents: %w"
	// ErrCheckingDocumentExists represents an error when checking if a document exists fails.
	ErrCheckingDocumentExists = "failed to check if document exists: %w"
)

// General Errors
var (
	// ErrMarshalingSource represents an error when marshaling a document source fails.
	ErrMarshalingSource = "failed to marshal document source: %w"
)
