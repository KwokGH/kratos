package de

// 定义data层所需的entity

type Base struct {
	ID        string `gorm:"id"`
	CreateAt  int64  `gorm:"create_at"`
	UpdatedAt int64  `gorm:"updated_at"`
	DeletedAt int64  `gorm:"deleted_at"`
}
