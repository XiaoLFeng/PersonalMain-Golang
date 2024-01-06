// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package api

import (
	"PersonalMain/api/request"
	"context"
)

type IAuthV1 interface {
	AuthRegister(ctx context.Context, req *request.RegisterReq) (res *request.RegisterRes, err error)
	AuthLogin(ctx context.Context, req *request.LoginReq) (res *request.LoginRes, err error)
	AuthCheck(ctx context.Context, req *request.CheckReq) (res *request.CheckRes, err error)
	AuthLogout(ctx context.Context, req *request.LogoutReq) (res *request.LogoutRes, err error)
}

type ITokenV1 interface {
	TokenCreate(ctx context.Context, req *request.TokenCreateReq) (res *request.TokenCreateRes, err error)
	TokenVerify(ctx context.Context, req *request.TokenVerifyReq) (res *request.TokenVerifyRes, err error)
}

type ILinkV1 interface {
	GetLinkFriend(ctx context.Context, req *request.GetLinkFriendReq) (res *request.GetLinkFriendRes, err error)
	AddLinkFriend(ctx context.Context, req *request.AddLinkFriendReq) (res *request.AddLinkFriendRes, err error)
	GetSortAndLink(ctx context.Context, req *request.GetSortAndLinkReq) (res *request.GetSortAndLinkRes, err error)
	DelLinkFriend(ctx context.Context, req *request.DelLinkFriendReq) (res *request.DelLinkFriendRes, err error)
}

type IUserV1 interface {
	GetUserCurrent(ctx context.Context, req *request.GetUserReq) (res *request.GetUserRes, err error)
}

type ISponsorV1 interface {
	GetSponsor(context.Context, *request.GetSponsorReq) (*request.GetSponsorRes, error)
	AddSponsor(context.Context, *request.AddSponsorReq) (*request.AddSponsorRes, error)
	GetCheckSponsor(context.Context, *request.GetCheckSponsorReq) (*request.GetCheckSponsorRes, error)
	CheckSponsor(context.Context, *request.CheckSponsorReq) (*request.CheckSponsorRes, error)
}
