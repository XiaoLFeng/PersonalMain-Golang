package token

import (
	"PersonalMain/api/request"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/service"
	"PersonalMain/utility"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

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
			utility.Success(req, "Token 依然有效", nil)
		} else {
			// 生成新的 Session
			token = tokenService.CreateToken()
			utility.Success(req, "Token 已重新生成", g.Map{"token": token.Token})
		}
	} else {
		// 生成新的 Session
		token = tokenService.CreateToken()
		utility.Success(req, "Token 已生成", g.Map{"token": token.Token})
	}
	return res, err
}
