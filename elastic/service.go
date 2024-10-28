package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/count"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// Service represents an Elasticsearch service with a configured client and timeout setting.
type Service struct {
	client  *elasticsearch.TypedClient
	timeout int64
}

// NewService initializes a new Elasticsearch service with the provided configuration.
// Returns an error if required configuration fields are missing or if the client cannot be created.
func NewService(cfg Config) (*Service, error) {
	if len(cfg.Addresses) == 0 {
		return nil, ErrNoAddresses
	}

	// Prepare Elasticsearch configuration
	esConfig := elasticsearch.Config{
		Addresses: cfg.Addresses,
		Username:  cfg.Username,
		Password:  cfg.Password,
	}

	// Optional: Add CertificateFingerprint if provided
	if cfg.CertificateFingerprint != "" {
		esConfig.CertificateFingerprint = cfg.CertificateFingerprint
	}

	// Optional: Load CA certificate
	if cfg.CACert != "" {
		caCert, err := os.ReadFile(cfg.CACert)
		if err != nil {
			return nil, ErrOpeningCACert
		}
		esConfig.CACert = caCert
	}

	// Set timeout
	timeout := cfg.Timeout
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

// IndexOne indexes or updates a single document in the specified index.
// The document must implement the Document interface, which provides a unique ID.
func (inst *Service) IndexOne(index string, doc Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Attempt to index the document with the specified ID
	_, err := inst.client.Index(index).Id(doc.GetID()).Request(doc).Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrIndexingDocument, err)
	}

	return nil
}

// Index indexes multiple documents in the specified index.
// Each document must implement the Document interface, which provides a unique ID for each document.
func (inst *Service) Index(index string, docs []Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Start a bulk request for multiple documents
	bulkRequest := inst.client.Bulk().Index(index)

	// Add each document to the bulk request with its custom ID
	for _, doc := range docs {
		id := new(string)
		*id = doc.GetID()
		bulkRequest.IndexOp(types.IndexOperation{
			DynamicTemplates: make(map[string]string),
			Id_:              id,
		}, doc)
	}

	// Execute the bulk indexing request
	response, err := bulkRequest.Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrIndexingDocuments, err)
	}

	// Aggregate any errors in the bulk response items
	var bulkErrors []string
	for _, item := range response.Items {
		for _, result := range item {
			if result.Error != nil {
				bulkErrors = append(bulkErrors, fmt.Sprintf("document ID %s: %v", *result.Id_, result.Error))
			}
		}
	}

	// If there were any bulk errors, return a combined error message
	if len(bulkErrors) > 0 {
		return fmt.Errorf("%w: %s", ErrIndexingDocuments, strings.Join(bulkErrors, "; "))
	}

	return nil
}

// SearchByID retrieves a single document by its unique ID from the specified index.
// Unmarshals the document into the provided result object. Returns an error if the document is not found.
func (inst *Service) SearchByID(index string, id string, result Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Attempt to retrieve the document by ID
	response, err := inst.client.Get(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrGettingDocument, err)
	}

	// Check if the document was found
	if !response.Found {
		return fmt.Errorf("%w with ID %s in index %s", ErrDocumentNotFound, id, index)
	}

	// Unmarshal the source into the result object
	if err := json.Unmarshal(response.Source_, result); err != nil {
		return fmt.Errorf("%w: %s", ErrUnmarshalingDocument, err)
	}

	result.SetID(response.Id_)

	return nil
}

