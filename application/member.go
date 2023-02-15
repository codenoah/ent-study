package application

import (
	"context"
	"ent-study/application/dto"
	"ent-study/domain/repository"
	"ent-study/errs"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type MemberApplication interface {
	GetMember(ctx context.Context, id int) (*dto.Member, error)
	JoinMember(ctx context.Context, request *dto.JoinMemberRequest) error
}

func NewMemberApplication(memberRepository repository.MemberRepository) *memberService {
	return &memberService{
		memberRepository: memberRepository,
	}
}

type memberService struct {
	memberRepository repository.MemberRepository
}

func (m memberService) GetMember(ctx context.Context, id int) (*dto.Member, error) {
	log.Println("memberService.GetMember")
	member, err := m.memberRepository.FindByID(ctx, id)
	if err != nil {
		return nil, errs.ErrNotFoundMember
	}

	result := new(dto.Member)
	result.Of(member)

	return result, nil
}

func (m memberService) JoinMember(ctx context.Context, request *dto.JoinMemberRequest) error {
	// 기존 사용자가 있는지 확인
	hasMember, err := m.memberRepository.ExistsByEmail(ctx, request.Email)
	if err != nil {
		return err
	}
	if hasMember {
		return errs.ErrAlreadyExistMember
	}

	// 패스워드 암호화
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.ErrPasswordGenerate
	}

	// repository 전달할 dto
	member := &dto.Member{
		Email:    request.Email,
		Password: password,
		Nickname: request.Nickname,
	}

	// 회원 등록
	if _, err = m.memberRepository.JoinMember(ctx, member); err != nil {
		return err
	}

	return nil
}
