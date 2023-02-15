package repository

import (
	"context"
	"ent-study/application/dto"
	"ent-study/ent"
)

type MemberRepository interface {
	// FindAll 전체 회원 조회
	FindAll(c context.Context) ([]*ent.Member, error)

	// FindByID ID 로 회원을 조회
	FindByID(c context.Context, id int) (*ent.Member, error)

	// ExistsByEmail 이메일로 회원이 존재하는지 확인
	ExistsByEmail(c context.Context, email string) (bool, error)

	// JoinMember 회원을 등록
	JoinMember(c context.Context, member *dto.Member) (*ent.Member, error)
}
