package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type CollectionSample struct {
	ent.Schema
}

func (CollectionSample) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}

func (CollectionSample) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("sample"),
		index.Edges("collection"),
		index.Edges("sample", "collection").Unique(),
	}
}

func (CollectionSample) Fields() []ent.Field {
	return []ent.Field{
		field.String("sample_id"),
		field.UUID("collection_id", uuid.UUID{}),
	}
}

func (CollectionSample) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sample", SourceSample.Type).Required().Unique().Field("sample_id"),
		edge.To("collection", Collection.Type).Required().Unique().Field("collection_id").
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
