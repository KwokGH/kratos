package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type OrderGoods struct {
	ID            string  `gorm:"column:id;primary_key"`
	OrderID       string  `gorm:"column:order_id;NOT NULL"`            // 订单id
	GoodsID       string  `gorm:"column:goods_id;NOT NULL"`            // 藏品id
	GoodsOptionID string  `gorm:"column:goods_option_id;NOT NULL"`     // 藏品编号id
	GoodsNoID     string  `gorm:"column:goods_no_id"`                  // 藏品编号id
	GoodsNo       string  `gorm:"column:goods_no"`                     // 藏品编号
	Price         float64 `gorm:"column:price;default:0.00;NOT NULL"`  // 单价
	Qty           int     `gorm:"column:qty;default:0;NOT NULL"`       // 购买数量
	Amount        float64 `gorm:"column:amount;default:0.00;NOT NULL"` // 总价
}

func (m *OrderGoods) TableName() string {
	return "order_goods"
}

type IOrderGoodsDao interface {
	dbo.DataAccesser
}

type OrderGoodsDao struct {
	dbo.BaseDA
}

var (
	_OrderGoodsOnce sync.Once
	_OrderGoodsDao  IOrderGoodsDao
)

func GetOrderGoodsDao() IOrderGoodsDao {
	_OrderGoodsOnce.Do(func() {
		_OrderGoodsDao = &OrderGoodsDao{}
	})
	return _OrderGoodsDao
}

type OrderGoodsCondition struct {
	OrderIDs dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c OrderGoodsCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.OrderIDs.Valid {
		wheres = append(wheres, "order_id in (?)")
		params = append(params, c.OrderIDs.Strings)
	}

	return wheres, params
}

func (c OrderGoodsCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c OrderGoodsCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
