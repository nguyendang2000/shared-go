package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Count returns the number of documents matching the given query.
// It uses the timeout field from the Service struct.
func (inst *Service) Count(dbName, collectionName string, query *Query) (int64, error) {
	// Use the timeout from the Service struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(inst.timeout)*time.Second)
	defer cancel()

	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Convert Query filter to BSON
	bsonFilter := bson.M(query.Filter)

	// Count the number of documents matching the query
	count, err := collection.CountDocuments(ctx, bsonFilter)
	if err != nil {
		return 0, fmt.Errorf(ErrFailedToCountDocuments, err)
	}

	return count, nil
}
