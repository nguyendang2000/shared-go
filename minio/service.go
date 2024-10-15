package minio

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// Service struct contains the MinIO client and a timeout field
type Service struct {
	client  *minio.Client
	timeout int64 // Timeout in seconds for requests
}

// NewService initializes a new MinIO connection using the given configuration
// and sets the timeout in the Service struct.
func NewService(conf Config) (*Service, error) {
	// Set timeout to DefaultTimeout if not provided or less than 0
	timeout := conf.Timeout
	if timeout <= 0 {
		timeout = DefaultTimeout
	}

	// Initialize the MinIO client
	minioClient, err := minio.New(conf.Address, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToInitializeClient, err)
	}

	return &Service{
		client:  minioClient,
		timeout: timeout,
	}, nil
}

// Client returns the MinIO client instance
func (inst *Service) Client() *minio.Client {
	return inst.client
}
