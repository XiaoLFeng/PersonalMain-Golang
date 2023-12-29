package entity

type LinkDelFriendVO struct {
	Id        uint64 `json:"id" v:"required|regex:^[0-9]+$#请输入id|id格式错误"`
	LinkEmail string `json:"link_email" v:"required|email#请输入邮箱|邮箱格式错误"`
}
