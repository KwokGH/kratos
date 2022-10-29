package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Promotion220610Days struct {
	ID   string `gorm:"column:id;primary_key"`
	Num1 int    `gorm:"column:num1;default:0;NOT NULL"` // 388元参与人数
	Num2 int    `gorm:"column:num2;default:0;NOT NULL"` // 288元参与人数
	Num3 int    `gorm:"column:num3;default:0;NOT NULL"` // 188元参与人数
}

func (m *Promotion220610Days) TableName() string {
	return "promotion220610_days"
}

type IPromotion220610DaysDao interface {
	dbo.DataAccesser
}

type Promotion220610DaysDao struct {
	dbo.BaseDA
}

var (
	_Promotion220610DaysOnce sync.Once
	_Promotion220610DaysDao  IPromotion220610DaysDao
)

func GetPromotion220610DaysDao() IPromotion220610DaysDao {
	_Promotion220610DaysOnce.Do(func() {
		_Promotion220610DaysDao = &Promotion220610DaysDao{}
	})
	return _Promotion220610DaysDao
}

type Promotion220610DaysCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c Promotion220610DaysCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c Promotion220610DaysCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c Promotion220610DaysCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
