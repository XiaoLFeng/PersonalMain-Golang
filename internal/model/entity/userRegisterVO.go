package entity

type UserRegisterVO struct {
	Username   string `json:"username" v:"required|regex:^[0-9A-Za-z-_]+$|length:6,30#请输入用户名|用户名只允许 0-9, A-Z, a-z, -, _|用户名长度为:6-30位"`
	Email      string `json:"email" v:"required|email#请输入邮箱|邮箱格式错误"`
	Password   string `json:"password" v:"required|length:6,30#请输入密码|密码长度为:6-30位"`
	RePassword string `json:"rePassword" v:"required|same:Password#请输入确认密码|两次密码不一致"`
}
