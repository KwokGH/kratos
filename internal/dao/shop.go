package dao

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
	"gorm.io/gorm"
)

type Shop struct {
	ID              int64     `gorm:"column:id;primary_key"`
	Name            string    `gorm:"column:name;NOT NULL"` // 店铺名称
	HeadPic         string    `gorm:"column:head_pic"`
	DetailItems     string    `gorm:"column:detail_items"`                         // 详情项目id，多个用英文逗号隔开，可以重复
	CollectNum      int       `gorm:"column:collect_num;default:0;NOT NULL"`       // 收藏次数
	ViewNum         int       `gorm:"column:view_num;default:0;NOT NULL"`          // 查看次数
	QrCode          string    `gorm:"column:qr_code"`                              // 小程序码
	ShareSubTitle   string    `gorm:"column:share_sub_title"`                      // 分享海报子标题
	SaleQty         int       `gorm:"column:sale_qty;default:0;NOT NULL"`          // 销量
	LastSaleTime    time.Time `gorm:"column:last_sale_time"`                       // 最近销售时间
	VCertification  int       `gorm:"column:v_certification;default:0;NOT NULL"`   // 是否V认证 0否，1是
	GoodsNum        int       `gorm:"column:goods_num;default:0;NOT NULL"`         // 商品数（系统自动统计）
	GoodsCollectNum int       `gorm:"column:goods_collect_num;default:0;NOT NULL"` // 商品总收藏数（系统自动统计）
	BackgroundImg   string    `gorm:"column:background_img"`
	Seq             int       `gorm:"column:seq;default:0;NOT NULL"` // 默认顺序值，越大越靠前，必须小于10000
}

func (m *Shop) TableName() string {
	return "shop"
}

type IShopDao interface {
	dbo.DataAccesser

	GetShopStat(ctx context.Context) ([]*ShopStatData, error)
	UpdateShopViewNum(ctx context.Context, shopId string) error
	UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, shopID string, isCollect bool) error
}

type ShopStatData struct {
	ShopID          string `gorm:"shop_id"`
	GoodsNum        int    `gorm:"goods_num"`
	GoodsCollectNum int    `gorm:"goods_collect_num"`
}

type ShopDao struct {
	dbo.BaseDA
}

var (
	_ShopOnce sync.Once
	_ShopDao  IShopDao
)

func GetShopDao() IShopDao {
	_ShopOnce.Do(func() {
		_ShopDao = &ShopDao{}
	})
	return _ShopDao
}

func (c *ShopDao) UpdateShopViewNum(ctx context.Context, shopId string) error {
	tx := dbo.MustGetDB(ctx)
	tx.ResetCondition()

	return tx.Model(&Shop{}).Where("id=?", shopId).UpdateColumn("view_num", gorm.Expr("view_num + ?", 1)).Error
}

func (c *ShopDao) UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, shopID string, isCollect bool) error {
	tx.ResetCondition()

	if isCollect {
		return tx.Model(&Shop{}).Where("id=?", shopID).UpdateColumn("collect_num", gorm.Expr("collect_num + ?", 1)).Error
	} else {
		return tx.Model(&Shop{}).Where("id=? and collect_num > 0", shopID).UpdateColumn("collect_num", gorm.Expr("collect_num - ?", 1)).Error
	}
}

func (c *ShopDao) GetShopStat(ctx context.Context) ([]*ShopStatData, error) {
	query := fmt.Sprintf(`SELECT 
	shop_id,sum(collect_num) goods_collect_num,count(*) goods_num  
	from %s 
	where shop_id is not null group by shop_id`, new(Goods).TableName())

	var result []*ShopStatData
	if err := c.QueryRawSQL(ctx, &result, query); err != nil {
		return nil, err
	}

	return result, nil
}

type ShopCondition struct {
	FuzzyName          sql.NullString
	UserIdAboutCollect sql.NullString
	IDs                dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c ShopCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.FuzzyName.Valid {
		wheres = append(wheres, "name like %?%")
		params = append(params, c.FuzzyName.String)
	}
	if c.UserIdAboutCollect.Valid {
		wheres = append(wheres, fmt.Sprintf(`
exists(select 1 from %s where user_id= ? and type= %d and %s.id = %s.item_id)`,
			TableNameUserCollection, 1, TableNameShop, TableNameUserCollection))
		params = append(params, c.UserIdAboutCollect.String)
	}

	if c.IDs.Valid {
		wheres = append(wheres, "id in (?)")
		params = append(params, c.IDs.Strings)
	}

	return wheres, params
}

func (c ShopCondition) GetOrderBy() string {
	var orderSql string
	switch c.OrderBy {
	case 0:
		orderSql = "seq desc"
	case 1:
		orderSql = "sale_qty desc"
	case 2:
		orderSql = "view_num desc"
	default:
		orderSql = "seq desc"
	}

	return orderSql + ",id desc"
}

func (c ShopCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
