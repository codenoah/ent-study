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
		field.Int("id").
			StorageKey("member_id"),
		field.String("nickname"),
		field.String("account_id"),
		field.String("account_type"),
		field.Time("created_at").
			Comment("사용자 계정 생성 시간").
			Default(time.Now).
			Immutable(), // 불변 필드
		field.Time("updated_at").
			Comment("사용자 계정 수정 시간").
			Optional().
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Comment("사용자 계정 삭제 시간").
			Optional(),
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
