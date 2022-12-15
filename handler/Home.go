package handler

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mang_srv/global"
	"mang_srv/model"
	"mang_srv/proto"
	"time"
)

func (u *UserServer) UserHomeParametersRegister(ctx context.Context, req *proto.HomeParametersRequest) (*empty.Empty, error) {
	var user model.User
	result := global.DB.First(&user, req.UserId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	var homeparameters model.HomeParameters
	homeparameters.UserID = int32(req.UserId)
	homeparameters.IsTab = false
	homeparameters.SubTitle = req.SubTitle
	if req.SubValue != 0 {
		homeparameters.SubValue = req.SubValue
	}
	homeparameters.Title = req.Title
	if req.Uint != "" {
		homeparameters.Uint = req.Uint
	}
	if req.UintColor != "" {
		homeparameters.UintColor = req.UintColor
	}
	if req.Value != "" {
		homeparameters.Value = req.Value
	}

	global.DB.Create(&homeparameters)
	return &empty.Empty{}, nil
}

func (u *UserServer) UserHomeParametersById(ctx context.Context, req *proto.HomeParametersUserId) (*proto.HomeParametersListResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.UserId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	var homeparameterss []model.HomeParameters
	result = global.DB.Where(&model.HomeParameters{IsTab: false, UserID: int32(req.UserId)}).Find(&homeparameterss)
	if result.Error != nil {
		return nil, status.Errorf(codes.Unknown, "内部错误")
	}
	data := &proto.HomeParametersListResponse{
		Total: uint32(len(homeparameterss)),
	}
	for _, value := range homeparameterss {
		data_ := proto.HomeParametersResponse{
			Id:        uint32(value.ID),
			SubTitle:  value.SubTitle,
			SubValue:  value.SubValue,
			Title:     value.Title,
			Uint:      value.Uint,
			UintColor: value.UintColor,
			Value:     value.Value,
			UserId:    uint32(value.UserID),
		}
		data.Data = append(data.Data, &data_)
	}
	return data, nil
}

func (u *UserServer) UserHomeParametersChange(ctx context.Context, req *proto.HomeParametersResponse) (*empty.Empty, error) {
	var homeparameters model.HomeParameters
	result := global.DB.First(&homeparameters, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "卡片不存在")
	}

	if req.SubTitle != "" {
		homeparameters.SubTitle = req.SubTitle
	}
	if req.SubValue != 0 {
		homeparameters.SubValue = req.SubValue
	}
	if req.Title != "" {
		homeparameters.Title = req.Title
	}
	if req.Uint != "" {
		homeparameters.Uint = req.Uint
	}
	if req.UintColor != "" {
		homeparameters.UintColor = req.UintColor
	}
	if req.Value != "" {
		homeparameters.Value = req.Value
	}

	global.DB.Save(&homeparameters)
	return &empty.Empty{}, nil

}

func (u *UserServer) UserHomeParametersDelete(ctx context.Context, req *proto.HomeParametersId) (*empty.Empty, error) {
	var homeparameters model.HomeParameters
	result := global.DB.First(&homeparameters, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "卡片不存在")
	}
	global.DB.Delete(&homeparameters)
	return &empty.Empty{}, nil
}

func (u *UserServer) UserHomeParametersIsTab(ctx context.Context, req *proto.HomeParametersId) (*empty.Empty, error) {
	var homeparameters model.HomeParameters
	result := global.DB.First(&homeparameters, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "卡片不存在")
	}
	if homeparameters.IsTab {
		homeparameters.IsTab = false
	} else {
		homeparameters.IsTab = true
	}
	global.DB.Save(&homeparameters)
	return &empty.Empty{}, nil
}

