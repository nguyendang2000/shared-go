package mongo

import "go.mongodb.org/mongo-driver/mongo"

// Error messages for the mongo package
const (
	ErrFailedToConnect        = "failed to connect to MongoDB: %v"             // Error when MongoDB connection fails
	ErrFailedToPing           = "failed to ping MongoDB: %v"                   // Error when a MongoDB ping operation fails
	ErrFailedToExecuteFind    = "failed to execute find query: %v"             // Error when a find query fails
	ErrFailedToDecodeDocument = "failed to decode document: %v"                // Error when decoding a document from MongoDB
	ErrFailedToInsertDocument = "failed to insert document: %v"                // Error when inserting a document into MongoDB
	ErrFailedToCountDocuments = "failed to count documents: %v"                // Error when counting documents in a collection
	ErrCursorError            = "cursor error: %v"                             // Error when iterating over a MongoDB cursor
	ErrFailedToFindOne        = "failed to execute find one query: %v"         // Error when a find one query fails
	ErrFailedToDeleteDocument = "failed to delete document: %v"                // Error when deleting a document from MongoDB
	ErrFailedToUpdateDocument = "failed to update document: %v"                // Error when updating a document in MongoDB
	ErrFailedToCheckExistence = "failed to check if document exists: %v"       // Error when checking for the existence of a document
	ErrInvalidResultArgument  = "result argument must be a pointer to a slice" // Error when the result argument is not a pointer to a slice
)

// Alias for mongo.ErrNoDocuments to represent a document not found error.
var ErrDocumentNotFound = mongo.ErrNoDocuments
