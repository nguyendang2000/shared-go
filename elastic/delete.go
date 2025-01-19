package elastic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// DeleteByID removes a document by its unique ID from the specified index.
// It returns an error if the document is not found or could not be deleted.
func (inst *Service) DeleteByID(index string, id string) error {
	// Set a timeout for the request using the configured timeout value
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute delete request by document ID
	response, err := inst.client.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrDeletingDocument, err)
	}

	// Ensure the document was successfully deleted
	if response.Result != result.Deleted {
		return fmt.Errorf(ErrDocumentNotDeleted, id, index)
	}

	return nil
}

// Delete removes all documents in the specified index that match the provided query.
// It executes a delete-by-query operation and returns an error if the operation fails.
func (inst *Service) Delete(index string, query *Query) error {
	// Set a timeout for the request using the configured timeout value
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute the delete-by-query request
	response, err := inst.client.DeleteByQuery(index).Query(query.q).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrDeletingDocuments, err)
	}

	// Check for failures in the delete response
	if len(response.Failures) > 0 {
		return errors.New("encountered failures during delete-by-query")
	}

	return nil
}
