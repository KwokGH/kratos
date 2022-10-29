package dao

import (
	"database/sql"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserAmountChangeMonth struct {
	ID          string  `gorm:"column:id;primary_key"`
	UserID      string  `gorm:"column:user_id;NOT NULL"`
	Month       string  `gorm:"column:month;default:0;NOT NULL"`          // 月份 格式yyyyMM
	Expenditure float64 `gorm:"column:expenditure;default:0.00;NOT NULL"` // 支出
	Income      float64 `gorm:"column:income;default:0.00;NOT NULL"`      // 收入
}

func (m *UserAmountChangeMonth) TableName() string {
	return "user_amount_change_month"
}

type IUserAmountChangeMonthDao interface {
	dbo.DataAccesser
}

type UserAmountChangeMonthDao struct {
	dbo.BaseDA
}

var (
	_UserAmountChangeMonthOnce sync.Once
	_UserAmountChangeMonthDao  IUserAmountChangeMonthDao
)

func GetUserAmountChangeMonthDao() IUserAmountChangeMonthDao {
	_UserAmountChangeMonthOnce.Do(func() {
		_UserAmountChangeMonthDao = &UserAmountChangeMonthDao{}
	})
	return _UserAmountChangeMonthDao
}

type UserAmountChangeMonthCondition struct {
	UserID  sql.NullString
	UserIDs dbo.NullStrings
	Months  dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c UserAmountChangeMonthCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.UserID.Valid {
		wheres = append(wheres, "`user_id`=?")
		params = append(params, c.UserID.String)
	}
	if c.UserIDs.Valid {
		wheres = append(wheres, "`user_id` in (?)")
		params = append(params, c.UserIDs.Strings)
	}

	if c.Months.Valid {
		wheres = append(wheres, "month in (?)")
		params = append(params, c.Months.Strings)
	}

	return wheres, params
}

func (c UserAmountChangeMonthCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserAmountChangeMonthCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
