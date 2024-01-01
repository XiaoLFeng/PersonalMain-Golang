package entity

type SponsorAddVO struct {
	Name               string  `json:"name" v:"required|length:1,50#请输入赞助名称|赞助名称长度为:1到:50位"`
	Url                *string `json:"url" v:"url#请输入正确的赞助链接"`
	Type               uint8   `json:"type" v:"required|between:1,4#请选择赞助类型|赞助类型只能为:1或:2"`
	Money              uint64  `json:"money" v:"required|between:1,1000000#请输入赞助金额|赞助金额只能为:1到1000000"`
	StatementOfAccount *string `json:"statement_of_account" v:"required#请输入流水号，以便博主可以查明账单来源"`
}
