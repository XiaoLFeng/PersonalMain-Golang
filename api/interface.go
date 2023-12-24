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
}

type ITokenV1 interface {
	TokenCreate(ctx context.Context, req *request.TokenCreateReq) (res *request.TokenCreateRes, err error)
}
