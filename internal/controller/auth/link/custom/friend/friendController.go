package friend

import (
	"PersonalMain/api"
	"PersonalMain/api/request"
	"PersonalMain/internal/model/entity"
	"PersonalMain/internal/service/LinkService"
	"PersonalMain/utility/ErrorCode"
	"PersonalMain/utility/ResultUtil"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type ControllerV1 struct{}

func NewLinkCustomFriendV1() api.ILinkV1 {
	return &ControllerV1{}
}

func linkService() LinkService.LinkService {
	return LinkService.NewLinkService()
}

// GetLinkFriend
//
// 获取友链
func (*ControllerV1) GetLinkFriend(ctx context.Context, _ *request.GetLinkFriendReq) (res *request.GetLinkFriendRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	linkService().GetLinkFriend(req)
	return res, err
}

// GetSortAndLink
//
// 获取分类和友链
func (*ControllerV1) GetSortAndLink(ctx context.Context, _ *request.GetSortAndLinkReq) (res *request.GetSortAndLinkRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	linkService().GetSortAndLink(req)
	return res, err
}

// AddLinkFriend
//
// 添加友链
func (*ControllerV1) AddLinkFriend(ctx context.Context, _ *request.AddLinkFriendReq) (res *request.AddLinkFriendRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	addFriendVO := entity.LinkAddFriendVO{}
	err = req.GetRequestStruct(&addFriendVO)
	if err == nil {
		// 检查对象
		errStruct := g.Validator().Data(addFriendVO).Run(ctx)
		if errStruct == nil {
			if addFriendVO.LinkAccept {
				hasAdd, blogDO, info := linkService().AddLinkFriendCustom(req, addFriendVO)
				if hasAdd {
					ResultUtil.Success(req, "添加成功", blogDO)
				} else {
					switch info {
					case "FriendLinkAlreadyExists":
						ResultUtil.ErrorNoData(req, ErrorCode.FriendLinkAlreadyExists)
					case "LinkAddressError":
						ResultUtil.ErrorNoData(req, ErrorCode.LinkAddressError)
					case "LikeLinkAlreadyExists":
						ResultUtil.ErrorNoData(req, ErrorCode.LikeLinkAlreadyExists)
					case "AddLinkFriendError":
						ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
					default:
						ResultUtil.ErrorNoData(req, ErrorCode.ServerUnknownError)
					}
				}
			} else {
				ResultUtil.ErrorNoData(req, ErrorCode.LinkAcceptDenied)
			}
		} else {
			g.Log().Cat("Struct").Cat("Link").Notice(ctx, errStruct.Map())
			ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
		}
	} else {
		g.Log().Cat("Struct").Cat("Link").Error(ctx, err.Error())
		ResultUtil.Error(req, ErrorCode.RequestBodyError, err.Error())
	}
	return res, err
}

// DelLinkFriend
//
// 删除友链
func (*ControllerV1) DelLinkFriend(ctx context.Context, _ *request.DelLinkFriendReq) (res *request.DelLinkFriendRes, err error) {
	req := ghttp.RequestFromCtx(ctx)
	// 获取业务
	delFriendVO := entity.LinkDelFriendVO{}
	err = req.GetRequestStruct(&delFriendVO)
	if err == nil {
		// 检查对象
		errStruct := g.Validator().Data(delFriendVO).Run(ctx)
		if errStruct == nil {
			hasDel, info := linkService().DelLinkFriendCustom(req, delFriendVO)
			if hasDel {
				ResultUtil.Success(req, "删除成功", nil)
			} else {
				switch info {
				case "FriendLinkDoesNotExist":
					ResultUtil.ErrorNoData(req, ErrorCode.FriendLinkDoesNotExist)
				case "DelLinkFriendError":
					ResultUtil.ErrorNoData(req, ErrorCode.ServerDatabaseInteriorError)
				default:
					ResultUtil.ErrorNoData(req, ErrorCode.ServerUnknownError)
				}
			}
		} else {
			g.Log().Cat("Struct").Cat("Link").Notice(ctx, errStruct.Map())
			ResultUtil.Error(req, ErrorCode.RequestBodyError, errStruct.Error())
		}
	} else {
		g.Log().Cat("Struct").Cat("Link").Error(ctx, err.Error())
		ResultUtil.Error(req, ErrorCode.RequestBodyError, err.Error())
	}
	return res, err
}
