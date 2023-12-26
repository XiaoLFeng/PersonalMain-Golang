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

func userService() UserService.UserService {
	return UserService.NewUserService()
}

// AuthRegister
//
// 用户注册
func (*ControllerV1) AuthRegister(ctx context.Context, _ *request.RegisterReq) (res *request.RegisterRes, err error) {
	userRegister := entity.UserRegisterVO{}
	req := ghttp.RequestFromCtx(ctx)
	// 获取 model 表单信息
	errStruct := req.GetRequestStruct(&userRegister)
	if errStruct == nil {
		errStruct := g.Validator().Data(userRegister).Run(ctx)
		if errStruct == nil {
			// 进行用户注册
			userService().UserRegister(req, &userRegister)
		} else {
			ResultUtil.Error(req, ErrorCode.RequestBodyMismatching, errStruct.Map())
		}
	} else {
		ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
	}
	return res, err
}

func (*ControllerV1) AuthLogin(ctx context.Context, _ *request.LoginReq) (res *request.LoginRes, err error) {
	userLogin := entity.UserLoginVO{}
	req := ghttp.RequestFromCtx(ctx)
	// 获取 model 表单信息
	errStruct := req.GetRequestStruct(&userLogin)
	if errStruct == nil {
		errStruct := g.Validator().Data(userLogin).Run(ctx)
		if errStruct == nil {
			// 进行用户注册
			userService().UserLogin(req, &userLogin)
		} else {
			ResultUtil.Error(req, ErrorCode.RequestBodyMismatching, errStruct.Map())
		}
	} else {
		ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
	}
	return res, err
}

// AuthCheck
//
// 检查登录
func (*ControllerV1) AuthCheck(ctx context.Context, _ *request.CheckReq) (res *request.CheckRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取数据库中用户信息
	userService().CheckLogin(req)
	return res, err
}

// AuthLogout
//
// 用户登出
func (*ControllerV1) AuthLogout(ctx context.Context, _ *request.LogoutReq) (res *request.LogoutRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取数据库中用户信息
	userService().UserLogout(req)
	return res, err
}
