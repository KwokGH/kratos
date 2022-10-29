package dao

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/KwokGH/kratos/pkg/dbo"
	"sync"
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID                  string    `gorm:"column:id;primary_key"`
	Title               string    `gorm:"column:title;NOT NULL"`                 // 标题
	ClassifyID          string    `gorm:"column:classify_id"`                    // 分类ID
	ShopID              string    `gorm:"column:shop_id"`                        // 店铺id
	HeadPic             string    `gorm:"column:head_pic"`                       // 列表图片地址
	Pics                string    `gorm:"column:pics"`                           // 顶部图片地址，多个用英文逗号隔开
	ShowTime            string    `gorm:"column:show_time"`                      // 展览时间
	ShowAddress         string    `gorm:"column:show_address"`                   // 展览地址
	IntroduceItems      string    `gorm:"column:introduce_items"`                // 展出介绍：detail_item的id列表，多个id用英文逗号隔开
	PurchaseNoticeItems string    `gorm:"column:purchase_notice_items"`          // 购票须知项目id，多个用英文逗号隔开
	CollectNum          int       `gorm:"column:collect_num;default:0;NOT NULL"` // 收藏数量
	Seq                 int       `gorm:"column:seq;default:0;NOT NULL"`         // 默认顺序值，越大越靠前，必须小于10000
	QrCode              string    `gorm:"column:qr_code"`                        // 小程序码
	ShareSubTitle       string    `gorm:"column:share_sub_title"`                // 分享海报子标题
	Status              int       `gorm:"column:status;default:0;NOT NULL"`      // 状态 0-已下架 1-即将开始 2-正在开始 3-已结束
	StartTime           time.Time `gorm:"column:start_time"`                     // 开始时间
	EndTime             time.Time `gorm:"column:end_time"`                       // 结束时间
	SubTitle            string    `gorm:"column:sub_title"`                      // 子标题
	Recommend           int       `gorm:"column:recommend;default:0;NOT NULL"`   // 推荐值，0不推荐，大于0为推荐活动，值越大越靠前
	HomeImg             string    `gorm:"column:home_img"`                       // 首页推荐图片
}

type ActivityStatus int

const (
	ActivityStatusOffline  ActivityStatus = 0
	ActivityStatusUpComing ActivityStatus = 1
	ActivityStatusGoing    ActivityStatus = 2
	ActivityStatusFinish   ActivityStatus = 3
)

func (m *Activity) TableName() string {
	return "activity"
}

// 操作数据库相关接口
type IActivityDao interface {
	dbo.DataAccesser

	GetByIdsWithShopInfo(ctx context.Context, ids []string) ([]*ActivityWithShopInfo, error)
	GetRecommend(ctx context.Context) ([]*ActivityWithShopInfo, error)
	UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, activityID string, isCollect bool) error
}

type ActivityDao struct {
	dbo.BaseDA
}

var (
	_activityOnce sync.Once
	_activityDao  IActivityDao
)

func GetActivityDao() IActivityDao {
	_activityOnce.Do(func() {
		_activityDao = &ActivityDao{}
	})
	return _activityDao
}

func (c *ActivityDao) GetRecommend(ctx context.Context) ([]*ActivityWithShopInfo, error) {
	selectFields := `t1.*, t2.name author, t2.head_pic authorHeadPic`
	query := fmt.Sprintf(`
select %s from %s as t1 inner join %s as t2 on t1.shop_id = t2.id
where t1.recommend > 0 order by t1.recommend desc,t1.id desc 
`, selectFields, TableNameActivity, TableNameShop)

	var result []*ActivityWithShopInfo
	if err := c.QueryRawSQL(ctx, &result, query); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ActivityDao) GetByIdsWithShopInfo(ctx context.Context, ids []string) ([]*ActivityWithShopInfo, error) {
	selectFields := `t1.*, t2.name author, t2.head_pic author_head_pic`
	query := fmt.Sprintf(`
	select %s from %s as t1 inner join %s as t2 on t1.shop_id = t2.id where t1.id in (?)
`, selectFields, TableNameActivity, TableNameShop)
	var result []*ActivityWithShopInfo
	if err := c.QueryRawSQL(ctx, &result, query, ids); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *ActivityDao) UpdateCollectNumTx(ctx context.Context, tx *dbo.DBContext, activityID string, isCollect bool) error {
	tx.ResetCondition()

	if isCollect {
		return tx.Model(&Activity{}).Where("id=?", activityID).UpdateColumn("collect_num", gorm.Expr("collect_num + ?", 1)).Error
	} else {
		return tx.Model(&Activity{}).Where("id=? and collect_num > 0", activityID).UpdateColumn("collect_num", gorm.Expr("collect_num - ?", 1)).Error
	}
}

type ActivityCondition struct {
	IDs                dbo.NullStrings
	ShopIds            dbo.NullStrings
	Status             sql.NullInt64
	StatusList         dbo.NullInts
	UserIdAboutCollect sql.NullString
	Keyword            sql.NullString

	OrderBy int
	Pager   dbo.Pager
}

func (c ActivityCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.ShopIds.Valid {
		wheres = append(wheres, fmt.Sprintf("shop_id in (%s)", c.ShopIds.SQLPlaceHolder()))
		params = append(params, c.ShopIds.ToInterfaceSlice()...)
	}

	if c.Status.Valid {
		wheres = append(wheres, "status = ?")
		params = append(params, c.Status.Int64)
	}

	if c.StatusList.Valid {
		wheres = append(wheres, "status in (?)")
		params = append(params, c.StatusList.Ints)
	}

	if c.Keyword.Valid {
		wheres = append(wheres, "title like ?")
		params = append(params, "%"+c.Keyword.String+"%")
	}

	if c.UserIdAboutCollect.Valid {
		wheres = append(wheres, fmt.Sprintf(`
exists(select 1 from %s where user_id= ? and type= %d and %s.id = %s.item_id)`,
			TableNameUserCollection, 2, TableNameActivity, TableNameUserCollection))
		params = append(params, c.UserIdAboutCollect.String)
	}

	return wheres, params
}

func (c ActivityCondition) GetOrderBy() string {
	orderBy := ""
	switch c.OrderBy {
	case 1:
		orderBy = " id desc "
	case 2:
		orderBy = " id asc "
	default:
		orderBy = " id desc "
	}

	return orderBy
}

func (c ActivityCondition) GetPager() *dbo.Pager {
	return &c.Pager
}

// db view

type ActivityWithShopInfo struct {
	Activity
	Author        string `gorm:"author"`
	AuthorHeadPic string `gorm:"authorHeadPic"`
}
