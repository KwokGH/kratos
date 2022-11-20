package de

import "github.com/KwokGH/kratos/pkg/utils"

// 定义data层所需的entity

type Base struct {
	ID        string `gorm:"id"`
	CreatedAt int64  `gorm:"created_at"`
	UpdatedAt int64  `gorm:"updated_at"`
	DeletedAt int64  `gorm:"deleted_at"`
}

func NewBase() Base {
	return Base{ID: utils.NewID()}
}
