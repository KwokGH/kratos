package service

import (
	"context"
	v1 "github.com/KwokGH/kratos/api/v1"
	pb "github.com/KwokGH/kratos/api/v1/user"
	"github.com/KwokGH/kratos/internal/biz"
	"github.com/KwokGH/kratos/internal/entity/be"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	log *log.Helper

	uuc *biz.UserUseCase
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserService(uuc *biz.UserUseCase, logger log.Logger) *UserService {
	return &UserService{
		uuc: uuc,
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
	}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	input := &be.LoginInput{
		Phone:    req.Phone,
		Password: req.Password,
	}
	if input.Valid() != nil {
		return nil, v1.ErrorBadRequest("参数不正确")
	}

	token, err := s.uuc.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Token: token,
	}, nil
}
