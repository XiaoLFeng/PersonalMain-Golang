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
	NoneDataResult = ErrorCode{output: "success", code: 200, message: "success"}
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
