package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Order struct {
	ID             string    `gorm:"column:id;primary_key"`
	Time           time.Time `gorm:"column:time;default:CURRENT_TIMESTAMP;NOT NULL"` // 下单时间
	UserID         string    `gorm:"column:user_id;NOT NULL"`                        // 用户id
	PurchaseType   int       `gorm:"column:purchase_type;default:0;NOT NULL"`        // 购买类型 0-购买 1-兑换
	Qty            int       `gorm:"column:qty;default:0;NOT NULL"`                  // 总数量
	Amount         float64   `gorm:"column:amount;default:0.00;NOT NULL"`            // 支付金额
	Cdkey          string    `gorm:"column:cdkey"`                                   // 兑换码
	Status         int       `gorm:"column:status;default:0;NOT NULL"`               // 订单状态 0-进行中 1-已完成 3-已取消
	PaymentTime    time.Time `gorm:"column:payment_time"`                            // 支付时间
	PaymentType    int       `gorm:"column:payment_type;default:0;NOT NULL"`         // 支付类型 0-微信支付 1-通联支付
	PaymentTradeNo string    `gorm:"column:payment_trade_no"`                        // 支付流水号
	SaleQtyStat    int       `gorm:"column:sale_qty_stat;default:0;NOT NULL"`        // 销量统计标志，0-未统计 1-已统计
	GoodsID        string    `gorm:"column:goods_id"`                                // 购买藏品id 数据版本大于等于1有效
	GoodsNoID      string    `gorm:"column:goods_no_id"`                             // 藏品编号id 数据版本大于等于1有效
	GoodsNo        string    `gorm:"column:goods_no"`                                // 藏品编号 数据版本大于等于1有效
	UserGoodsID    string    `gorm:"column:user_goods_id"`                           // 订单成功后，生成的user_goods表数据项的id，数据版本大于等于1有效
	Price          float64   `gorm:"column:price;default:0.00;NOT NULL"`             // 购买单价 数据版本大于等于1有效
	Version        int       `gorm:"column:version;default:0;NOT NULL"`              // 数据版本，当前请填1
}

func (m *Order) TableName() string {
	return "order"
}

type OrderStatus int

const (
	OrderStatusGoing    OrderStatus = 0
	OrderStatusFinished OrderStatus = 1
	OrderStatusCanceled OrderStatus = 2
)

const (
	CouponIDLength = 10
)

type IOrderDao interface {
	dbo.DataAccesser
}

type OrderDao struct {
	dbo.BaseDA
}

var (
	_OrderOnce sync.Once
	_OrderDao  IOrderDao
)

func GetOrderDao() IOrderDao {
	_OrderOnce.Do(func() {
		_OrderDao = &OrderDao{}
	})
	return _OrderDao
}

type OrderCondition struct {
	Status      sql.NullInt64
	Statuss     dbo.NullInts
	UserID      sql.NullString
	SaleQtyStat sql.NullInt64
	GoodsID     sql.NullString
	GoodsIDs    dbo.NullStrings
	CDkey       sql.NullString
	OrderID     sql.NullString

	OrderBy int
	Pager   dbo.Pager
}

func (c OrderCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.Status.Valid {
		wheres = append(wheres, "status = ?")
		params = append(params, c.Status.Int64)
	}

	if c.Statuss.Valid {
		wheres = append(wheres, "status in (?)")
		params = append(params, c.Statuss.Ints)
	}

	if c.UserID.Valid {
		wheres = append(wheres, "user_id = ?")
		params = append(params, c.UserID.String)
	}

	if c.SaleQtyStat.Valid {
		wheres = append(wheres, "sale_qty_stat=?")
		params = append(params, c.SaleQtyStat.Int64)
	}

	if c.GoodsID.Valid {
		wheres = append(wheres, "goods_id = ?")
		params = append(params, c.GoodsID.String)
	}
	if c.GoodsIDs.Valid {
		wheres = append(wheres, "goods_id in (?)")
		params = append(params, c.GoodsIDs.Strings)
	}

	if c.CDkey.Valid {
		wheres = append(wheres, "cdkey=?")
		params = append(params, c.CDkey.String)
	}

	if c.OrderID.Valid {
		wheres = append(wheres, "order_id = ?")
		params = append(params, c.OrderID.String)
	}

	return wheres, params
}

func (c OrderCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c OrderCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
