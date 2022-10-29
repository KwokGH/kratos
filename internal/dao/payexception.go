package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type PayException struct {
	ID        string    `gorm:"column:id;primary_key"`
	Time      time.Time `gorm:"column:time;NOT NULL"`
	Type      int       `gorm:"column:type;default:0;NOT NULL"` // 0-微信支付 1-通联支付
	ReturnMsg string    `gorm:"column:return_msg"`              // 支付返回消息
	OrderID   string    `gorm:"column:order_id"`                // 关联订单id
	UserID    string    `gorm:"column:user_id"`                 // 光联用户id
	Msg       string    `gorm:"column:msg"`                     // 异常消息
}

func (m *PayException) TableName() string {
	return "pay_exception"
}

type IPayExceptionDao interface {
	dbo.DataAccesser
}

type PayExceptionDao struct {
	dbo.BaseDA
}

var (
	_PayExceptionOnce sync.Once
	_PayExceptionDao  IPayExceptionDao
)

func GetPayExceptionDao() IPayExceptionDao {
	_PayExceptionOnce.Do(func() {
		_PayExceptionDao = &PayExceptionDao{}
	})
	return _PayExceptionDao
}

type PayExceptionCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c PayExceptionCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c PayExceptionCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c PayExceptionCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
