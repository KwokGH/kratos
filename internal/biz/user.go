package biz

import (
	"context"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/internal/entity"
	"github.com/KwokGH/kratos/internal/entity/be"
	"github.com/KwokGH/kratos/internal/entity/de"
	"github.com/KwokGH/kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

type IUserRepo interface {
	FindByPhone(ctx context.Context, phone string) (*de.User, error)
	CreateUser(ctx context.Context, input *be.RegisterInput) (string, error)
}

type UserUseCase struct {
	authConfig *conf.Auth
	userRepo   IUserRepo
	logger     *log.Helper
}

func NewUserUseCase(authConfig *conf.Auth, userRepo IUserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{
		authConfig: authConfig,
		userRepo:   userRepo,
		logger:     log.NewHelper(logger),
	}
}

func (a *UserUseCase) Login(ctx context.Context, input *be.LoginInput) (token string, err error) {
	userInfo, err := a.userRepo.FindByPhone(ctx, input.Phone)
	if err != nil {
		if err == entity.ErrRecordNotFound {
			return "", entity.ErrRecordNotFound
		}
		return "", err
	}

	if err := userInfo.CheckPassword(ctx, input.Password); err != nil {
		return "", err
	}

	claim := &utils.LoginClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.authConfig.GetExpireDuration().AsDuration())),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Kwok",
		},
		UserID: userInfo.ID,
	}
	token, err = claim.Token(a.authConfig)
	if err != nil {
		a.logger.Errorf("登录失败，生成token失败：%v", err)
		return "", err
	}

	return token, nil
}

func (a *UserUseCase) Register(ctx context.Context, input *be.RegisterInput) (string, error) {
	if strings.TrimSpace(input.NickName) == "" {
		input.NickName = "昵称"
	}

	userID, err := a.userRepo.CreateUser(ctx, input)
	if err != nil {
		return "", err
	}

	return userID, nil
}
