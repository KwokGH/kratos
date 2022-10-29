package dao

import (
	"database/sql"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserGoodsNum struct {
	ID      string `gorm:"column:id;primary_key"`
	UserID  string `gorm:"column:user_id;NOT NULL"`
	GoodsID string `gorm:"column:goods_id;NOT NULL"`
	Num     int    `gorm:"column:num;default:0;NOT NULL"` // 购买数量
}

func (m *UserGoodsNum) TableName() string {
	return "user_goods_num"
}

type IUserGoodsNumDao interface {
	dbo.DataAccesser
}

type UserGoodsNumDao struct {
	dbo.BaseDA
}

var (
	_UserGoodsNumOnce sync.Once
	_UserGoodsNumDao  IUserGoodsNumDao
)

func GetUserGoodsNumDao() IUserGoodsNumDao {
	_UserGoodsNumOnce.Do(func() {
		_UserGoodsNumDao = &UserGoodsNumDao{}
	})
	return _UserGoodsNumDao
}

type UserGoodsNumCondition struct {
	UserID   sql.NullString
	GoodsID  sql.NullString
	GoodsIDs dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c UserGoodsNumCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.UserID.Valid {
		wheres = append(wheres, "user_id=?")
		params = append(params, c.UserID.String)
	}
	if c.GoodsID.Valid {
		wheres = append(wheres, "goods_id=?")
		params = append(params, c.GoodsID.String)
	}

	if c.GoodsIDs.Valid {
		wheres = append(wheres, "goods_id in (?)")
		params = append(params, c.GoodsIDs.Strings)
	}

	return wheres, params
}

func (c UserGoodsNumCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserGoodsNumCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
