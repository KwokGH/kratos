package dao

import (
	"context"
	"database/sql"
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserCollection struct {
	ID     string `gorm:"column:id;primary_key"`
	UserID string `gorm:"column:user_id;NOT NULL"`
	Type   int    `gorm:"column:type;default:0;NOT NULL"` // 类型：0-藏品 1-店铺 2-活动
	ItemID string `gorm:"column:item_id;NOT NULL"`        // 藏品id、店铺id或活动id
}

func (m *UserCollection) TableName() string {
	return "user_collection"
}

const (
	UserCollectionTypeGoods    = 0
	UserCollectionTypeShop     = 1
	UserCollectionTypeActivity = 2
)

type IUserCollectionDao interface {
	dbo.DataAccesser
	DeleteByCondition(ctx context.Context, userId string, collectionType int, itemId string) error
}

type UserCollectionDao struct {
	dbo.BaseDA
}

var (
	_UserCollectionOnce sync.Once
	_UserCollectionDao  IUserCollectionDao
)

func GetUserCollectionDao() IUserCollectionDao {
	_UserCollectionOnce.Do(func() {
		_UserCollectionDao = &UserCollectionDao{}
	})
	return _UserCollectionDao
}

func (c *UserCollectionDao) DeleteByCondition(ctx context.Context, userId string, collectionType int, itemId string) error {
	tx := dbo.MustGetDB(ctx)
	tx.ResetCondition()

	err := tx.Where("user_id = ? and type =? and item_id = ?", userId, collectionType, itemId).Delete(&UserCollection{}).Error

	return err
}

type UserCollectionCondition struct {
	UserID  sql.NullString
	Type    sql.NullInt32
	ItemID  sql.NullString
	ItemIDs dbo.NullStrings

	OrderBy int
	Pager   dbo.Pager
}

func (c UserCollectionCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.UserID.Valid {
		wheres = append(wheres, "user_id =?")
		params = append(params, c.UserID.String)
	}

	if c.Type.Valid {
		wheres = append(wheres, "type =?")
		params = append(params, c.Type.Int32)
	}

	if c.ItemID.Valid {
		wheres = append(wheres, "item_id =?")
		params = append(params, c.ItemID.String)
	}

	if c.ItemIDs.Valid {
		wheres = append(wheres, "item_id in (?)")
		params = append(params, c.ItemIDs.Strings)
	}

	return wheres, params
}

func (c UserCollectionCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserCollectionCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
