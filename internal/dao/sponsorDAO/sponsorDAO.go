package sponsorDAO

import (
	"PersonalMain/internal/model/do"
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
