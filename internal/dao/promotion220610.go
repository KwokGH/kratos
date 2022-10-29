package dao

import (
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Promotion220610 struct {
	ID          string    `gorm:"column:id;primary_key"` // 等同于user_id
	Time        time.Time `gorm:"column:time;NOT NULL"`
	UserGoodsID string    `gorm:"column:user_goods_id;NOT NULL"`
	GoodsID     string    `gorm:"column:goods_id;NOT NULL"`
	Amount      float64   `gorm:"column:amount;default:0.00;NOT NULL"`
}

func (m *Promotion220610) TableName() string {
	return "promotion220610"
}

type IPromotion220610Dao interface {
	dbo.DataAccesser
}

type Promotion220610Dao struct {
	dbo.BaseDA
}

var (
	_Promotion220610Once sync.Once
	_Promotion220610Dao  IPromotion220610Dao
)

func GetPromotion220610Dao() IPromotion220610Dao {
	_Promotion220610Once.Do(func() {
		_Promotion220610Dao = &Promotion220610Dao{}
	})
	return _Promotion220610Dao
}

type Promotion220610Condition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c Promotion220610Condition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c Promotion220610Condition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c Promotion220610Condition) GetPager() *dbo.Pager {
	return &c.Pager
}
