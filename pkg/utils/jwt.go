package utils

import (
	"github.com/KwokGH/kratos/internal/conf"
	"github.com/golang-jwt/jwt/v4"
)

type LoginClaim struct {
	jwt.RegisteredClaims
	UserID string
}

func (lc *LoginClaim) Token(authConfig *conf.Auth) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, lc)
	return claims.SignedString([]byte(authConfig.GetJwtSecret()))
}

//jwt.RegisteredClaims{
//ExpiresAt: jwt.NewNumericDate(time.Now().Add(authConfig.GetExpireDuration().AsDuration())), // 设置token的过期时间
//}
