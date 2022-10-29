package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Goods struct {
	ID                   string            `gorm:"column:id;primary_key"`
	Name                 string            `gorm:"column:name;NOT NULL"`                           // 名称
	No                   string            `gorm:"column:no"`                                      // 藏品编号 数据版本大于等于1有效
	HeadPic              string            `gorm:"column:head_pic"`                                // 列表页图片地址
	Time                 time.Time         `gorm:"column:time;default:CURRENT_TIMESTAMP;NOT NULL"` // 上市时间，默认按改时间反向排序
	Type                 int               `gorm:"column:type;default:0;NOT NULL"`                 // 藏品类型 0-普通藏品 1-盲盒藏品
	PurchaseType         GoodsPurchaseType `gorm:"column:purchase_type;default:0;NOT NULL"`        // 购买类型 0-只能购买 1-只能兑换 2-即可购买又可兑换
	Price                float64           `gorm:"column:price;default:0.00;NOT NULL"`             // 列表页展示价格
	PurchaseLimit        int               `gorm:"column:purchase_limit;default:0;NOT NULL"`       // 限购数量
	ShopID               string            `gorm:"column:shop_id"`                                 // 所属店铺ID
	Pics                 string            `gorm:"column:pics"`                                    // 藏品主图片，多个用引文逗号隔开
	IntroduceItems       string            `gorm:"column:introduce_items"`                         // 藏品介绍项目id，多个用英文逗号隔开
	PurchaseNoticeItems  string            `gorm:"column:purchase_notice_items"`                   // 购买须知项目id，多个用英文逗号隔开
	Status               GoodsStatus       `gorm:"column:status;default:0;NOT NULL"`               // 状态 0-待上链 1-即将发售（等待开放登记） 2-在售（公布中） 3-售罄 4-已下架 5-开放登记中 6-开放登记结束待公布
	SaleTime             time.Time         `gorm:"column:sale_time"`                               // 发售时间
	CollectNum           int               `gorm:"column:collect_num;default:0;NOT NULL"`          // 收藏次数
	ViewNum              int               `gorm:"column:view_num;default:0"`                      // 查看次数
	Labels               string            `gorm:"column:labels"`                                  // 藏品标签，多个用英文逗号隔开，用于同类藏品的检索
	ChainID              string            `gorm:"column:chain_id;default:bsc;NOT NULL"`           // 区块链id，bsc或mychain，默认bsc
	ChainName            string            `gorm:"column:chain_name"`                              // 链上藏品名称，不能是中文
	ChainStatus          int               `gorm:"column:chain_status;default:0;NOT NULL"`         // 区块链状态 0-未上链 1-上链中 2-上链成功 3-上链失败
	ChainTime            time.Time         `gorm:"column:chain_time"`                              // 提交上链时间
	ChainContractAddress string            `gorm:"column:chain_contract_address"`                  // 区块链合约地址
	ChainHash            string            `gorm:"column:chain_hash"`                              // 区块链交易hash
	CertificationDate    string            `gorm:"column:certification_date"`                      // 区块链认证时间
	QrCode               string            `gorm:"column:qr_code"`                                 // 小程序码
	ShareSubTitle        string            `gorm:"column:share_sub_title"`                         // 分享海报子标题
	PurchaseScene        string            `gorm:"column:purchase_scene"`                          // 购买场景值
	SaleQty              int               `gorm:"column:sale_qty;default:0;NOT NULL"`             // 销量
	LastSaleTime         time.Time         `gorm:"column:last_sale_time"`                          // 最近销售时间
	AutoGenNo            int               `gorm:"column:auto_gen_no;default:0;NOT NULL"`          // 自定生成编号， 0-否 1-是
	VideoCoverPic        string            `gorm:"column:video_cover_pic"`                         // 视频默认图片
	Recommend            int               `gorm:"column:recommend;default:0;NOT NULL"`            // 推荐值，0不推荐，大于0为推荐商品，值越大越靠前
	WillSoldOut          int               `gorm:"column:will_sold_out;default:0;NOT NULL"`        // 是否即将售罄 0-否 1-是 （系统自动修正）
	RemainQty            int               `gorm:"column:remain_qty;default:0;NOT NULL"`           // 剩余可销售数量，系统自动统计
	Hidden               int               `gorm:"column:hidden;default:0;NOT NULL"`               // 是否隐藏，隐藏商品不在列表页显示
	Qty                  int               `gorm:"column:qty;default:0;NOT NULL"`                  // 总数量 数据版本大于等于1有效
	LockQty              int               `gorm:"column:lock_qty;default:0;NOT NULL"`             // 锁定数量 数据版本大于等于1有效
	SoldQty              int               `gorm:"column:sold_qty;default:0;NOT NULL"`             // 销售数量 数据版本大于等于1有效
	MysteryBoxGoodsIds   string            `gorm:"column:mystery_box_goods_ids"`                   // 盲盒关联的藏品id列表，多个id用英文逗号隔开，范围id使用英文减号号连接起始id和结束id
	MysteryBoxOpenTime   time.Time         `gorm:"column:mystery_box_open_time"`                   // 盲盒开启时间 如果为null则购买后立即开启
	Version              int               `gorm:"column:version;NOT NULL"`                        // 数据版本 目前请填1
	NeedBuyRight         int               `gorm:"column:need_buy_right;default:0;NOT NULL"`       // 是否需要抽取购买权 0-不需要 1-需要
	OpenStartTime        time.Time         `gorm:"column:open_start_time"`                         // 开放登记开始时间，该时间到达后自动将状态1改为5
	OpenEndTime          time.Time         `gorm:"column:open_end_time"`                           // 开放登记结束时间，该时间到达后自动将状态5改为6
	PublicTime           time.Time         `gorm:"column:public_time"`                             // 公布时间，该时间到达后自动将状态6改为2
	PaymentTime          time.Time         `gorm:"column:payment_time"`                            // 最后支付时间
	BallotFlag           int               `gorm:"column:ballot_flag;default:0;NOT NULL"`          // 需要购买权的藏品的抽签标志 0未抽签 1已抽签未通知 2已通知
}

