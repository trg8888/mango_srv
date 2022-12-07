package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
)

func (u *UserServer) UserResetPassword(ctx context.Context, req *proto.UserInfoResetPassword) (*proto.UserInfo, error) {
	var user model.User

	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.NewPassword, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	global.DB.Save(&user)
	ModelToRsponse(user)
	return ModelToRsponse(user), nil
}
