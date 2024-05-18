// Code generated by ent, DO NOT EDIT.

package sierraent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sample"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/source"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sourcesample"
)

// SourceSample is the model entity for the SourceSample schema.
type SourceSample struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// RelativePath holds the value of the "relative_path" field.
	RelativePath string `json:"relative_path,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SourceSampleQuery when eager-loading is set.
	Edges        SourceSampleEdges `json:"edges"`
	source       *string
	sample       *string
	selectValues sql.SelectValues
}

// SourceSampleEdges holds the relations/edges for other nodes in the graph.
type SourceSampleEdges struct {
	// Source holds the value of the source edge.
	Source *Source `json:"source,omitempty"`
	// Sample holds the value of the sample edge.
	Sample *Sample `json:"sample,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SourceOrErr returns the Source value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceSampleEdges) SourceOrErr() (*Source, error) {
	if e.Source != nil {
		return e.Source, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: source.Label}
	}
	return nil, &NotLoadedError{edge: "source"}
}

// SampleOrErr returns the Sample value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SourceSampleEdges) SampleOrErr() (*Sample, error) {
	if e.Sample != nil {
		return e.Sample, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: sample.Label}
	}
	return nil, &NotLoadedError{edge: "sample"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SourceSample) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sourcesample.FieldID, sourcesample.FieldRelativePath, sourcesample.FieldFilename:
			values[i] = new(sql.NullString)
		case sourcesample.ForeignKeys[0]: // source
			values[i] = new(sql.NullString)
		case sourcesample.ForeignKeys[1]: // sample
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SourceSample fields.
func (ss *SourceSample) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sourcesample.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				ss.ID = value.String
			}
		case sourcesample.FieldRelativePath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field relative_path", values[i])
			} else if value.Valid {
				ss.RelativePath = value.String
			}
		case sourcesample.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				ss.Filename = value.String
			}
		case sourcesample.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				ss.source = new(string)
				*ss.source = value.String
			}
		case sourcesample.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sample", values[i])
			} else if value.Valid {
				ss.sample = new(string)
				*ss.sample = value.String
			}
		default:
			ss.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SourceSample.
// This includes values selected through modifiers, order, etc.
func (ss *SourceSample) Value(name string) (ent.Value, error) {
	return ss.selectValues.Get(name)
}

// QuerySource queries the "source" edge of the SourceSample entity.
func (ss *SourceSample) QuerySource() *SourceQuery {
	return NewSourceSampleClient(ss.config).QuerySource(ss)
}

// QuerySample queries the "sample" edge of the SourceSample entity.
func (ss *SourceSample) QuerySample() *SampleQuery {
	return NewSourceSampleClient(ss.config).QuerySample(ss)
}

// Update returns a builder for updating this SourceSample.
// Note that you need to call SourceSample.Unwrap() before calling this method if this SourceSample
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SourceSample) Update() *SourceSampleUpdateOne {
	return NewSourceSampleClient(ss.config).UpdateOne(ss)
}

// Unwrap unwraps the SourceSample entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ss *SourceSample) Unwrap() *SourceSample {
	_tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("sierraent: SourceSample is not a transactional entity")
	}
	ss.config.driver = _tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SourceSample) String() string {
	var builder strings.Builder
	builder.WriteString("SourceSample(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ss.ID))
	builder.WriteString("relative_path=")
	builder.WriteString(ss.RelativePath)
	builder.WriteString(", ")
	builder.WriteString("filename=")
	builder.WriteString(ss.Filename)
	builder.WriteByte(')')
	return builder.String()
}

// SourceSamples is a parsable slice of SourceSample.
type SourceSamples []*SourceSample
