package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("comment_id"),
		field.String("content"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional().UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("member", Member.Type). // Comment 를 조회 했을때 Member 에 접근하기 위한 변수 이름 설정
							Ref("comments"). // Member 에서의 Comment 엣지 이름
							Unique(),        // 1to1, 1toM 관계에서 사용
		edge.From("post", Post.Type).
			Ref("comments").
			Unique(), // 1to1, 1toM 관계에서 사용
	}
}
