package request

import "github.com/gogf/gf/v2/frame/g"

type GetSponsorReq struct {
	g.Meta `path:"/list" tags:"获取赞助" method:"get" summary:"获取赞助"`
}

type GetSponsorRes struct{}
