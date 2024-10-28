package elastic

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// Query wraps an Elasticsearch query object, providing methods to build complex queries.
type Query struct {
	q *types.Query
}

// NewQuery initializes a new Query object, setting up an empty query structure.
func NewQuery() *Query {
	return &Query{
		q: &types.Query{},
	}
}

// Match adds a Match query to the Query, matching documents where the specified field contains the given value.
// This is useful for finding documents with similar text.
func (inst *Query) Match(field string, value string) *Query {
	inst.q.Match = map[string]types.MatchQuery{
		field: {Query: value},
	}
	return inst
}

// MatchAll adds a MatchAll query to the Query, matching all documents in the index.
func (inst *Query) MatchAll() *Query {
	inst.q.MatchAll = types.NewMatchAllQuery()
	return inst
}

// Term adds a Term query to the Query, matching documents where the specified field has an exact value.
// Useful for exact matches on fields like keywords, IDs, etc.
func (inst *Query) Term(field string, value interface{}) *Query {
	inst.q.Term = map[string]types.TermQuery{
		field: {Value: value},
	}
	return inst
}

// Range adds a Range query to the Query, matching documents where the specified field has values within a range.
// Parameters gte (greater than or equal) and lte (less than or equal) specify the range boundaries.
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

// Must adds one or more 'must' conditions to the Bool query, where all conditions must match (AND).
func (inst *Query) Must(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Must = append(inst.q.Bool.Must, convertQueries(queries)...)
	return inst
}

// Should adds one or more 'should' conditions to the Bool query, where at least one condition should match (OR).
func (inst *Query) Should(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Should = append(inst.q.Bool.Should, convertQueries(queries)...)
	return inst
}

// MustNot adds one or more 'must not' conditions to the Bool query, where documents matching any condition will be excluded (NOT).
func (inst *Query) MustNot(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.MustNot = append(inst.q.Bool.MustNot, convertQueries(queries)...)
	return inst
}

// Filter adds one or more 'filter' conditions to the Bool query, where all conditions must match but do not affect scoring.
func (inst *Query) Filter(queries ...*Query) *Query {
	if inst.q.Bool == nil {
		inst.q.Bool = types.NewBoolQuery()
	}

	inst.q.Bool.Filter = append(inst.q.Bool.Filter, convertQueries(queries)...)
	return inst
}

// Helper function to convert variadic []*Query to []types.Query
func convertQueries(queries []*Query) []types.Query {
	result := make([]types.Query, len(queries))
	for i := range queries {
		result[i] = *queries[i].q
	}
	return result
}
