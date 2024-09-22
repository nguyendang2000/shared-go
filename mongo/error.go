package mongo

// Error messages for the mongo package
const (
	ErrFailedToConnect        = "failed to connect to MongoDB: %v"
	ErrFailedToPing           = "failed to ping MongoDB: %v"
	ErrFailedToExecuteFind    = "failed to execute find query: %v"
	ErrFailedToDecodeDocument = "failed to decode document: %v"
	ErrFailedToInsertDocument = "failed to insert document: %v"
	ErrFailedToCountDocuments = "failed to count documents: %v"
	ErrCursorError            = "cursor error: %v"
	ErrFailedToFindOne        = "failed to execute find one query: %v"
	ErrFailedToDeleteDocument = "failed to delete document: %v"
	ErrFailedToUpdateDocument = "failed to update document: %v"
)
