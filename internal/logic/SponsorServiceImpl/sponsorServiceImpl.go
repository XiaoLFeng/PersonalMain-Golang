package SponsorServiceImpl

import (
	"PersonalMain/internal/dao/sponsorDAO"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/net/ghttp"
)

type SponsorServiceImpl struct{}

// GetSponsor
//
// 获取赞助
func (*SponsorServiceImpl) GetSponsor(req *ghttp.Request) {
	// 获取赞助
	getSponorDO := sponsorDAO.GetSponsor()
	if getSponorDO != nil {
		ResultUtil.Success(req, "Success", getSponorDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.NoSponsor)
	}
}
