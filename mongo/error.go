package mongo

import "go.mongodb.org/mongo-driver/mongo"

// Error messages for the mongo package.
const (
	// ErrFailedToConnect represents an error when the MongoDB connection fails.
	ErrFailedToConnect = "failed to connect to MongoDB: %v"

	// ErrFailedToPing represents an error when a MongoDB ping operation fails.
	ErrFailedToPing = "failed to ping MongoDB: %v"

	// ErrFailedToExecuteFind represents an error when a find query fails.
	ErrFailedToExecuteFind = "failed to execute find query: %v"

	// ErrFailedToDecodeDocument represents an error when decoding a document from MongoDB fails.
	ErrFailedToDecodeDocument = "failed to decode document: %v"

	// ErrFailedToInsertDocument represents an error when inserting a document into MongoDB fails.
	ErrFailedToInsertDocument = "failed to insert document: %v"

	// ErrFailedToCountDocuments represents an error when counting documents in a collection fails.
	ErrFailedToCountDocuments = "failed to count documents: %v"

	// ErrCursorError represents an error when iterating over a MongoDB cursor.
	ErrCursorError = "cursor error: %v"

	// ErrFailedToFindOne represents an error when a find one query fails.
	ErrFailedToFindOne = "failed to execute find one query: %v"

	// ErrFailedToDeleteDocument represents an error when deleting a document from MongoDB fails.
	ErrFailedToDeleteDocument = "failed to delete document: %v"

	// ErrFailedToUpdateDocument represents an error when updating a document in MongoDB fails.
	ErrFailedToUpdateDocument = "failed to update document: %v"

	// ErrFailedToCheckExistence represents an error when checking for the existence of a document fails.
	ErrFailedToCheckExistence = "failed to check if document exists: %v"

	// ErrInvalidResultArgument represents an error when the result argument is not a pointer to a slice.
	ErrInvalidResultArgument = "result argument must be a pointer to a slice"
)

// ErrDocumentNotFound is an alias for mongo.ErrNoDocuments to represent a document not found error.
var ErrDocumentNotFound = mongo.ErrNoDocuments
