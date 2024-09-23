package minio

// Error constants for the minio package
const (
	ErrFailedToInitializeClient = "failed to initialize MinIO client: %v"         // Error when MinIO client initialization fails
	ErrFailedToGetObject        = "failed to get object from bucket %s: %v"       // Error when fetching an object from a bucket fails
	ErrFailedToReadObject       = "failed to read object %s: %v"                  // Error when reading from the fetched object fails
	ErrFailedToPutObject        = "failed to put object in bucket %s: %v"         // Error when uploading an object to a bucket fails
	ErrFailedToCopyObject       = "failed to copy object from %s/%s to %s/%s: %v" // Error when copying an object between buckets fails
	ErrFailedToStatObject       = "failed to stat object %s in bucket %s: %v"     // Error when retrieving object metadata fails
	ErrFailedToRemoveObject     = "failed to remove object %s from bucket %s: %v" // Error when deleting a single object fails
	ErrFailedToConnect          = "failed to connect to MinIO: %v"                // Error when connecting to MinIO fails
)
