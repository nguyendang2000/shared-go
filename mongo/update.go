package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne updates a single document in the collection that matches the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching document is found.
// The function uses the timeout defined in the Service struct to create a context for the operation.
func (inst *Service) UpdateOne(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database.
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Set upsert option.
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the document that matches the filter.
	_, err := collection.UpdateOne(ctx, query.Filter, update.Filter, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}

// UpdateMany updates multiple documents in the collection that match the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching documents are found.
// The function uses the timeout defined in the Service struct to create a context for the operation.
func (inst *Service) UpdateMany(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database.
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Set upsert option.
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the documents that match the filter.
	_, err := collection.UpdateMany(ctx, query.Filter, update.Filter, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}
