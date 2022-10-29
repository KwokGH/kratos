package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type ITempDao interface {
	dbo.DataAccesser
}

type TempDao struct {
	dbo.BaseDA
}

var (
	_tempOnce sync.Once
	_tempDao  ITempDao
)

func GetTempDao() ITempDao {
	_tempOnce.Do(func() {
		_tempDao = &TempDao{}
	})
	return _tempDao
}

type TempCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c TempCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c TempCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c TempCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
