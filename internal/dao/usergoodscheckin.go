package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserGoodsCheckin struct {
	ID      string                 `gorm:"column:id;primary_key"`
	Time    time.Time              `gorm:"column:time"`                      // 登记时间
	UserID  string                 `gorm:"column:user_id;NOT NULL"`          // 用户id
	GoodsID string                 `gorm:"column:goods_id;NOT NULL"`         // 商品id
	Status  UserGoodsCheckinStatus `gorm:"column:status;default:0;NOT NULL"` // 登记状态 0-未抽签或未抽中 1-抽中 2-付款超时 3-已付款
}

type UserGoodsCheckinStatus int

const (
	UserGoodsCheckinStatusNotWin  UserGoodsCheckinStatus = 0
	UserGoodsCheckinStatusWin     UserGoodsCheckinStatus = 1
	UserGoodsCheckinStatusTimeOut UserGoodsCheckinStatus = 2
	UserGoodsCheckinStatusFinish  UserGoodsCheckinStatus = 3
)

func (m *UserGoodsCheckin) TableName() string {
	return "user_goods_checkin"
}

type IUserGoodsCheckinDao interface {
	dbo.DataAccesser
}

type UserGoodsCheckinDao struct {
	dbo.BaseDA
}

var (
	_UserGoodsCheckinOnce sync.Once
	_UserGoodsCheckinDao  IUserGoodsCheckinDao
)

func GetUserGoodsCheckinDao() IUserGoodsCheckinDao {
	_UserGoodsCheckinOnce.Do(func() {
		_UserGoodsCheckinDao = &UserGoodsCheckinDao{}
	})
	return _UserGoodsCheckinDao
}

type UserGoodsCheckinCondition struct {
	UserId   sql.NullString
	GoodsId  sql.NullString
	GoodsIds dbo.NullStrings
	Status   sql.NullInt64

	OrderBy int
	Pager   dbo.Pager
}

func (c UserGoodsCheckinCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.GoodsId.Valid {
		wheres = append(wheres, "goods_id = ?")
		params = append(params, c.GoodsId.String)
	}

	if c.UserId.Valid {
		wheres = append(wheres, "user_id = ?")
		params = append(params, c.UserId.String)
	}

	if c.Status.Valid {
		wheres = append(wheres, "status = ?")
		params = append(params, c.Status.Int64)
	}

	if c.GoodsIds.Valid {
		wheres = append(wheres, "goods_id in (?)")
		params = append(params, c.GoodsIds.Strings)
	}

	return wheres, params
}

func (c UserGoodsCheckinCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserGoodsCheckinCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
