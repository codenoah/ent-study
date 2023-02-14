package dto

import "ent-study/ent"

// JoinMemberRequest 회원가입시 사용하는 DTO
type JoinMemberRequest struct {
	ID       int    `json:"id" validate:"required"`
	Password string `json:"password" `
	Nickname string `json:"nickname" `
}

type JoinMemberResponse struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type Member struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
}

func (m *Member) Of(member *ent.Member) {
	m.ID = member.ID
	m.Nickname = member.Nickname
}
