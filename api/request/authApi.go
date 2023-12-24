package request

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta `path:"/register" tags:"注册" method:"get" summary:"注册账号"`

	Username string `json:"username"`
}

type RegisterRes struct {
	g.Meta `mime:"application/json" example:"string"`

	Output  string `json:"output"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
