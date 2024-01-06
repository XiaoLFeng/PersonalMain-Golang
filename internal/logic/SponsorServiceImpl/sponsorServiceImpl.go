package SponsorServiceImpl

import (
	"PersonalMain/internal/dao/sponsorDAO"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"PersonalMain/internal/service/TokenService"
	"PersonalMain/internal/service/UserService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

type SponsorServiceImpl struct{}

func tokenService() TokenService.TokenService {
	return TokenService.NewTokenService()
}

func userService() UserService.UserService {
	return UserService.NewUserService()
}

// AddSponsor
//
// 添加赞助
func (*SponsorServiceImpl) AddSponsor(req *ghttp.Request, entity entity.SponsorAddVO) {
	// 检查用户是否登陆
	getTokenDO := tokenService().GetToken(req)
	var getUserDO *do.UserDO
	var check = false
	if getTokenDO != nil {
		// 检查token是否有效
		if tokenService().VerifyToken(getTokenDO) {
			// 配置用户信息
			getUserDO = userService().UserCurrent(req)
			if getUserDO.Permission {
				check = true
			}
		} else {
			// token无效
			tokenService().DeleteToken(req)
			ResultUtil.ErrorNoData(req, ErrorCode.TokenVerifyFailed)
			return
		}
	} else {
		getUserDO = new(do.UserDO)
		getUserDO.UserName = entity.Name
	}
	// 数据检查
	getSponsorTypeInfo := sponsorDAO.GetSponsorType(entity.Type)
	if getSponsorTypeInfo == nil {
		ResultUtil.ErrorNoData(req, ErrorCode.NoSponsorType)
		return
	}
	// 数据处理
	newSponsorDO := do.SponsorDO{
		Name:               getUserDO.UserName,
		UserId:             getUserDO.Id,
		Url:                entity.Url,
		Type:               entity.Type,
		Money:              entity.Money,
		StatementOfAccount: entity.StatementOfAccount,
		CreatedAt:          time.Time{},
		Check:              check,
	}
	if sponsorDAO.AddSponsor(newSponsorDO) {
		//TODO：发送邮件
		ResultUtil.SuccessNoData(req, "添加成功")
	} else {
		//TODO：发送邮件
		ResultUtil.ErrorNoData(req, ErrorCode.AddSponsorFailed)
	}
}

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
