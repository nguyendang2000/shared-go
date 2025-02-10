package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Projection defines a structure for MongoDB field projections.
// It allows selecting which fields to include or exclude in query results.
type Projection struct {
	projectionMap map[string]int // Stores field inclusion/exclusion mapping
}

// NewProjection initializes and returns a new Projection instance.
// It creates an empty projection map to manage included/excluded fields.
func NewProjection() *Projection {
	return &Projection{
		projectionMap: make(map[string]int),
	}
}

// Include adds fields to the projection, marking them for inclusion in query results.
// Fields specified here will be assigned a value of 1 in the projection map.
//
// Example:
//
//	proj := NewProjection().Include("name", "email")  // Only includes "name" and "email".
func (inst *Projection) Include(fields ...string) *Projection {
	for _, field := range fields {
		inst.projectionMap[field] = 1 // 1 means the field is included
	}
	return inst
}

// Exclude adds fields to the projection, marking them for exclusion from query results.
// Fields specified here will be assigned a value of 0 in the projection map.
//
// Example:
//
//	proj := NewProjection().Exclude("password", "ssn")  // Excludes "password" and "ssn".
func (inst *Projection) Exclude(fields ...string) *Projection {
	for _, field := range fields {
		inst.projectionMap[field] = 0 // 0 means the field is excluded
	}
	return inst
}

// parse converts the projection map into a BSON document (bson.D).
// This format is used in MongoDB queries to specify field inclusion/exclusion.
//
// Example Output:
//
//	bson.D{{"name", 1}, {"email", 1}}  // Includes "name" and "email"
//	bson.D{{"password", 0}, {"ssn", 0}} // Excludes "password" and "ssn"
func (inst *Projection) parse() bson.D {
	parsed := bson.D{}
	for field, include := range inst.projectionMap {
		parsed = append(parsed, primitive.E{Key: field, Value: include})
	}
	return parsed
}
