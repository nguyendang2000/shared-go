package mongo

import (
	"context"
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

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Execute FindOne
	err := collection.FindOne(ctx, bsonFilter).Decode(result)
	if err != nil && err != mongo.ErrNoDocuments {
		return fmt.Errorf(ErrFailedToFindOne, err)
	}

	// No documents found is not considered an error, return nil
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

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Set options for the query, including limit and offset
	findOptions := options.Find()
	if limit > 0 {
		findOptions.SetLimit(limit)
	}
	if offset > 0 {
		findOptions.SetSkip(offset)
	}

	// Execute the query using Find for multiple results
	cursor, err := collection.Find(ctx, bsonFilter, findOptions)
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
