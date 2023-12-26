package LinkServiceImpl

import (
	"PersonalMain/internal/dao/linkDAO"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"PersonalMain/internal/service/UserService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"github.com/gogf/gf/v2/net/ghttp"
	regexp "regexp"
)

// GetLinkFriend
//
// 获取友链
func (*DefaultLinkImpl) GetLinkFriend(req *ghttp.Request) {
	// 数据库读取信息
	blogDO := linkDAO.GetBlogInformation()
	if blogDO != nil {
		ResultUtil.Success(req, "查询成功", blogDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
	}
}

// GetSortAndLink
//
// 获取分类和友链
func (*DefaultLinkImpl) GetSortAndLink(req *ghttp.Request) {
	// 获取Sort
	sortDO := linkDAO.GetSort()
	if sortDO != nil {
		// 博客内容归类
		sortDO = linkDAO.GetBlogForSort(sortDO)
		ResultUtil.Success(req, "查询成功", sortDO)
	} else {
		ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
	}
}

// AddLinkFriendCustom
//
// 添加友链
func (*DefaultLinkImpl) AddLinkFriendCustom(req *ghttp.Request, addFriendVO entity.LinkAddFriendVO) (bool, *do.BlogDO, string) {
	// 检查是否已存在此博客
	blogDO := linkDAO.GetBlogForName(addFriendVO.LinkName)
	if blogDO == nil {
		// 获取链接域名
		regex, _ := regexp.Compile("^(https?://)?([^:/\\s]+)")
		matches := regex.FindStringSubmatch(addFriendVO.LinkAddress)
		if len(matches) >= 3 {
			// 检查是否存在相似链接
			likeBlogDO := linkDAO.GetBlogForDomain(matches[2])
			if likeBlogDO == nil {
				// 数据处理
				var email *string
				var userId *uint64
				if !req.Cookie.Get("Token").IsEmpty() {
					email = &UserService.NewUserService().UserCurrent(req).Email
					userId = UserService.NewUserService().UserCurrent(req).Id
				} else {
					email = addFriendVO.LinkEmail
					userId = nil
				}
				var getRss *string
				if addFriendVO.LinkRssJudge {
					getRss = addFriendVO.LinkRss
				} else {
					getRss = nil
				}
				// 创建博客
				newBlogDO := do.BlogDO{
					BlogTitle:    addFriendVO.LinkName,
					BlogUrl:      addFriendVO.LinkAddress,
					BlogDesc:     addFriendVO.LinkDesc,
					BlogEmail:    email,
					BlogAvatar:   addFriendVO.LinkAvatar,
					BlogRssJudge: addFriendVO.LinkRssJudge,
					BlogRss:      getRss,
					BlogHost:     &addFriendVO.LinkHost,
					BlogLocation: addFriendVO.LinkUserLocation,
					BlogAddType:  0,
					BlogColor:    addFriendVO.LinkUserColor,
					BlogUserId:   userId,
					BlogRemark:   addFriendVO.LinkRemark,
				}
				if linkDAO.CreateBlog(newBlogDO) {
					return true, &newBlogDO, ""
				} else {
					return false, nil, "AddLinkFriendError"
				}
			} else {
				return false, nil, "LikeLinkAlreadyExists"

			}
		} else {
			return false, nil, "LinkAddressError"
		}
	} else {
		return false, nil, "FriendLinkAlreadyExists"
	}
}
