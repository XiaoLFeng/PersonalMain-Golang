package request

import "github.com/gogf/gf/v2/frame/g"

type GetUserReq struct {
	g.Meta `path:"/current" tags:"获取用户信息" method:"get" summary:"获取用户信息"`
}

type GetUserRes struct{}
