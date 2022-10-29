package dao

import (
	"database/sql"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type Message struct {
	ID      string    `gorm:"column:id;primary_key"`
	Time    time.Time `gorm:"column:time;NOT NULL"` // 产生时间
	UserID  string    `gorm:"column:user_id;NOT NULL"`
	Type    int       `gorm:"column:type;default:0;NOT NULL"` // 类型 0-系统消息
	Title   string    `gorm:"column:title;NOT NULL"`          // 标题
	Content string    `gorm:"column:content"`                 // 内容
	Read    int       `gorm:"column:read;default:0;NOT NULL"` // 是否已读 0-未读 1-已读
}

func (m *Message) TableName() string {
	return "message"
}

type IMessageDao interface {
	dbo.DataAccesser
}

type MessageDao struct {
	dbo.BaseDA
}

var (
	_MessageOnce sync.Once
	_MessageDao  IMessageDao
)

func GetMessageDao() IMessageDao {
	_MessageOnce.Do(func() {
		_MessageDao = &MessageDao{}
	})
	return _MessageDao
}

type MessageCondition struct {
	UserID sql.NullString

	OrderBy int
	Pager   dbo.Pager
}

func (c MessageCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	if c.UserID.Valid {
		wheres = append(wheres, "user_id = ?")
		params = append(params, c.UserID.String)
	}

	return wheres, params
}

func (c MessageCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return "id desc"
	}
}

func (c MessageCondition) GetPager() *dbo.Pager {
	return &c.Pager
}
