package sponsorDAO

import (
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// GetSponsor
//
// 获取赞助
func GetSponsor() *[]do.SponsorDO {
	// 获取数据表全部数据
	var getSponorDO []do.SponsorDO
	result, err := g.Model("xf_sponsor").OrderDesc("created_at").All()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Structs(&getSponorDO)
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表数据提取成功")
			return &getSponorDO
		} else {
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表中没有赞助相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return nil
	}
}

// GetSponsorById
//
// 获取赞助
func GetSponsorById(id uint64) *do.SponsorDO {
	// 获取相应id数据信息
	var sponsorDO do.SponsorDO
	result, err := g.Model("xf_sponsor").Where("id", id).One()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Struct(&sponsorDO)
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表", id, "数据提取成功")
			return &sponsorDO
		} else {
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表中没有", id, "赞助信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return nil
	}
}

// AddSponsor
//
// 添加赞助
func AddSponsor(getSponsorDO do.SponsorDO) bool {
	//
	_, err := g.Model("xf_sponsor").Data(getSponsorDO).Insert()
	if err == nil {
		g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表数据添加成功")
		return true
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return false
	}
}

// GetSponsorType
//
// 获取检查赞助
func GetSponsorType(t uint8) *do.SponsorTypeDO {
	var getSponsorTypeDO do.SponsorTypeDO
	result, err := g.Model("xf_sponsor_type").Where("id", t).One()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Struct(&getSponsorTypeDO)
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponsor_type 数据表数据提取成功")
			return &getSponsorTypeDO
		} else {
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponsor_type 数据表中没有赞助类型相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return nil
	}
}

// GetCheckSponsor
//
// 获取检查赞助
func GetCheckSponsor() *[]do.SponsorDO {
	// 获取数据表全部数据
	var getSponorDO []do.SponsorDO
	result, err := g.Model("xf_sponsor").Where("check", false).OrderDesc("created_at").All()
	if err == nil {
		if !result.IsEmpty() {
			_ = result.Structs(&getSponorDO)
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表数据提取成功")
			return &getSponorDO
		} else {
			g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表中没有赞助相关信息")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return nil
	}
}

// CheckSponsorSuccess
//
// 检查赞助
func CheckSponsorSuccess(id uint64, check bool) bool {
	// 获取相应id数据信息
	_, err := g.Model("xf_sponsor").Data(g.Map{"check": check}).Where("id", id).Update()
	if err == nil {
		g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表", id, "数据更新成功")
		return true
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return false
	}
}

// DeleteSponsor
//
// 删除赞助
func DeleteSponsor(id uint64) bool {
	// 获取相应id数据信息
	_, err := g.Model("xf_sponsor").Where("id", id).Delete()
	if err == nil {
		g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表", id, "数据删除成功")
		return true
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return false
	}
}

// EditSponsorNoRegisterUser
//
// 编辑赞助
func EditSponsorNoRegisterUser(entity entity.SponsorEditVO) bool {
	// 获取相应id数据信息
	_, err := g.Model("xf_sponsor").Data(
		g.Map{
			"name":                 entity.Name,
			"type":                 entity.Type,
			"url":                  entity.Url,
			"money":                entity.Money,
			"statement_of_account": entity.StatementOfAccount,
			"created_at":           entity.CreatedAt,
			"check":                1,
		}).Where("id", entity.Id).Update()
	if err == nil {
		g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表", entity.Id, "数据更新成功")
		return true
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return false
	}
}

// EditSponsorRegisterUser
//
// 编辑赞助
func EditSponsorRegisterUser(entity entity.SponsorEditVO) bool {
	// 获取相应id数据信息
	_, err := g.Model("xf_sponsor").Data(
		g.Map{
			"name":                 entity.Name,
			"type":                 entity.Type,
			"user_id":              entity.UserId,
			"url":                  entity.UserId,
			"money":                entity.Money,
			"statement_of_account": entity.StatementOfAccount,
			"created_at":           entity.CreatedAt,
			"check":                0,
		}).Where("id", entity.Id).Update()
	if err == nil {
		g.Log().Cat("Database").Cat("Sponor").Notice(context.Background(), "xf_sponor 数据表", entity.Id, "数据更新成功")
		return true
	} else {
		g.Log().Cat("Database").Cat("Sponor").Error(context.Background(), err.Error())
		return false
	}
}
