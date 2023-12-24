package ErrorCode

type ErrorCode struct {
	output  string
	code    int
	message string
}

type Errors interface {
	Output() string
	Code() int
	Message() string
}

var (
	TokenExpired      = ErrorCode{output: "error", code: 40100, message: "Token 已过期"}
	TokenVerifyFailed = ErrorCode{output: "error", code: 40101, message: "Token 验证失败"}
)

func (e ErrorCode) Output() string {
	return e.output
}

func (e ErrorCode) Code() int {
	return e.code
}

func (e ErrorCode) Message() string {
	return e.message
}
