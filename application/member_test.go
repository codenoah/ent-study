package application

import (
	"context"
	"ent-study/domain/repository"
	"ent-study/ent"
	"ent-study/errs"
	"testing"

	"ent-study/application/dto"
	"ent-study/ent/enttest"
	"ent-study/infrastructure/persistence"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type MemberServiceTestSuite struct {
	suite.Suite
	client           *ent.Client
	memberService    MemberApplication
	memberRepository repository.MemberRepository
}

func (s *MemberServiceTestSuite) SetupSuite() {
	s.client = enttest.Open(s.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	s.memberRepository = persistence.NewMemberRepository(s.client)
	s.memberService = NewMemberApplication(s.memberRepository)
}

func (s *MemberServiceTestSuite) TearDownSuite() {
	s.client.Close()
}

func TestMemberServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MemberServiceTestSuite))
}

func (s *MemberServiceTestSuite) TestJoinMember() {
	s.Run("회원 등록 성공", func() {
		err := s.memberService.JoinMember(context.Background(), &dto.JoinMemberRequest{
			Email:    "noah@gmail.com",
			Password: "noahpassword",
			Nickname: "noah",
		})
		s.NoError(err)
	})

	s.Run("회원 등록 실패 - 이메일 중복", func() {
		err := s.memberService.JoinMember(context.Background(), &dto.JoinMemberRequest{
			Email:    "noah@gmail.com",
			Password: "noahpassword",
			Nickname: "noah",
		})
		s.ErrorIs(err, errs.ErrAlreadyExistMember)
	})

}

func (s *MemberServiceTestSuite) TestGetMember() {
	s.Run("회원 아이디로 조회", func() {

	})
}
