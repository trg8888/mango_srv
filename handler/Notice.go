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

func (u *UserServer) AdminNoticeCreate(ctx context.Context, req *proto.NoticeInfoRequest) (*empty.Empty, error) {
	Notice := model.Notice{}
	Notice.Title = req.Title
	Notice.Describe = req.Describe
	if result := global.DB.Create(&Notice); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil
}

func (u *UserServer) AdminNoticeUpdate(ctx context.Context, req *proto.NoticeInfoRequest) (*empty.Empty, error) {
	Notice := model.Notice{}
	if res := global.DB.First(&Notice, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该图片")
	}
	Notice.Title = req.Title
	Notice.Describe = req.Describe
	if result := global.DB.Save(&Notice); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil
}

func (u *UserServer) AdminNoticeDelete(ctx context.Context, req *proto.NoticeInfoRequest) (*empty.Empty, error) {
	Notice := model.Notice{}
	if res := global.DB.First(&Notice, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该图片")
	}
	if result := global.DB.Delete(&Notice); result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "未知错误")
	}
	return &empty.Empty{}, nil
}

func (u *UserServer) AdminNoticeQueryId(ctx context.Context, req *proto.NoticeInfoRequestId) (*proto.NoticeInfoResponse, error) {
	Notice := model.Notice{}
	if res := global.DB.First(&Notice, req.Id); res.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "没有找到该图片")
	}

	return &proto.NoticeInfoResponse{
		Id:       uint32(Notice.ID),
		Title:    Notice.Title,
		Describe: Notice.Describe,
		Time:     Notice.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (u *UserServer) AdminNoticeQueryList(ctx context.Context, req *proto.NoticeInfoRequestList) (*proto.NoticeInfoResponseList, error) {
	var Notice []model.Notice
	res := global.DB.Scopes(model.Paginate(int(req.Page), int(req.Count))).Find(&Notice)
	if res.Error != nil {
		return nil, res.Error
	}

	var ImageListResponse []*proto.NoticeInfoResponse

	for _, value := range Notice {
		ImageListResponse = append(ImageListResponse, &proto.NoticeInfoResponse{
			Id:       uint32(value.ID),
			Title:    value.Title,
			Describe: value.Describe,
			Time:     value.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	var total int64
	global.DB.Model(&model.Notice{}).Count(&total)

	return &proto.NoticeInfoResponseList{
		Total: uint32(total),
		Data:  ImageListResponse,
	}, nil
}
