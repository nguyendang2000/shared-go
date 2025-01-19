package elastic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// IndexOne indexes or updates a single document in the specified index.
// The document must implement the Document interface, which provides a unique ID.
func (inst *Service) IndexOne(index string, doc Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Attempt to index the document with the specified ID
	_, err := inst.client.Index(index).Id(doc.GetID()).Request(doc).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrIndexingDocument, err)
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
		return fmt.Errorf(ErrIndexingDocuments, err)
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
		return fmt.Errorf(ErrIndexingDocuments, strings.Join(bulkErrors, "; "))
	}

	return nil
}
