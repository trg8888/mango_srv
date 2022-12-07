package handler

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
)

func (u *UserServer) UserByID(ctx context.Context, req *proto.UserInfoByID) (*proto.UserInfo, error) {
	var user model.User

	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	ModelToRsponse(user)
	return ModelToRsponse(user), nil
}

func (u *UserServer) UserByName(ctx context.Context, req *proto.UserInfoByName) (*proto.UserInfo, error) {
	var user model.User

	result := global.DB.Where(&model.User{Name: req.Name}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	ModelToRsponse(user)
	return ModelToRsponse(user), nil
}

func (u *UserServer) UserByMobile(ctx context.Context, req *proto.UserInfoByMobile) (*proto.UserInfo, error) {
	var user model.User

	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	ModelToRsponse(user)
	return ModelToRsponse(user), nil
}
