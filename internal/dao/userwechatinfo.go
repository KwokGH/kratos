package dao

import (
	"sync"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type UserWechatInfo struct {
	ID        string `gorm:"column:id;primary_key"`
	UserID    string `gorm:"column:user_id;NOT NULL"`
	NickName  string `gorm:"column:nick_name"`
	AvatarUrl string `gorm:"column:avatar_url"`
	Gender    string `gorm:"column:gender"`
	Province  string `gorm:"column:province"`
	City      string `gorm:"column:city"`
	Country   string `gorm:"column:country"`
}

func (m *UserWechatInfo) TableName() string {
	return "user_wechat_info"
}

type IUserWechatInfoDao interface {
	dbo.DataAccesser
}

type UserWechatInfoDao struct {
	dbo.BaseDA
}

var (
	_UserWechatInfoOnce sync.Once
	_UserWechatInfoDao  IUserWechatInfoDao
)

func GetUserWechatInfoDao() IUserWechatInfoDao {
	_UserWechatInfoOnce.Do(func() {
		_UserWechatInfoDao = &UserWechatInfoDao{}
	})
	return _UserWechatInfoDao
}

type UserWechatInfoCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c UserWechatInfoCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c UserWechatInfoCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c UserWechatInfoCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
