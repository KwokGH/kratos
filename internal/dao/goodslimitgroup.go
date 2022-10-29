package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type GoodsLimitGroup struct {
	ID          string `gorm:"column:id;primary_key"`
	GoodsIds    string `gorm:"column:goods_ids"`   // 多个id用英文逗号隔开，范围id使用英文减号号连接起始id和结束id
	Description string `gorm:"column:description"` // 描述信息，当用户购买被限制是提示
}

func (m *GoodsLimitGroup) TableName() string {
	return "goods_limit_group"
}

type IGoodsLimitGroupDao interface {
	dbo.DataAccesser
}

type GoodsLimitGroupDao struct {
	dbo.BaseDA
}

var (
	_GoodsLimitGroupOnce sync.Once
	_GoodsLimitGroupDao  IGoodsLimitGroupDao
)

func GetGoodsLimitGroupDao() IGoodsLimitGroupDao {
	_GoodsLimitGroupOnce.Do(func() {
		_GoodsLimitGroupDao = &GoodsLimitGroupDao{}
	})
	return _GoodsLimitGroupDao
}

type GoodsLimitGroupCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c GoodsLimitGroupCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c GoodsLimitGroupCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c GoodsLimitGroupCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
