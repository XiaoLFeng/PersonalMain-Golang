package SponsorService

import (
	"PersonalMain/internal/logic/SponsorServiceImpl"
	"PersonalMain/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewSponsorService() SponsorService {
	return &SponsorServiceImpl.SponsorServiceImpl{}
}

type SponsorService interface {
	GetSponsor(*ghttp.Request)
	AddSponsor(*ghttp.Request, entity.SponsorAddVO)
	GetCheckSponsor(*ghttp.Request)
	CheckSponsor(*ghttp.Request, entity.CheckSponsorVO)
}
