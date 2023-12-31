package request

import "github.com/gogf/gf/v2/frame/g"

type GetSponsorReq struct {
	g.Meta `path:"/list" tags:"获取赞助" method:"get" summary:"获取赞助"`
}
type AddSponsorReq struct {
	g.Meta `path:"/add" tags:"添加赞助" method:"post" summary:"添加赞助"`
}
type GetCheckSponsorReq struct {
	g.Meta `path:"/list" tags:"获取检查赞助" method:"get" summary:"获取检查赞助"`
}
type CheckSponsorReq struct {
	g.Meta `path:"/" tags:"检查赞助" method:"patch" summary:"检查赞助"`
}
type EditSponsorReq struct {
	g.Meta `path:"/edit" tags:"编辑赞助" method:"put" summary:"编辑赞助"`
}
type DeleteSponsorReq struct {
	g.Meta `path:"/delete" tags:"删除赞助" method:"delete" summary:"删除赞助"`
}

type GetSponsorRes struct{}
type AddSponsorRes struct{}
type GetCheckSponsorRes struct{}
type CheckSponsorRes struct{}
type EditSponsorRes struct{}
type DeleteSponsorRes struct{}
