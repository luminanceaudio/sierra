// Code generated by ent, DO NOT EDIT.

package sourcesample

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sourcesample type in the database.
	Label = "source_sample"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uri"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// EdgeSample holds the string denoting the sample edge name in mutations.
	EdgeSample = "sample"
	// SampleFieldID holds the string denoting the ID field of the Sample.
	SampleFieldID = "sha256"
	// Table holds the table name of the sourcesample in the database.
	Table = "source_samples"
	// SourceTable is the table that holds the source relation/edge.
	SourceTable = "source_samples"
	// SourceInverseTable is the table name for the Source entity.
	// It exists in this package in order to avoid circular dependency with the "source" package.
	SourceInverseTable = "sources"
	// SourceColumn is the table column denoting the source relation/edge.
	SourceColumn = "source"
	// SampleTable is the table that holds the sample relation/edge.
	SampleTable = "source_samples"
	// SampleInverseTable is the table name for the Sample entity.
	// It exists in this package in order to avoid circular dependency with the "sample" package.
	SampleInverseTable = "samples"
	// SampleColumn is the table column denoting the sample relation/edge.
	SampleColumn = "sample"
)

// Columns holds all SQL columns for sourcesample fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "source_samples"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"source",
	"sample",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the SourceSample queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySourceField orders the results by source field.
func BySourceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceStep(), sql.OrderByField(field, opts...))
	}
}

// BySampleField orders the results by sample field.
func BySampleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSampleStep(), sql.OrderByField(field, opts...))
	}
}
func newSourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SourceTable, SourceColumn),
	)
}
func newSampleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SampleInverseTable, SampleFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SampleTable, SampleColumn),
	)
}
