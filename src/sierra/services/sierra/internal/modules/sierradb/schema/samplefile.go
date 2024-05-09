package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SampleFile struct {
	ent.Schema
}

func (SampleFile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sha256").Unique(),
	}
}

func (SampleFile) Fields() []ent.Field {
	return []ent.Field{
		field.String("sha256"),
		field.String("format").Optional(),
		field.Int64("length").Optional(),
	}
}

func (SampleFile) Edges() []ent.Edge {
	return nil
}
