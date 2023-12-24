package user

import (
	"PersonalMain/api/request"
	"context"
)

func (c *ControllerV1) AuthRegister(_ context.Context, req *request.RegisterReq) (res *request.RegisterRes, err error) {
	res = &request.RegisterRes{
		Output:  "Success",
		Code:    200,
		Message: "Success",
		Data:    req,
	}
	return res, err
}
