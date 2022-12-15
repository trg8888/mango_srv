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

// DirectoryHomeRegister 创建一级目录
func (u *UserServer) DirectoryHomeRegister(ctx context.Context, req *proto.DirectoryInfoRequest) (*empty.Empty, error) {
	var directory model.HomeDirectory
	result := global.DB.Where(&model.Directory{Name: req.Name}).First(&directory)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "目录已存在")
	}
	directory.Name = req.Name
	directory.Path = req.Path
	if req.Icon != "" {
		directory.Icon = req.Icon
	}
	if result := global.DB.Create(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}

	return &empty.Empty{}, nil

}

// DirectoryRegister 创建二级目录
func (u *UserServer) DirectoryRegister(ctx context.Context, req *proto.DirectoryInfoRequest) (*empty.Empty, error) {
	var directory model.Directory
	result := global.DB.Where(&model.Directory{Name: req.Name}).First(&directory)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "目录已存在")
	}
	directory.Name = req.Name
	directory.Path = req.Path
	if req.Icon != "" {
		directory.Icon = req.Icon
	}
	if result := global.DB.Create(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}

	return &empty.Empty{}, nil

}

// DirectoryHomeChange 修改一级目录
func (u *UserServer) DirectoryHomeChange(ctx context.Context, req *proto.DirectoryInfoRequest) (*empty.Empty, error) {
	var directory model.HomeDirectory
	result := global.DB.First(&directory, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	directory.Name = req.Name
	directory.Path = req.Path
	if req.Icon != "" {
		directory.Icon = req.Icon
	}
	if result := global.DB.Save(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil

}

// DirectoryChange 修改二级目录
func (u *UserServer) DirectoryChange(ctx context.Context, req *proto.DirectoryInfoRequest) (*empty.Empty, error) {
	var directory model.Directory
	result := global.DB.First(&directory, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	directory.Name = req.Name
	directory.Path = req.Path
	if req.Icon != "" {
		directory.Icon = req.Icon
	}
	if result := global.DB.Save(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil

}

// DirectoryHomeList 查询一级所有目录
func (u *UserServer) DirectoryHomeList(ctx context.Context, in *empty.Empty) (*proto.DirectoryListIDResponse, error) {
	var directorys []model.HomeDirectory
	if result := global.DB.Find(&directorys); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	} else {
		var DirectoryInfoRequest []*proto.DirectoryInfoRequest
		for _, value := range directorys {
			DirectoryInfoRequest = append(DirectoryInfoRequest, &proto.DirectoryInfoRequest{
				Name: value.Name,
				Icon: value.Icon,
				Path: value.Path,
			})
		}

		directorylist := proto.DirectoryListIDResponse{
			Total: uint32(result.RowsAffected),
			Data:  DirectoryInfoRequest,
		}

		return &directorylist, nil
	}

}

// DirectoryList 查询二级所有目录
func (u *UserServer) DirectoryList(ctx context.Context, in *empty.Empty) (*proto.DirectoryListIDResponse, error) {
	var directorys []model.Directory
	if result := global.DB.Find(&directorys); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	} else {
		var DirectoryInfoRequest []*proto.DirectoryInfoRequest
		for _, value := range directorys {
			DirectoryInfoRequest = append(DirectoryInfoRequest, &proto.DirectoryInfoRequest{
				Name: value.Name,
				Icon: value.Icon,
				Path: value.Path,
			})
		}

		directorylist := proto.DirectoryListIDResponse{
			Total: uint32(result.RowsAffected),
			Data:  DirectoryInfoRequest,
		}

		return &directorylist, nil
	}

}

// DirectoryDelete 删除一级ID目录
func (u *UserServer) DirectoryHomeDelete(ctx context.Context, req *proto.DirectoryDeleteID) (*empty.Empty, error) {
	var directory model.HomeDirectory
	result := global.DB.First(&directory, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	if result := global.DB.Delete(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil
}

// DirectoryDelete 删除二级ID目录
func (u *UserServer) DirectoryDelete(ctx context.Context, req *proto.DirectoryDeleteID) (*empty.Empty, error) {
	var directory model.Directory
	result := global.DB.First(&directory, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}
	if result := global.DB.Delete(&directory); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil
}
