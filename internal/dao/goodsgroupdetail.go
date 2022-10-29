package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type GoodsGroupDetail struct {
	ID           string    `gorm:"column:id;primary_key"`
	GoodsGroupID string    `gorm:"column:goods_group_id;NOT NULL"` // 商品组id
	GoodsID      string    `gorm:"column:goods_id;NOT NULL"`       // 商品id
	Time         time.Time `gorm:"column:time;NOT NULL"`
}

func (m *GoodsGroupDetail) TableName() string {
	return "goods_group_detail"
}

type IGoodsGroupDetailDao interface {
	dbo.DataAccesser
}

type GoodsGroupDetailDao struct {
	dbo.BaseDA
}

var (
	_GoodsGroupDetailOnce sync.Once
	_GoodsGroupDetailDao  IGoodsGroupDetailDao
)

func GetGoodsGroupDetailDao() IGoodsGroupDetailDao {
	_GoodsGroupDetailOnce.Do(func() {
		_GoodsGroupDetailDao = &GoodsGroupDetailDao{}
	})
	return _GoodsGroupDetailDao
}

type GoodsGroupDetailCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c GoodsGroupDetailCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c GoodsGroupDetailCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c GoodsGroupDetailCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
