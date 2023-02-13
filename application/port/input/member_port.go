package input

type MemberPort interface {
	// JoinMember 회원가입
	JoinMember() error
}
