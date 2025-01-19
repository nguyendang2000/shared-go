package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// SearchByID retrieves a single document by its unique ID from the specified index.
// Unmarshals the document into the provided result object. Returns an error if the document is not found.
func (inst *Service) SearchByID(index string, id string, result Document) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Millisecond)
	defer cancel()

	// Attempt to retrieve the document by ID
	response, err := inst.client.Get(index, id).Do(ctx)
	if err != nil {
		return fmt.Errorf(ErrGettingDocument, err)
	}

	// Check if the document was found
	if !response.Found {
		return fmt.Errorf(ErrDocumentNotFound, id, index)
	}

	// Unmarshal the source into the result object
	if err := json.Unmarshal(response.Source_, result); err != nil {
		return fmt.Errorf(ErrUnmarshalingDocument, err)
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
		return fmt.Errorf(ErrSearchingDocuments, err)
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

	// Iterate through search hits
	for _, hit := range response.Hits.Hits {
		// Create a new instance of *Test
		newElem := reflect.New(elemType.Elem()).Interface() // Creates *Test

		// Unmarshal JSON response into *Test
		if err := json.Unmarshal(hit.Source_, newElem); err != nil {
			return fmt.Errorf(ErrUnmarshalingDocuments, err)
		}

		// If the type implements Document, set its ID
		if doc, ok := newElem.(Document); ok {
			doc.SetID(*hit.Id_)
		}

		// Append the new instance to the result slice
		resultSlice = reflect.Append(resultSlice, reflect.ValueOf(newElem))
	}

	// Set the modified slice back to the result pointer
	resultVal.Elem().Set(resultSlice)

	return nil
}
