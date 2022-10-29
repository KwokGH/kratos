package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type DetailItem struct {
	ID      string `gorm:"column:id;NOT NULL"`
	Type    int    `gorm:"column:type;default:0;NOT NULL"` // 类型：0-文本，1-图片
	Content string `gorm:"column:content"`                 // 内容：类型0时为文本，类型1时为图片地址
}

func (m *DetailItem) TableName() string {
	return "detail_item"
}

type IDetailItemDao interface {
	dbo.DataAccesser
}

type DetailItemDao struct {
	dbo.BaseDA
}

var (
	_DetailItemOnce sync.Once
	_DetailItemDao  IDetailItemDao
)

func GetDetailItemDao() IDetailItemDao {
	_DetailItemOnce.Do(func() {
		_DetailItemDao = &DetailItemDao{}
	})
	return _DetailItemDao
}

type DetailItemCondition struct {
	IDs dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c DetailItemCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.IDs.Valid {
		wheres = append(wheres, "id IN (?)")
		params = append(params, c.IDs.Strings)
	}

	return wheres, params
}

func (c DetailItemCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c DetailItemCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
