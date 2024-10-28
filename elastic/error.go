package elastic

import "errors"

// Configuration Errors
var (
	ErrNoAddresses           = errors.New("no addresses provided in the configuration")
	ErrOpeningCACert         = errors.New("error opening CA certificate file")
	ErrCreatingElasticClient = errors.New("error creating Elasticsearch client")
)

// Indexing Errors
var (
	ErrIndexingDocument    = errors.New("failed to index document")
	ErrIndexingDocuments   = errors.New("failed to index multiple documents")
	ErrMarshalingDocument  = errors.New("failed to marshal document")
	ErrAssigningDocument   = errors.New("failed to assign document to result")
	ErrMarshalingDocuments = errors.New("failed to marshal documents")
	ErrAssigningDocuments  = errors.New("failed to assign documents to result")
)

// Document Retrieval Errors
var (
	ErrGettingDocument       = errors.New("failed to get document")
	ErrDocumentNotFound      = errors.New("document not found in specified index")
	ErrUnmarshalingDocument  = errors.New("failed to unmarshal document into result")
	ErrUnmarshalingDocuments = errors.New("failed to unmarshal documents")
)

// Document Deletion Errors
var (
	ErrDeletingDocument   = errors.New("failed to delete document")
	ErrDeletingDocuments  = errors.New("failed to delete documents by query")
	ErrDocumentNotDeleted = errors.New("document not found or could not be deleted in specified index")
)

// Search and Query Errors
var (
	ErrSearchingDocuments     = errors.New("failed to execute search query")
	ErrDecodingSearchResponse = errors.New("failed to decode search response into result")
	ErrCountingDocuments      = errors.New("failed to count documents")
	ErrCheckingDocumentExists = errors.New("failed to check if document exists")
)

// General Errors
var (
	ErrMarshalingSource = errors.New("failed to marshal document source")
)
