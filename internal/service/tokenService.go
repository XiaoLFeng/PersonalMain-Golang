package service

import (
	"PersonalMain/internal/logic"
	"PersonalMain/internal/model/do"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewTokenService() TokenService {
	return &logic.DefaultTokenImpl{}
}

// TokenService
//
// Token 服务接口
type TokenService interface {
	CreateToken() *do.TokenDO
	GetToken(req *ghttp.Request) *do.TokenDO
	VerifyToken(token *do.TokenDO) bool
}
