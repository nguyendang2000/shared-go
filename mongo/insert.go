package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne inserts a single document into the collection.
// It uses the timeout defined in the Service struct to create a context for the operation.
func (inst *Service) InsertOne(dbName, collectionName string, document interface{}) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database.
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Insert the document into the collection.
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}

// InsertMany inserts multiple documents into the collection using variadic arguments.
// It uses the timeout defined in the Service struct to create a context for the operation.
func (inst *Service) InsertMany(dbName, collectionName string, documents ...interface{}) error {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database.
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Insert the documents into the collection.
	_, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}

// InsertManyUnordered inserts multiple documents into the specified collection in an unordered manner.
// MongoDB will attempt to insert all documents even if some fail due to errors like duplicate keys.
// It uses the timeout defined in the Service struct to create a context for the operation.
//
// Returns the number of successfully inserted documents and an error only if all insertions fail.
func (inst *Service) InsertManyUnordered(dbName, collectionName string, documents ...interface{}) (int64, error) {
	// Create a context with the specified timeout from the Service struct.
	ctx, cancel := context.WithTimeout(inst.context, time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database.
	collection := inst.client.Database(dbName).Collection(collectionName)

	// Attempt to insert documents with Ordered set to false.
	res, err := collection.InsertMany(ctx, documents, options.InsertMany().SetOrdered(false))

	// Count the successfully inserted documents.
	insertedCount := int64(len(res.InsertedIDs))

	// If some documents were inserted, return count without treating it as a complete failure.
	// Only return an error if *none* of the documents were inserted.
	if insertedCount > 0 {
		return insertedCount, nil
	}

	return 0, fmt.Errorf(ErrFailedToInsertDocument, err)
}
