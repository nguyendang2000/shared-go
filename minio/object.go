package minio

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
)

// GetObject retrieves an object from the specified bucket using the provided object name.
// It returns the object as a byte array, allowing for further processing.
// It uses the timeout from the Service struct.
func (inst *Service) GetObject(bucketName, objectName string) ([]byte, error) {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Use MinIO's GetObject method to retrieve the object
	object, err := inst.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToGetObject, bucketName, err)
	}

	// Read the object into a byte array
	data, err := io.ReadAll(object)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToReadObject, objectName, err)
	}

	return data, nil
}

// FGetObject downloads an object from the specified bucket and saves it to the provided file path.
// It uses the timeout from the Service struct.
func (inst *Service) FGetObject(bucketName, objectName, filePath string) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Use MinIO's FGetObject to download the object and save it locally
	err := inst.client.FGetObject(ctx, bucketName, objectName, filePath, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf(ErrFailedToGetObject, bucketName, err)
	}

	return nil
}

// PutObject uploads an object to the specified bucket using the provided object name and reader.
// It accepts a pointer to minio.PutObjectOptions for additional options and uses the timeout from the Service struct.
func (inst *Service) PutObject(bucketName, objectName string, reader io.Reader, objectSize int64, opts *minio.PutObjectOptions) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// If opts is nil, initialize an empty minio.PutObjectOptions struct
	if opts == nil {
		opts = &minio.PutObjectOptions{}
	}

	// Upload the object to the bucket
	_, err := inst.client.PutObject(ctx, bucketName, objectName, reader, objectSize, *opts)
	if err != nil {
		return fmt.Errorf(ErrFailedToPutObject, bucketName, err)
	}

	return nil
}

// FPutObject uploads a file from the local filesystem to the specified bucket.
// It accepts a pointer to minio.PutObjectOptions for additional options and uses the timeout from the Service struct.
func (inst *Service) FPutObject(bucketName, objectName, filePath string, opts *minio.PutObjectOptions) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// If opts is nil, initialize an empty minio.PutObjectOptions struct
	if opts == nil {
		opts = &minio.PutObjectOptions{}
	}

	// Upload the file to the bucket
	_, err := inst.client.FPutObject(ctx, bucketName, objectName, filePath, *opts)
	if err != nil {
		return fmt.Errorf(ErrFailedToPutObject, bucketName, err)
	}

	return nil
}

// CopyObject copies an object from a source bucket to a destination bucket.
// It uses the src and dest bucket/object names as parameters, along with the timeout from the Service struct.
func (inst *Service) CopyObject(srcBucket, srcObject, destBucket, destObject string) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Set up the source and destination options
	srcOpts := minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: srcObject,
	}
	destOpts := minio.CopyDestOptions{
		Bucket: destBucket,
		Object: destObject,
	}

	// Perform the object copy operation
	_, err := inst.client.CopyObject(ctx, destOpts, srcOpts)
	if err != nil {
		return fmt.Errorf(ErrFailedToCopyObject, srcBucket, srcObject, destBucket, destObject, err)
	}

	return nil
}

// StatObject retrieves metadata about an object in the specified bucket.
// It uses the timeout from the Service struct.
func (inst *Service) StatObject(bucketName, objectName string) (minio.ObjectInfo, error) {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get object metadata
	objectInfo, err := inst.client.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return minio.ObjectInfo{}, fmt.Errorf(ErrFailedToStatObject, objectName, bucketName, err)
	}

	return objectInfo, nil
}

// RemoveObject deletes a single object from the specified bucket.
// It uses the timeout from the Service struct.
func (inst *Service) RemoveObject(bucketName, objectName string) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Remove the object from the bucket
	err := inst.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf(ErrFailedToRemoveObject, objectName, bucketName, err)
	}

	return nil
}
