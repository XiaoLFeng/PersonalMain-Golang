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

// GetCheckSponsor
//
// 获取检查赞助
func (*SponsorServiceImpl) GetCheckSponsor(req *ghttp.Request) {
	// 检查用户是否是管理员
	if userService().CheckAdministrator(req) {
		// 获取赞助
		getSponorDO := sponsorDAO.GetCheckSponsor()
		if getSponorDO != nil {
			ResultUtil.Success(req, "Success", getSponorDO)
		} else {
			ResultUtil.ErrorNoData(req, ErrorCode.NoSponsorInNoCheck)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.NoPermission)
	}
}

// CheckSponsor
//
// 检查赞助
func (*SponsorServiceImpl) CheckSponsor(req *ghttp.Request, entity entity.CheckSponsorVO) {
	// 检查用户是否是管理员
	if userService().CheckAdministrator(req) {
		// 获取此单位注册信息
		getSponsorDO := sponsorDAO.GetSponsorById(entity.Id)
		if getSponsorDO != nil {
			// 检查是否已经审核过
			if getSponsorDO.Check {
				ResultUtil.ErrorNoData(req, ErrorCode.SponsorAlreadyCheck)
			} else {
				// 对内容内容进行审核管理
				if entity.Check {
					// 更新数据
					if sponsorDAO.CheckSponsorSuccess(entity.Id, true) {
						ResultUtil.SuccessOther(req, "CheckSuccess", "审核通过，内容已处理")
					} else {
						ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
					}
				} else {
					// 删除数据
					if sponsorDAO.DeleteSponsor(entity.Id) {
						ResultUtil.SuccessOther(req, "CheckDenied", "审核不通过，内容已处理")
					} else {
						ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
					}
				}
			}
		} else {
			ResultUtil.ErrorNoData(req, ErrorCode.NoSponsor)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.NoPermission)
	}
}

// EditSponsor
//
// 编辑赞助
func (*SponsorServiceImpl) EditSponsor(req *ghttp.Request, entity entity.SponsorEditVO) {
	// 检查用户是否是管理员
	if userService().CheckAdministrator(req) {
		// 获取此单位注册信息
		getSponsorDO := sponsorDAO.GetSponsorById(entity.Id)
		if getSponsorDO != nil {
			// 检查类型是否存在
			if sponsorDAO.GetSponsorType(entity.Type) != nil {
				// 检查为注册用户或非注册用户
				if entity.UserId == nil {
					// 对数据进行上传(NoRegisterUser)
					if sponsorDAO.EditSponsorNoRegisterUser(entity) {
						ResultUtil.SuccessNoData(req, "编辑成功")
					} else {
						ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
					}
				} else {
					// 检查用户是否存在
					getUserDO := userService().GetUserById(*entity.UserId)
					if getUserDO != nil {
						// 对数据进行上传(RegisterUser)
						if sponsorDAO.EditSponsorRegisterUser(entity) {
							ResultUtil.SuccessNoData(req, "编辑成功")
						} else {
							ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
						}
					} else {
						ResultUtil.ErrorNoData(req, ErrorCode.UserNotExist)
					}
				}
			} else {
				ResultUtil.ErrorNoData(req, ErrorCode.NoSponsorType)
			}
		} else {
			ResultUtil.ErrorNoData(req, ErrorCode.NoSponsor)
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.NoPermission)
	}
}

// DeleteSponsor
//
// 删除赞助
func (*SponsorServiceImpl) DeleteSponsor(req *ghttp.Request, entity entity.SponsorDelVO) {
	// 检查用户是否是管理员
	if userService().CheckAdministrator(req) {
		// 获取此单位注册信息
		getSponsorDO := sponsorDAO.GetSponsorById(entity.Id)
		if getSponsorDO != nil {
			// 删除数据
			if sponsorDAO.DeleteSponsor(entity.Id) {
				ResultUtil.SuccessNoData(req, "删除成功")
			} else {
				ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
			}
		}
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.NoPermission)
	}
}
