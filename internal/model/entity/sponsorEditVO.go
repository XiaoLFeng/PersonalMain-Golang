package entity

import "time"

type SponsorEditVO struct {
	Id                 uint64    `json:"id" v:"required|between:1,18446744073709551615#请输入赞助id|赞助id只能为:1到18446744073709551615"`
	Name               string    `json:"name" v:"required|length:1,50#请输入赞助名称|赞助名称长度为:1到:50位"`
	UserId             *uint64   `json:"user_id"`
	Url                *string   `json:"url" v:"url#请输入正确的赞助链接"`
	Type               uint8     `json:"type" v:"required|between:1,255#请选择赞助类型|赞助类型参考列表内容"`
	Money              uint64    `json:"money" v:"required|between:1,18446744073709551615#请输入赞助金额|赞助金额只能为:1到18446744073709551615"`
	StatementOfAccount *string   `json:"statement_of_account"`
	CreatedAt          time.Time `json:"created_at" v:"required#请输入赞助时间"`
}
