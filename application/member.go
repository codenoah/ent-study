package application

import (
	"context"
	"ent-study/application/dto"
	"ent-study/domain/repository"
	"github.com/pkg/errors"
	"log"
)

func NewMemberApplication(memberRepository repository.MemberRepository) *memberService {
	return &memberService{
		memberRepository: memberRepository,
	}
}

type MemberApplication interface {
	GetMember(ctx context.Context, id int) (*dto.Member, error)
}

type memberService struct {
	memberRepository repository.MemberRepository
}

func (m memberService) GetMember(ctx context.Context, id int) (*dto.Member, error) {
	log.Println("memberService.GetMember")
	member, err := m.memberRepository.FindByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get member")
	}

	result := new(dto.Member)
	result.Of(member)

	return result, nil
}
