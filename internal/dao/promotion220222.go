package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Promotion220222 struct {
	ID        string `gorm:"column:id;primary_key"` // 等同于user_id
	Goods106  int    `gorm:"column:goods106;default:0;NOT NULL"`
	Goods107  int    `gorm:"column:goods107;default:0;NOT NULL"`
	Goods108  int    `gorm:"column:goods108;default:0;NOT NULL"`
	Goods109  int    `gorm:"column:goods109;default:0;NOT NULL"`
	Goods110  int    `gorm:"column:goods110;default:0;NOT NULL"`
	Goods111  int    `gorm:"column:goods111;default:0;NOT NULL"`
	Status    int    `gorm:"column:status;default:0;NOT NULL"` // 状态 0-尚未达成 1-已达成
	Cdkey     string `gorm:"column:cdkey"`                     // 发放的兑换码
	Mobile    string `gorm:"column:mobile"`                    // 接收短信的手机号
	Sms       string `gorm:"column:sms"`                       // 发送的短信内容
	SmsResult string `gorm:"column:sms_result"`                // 短信反馈
}

func (m *Promotion220222) TableName() string {
	return "promotion220222"
}

type IPromotion220222Dao interface {
	dbo.DataAccesser
}

type Promotion220222Dao struct {
	dbo.BaseDA
}

var (
	_Promotion220222Once sync.Once
	_Promotion220222Dao  IPromotion220222Dao
)

func GetPromotion220222Dao() IPromotion220222Dao {
	_Promotion220222Once.Do(func() {
		_Promotion220222Dao = &Promotion220222Dao{}
	})
	return _Promotion220222Dao
}

type Promotion220222Condition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c Promotion220222Condition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c Promotion220222Condition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c Promotion220222Condition) GetPager() *dbo.Pager {
	return &c.Pager
}
