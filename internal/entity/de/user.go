package de

import (
	"context"
	"github.com/KwokGH/kratos/internal/entity"
	"github.com/KwokGH/kratos/pkg/utils"
)

type User struct {
	Base

	NickName string `gorm:"nick_name"`
	Phone    string `gorm:"phone"`
	Password string `gorm:"password"`
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) CheckPassword(ctx context.Context, password string) error {
	pwd := utils.GetMd5(password)
	if u.Password == pwd {
		return nil
	}

	return entity.ErrInvalidArgs
}
