package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Member holds the schema definition for the Member entity.
type Member struct {
	ent.Schema
}

// Fields of the Member.
func (Member) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("member_id"),
		field.String("nickname"),
		field.String("email").Unique(),
		field.Bytes("password"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

func (Member) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "member"},
	}
}

// Edges of the Member.
func (Member) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).
			StorageKey(edge.Column("member_id")),
		edge.To("comments", Comment.Type).
			StorageKey(edge.Column("member_id")),
	}
}
