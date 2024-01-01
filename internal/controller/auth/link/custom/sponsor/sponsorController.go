package sponsor

import (
	"PersonalMain/api"
	"PersonalMain/api/request"
	"PersonalMain/internal/model/entity"
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

// AddSponsor
//
// 添加赞助
func (*ControllerV1) AddSponsor(ctx context.Context, _ *request.AddSponsorReq) (res *request.AddSponsorRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	sponsorAddVO := entity.SponsorAddVO{}
	err = req.GetRequestStruct(&sponsorAddVO)
	if err == nil {
		sponorService().AddSponsor(req, sponsorAddVO)
	}
	return res, err
}
