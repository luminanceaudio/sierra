// Code generated by ent, DO NOT EDIT.

package sierraent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/collection"
)

// Collection is the model entity for the Collection schema.
type Collection struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CollectionQuery when eager-loading is set.
	Edges        CollectionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CollectionEdges holds the relations/edges for other nodes in the graph.
type CollectionEdges struct {
	// Sample holds the value of the sample edge.
	Sample []*SourceSample `json:"sample,omitempty"`
	// CollectionSamples holds the value of the collection_samples edge.
	CollectionSamples []*CollectionSample `json:"collection_samples,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// SampleOrErr returns the Sample value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) SampleOrErr() ([]*SourceSample, error) {
	if e.loadedTypes[0] {
		return e.Sample, nil
	}
	return nil, &NotLoadedError{edge: "sample"}
}

// CollectionSamplesOrErr returns the CollectionSamples value or an error if the edge
// was not loaded in eager-loading.
func (e CollectionEdges) CollectionSamplesOrErr() ([]*CollectionSample, error) {
	if e.loadedTypes[1] {
		return e.CollectionSamples, nil
	}
	return nil, &NotLoadedError{edge: "collection_samples"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Collection) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case collection.FieldName:
			values[i] = new(sql.NullString)
		case collection.FieldCreateTime:
			values[i] = new(sql.NullTime)
		case collection.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Collection fields.
func (c *Collection) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case collection.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case collection.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case collection.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Collection.
// This includes values selected through modifiers, order, etc.
func (c *Collection) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QuerySample queries the "sample" edge of the Collection entity.
func (c *Collection) QuerySample() *SourceSampleQuery {
	return NewCollectionClient(c.config).QuerySample(c)
}

// QueryCollectionSamples queries the "collection_samples" edge of the Collection entity.
func (c *Collection) QueryCollectionSamples() *CollectionSampleQuery {
	return NewCollectionClient(c.config).QueryCollectionSamples(c)
}

// Update returns a builder for updating this Collection.
// Note that you need to call Collection.Unwrap() before calling this method if this Collection
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Collection) Update() *CollectionUpdateOne {
	return NewCollectionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Collection entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Collection) Unwrap() *Collection {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("sierraent: Collection is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Collection) String() string {
	var builder strings.Builder
	builder.WriteString("Collection(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Collections is a parsable slice of Collection.
type Collections []*Collection