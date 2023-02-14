package errs

var (
	NotFoundMember = New("4000", "회원을 찾을 수 없습니다")
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

func (e Error) Code() string {
	return e.code
}
func (e Error) Message() string {
	return e.message
}
