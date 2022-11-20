package data

import (
	"context"
	"github.com/KwokGH/kratos/internal/biz"
	de2 "github.com/KwokGH/kratos/internal/entity/de"
	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u *userRepo) FindByPhone(ctx context.Context, phone string) (*de2.User, error) {
	return &de2.User{
		Base: de2.Base{
			ID: "123",
		},
		NickName: "",
		Phone:    "",
		Password: "",
	}, nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.IUserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
