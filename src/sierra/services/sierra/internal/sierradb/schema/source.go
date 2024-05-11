package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type Source struct {
	ent.Schema
}

func (Source) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (Source) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
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
