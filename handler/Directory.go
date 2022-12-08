package Directory

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"mang_srv/proto"
)

func (u *handler.UserServer) DirectoryRegister(ctx context.Context, in *proto.DirectoryInfoRequest) (*empty.Empty, error) {
}
