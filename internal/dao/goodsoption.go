package dao

import (
	"context"
	"database/sql"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
	"gorm.io/gorm"
)

type GoodsOption struct {
	ID      string            `gorm:"column:id;primary_key"`
	GoodsID string            `gorm:"column:goods_id;NOT NULL"`           // 商品id
	Name    string            `gorm:"column:name;NOT NULL"`               // 名称（编号）
	Price   float64           `gorm:"column:price;default:0.00;NOT NULL"` // 价格
	Qty     int               `gorm:"column:qty;default:0;NOT NULL"`      // 总数量
	LockQty int               `gorm:"column:lock_qty;default:0;NOT NULL"` // 锁定数量
	SoldQty int               `gorm:"column:sold_qty;default:0;NOT NULL"` // 卖出数量
	Status  GoodsOptionStatus `gorm:"column:status;default:0;NOT NULL"`   // 状态： 0-即将发售 1-发售中（或可兑换） 3-已售罄（或已兑换）
	ChainID string            `gorm:"column:chain_id"`                    // 区块链id，系统自动生成
	Seq     int               `gorm:"column:seq;default:0;NOT NULL"`      // 顺序值，越大越靠前
}

func (m *GoodsOption) TableName() string {
	return "goods_option"
}

type GoodsOptionStatus int

const (
	GoodsOptionStatusUpComing GoodsOptionStatus = 0
	GoodsOptionStatusOnSale   GoodsOptionStatus = 1
	GoodsOptionStatusLock     GoodsOptionStatus = 2
	GoodsOptionStatusSoldOut  GoodsOptionStatus = 3
)

type IGoodsOptionDao interface {
	dbo.DataAccesser

	LockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsOptionID string, lockQty int) error
	UnlockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsOptionID string, qty int) error
}

type GoodsOptionDao struct {
	dbo.BaseDA
}

var (
	_GoodsOptionOnce sync.Once
	_GoodsOptionDao  IGoodsOptionDao
)

func GetGoodsOptionDao() IGoodsOptionDao {
	_GoodsOptionOnce.Do(func() {
		_GoodsOptionDao = &GoodsOptionDao{}
	})
	return _GoodsOptionDao
}

func (c GoodsOptionDao) LockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsOptionID string, lockQty int) error {
	tx.ResetCondition()

	return tx.Model(&GoodsOption{}).
		Where("id=? and lock_qty+sold_qty+? <= qty", goodsOptionID, lockQty).
		UpdateColumn("lock_qty", gorm.Expr("lock_qty + ?", lockQty)).
		Error
}

func (c *GoodsOptionDao) UnlockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsOptionID string, qty int) error {
	tx.ResetCondition()

	return tx.Model(&GoodsOption{}).
		Where("id=? and lock_qty-? >= 0", goodsOptionID, qty).
		UpdateColumn("lock_qty", gorm.Expr("lock_qty - ?", qty)).
		Error
}

type GoodsOptionCondition struct {
	IDs       dbo.NullStrings
	GoodsIDs  dbo.NullStrings
	NotStatus sql.NullInt64

	OrderBy int
	Pager   dbo.Pager
}

func (c GoodsOptionCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.GoodsIDs.Valid {
		wheres = append(wheres, "goods_id in (?)")
		params = append(params, c.GoodsIDs.Strings)
	}

	if c.NotStatus.Valid {
		wheres = append(wheres, "status != ?")
		params = append(params, c.NotStatus.Int64)
	}

	if c.IDs.Valid {
		wheres = append(wheres, "id in (?)")
		params = append(params, c.IDs.Strings)
	}

	return wheres, params
}

func (c GoodsOptionCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c GoodsOptionCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
