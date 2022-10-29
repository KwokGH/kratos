package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Promotion220528 struct {
	ID        string    `gorm:"column:id;primary_key"`
	Time      time.Time `gorm:"column:time"`
	UserID    string    `gorm:"column:user_id"`
	CouponID  string    `gorm:"column:coupon_id"` // 优惠券id
	Mobile    string    `gorm:"column:mobile"`    // 接收短信的手机号
	Sms       string    `gorm:"column:sms"`
	SmsResult string    `gorm:"column:sms_result"`
}

func (m *Promotion220528) TableName() string {
	return "promotion220528"
}

type IPromotion220528Dao interface {
	dbo.DataAccesser
}

type Promotion220528Dao struct {
	dbo.BaseDA
}

var (
	_Promotion220528Once sync.Once
	_Promotion220528Dao  IPromotion220528Dao
)

func GetPromotion220528Dao() IPromotion220528Dao {
	_Promotion220528Once.Do(func() {
		_Promotion220528Dao = &Promotion220528Dao{}
	})
	return _Promotion220528Dao
}

type Promotion220528Condition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c Promotion220528Condition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c Promotion220528Condition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c Promotion220528Condition) GetPager() *dbo.Pager {
	return &c.Pager
}
