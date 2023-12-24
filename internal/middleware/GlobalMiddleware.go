package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gregex"
	"strconv"
	"time"
)

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
				r.Response.WriteJson(g.Map{
					"code":    40011,
					"message": "时间戳过期",
					"data":    nil,
				})
			}
		}
	} else {
		if err != nil {
			r.Response.WriteJson(g.Map{
				"code":    40010,
				"message": "时间戳格式错误",
				"data":    nil,
			})
		}
	}
}
