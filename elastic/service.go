package elastic

import (
	"context"
	"fmt"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
)

// Service represents an Elasticsearch service with a configured client and timeout setting.
type Service struct {
	client  *elasticsearch.TypedClient
	timeout int64
}

// NewService initializes a new Elasticsearch service with the provided configuration.
// Returns an error if required configuration fields are missing or if the client cannot be created.
func NewService(conf Config) (*Service, error) {
	if len(conf.Addresses) == 0 {
		return nil, ErrNoAddresses
	}

	// Prepare Elasticsearch configuration
	esConfig := elasticsearch.Config{
		Addresses: conf.Addresses,
		Username:  conf.Username,
		Password:  conf.Password,
	}

	// Optional: Add CertificateFingerprint if provided
	if conf.CertificateFingerprint != "" {
		esConfig.CertificateFingerprint = conf.CertificateFingerprint
	}

	// Optional: Load CA certificate
	if conf.CACert != "" {
		caCert, err := os.ReadFile(conf.CACert)
		if err != nil {
			return nil, ErrOpeningCACert
		}
		esConfig.CACert = caCert
	}

	// Set timeout
	timeout := conf.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// Create Elasticsearch client
	client, err := elasticsearch.NewTypedClient(esConfig)
	if err != nil {
		return nil, ErrCreatingElasticClient
	}

	return &Service{client: client, timeout: timeout}, nil
}

// Client returns the internal Elasticsearch client, allowing direct API access.
func (inst *Service) Client() *elasticsearch.TypedClient {
	return inst.client
}

// Count returns the number of documents in a specified index that match the provided query.
func (inst *Service) Count(index string, query *Query) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute the count request with the provided query
	response, err := inst.client.Count().Index(index).Request(&count.Request{
		Query: query.q,
	}).Do(ctx)
	if err != nil {
		return 0, fmt.Errorf("%w: %s", ErrCountingDocuments, err)
	}

	return response.Count, nil
}
