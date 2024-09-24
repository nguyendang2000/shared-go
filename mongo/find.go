package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindOne allows the user to specify search criteria and unmarshal the result into the specified struct.
// It uses the timeout from the Service struct.
func (inst *Service) FindOne(dbName, collectionName string, query *Query, result interface{}) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Execute FindOne
	err := collection.FindOne(ctx, query.Filter).Decode(result)
	if err != nil {
		// Return ErrDocumentNotFound (the alias) if no documents are found
		if err == mongo.ErrNoDocuments {
			return ErrDocumentNotFound
		}
		// Return other errors as is
		return fmt.Errorf(ErrFailedToFindOne, err)
	}

	// Return nil when the document is successfully found and decoded
	return nil
}

// FindMany allows the user to specify search criteria, number of items, and unmarshal the results into the specified struct.
// It uses the timeout from the Service struct.
func (inst *Service) FindMany(dbName, collectionName string, query *Query, limit int64, offset int64, result interface{}) error {
	// Create a context with the specified timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Set options for the query, including limit and offset
	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(limit)
	}
	if offset > 0 {
		findOptions.SetSkip(offset)
	}

	// Execute the query using Find for multiple results
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

// Exists checks if a document matching the provided filter exists in the collection.
func (inst *Service) Exists(dbName, collectionName string, query *Query) (bool, error) {
	var result bson.M // We don't care about the result, just want to check if it exists

	// Use the modified FindOne method
	err := inst.FindOne(dbName, collectionName, query, &result)

	if err != nil {
		if errors.Is(err, ErrDocumentNotFound) {
			// No document found, return false without error
			return false, nil
		}
		// Return false and the error if there's another issue
		return false, fmt.Errorf(ErrFailedToCheckExistence, err)
	}

	// Document found, return true
	return true, nil
}
