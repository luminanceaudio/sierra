// Code generated by ent, DO NOT EDIT.

package sierraent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/predicate"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/source"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/sierradb/sierraent/sourcesample"
)

// SourceUpdate is the builder for updating Source entities.
type SourceUpdate struct {
	config
	hooks    []Hook
	mutation *SourceMutation
}

// Where appends a list predicates to the SourceUpdate builder.
func (su *SourceUpdate) Where(ps ...predicate.Source) *SourceUpdate {
	su.mutation.Where(ps...)
	return su
}

// AddSampleIDs adds the "sample" edge to the SourceSample entity by IDs.
func (su *SourceUpdate) AddSampleIDs(ids ...string) *SourceUpdate {
	su.mutation.AddSampleIDs(ids...)
	return su
}

// AddSample adds the "sample" edges to the SourceSample entity.
func (su *SourceUpdate) AddSample(s ...*SourceSample) *SourceUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSampleIDs(ids...)
}

// Mutation returns the SourceMutation object of the builder.
func (su *SourceUpdate) Mutation() *SourceMutation {
	return su.mutation
}

// ClearSample clears all "sample" edges to the SourceSample entity.
func (su *SourceUpdate) ClearSample() *SourceUpdate {
	su.mutation.ClearSample()
	return su
}

// RemoveSampleIDs removes the "sample" edge to SourceSample entities by IDs.
func (su *SourceUpdate) RemoveSampleIDs(ids ...string) *SourceUpdate {
	su.mutation.RemoveSampleIDs(ids...)
	return su
}

// RemoveSample removes "sample" edges to SourceSample entities.
func (su *SourceUpdate) RemoveSample(s ...*SourceSample) *SourceUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSampleIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SourceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SourceUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SourceUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SourceUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SourceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(source.Table, source.Columns, sqlgraph.NewFieldSpec(source.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if su.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSampleIDs(); len(nodes) > 0 && !su.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{source.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SourceUpdateOne is the builder for updating a single Source entity.
type SourceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SourceMutation
}

// AddSampleIDs adds the "sample" edge to the SourceSample entity by IDs.
func (suo *SourceUpdateOne) AddSampleIDs(ids ...string) *SourceUpdateOne {
	suo.mutation.AddSampleIDs(ids...)
	return suo
}

// AddSample adds the "sample" edges to the SourceSample entity.
func (suo *SourceUpdateOne) AddSample(s ...*SourceSample) *SourceUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSampleIDs(ids...)
}

// Mutation returns the SourceMutation object of the builder.
func (suo *SourceUpdateOne) Mutation() *SourceMutation {
	return suo.mutation
}

// ClearSample clears all "sample" edges to the SourceSample entity.
func (suo *SourceUpdateOne) ClearSample() *SourceUpdateOne {
	suo.mutation.ClearSample()
	return suo
}

// RemoveSampleIDs removes the "sample" edge to SourceSample entities by IDs.
func (suo *SourceUpdateOne) RemoveSampleIDs(ids ...string) *SourceUpdateOne {
	suo.mutation.RemoveSampleIDs(ids...)
	return suo
}

// RemoveSample removes "sample" edges to SourceSample entities.
func (suo *SourceUpdateOne) RemoveSample(s ...*SourceSample) *SourceUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSampleIDs(ids...)
}

// Where appends a list predicates to the SourceUpdate builder.
func (suo *SourceUpdateOne) Where(ps ...predicate.Source) *SourceUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SourceUpdateOne) Select(field string, fields ...string) *SourceUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Source entity.
func (suo *SourceUpdateOne) Save(ctx context.Context) (*Source, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SourceUpdateOne) SaveX(ctx context.Context) *Source {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SourceUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SourceUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SourceUpdateOne) sqlSave(ctx context.Context) (_node *Source, err error) {
	_spec := sqlgraph.NewUpdateSpec(source.Table, source.Columns, sqlgraph.NewFieldSpec(source.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`sierraent: missing "Source.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, source.FieldID)
		for _, f := range fields {
			if !source.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("sierraent: invalid field %q for query", f)}
			}
			if f != source.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if suo.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSampleIDs(); len(nodes) > 0 && !suo.mutation.SampleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.SampleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   source.SampleTable,
			Columns: []string{source.SampleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Source{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{source.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
