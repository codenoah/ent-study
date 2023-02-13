package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("post_id"),
		field.String("title"),
		field.String("content"),
		field.Time("created_at").
			Default(time.Now).
			// 불변 필드
			Immutable(),
		field.Time("updated_at").
			Optional().
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional(),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "post"},
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		// ent.PostEdges 객체에서 사용할 이름과 타입을 지정(ent.post.go - Post.PostEdges)
		edge.From("member", Member.Type).
			// Member 의 Edges 에서 지정한 이름( edge.To("posts", Post.Type) )
			Ref("posts").
			// FK 설정
			Unique(),
		edge.To("comments", Comment.Type).
			StorageKey(edge.Column("post_id")),
	}
}
