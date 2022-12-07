package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
)

func (u *UserServer) UserPurview(ctx context.Context, req *proto.UserInfoPurview) (*empty.Empty, error) {
	var user model.User

	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	user.Position = int32(req.Position)
	global.DB.Save(&user)

	return &empty.Empty{}, nil
}
