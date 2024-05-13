// Code generated by ent, DO NOT EDIT.

package sierraent

import (
	"context"
	"errors"
	"fmt"
	"sierra/services/sierra/internal/sierradb/sierraent/predicate"
	"sierra/services/sierra/internal/sierradb/sierraent/sample"
	"sierra/services/sierra/internal/sierradb/sierraent/sourcesample"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SampleUpdate is the builder for updating Sample entities.
type SampleUpdate struct {
	config
	hooks    []Hook
	mutation *SampleMutation
}

// Where appends a list predicates to the SampleUpdate builder.
func (su *SampleUpdate) Where(ps ...predicate.Sample) *SampleUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetFormat sets the "format" field.
func (su *SampleUpdate) SetFormat(s string) *SampleUpdate {
	su.mutation.SetFormat(s)
	return su
}

// SetNillableFormat sets the "format" field if the given value is not nil.
func (su *SampleUpdate) SetNillableFormat(s *string) *SampleUpdate {
	if s != nil {
		su.SetFormat(*s)
	}
	return su
}

// ClearFormat clears the value of the "format" field.
func (su *SampleUpdate) ClearFormat() *SampleUpdate {
	su.mutation.ClearFormat()
	return su
}

// SetLength sets the "length" field.
func (su *SampleUpdate) SetLength(i int64) *SampleUpdate {
	su.mutation.ResetLength()
	su.mutation.SetLength(i)
	return su
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (su *SampleUpdate) SetNillableLength(i *int64) *SampleUpdate {
	if i != nil {
		su.SetLength(*i)
	}
	return su
}

// AddLength adds i to the "length" field.
func (su *SampleUpdate) AddLength(i int64) *SampleUpdate {
	su.mutation.AddLength(i)
	return su
}

// ClearLength clears the value of the "length" field.
func (su *SampleUpdate) ClearLength() *SampleUpdate {
	su.mutation.ClearLength()
	return su
}

// SetWaveformStoragePath sets the "waveform_storage_path" field.
func (su *SampleUpdate) SetWaveformStoragePath(s string) *SampleUpdate {
	su.mutation.SetWaveformStoragePath(s)
	return su
}

// SetNillableWaveformStoragePath sets the "waveform_storage_path" field if the given value is not nil.
func (su *SampleUpdate) SetNillableWaveformStoragePath(s *string) *SampleUpdate {
	if s != nil {
		su.SetWaveformStoragePath(*s)
	}
	return su
}

// ClearWaveformStoragePath clears the value of the "waveform_storage_path" field.
func (su *SampleUpdate) ClearWaveformStoragePath() *SampleUpdate {
	su.mutation.ClearWaveformStoragePath()
	return su
}

// AddSourceIDs adds the "source" edge to the SourceSample entity by IDs.
func (su *SampleUpdate) AddSourceIDs(ids ...string) *SampleUpdate {
	su.mutation.AddSourceIDs(ids...)
	return su
}

// AddSource adds the "source" edges to the SourceSample entity.
func (su *SampleUpdate) AddSource(s ...*SourceSample) *SampleUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddSourceIDs(ids...)
}

// Mutation returns the SampleMutation object of the builder.
func (su *SampleUpdate) Mutation() *SampleMutation {
	return su.mutation
}

// ClearSource clears all "source" edges to the SourceSample entity.
func (su *SampleUpdate) ClearSource() *SampleUpdate {
	su.mutation.ClearSource()
	return su
}

// RemoveSourceIDs removes the "source" edge to SourceSample entities by IDs.
func (su *SampleUpdate) RemoveSourceIDs(ids ...string) *SampleUpdate {
	su.mutation.RemoveSourceIDs(ids...)
	return su
}

// RemoveSource removes "source" edges to SourceSample entities.
func (su *SampleUpdate) RemoveSource(s ...*SourceSample) *SampleUpdate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveSourceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SampleUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SampleUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SampleUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SampleUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SampleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(sample.Table, sample.Columns, sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Format(); ok {
		_spec.SetField(sample.FieldFormat, field.TypeString, value)
	}
	if su.mutation.FormatCleared() {
		_spec.ClearField(sample.FieldFormat, field.TypeString)
	}
	if value, ok := su.mutation.Length(); ok {
		_spec.SetField(sample.FieldLength, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedLength(); ok {
		_spec.AddField(sample.FieldLength, field.TypeInt64, value)
	}
	if su.mutation.LengthCleared() {
		_spec.ClearField(sample.FieldLength, field.TypeInt64)
	}
	if value, ok := su.mutation.WaveformStoragePath(); ok {
		_spec.SetField(sample.FieldWaveformStoragePath, field.TypeString, value)
	}
	if su.mutation.WaveformStoragePathCleared() {
		_spec.ClearField(sample.FieldWaveformStoragePath, field.TypeString)
	}
	if su.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedSourceIDs(); len(nodes) > 0 && !su.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
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
	if nodes := su.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
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
			err = &NotFoundError{sample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SampleUpdateOne is the builder for updating a single Sample entity.
type SampleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SampleMutation
}

// SetFormat sets the "format" field.
func (suo *SampleUpdateOne) SetFormat(s string) *SampleUpdateOne {
	suo.mutation.SetFormat(s)
	return suo
}

// SetNillableFormat sets the "format" field if the given value is not nil.
func (suo *SampleUpdateOne) SetNillableFormat(s *string) *SampleUpdateOne {
	if s != nil {
		suo.SetFormat(*s)
	}
	return suo
}

// ClearFormat clears the value of the "format" field.
func (suo *SampleUpdateOne) ClearFormat() *SampleUpdateOne {
	suo.mutation.ClearFormat()
	return suo
}

// SetLength sets the "length" field.
func (suo *SampleUpdateOne) SetLength(i int64) *SampleUpdateOne {
	suo.mutation.ResetLength()
	suo.mutation.SetLength(i)
	return suo
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (suo *SampleUpdateOne) SetNillableLength(i *int64) *SampleUpdateOne {
	if i != nil {
		suo.SetLength(*i)
	}
	return suo
}

// AddLength adds i to the "length" field.
func (suo *SampleUpdateOne) AddLength(i int64) *SampleUpdateOne {
	suo.mutation.AddLength(i)
	return suo
}

// ClearLength clears the value of the "length" field.
func (suo *SampleUpdateOne) ClearLength() *SampleUpdateOne {
	suo.mutation.ClearLength()
	return suo
}

// SetWaveformStoragePath sets the "waveform_storage_path" field.
func (suo *SampleUpdateOne) SetWaveformStoragePath(s string) *SampleUpdateOne {
	suo.mutation.SetWaveformStoragePath(s)
	return suo
}

// SetNillableWaveformStoragePath sets the "waveform_storage_path" field if the given value is not nil.
func (suo *SampleUpdateOne) SetNillableWaveformStoragePath(s *string) *SampleUpdateOne {
	if s != nil {
		suo.SetWaveformStoragePath(*s)
	}
	return suo
}

// ClearWaveformStoragePath clears the value of the "waveform_storage_path" field.
func (suo *SampleUpdateOne) ClearWaveformStoragePath() *SampleUpdateOne {
	suo.mutation.ClearWaveformStoragePath()
	return suo
}

// AddSourceIDs adds the "source" edge to the SourceSample entity by IDs.
func (suo *SampleUpdateOne) AddSourceIDs(ids ...string) *SampleUpdateOne {
	suo.mutation.AddSourceIDs(ids...)
	return suo
}

// AddSource adds the "source" edges to the SourceSample entity.
func (suo *SampleUpdateOne) AddSource(s ...*SourceSample) *SampleUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddSourceIDs(ids...)
}

// Mutation returns the SampleMutation object of the builder.
func (suo *SampleUpdateOne) Mutation() *SampleMutation {
	return suo.mutation
}

// ClearSource clears all "source" edges to the SourceSample entity.
func (suo *SampleUpdateOne) ClearSource() *SampleUpdateOne {
	suo.mutation.ClearSource()
	return suo
}

// RemoveSourceIDs removes the "source" edge to SourceSample entities by IDs.
func (suo *SampleUpdateOne) RemoveSourceIDs(ids ...string) *SampleUpdateOne {
	suo.mutation.RemoveSourceIDs(ids...)
	return suo
}

// RemoveSource removes "source" edges to SourceSample entities.
func (suo *SampleUpdateOne) RemoveSource(s ...*SourceSample) *SampleUpdateOne {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveSourceIDs(ids...)
}

// Where appends a list predicates to the SampleUpdate builder.
func (suo *SampleUpdateOne) Where(ps ...predicate.Sample) *SampleUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SampleUpdateOne) Select(field string, fields ...string) *SampleUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Sample entity.
func (suo *SampleUpdateOne) Save(ctx context.Context) (*Sample, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SampleUpdateOne) SaveX(ctx context.Context) *Sample {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SampleUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SampleUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SampleUpdateOne) sqlSave(ctx context.Context) (_node *Sample, err error) {
	_spec := sqlgraph.NewUpdateSpec(sample.Table, sample.Columns, sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`sierraent: missing "Sample.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sample.FieldID)
		for _, f := range fields {
			if !sample.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("sierraent: invalid field %q for query", f)}
			}
			if f != sample.FieldID {
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
	if value, ok := suo.mutation.Format(); ok {
		_spec.SetField(sample.FieldFormat, field.TypeString, value)
	}
	if suo.mutation.FormatCleared() {
		_spec.ClearField(sample.FieldFormat, field.TypeString)
	}
	if value, ok := suo.mutation.Length(); ok {
		_spec.SetField(sample.FieldLength, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedLength(); ok {
		_spec.AddField(sample.FieldLength, field.TypeInt64, value)
	}
	if suo.mutation.LengthCleared() {
		_spec.ClearField(sample.FieldLength, field.TypeInt64)
	}
	if value, ok := suo.mutation.WaveformStoragePath(); ok {
		_spec.SetField(sample.FieldWaveformStoragePath, field.TypeString, value)
	}
	if suo.mutation.WaveformStoragePathCleared() {
		_spec.ClearField(sample.FieldWaveformStoragePath, field.TypeString)
	}
	if suo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sourcesample.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedSourceIDs(); len(nodes) > 0 && !suo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
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
	if nodes := suo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   sample.SourceTable,
			Columns: []string{sample.SourceColumn},
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
	_node = &Sample{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sample.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
