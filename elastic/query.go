package elastic

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// Query wraps an Elasticsearch query object, providing methods to build complex queries.
// It supports various query types, including term, range, match, and boolean queries.
type Query struct {
	// q holds the underlying Elasticsearch query object.
	q *types.Query
}

// NewQuery initializes a new Query object, setting up an empty query structure.
// This function provides a starting point for constructing queries dynamically.
func NewQuery() *Query {
	return &Query{
		q: &types.Query{},
	}
}

// Match adds a Match query to the Query, matching documents where the specified field contains the given value.
// This is useful for performing full-text searches or fuzzy matches.
func (inst *Query) Match(field string, value string) *Query {
	inst.q.Match = map[string]types.MatchQuery{
		field: {Query: value},
	}
	return inst
}

// MatchAll adds a MatchAll query to the Query, matching all documents in the index.
// This query is useful when retrieving all documents without filtering.
func (inst *Query) MatchAll() *Query {
	inst.q.MatchAll = types.NewMatchAllQuery()
	return inst
}

// Term adds a Term query to the Query, matching documents where the specified field has an exact value.
// This is particularly useful for filtering results based on exact matches, such as IDs or keywords.
func (inst *Query) Term(field string, value interface{}) *Query {
	inst.q.Term = map[string]types.TermQuery{
		field: {Value: value},
	}
	return inst
}

// Range adds a Range query to the Query, matching documents where the specified field falls within the given range.
// The gte (greater than or equal) and lte (less than or equal) parameters define the range boundaries.
func (inst *Query) Range(field string, gte interface{}, lte interface{}) *Query {
	inst.q.Range = map[string]types.RangeQuery{
		field: map[string]interface{}{
			"gte": gte,
			"lte": lte,
		},
	}
	return inst
}

// Gt adds a "greater than" condition to the Range query, matching documents where the specified field is greater than the given value.
func (inst *Query) Gt(field string, value interface{}) *Query {
	if inst.q.Range == nil {
		inst.q.Range = make(map[string]types.RangeQuery)
	}
	if inst.q.Range[field] == nil {
		inst.q.Range[field] = make(map[string]interface{})
	}
	inst.q.Range[field].(map[string]interface{})["gt"] = value
	return inst
}

// Lt adds a "less than" condition to the Range query, matching documents where the specified field is less than the given value.
func (inst *Query) Lt(field string, value interface{}) *Query {
	if inst.q.Range == nil {
		inst.q.Range = make(map[string]types.RangeQuery)
	}
	if inst.q.Range[field] == nil {
		inst.q.Range[field] = make(map[string]interface{})
	}
	inst.q.Range[field].(map[string]interface{})["lt"] = value
	return inst
}

// Gte adds a "greater than or equal" condition to the Range query, matching documents where the specified field is greater than or equal to the given value.
func (inst *Query) Gte(field string, value interface{}) *Query {
	if inst.q.Range == nil {
		inst.q.Range = make(map[string]types.RangeQuery)
	}
	if inst.q.Range[field] == nil {
		inst.q.Range[field] = make(map[string]interface{})
	}
	inst.q.Range[field].(map[string]interface{})["gte"] = value
	return inst
}

// Lte adds a "less than or equal" condition to the Range query, matching documents where the specified field is less than or equal to the given value.
func (inst *Query) Lte(field string, value interface{}) *Query {
	if inst.q.Range == nil {
		inst.q.Range = make(map[string]types.RangeQuery)
	}
	if inst.q.Range[field] == nil {
		inst.q.Range[field] = make(map[string]interface{})
	}
	inst.q.Range[field].(map[string]interface{})["lte"] = value
	return inst
}

// Must adds one or more 'must' conditions to the Bool query, where all conditions must match (logical AND).
// This is useful for ensuring multiple criteria are met in a query.
func (inst *Query) Must(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Must = append(inst.q.Bool.Must, convertQueries(queries)...)
	return inst
}

// Should adds one or more 'should' conditions to the Bool query, where at least one condition should match (logical OR).
// This is useful for boosting relevance when multiple conditions are met.
func (inst *Query) Should(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Should = append(inst.q.Bool.Should, convertQueries(queries)...)
	return inst
}

// MustNot adds one or more 'must not' conditions to the Bool query, excluding documents that match any of the specified conditions (logical NOT).
// This is useful for filtering out specific values from the results.
func (inst *Query) MustNot(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.MustNot = append(inst.q.Bool.MustNot, convertQueries(queries)...)
	return inst
}

// Filter adds one or more 'filter' conditions to the Bool query, ensuring all conditions match without affecting scoring.
// This is useful for filtering results based on criteria such as date ranges or categories.
func (inst *Query) Filter(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Filter = append(inst.q.Bool.Filter, convertQueries(queries)...)
	return inst
}

// convertQueries is a helper function that converts a slice of Query pointers to a slice of Elasticsearch Query objects.
// This function is used when constructing Bool queries that accept multiple conditions.
func convertQueries(queries []*Query) []types.Query {
	result := make([]types.Query, len(queries))
	for i := range queries {
		result[i] = *queries[i].q
	}
	return result
}
