package entity

type UserLoginVO struct {
	User     string `p:"user" v:"required|length:6,100#请输入用户名或邮箱"`
	Password string `p:"password" v:"required|length:6,30#请输入密码|密码长度为:6-30位"`
}
