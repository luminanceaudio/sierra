package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SourceSample struct {
	ent.Schema
}

func (SourceSample) Indexes() []ent.Index {
	return []ent.Index{}
}

func (SourceSample) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("relative_path").Immutable().Unique(),
	}
}

func (SourceSample) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("source", Source.Type).Unique().StorageKey(edge.Column("source")),
		edge.To("sample", Sample.Type).Unique().StorageKey(edge.Column("sample")),
	}
}
