package mongo

import "go.mongodb.org/mongo-driver/bson"

// Query is a wrapper around bson.M to help build MongoDB query filters.
type Query struct {
	// Filter represents the filter used to query documents.
	Filter bson.M
}

// NewQuery initializes and returns a new Query with an empty filter.
func NewQuery() *Query {
	return &Query{
		Filter: bson.M{}, // Initialize an empty bson.M.
	}
}

// Field adds a key-value pair to the Query filter for equality matching.
func (q *Query) Field(key string, value interface{}) *Query {
	q.Filter[key] = value
	return q
}

// In adds an $in operator to the Query filter for matching any of the provided values.
func (q *Query) In(key string, values ...interface{}) *Query {
	q.Filter[key] = bson.M{"$in": values}
	return q
}

// NotIn adds a $nin operator to the Query filter for excluding specified values.
func (q *Query) NotIn(key string, values ...interface{}) *Query {
	q.Filter[key] = bson.M{"$nin": values}
	return q
}

// GreaterThan adds a $gt operator to the Query filter for matching values greater than the provided value.
func (q *Query) GreaterThan(key string, value interface{}) *Query {
	q.Filter[key] = bson.M{"$gt": value}
	return q
}

// LessThan adds a $lt operator to the Query filter for matching values less than the provided value.
func (q *Query) LessThan(key string, value interface{}) *Query {
	q.Filter[key] = bson.M{"$lt": value}
	return q
}

// GreaterThanOrEqual adds a $gte operator to the Query filter for matching values greater than or equal to the provided value.
func (q *Query) GreaterThanOrEqual(key string, value interface{}) *Query {
	q.Filter[key] = bson.M{"$gte": value}
	return q
}

// LessThanOrEqual adds a $lte operator to the Query filter for matching values less than or equal to the provided value.
func (q *Query) LessThanOrEqual(key string, value interface{}) *Query {
	q.Filter[key] = bson.M{"$lte": value}
	return q
}

// Or adds an $or operator to the Query filter with multiple conditions.
func (q *Query) Or(queries ...*Query) *Query {
	conditions := make([]bson.M, len(queries))
	for i, query := range queries {
		conditions[i] = query.Filter
	}
	q.Filter["$or"] = conditions
	return q
}

// And adds an $and operator to the Query filter with multiple conditions.
func (q *Query) And(queries ...*Query) *Query {
	conditions := make([]bson.M, len(queries))
	for i, query := range queries {
		conditions[i] = query.Filter
	}
	q.Filter["$and"] = conditions
	return q
}

// Exists adds an $exists operator to the Query filter to check for the existence of a field.
func (q *Query) Exists(key string, exists bool) *Query {
	q.Filter[key] = bson.M{"$exists": exists}
	return q
}

// Ne adds a $ne (Not Equal) operator to the Query filter.
func (q *Query) Ne(key string, value interface{}) *Query {
	q.Filter[key] = bson.M{"$ne": value}
	return q
}

// Regex adds a $regex (Regular Expression) operator to the Query filter for pattern matching.
func (q *Query) Regex(key string, pattern string) *Query {
	q.Filter[key] = bson.M{"$regex": pattern}
	return q
}

// ElemMatch adds an $elemMatch operator to the Query filter to match elements in an array that satisfy the given Query conditions.
func (q *Query) ElemMatch(key string, match *Query) *Query {
	q.Filter[key] = bson.M{"$elemMatch": match.Filter}
	return q
}

// All adds an $all operator to the Query filter for matching arrays that contain all the provided values.
func (q *Query) All(key string, values ...interface{}) *Query {
	q.Filter[key] = bson.M{"$all": values}
	return q
}

// Set adds a $set operator to the Query filter for setting a field to a specific value.
func (q *Query) Set(key string, value interface{}) *Query {
	if existing, ok := q.Filter["$set"]; ok {
		// If $set already exists, merge the new key-value into the existing map.
		existingMap := existing.(bson.M)
		existingMap[key] = value
	} else {
		// Otherwise, create a new $set map.
		q.Filter["$set"] = bson.M{key: value}
	}
	return q
}

// Incr adds an $inc operator to the Query filter for incrementing a field's value.
func (q *Query) Incr(key string, value interface{}) *Query {
	if existing, ok := q.Filter["$inc"]; ok {
		// If $inc already exists, merge the new key-value into the existing map.
		existingMap := existing.(bson.M)
		existingMap[key] = value
	} else {
		// Otherwise, create a new $inc map.
		q.Filter["$inc"] = bson.M{key: value}
	}
	return q
}
