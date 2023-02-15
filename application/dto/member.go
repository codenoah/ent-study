package dto

import "ent-study/ent"

// JoinMemberRequest 회원가입시 사용하는 DTO
type JoinMemberRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Nickname string `json:"nickname" `
}

type Member struct {
	ID       int    `json:"id"`
	Email    string `json:"accountId"`
	Password []byte `json:"-"`
	Nickname string `json:"nickname"`
}

func (m *Member) Of(member *ent.Member) {
	m.ID = member.ID
	m.Email = member.Email
	m.Nickname = member.Nickname
}
