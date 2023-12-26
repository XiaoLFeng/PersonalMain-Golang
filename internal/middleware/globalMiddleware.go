package middleware

import (
	"PersonalMain/internal/service/TokenService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"net/http"
	"strconv"
	"time"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
	Output  string      `json:"output"  dc:"Output data for certain request according API definition"`
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
}

// TimestampMiddleware
//
// 全局中间件
// 检查时间戳误差是否在
func TimestampMiddleware(r *ghttp.Request) {
	// 检查时间戳误差是否在
	timestampString := r.GetHeader("timestamp")
	if len(timestampString) >= 13 && len(timestampString) <= 14 {
		timestamp, err := strconv.ParseInt(timestampString, 10, 64)
		if gregex.IsMatch(`^[0-9]+$`, []byte(strconv.FormatInt(timestamp, 10))) {
			if timestamp+int64(2000) > time.Now().UnixMilli() && timestamp-int64(2000) < time.Now().UnixMilli() {
				r.Middleware.Next()
			} else {
				if err != nil {
					ResultUtil.ErrorNoData(r, ErrorCode.TimestampExpired)
				}
			}
		} else {
			if err != nil {
				ResultUtil.ErrorNoData(r, ErrorCode.TimestampVerifyFailed)
			}
		}
	} else {
		ResultUtil.ErrorNoData(r, ErrorCode.TimestampVerifyFailed)
	}
}

// VerifyTokenMiddleware
//
// 校验 TokenServiceImpl 是否有效
func VerifyTokenMiddleware(r *ghttp.Request) {
	// 校验 token
	tokenService := TokenService.NewTokenService()
	getToken := tokenService.GetToken(r)
	if getToken != nil {
		// 检查 TokenServiceImpl 是否有效
		if tokenService.VerifyToken(getToken) {
			r.Middleware.Next()
		} else {
			ResultUtil.ErrorNoData(r, ErrorCode.TokenExpired)
		}
	} else {
		ResultUtil.ErrorNoData(r, ErrorCode.TokenNotFound)
	}
}

func JsonResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		output = r.GetRequest("output").String()
		msg    = r.GetRequest("message").String()
		err    = r.GetError()
		res    = r.GetHandlerResponse()
		code   int
	)
	if err != nil {
		if r.GetRequest("code") == nil {
			code = 500
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != 200 {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = 404
			case http.StatusForbidden:
				code = 403
			default:
				code = r.GetRequest("code").Int()
			}
			// 处理错误
			err = gerror.New(r.GetRequest("message").String())
			r.SetError(err)
		} else {
			code = ErrorCode.ServerUnknownError.Code()
			if msg == "" {
				msg = ErrorCode.ServerUnknownError.Message()
				output = ErrorCode.ServerUnknownError.Output()
			}
			g.Log().Cat("Unknown").Info(context.Background(), output, msg)
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Output:  output,
		Code:    code,
		Message: msg,
		Data:    res,
	})
}
