package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Sample struct {
	ent.Schema
}

func (Sample) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}

func (Sample) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("sha256").Immutable().Unique(),
		field.String("format").Optional(),
		field.Int64("length").Optional(),
		field.String("waveform_storage_path").Optional(),
	}
}

func (Sample) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", SourceSample.Type).Ref("sample"),
	}
}
