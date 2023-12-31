package userDAO

import (
	"PersonalMain/internal/model/do"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// InsertUser
//
// 插入用户
func InsertUser(userDO do.UserDO) string {
	// 检查数据库是否存在用户
	result, err := g.Model("xf_user").Where("user_name = ? or email = ?", userDO.UserName, userDO.Email).One()
	if err == nil {
		if result.IsEmpty() {
			_, err := g.Model("xf_user").Data(userDO).Insert()
			if err != nil {
				g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
				return "DatabaseError"
			} else {
				g.Log().Cat("Database").Cat("User").Notice(context.Background(), "User", userDO.UserName, "创建成功")
				return "Success"
			}
		} else {
			g.Log().Cat("Database").Cat("User").Notice(context.Background(), "无法创建", userDO.UserName, "用户。原因：已存在此用户")
			return "UserExist"
		}
	} else {
		g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
		return "DatabaseError"
	}
}

// GetUser
//
// 获取用户信息
func GetUser(user string) *do.UserDO {
	userDO := do.UserDO{}
	result, err := g.Model("xf_user").Where("user_name = ? or email = ?", user, user).One()
	if err == nil {
		if result.IsEmpty() {
			g.Log().Cat("Database").Cat("User").Notice(context.Background(), "无法获取", user, "用户。原因：不存在此用户")
			return nil
		} else {
			err := result.Struct(&userDO)
			if err != nil {
				g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
				return nil
			} else {
				g.Log().Cat("Database").Cat("User").Notice(context.Background(), "获取", user, "用户成功")
				return &userDO
			}
		}
	} else {
		g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
		return nil
	}
}

// GetUserByToken
//
// 通过Token获取用户信息
func GetUserByToken(token string) *do.UserDO {
	userDO := do.UserDO{}
	getTokenResult, err := g.Model("xf_token").Fields("user_id").Where("token = ?", token).One()
	if err == nil {
		result, err := g.Model("xf_user").Where("id = ?", getTokenResult["user_id"]).One()
		if err == nil {
			if result.IsEmpty() {
				g.Log().Cat("Database").Cat("User").Notice(context.Background(), "无法获取", token, "用户。原因：不存在此用户")
				return nil
			} else {
				err := result.Struct(&userDO)
				if err != nil {
					g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
					return nil
				} else {
					g.Log().Cat("Database").Cat("User").Notice(context.Background(), "Token", token, "获取用户", userDO.UserName, "成功")
					return &userDO
				}
			}
		} else {
			g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
			return nil
		}
	} else {
		g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
		return nil
	}
}

// GetUserById
//
// 通过ID获取用户信息
func GetUserById(u uint64) *do.UserDO {
	userDO := do.UserDO{}
	result, err := g.Model("xf_user").Where("id = ?", u).One()
	if err == nil {
		if result.IsEmpty() {
			g.Log().Cat("Database").Cat("User").Notice(context.Background(), "无法获取ID为", u, "用户。原因：不存在此用户")
			return nil
		} else {
			err := result.Struct(&userDO)
			if err != nil {
				g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
				return nil
			} else {
				g.Log().Cat("Database").Cat("User").Notice(context.Background(), "获取ID为", u, "用户成功")
				return &userDO
			}
		}
	} else {
		g.Log().Cat("Database").Cat("User").Error(context.Background(), err.Error())
		return nil
	}
}
