package mongo

import "go.mongodb.org/mongo-driver/mongo"

// Error messages for the mongo package
const (
	ErrFailedToConnect        = "failed to connect to MongoDB: %v"       // Error when MongoDB connection fails
	ErrFailedToPing           = "failed to ping MongoDB: %v"             // Error when MongoDB ping operation fails
	ErrFailedToExecuteFind    = "failed to execute find query: %v"       // Error during a find query
	ErrFailedToDecodeDocument = "failed to decode document: %v"          // Error when decoding a document from the database
	ErrFailedToInsertDocument = "failed to insert document: %v"          // Error when inserting a document into the database
	ErrFailedToCountDocuments = "failed to count documents: %v"          // Error when counting documents in a collection
	ErrCursorError            = "cursor error: %v"                       // Error when iterating over a MongoDB cursor
	ErrFailedToFindOne        = "failed to execute find one query: %v"   // Error during a find one query
	ErrFailedToDeleteDocument = "failed to delete document: %v"          // Error when deleting a document
	ErrFailedToUpdateDocument = "failed to update document: %v"          // Error when updating a document
	ErrFailedToCheckExistence = "failed to check if document exists: %v" // Error when checking document existence
)

// Alias for mongo.ErrNoDocuments to represent a document not found error.
var ErrDocumentNotFound = mongo.ErrNoDocuments
