package elastic

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
)

// Service represents an Elasticsearch service that manages an Elasticsearch client
// and provides a timeout setting for request operations.
type Service struct {
	client  *elasticsearch.TypedClient // The Elasticsearch client for executing requests
	timeout int64                      // Timeout duration in seconds for requests
	context context.Context            // Context used for managing request timeouts
}

// NewService initializes a new Elasticsearch service using the provided configuration.
// It validates required fields, loads optional TLS settings, and creates an Elasticsearch client.
// Returns an error if no addresses are provided, if the CA certificate file cannot be read,
// or if the Elasticsearch client fails to initialize.
func NewService(ctx context.Context, conf Config) (*Service, error) {
	// Validate configuration: Ensure at least one address is provided
	if len(conf.Addresses) == 0 {
		return nil, errors.New(ErrNoAddresses)
	}

	// Prepare Elasticsearch client configuration
	esConfig := elasticsearch.Config{
		Addresses: conf.Addresses, // Elasticsearch server addresses
		Username:  conf.Username,  // Optional: Basic authentication username
		Password:  conf.Password,  // Optional: Basic authentication password
	}

	// Optional: Set Certificate Fingerprint for secure connections
	if conf.CertificateFingerprint != "" {
		esConfig.CertificateFingerprint = conf.CertificateFingerprint
	}

	// Optional: Load CA certificate from the specified file
	if conf.CACert != "" {
		caCert, err := os.ReadFile(conf.CACert)
		if err != nil {
			return nil, fmt.Errorf(ErrOpeningCACert, err)
		}
		esConfig.CACert = caCert
	}

	// Set timeout from configuration, using the default if none is provided
	timeout := conf.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// Create a new Elasticsearch client with the configured settings
	client, err := elasticsearch.NewTypedClient(esConfig)
	if err != nil {
		return nil, fmt.Errorf(ErrCreatingElasticClient, err)
	}

	// Return the initialized Service instance
	return &Service{
		client:  client,
		context: ctx,
		timeout: timeout}, nil
}

// Client returns the internal Elasticsearch client, allowing direct access to its API methods.
func (inst *Service) Client() *elasticsearch.TypedClient {
	return inst.client
}

// Count returns the number of documents in a specified index that match the provided query.
// It executes a count request using the given query and returns the document count.
// The function uses the timeout field from the Service struct to set a request deadline.
func (inst *Service) Count(index string, query *Query) (int64, error) {
	// Set a timeout for the request using the configured timeout value
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Execute the count request with the provided query
	response, err := inst.client.Count().Index(index).Request(&count.Request{
		Query: query.q,
	}).Do(ctx)
	if err != nil {
		return 0, fmt.Errorf(ErrCountingDocuments, err)
	}

	return response.Count, nil
}

// Exists checks if at least one document matching the provided query exists in the specified index.
// It calls the Count method and returns true if the document count is greater than zero.
// The function propagates any errors encountered while executing the count request.
func (inst *Service) Exists(index string, query *Query) (bool, error) {
	// Get the document count using the Count method
	count, err := inst.Count(index, query)
	if err != nil {
		return false, fmt.Errorf(ErrCheckingDocumentExists, err)
	}

	// Return true if at least one document exists, otherwise return false
	return count != 0, nil
}
