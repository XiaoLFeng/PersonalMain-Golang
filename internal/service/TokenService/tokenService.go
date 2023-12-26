package TokenService

import (
	"PersonalMain/internal/logic/TokenServiceImpl"
	"PersonalMain/internal/model/do"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewTokenService() TokenService {
	return &TokenServiceImpl.DefaultTokenImpl{}
}

// TokenService
//
// TokenServiceImpl 服务接口
type TokenService interface {
	CreateToken() *do.TokenDO
	GetToken(*ghttp.Request) *do.TokenDO
	VerifyToken(*do.TokenDO) bool
	LoginToken(*ghttp.Request, do.UserDO) *do.TokenDO
}
