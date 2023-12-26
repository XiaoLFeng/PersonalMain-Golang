package UserServiceImpl

import (
	"PersonalMain/internal/dao/userDAO"
	"PersonalMain/internal/logic/TokenServiceImpl"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"PersonalMain/internal/service/TokenService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type DefaultUserImpl struct{}

func tokenService() TokenService.TokenService {
	return &TokenServiceImpl.DefaultTokenImpl{}
}

// UserRegister
//
// 用户注册
func (_ *DefaultUserImpl) UserRegister(req *ghttp.Request, userVO *entity.UserRegisterVO) {
	// 密码加密
	enPassword, err := bcrypt.GenerateFromPassword([]byte(userVO.Password), bcrypt.DefaultCost)
	if err == nil {
		// 注册用户
		newUserDO := do.UserDO{
			Id:          nil,
			UserName:    userVO.Username,
			DisplayName: nil,
			Email:       userVO.Email,
			Qq:          nil,
			Password:    string(enPassword),
			OldPassword: nil,
			EmailVerify: false,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
		}
		// 注入注册
		switch userDAO.InsertUser(newUserDO) {
		case "Success":
			userVO.Password = ""
			userVO.RePassword = ""
			ResultUtil.Success(req, "注册成功", userVO)
			break
		case "UserExist":
			ResultUtil.ErrorNoData(req, ErrorCode.UserExist)
		case "DatabaseError":
			ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
		default:
			ResultUtil.ErrorNoData(req, ErrorCode.ServerUnknownError)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.PasswordEncodeError)
	}
}

// UserLogin
//
// 用户登录
func (_ *DefaultUserImpl) UserLogin(req *ghttp.Request, userVO *entity.UserLoginVO) {
	// 获取数据库中用户信息
	getUserDO := userDAO.GetUser(userVO.User)
	if getUserDO != nil {
		// 比对密码-=
		err := bcrypt.CompareHashAndPassword([]byte(getUserDO.Password), []byte(userVO.Password))
		if err == nil {
			// Token 注册更新
			getTokenDO := tokenService().LoginToken(req, *getUserDO)
			if getTokenDO != nil {
				getUserDO.Password = ""
				getUserDO.OldPassword = nil
				ResultUtil.Success(req, "登录成功", g.Map{"user": getUserDO, "token": getTokenDO})
			}
		} else {
			ResultUtil.ErrorNoData(req, ErrorCode.PasswordNotMatch)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.UserNotExist)
	}
}

// CheckLogin
//
// 检查登录
func (*DefaultUserImpl) CheckLogin(req *ghttp.Request) {
	// 获取 Token
	userDO := userDAO.GetUserByToken(req.Cookie.Get("token").String())
	if userDO != nil {
		userDO.Password = ""
		userDO.OldPassword = nil
		ResultUtil.Success(req, "用户已处于登陆状态", userDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.UserNotExist)
	}
}

// UserLogout
//
// 用户登出
func (*DefaultUserImpl) UserLogout(req *ghttp.Request) {
	if tokenService().DeleteToken(req) {
		ResultUtil.Success(req, "用户已退出登录", nil)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.TokenNotFound)
	}
}