// Search performs a search query on the specified index with pagination and sorting options.
// The matching documents are unmarshaled into the specified result slice, and document IDs are set.
func (inst *Service) Search(index string, query *Query, limit int64, offset int64, sort []string, result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Prepare sorting options based on field prefixes
	sortOptions := make(map[string]string, len(sort))
	for _, field := range sort {
		if len(field) > 0 {
			if field[0] == '+' {
				sortOptions[field[1:]] = "asc"
			} else if field[0] == '-' {
				sortOptions[field[1:]] = "desc"
			} else {
				sortOptions[field] = "asc" // Default to ascending if no prefix is provided
			}
		}
	}

	// Execute the search request with pagination and sorting
	response, err := inst.client.Search().Index(index).Query(query.q).Size(int(limit)).From(int(offset)).Sort(sortOptions).Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrSearchingDocuments, err)
	}

	// Ensure result is a pointer to a slice of Document
	resultVal := reflect.ValueOf(result)
	if resultVal.Kind() != reflect.Ptr || resultVal.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result must be a pointer to a slice")
	}
	resultSlice := resultVal.Elem()
	elemType := resultSlice.Type().Elem()

	// Ensure that the slice element implements the Document interface
	docType := reflect.TypeOf((*Document)(nil)).Elem()
	if !elemType.Implements(docType) {
		return fmt.Errorf("result slice elements must implement the Document interface")
	}

	// Populate result slice with documents, setting IDs
	for _, hit := range response.Hits.Hits {
		elem := reflect.New(elemType).Interface()

		// Unmarshal document data into the element
		if err := json.Unmarshal(hit.Source_, elem); err != nil {
			return fmt.Errorf("%w: %s", ErrUnmarshalingDocuments, err)
		}

		// Set document ID using SetID
		doc := elem.(Document)
		doc.SetID(*hit.Id_)

		// Append the populated element to the result slice
		resultSlice = reflect.Append(resultSlice, reflect.ValueOf(elem).Elem())
	}

	// Set the modified result slice back to the original result pointer
	resultVal.Elem().Set(resultSlice)

	return nil
}

// SearchOne performs a search query on the specified index and retrieves a single matching document.
// The document is unmarshaled into the specified result object. Returns an error if no document is found.
func (inst *Service) SearchOne(index string, query *Query, result interface{}) error {
	var results []map[string]interface{}

	// Perform a search with limit 1 using the Search function
	err := inst.Search(index, query, 1, 0, nil, &results)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrSearchingDocuments, err)
	}

	// Ensure a document was found
	if len(results) == 0 {
		return ErrDocumentNotFound
	}

	// Marshal the first document and unmarshal into result
	bts, err := json.Marshal(results[0])
	if err != nil {
		return fmt.Errorf("%w: %s", ErrMarshalingDocument, err)
	}

	if err = json.Unmarshal(bts, result); err != nil {
		return fmt.Errorf("%w: %s", ErrAssigningDocument, err)
	}

	return nil
}

// DeleteByID deletes a document by its unique ID from the specified index.
// Returns an error if the document could not be deleted.
func (inst *Service) DeleteByID(index string, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute delete request by document ID
	response, err := inst.client.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDeletingDocument, err)
	}

	// Ensure the document was deleted
	if response.Result != result.Deleted {
		return fmt.Errorf("%w with ID %s in index %s", ErrDocumentNotDeleted, id, index)
	}

	return nil
}

// DeleteDocuments deletes all documents in the specified index that match the provided query.
// Returns an error if the delete-by-query operation encounters issues.
func (inst *Service) DeleteDocuments(index string, query *Query) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute the delete-by-query request
	response, err := inst.client.DeleteByQuery(index).Query(query.q).Do(ctx)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrDeletingDocuments, err)
	}

	// Check for errors in the delete response
	if len(response.Failures) > 0 {
		return fmt.Errorf("%w: encountered failures during delete-by-query", ErrDeletingDocuments)
	}

	return nil
}

// Count returns the number of documents in a specified index that match the provided query.
func (inst *Service) Count(index string, query Query) (int64, error) {
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

// Exists checks if there is at least one document in the specified index that matches the provided query.
// Returns true if any matching document exists, false otherwise.
func (inst *Service) Exists(index string, query *Query) (bool, error) {
	var result []map[string]interface{}

	// Perform a search with limit 1 to check for document existence
	err := inst.Search(index, query, 1, 0, nil, &result)
	if err != nil {
		return false, fmt.Errorf("%w: %s", ErrCheckingDocumentExists, err)
	}

	// Return true if any document was found
	return len(result) > 0, nil
}
