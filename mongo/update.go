package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne updates a single document in the collection that matches the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching document is found.
// It uses the timeout from the Service struct.
func (inst *Service) UpdateOne(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filters and update to BSON
	bsonFilter := bson.M(query.Filter)
	bsonUpdate := bson.M(update.Filter)

	// Set upsert option
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the document that matches the filter
	_, err := collection.UpdateOne(ctx, bsonFilter, bsonUpdate, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}

// UpdateMany updates multiple documents in the collection that match the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching documents are found.
// It uses the timeout from the Service struct.
func (inst *Service) UpdateMany(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filters and update to BSON
	bsonFilter := bson.M(query.Filter)
	bsonUpdate := bson.M(update.Filter)

	// Set upsert option
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the documents that match the filter
	_, err := collection.UpdateMany(ctx, bsonFilter, bsonUpdate, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}
