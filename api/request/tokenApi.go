package request

import "github.com/gogf/gf/v2/frame/g"

type TokenCreateReq struct {
	g.Meta `path:"/create" tags:"创建" method:"get" summary:"创建 Token"`
}
type TokenCreateRes struct{}

type TokenVerifyReq struct {
	g.Meta `path:"/verify" tags:"验证" method:"get" summary:"验证 Token"`
}
type TokenVerifyRes struct{}
