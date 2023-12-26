package user

import (
	"PersonalMain/api/request"
	"PersonalMain/internal/model/entity"
	"PersonalMain/internal/service/UserService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AuthRegister
//
// 用户注册
func (_ *ControllerV1) AuthRegister(ctx context.Context, _ *request.RegisterReq) (res *request.RegisterRes, err error) {
	userRegister := entity.UserRegisterVO{}
	req := ghttp.RequestFromCtx(ctx)
	// 获取 model 表单信息
	errStruct := req.GetRequestStruct(&userRegister)
	if errStruct == nil {
		errStruct := g.Validator().Data(userRegister).Run(ctx)
		if errStruct == nil {
			// 进行用户注册
			userService := UserService.NewUserService()
			userService.UserRegister(req, &userRegister)
		} else {
			ResultUtil.Error(req, ErrorCode.RequestBodyMismatching, errStruct.Map())
		}
	} else {
		ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
	}
	return res, err
}

func (_ *ControllerV1) AuthLogin(ctx context.Context, _ *request.LoginReq) (res *request.LoginRes, err error) {
	userLogin := entity.UserLoginVO{}
	req := ghttp.RequestFromCtx(ctx)
	// 获取 model 表单信息
	errStruct := req.GetRequestStruct(&userLogin)
	if errStruct == nil {
		errStruct := g.Validator().Data(userLogin).Run(ctx)
		if errStruct == nil {
			// 进行用户注册
			userService := UserService.NewUserService()
			userService.UserLogin(req, &userLogin)
		} else {
			ResultUtil.Error(req, ErrorCode.RequestBodyMismatching, errStruct.Map())
		}
	} else {
		ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
	}
	return res, err
}
