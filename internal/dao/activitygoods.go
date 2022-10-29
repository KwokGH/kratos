package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type ActivityGoods struct {
	ID         string `gorm:"column:id;primary_key"`
	ActivityID string `gorm:"column:activity_id;NOT NULL"`   // 活动id
	GoodsID    string `gorm:"column:goods_id;NOT NULL"`      // 商品id
	Seq        int    `gorm:"column:seq;default:0;NOT NULL"` // 顺序值，越大越靠前，不能大于10000
}

func (m *ActivityGoods) TableName() string {
	return "activity_goods"
}

type IActivityGoodsDao interface {
	dbo.DataAccesser
}

type ActivityGoodsDao struct {
	dbo.BaseDA
}

var (
	_ActivityGoodsOnce sync.Once
	_ActivityGoodsDao  IActivityGoodsDao
)

func GetActivityGoodsDao() IActivityGoodsDao {
	_ActivityGoodsOnce.Do(func() {
		_ActivityGoodsDao = &ActivityGoodsDao{}
	})
	return _ActivityGoodsDao
}

type ActivityGoodsCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c ActivityGoodsCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c ActivityGoodsCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c ActivityGoodsCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
