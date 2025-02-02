package minio

// Error constants for the MinIO package, representing various failure scenarios.
const (
	// ErrClientInitialization occurs when the MinIO client initialization fails.
	ErrClientInitialization = "failed to initialize MinIO client: %w"
)

// Bucket-related error constants.
const (
	// ErrMakeBucket occurs when bucket creation fails.
	ErrMakeBucket = "failed to make bucket: %w"

	// ErrListBuckets occurs when listing buckets fails.
	ErrListBuckets = "failed to list buckets: %w"

	// ErrBucketExists occurs when checking for a bucket's existence fails.
	ErrBucketExists = "failed to check bucket exists: %w"

	// ErrRemoveBucket occurs when bucket deletion fails.
	ErrRemoveBucket = "failed to remove bucket: %w"
)

// Object-related error constants.
const (
	// ErrListObjects occurs when listing objects in a bucket fails.
	ErrListObjects = "failed to list objects: %w"

	// ErrListIncompleteUploads occurs when listing incomplete uploads fails.
	ErrListIncompleteUploads = "failed to list incomplete uploads: %w"
)

// Object operation error constants.
const (
	// ErrGetObject occurs when fetching an object from a bucket fails.
	ErrGetObject = "failed to get object from bucket %s: %w"

	// ErrReadObject occurs when reading from the fetched object fails.
	ErrReadObject = "failed to read object %s: %w"

	// ErrPutObject occurs when uploading an object to a bucket fails.
	ErrPutObject = "failed to put object in bucket %s: %v"

	// ErrCopyObject occurs when copying an object between buckets fails.
	ErrCopyObject = "failed to copy object from %s/%s to %s/%s: %w"

	// ErrStatObject occurs when retrieving object metadata fails.
	ErrStatObject = "failed to stat object %s in bucket %s: %w"

	// ErrRemoveObject occurs when deleting an object from a bucket fails.
	ErrRemoveObject = "failed to remove object %s from bucket %s: %w"

	// ErrRemoveIncompleteUpload occurs when removing an incomplete upload fails.
	ErrRemoveIncompleteUpload = "failed to remove incomplete upload %s from bucket %s: %w"
)

// Tagging-related error constants.
const (
	// ErrCreateTags occurs when creating tags fails.
	ErrCreateTags = "failed to create tags: %w"

	// ErrSetTags occurs when setting tags on a bucket or object fails.
	ErrSetTags = "failed to set tags: %w"

	// ErrGetTags occurs when retrieving tags from a bucket or object fails.
	ErrGetTags = "failed to get tags: %w"

	// ErrRemoveTags occurs when removing tags from a bucket or object fails.
	ErrRemoveTags = "failed to remove tags: %w"
)
