package request

import "github.com/gogf/gf/v2/frame/g"

type GetLinkFriendReq struct {
	g.Meta `path:"/list" tags:"获取友链" method:"get" summary:"获取友链"`
}
type AddLinkFriendReq struct {
	g.Meta `path:"/add" tags:"添加友链" method:"post" summary:"添加友链"`
}
type GetSortAndLinkReq struct {
	g.Meta `path:"/sort" tags:"获取分类和友链" method:"get" summary:"获取分类和友链"`
}
type DelLinkFriendReq struct {
	g.Meta `path:"/delete" tags:"删除友链" method:"delete" summary:"删除友链"`
}

type GetLinkFriendRes struct{}
type AddLinkFriendRes struct{}
type GetSortAndLinkRes struct{}
type DelLinkFriendRes struct{}
