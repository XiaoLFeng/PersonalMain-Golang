package SponsorService

import (
	"PersonalMain/internal/logic/SponsorServiceImpl"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewSponsorService() SponsorService {
	return &SponsorServiceImpl.SponsorServiceImpl{}
}

type SponsorService interface {
	GetSponsor(*ghttp.Request)
}
