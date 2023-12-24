package hello

import (
	"PersonalMain/api/request"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Hello(ctx context.Context, _ *request.HelloReq) (res *request.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
