package mongo

import (
	"context"
	"fmt"
	"time"
)

// DeleteOne deletes a single document from the collection that matches the filter in the Query struct.
// It uses the timeout from the Service struct.
func (inst *Service) DeleteOne(dbName, collectionName string, query *Query) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Delete the document that matches the filter
	_, err := collection.DeleteOne(ctx, query.Filter)
	if err != nil {
		return fmt.Errorf(ErrFailedToDeleteDocument, err)
	}

	return nil
}

// DeleteMany deletes multiple documents from the collection that match the filter in the Query struct.
// It uses the timeout from the Service struct.
func (inst *Service) DeleteMany(dbName, collectionName string, query *Query) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Delete the documents that match the filter
	_, err := collection.DeleteMany(ctx, query.Filter)
	if err != nil {
		return fmt.Errorf(ErrFailedToDeleteDocument, err)
	}

	return nil
}
