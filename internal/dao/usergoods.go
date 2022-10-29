package dao

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserGoods struct {
	ID                    string    `gorm:"column:id;primary_key"`
	Time                  time.Time `gorm:"column:time;default:CURRENT_TIMESTAMP;NOT NULL"`
	UserID                string    `gorm:"column:user_id;NOT NULL"`                           // 用户id
	GoodsID               string    `gorm:"column:goods_id;NOT NULL"`                          // 藏品id
	GoodsType             int       `gorm:"column:goods_type;default:0;NOT NULL"`              // 藏品类型 0-普通藏品 1-盲盒藏品
	GoodsOptionID         string    `gorm:"column:goods_option_id"`                            // 藏品编号id
	GoodsNo               string    `gorm:"column:goods_no"`                                   // 藏品编号名称
	GoodsHeadPic          string    `gorm:"column:goods_head_pic"`                             // 商品图像
	PurchaseType          int       `gorm:"column:purchase_type;default:0;NOT NULL"`           // 购买类型 0-购买 1-兑换
	Price                 float64   `gorm:"column:price;default:0.00;NOT NULL"`                // 购买价格
	ChainHash             string    `gorm:"column:chain_hash"`                                 // 区块链交易hash
	CertificationDate     string    `gorm:"column:certification_date"`                         // 区块链认证时间
	Version               int       `gorm:"column:version;default:0;NOT NULL"`                 // 数据版本 目前请填1
	NumStatFlag           int       `gorm:"column:num_stat_flag;default:1;NOT NULL"`           // 购买数量统计标记 手动添加数据请填0
	MysteryBoxOpenTime    time.Time `gorm:"column:mystery_box_open_time"`                      // 盲盒开启时间
	MysteryBoxOpened      int       `gorm:"column:mystery_box_opened;default:0;NOT NULL"`      // 盲盒是否已经开启
	MysteryBoxGoodsID     string    `gorm:"column:mystery_box_goods_id"`                       // 盲盒对应的商品id，购买时即确定
	MysteryBoxUserGoodsID string    `gorm:"column:mystery_box_user_goods_id"`                  // 盲盒开启后生成的user_goods表的id
	MysteryBoxOpenNotify  int       `gorm:"column:mystery_box_open_notify;default:0;NOT NULL"` // 盲盒开启提醒标志 0未提醒 1-已提醒
}

func (m *UserGoods) TableName() string {
	return "user_goods"
}

type IUserGoodsDao interface {
	dbo.DataAccesser

	PageUserGoodsByUserId(ctx context.Context, page, pageSize int, userId string) ([]*PageUserGoods, error)
}

type UserGoodsDao struct {
	dbo.BaseDA
}

var (
	_UserGoodsOnce sync.Once
	_UserGoodsDao  IUserGoodsDao
)

func GetUserGoodsDao() IUserGoodsDao {
	_UserGoodsOnce.Do(func() {
		_UserGoodsDao = &UserGoodsDao{}
	})
	return _UserGoodsDao
}

type PageUserGoods struct {
	Id                   string
	No                   string
	HeadPic              string `gorm:"column:head_pic"`
	Type                 int
	Price                string
	PurchaseType         int
	ChainHash            string
	Name                 string
	Author               string
	AuthorHeadPic        string `gorm:"column:author_Head_Pic"`
	ShopId               string
	ChainContractAddress string `gorm:"column:chain_contract_address"`
	CollectNum           int    `gorm:"column:collect_num"`
}

func (c *UserGoodsDao) PageUserGoodsByUserId(ctx context.Context, page, pageSize int, userId string) ([]*PageUserGoods, error) {
	queryPage, queryPageSize := 0, 10
	if page > 0 {
		queryPage = page
	}
	if pageSize > 0 {
		if pageSize < 10 {
			queryPageSize = pageSize
		}
	}

	var result []*PageUserGoods
	selectSql := `
	user_goods.id as id,
	user_goods.goods_no       as no,
	user_goods.goods_head_pic as head_pic,
	user_goods.goods_type     as type,
	user_goods.price as price,
	user_goods.purchase_type as purchase_type,
	user_goods.chain_hash as chain_hash,
	goods.name  as name, 
	shop.name           as author,
	shop.head_pic       as author_Head_Pic,
goods.shop_id   	  as shop_id,
goods.chain_contract_address as chain_contract_address,
goods.collect_num 	  as collect_num
`
	queryPage = queryPage - 1
	query := fmt.Sprintf(`select %s from 
              %s,
              %s,
              %s
              where user_goods.goods_id=goods.id and goods.shop_id = shop.id and (user_goods.goods_type = 0 or (user_goods.goods_type = 1 and user_goods.mystery_box_opened = 0))
              and user_goods.user_id = ? order by user_goods.id desc limit ? offset ?`, selectSql, TableNameUserGoods, TableNameGoods, TableNameShop)
	if err := c.QueryRawSQL(ctx, &result, query, userId, queryPageSize, queryPage*queryPageSize); err != nil {
		return nil, err
	}

	return result, nil
}

type UserGoodsCondition struct {
	ID                   sql.NullString
	UserID               sql.NullString
	GoodsID              sql.NullString
	NumStatFlag          sql.NullInt64
	GoodsType            sql.NullInt64
	MysteryBoxOpened     sql.NullInt64
	MysteryBoxOpenNotify sql.NullInt64
	OrderBy              int
	Pager                dbo.Pager
}

func (cnd UserGoodsCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if cnd.UserID.Valid {
		wheres = append(wheres, "user_id = ?")
		params = append(params, cnd.UserID.String)
	}

	if cnd.GoodsID.Valid {
		wheres = append(wheres, "goods_id = ?")
		params = append(params, cnd.GoodsID.String)
	}
	if cnd.GoodsType.Valid {
		wheres = append(wheres, "goods_type=?")
		params = append(params, cnd.GoodsType.Int64)
	}

	if cnd.MysteryBoxOpened.Valid {
		wheres = append(wheres, "mystery_box_opened=?")
		params = append(params, cnd.MysteryBoxOpened.Int64)
	}
	if cnd.MysteryBoxOpenNotify.Valid {
		wheres = append(wheres, "mystery_box_open_notify=?")
		params = append(params, cnd.MysteryBoxOpenNotify.Int64)
	}

	return wheres, params
}

func (c UserGoodsCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserGoodsCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
