package mongo

import "go.mongodb.org/mongo-driver/mongo"

// Error messages for the mongo package.
const (
	// ErrFailedToConnect represents an error when the MongoDB connection fails.
	ErrFailedToConnect = "failed to connect to MongoDB: %w"

	// ErrFailedToPing represents an error when a MongoDB ping operation fails.
	ErrFailedToPing = "failed to ping MongoDB: %w"

	// ErrFailedToExecuteFind represents an error when a find query fails.
	ErrFailedToExecuteFind = "failed to execute find query: %w"

	// ErrFailedToDecodeDocument represents an error when decoding a document from MongoDB fails.
	ErrFailedToDecodeDocument = "failed to decode document: %w"

	// ErrFailedToInsertDocument represents an error when inserting a document into MongoDB fails.
	ErrFailedToInsertDocument = "failed to insert document: %w"

	// ErrFailedToCountDocuments represents an error when counting documents in a collection fails.
	ErrFailedToCountDocuments = "failed to count documents: %w"

	// ErrFailedToEstimateCount represents an error when retrieving an estimated count of documents fails.
	ErrFailedToEstimateCount = "failed to estimate number of documents: %w"

	// ErrFailedToGetDistinct represents an error when retrieving distinct values from a collection fails.
	ErrFailedToGetDistinct = "failed to get distinct values: %w"

	// ErrCursorError represents an error when iterating over a MongoDB cursor.
	ErrCursorError = "cursor error: %w"

	// ErrFailedToFindOne represents an error when a find one query fails.
	ErrFailedToFindOne = "failed to execute find one query: %w"

	// ErrFailedToFindOneAndDelete represents an error when a find one and delete query fails.
	ErrFailedToFindOneAndDelete = "failed to execute find one and delete query: %w"

	// ErrFailedToFindOneAndReplace represents an error when a find one and replace query fails.
	ErrFailedToFindOneAndReplace = "failed to execute find one and replace query: %w"

	// ErrFailedToFindOneAndUpdate represents an error when a find one and update query fails.
	ErrFailedToFindOneAndUpdate = "failed to execute find one and update query: %w"

	// ErrFailedToDeleteDocument represents an error when deleting a document from MongoDB fails.
	ErrFailedToDeleteDocument = "failed to delete document: %w"

	// ErrFailedToUpdateDocument represents an error when updating a document in MongoDB fails.
	ErrFailedToUpdateDocument = "failed to update document: %w"

	// ErrFailedToCheckExistence represents an error when checking for the existence of a document fails.
	ErrFailedToCheckExistence = "failed to check if document exists: %w"

	// ErrInvalidResultArgument represents an error when the result argument is not a pointer to a slice.
	ErrInvalidResultArgument = "result argument must be a pointer to a slice"
)

// ErrDocumentNotFound is an alias for mongo.ErrNoDocuments to represent a document not found error.
var ErrDocumentNotFound = mongo.ErrNoDocuments
