package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	minio_tags "github.com/minio/minio-go/v7/pkg/tags"
)

// GetObject retrieves an object from the specified bucket using the provided object name.
// It returns the object as a byte array, allowing for further processing.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) GetObject(bucket, object string, opts ...GetObjectOptions) ([]byte, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []GetObjectOptions{{}}
	}

	// Use MinIO's GetObject method to retrieve the object.
	minioObject, err := inst.client.GetObject(ctx, bucket, object, minio.GetObjectOptions(opts[0]))
	if err != nil {
		return nil, fmt.Errorf(ErrGetObject, bucket, err)
	}

	// Read the object into a byte array.
	data, err := io.ReadAll(minioObject)
	if err != nil {
		return nil, fmt.Errorf(ErrReadObject, object, err)
	}

	return data, nil
}

// FGetObject downloads an object from the specified bucket and saves it to the provided file path.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) FGetObject(bucket, object, filePath string, opts ...GetObjectOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []GetObjectOptions{{}}
	}

	// Use MinIO's FGetObject method to download the object and save it locally.
	err := inst.client.FGetObject(ctx, bucket, object, filePath, minio.GetObjectOptions(opts[0]))
	if err != nil {
		return fmt.Errorf(ErrGetObject, bucket, err)
	}

	return nil
}

// PutObject uploads an object to the specified bucket using the provided object name and byte data.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) PutObject(bucket, object string, objectBytes []byte, objectSize int64, opts ...PutObjectOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []PutObjectOptions{{}}
	}

	// Upload the object to the specified bucket using MinIO's PutObject method.
	_, err := inst.client.PutObject(ctx, bucket, object, bytes.NewReader(objectBytes), objectSize, minio.PutObjectOptions(opts[0]))
	if err != nil {
		return fmt.Errorf(ErrPutObject, bucket, err)
	}

	return nil
}

// FPutObject uploads a file from the local filesystem to the specified bucket.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) FPutObject(bucket, object, filePath string, opts ...PutObjectOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []PutObjectOptions{{}}
	}

	// Upload the file from the local filesystem to the specified bucket using MinIO's FPutObject method.
	_, err := inst.client.FPutObject(ctx, bucket, object, filePath, minio.PutObjectOptions(opts[0]))
	if err != nil {
		return fmt.Errorf(ErrPutObject, bucket, err)
	}

	return nil
}

// CopyObject copies an object from a source bucket to a destination bucket.
// It uses the src and dest bucket/object names as parameters, along with the timeout from the Service struct.
func (inst *Service) CopyObject(srcBucket, srcObject, destBucket, destObject string) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Set up the source and destination options.
	srcOpts := minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: srcObject,
	}
	destOpts := minio.CopyDestOptions{
		Bucket: destBucket,
		Object: destObject,
	}

	// Perform the object copy operation.
	_, err := inst.client.CopyObject(ctx, destOpts, srcOpts)
	if err != nil {
		return fmt.Errorf(ErrCopyObject, srcBucket, srcObject, destBucket, destObject, err)
	}

	return nil
}

// StatObject retrieves metadata about an object in the specified bucket.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) StatObject(bucket, object string, opts ...StatObjectOptions) (minio.ObjectInfo, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []StatObjectOptions{{}}
	}

	// Get object metadata from the specified bucket.
	objectInfo, err := inst.client.StatObject(ctx, bucket, object, minio.StatObjectOptions(opts[0]))
	if err != nil {
		return minio.ObjectInfo{}, fmt.Errorf(ErrStatObject, object, bucket, err)
	}

	return objectInfo, nil
}

// RemoveObject deletes a single object from the specified bucket.
// It supports optional MinIO object options for customization and uses the timeout from the Service struct.
func (inst *Service) RemoveObject(bucket, object string, opts ...RemoveObjectOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []RemoveObjectOptions{{}}
	}

	// Remove the specified object from the bucket.
	err := inst.client.RemoveObject(ctx, bucket, object, minio.RemoveObjectOptions(opts[0]))
	if err != nil {
		return fmt.Errorf(ErrRemoveObject, object, bucket, err)
	}

	return nil
}

// PutObjectTags assigns a set of key-value tags to the specified object in a bucket.
// It converts the provided map into MinIO's tag format and applies it to the object.
func (inst *Service) PutObjectTags(bucket, object string, tags map[string]string, opts ...PutObjectTaggingOptions) error {
	// Convert the provided map into MinIO's tag format.
	objectTags, err := minio_tags.NewTags(tags, false)
	if err != nil {
		return fmt.Errorf(ErrCreateTags, err)
	}

	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []PutObjectTaggingOptions{{}}
	}

	// Apply the tags to the specified object.
	if err := inst.client.PutObjectTagging(ctx, bucket, object, objectTags, minio.PutObjectTaggingOptions(opts[0])); err != nil {
		return fmt.Errorf(ErrCreateTags, err)
	}

	return nil
}

// GetObjectTags retrieves the tags assigned to the specified object in a bucket.
// It returns the tags as a map of key-value pairs.
func (inst *Service) GetObjectTags(bucket, object string, opts ...GetObjectTaggingOptions) (map[string]string, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []GetObjectTaggingOptions{{}}
	}

	// Retrieve the tags from the specified object.
	tags, err := inst.client.GetObjectTagging(ctx, bucket, object, minio.GetObjectTaggingOptions(opts[0]))
	if err != nil {
		return nil, fmt.Errorf(ErrGetTags, err)
	}

	// Convert the retrieved tags into a map and return them.
	return tags.ToMap(), nil
}

// RemoveObjectTags removes all tags assigned to the specified object in a bucket.
// It uses the timeout from the Service struct to limit execution time.
func (inst *Service) RemoveObjectTags(bucket, object string, opts ...RemoveObjectTaggingOptions) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Initialize opts with a default empty struct if not provided.
	if len(opts) == 0 {
		opts = []RemoveObjectTaggingOptions{{}}
	}

	// Remove all tags from the specified object.
	if err := inst.client.RemoveObjectTagging(ctx, bucket, object, minio.RemoveObjectTaggingOptions(opts[0])); err != nil {
		return fmt.Errorf(ErrRemoveTags, err)
	}

	return nil
}

// RemoveIncompleteUpload deletes an incomplete multipart upload for the specified object in a bucket.
// It uses the timeout from the Service struct to limit execution time.
func (inst *Service) RemoveIncompleteUpload(bucket, object string) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Remove the incomplete upload for the specified object.
	if err := inst.client.RemoveIncompleteUpload(ctx, bucket, object); err != nil {
		return fmt.Errorf(ErrRemoveIncompleteUpload, object, bucket, err)
	}

	return nil
}
