package ResultUtil

import (
	"PersonalMain/utility/ErrorCode"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type DataResult struct{}

// Success
//
// 内容输出（包含 data）
func Success(req *ghttp.Request, message string, data interface{}) {
	req.Response.WriteJson(g.Map{
		"output":  "Success",
		"code":    200,
		"message": message,
		"data":    data,
	})
}

// SuccessNoData
//
// 内容输出（不含 data）
func SuccessNoData(req *ghttp.Request, message string) {
	req.Response.WriteJson(g.Map{
		"output":  "Success",
		"code":    200,
		"message": message,
	})
}

// Error
//
// 错误输出（包含 data）
func Error(req *ghttp.Request, errorCode ErrorCode.ErrorCode, data interface{}) {
	req.Response.WriteJson(g.Map{
		"output":  errorCode.Output(),
		"code":    errorCode.Code(),
		"message": errorCode.Message(),
		"data":    data,
	})
}

// ErrorNoData
//
// 错误输出（不含 data）
func ErrorNoData(req *ghttp.Request, errorCode ErrorCode.ErrorCode) {
	req.Response.WriteJson(g.Map{
		"output":  errorCode.Output(),
		"code":    errorCode.Code(),
		"message": errorCode.Message(),
	})
}

// ErrorDefault
//
// 默认错误输出（包含 data）
func ErrorDefault(req *ghttp.Request, output string, code int, message string, data interface{}) {
	req.Response.WriteJson(g.Map{
		"output":  output,
		"code":    code,
		"message": message,
		"data":    data,
	})
}
