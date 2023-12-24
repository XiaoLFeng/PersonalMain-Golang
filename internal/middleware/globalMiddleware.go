package middleware

import (
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"net/http"
	"strconv"
	"time"
)

// DefaultHandlerResponse is the default implementation of HandlerResponse.
type DefaultHandlerResponse struct {
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
	timestamp, err := strconv.ParseInt(r.GetHeader("timestamp"), 10, 64)
	if gregex.IsMatch(`^[0-9]{13,14}$`, []byte(strconv.FormatInt(timestamp, 10))) {
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
}

func JsonResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  = r.GetRequest("message").String()
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code int
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
			code = 200
			if msg == "" {
				msg = "success"
			}
		}
	}

	r.Response.WriteJson(DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    res,
	})
}
