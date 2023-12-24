package token

import (
	"PersonalMain/api/request"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/service"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// TokenCreate
//
// 生成 Token
func (_ *ControllerV1) TokenCreate(ctx context.Context, _ *request.TokenCreateReq) (res *request.TokenCreateRes, err error) {
	// 获取 Cookie 是否存在
	req := ghttp.RequestFromCtx(ctx)
	hasCookie := req.Cookie.Contains("token")
	tokenService := service.NewTokenService()
	var token *do.TokenDO
	if hasCookie {
		// 检查 Session 是否有效
		token = tokenService.GetToken(req)
		if tokenService.VerifyToken(token) {
			// 有效则输出Token依然有效
			ResultUtil.Success(req, "Token 依然有效", nil)
		} else {
			// 生成新的 Session
			token = tokenService.CreateToken()
			ResultUtil.Success(req, "Token 已重新生成", g.Map{"token": token.Token})
		}
	} else {
		// 生成新的 Session
		token = tokenService.CreateToken()
		ResultUtil.Success(req, "Token 已生成", g.Map{"token": token.Token})
	}
	return res, err
}

// TokenVerify
//
// 验证 Token
func (_ *ControllerV1) TokenVerify(ctx context.Context, _ *request.TokenVerifyReq) (res *request.TokenVerifyRes, err error) {
	// 获取 Cookie 是否存在
	req := ghttp.RequestFromCtx(ctx)
	hasCookie := req.Cookie.Contains("token")
	tokenService := service.NewTokenService()
	var token *do.TokenDO
	if hasCookie {
		// 检查 Session 是否有效
		token = tokenService.GetToken(req)
		if tokenService.VerifyToken(token) {
			ResultUtil.SuccessNoData(req, "Token 有效")
		} else {
			ResultUtil.ErrorNoData(req, ErrorCode.TokenVerifyFailed)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.TokenExpired)
	}
	return res, err
}
