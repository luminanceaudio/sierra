package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type SourceSample struct {
	ent.Schema
}

func (SourceSample) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("relative_path"),
		index.Fields("filename"),
	}
}

func (SourceSample) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("uri").Immutable().Unique(),
		field.String("relative_path").Immutable(),
		field.String("filename").Immutable(),
	}
}

func (SourceSample) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("source", Source.Type).Unique().StorageKey(edge.Column("source")),
		edge.To("sample", Sample.Type).Unique().StorageKey(edge.Column("sample")),
	}
}
