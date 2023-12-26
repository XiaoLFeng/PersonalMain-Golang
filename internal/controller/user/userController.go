package user

import (
	"PersonalMain/api"
	"PersonalMain/api/request"
	"PersonalMain/internal/service/UserService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func userService() UserService.UserService {
	return UserService.NewUserService()
}

type ControllerV1 struct{}

func NewUserV1() api.IUserV1 {
	return &ControllerV1{}
}

// GetUserCurrent
//
// 获取当前用户
func (*ControllerV1) GetUserCurrent(ctx context.Context, _ *request.GetUserReq) (res *request.GetUserRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	userDO := userService().UserCurrent(req)
	if userDO != nil {
		ResultUtil.Success(req, "获取成功", userDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.UserNotExist)
	}
	return res, err
}
