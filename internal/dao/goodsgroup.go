package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type GoodsGroup struct {
	ID        string    `gorm:"column:id;primary_key"`
	GroupName string    `gorm:"column:group_name;NOT NULL"` // 商品组名称
	Time      time.Time `gorm:"column:time;NOT NULL"`
}

func (m *GoodsGroup) TableName() string {
	return "goods_group"
}

type IGoodsGroupDao interface {
	dbo.DataAccesser
}

type GoodsGroupDao struct {
	dbo.BaseDA
}

var (
	_GoodsGroupOnce sync.Once
	_GoodsGroupDao  IGoodsGroupDao
)

func GetGoodsGroupDao() IGoodsGroupDao {
	_GoodsGroupOnce.Do(func() {
		_GoodsGroupDao = &GoodsGroupDao{}
	})
	return _GoodsGroupDao
}

type GoodsGroupCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c GoodsGroupCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c GoodsGroupCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c GoodsGroupCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