func (m *Goods) TableName() string {
	return "goods"
}

type GoodsStatus int

const (
	GoodsStatusPending    GoodsStatus = 0
	GoodsStatusUpcoming   GoodsStatus = 1
	GoodsStatusOnSale     GoodsStatus = 2
	GoodsStatusSoldOut    GoodsStatus = 3
	GoodsStatusOffline    GoodsStatus = 4
	GoodsStatusOpening    GoodsStatus = 5
	GoodsStatusOpenFinish GoodsStatus = 6
)

type GoodsPurchaseType int

const (
	GoodsPurchaseTypePurchase         GoodsPurchaseType = 0
	GoodsPurchaseTypeCDkey            GoodsPurchaseType = 1
	GoodsPurchaseTypePurchaseAndCDkey GoodsPurchaseType = 2
)

type IGoodsDao interface {
	dbo.DataAccesser

	GetByIdsWithShopInfo(ctx context.Context, ids dbo.NullStrings) ([]*GoodsWithShop, error)
	UpdateViewNum(ctx context.Context, shopID string) error
	UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, goodsID string, isCollect bool) error

	LockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error
	UnlockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error

	DeductQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error
	IncreaseSoldQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error
}

type GoodsDao struct {
	dbo.BaseDA
}

var (
	_GoodsOnce sync.Once
	_GoodsDao  IGoodsDao
)

func GetGoodsDao() IGoodsDao {
	_GoodsOnce.Do(func() {
		_GoodsDao = &GoodsDao{}
	})
	return _GoodsDao
}

type GoodsWithShop struct {
	Goods

	Author        string `gorm:"shopName"`
	AuthorHeadPic string `gorm:"shopHeadPic"`
}

func (c *GoodsDao) GetByIdsWithShopInfo(ctx context.Context, ids dbo.NullStrings) ([]*GoodsWithShop, error) {
	if !ids.Valid {
		return make([]*GoodsWithShop, 0), nil
	}

	selectFields := fmt.Sprintf(`t1.*, t2.name as shopName, t2.head_pic as shopHeadPic`)
	query := fmt.Sprintf(`
		select %s from %s as t1 inner join %s as t2 on t1.shop_id = t2.id where t1.id in (%s)`, selectFields, TableNameGoods, TableNameShop, ids.SQLPlaceHolder())

	var result = make([]*GoodsWithShop, 0)
	if err := c.QueryRawSQL(ctx, &result, query, ids.ToInterfaceSlice()...); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *GoodsDao) UpdateViewNum(ctx context.Context, shopId string) error {
	tx := dbo.MustGetDB(ctx)
	tx.ResetCondition()

	return tx.Model(&Goods{}).Where("id=?", shopId).UpdateColumn("view_num", gorm.Expr("view_num + ?", 1)).Error
}

func (c *GoodsDao) UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, goodsID string, isCollect bool) error {
	tx.ResetCondition()

	if isCollect {
		return tx.Model(&Goods{}).Where("id=?", goodsID).UpdateColumn("collect_num", gorm.Expr("collect_num + ?", 1)).Error
	} else {
		return tx.Model(&Goods{}).Where("id=? and collect_num-1 >= 0", goodsID).UpdateColumn("collect_num", gorm.Expr("collect_num - ?", 1)).Error
	}
}

func (c *GoodsDao) UnlockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error {
	tx.ResetCondition()

	return tx.Model(&Goods{}).
		Where("id=? and lock_qty-? >= 0", goodsID, qty).
		UpdateColumn("lock_qty", gorm.Expr("lock_qty - ?", qty)).
		Error
}

func (c *GoodsDao) LockQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error {
	tx.ResetCondition()

	updateTx := tx.Model(&Goods{}).
		Where("id=? and lock_qty+sold_qty+? <= qty", goodsID, qty).
		UpdateColumn("lock_qty", gorm.Expr("lock_qty + ?", qty))

	if updateTx.Error != nil {
		return updateTx.Error
	}

	if updateTx.RowsAffected <= 0 {
		return dbo.ErrRecordNotFound
	}

	return nil
}

