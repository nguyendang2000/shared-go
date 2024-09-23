package mongo

import (
	"context"
	"fmt"
	"time"
)

// InsertOne inserts a single document into the collection.
// It uses the timeout from the Service struct.
func (inst *Service) InsertOne(dbName, collectionName string, document interface{}) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Insert the document into the collection
	_, err := collection.InsertOne(ctx, document)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}

// InsertMany inserts multiple documents into the collection using variadic arguments.
// It uses the timeout from the Service struct.
func (inst *Service) InsertMany(dbName, collectionName string, documents ...interface{}) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Insert the documents into the collection
	_, err := collection.InsertMany(ctx, documents)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}
