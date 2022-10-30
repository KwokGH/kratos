package biz

import (
	"context"
	"fmt"
	"github.com/KwokGH/kratos/api/account"
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/KwokGH/kratos/pkg/utils"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AccountUseCase struct {
	authConfig *conf.Auth
	userRepo   IUserRepo
	logger     *log.Helper
}

func NewAccountUseCase(authConfig *conf.Auth, userRepo IUserRepo, logger log.Logger) *AccountUseCase {
	return &AccountUseCase{
		authConfig: authConfig,
		userRepo:   userRepo,
		logger:     log.NewHelper(logger),
	}
}

func (a *AccountUseCase) Login(ctx context.Context, req *account.LoginReq) (token string, err error) {
	// 校验参数
	if req.Phone == "" || req.Password == "" {
		return "", fmt.Errorf("登录失败")
	}
	// 获取用户信息
	user, err := a.userRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		return "", fmt.Errorf("登录失败：%w", err)
	}
	// 校验密码
	encrypt, err := utils.Encrypt([]byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("登录失败:%w", err)
	}
	if user.Password != string(encrypt) {
		return "", fmt.Errorf("登录失败")
	}
	// 生成token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.authConfig.GetExpireDuration().AsDuration())), // 设置token的过期时间
	})
	token, err = claims.SignedString([]byte(a.authConfig.GetJwtSecret()))
	if err != nil {
		a.logger.Errorf("登录失败，生成token失败：%v", err)
		return "", fmt.Errorf("登录失败")
	}
	return token, nil
}
