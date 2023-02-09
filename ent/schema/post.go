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
		field.Int("id").Comment("Post table PK").
			StorageKey("post_id"),

		field.String("title").Comment("게시글의 타이틀"),

		field.String("content").Comment("게시글의 내용"),

		field.Time("created_at").Comment("게시글 생성 시간").
			Default(time.Now).
			Immutable(), // 불변 필드

		field.Time("updated_at").Comment("게시글 수정 시간").
			Optional().
			UpdateDefault(time.Now),

		field.Time("deleted_at").Comment("게시글 삭제 시간").
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
		edge.From("member", Member.Type). // ent.Post 객체에서 사용할 이름과 타입을 지정
							Ref("posts"). // Member 의 Edges 에서 지정한 이름( edge.To("post", Post.Type) )
							Unique(),     // FK 설정
		edge.To("comments", Comment.Type).
			StorageKey(edge.Column("post_id")),
	}
}
