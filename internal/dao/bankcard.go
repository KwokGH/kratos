package dao

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/KwokGH/kratos/pkg/dbo"
)

type BankCard struct {
	ID           string    `gorm:"column:id;primary_key"`
	Time         time.Time `gorm:"column:time"`                      // 录入时间
	UserID       string    `gorm:"column:user_id;NOT NULL"`          // 所属用户id
	BankID       string    `gorm:"column:bank_id"`                   // 银行id
	Name         string    `gorm:"column:name"`                      // 账户名
	Account      string    `gorm:"column:account"`                   // 账号
	IDCard       string    `gorm:"column:id_card"`                   // 身份证号
	Mobile       string    `gorm:"column:mobile"`                    // 预留手机号
	WtsTradeNo   string    `gorm:"column:wts_trade_no"`              // 通联流水号
	WtsTransDate string    `gorm:"column:wts_trans_date"`            // 通联交易日期
	Status       int       `gorm:"column:status;default:0;NOT NULL"` // 状态 0-待上送验证码 1-绑定成功
}

func (m *BankCard) TableName() string {
	return "bank_card"
}

type IBankCardDao interface {
	dbo.DataAccesser
	FirstUserBankCard(ctx context.Context, userId string) (*UserBankCardDB, error)
}

type BankCardDao struct {
	dbo.BaseDA
}

var (
	_BankCardOnce sync.Once
	_BankCardDao  IBankCardDao
)

func GetBankCardDao() IBankCardDao {
	_BankCardOnce.Do(func() {
		_BankCardDao = &BankCardDao{}
	})
	return _BankCardDao
}

type BankCardCondition struct {
	OrderBy int
	Pager   dbo.Pager
}

func (c BankCardCondition) GetConditions() ([]string, []interface{}) {
	var wheres []string
	var params []interface{}

	return wheres, params
}

func (c BankCardCondition) GetOrderBy() string {
	switch c.OrderBy {
	default:
		return ""
	}
}

func (c BankCardCondition) GetPager() *dbo.Pager {
	return &c.Pager
}

type UserBankCardDB struct {
	BankName string `gorm:"bank_name"`
	BankLogo string `gorm:"bank_logo"`
	Account  string `gorm:"account"`
	Limit    int    `gorm:"limit"`
}

func (c *BankCardDao) FirstUserBankCard(ctx context.Context, userId string) (*UserBankCardDB, error) {
	var result []*UserBankCardDB
	selectSql := "bank.name bank_name, bank.logo bank_logo, bank_card.account, bank.limit"
	query := fmt.Sprintf(`
select %s from %s as bank_card
inner join %s as bank on bank_card.bank_id=bank.id
where bank_card.user_id = ? order by bank_card.id desc limit 1
`, selectSql, TableNameBankCard, TableNameBank)

	err := c.QueryRawSQL(ctx, &result, query, userId)
	if err != nil {
		return nil, err
	}

	if len(result) <= 0 {
		return nil, dbo.ErrRecordNotFound
	}

	return result[0], nil
}
