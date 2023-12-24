package tokenDAO

import (
	"PersonalMain/internal/model/do"
	"PersonalMain/utility/Processing"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// CreateToken
//
// 创建Token业务
func CreateToken() do.TokenDO {
	token := do.TokenDO{
		Id:        nil,
		UserId:    nil,
		Token:     Processing.CreateToken(),
		ExpiredAt: time.Now().Add(time.Minute * 10),
		CreatedAt: time.Now(),
	}
	_, err := g.Model("xf_token").Data(token).OmitEmpty().Insert()
	if err == nil {
		g.Log().Cat("Database").Cat("Token").Debug(context.TODO(), "Token", token.Token, "创建成功")
		return token
	} else {
		g.Log().Cat("Database").Cat("Token").Error(context.TODO(), err.Error())
		return do.TokenDO{}
	}
}

// GetToken
//
// 获取Token业务
func GetToken(token string) *do.TokenDO {
	var tokenDO do.TokenDO
	result, err := g.Model("xf_token").Where("token = ?", token).One()
	if err == nil {
		// 检查数据库是否为空
		if !result.IsEmpty() {
			err := result.Struct(&tokenDO)
			if err != nil {
				g.Log().Cat("Database").Cat("Token").Error(context.TODO(), err.Error())
				return nil
			} else {
				g.Log().Cat("Database").Cat("Token").Debug(context.TODO(), "Token", tokenDO.Token, "获取成功")
				return &tokenDO
			}
		} else {
			g.Log().Cat("Database").Cat("Token").Debug(context.TODO(), "xf_token 数据表为空")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Token").Error(context.TODO(), err.Error())
		return nil
	}
}

// DeleteToken
//
// 删除Token业务
func DeleteToken(token string) {
	_, err := g.Model("xf_token").Where("token = ?", token).Delete()
	if err != nil {
		g.Log().Cat("Database").Cat("Token").Error(context.TODO(), err.Error())
	} else {
		g.Log().Cat("Database").Cat("Token").Debug(context.TODO(), "Token", token, "删除成功")
	}
}
