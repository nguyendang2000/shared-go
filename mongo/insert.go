package mongo

import "fmt"

// InsertOne inserts a single document into the collection
func (inst *Service) InsertOne(dbName, collectionName string, document interface{}) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Insert the document into the collection
	_, err := collection.InsertOne(inst.ctx, document)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}

// InsertMany inserts multiple documents into the collection using variadic arguments
func (inst *Service) InsertMany(dbName, collectionName string, documents ...interface{}) error {
	// Get the collection from the specified database
	collection := inst.Database(dbName).Collection(collectionName)

	// Insert the documents into the collection
	_, err := collection.InsertMany(inst.ctx, documents)
	if err != nil {
		return fmt.Errorf(ErrFailedToInsertDocument, err)
	}

	return nil
}
