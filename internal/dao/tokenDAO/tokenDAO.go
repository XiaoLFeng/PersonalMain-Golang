package tokenDAO

import (
	"PersonalMain/internal/model/do"
	"PersonalMain/utility/CustomError"
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
		g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", token.Token, "创建成功")
		return token
	} else {
		g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
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
				g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
				return nil
			} else {
				g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", tokenDO.Token, "获取成功")
				return &tokenDO
			}
		} else {
			g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", token, "获取失败，原因：Token不存在")
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
		return nil
	}
}

// DeleteToken
//
// 删除Token业务
func DeleteToken(token string) bool {
	_, err := g.Model("xf_token").Where("token = ?", token).Delete()
	if err != nil {
		g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
		return false
	} else {
		g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", token, "删除成功")
		return true
	}
}

// UpdateToken
//
// 更新Token业务
func UpdateToken(token string, userId *uint64) (*do.TokenDO, error) {
	// 查找 token
	getTokenDO := GetToken(token)
	if getTokenDO != nil {
		if getTokenDO.UserId == nil {
			// 获取数据库信息
			result, err := g.Model("xf_token").Where("token = ?", getTokenDO.Token).One()
			if err != nil {
				g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
				errorData := &CustomError.CustomError{Message: "DatabaseError"}
				return nil, errorData
			}
			// 更新数据库信息
			var newTokenDO = do.TokenDO{}
			_ = result.Struct(&newTokenDO)
			newTokenDO.UserId = userId
			newTokenDO.ExpiredAt = time.Now().Add(time.Hour * 24)
			_, err = g.Model("xf_token").Data(newTokenDO).Where("token = ?", getTokenDO.Token).Update()
			if err != nil {
				g.Log().Cat("Database").Cat("Token").Error(context.Background(), err.Error())
				errorData := &CustomError.CustomError{Message: "DatabaseError"}
				return nil, errorData
			} else {
				g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", getTokenDO.Token, "更新成功")
				return &newTokenDO, nil
			}
		} else {
			g.Log().Cat("Database").Cat("Token").Notice(context.Background(), "Token", getTokenDO.Token, "更新失败，原因：该Token已登陆")
			errorData := &CustomError.CustomError{Message: "AlreadyLogin"}
			return nil, errorData
		}
	} else {
		errorData := &CustomError.CustomError{Message: "TokenNotFound"}
		return nil, errorData
	}
}
