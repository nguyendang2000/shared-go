package mongo

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne retrieves a single document from the specified collection using the provided query filter.
// The result is unmarshaled into the specified struct. It uses the timeout defined in the Service struct.
func (inst *Service) FindOne(dbName, collectionName string, query *Query, result interface{}) error {
	// Create a context with the timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Execute FindOne and decode the result
	err := collection.FindOne(ctx, query.Filter).Decode(result)
	if err != nil {
		// Return ErrDocumentNotFound if no documents are found
		if err == mongo.ErrNoDocuments {
			return ErrDocumentNotFound
		}
		// Return other errors with context
		return fmt.Errorf(ErrFailedToFindOne, err)
	}

	// Return nil when the document is found and decoded
	return nil
}

// FindMany retrieves multiple documents from the specified collection using the provided query filter.
// It allows the user to specify a limit, offset, sorting criteria, and unmarshals the results into the provided struct.
// The function uses the timeout defined in the Service struct.
func (inst *Service) FindMany(dbName, collectionName string, query *Query, limit int64, offset int64, sort []string, result interface{}) error {
	// Create a context with the timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Set query options: limit, offset, and sorting
	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(limit)
	}
	if offset > 0 {
		findOptions.SetSkip(offset)
	}

	// Parse the sort parameter and convert it to MongoDB sort format
	sortFields := bson.D{}
	for _, s := range sort {
		order := 1 // Default to ascending order
		field := s

		// Check for a + or - sign to set the sorting order
		if len(s) > 1 && (s[0] == '+' || s[0] == '-') {
			field = s[1:] // Remove the first character (+ or -)
			if s[0] == '-' {
				order = -1 // Descending order
			}
		}

		sortFields = append(sortFields, bson.E{Key: field, Value: order})
	}

	// Apply the sort options if provided
	if len(sortFields) > 0 {
		findOptions.SetSort(sortFields)
	}

	// Execute the query and retrieve the cursor for the results
	cursor, err := collection.Find(ctx, query.Filter, findOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToExecuteFind, err)
	}
	defer cursor.Close(ctx)

	// Unmarshal the results into the provided struct
	if err := cursor.All(ctx, result); err != nil {
		return fmt.Errorf(ErrFailedToDecodeDocument, err)
	}

	return nil
}

// FindAll retrieves all documents from a collection using pagination to avoid memory overload.
// It iteratively calls FindMany in batches until all records are retrieved.
// The function ensures that the result argument is a pointer to a slice.
func (inst *Service) FindAll(dbName, collectionName string, query *Query, sort []string, batchSize int64, result interface{}) error {
	// Set a default batch size if the provided batch size is 0 or less
	if batchSize <= 0 {
		batchSize = DefaultBatchSize // Use the default batch size
	}

	// Ensure the result argument is a pointer to a slice
	resultValue := reflect.ValueOf(result)
	resultSlice := resultValue.Elem()
	if resultSlice.Kind() != reflect.Slice {
		return errors.New(ErrInvalidResultArgument)
	}

	offset := int64(0)

	// Fetch documents in batches until all records are retrieved
	for {
		// Create a batch result placeholder for decoding
		batchResultPtr := reflect.New(resultSlice.Type())
		batchResult := batchResultPtr.Elem()

		// Fetch a batch of documents
		err := inst.FindMany(dbName, collectionName, query, batchSize, offset, sort, batchResult.Addr().Interface())
		if err != nil {
			return err
		}

		// Append the batch result to the main result slice
		for i := 0; i < batchResult.Len(); i++ {
			resultSlice.Set(reflect.Append(resultSlice, batchResult.Index(i)))
		}

		// Stop when the batch size is less than the requested size
		if int64(batchResult.Len()) < batchSize {
			break // No more data to fetch
		}

		// Update the offset for the next batch
		offset += batchSize
	}

	return nil
}

// Exists checks whether a document matching the provided query filter exists in the collection.
// It returns a boolean indicating the existence of the document and an error if any occurs during execution.
func (inst *Service) Exists(dbName, collectionName string, query *Query) (bool, error) {
	var result bson.M // Placeholder for the result

	// Use FindOne to check for the document
	err := inst.FindOne(dbName, collectionName, query, &result)
	if err != nil {
		if errors.Is(err, ErrDocumentNotFound) {
			// No document found, return false without error
			return false, nil
		}
		// Return false and the error if any other issue occurs
		return false, fmt.Errorf(ErrFailedToCheckExistence, err)
	}

	// Document found, return true
	return true, nil
}
