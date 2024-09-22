package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateOne updates a single document in the collection that matches the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching document is found.
func (inst *Service) UpdateOne(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filters and update to BSON
	bsonFilter := bson.M(query.Filter)
	bsonUpdate := bson.M(update.Filter)

	// Set upsert option
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the document that matches the filter
	_, err := collection.UpdateOne(inst.ctx, bsonFilter, bsonUpdate, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}

// UpdateMany updates multiple documents in the collection that match the filter and applies the update in the Query struct.
// If upsert is true, it will insert the document if no matching documents are found.
func (inst *Service) UpdateMany(dbName, collectionName string, query *Query, update *Query, upsert bool) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filters and update to BSON
	bsonFilter := bson.M(query.Filter)
	bsonUpdate := bson.M(update.Filter)

	// Set upsert option
	updateOptions := options.Update().SetUpsert(upsert)

	// Update the documents that match the filter
	_, err := collection.UpdateMany(inst.ctx, bsonFilter, bsonUpdate, updateOptions)
	if err != nil {
		return fmt.Errorf(ErrFailedToUpdateDocument, err)
	}

	return nil
}
