package service

import (
	"context"
	"fmt"
	pb "github.com/KwokGH/kratos/api/v1/user"
	"github.com/KwokGH/kratos/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

type UserService struct {
	pb.UnimplementedUserServer
	log *log.Helper

	userRepo biz.IUserRepo
}

func NewUserService(ur biz.IUserRepo, logger log.Logger) *UserService {
	return &UserService{
		userRepo: ur,
		log:      log.NewHelper(log.With(logger, "module", "service/interface")),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return s.userRepo.Login(ctx, req)
}

func (s *UserService) GetUserDetail(ctx context.Context, req *pb.GetUserDetailReq) (*pb.UserDetailReply, error) {
	t, ok := jwt.FromContext(ctx)
	if ok {
		fmt.Println(t.(*jwt2.MapClaims))
	} else {
		fmt.Println("未找到数据")
	}

	return &pb.UserDetailReply{
		UserId:   "999999",
		UserName: "Kwok",
	}, nil
}
