package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
	"strings"
)

func (u *UserServer) UserRegister(ctx context.Context, req *proto.UserInfoRegister) (*empty.Empty, error) {

	var user model.User

	result := global.DB.Where(&model.User{Name: req.Name}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	result = global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "手机号码已存在")
	}
	user.Name = req.Name
	user.Mobile = req.Mobile
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	if req.Position != 0 {
		user.Position = int32(req.Position)
	}
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil

}

func (u *UserServer) UserCheckPassword(ctx context.Context, req *proto.UserInfoCheckPassword) (*proto.UserInfoCheckResponse, error) {
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	return &proto.UserInfoCheckResponse{Success: check}, nil
}
