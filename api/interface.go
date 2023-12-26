// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package api

import (
	"PersonalMain/api/request"
	"context"
)

type IHelloV1 interface {
	Hello(ctx context.Context, req *request.HelloReq) (res *request.HelloRes, err error)
}

type IAuthV1 interface {
	AuthRegister(ctx context.Context, req *request.RegisterReq) (res *request.RegisterRes, err error)
	AuthLogin(ctx context.Context, req *request.LoginReq) (res *request.LoginRes, err error)
	AuthCheck(ctx context.Context, req *request.CheckReq) (res *request.CheckRes, err error)
	AuthLogout(ctx context.Context, req *request.LogoutReq) (res *request.LogoutRes, err error)
}

type ITokenV1 interface {
	TokenCreate(ctx context.Context, req *request.TokenCreateReq) (res *request.TokenCreateRes, err error)
	TokenVerify(ctx context.Context, req *request.TokenVerifyReq) (res *request.TokenVerifyRes, err error)
}
