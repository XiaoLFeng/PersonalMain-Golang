package entity

type LinkAddFriendVO struct {
	LinkEmail        *string `json:"link_email" v:"required|email#请输入邮箱|邮箱格式错误"`
	LinkHost         string  `json:"link_host" v:"required#请输入所用（主机）服务商"`
	LinkName         string  `json:"link_name" v:"required#请输入博客名称"`
	LinkAddress      string  `json:"link_address" v:"required|regex:[a-zA-z]+://[^\\s]*#请输入博客地址|网址不正确（例：https://blog.x-lf.com/）"`
	LinkDesc         string  `json:"link_desc" v:"required#请输入博客描述"`
	LinkAvatar       string  `json:"link_avatar" v:"required#请输入博客头像"`
	LinkRssJudge     bool    `json:"link_rss_judge" v:"required#请输入是否开启Rss"`
	LinkRss          *string `json:"link_rss" v:"required-if:link_rss_judge,true|regex:[a-zA-z]+://[^\\s]*#请输入Rss地址|Rss地址不正确（例：https://blog.x-lf.com/atom.xml）"`
	LinkUserLocation uint    `json:"link_user_location" v:"required#请输入期望板块"`
	LinkUserColor    uint    `json:"link_user_color" v:"required#请输入期望颜色"`
	LinkRemark       *string `json:"link_remark"`
	LinkAccept       bool    `json:"link_accept" v:"required#请同意协议"`
}
