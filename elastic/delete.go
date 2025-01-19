package elastic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// DeleteByID deletes a document by its unique ID from the specified index.
// Returns an error if the document could not be deleted.
func (inst *Service) DeleteByID(index string, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute delete request by document ID
	response, err := inst.client.Delete(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrDeletingDocument, err)
	}

	// Ensure the document was deleted
	if response.Result != result.Deleted {
		return fmt.Errorf(ErrDocumentNotDeleted, id, index)
	}

	return nil
}

// Delete deletes all documents in the specified index that match the provided query.
// Returns an error if the delete-by-query operation encounters issues.
func (inst *Service) Delete(index string, query *Query) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Execute the delete-by-query request
	response, err := inst.client.DeleteByQuery(index).Query(query.q).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrDeletingDocuments, err)
	}

	// Check for errors in the delete response
	if len(response.Failures) > 0 {
		return errors.New("encountered failures during delete-by-query")
	}

	return nil
}
