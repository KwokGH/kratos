package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type ActivityClassify struct {
	ID    string `gorm:"column:id;primary_key"`
	Title string `gorm:"column:title;NOT NULL"`         // 标题
	Pic   string `gorm:"column:pic"`                    // 顶部图片地址
	Desc  string `gorm:"column:desc"`                   // 描述
	Seq   int    `gorm:"column:seq;default:0;NOT NULL"` // 顺序值，越大越靠前，必须小于10000
}

func (m *ActivityClassify) TableName() string {
	return "activity_classify"
}

type IActivityClassifyDao interface {
	dbo.DataAccesser
}

type ActivityClassifyDao struct {
	dbo.BaseDA
}

var (
	_ActivityClassifyOnce sync.Once
	_ActivityClassifyDao  IActivityClassifyDao
)

func GetActivityClassifyDao() IActivityClassifyDao {
	_ActivityClassifyOnce.Do(func() {
		_ActivityClassifyDao = &ActivityClassifyDao{}
	})
	return _ActivityClassifyDao
}

type ActivityClassifyCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c ActivityClassifyCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c ActivityClassifyCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c ActivityClassifyCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
