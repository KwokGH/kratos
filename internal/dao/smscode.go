package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type SmsCode struct {
	ID     string    `gorm:"column:id;primary_key"`
	Time   time.Time `gorm:"column:time"`   // 发送时间
	Mobile string    `gorm:"column:mobile"` // 手机号
	Sms    string    `gorm:"column:sms"`    // 短消息
	Result string    `gorm:"column:result"` // 发送短信返回结果
}

func (m *SmsCode) TableName() string {
	return "sms_code"
}

type ISmsCodeDao interface {
	dbo.DataAccesser
}

type SmsCodeDao struct {
	dbo.BaseDA
}

var (
	_SmsCodeOnce sync.Once
	_SmsCodeDao  ISmsCodeDao
)

func GetSmsCodeDao() ISmsCodeDao {
	_SmsCodeOnce.Do(func() {
		_SmsCodeDao = &SmsCodeDao{}
	})
	return _SmsCodeDao
}

type SmsCodeCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c SmsCodeCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c SmsCodeCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c SmsCodeCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
