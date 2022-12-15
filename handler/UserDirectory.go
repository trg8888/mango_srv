package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
)

// UserCraterHomeDirectory 用户关联一级目录
func (u *UserServer) UserCraterHomeDirectory(ctx context.Context, req *proto.UserCraterHomeDirectoryID) (*empty.Empty, error) {
	var category model.HomeDirectory
	if res := global.DB.First(&category, req.Homedirectoryid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "一级目录分类不存在")
	}

	var user model.User
	if res := global.DB.First(&user, req.Userid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户不存在")
	}
	categoryBrand := model.HomeDirectoryCategoryDirectory{
		UserID:          int32(req.Userid),
		HomeDirectoryID: int32(req.Homedirectoryid),
	}
	if result := global.DB.Create(&categoryBrand); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}

	return &empty.Empty{}, nil

}

// UserCraterDirectory 用户关联二级目录
func (u *UserServer) UserCraterDirectory(ctx context.Context, req *proto.UserCraterDirectoryID) (*empty.Empty, error) {
	var HomeDirectory model.HomeDirectory
	if res := global.DB.First(&HomeDirectory, req.Homedirectoryid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "一级目录分类不存在")
	}

	var user model.User
	if res := global.DB.First(&user, req.Userid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "用户不存在")
	}

	var directory model.Directory
	if res := global.DB.First(&directory, req.DirectoryID); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "二级目录分类不存在")
	}

	categoryBrand := model.UserCategoryDirectory{
		UserID:          int32(req.Userid),
		HomeDirectoryID: int32(req.Homedirectoryid),
		DirectoryID:     int32(req.DirectoryID),
	}
	if result := global.DB.Create(&categoryBrand); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}

	return &empty.Empty{}, nil

}

// UserDeleteHomeDirectory 用户删除一级目录
func (u *UserServer) UserDeleteHomeDirectory(ctx context.Context, req *proto.UserCraterHomeDirectoryID) (*empty.Empty, error) {
	if res := global.DB.Where(&model.HomeDirectoryCategoryDirectory{UserID: int32(req.Userid), HomeDirectoryID: int32(req.Homedirectoryid)}).Delete(&model.UserCategoryDirectory{}); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	return &emptypb.Empty{}, nil
}

// UserDeleteDirectory 用户删除二级目录
func (u *UserServer) UserDeleteDirectory(ctx context.Context, req *proto.UserCraterDirectoryID) (*empty.Empty, error) {
	if res := global.DB.Where(&model.UserCategoryDirectory{UserID: int32(req.Userid), DirectoryID: int32(req.DirectoryID), HomeDirectoryID: int32(req.Homedirectoryid)}).Delete(&model.UserCategoryDirectory{}); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	return &emptypb.Empty{}, nil
}

// UserListDirectory 通过用户查询目录
func (u *UserServer) UserListHomeDirectory(ctx context.Context, req *proto.UserListDirectoryID) (*proto.DirectoryListResponse, error) {
	var HomeDirectorys []model.HomeDirectoryCategoryDirectory
	if res := global.DB.Preload("HomeDirectory").Where(&model.HomeDirectoryCategoryDirectory{UserID: int32(req.Userid)}).Find(&HomeDirectorys); res.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	} else {
		data := &proto.DirectoryListResponse{}
		var data_ []*proto.Result
		for _, value := range HomeDirectorys {
			var data__ = proto.Result{}
			var Directorys []model.UserCategoryDirectory
			if res := global.DB.Preload("Directory").Where(&model.UserCategoryDirectory{UserID: int32(req.Userid), HomeDirectoryID: value.HomeDirectoryID}).Find(&Directorys); res.Error != nil {
				return nil, status.Errorf(codes.Unknown, "未知错误")
			}

			data__.Name = value.HomeDirectory.Name
			data__.Path = value.HomeDirectory.Path
			if value.HomeDirectory.Icon != "" {
				data__.Icon = value.HomeDirectory.Icon
			}
			data__.IsTab = value.HomeDirectory.IsTab
			for _, value_ := range Directorys {
				data__.Result = append(data__.Result, &proto.Result{
					Name:  value_.Directory.Name,
					Path:  value_.Directory.Path,
					Icon:  value_.Directory.Icon,
					IsTab: value_.Directory.IsTab,
				})
			}
			data_ = append(data_, &data__)
		}
		data.Result = data_
		return data, nil
	}
}
