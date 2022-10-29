package biz

import (
	"context"
	pb "github.com/KwokGH/kratos/api/v1/user"
)

type IUserRepo interface {
	Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error)
}
