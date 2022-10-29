package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserCashApply struct {
	ID     string    `gorm:"column:id;primary_key"`
	Time   time.Time `gorm:"column:time;NOT NULL"`             // 申请时间
	UserID string    `gorm:"column:user_id;NOT NULL"`          // 用户id
	Amount int       `gorm:"column:amount;default:0;NOT NULL"` // 申请金额
	Status int       `gorm:"column:status;default:0;NOT NULL"` // 状态 0待审核 1审核通过 2审核不通过
}

func (m *UserCashApply) TableName() string {
	return "user_cash_apply"
}

type IUserCashApplyDao interface {
	dbo.DataAccesser
}

type UserCashApplyDao struct {
	dbo.BaseDA
}

var (
	_UserCashApplyOnce sync.Once
	_UserCashApplyDao  IUserCashApplyDao
)

func GetUserCashApplyDao() IUserCashApplyDao {
	_UserCashApplyOnce.Do(func() {
		_UserCashApplyDao = &UserCashApplyDao{}
	})
	return _UserCashApplyDao
}

type UserCashApplyCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c UserCashApplyCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c UserCashApplyCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserCashApplyCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
