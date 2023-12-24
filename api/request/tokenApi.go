package request

import "github.com/gogf/gf/v2/frame/g"

type TokenCreateReq struct {
	g.Meta `path:"/create" tags:"创建" method:"get" summary:"创建 Token"`
}

type TokenCreateRes struct{}
