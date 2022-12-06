package handler

import (
	"context"
	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
	"mang_srv/proto"
	"strings"
)

//func (u *UserServer) UserLogin(context.Context, *proto.UserInfoLogin) (*proto.UserInfo, error) {
//
//}

func (u *UserServer) UserCheckPassword(ctx context.Context, req *proto.UserInfoCheckPassword) (*proto.UserInfoCheckResponse, error) {
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.UserInfoCheckResponse{Success: check}, nil
}
