package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type WtsTradeNo struct {
	ID      string           `gorm:"column:id;primary_key"`
	Time    time.Time        `gorm:"column:time;NOT NULL"`
	OrderID string           `gorm:"column:order_id;NOT NULL"`
	TradeNo string           `gorm:"column:trade_no;NOT NULL"`
	Type    int              `gorm:"column:type;default:0;NOT NULL"`   // 0-微信jsapi支付，1-微信扫码字符，2-微信小程序字符，3-支付宝扫码支付
	PayInfo string           `gorm:"column:pay_info"`                  // 支付串
	Status  WtsTradeNoStatus `gorm:"column:status;default:0;NOT NULL"` // 0-待支付 1-支付成功 2-支付失败
}

func (m *WtsTradeNo) TableName() string {
	return "wts_trade_no"
}

type WtsTradeNoStatus int

const (
	WtsTradeNoStatusPaying  WtsTradeNoStatus = 0
	WtsTradeNoStatusSuccess WtsTradeNoStatus = 1
	WtsTradeNoStatusFail    WtsTradeNoStatus = 2
)

type IWtsTradeNoDao interface {
	dbo.DataAccesser
}

type WtsTradeNoDao struct {
	dbo.BaseDA
}

var (
	_WtsTradeNoOnce sync.Once
	_WtsTradeNoDao  IWtsTradeNoDao
)

func GetWtsTradeNoDao() IWtsTradeNoDao {
	_WtsTradeNoOnce.Do(func() {
		_WtsTradeNoDao = &WtsTradeNoDao{}
	})
	return _WtsTradeNoDao
}

type WtsTradeNoCondition struct {
	Type    sql.NullInt64
	OrderID sql.NullString
	Status  sql.NullInt64

	OrderBy int
	Pager   dbo.Pager
}

func (c WtsTradeNoCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.Type.Valid {
		wheres = append(wheres, "type = ?")
		params = append(params, c.Type.Int64)
	}

	if c.OrderID.Valid {
		wheres = append(wheres, "order_id = ?")
		params = append(params, c.OrderID.String)
	}

	if c.Status.Valid {
		wheres = append(wheres, "status = ?")
		params = append(params, c.Status.Int64)
	}

	return wheres, params
}

func (c WtsTradeNoCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c WtsTradeNoCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
