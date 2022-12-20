package handler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
)

func (u *UserServer) UserImageUpdate(ctx context.Context, req *proto.ImageRequest) (*empty.Empty, error) {

	var image model.Image
	result := global.DB.Where(&model.Image{Level: 2}).First(&image, req.Userid)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "目录不存在")
	}

	if uint32(image.ID) != req.Userid {
		return nil, status.Errorf(codes.PermissionDenied, "用户跟图片归属不一致")
	}
	if image.Level != 2 {
		return nil, status.Errorf(codes.Aborted, "图片层级错误")
	}
	var newimage model.Image
	newimage.ImageID = int32(req.Userid)
	newimage.Level = 3
	newimage.Name = req.Name
	newimage.UserID = int32(req.Id)
	newimage.URL = req.Url

	if result := global.DB.Create(&newimage); result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "内部错误")
	}

	return &empty.Empty{}, nil
}

func (u *UserServer) UserImageCreateName(ctx context.Context, req *proto.ImageRequest) (*proto.ImageResponse, error) {
	var image model.Image
	result := global.DB.Where(&model.Image{Level: 1, UserID: int32(req.Userid)}).First(&image)
	if result.RowsAffected == 0 {
		image.UserID = int32(req.Userid)
		image.Name = fmt.Sprintf("%d", req.Userid)
		image.Level = 1
		global.DB.Omit("ImageID").Create(&image)
	}
	var createimage model.Image
	createimage.Name = req.Name
	createimage.UserID = int32(req.Userid)
	createimage.ImageID = image.ID
	createimage.Level = 2
	global.DB.Create(&createimage)

	return &proto.ImageResponse{
		Id:     uint32(createimage.ID),
		Name:   createimage.Name,
		Userid: uint32(createimage.UserID),
		Level:  uint32(createimage.Level),
	}, nil
}

func (u *UserServer) UserImageQueryId(ctx context.Context, req *proto.ImageRequestId) (*proto.ImageListResponse, error) {
	var category model.Image

	if res := global.DB.Where(&model.Image{Level: 1, UserID: int32(req.Userid)}).First(&category); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户暂时没有图片")
	}

	var image []model.Image
	res := global.DB.Where(&model.Image{Level: 2, UserID: int32(req.Userid)}).Scopes(model.Paginate(int(req.Page), int(req.Count))).Find(&image)
	if res.Error != nil {
		return nil, res.Error
	}

	var total int64
	global.DB.Model(&model.Image{}).Where(&model.Image{Level: 2, UserID: int32(req.Userid)}).Count(&total)

	var ImageListResponse []*proto.ImageResponse

	for _, value := range image {
		ImageListResponse = append(ImageListResponse, &proto.ImageResponse{
			Id:     uint32(value.ID),
			Name:   value.Name,
			Userid: uint32(value.UserID),
			Level:  uint32(value.Level),
		})
	}

	return &proto.ImageListResponse{
		Total: uint32(total),
		Data:  ImageListResponse,
	}, nil
}

func (u *UserServer) UserImageQueryImageId(ctx context.Context, req *proto.ImageRequestId) (*proto.ImageListResponse, error) {
	// Userid 当id 使用

	var category model.Image
	if res := global.DB.Where(&model.Image{Level: 2}).First(&category, req.Userid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该分类")
	}

	var image []model.Image
	res := global.DB.Where(&model.Image{Level: 3, ImageID: category.ID}).Scopes(model.Paginate(int(req.Page), int(req.Count))).Find(&image)
	if res.Error != nil {
		return nil, res.Error
	}

	var ImageListResponse []*proto.ImageResponse

	for _, value := range image {
		ImageListResponse = append(ImageListResponse, &proto.ImageResponse{
			Id:     uint32(value.ID),
			Name:   value.Name,
			Url:    value.URL,
			Userid: uint32(value.UserID),
			Level:  uint32(value.Level),
		})
	}

	var total int64
	global.DB.Model(&model.Image{}).Where(&model.Image{Level: 3, ImageID: category.ID}).Count(&total)

	return &proto.ImageListResponse{
		Total: uint32(total),
		Data:  ImageListResponse,
	}, nil

}

func (u *UserServer) UserImageId(ctx context.Context, req *proto.ImageRequestId) (*empty.Empty, error) {

	// Userid 当id 使用
	var category model.Image
	if res := global.DB.Where(&model.Image{Level: 3}).First(&category, req.Userid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该图片")
	}
	global.DB.Delete(&category)
	return &empty.Empty{}, nil
}

func (u *UserServer) UserImageNameId(ctx context.Context, req *proto.ImageRequestId) (*empty.Empty, error) {
	// Userid 当id 使用

	var category model.Image
	if res := global.DB.Preload("SubImage").Where(&model.Image{Level: 2}).First(&category, req.Userid); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该分类")
	}
	var comments []model.Image
	global.DB.Where(&model.Image{Level: 3, ImageID: int32(req.Userid)}).Find(&comments)
	global.DB.Delete(&comments)
	global.DB.Delete(&category)

	return &empty.Empty{}, nil

}

func (u *UserServer) UserImageResetId(ctx context.Context, req *proto.ImageRequest) (*empty.Empty, error) {
	var category model.Image
	if res := global.DB.First(&category, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "未找到")
	}
	category.Name = req.Name
	if result := global.DB.Save(&category); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, result.Error.Error())
	}
	return &empty.Empty{}, nil
}
