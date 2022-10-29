package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Coupon struct {
	ID       string    `gorm:"column:id;primary_key"`
	Time     time.Time `gorm:"column:time;NOT NULL"`
	GoodsIds string    `gorm:"column:goods_ids"`                 // 关联商品，多个用逗号隔开
	Status   int       `gorm:"column:status;default:0;NOT NULL"` // 状态 0-未使用，1-已使用
}

func (m *Coupon) TableName() string {
	return "coupon"
}

type ICouponDao interface {
	dbo.DataAccesser
}

type CouponDao struct {
	dbo.BaseDA
}

var (
	_CouponOnce sync.Once
	_CouponDao  ICouponDao
)

func GetCouponDao() ICouponDao {
	_CouponOnce.Do(func() {
		_CouponDao = &CouponDao{}
	})
	return _CouponDao
}

type CouponCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c CouponCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c CouponCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c CouponCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
