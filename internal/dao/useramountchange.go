package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserAmountChange struct {
	ID           string    `gorm:"column:id;primary_key"`
	Time         time.Time `gorm:"column:time;NOT NULL"`           // 变动时间
	UserID       string    `gorm:"column:user_id;NOT NULL"`        // 用户id
	Type         int       `gorm:"column:type;default:0;NOT NULL"` // 类型 0藏品出售或回购 1藏品购买 2提现
	Amount       float64   `gorm:"column:amount;default:0.00;NOT NULL"`
	RemainAmount float64   `gorm:"column:remain_amount;default:0.00;NOT NULL"` // 变动后余额
	GoodsID      string    `gorm:"column:goods_id"`                            // 关联的藏品id
	Description  string    `gorm:"column:description"`                         // 变动描述
}

func (m *UserAmountChange) TableName() string {
	return "user_amount_change"
}

type IUserAmountChangeDao interface {
	dbo.DataAccesser
}

type UserAmountChangeDao struct {
	dbo.BaseDA
}

var (
	_UserAmountChangeOnce sync.Once
	_UserAmountChangeDao  IUserAmountChangeDao
)

func GetUserAmountChangeDao() IUserAmountChangeDao {
	_UserAmountChangeOnce.Do(func() {
		_UserAmountChangeDao = &UserAmountChangeDao{}
	})
	return _UserAmountChangeDao
}

type UserAmountChangeCondition struct {
	UserID sql.NullString
	GtID   sql.NullString

	OrderBy int
	Pager   dbo.Pager
}

func (c UserAmountChangeCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.UserID.Valid {
		wheres = append(wheres, "`user_id`=?")
		params = append(params, c.UserID.String)
	}

	if c.GtID.Valid {
		wheres = append(wheres, "id > ?")
		params = append(params, c.GtID.String)
	}

	return wheres, params
}

func (c UserAmountChangeCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return "id asc"
	}
}

func (c UserAmountChangeCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
