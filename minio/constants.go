package minio

import (
	"github.com/minio/minio-go/v7"
)

// DefaultTimeout defines the default request timeout in seconds.
const DefaultTimeout int64 = 24 * 60 * 60

// MakeBucketOptions represents options for creating a new bucket.
type MakeBucketOptions minio.MakeBucketOptions

// ListObjectsOptions represents options for listing objects in a bucket.
type ListObjectsOptions minio.ListObjectsOptions

// GetObjectOptions represents options for retrieving an object from a bucket.
type GetObjectOptions minio.GetObjectOptions

// PutObjectOptions represents options for uploading an object to a bucket.
type PutObjectOptions minio.PutObjectOptions

// StatObjectOptions represents options for retrieving metadata of an object.
type StatObjectOptions minio.StatObjectOptions

// RemoveObjectOptions represents options for deleting an object from a bucket.
type RemoveObjectOptions minio.RemoveObjectOptions

// PutObjectTaggingOptions represents options for setting tags on an object.
type PutObjectTaggingOptions minio.PutObjectTaggingOptions

// GetObjectTaggingOptions represents options for retrieving tags assigned to an object.
type GetObjectTaggingOptions minio.GetObjectTaggingOptions

// RemoveObjectTaggingOptions represents options for removing tags from an object.
type RemoveObjectTaggingOptions minio.RemoveObjectTaggingOptions
