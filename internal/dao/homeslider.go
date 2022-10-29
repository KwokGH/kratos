package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type HomeSlider struct {
	ID      string `gorm:"column:id;primary_key"`
	Type    int    `gorm:"column:type;default:0;NOT NULL"` // 首页项目类型，0-活动，1-藏品，2-网址
	ItemID  string `gorm:"column:item_id"`                 // 项目id，藏品为藏品id，活动为活动id
	Pic     string `gorm:"column:pic"`                     // 图片地址
	LinkUrl string `gorm:"column:link_url"`                // 链接地址
	Seq     int    `gorm:"column:seq;default:0;NOT NULL"`  // 顺序值，越大越靠前
}

func (m *HomeSlider) TableName() string {
	return "home_slider"
}

type IHomeSliderDao interface {
	dbo.DataAccesser
}

type HomeSliderDao struct {
	dbo.BaseDA
}

var (
	_HomeSliderOnce sync.Once
	_HomeSliderDao  IHomeSliderDao
)

func GetHomeSliderDao() IHomeSliderDao {
	_HomeSliderOnce.Do(func() {
		_HomeSliderDao = &HomeSliderDao{}
	})
	return _HomeSliderDao
}

type HomeSliderCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c HomeSliderCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c HomeSliderCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return "seq desc, id desc"
	}
}

func (c HomeSliderCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
