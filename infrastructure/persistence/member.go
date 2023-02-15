package persistence

import (
	"context"
	"ent-study/application/dto"
	"ent-study/domain/repository"
	"ent-study/ent"
	"ent-study/ent/member"
	"ent-study/errs"
)

var _ repository.MemberRepository = &MemberRepository{}

type MemberRepository struct {
	*ent.Client
}

func NewMemberRepository(client *ent.Client) *MemberRepository {
	return &MemberRepository{
		Client: client,
	}
}

func (r *MemberRepository) FindAll(c context.Context) ([]*ent.Member, error) {
	members, err := r.Member.Query().
		Where(member.DeletedAtIsNil()).
		All(c)

	if err != nil {
		return nil, errs.ErrNotFoundMember
	}

	return members, nil
}

func (r *MemberRepository) FindByID(c context.Context, id int) (*ent.Member, error) {
	member, err := r.Member.Get(c, id)

	if err != nil {
		return nil, err
	}

	return member, nil
}

func (r *MemberRepository) ExistsByEmail(c context.Context, email string) (bool, error) {
	exists, err := r.Member.Query().
		Where(member.Email(email)).
		Exist(c)

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *MemberRepository) JoinMember(c context.Context, member *dto.Member) (*ent.Member, error) {
	joinedMember, err := r.Member.Create().
		SetEmail(member.Email).
		SetPassword(member.Password).
		SetNickname(member.Nickname).
		Save(c)

	if err != nil {
		return nil, err
	}

	return joinedMember, nil
}
