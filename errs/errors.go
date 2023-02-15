package errs

var (
	ErrPasswordGenerate   = New("1001", "비밀번호를 생성할 수 없습니다")
	ErrNotFoundMember     = New("4000", "회원을 찾을 수 없습니다")
	ErrAlreadyExistMember = New("4001", "이미 존재하는 회원입니다")
)

func New(code string, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

type Error struct {
	code    string `json:"code"`
	message string `json:"message"`
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() string {
	return e.code
}
func (e Error) Message() string {
	return e.message
}
