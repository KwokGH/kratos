package be

import (
	"github.com/KwokGH/kratos/internal/entity"
	"strings"
)

type LoginInput struct {
	Phone    string
	Password string
}

func (input *LoginInput) Valid() error {
	if strings.TrimSpace(input.Phone) == "" || input.Password == "" {
		return entity.ErrInvalidArgs
	}

	return nil
}

type LoginOutput struct {
	Token string
}

type RegisterInput struct {
	NickName string
	Phone    string
	Password string
}

func (input *RegisterInput) Valid() error {
	if strings.TrimSpace(input.Phone) == "" || input.Password == "" {
		return entity.ErrInvalidArgs
	}

	return nil
}
