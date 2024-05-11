package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Source struct {
	ent.Schema
}

func (Source) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Source) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("uri").Immutable().Unique(),
	}
}

func (Source) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sample", SourceSample.Type).Ref("source"),
	}
}
