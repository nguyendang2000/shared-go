package minio

// Error constants for the minio package.
const (
	// ErrFailedToInitializeClient represents an error when the MinIO client initialization fails.
	ErrFailedToInitializeClient = "failed to initialize MinIO client: %v"

	// ErrFailedToGetObject represents an error when fetching an object from a bucket fails.
	ErrFailedToGetObject = "failed to get object from bucket %s: %v"

	// ErrFailedToReadObject represents an error when reading from the fetched object fails.
	ErrFailedToReadObject = "failed to read object %s: %v"

	// ErrFailedToPutObject represents an error when uploading an object to a bucket fails.
	ErrFailedToPutObject = "failed to put object in bucket %s: %v"

	// ErrFailedToCopyObject represents an error when copying an object between buckets fails.
	ErrFailedToCopyObject = "failed to copy object from %s/%s to %s/%s: %v"

	// ErrFailedToStatObject represents an error when retrieving object metadata fails.
	ErrFailedToStatObject = "failed to stat object %s in bucket %s: %v"

	// ErrFailedToRemoveObject represents an error when deleting a single object fails.
	ErrFailedToRemoveObject = "failed to remove object %s from bucket %s: %v"

	// ErrFailedToConnect represents an error when connecting to MinIO fails.
	ErrFailedToConnect = "failed to connect to MinIO: %v"
)
