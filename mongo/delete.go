package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteOne deletes a single document from the collection that matches the filter in the Query struct
func (inst *Service) DeleteOne(dbName, collectionName string, query *Query) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Delete the document that matches the filter
	_, err := collection.DeleteOne(inst.ctx, bsonFilter)
	if err != nil {
		return fmt.Errorf(ErrFailedToDeleteDocument, err)
	}

	return nil
}

// DeleteMany deletes multiple documents from the collection that match the filter in the Query struct
func (inst *Service) DeleteMany(dbName, collectionName string, query *Query) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Delete the documents that match the filter
	_, err := collection.DeleteMany(inst.ctx, bsonFilter)
	if err != nil {
		return fmt.Errorf(ErrFailedToDeleteDocument, err)
	}

	return nil
}
