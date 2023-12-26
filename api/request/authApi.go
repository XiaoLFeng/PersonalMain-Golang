package request

import (
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta `path:"/register" tags:"注册" method:"post" summary:"注册账号"`
}
type LoginReq struct {
	g.Meta `path:"/login" tags:"登录" method:"get" summary:"登录账号"`
}
type CheckReq struct {
	g.Meta `path:"/check" tags:"检查登录" method:"get" summary:"检查登录"`
}

type RegisterRes struct{}
type LoginRes struct{}
type CheckRes struct{}
