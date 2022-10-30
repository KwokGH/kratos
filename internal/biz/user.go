package biz

import (
	"context"
	"github.com/KwokGH/kratos/internal/entity"
)

type IUserRepo interface {
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
}
