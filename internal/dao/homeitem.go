package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type HomeItem struct {
	ID               string `gorm:"column:id;primary_key"`
	Type             int    `gorm:"column:type;default:0;NOT NULL"`               // 首页项目类型，0-活动，1-藏品
	ItemID           string `gorm:"column:item_id"`                               // 项目id，藏品为藏品id，活动为活动id
	Pic              string `gorm:"column:pic"`                                   // 图片地址
	ShowRelatedGoods int    `gorm:"column:show_related_goods;default:1;NOT NULL"` // 是否显示相关藏品列表 0-不显示 1-显示
	Seq              int    `gorm:"column:seq;default:0;NOT NULL"`                // 顺序值，越大越靠前，不能大于10000
}

func (m *HomeItem) TableName() string {
	return "home_item"
}

type IHomeItemDao interface {
	dbo.DataAccesser
}

type HomeItemDao struct {
	dbo.BaseDA
}

var (
	_HomeItemOnce sync.Once
	_HomeItemDao  IHomeItemDao
)

func GetHomeItemDao() IHomeItemDao {
	_HomeItemOnce.Do(func() {
		_HomeItemDao = &HomeItemDao{}
	})
	return _HomeItemDao
}

type HomeItemCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c HomeItemCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c HomeItemCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return "seq desc, id desc"
	}
}

func (c HomeItemCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
