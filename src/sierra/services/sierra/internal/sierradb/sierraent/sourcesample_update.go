// Code generated by ent, DO NOT EDIT.

package sierraent

import (
	"context"
	"errors"
	"fmt"
	"sierra/services/sierra/internal/sierradb/sierraent/predicate"
	"sierra/services/sierra/internal/sierradb/sierraent/sample"
	"sierra/services/sierra/internal/sierradb/sierraent/source"
	"sierra/services/sierra/internal/sierradb/sierraent/sourcesample"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SourceSampleUpdate is the builder for updating SourceSample entities.
type SourceSampleUpdate struct {
	config
	hooks    []Hook
	mutation *SourceSampleMutation
}

// Where appends a list predicates to the SourceSampleUpdate builder.
func (ssu *SourceSampleUpdate) Where(ps ...predicate.SourceSample) *SourceSampleUpdate {
	ssu.mutation.Where(ps...)
	return ssu
}

// SetSourceID sets the "source" edge to the Source entity by ID.
func (ssu *SourceSampleUpdate) SetSourceID(id string) *SourceSampleUpdate {
	ssu.mutation.SetSourceID(id)
	return ssu
}

// SetNillableSourceID sets the "source" edge to the Source entity by ID if the given value is not nil.
func (ssu *SourceSampleUpdate) SetNillableSourceID(id *string) *SourceSampleUpdate {
	if id != nil {
		ssu = ssu.SetSourceID(*id)
	}
	return ssu
}

// SetSource sets the "source" edge to the Source entity.
func (ssu *SourceSampleUpdate) SetSource(s *Source) *SourceSampleUpdate {
	return ssu.SetSourceID(s.ID)
}

// SetSampleID sets the "sample" edge to the Sample entity by ID.
func (ssu *SourceSampleUpdate) SetSampleID(id string) *SourceSampleUpdate {
	ssu.mutation.SetSampleID(id)
	return ssu
}

// SetNillableSampleID sets the "sample" edge to the Sample entity by ID if the given value is not nil.
func (ssu *SourceSampleUpdate) SetNillableSampleID(id *string) *SourceSampleUpdate {
	if id != nil {
		ssu = ssu.SetSampleID(*id)
	}
	return ssu
}

// SetSample sets the "sample" edge to the Sample entity.
func (ssu *SourceSampleUpdate) SetSample(s *Sample) *SourceSampleUpdate {
	return ssu.SetSampleID(s.ID)
}

// Mutation returns the SourceSampleMutation object of the builder.
func (ssu *SourceSampleUpdate) Mutation() *SourceSampleMutation {
	return ssu.mutation
}

// ClearSource clears the "source" edge to the Source entity.
func (ssu *SourceSampleUpdate) ClearSource() *SourceSampleUpdate {
	ssu.mutation.ClearSource()
	return ssu
}

// ClearSample clears the "sample" edge to the Sample entity.
func (ssu *SourceSampleUpdate) ClearSample() *SourceSampleUpdate {
	ssu.mutation.ClearSample()
	return ssu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ssu *SourceSampleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ssu.sqlSave, ssu.mutation, ssu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *SourceSampleUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *SourceSampleUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *SourceSampleUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssu *SourceSampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourcesample.Table, sourcesample.Columns, sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString))
	if ps := ssu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ssu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SourceTable,
			Columns: []string{sourcesample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssu.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SourceTable,
			Columns: []string{sourcesample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ssu.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SampleTable,
			Columns: []string{sourcesample.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssu.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SampleTable,
			Columns: []string{sourcesample.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourcesample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ssu.mutation.done = true
	return n, nil
}

// SourceSampleUpdateOne is the builder for updating a single SourceSample entity.
type SourceSampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SourceSampleMutation
}

// SetSourceID sets the "source" edge to the Source entity by ID.
func (ssuo *SourceSampleUpdateOne) SetSourceID(id string) *SourceSampleUpdateOne {
	ssuo.mutation.SetSourceID(id)
	return ssuo
}

// SetNillableSourceID sets the "source" edge to the Source entity by ID if the given value is not nil.
func (ssuo *SourceSampleUpdateOne) SetNillableSourceID(id *string) *SourceSampleUpdateOne {
	if id != nil {
		ssuo = ssuo.SetSourceID(*id)
	}
	return ssuo
}

// SetSource sets the "source" edge to the Source entity.
func (ssuo *SourceSampleUpdateOne) SetSource(s *Source) *SourceSampleUpdateOne {
	return ssuo.SetSourceID(s.ID)
}

// SetSampleID sets the "sample" edge to the Sample entity by ID.
func (ssuo *SourceSampleUpdateOne) SetSampleID(id string) *SourceSampleUpdateOne {
	ssuo.mutation.SetSampleID(id)
	return ssuo
}

// SetNillableSampleID sets the "sample" edge to the Sample entity by ID if the given value is not nil.
func (ssuo *SourceSampleUpdateOne) SetNillableSampleID(id *string) *SourceSampleUpdateOne {
	if id != nil {
		ssuo = ssuo.SetSampleID(*id)
	}
	return ssuo
}

// SetSample sets the "sample" edge to the Sample entity.
func (ssuo *SourceSampleUpdateOne) SetSample(s *Sample) *SourceSampleUpdateOne {
	return ssuo.SetSampleID(s.ID)
}

// Mutation returns the SourceSampleMutation object of the builder.
func (ssuo *SourceSampleUpdateOne) Mutation() *SourceSampleMutation {
	return ssuo.mutation
}

// ClearSource clears the "source" edge to the Source entity.
func (ssuo *SourceSampleUpdateOne) ClearSource() *SourceSampleUpdateOne {
	ssuo.mutation.ClearSource()
	return ssuo
}

// ClearSample clears the "sample" edge to the Sample entity.
func (ssuo *SourceSampleUpdateOne) ClearSample() *SourceSampleUpdateOne {
	ssuo.mutation.ClearSample()
	return ssuo
}

// Where appends a list predicates to the SourceSampleUpdate builder.
func (ssuo *SourceSampleUpdateOne) Where(ps ...predicate.SourceSample) *SourceSampleUpdateOne {
	ssuo.mutation.Where(ps...)
	return ssuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ssuo *SourceSampleUpdateOne) Select(field string, fields ...string) *SourceSampleUpdateOne {
	ssuo.fields = append([]string{field}, fields...)
	return ssuo
}

// Save executes the query and returns the updated SourceSample entity.
func (ssuo *SourceSampleUpdateOne) Save(ctx context.Context) (*SourceSample, error) {
	return withHooks(ctx, ssuo.sqlSave, ssuo.mutation, ssuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *SourceSampleUpdateOne) SaveX(ctx context.Context) *SourceSample {
	node, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ssuo *SourceSampleUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *SourceSampleUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssuo *SourceSampleUpdateOne) sqlSave(ctx context.Context) (_node *SourceSample, err error) {
	_spec := sqlgraph.NewUpdateSpec(sourcesample.Table, sourcesample.Columns, sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString))
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`sierraent: missing "SourceSample.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ssuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sourcesample.FieldID)
		for _, f := range fields {
			if !sourcesample.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("sierraent: invalid field %q for query", f)}
			}
			if f != sourcesample.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ssuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ssuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SourceTable,
			Columns: []string{sourcesample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssuo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SourceTable,
			Columns: []string{sourcesample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(source.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ssuo.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SampleTable,
			Columns: []string{sourcesample.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ssuo.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   sourcesample.SampleTable,
			Columns: []string{sourcesample.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SourceSample{config: ssuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sourcesample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ssuo.mutation.done = true
	return _node, nil
}
