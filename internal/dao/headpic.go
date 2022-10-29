package dao

import (
	"context"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type HeadPic struct {
	ID      string `gorm:"column:id;primary_key"`
	Url     string `gorm:"column:url;NOT NULL"`               // 图片地址
	Default int    `gorm:"column:default;default:0;NOT NULL"` // 是否为默认头像：0否 1是
	Seq     int    `gorm:"column:seq;default:0;NOT NULL"`     // 顺序，越大越靠前，不能大于10000
}

func (m *HeadPic) TableName() string {
	return "head_pic"
}

type IHeadPicDao interface {
	dbo.DataAccesser

	GetLastHeadPic(ctx context.Context) (*HeadPic, error)
}

type HeadPicDao struct {
	dbo.BaseDA
}

var (
	_HeadPicOnce sync.Once
	_HeadPicDao  IHeadPicDao
)

func GetHeadPicDao() IHeadPicDao {
	_HeadPicOnce.Do(func() {
		_HeadPicDao = &HeadPicDao{}
	})
	return _HeadPicDao
}

func (c *HeadPicDao) GetLastHeadPic(ctx context.Context) (*HeadPic, error) {
	tx := dbo.MustGetDB(ctx)
	tx.ResetCondition()

	var resp HeadPic
	err := tx.Last(&resp).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

type HeadPicCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c HeadPicCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c HeadPicCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c HeadPicCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
