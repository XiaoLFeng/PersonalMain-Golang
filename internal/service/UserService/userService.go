package UserService

import (
	"PersonalMain/internal/logic/UserServiceImpl"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewUserService() UserService {
	return &UserServiceImpl.DefaultUserImpl{}
}

type UserService interface {
	UserRegister(*ghttp.Request, *entity.UserRegisterVO)
	UserLogin(*ghttp.Request, *entity.UserLoginVO)
	CheckLogin(*ghttp.Request)
	UserLogout(*ghttp.Request)
	UserCurrent(*ghttp.Request) *do.UserDO
}
