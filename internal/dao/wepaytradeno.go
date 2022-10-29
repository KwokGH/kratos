package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type WepayTradeNo struct {
	ID       string `gorm:"column:id;primary_key"`
	OrderID  string `gorm:"column:order_id;NOT NULL"`       // 订单id
	TradeNo  string `gorm:"column:trade_no;NOT NULL"`       // 支付编号
	Type     int    `gorm:"column:type;default:0;NOT NULL"` // 类型 0-jsapi 1-h5
	PrepayID string `gorm:"column:prepay_id"`
}

func (m *WepayTradeNo) TableName() string {
	return "wepay_trade_no"
}

type IWepayTradeNoDao interface {
	dbo.DataAccesser
}

type WepayTradeNoDao struct {
	dbo.BaseDA
}

var (
	_WepayTradeNoOnce sync.Once
	_WepayTradeNoDao  IWepayTradeNoDao
)

func GetWepayTradeNoDao() IWepayTradeNoDao {
	_WepayTradeNoOnce.Do(func() {
		_WepayTradeNoDao = &WepayTradeNoDao{}
	})
	return _WepayTradeNoDao
}

type WepayTradeNoCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c WepayTradeNoCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c WepayTradeNoCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c WepayTradeNoCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
