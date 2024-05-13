// Code generated by ent, DO NOT EDIT.

package sample

import (
	"sierra/services/sierra/internal/sierradb/sierraent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Sample {
	return predicate.Sample(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Sample {
	return predicate.Sample(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Sample {
	return predicate.Sample(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Sample {
	return predicate.Sample(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Sample {
	return predicate.Sample(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Sample {
	return predicate.Sample(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Sample {
	return predicate.Sample(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Sample {
	return predicate.Sample(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Sample {
	return predicate.Sample(sql.FieldContainsFold(FieldID, id))
}

// Format applies equality check predicate on the "format" field. It's identical to FormatEQ.
func Format(v string) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldFormat, v))
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldLength, v))
}

// WaveformSvg applies equality check predicate on the "waveform_svg" field. It's identical to WaveformSvgEQ.
func WaveformSvg(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldWaveformSvg, v))
}

// FormatEQ applies the EQ predicate on the "format" field.
func FormatEQ(v string) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldFormat, v))
}

// FormatNEQ applies the NEQ predicate on the "format" field.
func FormatNEQ(v string) predicate.Sample {
	return predicate.Sample(sql.FieldNEQ(FieldFormat, v))
}

// FormatIn applies the In predicate on the "format" field.
func FormatIn(vs ...string) predicate.Sample {
	return predicate.Sample(sql.FieldIn(FieldFormat, vs...))
}

// FormatNotIn applies the NotIn predicate on the "format" field.
func FormatNotIn(vs ...string) predicate.Sample {
	return predicate.Sample(sql.FieldNotIn(FieldFormat, vs...))
}

// FormatGT applies the GT predicate on the "format" field.
func FormatGT(v string) predicate.Sample {
	return predicate.Sample(sql.FieldGT(FieldFormat, v))
}

// FormatGTE applies the GTE predicate on the "format" field.
func FormatGTE(v string) predicate.Sample {
	return predicate.Sample(sql.FieldGTE(FieldFormat, v))
}

// FormatLT applies the LT predicate on the "format" field.
func FormatLT(v string) predicate.Sample {
	return predicate.Sample(sql.FieldLT(FieldFormat, v))
}

// FormatLTE applies the LTE predicate on the "format" field.
func FormatLTE(v string) predicate.Sample {
	return predicate.Sample(sql.FieldLTE(FieldFormat, v))
}

// FormatContains applies the Contains predicate on the "format" field.
func FormatContains(v string) predicate.Sample {
	return predicate.Sample(sql.FieldContains(FieldFormat, v))
}

// FormatHasPrefix applies the HasPrefix predicate on the "format" field.
func FormatHasPrefix(v string) predicate.Sample {
	return predicate.Sample(sql.FieldHasPrefix(FieldFormat, v))
}

// FormatHasSuffix applies the HasSuffix predicate on the "format" field.
func FormatHasSuffix(v string) predicate.Sample {
	return predicate.Sample(sql.FieldHasSuffix(FieldFormat, v))
}

// FormatIsNil applies the IsNil predicate on the "format" field.
func FormatIsNil() predicate.Sample {
	return predicate.Sample(sql.FieldIsNull(FieldFormat))
}

// FormatNotNil applies the NotNil predicate on the "format" field.
func FormatNotNil() predicate.Sample {
	return predicate.Sample(sql.FieldNotNull(FieldFormat))
}

// FormatEqualFold applies the EqualFold predicate on the "format" field.
func FormatEqualFold(v string) predicate.Sample {
	return predicate.Sample(sql.FieldEqualFold(FieldFormat, v))
}

// FormatContainsFold applies the ContainsFold predicate on the "format" field.
func FormatContainsFold(v string) predicate.Sample {
	return predicate.Sample(sql.FieldContainsFold(FieldFormat, v))
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldLength, v))
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldNEQ(FieldLength, v))
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...int64) predicate.Sample {
	return predicate.Sample(sql.FieldIn(FieldLength, vs...))
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...int64) predicate.Sample {
	return predicate.Sample(sql.FieldNotIn(FieldLength, vs...))
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldGT(FieldLength, v))
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldGTE(FieldLength, v))
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldLT(FieldLength, v))
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v int64) predicate.Sample {
	return predicate.Sample(sql.FieldLTE(FieldLength, v))
}

// LengthIsNil applies the IsNil predicate on the "length" field.
func LengthIsNil() predicate.Sample {
	return predicate.Sample(sql.FieldIsNull(FieldLength))
}

// LengthNotNil applies the NotNil predicate on the "length" field.
func LengthNotNil() predicate.Sample {
	return predicate.Sample(sql.FieldNotNull(FieldLength))
}

// WaveformSvgEQ applies the EQ predicate on the "waveform_svg" field.
func WaveformSvgEQ(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldEQ(FieldWaveformSvg, v))
}

// WaveformSvgNEQ applies the NEQ predicate on the "waveform_svg" field.
func WaveformSvgNEQ(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldNEQ(FieldWaveformSvg, v))
}

// WaveformSvgIn applies the In predicate on the "waveform_svg" field.
func WaveformSvgIn(vs ...[]byte) predicate.Sample {
	return predicate.Sample(sql.FieldIn(FieldWaveformSvg, vs...))
}

// WaveformSvgNotIn applies the NotIn predicate on the "waveform_svg" field.
func WaveformSvgNotIn(vs ...[]byte) predicate.Sample {
	return predicate.Sample(sql.FieldNotIn(FieldWaveformSvg, vs...))
}

// WaveformSvgGT applies the GT predicate on the "waveform_svg" field.
func WaveformSvgGT(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldGT(FieldWaveformSvg, v))
}

// WaveformSvgGTE applies the GTE predicate on the "waveform_svg" field.
func WaveformSvgGTE(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldGTE(FieldWaveformSvg, v))
}

// WaveformSvgLT applies the LT predicate on the "waveform_svg" field.
func WaveformSvgLT(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldLT(FieldWaveformSvg, v))
}

// WaveformSvgLTE applies the LTE predicate on the "waveform_svg" field.
func WaveformSvgLTE(v []byte) predicate.Sample {
	return predicate.Sample(sql.FieldLTE(FieldWaveformSvg, v))
}

// WaveformSvgIsNil applies the IsNil predicate on the "waveform_svg" field.
func WaveformSvgIsNil() predicate.Sample {
	return predicate.Sample(sql.FieldIsNull(FieldWaveformSvg))
}

// WaveformSvgNotNil applies the NotNil predicate on the "waveform_svg" field.
func WaveformSvgNotNil() predicate.Sample {
	return predicate.Sample(sql.FieldNotNull(FieldWaveformSvg))
}

// HasSource applies the HasEdge predicate on the "source" edge.
func HasSource() predicate.Sample {
	return predicate.Sample(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, SourceTable, SourceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSourceWith applies the HasEdge predicate on the "source" edge with a given conditions (other predicates).
func HasSourceWith(preds ...predicate.SourceSample) predicate.Sample {
	return predicate.Sample(func(s *sql.Selector) {
		step := newSourceStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Sample) predicate.Sample {
	return predicate.Sample(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Sample) predicate.Sample {
	return predicate.Sample(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Sample) predicate.Sample {
	return predicate.Sample(sql.NotPredicates(p))
}
