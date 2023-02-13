package service

import (
	"ent-study/application/port/input"
	"log"
)

// 인터페이스 구현 확인
var _ input.MemberPort = &memberService{}

// memberService 회원 관련 유즈케이스를 담당하는 구조체
type memberService struct {
}

func NewMemberService() *memberService {
	return &memberService{}
}

func (m memberService) JoinMember() error {
	log.Println("ㅋㅋㅋㅋㅋ")

	return nil
}