// UserHomeUpdateLog 记录单个用户查询的
func (u *UserServer) UserHomeUpdateLog(ctx context.Context, req *proto.HomeUpdateLogRequest) (*proto.HomeUpdateLogListResponse, error) {
	var user model.User
	result := global.DB.First(&user, req.UserId)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	var homeUpdateLogs []model.HomeUpdateLog
	var homeUpdateLogs_ []model.HomeUpdateLog
	switch req.Query {
	case "week":
		global.DB.Where("user_id = ?", req.UserId).Where("add_time >= ?", time.Now().AddDate(0, 0, -7)).Order("add_time asc").Find(&homeUpdateLogs)

		if len(homeUpdateLogs) != 7 {
			for i := 0; i < 7; i++ {
				isTab := false
				TimeData := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
				for _, v := range homeUpdateLogs {
					if v.Data == TimeData {
						isTab = true
						homeUpdateLogs_ = append(homeUpdateLogs_, v)
					}
				}
				if !isTab {
					homeUpdateLogs_ = append(homeUpdateLogs_, model.HomeUpdateLog{
						UserID: int32(req.UserId),
						Amount: 0,
						Data:   TimeData,
					})
				}
			}
		} else {
			homeUpdateLogs_ = homeUpdateLogs
		}

	case "month":
		global.DB.Where("user_id = ?", req.UserId).Where("add_time >= ?", time.Now().AddDate(0, -1, 0)).Order("add_time asc").Find(&homeUpdateLogs)

		if len(homeUpdateLogs) != 30 {
			for i := 0; i < 30; i++ {
				isTab := false
				TimeData := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
				for _, v := range homeUpdateLogs {
					if v.Data == TimeData {
						isTab = true
						homeUpdateLogs_ = append(homeUpdateLogs_, v)
					}
				}
				if !isTab {
					homeUpdateLogs_ = append(homeUpdateLogs_, model.HomeUpdateLog{
						UserID: int32(req.UserId),
						Amount: 0,
						Data:   TimeData,
					})
				}
			}
		} else {
			homeUpdateLogs_ = homeUpdateLogs
		}

	case "year":
		global.DB.Where("user_id = ?", req.UserId).Where("add_time >= ?", time.Now().AddDate(-1, 0, 0)).Order("add_time asc").Find(&homeUpdateLogs)
		if len(homeUpdateLogs) != 365 {
			for i := 0; i < 365; i++ {
				isTab := false
				TimeData := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
				for _, v := range homeUpdateLogs {
					if v.Data == TimeData {
						isTab = true
						homeUpdateLogs_ = append(homeUpdateLogs_, v)
					}
				}
				if !isTab {
					homeUpdateLogs_ = append(homeUpdateLogs_, model.HomeUpdateLog{
						UserID: int32(req.UserId),
						Amount: 0,
						Data:   TimeData,
					})
				}
			}
		} else {
			homeUpdateLogs_ = homeUpdateLogs
		}

	default:
		global.DB.Where("user_id = ?", req.UserId).Where("add_time >= ?", time.Now().AddDate(0, 0, -7)).Order("add_time asc").Find(&homeUpdateLogs)
		if len(homeUpdateLogs) != 7 {
			for i := 0; i < 7; i++ {
				isTab := false
				TimeData := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
				for _, v := range homeUpdateLogs {
					if v.Data == TimeData {
						isTab = true
						homeUpdateLogs_ = append(homeUpdateLogs_, v)
					}
				}
				if !isTab {
					homeUpdateLogs_ = append(homeUpdateLogs_, model.HomeUpdateLog{
						UserID: int32(req.UserId),
						Amount: 0,
						Data:   TimeData,
					})
				}
			}
		} else {
			homeUpdateLogs_ = homeUpdateLogs
		}

	}

	var data_ []*proto.HomeUpdateLogResponse
	for _, v := range homeUpdateLogs_ {
		data_ = append(data_, &proto.HomeUpdateLogResponse{
			Id:     uint32(v.ID),
			UserId: uint32(v.UserID),
			Amount: uint32(v.Amount),
			Data:   v.Data,
		})
	}
	data := proto.HomeUpdateLogListResponse{
		Total: uint32(len(homeUpdateLogs_)),
		Data:  data_,
	}

	return &data, nil

}
