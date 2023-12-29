package sponsor

import (
	"PersonalMain/api"
	"PersonalMain/api/request"
	"PersonalMain/internal/service/SponsorService"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ControllerV1 struct{}

func NewSponsorV1() api.ISponsorV1 {
	return &ControllerV1{}
}

func sponorService() SponsorService.SponsorService {
	return SponsorService.NewSponsorService()
}

// GetSponsor
//
// 获取赞助
func (*ControllerV1) GetSponsor(ctx context.Context, _ *request.GetSponsorReq) (res *request.GetSponsorRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	sponorService().GetSponsor(req)
	return res, err
}
