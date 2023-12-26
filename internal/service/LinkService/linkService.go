package LinkService

import (
	"PersonalMain/internal/logic/LinkServiceImpl"
	"PersonalMain/internal/model/do"
	"PersonalMain/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

func NewLinkService() LinkService {
	return &LinkServiceImpl.DefaultLinkImpl{}
}

type LinkService interface {
	GetLinkFriend(*ghttp.Request)
	GetSortAndLink(*ghttp.Request)
	AddLinkFriendCustom(*ghttp.Request, entity.LinkAddFriendVO) (bool, *do.BlogDO, string)
}
