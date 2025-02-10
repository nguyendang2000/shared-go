package minio

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	minio_tags "github.com/minio/minio-go/v7/pkg/tags"
)

// MakeBucket creates a new bucket with the specified name.
// It supports optional MinIO bucket options for customization and uses the timeout from the Service struct.
func (inst *Service) MakeBucket(bucket string, opts ...MakeBucketOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if no options are provided.
	if len(opts) == 0 {
		opts = []MakeBucketOptions{{}}
	}

	// Create the bucket using MinIO's MakeBucket method.
	if err := inst.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions(opts[0])); err != nil {
		return fmt.Errorf(ErrMakeBucket, err)
	}

	return nil
}

// ListBuckets retrieves a list of all available buckets.
// It uses the timeout from the Service struct to limit execution time.
func (inst *Service) ListBuckets() ([]minio.BucketInfo, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Retrieve the list of buckets from MinIO.
	buckets, err := inst.client.ListBuckets(ctx)
	if err != nil {
		return nil, fmt.Errorf(ErrListBuckets, err)
	}

	return buckets, nil
}

// BucketExists checks if a bucket exists in the MinIO storage.
// It returns true if the bucket exists, otherwise false, and uses the timeout from the Service struct.
func (inst *Service) BucketExists(bucket string) (bool, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Check if the bucket exists in MinIO.
	exists, err := inst.client.BucketExists(ctx, bucket)
	if err != nil {
		return false, fmt.Errorf(ErrBucketExists, err)
	}

	return exists, nil
}

// RemoveBucket deletes a bucket from the MinIO storage.
// It uses the timeout from the Service struct to limit execution time.
func (inst *Service) RemoveBucket(bucket string) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Remove the specified bucket using MinIO's RemoveBucket method.
	if err := inst.client.RemoveBucket(ctx, bucket); err != nil {
		return fmt.Errorf(ErrRemoveBucket, err)
	}

	return nil
}

// ListObjects retrieves a list of object keys from the specified bucket based on a given prefix.
// It supports recursive listing and optional parameters for additional filtering.
// The function uses the timeout from the Service struct to limit execution time.
func (inst *Service) ListObjects(bucket, prefix string, recursive bool, opts ...ListObjectsOptions) ([]string, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if no options are provided.
	if len(opts) == 0 {
		opts = []ListObjectsOptions{{
			Prefix:    prefix,
			Recursive: recursive,
		}}
	}

	// Retrieve a channel of objects matching the specified criteria.
	objectCh := inst.client.ListObjects(ctx, bucket, minio.ListObjectsOptions(opts[0]))

	// Collect object keys from the result channel.
	objects := make([]string, 0)
	for object := range objectCh {
		if object.Err != nil {
			// Return an error if object retrieval fails.
			return nil, fmt.Errorf(ErrListObjects, object.Err)
		}
		objects = append(objects, object.Key)
	}

	// Return the collected list of object keys.
	return objects, nil
}

// ListIncompleteUploads retrieves a list of incomplete multipart uploads in the specified bucket.
// It supports filtering by prefix and recursive listing, using the timeout from the Service struct.
func (inst *Service) ListIncompleteUploads(bucket, prefix string, recursive bool) ([]string, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize the object list to store results.
	objects := make([]string, 0)
	uploadsCh := inst.client.ListIncompleteUploads(ctx, bucket, prefix, recursive)

	// Iterate over the incomplete uploads and collect their keys.
	for upload := range uploadsCh {
		if upload.Err != nil {
			return nil, fmt.Errorf(ErrListIncompleteUploads, upload.Err)
		}
		objects = append(objects, upload.Key)
	}

	// Return the list of incomplete uploads.
	return objects, nil
}

// SetBucketTags assigns a set of key-value tags to the specified bucket.
// It converts the provided map into MinIO's tag format and applies it to the bucket.
func (inst *Service) SetBucketTags(bucket string, tags map[string]string) error {
	// Convert the provided map into MinIO's tag format.
	bucketTags, err := minio_tags.NewTags(tags, false)
	if err != nil {
		return fmt.Errorf(ErrCreateTags, err)
	}

	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Apply the tags to the specified bucket.
	if err = inst.client.SetBucketTagging(ctx, bucket, bucketTags); err != nil {
		return fmt.Errorf(ErrSetTags, err)
	}

	return nil
}

// GetBucketTags retrieves the tags assigned to the specified bucket.
// It returns the tags as a map of key-value pairs.
func (inst *Service) GetBucketTags(bucket string) (map[string]string, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Retrieve the tags from the specified bucket.
	tags, err := inst.client.GetBucketTagging(ctx, bucket)
	if err != nil {
		return nil, fmt.Errorf(ErrGetTags, err)
	}

	// Convert the retrieved tags into a map and return them.
	return tags.ToMap(), nil
}

// RemoveBucketTags removes all tags assigned to the specified bucket.
// It uses the timeout from the Service struct to limit execution time.
func (inst *Service) RemoveBucketTags(bucket string) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Remove all tags from the specified bucket.
	if err := inst.client.RemoveBucketTagging(ctx, bucket); err != nil {
		return fmt.Errorf(ErrRemoveTags, err)
	}

	return nil
}
