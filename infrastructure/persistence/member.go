package persistence

import (
	"context"
	"ent-study/ent"
	"github.com/pkg/errors"
	"log"
)

type MemberRepository struct {
	*ent.Client
}

func NewMemberRepository(client *ent.Client) *MemberRepository {
	return &MemberRepository{
		Client: client,
	}
}

func (r *MemberRepository) FindByID(c context.Context, id int) (*ent.Member, error) {
	member, err := r.Member.Get(c, id)
	if err != nil {
		log.Println("ERROR: persistence find by id", err)
		return nil, errors.Wrap(err, "failed to find member")
	}
	return member, nil
}
