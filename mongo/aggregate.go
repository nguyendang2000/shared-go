package mongo

import (
	"context"
	"fmt"
	"time"
)

// Count returns the number of documents matching the given query.
// It uses the timeout field from the Service struct.
func (inst *Service) Count(dbName, collectionName string, query *Query) (int64, error) {
	// Use the timeout from the Service struct
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Count the number of documents matching the query
	count, err := collection.CountDocuments(ctx, query.Filter)
	if err != nil {
		return 0, fmt.Errorf(ErrFailedToCountDocuments, err)
	}

	return count, nil
}

// EstimatedDocumentCount returns an estimated count of documents in the specified collection.
// Unlike Count, this method provides a faster approximation without scanning all documents.
// It uses the timeout field from the Service struct.
func (inst *Service) EstimatedDocumentCount(dbName, collectionName string) (int64, error) {
	// Use the timeout from the Service struct
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Retrieve an estimated count of documents in the collection
	count, err := collection.EstimatedDocumentCount(ctx)
	if err != nil {
		return 0, fmt.Errorf(ErrFailedToEstimateCount, err)
	}

	return count, nil
}

// Distinct retrieves the distinct values for a specified field in a collection,
// based on the provided query filter. It returns a slice of unique values.
// This method uses the timeout field from the Service struct.
func (inst *Service) Distinct(dbName, collectionName, fieldName string, query *Query) ([]interface{}, error) {
	// Use the timeout from the Service struct
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Retrieve distinct values for the specified field, filtered by the given query
	distinct, err := collection.Distinct(ctx, fieldName, query.Filter)
	if err != nil {
		return nil, fmt.Errorf(ErrFailedToGetDistinct, err)
	}

	return distinct, nil
}
