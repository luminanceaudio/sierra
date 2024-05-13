// Code generated by ent, DO NOT EDIT.

package sample

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the sample type in the database.
	Label = "sample"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "sha256"
	// FieldFormat holds the string denoting the format field in the database.
	FieldFormat = "format"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldWaveformStoragePath holds the string denoting the waveform_storage_path field in the database.
	FieldWaveformStoragePath = "waveform_storage_path"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// SourceSampleFieldID holds the string denoting the ID field of the SourceSample.
	SourceSampleFieldID = "uri"
	// Table holds the table name of the sample in the database.
	Table = "samples"
	// SourceTable is the table that holds the source relation/edge.
	SourceTable = "source_samples"
	// SourceInverseTable is the table name for the SourceSample entity.
	// It exists in this package in order to avoid circular dependency with the "sourcesample" package.
	SourceInverseTable = "source_samples"
	// SourceColumn is the table column denoting the source relation/edge.
	SourceColumn = "sample"
)

// Columns holds all SQL columns for sample fields.
var Columns = []string{
	FieldID,
	FieldFormat,
	FieldLength,
	FieldWaveformStoragePath,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Sample queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFormat orders the results by the format field.
func ByFormat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFormat, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

// ByWaveformStoragePath orders the results by the waveform_storage_path field.
func ByWaveformStoragePath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWaveformStoragePath, opts...).ToFunc()
}

// BySourceCount orders the results by source count.
func BySourceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSourceStep(), opts...)
	}
}

// BySource orders the results by source terms.
func BySource(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceInverseTable, SourceSampleFieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SourceTable, SourceColumn),
	)
}