func (c *GoodsDao) DeductQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error {
	tx.ResetCondition()

	return tx.Model(&Goods{}).
		Where("id=? and lock_qty-? >= 0", goodsID, qty).
		UpdateColumn("lock_qty", gorm.Expr("lock_qty - ?", qty)).
		UpdateColumn("sold_qty", gorm.Expr("sold_qty + ?", qty)).Error
}

func (c *GoodsDao) IncreaseSoldQtyTx(ctx context.Context, tx *dbo.DBContext, goodsID string, qty int) error {
	tx.ResetCondition()

	return tx.Model(&Goods{}).
		Where("id=?", goodsID).
		UpdateColumn("sold_qty", gorm.Expr("sold_qty + ?", qty)).Error
}

type GoodsCondition struct {
	IDs     dbo.NullStrings
	ShopID  sql.NullString
	ShopIDs dbo.NullStrings
	Type    sql.NullInt64

	UserIdAboutCollect sql.NullString
	GoodsName          sql.NullString
	Status             sql.NullInt64
	Statuses           dbo.NullInts
	SimilarGoodsCond   *SimilarGoodsCondition
	NeedBuyRight       sql.NullInt64
	BallotFlag         sql.NullInt64
	RemainQtyGT        sql.NullInt64
	MysteryBoxGoodsIds dbo.NullStrings

	Hidden bool

	OrderBy int
	Pager   dbo.Pager
}

type SimilarGoodsCondition struct {
	GoodsId string
	Labels  dbo.NullStrings
}

func (condition GoodsCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if condition.IDs.Valid {
		wheres = append(wheres, "id in (?)")
		params = append(params, condition.IDs.Strings)
	}

	if condition.ShopIDs.Valid {
		wheres = append(wheres, "shop_id in (?)")
		params = append(params, condition.ShopIDs.Strings)
	}

	if condition.Type.Valid {
		wheres = append(wheres, "type = ?")
		params = append(params, condition.Type.Int64)
	}

	if condition.SimilarGoodsCond != nil {
		wheres = append(wheres, "id != ?")
		params = append(params, condition.SimilarGoodsCond.GoodsId)

		orWheres := make([]string, 0)
		for _, label := range condition.SimilarGoodsCond.Labels.Strings {
			orWheres = append(orWheres, "labels like ?")
			params = append(params, "%"+label+"%")
		}
		wheres = append(wheres, strings.Join(orWheres, " or "))
	}

	if condition.GoodsName.Valid {
		wheres = append(wheres, "name like ?")
		params = append(params, "%"+condition.GoodsName.String+"%")
	}

	if condition.Status.Valid {
		wheres = append(wheres, "status = ?")
		params = append(params, condition.Status.Int64)
	}
	if condition.Statuses.Valid {
		wheres = append(wheres, "status in (?)")
		params = append(params, condition.Statuses.Ints)
	}

	if condition.UserIdAboutCollect.Valid {
		wheres = append(wheres, fmt.Sprintf(`
exists(select 1 from %s where user_id= ? and type= %d and %s.id = %s.item_id)`,
			TableNameUserCollection, 0, TableNameGoods, TableNameUserCollection))
		params = append(params, condition.UserIdAboutCollect.String)
	}

	if condition.NeedBuyRight.Valid {
		wheres = append(wheres, "need_buy_right=?")
		params = append(params, condition.NeedBuyRight.Int64)
	}

	if condition.BallotFlag.Valid {
		wheres = append(wheres, "ballot_flag=?")
		params = append(params, condition.BallotFlag.Int64)
	}

	if condition.MysteryBoxGoodsIds.Valid {
		ors := make([]string, 0)
		for _, idsStr := range condition.MysteryBoxGoodsIds.Strings {
			if idsStr == "" {
				continue
			}
			ids := strings.Split(idsStr, "-")
			if len(ids) == 1 {
				ors = append(ors, " id = ? ")
				params = append(params, ids[0])
			} else if len(ids) > 1 {
				ors = append(ors, " id  between ? and ? ")
				params = append(params, ids[0], ids[1])
			}
		}

		if len(ors) > 0 {
			wheres = append(wheres, "("+strings.Join(ors, " or ")+")")
		}
	}

	if condition.RemainQtyGT.Valid {
		wheres = append(wheres, "remain_qty > ?")
		params = append(params, condition.RemainQtyGT.Int64)
	}

	if condition.Hidden {
		wheres = append(wheres, "hidden = 1")
	} else {
		wheres = append(wheres, "hidden = 0")
	}

	return wheres, params
}

func (c GoodsCondition) GetOrderBy() string {
	switch c.OrderBy {
	case 0:
		return "time desc,id desc"
	case 1:
		return "collect_num desc,id desc"
	case 2:
		return "view_num desc,id desc"
	case 3:
		return "price asc,id asc"
	case 4:
		return "price desc,id desc"
	case 5:
		return "time asc,id asc"
	default:
		return "id desc"
	}
}

func (c GoodsCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
