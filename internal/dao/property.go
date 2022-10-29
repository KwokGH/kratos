package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Property struct {
	ID          string `gorm:"column:id;primary_key"`
	Value       string `gorm:"column:value"`
	Description string `gorm:"column:description"`
}

var (
	MetaPassIDNext = "user_metapass_id_next"
)

func (m *Property) TableName() string {
	return "property"
}

type IPropertyDao interface {
	dbo.DataAccesser
}

type PropertyDao struct {
	dbo.BaseDA
}

var (
	_PropertyOnce sync.Once
	_PropertyDao  IPropertyDao
)

func GetPropertyDao() IPropertyDao {
	_PropertyOnce.Do(func() {
		_PropertyDao = &PropertyDao{}
	})
	return _PropertyDao
}

type PropertyCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c PropertyCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c PropertyCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c PropertyCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
