package handler

import (
	"mang_srv/model"
	"mang_srv/proto"
)

func ModelToRsponse(user model.User) *proto.UserInfo {
	userInfoRsp := proto.UserInfo{
		Id:       uint32(user.ID),
		Name:     user.Name,
		Mobile:   user.Mobile,
		Password: user.Password,
		Position: uint32(user.Position),
	}

	return &userInfoRsp
}

type UserServer struct {
	proto.UnimplementedUserServer
}
