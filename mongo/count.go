package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

// Count returns the number of documents matching the given query
func (inst *Service) Count(dbName, collectionName string, query *Query) (int64, error) {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Count the number of documents matching the query
	count, err := collection.CountDocuments(inst.ctx, bsonFilter)
	if err != nil {
		return 0, fmt.Errorf(ErrFailedToCountDocuments, err)
	}

	return count, nil
}
