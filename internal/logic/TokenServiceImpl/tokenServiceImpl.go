package TokenServiceImpl

import (
	"PersonalMain/internal/dao/tokenDAO"
	"PersonalMain/internal/model/do"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

type DefaultTokenImpl struct{}

// CreateToken
//
// 创建Token业务
func (_ DefaultTokenImpl) CreateToken() *do.TokenDO {
	token := tokenDAO.CreateToken()
	// 检查数据库中是否存在该 token
	return &token
}

// GetToken
//
// 获取Token业务
func (_ DefaultTokenImpl) GetToken(req *ghttp.Request) *do.TokenDO {
	// 获取 Cookie 中的 token
	cookieToken := req.Cookie.Get("token")
	if cookieToken != nil {
		// 数据库查找 token
		token := tokenDAO.GetToken(cookieToken.String())
		// 检查数据库中是否存在该 token
		if token != nil {
			// 检查Token是否超时
			return token
		}
	}
	return nil
}

// VerifyToken
//
// 验证Token是否有效
func (_ DefaultTokenImpl) VerifyToken(token *do.TokenDO) bool {
	if token != nil {
		if token.ExpiredAt.After(time.Now()) {
			return true
		} else {
			// 删除数据库
			tokenDAO.DeleteToken(token.Token)
			return false
		}
	} else {
		return false
	}
}

// LoginToken
//
// 登录Token业务
func (_ DefaultTokenImpl) LoginToken(req *ghttp.Request, userDO do.UserDO) *do.TokenDO {
	// 更新数据库
	newTokenDO, err := tokenDAO.UpdateToken(req.Cookie.Get("token").String(), userDO.Id)
	if err == nil {
		if newTokenDO != nil {
			return newTokenDO
		}
	} else {
		switch err.Error() {
		case "DatabaseError":
			ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
		case "AlreadyLogin":
			ResultUtil.ErrorNoData(req, ErrorCode.AlreadyLogin)
		case "TokenNotFound":
			ResultUtil.ErrorNoData(req, ErrorCode.TokenNotFound)
		}
	}
	return nil
}

// DeleteToken
//
// 删除Token业务
func (_ DefaultTokenImpl) DeleteToken(req *ghttp.Request) bool {
	// 获取 Cookie 中的 token
	cookieToken := req.Cookie.Get("token")
	if cookieToken != nil {
		// 数据库查找 token
		token := tokenDAO.GetToken(cookieToken.String())
		// 检查数据库中是否存在该 token
		if token != nil {
			// 删除数据库
			return tokenDAO.DeleteToken(token.Token)
		}
	}
	return false
}
