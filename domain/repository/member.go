package repository

import (
	"context"
	"ent-study/ent"
)

type MemberRepository interface {
	FindByID(c context.Context, id int) (*ent.Member, error)
}
