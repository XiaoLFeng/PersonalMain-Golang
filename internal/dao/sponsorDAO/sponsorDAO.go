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
