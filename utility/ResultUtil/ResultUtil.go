package ResultUtil

import (
	"PersonalMain/utility/ErrorCode"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Success
//
// 内容输出（包含 data）
func Success(req *ghttp.Request, message string, data interface{}) {
	g.Log().Cat("Result").Debug(context.WithValue(context.Background(), req.RequestURI, req.RequestURI), req.RequestURI, "<Success[200]>", message)
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
	g.Log().Cat("Result").Debug(context.WithValue(context.Background(), req.RequestURI, req.RequestURI), req.RequestURI, "<SuccessNoData[200]>", message)
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
	g.Log().Cat("Result").Debug(context.WithValue(context.Background(), req.RequestURI, req.RequestURI), req.RequestURI, "<Error[", errorCode.Code(), "]>", errorCode.Message())
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
	g.Log().Cat("Result").Debug(context.WithValue(context.Background(), req.RequestURI, req.RequestURI), req.RequestURI, "ErrorNoData[", errorCode.Code(), "]>", errorCode.Message())
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
	g.Log().Cat("Result").Debug(context.WithValue(context.Background(), req.RequestURI, req.RequestURI), req.RequestURI, "ErrorDefault[", code, "]>", message)
	req.Response.WriteJson(g.Map{
		"output":  output,
		"code":    code,
		"message": message,
		"data":    data,
	})
}
