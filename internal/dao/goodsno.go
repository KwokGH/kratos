package dao

import (
	"context"
	"database/sql"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type GoodsNo struct {
	ID      string `gorm:"column:id;primary_key"`
	GoodsID string `gorm:"column:goods_id;NOT NULL"`       // 藏品id
	No      string `gorm:"column:no;NOT NULL"`             // 编号
	Used    int    `gorm:"column:used;default:0;NOT NULL"` // 是否使用 0-未使用 1-使用
}

func (m *GoodsNo) TableName() string {
	return "goods_no"
}

type IGoodsNoDao interface {
	dbo.DataAccesser

	UpdateUsedTx(ctx context.Context, tx *dbo.DBContext, goodsNoIDs []string, used int) error
}

type GoodsNoDao struct {
	dbo.BaseDA
}

var (
	_GoodsNoOnce sync.Once
	_GoodsNoDao  IGoodsNoDao
)

func GetGoodsNoDao() IGoodsNoDao {
	_GoodsNoOnce.Do(func() {
		_GoodsNoDao = &GoodsNoDao{}
	})
	return _GoodsNoDao
}

func (m *GoodsNoDao) UpdateUsedTx(ctx context.Context, tx *dbo.DBContext, goodsNoIDs []string, used int) error {
	tx.ResetCondition()

	return tx.Model(&GoodsNo{}).
		Where("id in (?)", goodsNoIDs).
		UpdateColumn("used", used).Error
}

type GoodsNoCondition struct {
	GoodsID sql.NullString
	Used    sql.NullInt64

	OrderBy int
	Pager   dbo.Pager
}

func (c GoodsNoCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.GoodsID.Valid {
		wheres = append(wheres, "goods_id = ?")
		params = append(params, c.GoodsID.String)
	}

	if c.Used.Valid {
		wheres = append(wheres, "used = ?")
		params = append(params, c.Used.Int64)
	}

	return wheres, params
}

func (c GoodsNoCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c GoodsNoCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
