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
	TimestampVerifyFailed       = ErrorCode{output: "TimestampVerifyFailed", code: 40010, message: "时间戳格式错误"}
	TimestampExpired            = ErrorCode{output: "TimestampExpired", code: 40011, message: "时间戳过期"}
	PasswordEncodeError         = ErrorCode{output: "PasswordEncodeError", code: 40012, message: "密码加密失败"}
	TokenExpired                = ErrorCode{output: "TokenExpired", code: 40100, message: "Token 已过期"}
	TokenVerifyFailed           = ErrorCode{output: "TokenVerifyFailed", code: 40101, message: "Token 验证失败"}
	TokenNotFound               = ErrorCode{output: "TokenNotFound", code: 40102, message: "Token 不存在"}
	PasswordNotMatch            = ErrorCode{output: "PasswordNotMatch", code: 40103, message: "密码错误"}
	AlreadyLogin                = ErrorCode{output: "AlreadyLogin", code: 40104, message: "已经登录"}
	NoPermission                = ErrorCode{output: "NoPermission", code: 40105, message: "没有权限"}
	RequestBodyMismatching      = ErrorCode{output: "RequestBodyMismatching", code: 40200, message: "请求体不匹配"}
	RequestBodyError            = ErrorCode{output: "RequestBodyError", code: 40201, message: "请求体错误"}
	UserExist                   = ErrorCode{output: "UserExists", code: 40300, message: "用户已存在"}
	UserNotExist                = ErrorCode{output: "UserNotExist", code: 40301, message: "用户不存在"}
	FriendLinkAlreadyExists     = ErrorCode{output: "FriendLinkAlreadyExists", code: 40302, message: "友链已存在"}
	LinkAddressError            = ErrorCode{output: "LinkAddressError", code: 40303, message: "链接地址错误"}
	LikeLinkAlreadyExists       = ErrorCode{output: "LikeLinkAlreadyExists", code: 40304, message: "链接已存在"}
	LinkAcceptDenied            = ErrorCode{output: "LinkAcceptDenied", code: 40305, message: "链接接受协议拒绝"}
	FriendLinkDoesNotExist      = ErrorCode{output: "FriendLinkDoesNotExist", code: 40306, message: "友链不存在"}
	NoSponsor                   = ErrorCode{output: "NoSponsor", code: 40307, message: "没有赞助"}
	AddSponsorFailed            = ErrorCode{output: "AddSponsorFailed", code: 40308, message: "添加赞助失败"}
	NoSponsorType               = ErrorCode{output: "NoSponsorType", code: 40309, message: "没有此赞助类型"}
	SponsorAlreadyCheck         = ErrorCode{output: "SponsorAlreadyCheck", code: 40310, message: "此赞助已经审核过"}
	NoSponsorInNoCheck          = ErrorCode{output: "NoSponsorInNoCheck", code: 40311, message: "没有未审核的赞助"}
	ServerUnknownError          = ErrorCode{output: "ServerUnknownError", code: 50000, message: "服务器未知错误"}
	ServerDatabaseInteriorError = ErrorCode{output: "ServerDatabaseInteriorError", code: 50001, message: "服务器数据库内部错误"}
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
